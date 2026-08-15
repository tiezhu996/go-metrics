package model

import "testing"

func TestValidSample(t *testing.T) {
	if !ValidSample(&Sample{ID: "a", Name: "cpu"}) || ValidSample(nil) || ValidSample(&Sample{}) {
		t.Fatal("ValidSample wrong")
	}
}

func TestSortSamples(t *testing.T) {
	in := []*Sample{{ID: "c"}, {ID: "a"}, {ID: "b"}}
	got := SortSamples(in)
	for i, id := range []string{"a", "b", "c"} {
		if got[i].ID != id {
			t.Fatalf("order=%v", got)
		}
	}
}

func TestBuildBucketsFresh(t *testing.T) {
	in := []*Sample{{ID: "1"}, {ID: "2"}, {ID: "3"}}
	b := BuildBuckets(in, 2)
	if len(b) != 2 {
		t.Fatalf("len=%d", len(b))
	}
	b[0][0] = &Sample{ID: "x"}
	if in[0].ID != "1" {
		t.Fatal("mutating bucket corrupted input")
	}
}

func TestMergeSummary(t *testing.T) {
	got := MergeSummary(Summary{Count: 1, Sum: 1, Max: 1, Min: 1}, Summary{Count: 1, Sum: 2, Max: 5, Min: 2})
	if got.Count != 2 || got.Sum != 3 || got.Max != 5 || got.Min != 1 {
		t.Fatalf("%+v", got)
	}
}
