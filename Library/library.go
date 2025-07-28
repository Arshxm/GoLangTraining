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

	books := make(map[string]string)
	for i := 0; i < n; i++ {
		scanner.Scan()
		command := scanner.Text()
		arrCommand := strings.Split(command, " ")

		if arrCommand[0] == "ADD" {
			books[arrCommand[1]] = strings.Join(arrCommand[2:], " ")
		} else if arrCommand[0] == "REMOVE" {
			if _, exists := books[arrCommand[1]]; exists {
				delete(books, arrCommand[1])
			}
		}
	}

	type BookPair struct {
		ID    string
		Title string
	}

	var bookPairs []BookPair
	for id, title := range books {
		bookPairs = append(bookPairs, BookPair{ID: id, Title: title})
	}

	sort.Slice(bookPairs, func(i, j int) bool {
		if bookPairs[i].Title == bookPairs[j].Title {
			id1, _ := strconv.Atoi(bookPairs[i].ID)
			id2, _ := strconv.Atoi(bookPairs[j].ID)
			return id1 < id2
		}
		return bookPairs[i].Title < bookPairs[j].Title
	})

	for _, pair := range bookPairs {
		fmt.Println(pair.ID)
	}
}
