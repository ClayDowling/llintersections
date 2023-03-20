package main

import (
	"bufio"
	"fmt"
	"os"
)

func Intersection([]string) bool {
	return false
}

func main() {
	source := bufio.NewScanner(os.Stdin)

	Parse(source)

	for _, q := range Questions {
		for idx, node := range q {
			if idx > 0 {
				fmt.Print(",")
			}
			fmt.Printf("%s", node)
		}
		fmt.Printf(": %t\n", Intersection(q))
	}

}
