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

func FromV1() {
	fmt.Println("I'm from v1.0.0")
}

func FromV11() {
	fmt.Println("I'm from v1.1.0")
}

func FromV111() {
	fmt.Println("I'm from v1.1.1")
}

func FromV123() {
	fmt.Println("I'm from v1.2.3")
}

func FromV134() {
	fmt.Println("I'm from v1.3.4")
}