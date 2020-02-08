package point

import "testing"

func TestSetDimensions(t *testing.T) {
	SetDimensions(7, 2, 3, 4)
	if XMin != 1 {
		t.Errorf("Point.XMin = %d, not 1", XMin)
	}
}
