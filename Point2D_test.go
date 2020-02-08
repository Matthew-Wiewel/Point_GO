package point

import "math"
import "testing"

const (
	epsilon float64 = 0.00001 //acceptable difference in these tests for floating point values
)

func TestConstructorPoint2D(t *testing.T) {
	var p1 Point2D = Point2D{2,3}
	if p1.X != 2 {
		t.Errorf("p1.X = %d, not 2", p1.X)
	}
	if p1.Y != 3 {
		t.Errorf("p1.Y = %d, not 3", p1.Y)
	}

	p2 := Point2D{4,5}
	if p2.X != 4 {
		t.Errorf("p2.X = %d, not 4", p2.X)
	}
	if p2.Y != 5 {
		t.Errorf("p2.Y = %d, not 5", p2.Y)
	}
}

func TestEqualityPoint2D(t *testing.T) {
	p1 := Point2D{5,7}
	p2 := Point2D{5,7}
	if (p1 == p2) != true {
		t.Errorf("p1 == p2 not true")
	}

	p3 := Point2D{2,2}
	p4 := Point2D{2,2}
	if (p3 == p4) != true {
		t.Errorf("p3 == p4 not true")
	}
}

func TestInequalityPoint2D(t *testing.T) {
	p1 := Point2D{8,9}
	p2 := Point2D{8,3}
	p3 := Point2D{2,9}

	if (p1 != p2) != true {
		t.Errorf("p1 != p2 not true")
	}
	if (p1 != p3) != true {
		t.Errorf("p1 != p3 not true")
	}
}

func TestToStringPoint2D(t *testing.T) {
	p := Point2D{3,4}
	expected := "(3,4)"

	if p.ToString() != expected {
		t.Errorf("p.ToString() = %s, not %s", p.ToString(), expected)
	}
}

func TestDistanceFromPoint2D(t *testing.T) {
	p1 := Point2D{0,0}
	p2 := Point2D{3,4}
	expected1 := 5.
	if p1.DistanceFrom(&p2) != p2.DistanceFrom(&p1) {
		t.Errorf("DistanceFrom is not commutative")
	}
	result1 := p1.DistanceFrom(&p2)
	difference1 := math.Abs(expected1 - result1)
	if difference1 > epsilon {
		t.Errorf("Distance from p1 to p2 is %g, not %g", difference1, expected1)
	}
	
}

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
