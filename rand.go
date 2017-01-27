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
	"os"
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

type Graph []Edge

func (g Graph) Len() int {
	return len(g)
}

func (g Graph) Swap(i, j int) {
	g[i], g[j] = g[j], g[i]
}

func (g Graph) Less(i, j int) bool {
	if i >= len(g) || j >= len(g) {
		fmt.Printf("graph_of_%d[%d/%d]\n", len(g), i, j)
	}
	return g[i].Length < g[j].Length
}

func findShortestPath(vertices []Vertex, edges Graph) Graph {
	var graph, graph_temp Graph

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
			continue
		}
		graph = graph_temp
	}

	return graph
}

func printGraph(graph Graph) {
	length := 0
	for _, edge := range graph {
		fmt.Printf("%d -> %d\n", (*edge.This).ID, (*edge.Next).ID)
		length += edge.Length
	}
	fmt.Printf("%d edges, length %d\n", len(graph), length)
}

// Returns whether the graph is proper or not
func evaluateGraph(graph Graph, vertices []Vertex) bool {
	if len(graph) < len(vertices) - 1 {
		fmt.Printf("Output graph is shorter than expected: got %d, want %d\n", len(graph), len(vertices) - 1)
		return false
	} else {
		fmt.Print("Output graph length is ok.\n")
		return true
	}
}

func main() {
	r := rand.New(rand.NewSource(0))
	noOfItems := 2 * 45 // 2 * num. palline
	vertices := make([]Vertex, noOfItems)
	var edges Graph

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

	graph := findShortestPath(vertices, edges)

	if evaluateGraph(graph, vertices) {
		os.Exit(0)
	}

	for i := 0; i < len(edges); i++ {
		// fmt.Printf("Blacklisting %d -> %d\n", (*edges[i].This).ID, (*edges[i].Next).ID)
		// http://stackoverflow.com/a/37359662
		edgeset := append(edges[:i], edges[i+1:]...)
		graph = findShortestPath(vertices, edgeset)
		if evaluateGraph(graph, vertices) {
			printGraph(graph)
			os.Exit(0)
		}
	}
}
