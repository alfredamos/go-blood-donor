package utils

import "errors"

func CheckForOwnership(userIdFromContext, userIdFromResource string, isAdmin bool) error {
	//----> Get same user flag.
	isEqual := isSameUser(userIdFromContext, userIdFromResource)
	
	//----> Check for same user and admin privilege.
	if !isEqual && !isAdmin {
		return errors.New("you are not allowed to view or delete this resource")
	}

	//----> Either is same user or admin.
	return nil
}

func isSameUser(userIdFirst, userIdSecond string) bool {
	return userIdFirst == userIdSecond
}
