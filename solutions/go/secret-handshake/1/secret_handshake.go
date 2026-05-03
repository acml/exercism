// Package secret provides a secret handsake
package secret

// Handshake given a decimal number, converts it to the appropriate sequence of
// events for a secret handshake.
func Handshake(code uint) []string {
	const (
		wink = 1 << iota
		doubleBlink
		closeYourEyes
		jump
		reverse
	)

	handshake := map[uint]string{
		wink:          "wink",
		doubleBlink:   "double blink",
		closeYourEyes: "close your eyes",
		jump:          "jump",
	}

	res := []string{}
	isReverse := code&reverse > 0
	var j uint
	for i := 0; i < 4; i++ {
		if isReverse {
			j = 1 << (3 - i)
		} else {
			j = 1 << i
		}
		if code&j > 0 {
			res = append(res, handshake[j])
		}
	}
	return res
}
