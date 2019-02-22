package levenshtein

import (
	"fmt"
	"testing"
)

func TestLevenshtein(t *testing.T) {
	cases := []struct{
		S1 string
		S2 string
		Expected int
	}{
		{
			S1:"Saturday",
			S2:"Sunday",
			Expected: 3,
		},
		{
			S1:"kitten",
			S2:"sitting",
			Expected: 3,
		},
	}

	for _,td := range cases{
		t.Run(fmt.Sprintf("Levenshtein %v,%v", td.S1, td.S2), func (t *testing.T){
			result1 := Levenshtein(td.S1, td.S2)
			result2 := Levenshtein(td.S2, td.S1)

			if result1 != result2 {
				t.Fatalf("results should be symmetrical.")
			}

			if result1 != td.Expected {
				t.Fatalf("%v doesn't match expected value %v", result1, td.Expected)
			}
		})
	}
}