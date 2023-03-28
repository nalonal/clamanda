package main

import (
	"fmt"

	"github.com/mbndr/figlet4go"
)

func main() {
	messageWelcome()
}

func messageWelcome() {
	ascii := figlet4go.NewAsciiRender()
	renderStr, _ := ascii.Render("Clamanda")
	fmt.Print(renderStr)
	fmt.Println("Threat Detection - APT Scanner - IoC Scanner - Log Scanner")
	fmt.Println("Version 1.0.0 (2023-03-28 10:58)")
	fmt.Println("(c) National Cyber and Crypto Agency - Indonesia")
}
