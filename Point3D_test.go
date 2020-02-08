package point

import "testing"

func TestToStringPoint3D(t *testing.T) {

}
func TestSetDimensions3D(t *testing.T) {
    SetDimensions3D(1,2,3,4,5,6)
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
    if ZMin != 5 {
        t.Errorf("ZMin = %d, not 5", ZMin)
    }
    if ZMax != 6 {
        t.Errorf("ZMax = %d, not 6", ZMax)
    }
}