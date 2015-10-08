package main

type AdjacencyEntries struct {
	inEdges      []int
	outEdgeCount int
}

func (a *AdjacencyEntries) addInEdge(outIndex int) {
	a.inEdges = append(a.inEdges, outIndex)
}

func (a *AdjacencyEntries) addOutEdge() {
	a.outEdgeCount++
}
