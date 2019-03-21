package del

import (
	"fmt"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/iam"
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

// DeleteIAMGroup adds the users in the IAM group to the EKS in the specified group
func DeleteIAMGroup(iamgroup string) {
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
		DeleteUser(*user.UserName)
	}

	fmt.Println(iamgroup, " IAM group users deleted from EKS")
}
