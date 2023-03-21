package main

import (
	"bufio"
	"strings"
	"testing"
)

func TestSearch(t *testing.T) {

	src := bufio.NewScanner(strings.NewReader("a -> b\nb -> c\nc -> a\nd -> e"))
	Parse(src)

	t.Run("AdvanceIncrementsByOne", func(t *testing.T) {
		s := NewSearch("a")
		actual := s.Advance()
		if actual != "b" {
			t.Errorf("Found %s but expected b", actual)
		}
	})

	t.Run("AdvanceReturnsEmptyNodeWhenNoConnectingNode", func(t *testing.T){
		s := NewSearch("d")
		actual := s.Advance()
		if actual != "e" {
			t.Fatalf("Expected e, got %s", actual)
		}
		actual = s.Advance()
		if actual != "" {
			t.Errorf("Expected empty node, got %s", actual)
		}
	})

	t.Run("AdvanceStopsWhenLoopDetected", func(t *testing.T){
		s := NewSearch("a")
		s.Advance()
		actual := s.Advance()
		if actual != "c" {
			t.Fatalf("Expected c, found %s", actual)
		}
		actual = s.Advance()
		if actual != "" {
			t.Errorf("Expected end of list, found %s", actual)
		}
	})

}

func TestIntersection(t *testing.T) {
	src := bufio.NewScanner(strings.NewReader("a -> b\nb -> c\nc -> a\nd -> e\nf -> g\ng -> h"))
	Parse(src)

	t.Run("NodesInSameGraphIntersect", func(t *testing.T){
		if Intersection([]string{"a","c"}) != true {
			t.Error("Expected a,c to intersect")
		}
	})
	t.Run("NodesInDifferentGraphsDoNotIntersect", func(t *testing.T){
		if Intersection([]string{"b","e"}) {
			t.Error("Did not expect b,e to intersect")
		}
	})
	t.Run("MultipleNodesCanBeProcessed", func(t *testing.T){
		if Intersection([]string{"a","e","h"}) {
			t.Error("Did not expect a,e,h to intersect")
		}
	})
}
