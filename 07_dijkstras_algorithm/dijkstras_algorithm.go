package main

import (
	"fmt"
	"math"
	"strings"
)

type (
	graph    map[string]map[string]int // to hold graph from exercise
	Costs    map[string]int            // to hold weights of graph edges
	Parents  map[string]string         // to hold "start", "A", "B", and "end" nodes
	searched map[string]struct{}       // to hold nodes have been searched

)

func findLowestCostNode(c Costs, searchedList searched) string {
	min := math.MaxInt
	var lowestCostNode string

	for node, cost := range c {
		// of node is not in the searched list
		if _, ok := searchedList[node]; cost < min && !ok {
			min = cost
			lowestCostNode = node
		}
	}

	return lowestCostNode
}

// dijkstrasAlgorithm finds the shortest path of the weighted graph
// and returns pares "node - nearest node"
func dijkstrasAlgorithm(g graph, start, end string) Parents {
	parents := Parents{}
	costs := Costs{}

	// take all neighbors of the "start" node
	for node, cost := range g[start] {
		fmt.Println("current node is:", node)
		// setting its cost
		costs[node] = cost // "a": 6, "b": 2, "end": ?

		// and setting its parent:
		parents[node] = start
	}

	// coz we don't know yet the cost of the last node "end"
	// we set it to maximum int
	costs[end] = math.MaxInt // "end": 9223372036854775807

	// create list to track checked nodes
	searchedList := searched{}

	// take the node with lowest cost
	node := findLowestCostNode(costs, searchedList)

	for node != "" { // while node has a neighbor

		// check all these nodes neighbors
		for neighbor, value := range g[node] {

			// evaluate the cost to reach this neighbor
			newCost := costs[node] + value

			if costs[neighbor] > newCost {
				costs[neighbor] = newCost
				parents[neighbor] = node // update a parent in this case
			}
		}
		// current node has checked
		searchedList[node] = struct{}{} // put it in the list
		node = findLowestCostNode(costs, searchedList)
	}
	return parents
}

// drawPath works with results of dijkstrasAlgorithm function
// returns resulting path in format: parent  --- > nearest_neighbor"
func drawPath(p Parents) string {
	start := "end"
	path := []string{}
	path = append(path, start)
	for start != "start" {
		start = p[start]
		path = append(path, start)
	}

	// reverse
	for i, j := 0, len(path)-1; i < j; i, j = i+1, j-1 {
		path[i], path[j] = path[j], path[i]
	}

	return strings.Join(path, " ---> ")
}

func main() {
	g := graph{
		"start": map[string]int{
			"a": 6,
			"b": 2,
		},
		"a": map[string]int{
			"end": 1,
		},
		"b": map[string]int{
			"a":   3,
			"end": 5,
		},
		"end": map[string]int{},
	}

	resultingPath := drawPath(dijkstrasAlgorithm(g, "start", "end"))

	fmt.Println(resultingPath)
}
