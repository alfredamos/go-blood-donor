package utils

import "errors"

func CheckForOwnership(userIdFromContext, userIdFromResource string) error {
	//----> Is not same user.
	if isEqual := isSameUser(userIdFromContext, userIdFromResource); !isEqual {
		return errors.New("you are not allowed to view or delete this resource")
	}

	//----> Is not same user.
	return nil
}

func isSameUser(userIdFirst, userIdSecond string) bool {
	return userIdFirst == userIdSecond
}
