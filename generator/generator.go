/// Generates random graph edges (x -> y)
/// and queries for reachability (a,b,c)

package main

import (
	"fmt"
	"math/rand"
	"time"
)

const validkeys = "abcdefghijklmnopqrstuvwxyz"

func main() {
	source := rand.NewSource(time.Now().Unix())
	random := rand.New(source)

	N := random.Intn(20) + 5

	edges := make(map[rune]rune)
	validkeys := []rune(validkeys)
	interval := len(validkeys) / N
	offset := 0
	for i := 0; i < N; i++ {
		skip := random.Intn(interval)
		edges[validkeys[skip+offset]] = ' '
		offset += interval
	}
	actualkeys := make([]rune, 0, N)
	for k := range edges {
		actualkeys = append(actualkeys, k)
	}

	for k := range edges {
		idx := random.Intn(N)
		target := actualkeys[idx]

		if random.Intn(10) > 0 {
			if target == k {
				idx++
				if idx == N {
					idx = 0
				}
				target = actualkeys[idx]
			}
			edges[k] = target
		}
	}

	for k, v := range edges {
		if v == ' ' {
			fmt.Printf("%c\n", k)
		} else {
			fmt.Printf("%c -> %c\n", k, v)
		}
	}

	Q := random.Intn(3) + 2

	for i := 0; i < Q; i++ {
		qn := random.Intn(3) + 2
		interval := len(actualkeys) / qn
		offset = 0
		for x := 0; x < qn; x++ {
			if x > 0 {
				fmt.Printf(",")
			}
			fmt.Printf("%c", actualkeys[offset+random.Intn(interval)])
			offset += interval
		}
		fmt.Printf("\n")
	}
}
