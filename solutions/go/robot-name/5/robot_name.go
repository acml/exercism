package robotname

import (
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
	random = rand.New(rand.NewSource(time.Now().UnixNano()))
	used   = map[string]bool{}
)

// Name generates a random name for each robot in the format of two uppercase
// letters followed by three digits, such as RX837 or BC811.
func (r *Robot) Name() (name string, err error) {
	if r.name != "" {
		return r.name, nil
	}

	if len(used) == maxRobotNames {
		return "", fmt.Errorf("used all names")
	}

	r.name = newName()
	for used[r.name] {
		r.name = newName()
	}
	used[r.name] = true

	return r.name, nil
}

// Reset wipes the robot name and generates a new one.
func (r *Robot) Reset() (string, error) {
	r.name = ""
	return r.name, nil
}

func newName() string {
	r1 := random.Intn(26) + 'A'
	r2 := random.Intn(26) + 'A'
	num := random.Intn(1000)
	return fmt.Sprintf("%c%c%03d", r1, r2, num)
}
