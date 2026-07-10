package auth

import "fmt"

func isAuthorized(username string) string {
	if username == "codedbygautam" {
		return fmt.Sprintf("User : %s is authorized to perform all actions.", username)
	}
	return fmt.Sprintf("User : %s isn't authorized beyond default permissions.", username)
}
