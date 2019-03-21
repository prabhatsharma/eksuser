package action

import (
	"fmt"
	"os"
	"path/filepath"

	yaml "gopkg.in/yaml.v2"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"

	"github.com/prabhatsharma/eksuser/pkg/utils"
)

// UpdateKubeConfigMap updates the aws-auth configmap that contains EKS users IAM mapping.
// userToActUpon: user to be added or deleted
// action: valid values for action: add, delete
func UpdateKubeConfigMap(userToActUpon utils.IamUser, action string) {

	var kubeConfigPath string
	if home := utils.HomeDir(); home != "" {
		kubeConfigPath = filepath.Join(home, ".kube", "config")
	}

	// use the current context in kubeconfig
	config, err := clientcmd.BuildConfigFromFlags("", kubeConfigPath)
	if err != nil {
		panic(err.Error())
	}

	// create the clientset
	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		panic(err.Error())
	}

	// get all the configmaps in kube-system namespace
	cm, cErr := clientset.CoreV1().ConfigMaps("kube-system").List(metav1.ListOptions{})
	if cErr != nil {
		if cErr.Error() == "Unauthorized" {
			fmt.Println("Current IAM user is not authorized to read/write kube-system/aws-auth configmap. You need access to kube-system/aws-auth configmap for you to use eksuser.")
			os.Exit(1)
		}
	}

	for _, v := range cm.Items {
		if v.Name == "aws-auth" {
			allUsers := utils.ConvertUsersStringListToStruct(v.Data["mapUsers"])

			// Remove existing user from the configmap struct so that we can update and re-insert it
			allUsers.Users = utils.RemoveIfAlreadyExists(allUsers.Users, userToActUpon.UserName)

			// add the requested user to the struct
			if action == "add" {
				allUsers.Users = append(allUsers.Users, userToActUpon)
			}

			byteAllUsers, _ := yaml.Marshal(allUsers.Users)
			strAllUsers := string(byteAllUsers)

			v.Data["mapUsers"] = strAllUsers

			// update kubernetes configmap.
			_, err5 := clientset.CoreV1().ConfigMaps("kube-system").Update(&v)
			if err5 != nil {
				fmt.Println("Error updating configmap")
			}
		}
	}
}
