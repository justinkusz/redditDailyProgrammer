package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	progs, _ := reader.ReadString('\n')
	progs = strings.Trim(progs, "\n\r")
	input, _ := reader.ReadString('\n')
	moves := strings.Split(input, ",")
	indeces := make([]int, len(progs))

	for i := 0; i < len(indeces); i++ {
		indeces[i] = i
	}
	tmp := make([]int, len(progs))
	copy(tmp, indeces)

	for _, move := range moves {
		tmp = parseMove(tmp, move)
	}

	retval := ""
	for i := 0; i < len(tmp); i++ {
		retval += progs[tmp[i] : tmp[i]+1]
	}
	fmt.Println(retval)
}

func parseMove(indeces []int, move string) []int {
	op := move[0:1]
	switch op {
	case "s":
		n, _ := strconv.Atoi(move[1:])
		return spin(indeces, n)
	case "x":
		pos := strings.Split(move[1:], "/")
		a, _ := strconv.Atoi(pos[0])
		b, _ := strconv.Atoi(pos[1])
		return exch(indeces, a, b)
	case "p":
		pos := strings.Split(move[1:], "/")
		a, _ := strconv.Atoi(pos[0])
		b, _ := strconv.Atoi(pos[1])
		return part(indeces, a, b)
	}
	return indeces
}

func spin(indeces []int, n int) []int {
	return append(indeces[len(indeces)-n:], indeces[:len(indeces)-n]...)
}

func exch(indeces []int, a int, b int) []int {
	tmp := make([]int, len(indeces))
	copy(tmp, indeces)
	tmp[a] = indeces[b]
	tmp[b] = indeces[a]
	return tmp
}

func part(indeces []int, a int, b int) []int {
	for i := 0; i < len(indeces); i++ {
		if indeces[i] == a {
			indeces[i] = b
			continue
		}
		if indeces[i] == b {
			indeces[i] = a
			continue
		}
	}
	return indeces
}
