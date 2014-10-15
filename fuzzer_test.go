package dontpanic

import (
	"encoding/json"
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

type jsonStrunct struct {
	Person  map[string]string
	Age     int
	Hobbies []string
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

	t.Logf("Init sample: ", ts)

	fuzzed, err := Fuzz(&ts)
	if err != nil {
		t.Error(err)
	}

	t.Logf("Fuzzed one: ", fuzzed)
}

func TestJson(t *testing.T) {
	var data jsonStrunct
	jsonData := `{"Person":{"name":"John", "surname":"Doe"}, "Age":42, "Hobbies":["hacking", "running", "painting"]}`
	err := json.Unmarshal([]byte(jsonData), &data)
	if err != nil {
		t.Error(err)
	}
	t.Logf("Init sample: ", data)

	fuzzed, err := Fuzz(&data)
	if err != nil {
		t.Error(err)
	}

	t.Logf("Fuzzed one: ", fuzzed)
}
