// pg 99
// write a program to report the frequency of each word in an input text file
// call input.Split(bufio.ScanWords) before the first call to Scan
// to break the input text into words instead of lines

package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func main() {
	count := make(map[string]int)

	file, err := os.Open("text.txt")
	check(err)
	defer file.Close()

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanWords)

	for scanner.Scan() {
		word := scanner.Text()
		count[word]++
	}
	fmt.Println(count)
}

func check(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
