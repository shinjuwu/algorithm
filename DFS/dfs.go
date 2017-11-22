package DFS

var TestData = map[string][]string{
	"A": {"B", "C", "D"},
	"B": {"E", "F"},
	"C": {"G"},
	"D": {"F"},
	"E": {"B"},
	"F": {"B", "D"},
	"G": {"C"},
}
var seem = make(map[string]bool)
var SearchList []string

func Dfs(vertex string) {

	if _, ok := seem[vertex]; !ok {
		seem[vertex] = true
		SearchList = append(SearchList, vertex)
	}

	for _, v := range TestData[vertex] {
		if !seem[v] {
			Dfs(v)
		}
	}

}
