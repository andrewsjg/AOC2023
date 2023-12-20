package puzzle8

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"sync"
)

type node struct {
	left  string
	right string
}

type nodeMap map[string]node

func Part1Solve(input string) {
	file, err := os.ReadFile(input)
	if err != nil {
		fmt.Println(err)
	}

	stringInput := string(file)

	myNodeMap, instructions := getNodeMap(stringInput)

	fmt.Printf("Total numer of steps: %d\n", navigateMap(myNodeMap, instructions))
}

func Part2Solve(input string) {

	file, err := os.ReadFile(input)
	if err != nil {
		fmt.Println(err)
	}

	stringInput := string(file)

	myNodeMap, instructions := getNodeMap(stringInput)

	fmt.Printf("Total numer of steps: %d\n", p2Navigate(myNodeMap, instructions))
}

func navigateMap(myNodeMap nodeMap, instructions string) (steps int) {

	currentNode := "AAA"
	destinationNode := "ZZZ"

	for currentNode != destinationNode {
		for i := 0; i < len(instructions); i++ {

			if currentNode == destinationNode {
				return steps
			}

			curentInstruction := string(instructions[i])
			currentNode = move(myNodeMap[currentNode], myNodeMap, curentInstruction)
			steps++

		}
	}

	// Catch all
	if currentNode == destinationNode {
		return steps
	} else {
		return -1
	}

}

func p2Navigate(myNodeMap nodeMap, instructions string) (steps int) {

	startNodes := []string{}
	for k := range myNodeMap {
		if k[2] == 'A' {

			startNodes = append(startNodes, k)
		}
	}

	// Buffered channel so we can process the results once the channel is closed.
	stepChannel := make(chan int, len(startNodes))
	var wg sync.WaitGroup

	for _, startNode := range startNodes {
		if startNode != "" {

			wg.Add(1)

			go concurrentNavigate(myNodeMap[startNode], myNodeMap, instructions, stepChannel, &wg)
		}

	}

	wg.Wait()
	close(stepChannel)

	steplist := []int{}

	// Make an array of the results
	for steps := range stepChannel {
		steplist = append(steplist, steps)
	}

	steps = LCM(steplist[0], steplist[1:])

	return

}

func move(fromNode node, nmap nodeMap, instruction string) (toNode string) {

	if instruction == "L" {
		return fromNode.left
	}

	return fromNode.right
}

func concurrentNavigate(startNode node, nmap nodeMap, instructions string, stepChan chan int, wg *sync.WaitGroup) {
	defer wg.Done()

	currentNode := startNode

	steps := 0
	dest := "  "

	for dest[len(dest)-1] != 'Z' {
		for _, instruction := range instructions {

			if instruction == 'L' {

				dest = currentNode.left
				currentNode = nmap[currentNode.left]
			} else {

				dest = currentNode.right
				currentNode = nmap[currentNode.right]
			}

			steps++

			if dest[len(dest)-1] == 'Z' {

				break
			}

		}

	}
	stepChan <- steps

}

func getNodeMap(input string) (nMap nodeMap, instructions string) {
	scanner := bufio.NewScanner(strings.NewReader(input))
	scanner.Scan()
	instructions = scanner.Text()
	scanner.Scan() // Skip the blank line

	nMap = make(nodeMap)

	nodeName := ""

	for scanner.Scan() {
		line := scanner.Text()
		nodeName = strings.Fields(line)[0]

		lr := strings.Split(strings.Replace(line[7:], ")", "", -1), ",")
		left := strings.TrimSpace(lr[0])
		right := strings.TrimSpace(lr[1])

		nMap[nodeName] = node{left: left, right: right}

	}

	return
}

// Functions to work out LCM
func GCD(a, b int) int {
	for b != 0 {
		t := b
		b = a % b
		a = t
	}
	return a
}

func LCM(first int, integers []int) int {
	result := first * integers[0] / GCD(first, integers[0])
	for i := 1; i < len(integers); i++ {
		result = LCM(result, []int{integers[i]})
	}
	return result
}

/*
// Thanks to good old stack overflow!
// find Least Common Multiple (LCM) via GCD
func LCM(a, b int, integers ...int) int {
	result := a * b / GCD(a, b)

	for i := 0; i < len(integers); i++ {
			result = LCM(result, integers[i])
	}

	return result
}
*/
