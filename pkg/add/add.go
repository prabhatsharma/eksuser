package add

import (
	"fmt"
	"os"
	"strings"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/awserr"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/iam"
	"github.com/prabhatsharma/eksuser/pkg/action"
	"github.com/prabhatsharma/eksuser/pkg/utils"
)

// IamUser is the format of the user return by GetGroups
type IamUser struct {
	Arn              string `json:"arn"`
	CreateDate       string `json:"createddate"`
	PasswordLastUsed string `json:"passwordlastused"`
	Path             string `json:"path"`
	UserID           string `json:"userid"`
	UserName         string `json:"username"`
}

// InsertUser adds the user to the EKS. It makes an entry of the user to aws-
func InsertUser(userName, iamgroup, kubegroups string) {
	userArn := getUserARN(userName)
	insertUserInKube(userName, iamgroup, userArn, kubegroups)
}

func insertUserInKube(userName, iamgroup, userArn, kubegroups string) {
	newUser := utils.IamUser{
		UserArn:  userArn,
		UserName: userName,
		IAMGroup: iamgroup,
		Groups:   strings.Split(kubegroups, ","), // get one or more groups that were passed on command line
	}

	action.UpdateKubeConfigMapUser(newUser, "add")

	fmt.Println(userName, " added/updated to EKS.")

}

// getUserARN gets the ARN of the user
func getUserARN(user string) string {
	svc := iam.New(session.New())
	input := &iam.GetUserInput{
		UserName: aws.String(user),
	}

	result, err := svc.GetUser(input)
	if err != nil {
		if aerr, ok := err.(awserr.Error); ok {
			switch aerr.Code() {
			case iam.ErrCodeNoSuchEntityException:
				// fmt.Println(iam.ErrCodeNoSuchEntityException, aerr.Error())
				fmt.Println("IAM user: ", user, " could not be found.")
			case iam.ErrCodeServiceFailureException:
				fmt.Println(iam.ErrCodeServiceFailureException, aerr.Error())
			default:
				fmt.Println(aerr.Error())
			}
			os.Exit(1)
		}
	}

	return *result.User.Arn
}

// InsertIAMGroup adds the users in the IAM group to the EKS in the specified group
func InsertIAMGroup(iamgroup, kubegroups string) {
	svc := iam.New(session.New())

	input := &iam.GetGroupInput{
		GroupName: aws.String(iamgroup),
	}

	result, err := svc.GetGroup(input)
	if err != nil {
		fmt.Println("Error occurred getting IAM Group details")
	}

	for index, user := range result.Users {
		fmt.Println(*user.Arn, index)
		insertUserInKube(*user.UserName, iamgroup, *user.Arn, kubegroups)
	}

	fmt.Println(iamgroup, " IAM group added/updated to EKS group ", kubegroups)
}

// UpdateIAMGroup updates (updates/deletes) users of an IAM group
func UpdateIAMGroup(iamgroup, kubegroups string) {

	// Remove all group users to begin with.

	// Reinsert all group users.
	InsertIAMGroup(iamgroup, kubegroups)

}

// removeIAMGroupUsers removes all the IAM users of the particular group from the aws-auth configmap
func removeIAMGroupUsers(group string) {

}
