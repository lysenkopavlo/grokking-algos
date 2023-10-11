package main

import "fmt"

func checkVoter(list map[string]bool, name string) {

	if list[name] {
		fmt.Println("Kick them out!")
	}

	if !list[name] {
		fmt.Println("let them voite!")
		list[name] = true
	}
}

func main() {

	voted := map[string]bool{}
	checkVoter(voted, "tom")
	checkVoter(voted, "mike")
	checkVoter(voted, "mike")

}
