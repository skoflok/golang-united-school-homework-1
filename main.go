package main

import (
	"github.com/kyokomi/emoji/v2"
)

func main() {
	mes := emoji.Sprint("Hello 🗺️")

	emoji.Print(mes)
}
