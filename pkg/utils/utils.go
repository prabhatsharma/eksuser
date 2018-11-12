package utils

import (
	"fmt"
	"os"

	yaml "gopkg.in/yaml.v2"
)

type IamUser struct {
	UserArn  string   `yaml:"userarn"`
	UserName string   `yaml:"username"`
	Groups   []string `yaml:"groups"`
}

type MapUsers struct {
	Users []IamUser `yaml:"mapUsers"`
}

// HomeDir gets the home directory of the user
func HomeDir() string {
	if h := os.Getenv("HOME"); h != "" {
		return h
	}
	return os.Getenv("USERPROFILE") // windows
}

// RemoveIfAlreadyExists removes already existing user so that we can add it again after making required modifications
func RemoveIfAlreadyExists(allUsers []IamUser, userToRemove string) []IamUser {
	userLocation := FindExistingUser(allUsers, userToRemove)

	if userLocation >= 0 {
		allUsers[userLocation] = allUsers[len(allUsers)-1]
		allUsers = allUsers[:len(allUsers)-1]
	}

	return allUsers
}

// FindExistingUser finds an existing user in the configmap mapUsers array and returns its localtions
func FindExistingUser(allUsers []IamUser, userToFind string) int {
	for i, user := range allUsers {
		if user.UserName == userToFind {
			return i
		}
	}

	return -1
}

// ConvertUsersStringListToStruct : since user details are stored as string in the configmap utils.MapUsers we need to convert it to struct before we can play (add/delete/modify) with it 
func ConvertUsersStringListToStruct(userList string) MapUsers {
	// add utils.MapUsers to the top so that mapping is easier.
	userList = "mapUsers: \n" + userList

	ebUsers := []byte(userList)

	var mappedUsers MapUsers

	// convert string representation of mappedUsers to struct so that we can manipulate it
	errU := yaml.Unmarshal(ebUsers, &mappedUsers)
	if errU != nil {
		fmt.Println(errU)
	}

	return mappedUsers
}
