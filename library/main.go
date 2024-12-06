package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	n, _ := strconv.Atoi(scanner.Text())
	var books = make(map[int]string)
	for i := 0; i < n; i++ {
		var isbn int
		var command, title string

		scanner.Scan()
		inputs := strings.Fields(scanner.Text())

		command = inputs[0]
		isbn, _ = strconv.Atoi(inputs[1])
		title = strings.Join(inputs[2:], " ")

		if command == "ADD" {
			books[isbn] = title
		} else if command == "REMOVE" {
			delete(books, isbn)
		}
	}
	isbns := make([]int, 0, len(books))
	for key := range books {
		isbns = append(isbns, key)
	}
	sort.Ints(isbns)
	sort.SliceStable(isbns, func(i, j int) bool {
		return books[isbns[i]] < books[isbns[j]]
	})

	for _, isbn := range isbns {
		fmt.Println(isbn)
	}
}
