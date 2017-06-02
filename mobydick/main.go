package main

import (
	"bufio"
	"fmt"
	"log"
	"net/http"
	"strings"
)

func main() {
	// get the book Moby Dick
	res, err := http.Get("http://www.gutenberg.org/files/2701/old/moby10b.txt")
	if err != nil {
		log.Fatal(err)
	}

	//scan the page
	scanner := bufio.NewScanner(res.Body)
	defer res.Body.Close()

	//set the split function for the scanning operation
	scanner.Split(bufio.ScanWords)
	buckets := make(map[string]int)

	for scanner.Scan() {
		word := ToString(scanner.Text())
		buckets[word]++
	}

	fmt.Println(buckets)
}

//ToString ... function to force scanner.Text() to be a string
func ToString(word string) string {
	word = strings.Replace(word, ".", " ", -1)
	word = strings.Replace(word, ",", " ", -1)
	word = strings.Replace(word, "!", " ", -1)
	word = strings.Replace(word, "-", " ", -1)
	word = strings.ToLower(word)
	return word
}
