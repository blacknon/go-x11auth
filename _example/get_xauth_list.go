package main

import (
	"fmt"
	"os"

	x11auth "github.com/blacknon/go-x11auth"
)

func main() {
	xauth := x11auth.XAuth{}

	xAuthorityFilePath = os.Getenv("XAUTHORITY")
	entries, err := xauth.GetXAuthList(xAuthorityFilePath)

	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(entries)
}
