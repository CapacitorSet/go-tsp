go-tsp
======

Solves the Travelling Salesman Problem on directed, incomplete graphs.

## Problem

The specific problem is the following:

>There are points with `Type` either `0` or `1`, and `Colour` either `0`, `1` or `2`. A type-0 point can connect to a type-1 point of the same colour; a type-1 point can connect to any type-0 points.

This represents carrying items from a pickup point to a delivery point; the result will be the optimal route to deliver every item to an appropriate delivery point. It also models having delivery points with capacities (eg. 2 red and 1 blue items), by simply representing each delivery "slot" with a vertex.

## Algorithm

The algorithm is detailed at http://lcm.csa.iisc.ernet.in/dsa/node186.html:

* Sort the edges in the increasing order of weights.
* Iterate the edges starting with the least-cost one. Add each one to the output graph if the result:
  * Has no vertices with degree three or more
  * Does not form a cycle

This is a heuristic that finds a good solution, though not a provably-optimal one.

## Implementation

The points are stored in an array of structs. The vertices are stored in an array of structs, taking care not to add impossible edges (edge to self, edge to same type); each vertex is represented as a pointer to the source, a pointer to the destination, and a "length" field (a random number, for now).

The list of edges is sorted, and then iterated.

An "output graph" is stored in memory; a "temporary graph" is created at each iteration. We attempt to add the edge to this graph, then verify that no vertex has degree three (by verifying that no existing edge has the same source or the same destination as the candidate), then verify that there are no loops (by traversing the graph starting with the candidate edge, then marking each edge as visited, then failing if we traverse an edge we already visited and passing if we eventually find an edge that leads to nowhere, i.e. to a vertex that is not in the temporary graph), then copy the temporary graph in the output graph and do the next iteration.
