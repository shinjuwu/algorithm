package DFS

var TestData = map[string][]string{
	"A": {"B", "C", "D"},
	"B": {"A", "E", "F"},
	"C": {"A", "G"},
	"D": {"A", "F"},
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
