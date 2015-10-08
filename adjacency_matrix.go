package main

//import "fmt"

type AdjacencyMatrix struct {
	rows []*AdjacencyEntries
}

func (a *AdjacencyMatrix) initWithSize(size int) {
	a.rows = make([]*AdjacencyEntries, size)
	for i, _ := range a.rows {
		a.rows[i] = &AdjacencyEntries{make([]int, 0, 10), 0}
	}
}

func (a *AdjacencyMatrix) addEdge(out int, in int) {
	// _, ok := a.rows[out]
	// if !ok {
	// 	a.rows[out] = &AdjacencyEntries{make([]int, 10), 0}
	// }
	a.rows[out].addOutEdge()

	// _, ok = a.rows[in]
	// if !ok {
	// 	a.rows[in] = &AdjacencyEntries{make([]int, 10), 0}
	// }
	a.rows[in].addInEdge(out)
}

func (a *AdjacencyMatrix) size() int {
	return len(a.rows)
}

func (a *AdjacencyMatrix) generateNewRank(rank []float64, rankNew []float64, beta float64) {
	for i, entry := range a.rows {
		r := 0.0
		for _, inEdge := range entry.inEdges {
			r += beta * rank[inEdge] / float64(a.rows[inEdge].outEdgeCount)
			//fmt.Println("r:", r)
		}
		rankNew[i] = r
	}
}
