package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {

	n, _ := strconv.Atoi(os.Args[1])
	m := ""
	alpha := "0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"

	for n > 0 {
		d := n / 62
		r := n % 62
		s := []string{string(alpha[r]), m}
		m = strings.Join(s, "")
		n = d
	}

	fmt.Println(m)
}
