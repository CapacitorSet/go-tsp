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
	"math/rand"
	"os"
)

func main() {
	r := rand.New(rand.NewSource(1))
	noOfItems := 2 * 45 // 2 * num. palline
	noOfContainers := 10
	vertices := make([]Vertex, noOfItems + 1)
	var edges Graph

	for i := 0; i < noOfItems; {
		X := r.Intn(500)
		Y := r.Intn(500 - 40)
		Colour := Colour(r.Intn(3))
		vertices[i] = Vertex{i, X, Y, Colour, 0}
		i++

		X = (500 / noOfContainers) * r.Intn(noOfContainers + 1)
//		X = r.Intn(500); _ = noOfContainers
		Y = 500 - 20
		vertices[i] = Vertex{i, X, Y, Colour, 1}
		i++
	}

	vertices[noOfItems] = Vertex{-1, 0, 0, 0, 1} // Starting point

	vfile, err := os.Create("vertices.json")
	if err != nil {
		panic(err)
	}
	text, err := json.Marshal(vertices)
	if err != nil {
		panic(err)
	}
	_, err = vfile.Write(text)
	if err != nil {
		panic(err)
	}

	for i, da := range vertices {
		for j, a := range vertices {
			if da.Kind == a.Kind {
				continue
			}
			if da.Kind == 0 && da.Colour != a.Colour {
				continue
			}
			if a.ID == -1 {
				continue
			}

			length := Distance(da, a)
			// fmt.Printf("%dc%d -> %dc%d has weight %d\n", da.ID, da.Colour, a.ID, a.Colour, length)
			edges = append(edges, Edge{&(vertices[i]), &(vertices[j]), length})
		}
	}

	efile, err := os.Create("edges.json")
	if err != nil {
		panic(err)
	}
	text, err = json.Marshal(edges)
	if err != nil {
		panic(err)
	}
	_, err = efile.Write(text)
	if err != nil {
		panic(err)
	}
}
