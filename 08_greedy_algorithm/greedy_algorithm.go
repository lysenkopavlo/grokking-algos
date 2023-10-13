package main

import "fmt"

type (
	States   map[string]struct{} // type to hold different states of country
	Stations map[string]States   // type to display Station to State coverage
)

func main() {

	// declaring amount fo states needs to be covered
	statesNeeded := States{
		"mt": {},
		"wa": {},
		"or": {},
		"id": {},
		"nv": {},
		"ut": {},
		"ca": {},
	}

	// declaring how many states does each station covers
	stations := Stations{
		"kone": {
			"id": {},
			"nv": {},
			"ut": {},
		},
		"ktwo": {
			"wa": {},
			"id": {},
			"mt": {},
		},
		"kthree": {
			"or": {},
			"nv": {},
			"ca": {},
		},
		"kfour": {
			"nv": {},
			"ut": {},
		},
		"kfive": {
			"ca": {},
			"az": {},
		},
	}

	//
	result := greedySearch(stations, statesNeeded)

	// print the result
	fmt.Printf("The list of stations to pick: %v\n", result)
}

// greedySearch implies greedy algorithm to find optimal number of radio stations
// and returns names of these stations
func greedySearch(stations Stations, statesNeeded States) map[string]struct{} {

	// declaring resulting variable
	optimalStationsList := map[string]struct{}{}

	// while not all the states are covered
	for len(statesNeeded) > 0 {

		var (
			bestStation   string //
			statesCovered States // to track already covered states
		)

		for station, states := range stations {
			covered := intersect(statesNeeded, states)
			if len(covered) > len(statesCovered) {
				bestStation = station
				statesCovered = covered
			}
			optimalStationsList[bestStation] = struct{}{}
			remove(statesNeeded, statesCovered)
		}
	}
	return optimalStationsList
}

// remove deletes an item from statesNeeded table
func remove(statesNeeded States, covered map[string]struct{}) {
	for k := range covered {
		delete(statesNeeded, k)
	}
}

// intersect finds the intersection between statesNeeded and statesCovered
// and returns the intersection as table
func intersect(statesNeeded, statesCovered States) States {
	numberOfStates := States{}
	from, to := statesNeeded, statesCovered
	if len(statesNeeded) < len(statesCovered) {
		from, to = statesCovered, statesNeeded
	}

	for k := range from {
		if _, ok := to[k]; ok {
			numberOfStates[k] = struct{}{}
		}
	}

	return numberOfStates
}
