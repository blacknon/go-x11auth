package main

import (
	"fmt"
	"os"

	x11auth "github.com/blacknon/go-x11auth"
)

func main() {
	xauth := x11auth.XAuth{}

	xAuthorityFilePath := os.Getenv("XAUTHORITY")
	xauth.Display = os.Getenv("DISPLAY")

	// untrusted cookie
	untrustedCookie, err := xauth.GetXAuthCookie(xAuthorityFilePath, false)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("untrusted cookie:", untrustedCookie)

	// trusted cookie
	trustedCookie, err := xauth.GetXAuthCookie(xAuthorityFilePath, true)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("trusted cookie:", trustedCookie)
}
