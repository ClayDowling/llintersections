package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
)

type Search struct {
	Current string
	found   []string
	ended   bool
}

func NewSearch(start string) Search {
	s := Search{start, []string{}, false}
	return s
}

func (s *Search) Advance() string {

	if s.ended {
		return ""
	}

	s.found = append(s.found, s.Current)
	s.Current = Edges[s.Current]

	if s.Current == "" {
		s.ended = true
		return ""
	}

	for _, v := range s.found {
		if v == s.Current {
			s.ended = true
			return ""
		}
	}

	return s.Current
}

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
