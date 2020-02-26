# RegexWriter
This is a simple helper to enable taking actions based on writer output. Personal use is parsing ffmpeg command-line output for video duration and completion percentage. Use it for all sorts of fun stuff!

## Getting it
	go get -u github.com/timdufrane/regexwriter
	
## Using it
	package main

	import (
		"github.com/timdufrane/regexwriter"
		"strings"
	)

	func main() {
		reWriter := regexwriter.RegexWriter{}

		// Add one match action
		reWriter.AddMatchAction("production", func([][]byte) {
			println("We're in production now, baby!")
		})

		// Maybe add another!
		reWriter.AddMatchAction("(?i)(te)(st)", func(b [][]byte) {
			println("Looks like this is a case-insensitive test string!")
			println("First part being", string(b[1]), "second part being", string(b[2]))
		})

		// Or add an action when something doesn't match? How novel!
		reWriter.AddNonMatchAction("production", func(b [][]byte) {
			println("This is definitely not a production string!")
			println("In fact, here's the string itself:", string(b[0]))
		})

		println("**Testing first string**")
		reader := strings.NewReader("This is a TEst string")

		reader.WriteTo(reWriter)

		println("**Testing second string**")
		reader = strings.NewReader("This is a production string")

		reader.WriteTo(reWriter)
	}
	
### Result
	**Testing first string**
	Looks like this is a case-insensitive test string!
	First part being TE second part being st
	This is definitely not a production string!
	In fact, here's the string itself: This is a TEst string
	**Testing second string**
	We're in production now, baby!
