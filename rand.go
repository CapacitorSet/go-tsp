/*
Copyright 2017 @CapacitorSet

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

 http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package main

import (
	"fmt"
	"math/rand"
	"sort"
)

type Colour int
type Kind int

type Vertex struct {
	ID     int
	X, Y   int
	Colour // 0, 1, 2
	Kind   // 0, 1: 0 -> 1
}

type Edge struct {
	This, Next *Vertex
	Length     int
}

type graph []Edge

func (g graph) Len() int {
	return len(g)
}

func (g graph) Swap(i, j int) {
	g[i], g[j] = g[j], g[i]
}

func (g graph) Less(i, j int) bool {
	return g[i].Length < g[j].Length
}

func main() {
	r := rand.New(rand.NewSource(0))
	noOfItems := 2 * 3 // 2 * num. palline
	vertices := make([]Vertex, noOfItems)
	var edges, graph, graph_temp graph

	for i := 0; i < noOfItems; {
		X := r.Intn(500)
		Y := r.Intn(500 - 40)
		Colour := Colour(r.Intn(3))
		vertices[i] = Vertex{i, X, Y, Colour, 0}
		i++

		X = r.Intn(500)
		Y = 500 - 20
		vertices[i] = Vertex{i, X, Y, Colour, 1}
		i++
	}

	for i, da := range vertices {
		for j, a := range vertices {
			if da.Kind == a.Kind {
				continue
			}
			if da.Kind == 0 && (da.Colour != a.Colour) {
				continue
			}
			if da == a {
				continue
			}

			length := r.Intn(100)
			edges = append(edges, Edge{&(vertices[i]), &(vertices[j]), length})
		}
	}

	sort.Sort(edges)

	for _, edge := range edges {
		graph_temp = graph

		graph_temp = append(graph_temp, edge)

		/*if edge.This == edge.Next {
			panic("Edge from self to self")
		}*/

		bailout := false
		for _, currentEdge := range graph {
			if currentEdge.This == edge.This {
				bailout = true
				break
			}
			if currentEdge.Next == edge.Next {
				bailout = true
				break
			}
		}
		if bailout == true {
			continue
		}

		currentEdge := edge
		visited := make(map[*Vertex]bool)
		doesCycle := false
		for {
			if visited[currentEdge.Next] {
				doesCycle = true
				break
			}
			visited[currentEdge.This] = true

			found := false
			for _, item := range graph_temp {
				if item.This == currentEdge.Next {
					currentEdge = item
					found = true
					break
				}
			}
			if !found {
				break
			}
		}
		if doesCycle {
			break
		}

		graph = graph_temp
	}

	length := 0
	for _, edge := range graph {
		// fmt.Printf("%d -> %d\n", (*edge.This).ID, (*edge.Next).ID)
		length += edge.Length
	}
	fmt.Printf("%d edges, length %d\n", len(graph), length)
	if len(graph) < (len(vertices) - 1) {
		fmt.Print("Warning: output graph is shorter than expected\n")
	}
}
