package main

import (
	"flag"
	"fmt"
	"os"
	"sort"
	"strings"

	"github.com/gazek/gonr/scanner"
)

// limit the number of sequence counts printed
const resultLimit = 100

func main() {
	// collect command line args
	filePaths, useStdin := parseFlags()
	// initialize the structure to track the counts
	countByKeys := make(map[string]int)
	// scan content from file paths
	for _, path := range filePaths {
		// open the file
		f, err := openFile(path)
		checkFileOpenErr(err)
		// defer close of file
		defer f.Close()
		// create a new scanner for the file
		s := scanner.NewScanner(f)
		// scan the file content
		scan(s, countByKeys)
	}
	// scan content from stdin, if required
	if useStdin {
		// grab the stdin file pointer
		stdin := os.Stdin
		// create a new scanner for the file
		s := scanner.NewScanner(stdin)
		// scan the file content
		scan(s, countByKeys)
	}
	// print the result
	sortAndPrint(countByKeys)
}

func openFile(path string) (*os.File, error) {
	// attempt to open the file at the provided path
	f, err := os.Open(path)
	// check for errors
	if err != nil {
		return nil, err
	}
	return f, nil
}

func checkFileOpenErr(err error) {
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func parseFlags() (files []string, useStdin bool) {
	// parses command line flag and the trailing list of paths
	// I couldn't find a good way to determine if there was content waiting on stdin
	// so I am requiring a command line flag to indicate that we should look for content on stdin
	stdinFlag := flag.Bool("stdin", false, "specifies that stdin should be used for input")
	flag.Parse()
	return flag.Args(), *stdinFlag
}

func scan(s *scanner.Scanner, countByKeys map[string]int) {
	// current triplet of words
	var words []string
	// the newly fetched word
	var word string
	// end of file
	var eof bool
	// scan the first 2 words
	for i := 0; i < 2; i++ {
		word, eof = s.Next()
		words = append(words, word)
	}
	// scan until EOF
	for {
		// get the next word
		word, eof = s.Next()
		// if EOF, break
		if eof {
			break
		}
		// add the new word to the slice
		words = append(words, word)
		// update the count for the current sequence
		updateCounter(words, countByKeys)
		// drop the first word from the sequence
		words = words[1:]
	}
}

func sortAndPrint(countByKeys map[string]int) {
	// sort the keys by their count
	keysByCount := getKeysByCount(countByKeys)
	// get the unique sequence counts sorted in descending order
	counts := getUniqueCounts(keysByCount)
	// print the results
	printResult(counts, keysByCount)
}

func printResult(counts []int, keysByCount map[int][]string) {
	if len(counts) == 0 {
		fmt.Println("There were not enough words in the input, no 3-word sequences found")
	}
	// count the numberof results printed
	count := 0
	// print the results up to the resultLimit
	// the sequence are printed by highest count
	// within a common count, the sequences are printed
	// in the order they were encountered in the text
	for _, c := range counts {
		for _, k := range keysByCount[c] {
			fmt.Printf("%v - %v\n", k, c)
			count++
			// check if limit is reached and return
			if count == resultLimit {
				return
			}
		}
	}
}

func getKeysByCount(countByKeys map[string]int) map[int][]string {
	// this isn't all that efficient but when I tried to
	// build up this data during the scanning, it was taking
	// up a lot of memory for large files and I didn't have time
	// to debug it

	// build the map
	keysByCount := make(map[int][]string)
	// populate he map
	for k, c := range countByKeys {
		keysByCount[c] = append(keysByCount[c], k)
	}
	return keysByCount
}

func getUniqueCounts(keysByCount map[int][]string) []int {
	// collect the unique sequence counts
	counts := make([]int, 0, 0)
	for c := range keysByCount {
		counts = append(counts, c)
	}
	// sort the counts in descending order
	sort.Slice(counts, func(i, j int) bool { return counts[i] > counts[j] })
	return counts
}

func getCounterKey(words []string) string {
	// build the key usedfor tracking sequence count
	return strings.Join(words, " ")
}

func updateCounter(words []string, countByKeys map[string]int) {
	// get the map key name
	key := getCounterKey(words)
	// increment the count
	countByKeys[key]++
}
