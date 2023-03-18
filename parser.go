package main

import "strings"

const UNKNOWN = -1
const EDGE = 1
const QUERY = 2

func LineType(line string) int {
	if strings.Contains(line, "->") {
		return EDGE
	}
	if strings.Contains(line, ",") {
		return QUERY
	}
	return UNKNOWN
}

func ParseEdge(line string) (string, string) {
	parts := strings.Split(line, "->")
	if len(parts) != 2 {
		return "", ""
	}
	return strings.TrimSpace(parts[0]), strings.TrimSpace(parts[1])
}

func ParseQuery(line string) []string {
	parts := strings.Split(line, ",")
	nodes := make([]string, len(parts))
	for i, val := range parts {
		nodes[i] = strings.TrimSpace(val)
	}
	return nodes
}
