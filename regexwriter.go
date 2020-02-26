package regexwriter

import "regexp"

// RegexAction is a combination of a regular expression and a function
type RegexAction struct {
	// Search expression
	regex *regexp.Regexp
	// Action to take on matched/non-matched bytes
	action func([][]byte)
}

// IsMatch returns whether the regular expresion is a match for the
// given bytes
func (r RegexAction) IsMatch(b []byte) bool {
	return r.regex.Match(b)
}

// Matches returns all regular expression matches in the given bytes
func (r RegexAction) Matches(b []byte) [][][]byte {
	return r.regex.FindAllSubmatch(b, -1)
}

// PerformAction executes the defined function on the given multidimensional
// byte array
func (r RegexAction) PerformAction(bytes [][]byte) {
	r.action(bytes)
}

// CreateAction returns a RegexAction object
func CreateAction(r *regexp.Regexp, a func([][]byte)) RegexAction {
	return RegexAction{r, a}
}

// RegexWriter allows binding of matching and non-matching actions to
// a writer interface
type RegexWriter struct {
	// RawOutput contains the full output of the writer
	RawOutput       string
	matchActions    []RegexAction
	nonMatchActions []RegexAction
}

// Reset clears this RegexWriter's set members
func (re *RegexWriter) Reset() {
	re.ClearMatchActions()
	re.ClearNonMatchActions()
	re.RawOutput = ""
}

// Write is the implementation of the Writer interface
func (re RegexWriter) Write(b []byte) (n int, err error) {
	out := string(b[:])

	re.RawOutput += out

	for _, v := range re.matchActions {
		if v.IsMatch(b) {
			for _, match := range v.Matches(b) {
				v.PerformAction(match)
			}
		}
	}

	for _, v := range re.nonMatchActions {
		if !v.IsMatch(b) {
			// Since there's no match here and we're reusing the same function
			// definition, simply make the bytes look like [][]byte
			v.PerformAction([][]byte{b[:]})
		}
	}

	return len(b), nil
}

// AddMatchAction adds a match action to the RegexWriter instance
func (re *RegexWriter) AddMatchAction(pattern string, handler func([][]byte)) {
	regex := regexp.MustCompile(pattern)

	if re.matchActions == nil {
		re.matchActions = make([]RegexAction, 0)
	}

	re.matchActions = append(re.matchActions, CreateAction(regex, handler))
}

// ClearMatchActions clears the matchActions array
func (re *RegexWriter) ClearMatchActions() {
	re.matchActions = nil
}

// AddNonMatchAction adds a non-matching action to the RegexWriter instance
func (re *RegexWriter) AddNonMatchAction(pattern string, handler func([][]byte)) {
	regex := regexp.MustCompile(pattern)

	if re.nonMatchActions == nil {
		re.nonMatchActions = make([]RegexAction, 0)
	}

	re.nonMatchActions = append(re.nonMatchActions, CreateAction(regex, handler))
}

// ClearNonMatchActions clears the nonMatchActions array
func (re *RegexWriter) ClearNonMatchActions() {
	re.nonMatchActions = nil
}

// ClearRawOutput clears the raw output member
func (re *RegexWriter) ClearRawOutput() {
	re.RawOutput = ""
}
