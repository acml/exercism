package robotname

import (
	"errors"
	"fmt"
	"math/rand"
	"time"
)

// Robot is identified with a name
type Robot struct {
	name string
}

const maxRobotNames = 26 * 26 * 10 * 10 * 10

var (
	random         = rand.New(rand.NewSource(time.Now().UnixNano()))
	names          = map[string]struct{}{}
	availableNames = maxRobotNames
)

// Name generates a random name for each robot in the format of two uppercase
// letters followed by three digits, such as RX837 or BC811.
func (r *Robot) Name() (name string, err error) {
	if r.name != "" {
		return r.name, nil
	}

	if availableNames == 0 {
		return "", errors.New("used all names")
	}

	for {
		r1 := random.Intn(26) + 'A'
		r2 := random.Intn(26) + 'A'
		num := random.Intn(1000)
		r.name = fmt.Sprintf("%c%c%03d", r1, r2, num)
		if _, ok := names[r.name]; ok {
			continue
		}
		break
	}
	names[r.name] = struct{}{}
	availableNames--

	return r.name, err
}

// Reset wipes the robot name and generates a new one.
func (r *Robot) Reset() (string, error) {
	r.name = ""
	return r.Name()
}
