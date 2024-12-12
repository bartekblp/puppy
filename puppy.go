package puppy

import (
	"fmt"

	"github.com/bartekblp/dog"
)

func Bark() string {
	return "Woof!"
}

func Barks() string {
	return "Woof! Woof! Woof!"
}

func BigBark() string {
	return dog.WhenGrownUp(Bark())
}

func BigBarks() string {
	return dog.WhenGrownUp(Barks())
}

func fromV1() {
	fmt.Println("I'm from v1.0.0")
}

func fromV11() {
	fmt.Println("I'm from v1.1.0")
}

func fromV111() {
	fmt.Println("I'm from v1.1.1")
}