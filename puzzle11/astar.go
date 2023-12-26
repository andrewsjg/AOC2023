package puzzle11

import (
	"container/heap"
)

type Pather interface {
	// PathNeighbors returns the direct neighboring nodes of this node which
	// can be pathed to.
	neighbours() []Pather
	// PathNeighborCost calculates the exact movement cost to neighbor nodes.
	pathToNeighBourCost(to Pather) float64
	// PathEstimatedCost is a heuristic method for estimating movement costs
	// between non-adjacent nodes.
	pathCostEstimate(to Pather) float64
}

// node is a wrapper to store A* data for a Pather node.
type node struct {
	pather Pather
	cost   float64
	rank   float64
	parent *node
	open   bool
	closed bool
	index  int
}

// nodeMap is a collection of nodes keyed by Pather nodes for quick reference.
type nodeMap map[Pather]*node

// get gets the Pather object wrapped in a node, instantiating if required.
func (nm nodeMap) get(p Pather) *node {
	n, ok := nm[p]
	if !ok {
		n = &node{
			pather: p,
		}
		nm[p] = n
	}
	return n
}

func path(from Pather, to Pather) (path []Pather, distance float64, found bool) {

	nm := nodeMap{}
	locationQueue := &priorityQueue{}
	heap.Init(locationQueue)

	fromNode := nm.get(from)

	fromNode.open = true
	heap.Push(locationQueue, fromNode)

	for {

		if locationQueue.Len() == 0 {
			// No path
			return
		}

		currentLocation := heap.Pop(locationQueue).(*node)
		currentLocation.open = false
		currentLocation.closed = true

		//fmt.Printf("Current Location idx: %d\n", currentLocation.index)
		if currentLocation == nm.get(to) {
			// Found a path
			p := []Pather{}
			curr := currentLocation

			for curr != nil {
				p = append(p, curr.pather)
				curr = curr.parent
			}

			return p, currentLocation.cost, true
		}

		for _, neighbour := range currentLocation.pather.neighbours() {
			cost := currentLocation.cost + currentLocation.pather.pathToNeighBourCost(neighbour)

			//fmt.Printf("Current location cost: %f Neighbour cost: %f Cost: %f\n", currentLocation.cost, neighbour.cost, cost)
			neighbourNode := nm.get(neighbour)
			if cost < neighbourNode.cost {

				if neighbourNode.open {
					heap.Remove(locationQueue, neighbourNode.index)
				}

				neighbourNode.open = false
				neighbourNode.closed = false
			}

			if !neighbourNode.open && !neighbourNode.closed {

				neighbourNode.cost = cost
				neighbourNode.rank = cost + neighbour.pathCostEstimate(to)
				neighbourNode.parent = currentLocation
				neighbourNode.open = true

				heap.Push(locationQueue, neighbourNode)
			}
		}
	}
}
