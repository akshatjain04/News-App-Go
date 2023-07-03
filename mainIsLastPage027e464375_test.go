// Test generated by RoostGPT for test go-test using AI Model gpt

package main

import (
	"testing"
)

type Search struct {
	NextPage   int
	TotalPages int
}

func (s *Search) IsLastPage() bool {
	return s.NextPage >= s.TotalPages
}

func TestIsLastPage027e464375(t *testing.T) {
	testCases := []struct {
		name  string
		input Search
		want  bool
	}{
		{
			name: "Test Case 1: NextPage is less than TotalPages",
			input: Search{
				NextPage:   3,
				TotalPages: 5,
			},
			want: false,
		},
		{
			name: "Test Case 2: NextPage is equal to TotalPages",
			input: Search{
				NextPage:   5,
				TotalPages: 5,
			},
			want: true,
		},
		{
			name: "Test Case 3: NextPage is greater than TotalPages",
			input: Search{
				NextPage:   6,
				TotalPages: 5,
			},
			want: true,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			if got := tc.input.IsLastPage(); got != tc.want {
				t.Errorf("IsLastPage() = %v, want %v", got, tc.want)
			}
		})
	}
}
