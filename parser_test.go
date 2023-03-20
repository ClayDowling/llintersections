package main

import (
	"bufio"
	"strings"
	"testing"
)

func Test_LineWithArrowIsEdge(t *testing.T) {
	lt := LineType("a -> b")
	if lt != EDGE {
		t.Errorf("Found %d, expected %d", lt, EDGE)
	}
}

func Test_LineWithCommaIsQuery(t *testing.T) {
	lt := LineType("a,b")
	if lt != QUERY {
		t.Errorf("Found %d, expected %d", lt, QUERY)
	}
}

func Test_LineWithNeitherMarkerReturnsUnknown(t *testing.T) {
	lt := LineType("some stuff without commas")
	if lt != UNKNOWN {
		t.Errorf("Found %d, expected %d", lt, UNKNOWN)
	}
}

func Test_EdgeLineReturnsSourceAndDest(t *testing.T) {
	s, d := ParseEdge("a -> b")
	if s != "a" {
		t.Errorf("Found source %s, expected a", s)
	}
	if d != "b" {
		t.Errorf("Found dest %s, expected b", d)
	}
}

func Test_QueryLineReturnsSliceOfElements(t *testing.T) {
	nodes := ParseQuery("c, d, e, f, g")
	if len(nodes) != 5 {
		t.Errorf("Expected 5 nodes, found %d", len(nodes))
		return
	}
	inorder := true
	if nodes[0] != "c" {
		inorder = false
	}
	if nodes[1] != "d" {
		inorder = false
	}
	if nodes[2] != "e" {
		inorder = false
	}
	if nodes[3] != "f" {
		inorder = false
	}
	if nodes[4] != "g" {
		inorder = false
	}

	if !inorder {
		t.Errorf("Expected [c,d,e,f,g] but got %v", nodes)
	}
}

func Test_ParseOnEdgePopulatesEdges(t *testing.T) {
	src := bufio.NewScanner(strings.NewReader("a -> b"))
	Parse(src)
	if Edges["a"] != "b" {
		t.Errorf("Expected b, got \"%s\"", Edges["a"])
	}
}
