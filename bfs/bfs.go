package bfs

import (
	"container/list"
	"fmt"
	"strings"
)

func initGraph() map[string][]string {
	var graph = make(map[string][]string)
	graph["you"] = []string{"alice", "bob", "claire"}
	graph["bob"] = []string{"anuj", "peggy"}
	graph["alice"] = []string{"peggy"}
	graph["claire"] = []string{"thom", "jonny"}
	return graph
}

func SearchQueue() bool {
	graph := initGraph()
	queue := list.New()
	for _, el := range graph["you"] {
		queue.PushBack(el)
	}
	var searched []string

	for queue.Len() > 0 {
		var person string
		person = queue.Front().Value.(string)
		fmt.Println(person)
		queue.Remove(queue.Front())
		if !find(searched, person) {
			if personIsSeller(person) {
				fmt.Println(person, "is a mango seller")
				return true
			} else {
				for _, el := range graph[person] {
					queue.PushBack(el)
				}
				searched = append(searched, person)
			}

		}
	}
	return false
}

func personIsSeller(name string) bool {
	return strings.HasSuffix(name, "m")
}

func find(slice []string, val string) bool {
	for _, item := range slice {
		if item == val {
			return true
		}
	}
	return false
}
