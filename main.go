package main

import (
	"github.com/kyokomi/emoji/v2"
)

func getMessage() string {
	return "Hello :world_map:!"
}
func main() {
	emoji.Println(emoji.Sprint(getMessage()))
}
