package solution

import (
	"github.com/kyokomi/emoji/v2"
)

func getMessage() string {
	return "Hello :world_map:!"
}

func main() {
	emoji.Sprint(getMessage())
}
