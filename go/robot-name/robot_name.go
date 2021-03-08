package robotname

import (
	"errors"
	"math/rand"
	"strings"
	"sync"
	"time"
)

// Robot is identified with a name
type Robot struct {
	name string
}

const maxRobotNames = 26 * 26 * 10 * 10 * 10

var names = map[int]struct{}{}
var availableNames = maxRobotNames

// Name generates a random name for each robot in the format of two uppercase
// letters followed by three digits, such as RX837 or BC811.
func (r *Robot) Name() (string, error) {
	if r.name == "" {
		if availableNames == 0 {
			return "", errors.New("used all names")
		}
		r.name = randRobotName()
	}
	return r.name, nil
}

// Reset wipes the robot name and generates a new one.
func (r *Robot) Reset() (string, error) {
	r.name = ""
	return r.Name()
}

func randRobotName() string {

	once := sync.Once{}
	once.Do(func() { rand.Seed(time.Now().UnixNano()) })

	var idx int
	for {
		idx = rand.Intn(maxRobotNames)
		if _, ok := names[idx]; ok {
			continue
		}
		break
	}

	names[idx] = struct{}{}
	availableNames--

	sb := strings.Builder{}
	sb.Grow(5)

	sb.WriteRune('A' + rune(idx/(26*1000)))
	sb.WriteRune('A' + rune((idx%(26*1000))/1000))
	sb.WriteRune('0' + rune((idx%1000)/100))
	sb.WriteRune('0' + rune(((idx%1000)%100)/10))
	sb.WriteRune('0' + rune((idx%1000)%10))
	return sb.String()
}
