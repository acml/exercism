// Package twofer should have a package comment that summarizes what it's about.
// https://golang.org/doc/effective_go.html#commentary
package twofer

// ShareWith should have a comment documenting it.
func ShareWith(name string) string {
	who := name
	if who == "" {
		who = "you"
	}

	return "One for " + who + ", one for me."
}
