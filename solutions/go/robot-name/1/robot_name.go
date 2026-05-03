package robotname

import (
	"fmt"
	"math/rand"
	"time"
)

// Robot is identified with a name
type Robot struct {
	name *string
}

var names []*string
var namesInitialized = false
var availableNames = 26 * 26 * 1000

// Name generates a random name for each robot in the format of two uppercase
// letters followed by three digits, such as RX837 or BC811.
func (r *Robot) Name() (string, error) {
	if !namesInitialized {
		i := 0
		names = make([]*string, availableNames)
		rand.Seed(time.Now().UnixNano())
		for n1 := 'A'; n1 <= 'Z'; n1++ {
			for n2 := 'A'; n2 <= 'Z'; n2++ {
				for n3 := 0; n3 < 1000; n3++ {
					temp := fmt.Sprintf("%c%c%03d", n1, n2, n3)
					names[i] = &temp
					i++
				}
			}
		}
		namesInitialized = true
	}

	if r.name == nil {
		if availableNames == 0 {
			return "", fmt.Errorf("used all names")
		}
		// select a random name
		i := rand.Intn(availableNames)
		r.name = names[i]

		// delete the assigned name from the list
		names[i] = names[availableNames-1] // Copy the last element to index i.
		availableNames--
	}
	return *r.name, nil
}

// Reset wipes the robot name and generates a new one.
func (r *Robot) Reset() (string, error) {
	r.name = nil
	return r.Name()
}
