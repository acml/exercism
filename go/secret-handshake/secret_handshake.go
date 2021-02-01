// Package secret provides a secret handsake
package secret

// Handshake given a decimal number, converts it to the appropriate sequence of
// events for a secret handshake.
func Handshake(code uint) []string {
	handshake := []string{"wink", "double blink", "close your eyes", "jump"}

	res := []string{}
	isReverse := code&(1<<4) > 0
	var j, k int
	for i := 0; i < 4; i++ {
		if isReverse {
			j = 1 << (3 - i)
			k = 3 - i
		} else {
			j = 1 << i
			k = i
		}
		if code&uint(j) > 0 {
			res = append(res, handshake[k])
		}
	}
	return res
}
