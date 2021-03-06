package twobucket

import "errors"

type bucket struct {
	level int
	size  int
	name  string
}

type waterPour struct {
	a                *bucket
	b                *bucket
	goal             *int
	goalBucket       *string
	otherBucketLevel *int
	steps            *int
}

// Solve given:
// - the size of bucket one
// - the size of bucket two
// - the desired number of liters to reach
// - which bucket to fill first, either bucket one or bucket two
// determines:
// - the total number of "moves" it should take to reach the desired number of liters, including the first fill
// - which bucket should end up with the desired number of liters (let's say this is bucket A) - either bucket one or bucket two
// - how many liters are left in the other bucket (bucket B)
func Solve(sizeBucketOne,
	sizeBucketTwo,
	goalAmount int,
	startBucket string) (goalBucket string, numSteps, otherBucketLevel int, e error) {
	if sizeBucketOne == 0 || sizeBucketTwo == 0 || goalAmount == 0 ||
		!(startBucket == "one" || startBucket == "two") ||
		goalAmount%gcd(sizeBucketOne, sizeBucketTwo) != 0 {
		return "", 0, 0, errors.New("invalid parameter")
	}
	wp := &waterPour{
		a:                &bucket{level: 0, size: sizeBucketOne, name: "one"},
		b:                &bucket{level: 0, size: sizeBucketTwo, name: "two"},
		goal:             &goalAmount,
		goalBucket:       &goalBucket,
		otherBucketLevel: &otherBucketLevel,
		steps:            &numSteps,
	}
	if startBucket != "one" {
		wp.a, wp.b = wp.b, wp.a
	}

	for {
		if wp.fill(wp.a, wp.b) {
			break
		}
		if wp.b.size == goalAmount {
			wp.fill(wp.b, wp.a)
			break
		}
		if wp.pour() {
			break
		}
	}

	return goalBucket, numSteps, otherBucketLevel, nil
}

func (wp *waterPour) fill(dest, other *bucket) bool {
	dest.level = dest.size
	*wp.steps++
	if dest.level == *wp.goal {
		*wp.goalBucket = dest.name
		*wp.otherBucketLevel = other.level
		return true
	}
	return false
}

func (wp *waterPour) pour() bool {
	for wp.a.level != 0 {
		delta := min(wp.a.level, wp.b.size-wp.b.level)
		wp.b.level += delta
		wp.a.level -= delta
		*wp.steps++
		if wp.a.level == *wp.goal {
			*wp.goalBucket = wp.a.name
			*wp.otherBucketLevel = wp.b.level
			return true
		}
		if wp.b.level == *wp.goal {
			*wp.goalBucket = wp.b.name
			*wp.otherBucketLevel = wp.a.level
			return true
		}
		if wp.b.level == wp.b.size {
			wp.b.level = 0
			*wp.steps++
		}
	}
	return false
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func gcd(a, b int) int {
	if b == 0 {
		return a
	}
	return gcd(b, a%b)
}
