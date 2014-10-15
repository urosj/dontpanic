package dontpanic

import (
	"fmt"
	"math/rand"
	"testing"
	"time"
)

type testSample struct {
	Num1  uint16
	Num2  int32
	Name  string
	Bork  map[string]int32
	Names []string
}

func init() {
	rand.Seed(time.Now().UTC().UnixNano())
}

func TestFuzz(t *testing.T) {
	ts := testSample{
		Num1: 0,
		Num2: 0,
		Name: "bla",
		Bork: map[string]int32{
			"hello": 1,
		},
		Names: []string{
			"bla", "bla", "bla",
		},
	}

	fmt.Println("Init sample: ", ts)

	fuzzed, err := Fuzz(&ts)
	if err != nil {
		t.Error(err)
	}

	fmt.Println("Fuzzed one: ", fuzzed)
}
