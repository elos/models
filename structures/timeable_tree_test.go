package structures

import (
	"math/rand"
	"testing"
	"time"
)

func RandTime() time.Time {
	return time.Date(
		rand.Intn(3000)+2,
		time.Month(rand.Intn(12)),
		rand.Intn(31),
		rand.Intn(24),
		rand.Intn(60),
		rand.Intn(60),
		rand.Intn(1000),
		time.UTC)
}

type testTimeable struct {
	s time.Time
	e time.Time
}

func (t *testTimeable) SetStartTime(st time.Time) {
	t.s = st
}

func (t *testTimeable) SetEndTime(et time.Time) {
	t.e = et
}

func (t *testTimeable) StartTime() time.Time {
	return t.s
}

func (t *testTimeable) EndTime() time.Time {
	return t.e
}

func TestTimeableTree(t *testing.T) {
	tree := new(TimeableTree)

	if !tree.Empty() {
		t.Errorf("Tree should be empty")
	}

	/*
		timeN := time.Now()

		tree.Add(&testTimeable{
			s: timeN,
			e: timeN.Add(1 * time.Hour),
		})

		tree.Add(&testTimeable{
			s: timeN.Add(2 * time.Hour),
			e: timeN.Add(3 * time.Hour),
		})

		tree.Add(&testTimeable{
			s: timeN.Add(4 * time.Hour),
			e: timeN.Add(5 * time.Hour),
		})

		log.Print(tree)
		log.Print(tree.Right)
	*/

	for i := 0; i < 3000; i++ {
		testTime := RandTime()
		testTimeable := &testTimeable{
			s: testTime,
			e: testTime.Add(time.Hour*time.Duration(rand.Intn(10)) + time.Minute*time.Duration(rand.Intn(60))),
		}
		tree.Add(testTimeable)
	}

	c := tree.Stream()

	lastTime := *new(time.Time)

	for i := 0; i < 3000; i++ {
		select {
		case timeable := <-c:
			if !lastTime.Before(timeable.StartTime()) {
				t.Errorf("Tree out of order, current start time %s, not after previous %s. Currently on iteration %d", timeable.StartTime(), lastTime, i)
			}
			lastTime = timeable.StartTime()
		case <-time.After(1 * time.Second):
			t.Fatalf("Timeout waiting on tree steram")
		}
	}
}
