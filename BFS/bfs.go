package BFS

var TestData = map[string][]string{
	"A": {"B", "C", "D"},
	"B": {"A", "E", "F"},
	"C": {"A", "G"},
	"D": {"A", "F"},
	"E": {"B"},
	"F": {"B", "D"},
	"G": {"C"},
}

var SearchList = []string{}
var queue = []string{}
var seem = map[string]bool{}

func Bfs(vertex string) {
	if _, ok := seem[vertex]; !ok {
		seem[vertex] = true
		SearchList = append(SearchList, vertex)
		queue = append(queue, vertex)
	}
	for len(queue) != 0 {
		v := queue[0]
		queue = queue[1:]
		for _, w := range TestData[v] {
			if _, ok := seem[w]; !ok {
				seem[w] = true
				SearchList = append(SearchList, w)
				queue = append(queue, w)
			}
		}
	}

}
