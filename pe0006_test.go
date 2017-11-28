package project_euler

import "testing"

func Test_PE0006(t *testing.T) {
	cases := []struct {
		Input    int
		Expected int
	}{
		{10, 2640},
		{100, 25164150},
	}
	for _, tc := range cases {
		if actual := PE0006(tc.Input); actual != tc.Expected {
			t.Errorf("expected=%v, actual=%v", tc.Expected, actual)
		}
	}
}
