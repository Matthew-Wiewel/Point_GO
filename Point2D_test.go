package point

import "testing"

func TestSetDimensions2D(t *testing.T) {
	SetDimensions2D(1, 2, 3, 4)
	if XMin != 1 {
		t.Errorf("XMin = %d, not 1", XMin)
	}
	if XMax != 2 {
		t.Errorf("XMax = %d, not 2", XMax)
	}
	if YMin != 3 {
		t.Errorf("YMin = %d, not 3", YMin)
	}
	if YMax != 4 {
		t.Errorf("YMax = %d, not 4", YMax)
	}
}
