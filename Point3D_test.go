package point

import (
	"math"
	"testing"
)

//epsilon definition is found in Point2D_test.go

/*
NOTE: most unit tests of Point3D are missing from this file at this time.
I am not sure exactly when I'll get around to adding them.
*/

func TestToStringPoint3D(t *testing.T) {
	p := Point3D{Point2D{1, 2}, 3}
	expected := "(1,2,3)"
	result := p.ToString()
	if result != expected {
		t.Errorf("point with x,y,z = 1,2,3 returned %s, not %s", result, expected)
	}
}

func TestDistanceFromPoint3D_1(t *testing.T) { //test case where same point
	p1 := Point3D{Point2D{1, 2}, 3}
	p2 := Point3D{Point2D{1, 2}, 3}
	expected := 0.
	result := p1.DistanceFrom(&p2)
	diff := math.Abs(result - expected)
	if diff > epsilon {
		t.Errorf("Distance from %s to %s was %g, not %g as expected",
			p1.ToString(), p2.ToString(), result, diff)
	}

	if p1.DistanceFrom(&p2) != p2.DistanceFrom(&p1) {
		t.Errorf("Distance From not commutative. Received %g and %g",
			p1.DistanceFrom(&p2), p2.DistanceFrom(&p1))
	}

}
func TestSetDimensions3D(t *testing.T) {
	SetDimensions3D(1, 2, 3, 4, 5, 6)
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
