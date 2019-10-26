package scanner

import (
	"os"
	"strings"
	"testing"
)

func TestCleanWord(t *testing.T) {
	type test struct {
		input string
		want  string
	}

	tests := []test{
		{input: "Foo", want: "foo"},
		{input: "sHouldn't", want: "shouldnt"},
		{input: "(foo)[bar]", want: "foobar"},
		{input: "(foo)[bar]", want: "foobar"},
		{input: "世界!!", want: "世界"},
	}

	for _, v := range tests {
		if result := cleanWord(v.input); result != v.want {
			t.Errorf("Got %v, wanted %v", result, v.want)
		}
	}
}

func TestNewScannerFromReader(t *testing.T) {
	text := "foo! (bar) [baz]"
	tests := strings.Split(text, " ")

	s := NewScannerFromReader(strings.NewReader(text))

	// we'll test this by making sure that the created
	// scanner in the Scanner struct returns the words from
	// the sample string
	for _, v := range tests {
		s.scanner.Scan()
		if result := s.scanner.Text(); result != v {
			t.Errorf("Got %v, wanted %v", result, v)
		}
	}
}

func TestNext(t *testing.T) {
	// again, at the expense of code duplication,
	// I will this test seperate
	text := "foo! (bar) ... [baz]"
	tests := []string{"foo", "bar", "baz"}

	s := NewScannerFromReader(strings.NewReader(text))

	// we'll test this by making sure that the created
	// scanner in the Scanner struct returns the words from
	// the sample string
	for _, v := range tests {
		if result, _ := s.Next(); result != v {
			t.Errorf("Got %v, wanted %v", result, v)
		}
	}
	// check EOF
	if _, eof := s.Next(); eof != true {
		t.Errorf("Failed tgo find EOF")
	}
}

func TestNewScanner(t *testing.T) {
	// this is awkward to test since it involves file I/O
	// I'm just going to open the source file and make sure
	// I see "package" as the first word scanned
	path := "./scanner_test.go"
	f, err := os.Open(path)
	if err != nil {
		t.Errorf("failed to open file for test: %v", path)
	}
	want := "package"
	s := NewScanner(f)

	// we'll test this by making sure that the created
	// scanner in the Scanner struct returns the words from
	// the sample string
	s.scanner.Scan()
	if result := s.scanner.Text(); result != want {
		t.Errorf("Got %v, wanted %v", result, want)
	}
}
