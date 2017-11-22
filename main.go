package main

import (
	"algorithm/BFS"
	"algorithm/DFS"
	"fmt"
)

func main() {
	//Dfs test
	DFS.Dfs("A")
	fmt.Println(DFS.SearchList)

	//Bfs test
	BFS.Bfs("A")
	fmt.Println(BFS.SearchList)
}
