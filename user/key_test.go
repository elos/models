package user_test

import (
	"testing"

	"github.com/elos/models/user"
)

func TestRandomString(t *testing.T) {
	sizes := map[int]int{
		-3:  0,
		0:   0,
		1:   1,
		200: 200,
	}

	for key, val := range sizes {
		s := user.RandomString(key)
		if len(s) != val {
			t.Errorf("Expected random string to be size %d got %d", val, len(s))
		}
	}
}

func TestKey(t *testing.T) {
	k1 := user.NewKey()
	k2 := user.NewKey()

	if len(k1) != len(k2) || len(k1) != 64 {
		t.Errorf("Keys should be of length 64")
	}

	if k1 == k2 {
		t.Errorf("You have just experienced a mathematical anomaly, and improbability as great as your existence")
	}
}
