package cmd

import (
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestFormat(t *testing.T) {
	lines := map[string]Line{
		"this is a pen.": {
			Text: "this is a pen.",
			Num:  1,
		},
		"This is wong": {
			Text: "This is wong",
			Num:  2,
		},
		"this is a apple.": {
			Text: "this is a apple.",
			Num:  3,
		},
	}

	resp := Response{
		Matches: []Match{
			{
				Message: "This sentence does not start with an uppercase letter.",
				Offset:  0,
				Length:  4,
				Context: Context{
					Text:   "this is a pen.",
					Offset: 0,
					Length: 4,
				},
			},
			{
				Message: "Possible spelling mistake found.",
				Offset:  8,
				Length:  4,
				Context: Context{
					Text:   "This is wong",
					Offset: 8,
					Length: 4,
				},
			},
			{
				Message: "This sentence does not start with an uppercase letter.",
				Offset:  0,
				Length:  4,
				Context: Context{
					Text:   "this is a apple.",
					Offset: 0,
					Length: 4,
				},
			},
		},
	}

	got := format("test", lines, resp)
	want := []string{
		"test:1:0: This sentence does not start with an uppercase letter.",
		"test:3:0: This sentence does not start with an uppercase letter.",
		"test:2:8: Possible spelling mistake found.",
	}

	if diff := cmp.Diff(want, got); diff != "" {
		t.Errorf("result mismatch (-want +got):\n%s", diff)
	}
}
