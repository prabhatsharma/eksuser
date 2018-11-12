package del

import (
	"fmt"

	"github.com/prabhatsharma/eksuser/pkg/action"
	"github.com/prabhatsharma/eksuser/pkg/utils"
)

func DeleteUser(userName string) {

	userToDelete := utils.IamUser{
		UserArn:  "", // userArn is not required as we are not calling AWS IAM
		UserName: userName,
		Groups:   nil,
	}

	action.UpdateKubeConfigMap(userToDelete, "delete")

	fmt.Println(userName, " deleted from EKS.")
}
