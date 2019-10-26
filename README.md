# gonr

## How to build it
`go build main.go`

## Usage
### To call it using file paths
Add as many paths as you want, separated by spaces.
`./main.exe <path1> [<path2> ...]`

### To call it using stdin
You must use the -stdin flag and the text must be piped to the function.
If you use the -stdin flag without piping, it will block and take content from stdin but it will not scan the text.
`cat <path to file> | ./main.exe -stdin`

### To call it using both stdin and file paths
You must use the -stdin flag and the text must be piped to the function.
The -stdin flag must come **before** the file paths.
Again, if you use the -stdin flag without piping, it will block and take content from stdin but it will not scan the text.
`cat <path to file> | ./main.exe -stdin path1 [path2 ...]`

## Given more time
I would probably explore using goroutines to add concurrency to the program. The easily implementable approach would be to just run each file on its own goroutine. This would also involve wrapping the sequence count data structure into a struct with a mutex. A more involved concurrency approach might involve sending similarly sized chuncks of the files to their own gorutines. This handling would get a bit involved because this would necessarily require some overlap of the chuncks to scan all the sequences. Also, the starting point of most of the chunks would likely fall in the middle of a word and this would take some effort to handle.

I would also like to try to refactor much of the code in main.go to see if there is a nicer way to structure it to make it more testable. There is a lot of I/O and printing which does not lend itself well to unit testing...not to mention the command line argument and stdin handling.

I would also like to explore the use of a structure similar to a trie but using words instead of letters. There would probably be some efficiency improvement in finding the most common seqences.

## Known bugs
It will confuse a hyphen with a dash or double dash. A dash or double dash is used for narative effect while a hyphon joins two words or splits words across a carriage return. We would like to split on dash and double dash but drop a hyphen
and join the words on either side of the hyphen but that is not currently handled.

I don't think this is a bug but I noticed I got a much lower count when I ran it on the Moby Dick file I found. I found the same top 3 sequences as in your example though. I was concerned about my low count and I manually stepped through my file with find and my result seems reasonable. Then I was concerned that I misread or misunderstood and maybe you were asking for subsequences but that would almost certainly resulted in a top sequence of "the the the".

## Extra Credit
I did attempt 2 of the extra credits. The code handles unicode charaters.  This can be seen in the unit tests.
Also, I tested it against 288MB (mobydick.txt psted into the same file a bunch of times) file as well as made it open and parse mobydick.txt 1000 times. The memory usage stayed below 20MB for these. This may not represent real world cases too well since the 3-word sequences will all have been seen after the first time through and only the count will increase.