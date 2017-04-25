package main

import (
	"fmt"

	"github.com/fouralarmfire/xsay"
)

func main() {
	dino := xsay.New("stego.txt", "RARW!")

	fmt.Print("\n")
	dino.Say()
	fmt.Print("\n")
}
