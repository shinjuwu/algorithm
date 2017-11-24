package main

import (
	"algorithm/BFS"
	"algorithm/DFS"
	"algorithm/TOPOSORT"
	"fmt"
)

func main() {
	//Dfs test
	DFS.Dfs("A")
	fmt.Println(DFS.SearchList)

	//Bfs test
	BFS.Bfs("A")
	fmt.Println(BFS.SearchList)

	//Topo sort via DFS test
	for i, course := range TOPOSORT.TopoSort(TOPOSORT.Prereqs) {
		fmt.Printf("%d:\t%s\n", i+1, course)
	}
}
