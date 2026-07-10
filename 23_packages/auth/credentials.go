package auth

import (
	"fmt"

	"github.com/fatih/color"
)

func LoginWithCredentials(username, password string) {
	fmt.Println("user loggedIn using:", username, password)

	color.Red(isAuthorized(username)) // using the function of an outside package
}

// This function can be used only inside this package
// since the name of the function starts with a lowercase letter
// func loginWithCredentials(username, password string) {
// 	fmt.Println("user loggedIn using:", username, password)
// }
