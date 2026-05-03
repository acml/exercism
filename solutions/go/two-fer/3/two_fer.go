// `Two-fer` or `2-fer` is short for two for one.
package twofer

/*
Given a name, return a string with the message:
One for X, one for me.

Where X is the given name.

However, if the name is missing, return the string:

One for you, one for me.
 */
func ShareWith(name string) string {
	if name == "" {
		name = "you"
	}

	return "One for " + name + ", one for me."
}
