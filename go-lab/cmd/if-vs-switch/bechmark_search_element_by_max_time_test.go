package test

// run this with: go test -bench="BenchmarkSearchMax*" -benchmem
// result:
// goos: darwin
// goarch: amd64
// BenchmarkSearchMaxOneLoop-4   	 2000000	       694 ns/op	     432 B/op	       6 allocs/op
// BenchmarkSearchMaxTwoLoop-4   	 5000000	       341 ns/op	     144 B/op	       2 allocs/op

import (
	"testing"
	"time"
)

const (
	DateLayout = "2006-01-02"
)

type Model struct {
	ID      int
	Name    string
	Created time.Time
}

func initModel() []Model {
	models := make([]Model, 6)

	t, _ := time.Parse(DateLayout, "2015-01-02")
	models[0] = Model{
		ID:      1,
		Name:    "2015-1",
		Created: t,
	}
	models[1] = Model{
		ID:      2,
		Name:    "2015-2",
		Created: t,
	}

	t, _ = time.Parse(DateLayout, "2016-01-02")
	models[2] = Model{
		ID:      3,
		Name:    "2016-1",
		Created: t,
	}
	models[3] = Model{
		ID:      4,
		Name:    "2016-2",
		Created: t,
	}

	t, _ = time.Parse(DateLayout, "2017-01-02")
	models[4] = Model{
		ID:      5,
		Name:    "2017-1",
		Created: t,
	}
	models[5] = Model{
		ID:      6,
		Name:    "2017-2",
		Created: t,
	}

	return models
}

func searchMaxTwoLoops(ins []Model) (outs []Model) {
	latestDate := time.Time{}
	for _, in := range ins {
		if in.Created.After(latestDate) {
			latestDate = in.Created
		}
	}

	for _, in := range ins {
		if in.Created.Equal(latestDate) {
			outs = append(outs, in)
		}
	}

	return
}

func searchMaxOneLoop(ins []Model) (result []Model) {
	latestDate := time.Time{}
	for _, in := range ins {
		if in.Created.After(latestDate) {
			latestDate = in.Created
			result = nil
		}
		if in.Created.Equal(latestDate) {
			result = append(result, in)
		}
	}

	return
}

func TestSearchMax(t *testing.T) {
	inputs := initModel()

	result1 := searchMaxOneLoop(inputs)
	if len(result1) != 2 {
		t.Errorf("size should be 2 but got %d", len(result1))
	}

	for _, r := range result1 {
		if r.ID < 5 {
			t.Errorf("id should be 5 or 6 but got", r.ID)
		}
	}

	result2 := searchMaxOneLoop(inputs)
	if len(result2) != 2 {
		t.Errorf("size should be 2 but got %d", len(result1))
	}

	for _, r := range result2 {
		if r.ID < 5 {
			t.Errorf("id should be 5 or 6 but got", r.ID)
		}
	}
}

func BenchmarkSearchMaxOneLoop(b *testing.B) {
	inputs := initModel()
	for n := 0; n < b.N; n++ {
		searchMaxOneLoop(inputs)
	}
}

func BenchmarkSearchMaxTwoLoop(b *testing.B) {
	inputs := initModel()
	for n := 0; n < b.N; n++ {
		searchMaxTwoLoops(inputs)
	}
}
