package main

//import "runtime"

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
	a.rows[out].addOutEdge()
	a.rows[in].addInEdge(out)
}

func (a *AdjacencyMatrix) size() int {
	return len(a.rows)
}

//x is the concurrency factor
//note: will error if number of nodes is too small relative to x
func (a *AdjacencyMatrix) generateNewRank(rank []float64, rankNew []float64, beta float64, x int) {
	dones := make(chan bool)

	sectionSize := (a.size() + x - 1) / x

	for i := 0; i < x; i++ {
		go func(i int) {
			top := sectionSize * (i + 1)
			if i == x-1 {
				top = a.size()
			}
			for index, entry := range a.rows[sectionSize*i : top] {
				r := 0.0
				for _, inEdge := range entry.inEdges {
					r += beta * rank[inEdge] / float64(a.rows[inEdge].outEdgeCount)
				}
				rankNew[index+sectionSize*i] = r
			}
			dones <- true
		}(i)
	}

	for i := 0; i < x; i++ {
		<-dones
	}
}
