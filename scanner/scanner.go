package scanner

import (
	"bufio"
	"io"
	"os"
	"unicode"
)

// Scanner reads words from the input source, removes punctuation and lowercases the word
type Scanner struct {
	scanner  *bufio.Scanner
	fileName string
}

// NewScanner creates a new scanner from a file path an error is returned if the path is invalid
func NewScanner(f *os.File) *Scanner {
	// create and return the Scanner struct
	return NewScannerFromReader(f)
}

// NewScannerFromReader is only public to facilitate testing
func NewScannerFromReader(r io.Reader) *Scanner {
	scanner := bufio.NewScanner(r)
	scanner.Split(bufio.ScanWords)
	return &Scanner{scanner: scanner}
}

// Next reads the returns the next word from the input source
func (s *Scanner) Next() (word string, eof bool) {
	// advance the scanner
	eof = !s.scanner.Scan()
	// check for end of file
	if eof {
		return "", true
	}
	// collect the new word
	rawWord := s.scanner.Text()
	// lowercase and remove punctuation
	cleanWord := cleanWord(rawWord)
	// if the cleaning resulted in an empty string
	// then scan the next word by recursing
	if cleanWord == "" {
		return s.Next()
	}
	// return the cleaned word
	return cleanWord, false
}

// cleanWord takes a string, lowercases it and removes punctuation
func cleanWord(word string) string {
	var cleaned string
	for _, r := range []rune(word) {
		// this will confuse a hyphen with a dash or double dash
		// a dash or double dash is used for narative effect while
		// a hyphon joins two words or splits words across a carriage return
		// We would like to split on dash and double dash but drop a hyphen
		// and join the words on either side of the hyphen
		if unicode.IsPunct(r) {
			continue
		}
		cleaned += string(unicode.ToLower(r))
	}
	return cleaned
}
