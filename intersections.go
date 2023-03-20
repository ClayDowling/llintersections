package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
)

func Intersection([]string) bool {
	return false
}

func main() {

	graphviz := flag.Bool("graph", false, "generate a graphviz graph")
	flag.Parse()

	source := bufio.NewScanner(os.Stdin)

	Parse(source)

	if *graphviz {
		fmt.Println("digraph intersection {")
		for k, v := range Edges {
			fmt.Printf("  %s -> %s\n", k, v)
		}
		fmt.Println("}")
		os.Exit(0)
	}

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
