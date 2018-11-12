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

// InsertUser adds the user to the EKS. It makes an entry of the user to aws-
func InsertUser(userName, groups string) {

	userArn := getUserARN(userName)

	newUser := utils.IamUser{
		UserArn:  userArn,
		UserName: userName,
		Groups:   strings.Split(groups, ","), // get one or more groups that were passed on command line
	}

	action.UpdateKubeConfigMap(newUser, "add")

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
