package solution

import (
	"github.com/kyokomi/emoji/v2"
)

func GetMessage() string {
	return "Hello :world_map:!"
}

func main() {
	pizzaMessage := emoji.Sprint(GetMessage())
	emoji.Print(pizzaMessage)
}
