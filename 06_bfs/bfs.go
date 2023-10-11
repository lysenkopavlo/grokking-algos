package main

import "fmt"

// personIsSeller checks if person is a mango seller
// person is mango seller if last symbol of person's name is "m"
func personIsSeller(name string) bool {
	lastSymbol := name[len(name)-1]
	return string(lastSymbol) == "m"
}

// search implements Breadth-First Search alike algorithm
// and returns if there is a mango-seller in queue structure
func search(graph map[string][]string, name string) string {

	var (
		searchQueue []string // implement queue-alike structure and add first-level persons
		searched    []string // add persons who has been searched already
	)
	searchQueue = append(searchQueue, graph[name]...) // add alice, bob, claire in first iteration

	// while there is someone searchQueue
	for len(searchQueue) > 0 {
		person := searchQueue[0]      // take first element in searchQueue
		searchQueue = searchQueue[1:] // exclude this element from searchQueue
		personSearched := false       // flag to exclude person from repeat

		// check to see if this person has already been searched
		for i := range searched {
			if searched[i] == person {
				personSearched = true
			}
		}

		// if it is first time search
		if !personSearched {
			if personIsSeller(person) {
				return fmt.Sprintf("%s is a mango seller!\n", person)
			}

			searchQueue = append(searchQueue, graph[person]...) // append next person to search queue
			searched = append(searched, person)                 // add have been just searched person in searched slice
		}
	}

	// otherwise there is no mango-seller
	return "There isn't any mango seller!"
}

func main() {
	// create graph alike struct via map:
	graph := make(map[string][]string)

	// add first level friends
	graph["you"] = []string{"alice", "bob", "claire"} // these are my  friends

	// add second level friends
	graph["bob"] = []string{"anuj", "peggy"}    // these are friends of my friend bob
	graph["alice"] = []string{"peggy"}          // this friend of my friend alice
	graph["claire"] = []string{"tom", "johnny"} // these are friends of my friend

	// add third level friends
	graph["anuj"] = []string{}
	graph["peggy"] = []string{}
	graph["tom"] = []string{}
	graph["johnny"] = []string{}

	fmt.Println(search(graph, "you"))
}
