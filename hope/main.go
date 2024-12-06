package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	line := scanner.Text()

	var p, q int
	fmt.Sscan(line, &p, &q)
	for i := 1; i <= q; i++ {
		if i%p == 0 {
			fmt.Println(strings.TrimSpace(strings.Repeat("Hope ", i/p)))
		} else {
			fmt.Println(i)
		}
	}
}
