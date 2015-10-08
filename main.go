package main

import (
	"flag"
	"fmt"
	"io"
	"io/ioutil"
	"math"
	"runtime"
	"strconv"
	"strings"
	"time"
)

func check(err error) {
	if err != nil && err != io.EOF {
		panic(err)
	}
}

var filename = flag.String("file", "web-Google.txt", "file to parse on")

// file format rules:
// first line is an integer for the number of nodes
// all following lines are a pain of integers
// the first int is the outIndex, second in the inIndex
func main() {
	flag.Parse()
	fmt.Println(*filename)

	start := time.Now()
	byteArray, err := ioutil.ReadFile(*filename)
	finish := time.Since(start)
	fmt.Println("raw read time:", finish)
	check(err)

	start = time.Now()

	s := string(byteArray[:])

	lines := strings.Split(s, "\r\n")

	var matrix AdjacencyMatrix

	nodeCount, err := strconv.Atoi(lines[0])
	check(err)

	matrix.initWithSize(nodeCount)

	indexMap := make(map[int]int)
	//backIndexMap := make(map[int]int)

	nextIndex := 0

	for _, line := range lines[1:] {
		twoNums := strings.Fields(line)

		if len(twoNums) != 2 {
			fmt.Println(line)
			panic(fmt.Sprintln("bad line is not 2 nums:", len(twoNums)))
		}

		outIndex, err := strconv.Atoi(twoNums[0])
		check(err)
		inIndex, err := strconv.Atoi(twoNums[1])
		check(err)

		out, ok := indexMap[outIndex]
		if !ok {
			out = nextIndex
			indexMap[outIndex] = nextIndex
			//backIndexMap[nextIndex] = outIndex
			nextIndex++
		}

		in, ok := indexMap[inIndex]
		if !ok {
			in = nextIndex
			indexMap[inIndex] = nextIndex
			//backIndexMap[nextIndex] = inIndex
			nextIndex++
		}

		matrix.addEdge(out, in)
	}

	finish = time.Since(start)
	//reading file done
	fmt.Println("time to read file:", finish)

	start = time.Now()

	iteration := 0
	const epsilon = 0.0000000001
	const beta = 0.8
	rank := make([]float64, nodeCount)
	rankNew := make([]float64, nodeCount)
	done := false

	concurrencyFactor := runtime.GOMAXPROCS(0) * 16

	for i, _ := range rank {
		rank[i] = 1.0 / float64(nodeCount)
	}

	for !done {
		matrix.generateNewRank(rank, rankNew, beta, concurrencyFactor)

		sum := 0.0

		for _, v := range rankNew {
			sum += v
		}

		jumpFactor := (1.0 - sum) / float64(nodeCount)

		for i, _ := range rankNew {
			rankNew[i] += jumpFactor
		}

		done = true
		for i, _ := range rankNew {
			if math.Abs(rankNew[i]-rank[i]) > epsilon {
				done = false
				break
			}
		}

		// fmt.Println(rank)
		// fmt.Println(rankNew)
		// break

		copy(rank, rankNew)
		iteration++
	}

	finish = time.Since(start)

	fmt.Println("time to run algorithm:", finish)

	fmt.Println("rank of node 99:", rank[indexMap[99]])
	fmt.Println("final iteration:", iteration)
}
