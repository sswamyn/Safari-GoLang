package main

import (
	"dcs.com/mainproj/messages"
	"dcs.com/support/utils"
	"fmt"
)

func main() {
	fmt.Println("Hello Go World!")
	fmt.Println(messages.MyPkgMessage)
	fmt.Println(messages.GetMessage())
	fmt.Println(utils.ModuleMessage)
}
