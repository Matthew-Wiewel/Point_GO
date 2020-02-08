package point

import (
	"math"
	"testing"
)

const (
	epsilon float64 = 0.00001 //acceptable difference in these tests for floating point values
)

func TestConstructorPoint2D(t *testing.T) {
	var p1 Point2D = Point2D{2, 3}
	if p1.X != 2 {
		t.Errorf("p1.X = %d, not 2", p1.X)
	}
	if p1.Y != 3 {
		t.Errorf("p1.Y = %d, not 3", p1.Y)
	}

	p2 := Point2D{4, 5}
	if p2.X != 4 {
		t.Errorf("p2.X = %d, not 4", p2.X)
	}
	if p2.Y != 5 {
		t.Errorf("p2.Y = %d, not 5", p2.Y)
	}
}

func TestEqualityPoint2D(t *testing.T) {
	p1 := Point2D{5, 7}
	p2 := Point2D{5, 7}
	if (p1 == p2) != true {
		t.Errorf("p1 == p2 not true")
	}

	p3 := Point2D{2, 2}
	p4 := Point2D{2, 2}
	if (p3 == p4) != true {
		t.Errorf("p3 == p4 not true")
	}
}

func TestInequalityPoint2D(t *testing.T) {
	p1 := Point2D{8, 9}
	p2 := Point2D{8, 3}
	p3 := Point2D{2, 9}

	if (p1 != p2) != true {
		t.Errorf("p1 != p2 not true")
	}
	if (p1 != p3) != true {
		t.Errorf("p1 != p3 not true")
	}
}

func TestToStringPoint2D(t *testing.T) {
	p := Point2D{3, 4}
	expected := "(3,4)"

	if p.ToString() != expected {
		t.Errorf("p.ToString() = %s, not %s", p.ToString(), expected)
	}
}

func TestDistanceFromPoint2D_1(t *testing.T) {
	p1 := Point2D{1, 1}
	p2 := Point2D{4, 5}
	expected := 5.
	if p1.DistanceFrom(&p2) != p2.DistanceFrom(&p1) {
		t.Errorf("DistanceFrom is not commutative")
	}
	result := p1.DistanceFrom(&p2)
	difference := math.Abs(expected - result)
	if difference > epsilon {
		t.Errorf("Distance from p1 to p2 is %g, not %g", result, expected)
	}
}

func TestDistanceFromPoint2D_2(t *testing.T) {
	p1 := Point2D{-7, -4}
	p2 := Point2D{17, 7}
	expected := 26.400758
	if p1.DistanceFrom(&p2) != p2.DistanceFrom(&p1) {
		t.Errorf("DistanceFrom is not commutative")
	}
	result := p1.DistanceFrom(&p2)
	difference := math.Abs(expected - result)
	if difference > epsilon {
		t.Errorf("Distance from p1 to p2 is %g, not %g", result, expected)
	}
}

func TestDistanceFromPoint2D_3(t *testing.T) {
	p1 := Point2D{-5, -10}
	p2 := Point2D{-3, -4}
	expected := 6.324555
	if p1.DistanceFrom(&p2) != p2.DistanceFrom(&p1) {
		t.Errorf("DistanceFrom is not commutative")
	}
	result := p1.DistanceFrom(&p2)
	difference := math.Abs(expected - result)
	if difference > epsilon {
		t.Errorf("Distance from p1 to p2 is %g, not %g", result, expected)
	}
}

func TestDistanceFromPoint2D_4(t *testing.T) {
	p1 := Point2D{3, -8}
	p2 := Point2D{3, -8}
	expected := 0.
	if p1.DistanceFrom(&p2) != p2.DistanceFrom(&p1) {
		t.Errorf("DistanceFrom is not commutative")
	}
	result := p1.DistanceFrom(&p2)
	difference := math.Abs(expected - result)
	if difference > epsilon {
		t.Errorf("Distance from p1 to p2 is %g, not %g", result, expected)
	}
}

func TestDistanceFromOriginPoint2D_1(t *testing.T) {
	p := Point2D{0, 0}
	expected := 0.
	result := p.DistanceFromOrigin()
	difference := math.Abs(result - expected)
	if difference > epsilon {
		t.Errorf("Distance from p to origin is %g, not %g", p.DistanceFromOrigin(), expected)
	}
}

func TestDistanceFromOriginPoint2D_2(t *testing.T) {
	p := Point2D{3, 4}
	expected := 5.
	result := p.DistanceFromOrigin()
	difference := math.Abs(result - expected)
	if difference > epsilon {
		t.Errorf("Distance from p to origin is %g, not %g", p.DistanceFromOrigin(), expected)
	}
}

func TestDistanceFromOriginPoint2D_3(t *testing.T) {
	p := Point2D{3, -9}
	expected := 9.486833
	result := p.DistanceFromOrigin()
	difference := math.Abs(result - expected)
	if difference > epsilon {
		t.Errorf("Distance from p to origin is %g, not %g", p.DistanceFromOrigin(), expected)
	}
}

func TestGetQuadrantPoint2D_1(t *testing.T) {
	p := Point2D{1,3}
	expected := 1
	result := p.GetQuadrant()
	if result != expected {
		t.Errorf("%s is in %d, not %d", p.ToString(), result, expected)
	}
}

func TestGetQuadrantPoint2D_2(t *testing.T) {
	p := Point2D{-1,3}
	expected := 2
	result := p.GetQuadrant()
	if result != expected {
		t.Errorf("%s is in %d, not %d", p.ToString(), result, expected)
	}
}

func TestGetQuadrantPoint2D_3(t *testing.T) {
	p := Point2D{-1,-3}
	expected := 3
	result := p.GetQuadrant()
	if result != expected {
		t.Errorf("%s is in %d, not %d", p.ToString(), result, expected)
	}
}

func TestGetQuadrantPoint2D_4(t *testing.T) {
	p := Point2D{1,-3}
	expected := 4
	result := p.GetQuadrant()
	if result != expected {
		t.Errorf("%s is in %d, not %d", p.ToString(), result, expected)
	}
}

func TestGetQuadrantPoint2D_5(t *testing.T) {
	p := Point2D{0,3}
	expected := -1
	result := p.GetQuadrant()
	if result != expected {
		t.Errorf("%s is in %d, not %d", p.ToString(), result, expected)
	}
}

func TestGetQuadrantPoint2D_6(t *testing.T) {
	p := Point2D{1,0}
	expected := -1
	result := p.GetQuadrant()
	if result != expected {
		t.Errorf("%s is in %d, not %d", p.ToString(), result, expected)
	}
}

func TestGetQuadrantPoint2D_7(t *testing.T) {
	p := Point2D{0,0}
	expected := -1
	result := p.GetQuadrant()
	if result != expected {
		t.Errorf("%s is in %d, not %d", p.ToString(), result, expected)
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
