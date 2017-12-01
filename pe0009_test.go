package project_euler

import "testing"

func Test_PE0009(t *testing.T) {
	expected := 31875000
	if actual := PE0009(); actual != expected {
		t.Errorf("expected=%v, actual=%v", expected, actual)
	}
}