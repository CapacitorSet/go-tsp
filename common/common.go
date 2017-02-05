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

import "math"

type Colour int
type Kind int

type Vertex struct {
	ID     int
	X, Y   int
	Colour // 0, 1, 2
	Kind   // 0, 1: 0 -> 1
	// MustContain []Colour
}

type Edge struct {
	This, Next *Vertex
	Length     float64
}

type Graph []Edge

func (g Graph) Len() int {
	return len(g)
}

func (g Graph) Swap(i, j int) {
	g[i], g[j] = g[j], g[i]
}

func (g Graph) Less(i, j int) bool {
	return g[i].Length < g[j].Length
}

func Distance(a, b Vertex) float64 {
	return math.Sqrt(math.Pow(float64(a.X - b.X), 2) + math.Pow(float64(a.Y - b.Y), 2))
}