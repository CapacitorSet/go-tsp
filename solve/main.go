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
	"encoding/json"
	"fmt"
	"os"
	"io/ioutil"
	"sort"
)

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
	// graph is sorted by length;
	// "sort" it by logical order (a -> b, b -> c, c -> d)
	var newGraph Graph
	thisId := -1
	for {
		done := true
		for _, edge := range graph {
			if (*edge.This).ID != thisId {
				continue;
			}
			done = false;
			newGraph = append(newGraph, edge)
			thisId = (*edge.Next).ID
		}
		if done {
			break;
		}
	}

	length := float64 (0)
	for _, edge := range newGraph {
		fmt.Printf("%d -> %d\n", (*edge.This).ID, (*edge.Next).ID)
		length += edge.Length
	}
	fmt.Printf("%d edges, length %#v\n", len(newGraph), length)

	sfile, err := os.Create("solution.json")
	if err != nil {
		panic(err)
	}
	text, err := json.Marshal(newGraph)
	if err != nil {
		panic(err)
	}
	_, err = sfile.Write(text)
	if err != nil {
		panic(err)
	}
}

// Returns whether the graph is proper or not
func evaluateGraph(graph Graph, vertices []Vertex) bool {
	for _, edge := range graph {
		if (*edge.Next).ID == -1 {
			panic("Edge pointing to -1!")
		}
	}
	if len(graph) < len(vertices) - 1 {
		fmt.Printf("Output graph is shorter than expected: got %d, want %d\n", len(graph), len(vertices) - 1)
		return false
	} else {
		fmt.Print("Output graph length is ok.\n")
		return true
	}
}

func main() {
	noOfItems := 2 * 45 // 2 * num. palline
	vertices := make([]Vertex, noOfItems + 1)
	id2vertex := make(map[int](*Vertex)) // For quicker deserialization

	vtext, err := ioutil.ReadFile("vertices.json")
	if err != nil {
		panic(err)
	}
	err = json.Unmarshal(vtext, &vertices)
	if err != nil {
		panic(err)
	}
	for i, vertex := range vertices {
		id2vertex[vertex.ID] = &vertices[i]
	}

	var edges Graph

	etext, err := ioutil.ReadFile("edges.json")
	if err != nil {
		panic(err)
	}
	err = json.Unmarshal(etext, &edges)
	if err != nil {
		panic(err)
	}
	// Note: at this time, the edge pointers are messed up.
	// They must be fixed to point to vertices.
	for i, edge := range edges {
		var found bool
		edges[i].This, found = id2vertex[(*edge.This).ID]
		if !found {
			panic("Couldn't remap pointers")
		}
		edges[i].Next, found = id2vertex[(*edge.Next).ID]
		if !found {
			panic("Couldn't remap pointers")
		}
	}

	graph := findShortestPath(vertices, edges)

	if evaluateGraph(graph, vertices) {
		printGraph(graph)
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
