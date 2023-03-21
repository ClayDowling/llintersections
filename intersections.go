package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
)

type ResultSet []string

type Search struct {
	Current string
	Found   ResultSet
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

	s.Found = append(s.Found, s.Current)
	s.Current = Edges[s.Current]

	if s.Current == "" {
		s.ended = true
		return ""
	}

	if s.Found.Contains(s.Current) {
		return ""
	}

	return s.Current
}

func (haystack *ResultSet) Contains(needle string) bool {
	for _, v := range *haystack {
		if v == needle {
			return true
		}
	}
	return false
}

func Intersection(question []string) bool {

	nq := len(question)
	found := []ResultSet{}
	for _, q := range question {
		a := NewSearch(q)
		for a.Advance() != "" {
		}
		found = append(found, a.Found)
	}

	for i := 0; i < nq-1; i++ {
		subject := found[i]
		for j := i + 1; j < nq; j++ {
			candidate := found[j]
			for _, target := range candidate {
				if subject.Contains(target) {
					return true
				}
			}
		}
	}

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
