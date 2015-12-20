package main

import (
	"io/ioutil"
	"strings"
	"strconv"
	"fmt"
)

type graph struct {
	verticies map[string]int
	edges map[string]int
	currentPath []node
	allPaths [][]node
	start string
}

type node struct {
	name string
	cost int
}

func newGraph() (*graph) {	
	g := new(graph)
	g.edges = make(map[string]int)
	g.verticies = make(map[string]int)
	g.allPaths = make([][]node, 0)
	g.currentPath = make([]node, 0)
	
	return g
}

func (g *graph) numberOfVerticies() int {
	return len(g.verticies)
}

func (g *graph) getLowestPath() ([]node, int) {
	numVerts := g.numberOfVerticies()
	
	var lowestPath []node 
	lowestCost := 999999999
	for _,path := range g.allPaths {
		if len(path) == numVerts {
			pathCost := 0
			for _,node := range path {
				pathCost += node.cost
			}
			
			if pathCost < lowestCost {
				lowestCost = pathCost
				lowestPath = path
			}
		}
	}
	
	return lowestPath, lowestCost
}

func (g *graph) getHighestPath() ([]node, int) {
	numVerts := g.numberOfVerticies()
	
	var highestPath []node 
	highestCost := -1
	for _,path := range g.allPaths {
		if len(path) == numVerts {
			pathCost := 0
			for _,node := range path {
				pathCost += node.cost
			}
			
			if pathCost > highestCost {
				highestCost = pathCost
				highestPath = path
			}
		}
	}
	
	return highestPath, highestCost
}

func (g *graph) calcPaths(vertex string, currentPath []node) {
	if(vertex == g.start){
		currentPath = make([]node, 0)
		currentPath = append(currentPath, node{vertex, 0})
	}
		
	n := g.getNeighbors(vertex)
	
	if (len(n) < 1 || len(currentPath) >= g.numberOfVerticies()) {
		g.allPaths = append(g.allPaths, currentPath)
	} else {
		for k,v := range n {
			if containsNode(currentPath, k) { continue }
			cp := append(currentPath, node{k,v})
			g.calcPaths(k, cp)
		}
	}
}

func containsNode(nodeList []node, s string) bool {
	for _,v := range nodeList {
		if v.name == s {
			return true
		}
	} 
	
	return false
}

func (g *graph) getNeighbors(vertex string) map[string]int {
	vertexKey := fmt.Sprintf("%s,", vertex)
	neighbors := make(map[string]int)
	
	for key,val := range g.edges {		
		if strings.Contains(key, vertexKey) {
			t := strings.Split(key, ",")
			neighbors[t[1]] = val
		}
	}
	
	return neighbors
}

func main() {
	g := newGraph()
	
	b, _ := ioutil.ReadFile("input.txt")
	input := string(b[:])
	lines := strings.Split(input, "\n")


	for _, line := range lines {
		if(len(line) < 1) {continue}
		
		//AlphaCentauri to Snowdin = 66
		distanceTokens := strings.Split(line, " = ")
		distance, _ := strconv.Atoi(distanceTokens[1])
		
		toFromTokens := strings.Split(distanceTokens[0], " to ")  
		from := toFromTokens[0]
		to := toFromTokens[1]
		
		key := fmt.Sprintf("%s,%s", from, to)
		g.edges[key] = distance
		
		key = fmt.Sprintf("%s,%s", to, from)
		g.edges[key] = distance
		
		g.verticies[from] = -1
		g.verticies[to] = -1
		
		fmt.Printf("%s -> %s = %d\n", from, to, distance)
	}
	
	
	for i := 0; i < 5; i++ {
	fmt.Println("")
	
	for k := range g.verticies {
		g.start = k
		g.calcPaths(k, nil)
		path, cost := g.getLowestPath()
		highestPath, highestCost := g.getHighestPath()
		fmt.Println("")
		fmt.Println("Lowest")
		fmt.Println(path)
		fmt.Println(cost)
		fmt.Println("")
		fmt.Println("Highest")
		fmt.Println(highestPath)
		fmt.Println(highestCost)
	}
	}
	
	fmt.Println("total num paths calculated", len(g.allPaths))
}
