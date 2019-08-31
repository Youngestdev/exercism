// https://golang.org/doc/effective_go.html#commentary
// package twofer implements the "One for [you], one for me."

package twofer

// ShareWith should have a comment documenting it.
// ShareWith function takes a name which is a string to substitue the [you] above.

func ShareWith(name string) string {		
	if name != "" {
		return "One for " + name+ ", one for me."
	} else {
		return "One for you, one for me."
	}
}
