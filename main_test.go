package main

// there is a lot of I/O, printing, command line parsing etc
// in this file and most of these are not interfaces so
// testing is quite difficult
// I'm trying to err on the side of having as much test
// coverage as possible even though doing things like touching
// files in unit tests is a bit dubious

import (
	"strings"
	"testing"

	"github.com/gazek/gonr/scanner"
)

func TestScan(t *testing.T) {
	text := "foo! bar?\nbaz 世界!!\tbar?\nbaz 世界 foo!\tBAR?\nbAz"
	want := map[string]int{
		"foo bar baz": 2,
		"bar baz 世界":  2,
		"baz 世界 bar":  1,
		"世界 bar baz":  1,
		"baz 世界 foo":  1,
		"世界 foo bar":  1,
	}
	s := scanner.NewScannerFromReader(strings.NewReader(text))
	countByKeys := make(map[string]int)
	scan(s, countByKeys)

	for k, v := range countByKeys {
		if v != want[k] {
			t.Errorf("%v - Expected: %v, found: %v", k, want[k], v)
		}
	}

}

func TestGetKeysByCount(t *testing.T) {
	input := map[string]int{
		"key1": 1,
		"key2": 1,
		"key3": 2,
		"key4": 5,
		"key5": 5,
		"key6": 5,
	}

	result := getKeysByCount(input)

	for c, keys := range result {
		for _, k := range keys {
			if input[k] != c {
				t.Errorf("Expected: %v, found: %v", input[k], c)
			}
		}
	}
}

func TestGetCounterKey(t *testing.T) {
	input := []string{"word1", "word2", "word3"}
	delimiter := " "
	want := strings.Join(input, delimiter)

	result := getCounterKey(input)

	if result != want {
		t.Errorf("Expected: %v, found: %v", want, result)
	}
}

func TestGetUniqueCount(t *testing.T) {
	input := map[int][]string{
		1: []string{"key1", "key2"},
		2: []string{"key3"},
		5: []string{"key4", "key5", "key6"},
	}

	want := []int{5, 2, 1}

	result := getUniqueCounts(input)

	for i, v := range result {
		if v != want[i] {
			t.Errorf("Expected: %v, found: %v", want[i], v)
		}
	}
}

func TestUpdateCounter(t *testing.T) {
	words := []string{"key1"}
	countByKeys := map[string]int{
		"key1": 1,
		"key2": 1,
		"key3": 2,
		"key4": 5,
		"key5": 5,
		"key6": 5,
	}

	want := countByKeys[words[0]] + 1

	updateCounter(words, countByKeys)

	if countByKeys[words[0]] != want {
		t.Errorf("Expected: %v, found: %v", want, countByKeys[words[0]])
	}
}

func TestOpenFile(t *testing.T) {
	// this is awkward to touch a file
	// in a test
	name := "main.go"
	path := "./" + name
	file, err := openFile(path)
	if err != nil {
		t.Errorf("Open file failed")
	}
	fi, _ := file.Stat()
	if fi.Name() != name {
		t.Errorf("Expected: %v, found: %v", name, fi.Name())
	}
}

func TestOpenFileFail(t *testing.T) {
	name := "fake.go"
	path := "./" + name
	_, err := openFile(path)
	if err == nil {
		t.Errorf("failed")
	}
}

func TestCheckFileOpenErr(t *testing.T) {
	// this is a silly test
	// I could use panic and then use recover to
	// make this more testable but I want to provide
	// an error message to the user, not a stack trace
	// that's why I'm calling os.Exit
	checkFileOpenErr(nil)
}
