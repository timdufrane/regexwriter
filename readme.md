# RegexWriter
This is a simple helper to enable taking actions based on writer output. Personal use is parsing ffmpeg command-line output for video duration and completion percentage. Use it for all sorts of fun stuff!

## Example
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
		reWriter.AddMatchAction("(?i)test", func([][]byte) {
			println("Looks like this is a case-insensitive test string!")
		})

	    	// Or add an action when something doesn't match? How novel!
		reWriter.AddNonMatchAction("production", func([][]byte) {
			println("This is definitely not a production string!")
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
	This is definitely not a production string!
	**Testing second string**
	We're in production now, baby!
