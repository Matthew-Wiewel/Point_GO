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

func TestEqualityPoint2D(t *testing.T) { //a test of my own to get a better idea of Go's comparison of structs
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

func TestDistanceFromPoint2D_1(t *testing.T) { //test when result is a whole number
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

func TestDistanceFromPoint2D_2(t *testing.T) { //test when one value is negative and the other positive
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

func TestDistanceFromPoint2D_3(t *testing.T) { //test when both values are negative
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

func TestDistanceFromPoint2D_4(t *testing.T) { //test when distance is 0
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

func TestDistanceFromOriginPoint2D_1(t *testing.T) { //test when is origin
	p := Point2D{0, 0}
	expected := 0.
	result := p.DistanceFromOrigin()
	difference := math.Abs(result - expected)
	if difference > epsilon {
		t.Errorf("Distance from p to origin is %g, not %g", p.DistanceFromOrigin(), expected)
	}
}

func TestDistanceFromOriginPoint2D_2(t *testing.T) { //test when both x and y are positive
	p := Point2D{3, 4}
	expected := 5.
	result := p.DistanceFromOrigin()
	difference := math.Abs(result - expected)
	if difference > epsilon {
		t.Errorf("Distance from p to origin is %g, not %g", p.DistanceFromOrigin(), expected)
	}
}

func TestDistanceFromOriginPoint2D_3(t *testing.T) { //test when there is a mix of positive and negative
	p := Point2D{3, -9}
	expected := 9.486833
	result := p.DistanceFromOrigin()
	difference := math.Abs(result - expected)
	if difference > epsilon {
		t.Errorf("Distance from p to origin is %g, not %g", p.DistanceFromOrigin(), expected)
	}
}

func TestGetQuadrantPoint2D_1(t *testing.T) { //test 1st quadrant
	p := Point2D{1,3}
	expected := 1
	result := p.GetQuadrant()
	if result != expected {
		t.Errorf("%s is in %d, not %d", p.ToString(), result, expected)
	}
}

func TestGetQuadrantPoint2D_2(t *testing.T) { //test 2nd quadrant
	p := Point2D{-1,3}
	expected := 2
	result := p.GetQuadrant()
	if result != expected {
		t.Errorf("%s is in %d, not %d", p.ToString(), result, expected)
	}
}

func TestGetQuadrantPoint2D_3(t *testing.T) { //test 3rd quadrant
	p := Point2D{-1,-3}
	expected := 3
	result := p.GetQuadrant()
	if result != expected {
		t.Errorf("%s is in %d, not %d", p.ToString(), result, expected)
	}
}

func TestGetQuadrantPoint2D_4(t *testing.T) { //test 4th quadrant
	p := Point2D{1,-3}
	expected := 4
	result := p.GetQuadrant()
	if result != expected {
		t.Errorf("%s is in %d, not %d", p.ToString(), result, expected)
	}
}

func TestGetQuadrantPoint2D_5(t *testing.T) { //test when on x axis
	p := Point2D{0,3}
	expected := -1
	result := p.GetQuadrant()
	if result != expected {
		t.Errorf("%s is in %d, not %d", p.ToString(), result, expected)
	}
}

func TestGetQuadrantPoint2D_6(t *testing.T) { //test when on y axis
	p := Point2D{1,0}
	expected := -1
	result := p.GetQuadrant()
	if result != expected {
		t.Errorf("%s is in %d, not %d", p.ToString(), result, expected)
	}
}

func TestGetQuadrantPoint2D_7(t *testing.T) { //test when origin
	p := Point2D{0,0}
	expected := -1
	result := p.GetQuadrant()
	if result != expected {
		t.Errorf("%s is in %d, not %d", p.ToString(), result, expected)
	}
}

func TestIsOnXaxis_1(t *testing.T) { //test when x is nonzero and y is 0
	p := Point2D{1,0}
	expected := false
	result := p.IsOnXaxis()
	if result != expected {
		t.Errorf("%s returns %t, not %t", p.ToString(), result, expected)
	}
}

func TestIsOnXaxis_2(t *testing.T) { //test when x and y are both nonzero
	p := Point2D{1,1}
	expected := false
	result := p.IsOnXaxis()
	if result != expected {
		t.Errorf("%s returns %t, not %t", p.ToString(), result, expected)
	}
}

func TestIsOnXaxis_3(t *testing.T) { //test when both x and y are 0
	p := Point2D{0,0}
	expected := true
	result := p.IsOnXaxis()
	if result != expected {
		t.Errorf("%s returns %t, not %t", p.ToString(), result, expected)
	}
}

func TestIsOnXaxis_4(t *testing.T) { //test when x is 0 and y is nonzero
	p := Point2D{0,5}
	expected := true
	result := p.IsOnXaxis()
	if result != expected {
		t.Errorf("%s returns %t, not %t", p.ToString(), result, expected)
	}
}

func TestIsOnXaxis_5(t *testing.T) { //test when x is less than 0
	p := Point2D{-1,0}
	expected := false
	result := p.IsOnXaxis()
	if result != expected {
		t.Errorf("%s returns %t, not %t", p.ToString(), result, expected)
	}
}

func TestIsOnXaxis_6(t *testing.T) { //test when both x and y are less than 0
	p := Point2D{-1,-1}
	expected := false
	result := p.IsOnXaxis()
	if result != expected {
		t.Errorf("%s returns %t, not %t", p.ToString(), result, expected)
	}
}

func TestIsOnYaxis_1(t *testing.T) { //test when x is 0 and y is greater than 0
	p := Point2D{0,1}
	expected := false
	result := p.IsOnYaxis()
	if result != expected {
		t.Errorf("%s return %t, not %t", p.ToString(), result, expected)
	}
}

func TestIsOnYaxis_2(t *testing.T) { //test when x is 0 ad y is less than 0
	p := Point2D{0,-1}
	expected := false
	result := p.IsOnYaxis()
	if result != expected {
		t.Errorf("%s return %t, not %t", p.ToString(), result, expected)
	}
}

func TestIsOnYaxis_3(t *testing.T) { //test when x is nonzero and y is 0
	p := Point2D{1,0}
	expected := true
	result := p.IsOnYaxis()
	if result != expected {
		t.Errorf("%s return %t, not %t", p.ToString(), result, expected)
	}
}

func TestIsOnYaxis_4(t *testing.T) { //test when origin
	p := Point2D{0,0}
	expected := true
	result := p.IsOnYaxis()
	if result != expected {
		t.Errorf("%s return %t, not %t", p.ToString(), result, expected)
	}
}

func TestIsOnYaxis_5(t *testing.T) { //test when x and y are nonzero and positive
	p := Point2D{1,1}
	expected := false
	result := p.IsOnYaxis()
	if result != expected {
		t.Errorf("%s return %t, not %t", p.ToString(), result, expected)
	}
}

func TestIsOnYaxis_6(t *testing.T) { //test when x and y are nonzero and negative
	p := Point2D{-1,-1}
	expected := false
	result := p.IsOnYaxis()
	if result != expected {
		t.Errorf("%s return %t, not %t", p.ToString(), result, expected)
	}
}

func TestIsOriginPoint2D_1(t *testing.T) { //test when in 1st quadrant
	p := Point2D{1,1}
	expected := false
	result := p.IsOrigin()
	if result != expected {
		t.Errorf("%s returns %t to IsOrigin, not %t",p.ToString(), result, expected)
	}
}

func TestIsOriginPoint2D_2(t *testing.T) { //test when in 2nd quadrant
	p := Point2D{-1,1}
	expected := false
	result := p.IsOrigin()
	if result != expected {
		t.Errorf("%s returns %t to IsOrigin, not %t",p.ToString(), result, expected)
	}
}

func TestIsOriginPoint2D_3(t *testing.T) { //test when in 3rd quadrant
	p := Point2D{-1,-1}
	expected := false
	result := p.IsOrigin()
	if result != expected {
		t.Errorf("%s returns %t to IsOrigin, not %t",p.ToString(), result, expected)
	}
}

func TestIsOriginPoint2D_4(t *testing.T) { //test when in 4th quadrant
	p := Point2D{1,-1}
	expected := false
	result := p.IsOrigin()
	if result != expected {
		t.Errorf("%s returns %t to IsOrigin, not %t",p.ToString(), result, expected)
	}
}

func TestIsOriginPoint2D_5(t *testing.T) { //test when on x axis only
	p := Point2D{0,1}
	expected := false
	result := p.IsOrigin()
	if result != expected {
		t.Errorf("%s returns %t to IsOrigin, not %t",p.ToString(), result, expected)
	}
}

func TestIsOriginPoint2D_6(t *testing.T) { //test when on y axis only
	p := Point2D{1,0}
	expected := false
	result := p.IsOrigin()
	if result != expected {
		t.Errorf("%s returns %t to IsOrigin, not %t",p.ToString(), result, expected)
	}
}

func TestIsOriginPoint2D_7(t *testing.T) { //test when is origin
	p := Point2D{0,0}
	expected := true
	result := p.IsOrigin()
	if result != expected {
		t.Errorf("%s returns %t to IsOrigin, not %t",p.ToString(), result, expected)
	}
}

func TestMidpointValuesPoint2D_1(t *testing.T) { //test when midpoint values are whole and positive
	p1 := Point2D{3,4}
	p2 := Point2D{7,2}
	expectedX, expectedY := 5., 3.

	resultX, resultY := p1.MidpointValues(&p2)

	diffX := math.Abs(expectedX - resultX)
	diffY := math.Abs(expectedY - resultY)

	if diffX > epsilon {
		t.Errorf("x value returned was %g, not %g", resultX, expectedX)
	}
	if diffY > epsilon {
		t.Errorf("y value returned was %g, not %g", resultY, expectedY)
	}
}

func TestMidpointValuesPoint2D_2(t *testing.T) { //test when midpoint is same point
	p1 := Point2D{3,4}
	p2 := Point2D{3,4}
	expectedX, expectedY := 3., 4.

	resultX, resultY := p1.MidpointValues(&p2)

	diffX := math.Abs(expectedX - resultX)
	diffY := math.Abs(expectedY - resultY)

	if diffX > epsilon {
		t.Errorf("x value returned was %g, not %g", resultX, expectedX)
	}
	if diffY > epsilon {
		t.Errorf("y value returned was %g, not %g", resultY, expectedY)
	}
}

func TestMidpointValuesPoint2D_3(t *testing.T) { //test when doing midpoints where 1 value is negative and the other is positive
	p1 := Point2D{-4,4}
	p2 := Point2D{7,-1}
	expectedX, expectedY := 1.5, 1.5

	resultX, resultY := p1.MidpointValues(&p2)

	diffX := math.Abs(expectedX - resultX)
	diffY := math.Abs(expectedY - resultY)

	if diffX > epsilon {
		t.Errorf("x value returned was %g, not %g", resultX, expectedX)
	}
	if diffY > epsilon {
		t.Errorf("y value returned was %g, not %g", resultY, expectedY)
	}
}

func TestMidpointValuesPoint2D_4(t *testing.T) { //test when midpoint values are whole and negative
	p1 := Point2D{-3,-4}
	p2 := Point2D{-7,-2}
	expectedX, expectedY := -5., -3.

	resultX, resultY := p1.MidpointValues(&p2)

	diffX := math.Abs(expectedX - resultX)
	diffY := math.Abs(expectedY - resultY)

	if diffX > epsilon {
		t.Errorf("x value returned was %g, not %g", resultX, expectedX)
	}
	if diffY > epsilon {
		t.Errorf("y value returned was %g, not %g", resultY, expectedY)
	}
}

func TestMidpointValuesPoint2D_5(t *testing.T) { //test midpoint values negative and a mix of whole and floating point
	p1 := Point2D{-1,-4}
	p2 := Point2D{-6,-2}
	expectedX, expectedY := -3.5, -3.

	resultX, resultY := p1.MidpointValues(&p2)

	diffX := math.Abs(expectedX - resultX)
	diffY := math.Abs(expectedY - resultY)

	if diffX > epsilon {
		t.Errorf("x value returned was %g, not %g", resultX, expectedX)
	}
	if diffY > epsilon {
		t.Errorf("y value returned was %g, not %g", resultY, expectedY)
	}
}

//the MidpointApprox tests use the same points as their correspoidning MidpointValues test

func TestMidpointApproxPoint2D_1(t *testing.T) {
	p1 := Point2D{3,4}
	p2 := Point2D{7,2}
	expected := Point2D{5,3}
	result := p1.MidpointApprox(&p2)

	if p1.MidpointApprox(&p2) != p2.MidpointApprox(&p1) {
		t.Errorf("MidpointApprox not commutative")
	}

	if result != expected {
		t.Errorf("Midpoint of %s and %s returned was %s, not %s", 
		p1.ToString(), p2.ToString(), result.ToString(), expected.ToString())
	}
}

func TestMidpointApproxPoint2D_2(t *testing.T) {
	p1 := Point2D{3,4}
	p2 := Point2D{3,4}
	expected := Point2D{3,4}
	result := p1.MidpointApprox(&p2)

	if p1.MidpointApprox(&p2) != p2.MidpointApprox(&p1) {
		t.Errorf("MidpointApprox not commutative")
	}

	if result != expected {
		t.Errorf("Midpoint of %s and %s returned was %s, not %s", 
		p1.ToString(), p2.ToString(), result.ToString(), expected.ToString())
	}
}

func TestMidpointApproxPoint2D_3(t *testing.T) {
	p1 := Point2D{-4,4}
	p2 := Point2D{7,-1}
	expected := Point2D{2,2}
	result := p1.MidpointApprox(&p2)

	if p1.MidpointApprox(&p2) != p2.MidpointApprox(&p1) {
		t.Errorf("MidpointApprox not commutative")
	}

	if result != expected {
		t.Errorf("Midpoint of %s and %s returned was %s, not %s", 
		p1.ToString(), p2.ToString(), result.ToString(), expected.ToString())
	}
}

func TestMidpointApproxPoint2D_4(t *testing.T) {
	p1 := Point2D{-3,-4}
	p2 := Point2D{-7,-2}
	expected := Point2D{-5,-3}
	result := p1.MidpointApprox(&p2)

	if p1.MidpointApprox(&p2) != p2.MidpointApprox(&p1) {
		t.Errorf("MidpointApprox not commutative")
	}

	if result != expected {
		t.Errorf("Midpoint of %s and %s returned was %s, not %s", 
		p1.ToString(), p2.ToString(), result.ToString(), expected.ToString())
	}
}

func TestMidpointApproxPoint2D_5(t *testing.T) {
	p1 := Point2D{-1,-4}
	p2 := Point2D{-6,-2}
	expected := Point2D{-4,-3}
	result := p1.MidpointApprox(&p2)

	if p1.MidpointApprox(&p2) != p2.MidpointApprox(&p1) {
		t.Errorf("MidpointApprox not commutative")
	}

	if result != expected {
		t.Errorf("Midpoint of %s and %s returned was %s, not %s", 
		p1.ToString(), p2.ToString(), result.ToString(), expected.ToString())
	}
}

func TestMidpointValuesPoint2D_6(t *testing.T) { //test when a midpoint value is 0
	p1 := Point2D{-3,-4}
	p2 := Point2D{3,4}
	expectedX, expectedY := 0., 0.

	resultX, resultY := p1.MidpointValues(&p2)

	diffX := math.Abs(expectedX - resultX)
	diffY := math.Abs(expectedY - resultY)

	if diffX > epsilon {
		t.Errorf("x value returned was %g, not %g", resultX, expectedX)
	}
	if diffY > epsilon {
		t.Errorf("y value returned was %g, not %g", resultY, expectedY)
	}
}

func TestSlopePoint2DEquality(t *testing.T) {
	p1 := Point2D{6,7}
	p2 := Point2D{-4,5}
	if p1.Slope(&p2) != p2.Slope(&p1) {
		t.Errorf("Slope is not commutative")
	}
}
func TestSlopePoint2D_1(t *testing.T) {
	p1 := Point2D{5,4}
	p2 := Point2D{7,6}
	expected := 1.
	result := p1.Slope(&p2)
	diff := math.Abs(expected - result)

	if diff > epsilon {
		t.Errorf("Slope between %s and %s was %g, not %g", p1.ToString(), p2.ToString(), result, expected)
	}
}

func TestSlopePoint2D_2(t *testing.T) {
	p1 := Point2D{1,1}
	p2 := Point2D{2,3}
	expected := 2.
	result := p1.Slope(&p2)
	diff := math.Abs(expected - result)

	if diff > epsilon {
		t.Errorf("Slope between %s and %s was %g, not %g", p1.ToString(), p2.ToString(), result, expected)
	}
}

func TestSlopePoint2D_3(t *testing.T) {
	p1 := Point2D{10,1}
	p2 := Point2D{32,2}
	expected := .045454545
	result := p1.Slope(&p2)
	diff := math.Abs(expected - result)

	if diff > epsilon {
		t.Errorf("Slope between %s and %s was %g, not %g", p1.ToString(), p2.ToString(), result, expected)
	}
}

func TestSlopePoint2D_4(t *testing.T) {
	p1 := Point2D{-5,8}
	p2 := Point2D{7,6}
	expected := -0.16666666666
	result := p1.Slope(&p2)
	diff := math.Abs(expected - result)

	if diff > epsilon {
		t.Errorf("Slope between %s and %s was %g, not %g", p1.ToString(), p2.ToString(), result, expected)
	}
}

func TestSlopePoint2D_5(t *testing.T) {
	p1 := Point2D{5,4}
	p2 := Point2D{7,4}
	expected := 0.
	result := p1.Slope(&p2)
	diff := math.Abs(expected - result)

	if diff > epsilon {
		t.Errorf("Slope between %s and %s was %g, not %g", p1.ToString(), p2.ToString(), result, expected)
	}
}

func TestSlopePoint2D_6(t *testing.T) {
	p1 := Point2D{5,4}
	p2 := Point2D{5,6}
	expected := math.Inf(0)
	result := p1.Slope(&p2)
	diff := math.Abs(expected - result)

	if diff > epsilon {
		t.Errorf("Slope between %s and %s was %g, not %g", p1.ToString(), p2.ToString(), result, expected)
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

//note: for all tests below, XMin, XMax, YMin, YMax = -5, 10, -4, 8

func TestIsValidPoint2D_1(t *testing.T) { //test both points below minimum
	SetDimensions2D(-5, 10, -4, 8)
	p := Point2D{ -6,-5}
	expected := false
	result := p.IsValid()

	if result != expected {
		t.Errorf("Validity of %s returned %t, not %t with bounds of XMin,XMax,Ymin,YMax = %d,%d,%d,%d",
		p.ToString(), result, expected, XMin, XMax, YMin, YMax)
	}
}

func TestIsValidPoint2D_2(t *testing.T) { //test x point below minimum
	SetDimensions2D(-5, 10, -4, 8)
	p := Point2D{ -6,-4}
	expected := false
	result := p.IsValid()

	if result != expected {
		t.Errorf("Validity of %s returned %t, not %t with bounds of XMin,XMax,Ymin,YMax = %d,%d,%d,%d",
		p.ToString(), result, expected, XMin, XMax, YMin, YMax)
	}
}

func TestIsValidPoint2D_3(t *testing.T) { //test y point below minimum
	SetDimensions2D(-5, 10, -4, 8)
	p := Point2D{ -5, -5}
	expected := false
	result := p.IsValid()

	if result != expected {
		t.Errorf("Validity of %s returned %t, not %t with bounds of XMin,XMax,Ymin,YMax = %d,%d,%d,%d",
		p.ToString(), result, expected, XMin, XMax, YMin, YMax)
	}
}

func TestIsValidPoint2D_4(t *testing.T) { //test both points above maximum
	SetDimensions2D(-5, 10, -4, 8)
	p := Point2D{ 11,9}
	expected := false
	result := p.IsValid()

	if result != expected {
		t.Errorf("Validity of %s returned %t, not %t with bounds of XMin,XMax,Ymin,YMax = %d,%d,%d,%d",
		p.ToString(), result, expected, XMin, XMax, YMin, YMax)
	}
}

func TestIsValidPoint2D_5(t *testing.T) { //test x point above maximum
	SetDimensions2D(-5, 10, -4, 8)
	p := Point2D{ 11, 8}
	expected := false
	result := p.IsValid()

	if result != expected {
		t.Errorf("Validity of %s returned %t, not %t with bounds of XMin,XMax,Ymin,YMax = %d,%d,%d,%d",
		p.ToString(), result, expected, XMin, XMax, YMin, YMax)
	}
}

func TestIsValidPoint2D_6(t *testing.T) { //test y point above maximum
	SetDimensions2D(-5, 10, -4, 8)
	p := Point2D{ 10,9}
	expected := false
	result := p.IsValid()

	if result != expected {
		t.Errorf("Validity of %s returned %t, not %t with bounds of XMin,XMax,Ymin,YMax = %d,%d,%d,%d",
		p.ToString(), result, expected, XMin, XMax, YMin, YMax)
	}
}

func TestIsValidPoint2D_7(t *testing.T) { //test x below and y above
	SetDimensions2D(-5, 10, -4, 8)
	p := Point2D{ -6,9}
	expected := false
	result := p.IsValid()

	if result != expected {
		t.Errorf("Validity of %s returned %t, not %t with bounds of XMin,XMax,Ymin,YMax = %d,%d,%d,%d",
		p.ToString(), result, expected, XMin, XMax, YMin, YMax)
	}
}

func TestIsValidPoint2D_8(t *testing.T) { //test x above and y below
	SetDimensions2D(-5, 10, -4, 8)
	p := Point2D{ 11,-5}
	expected := false
	result := p.IsValid()

	if result != expected {
		t.Errorf("Validity of %s returned %t, not %t with bounds of XMin,XMax,Ymin,YMax = %d,%d,%d,%d",
		p.ToString(), result, expected, XMin, XMax, YMin, YMax)
	}
}

func TestIsValidPoint2D_9(t *testing.T) { //test x is min and y is max
	SetDimensions2D(-5, 10, -4, 8)
	p := Point2D{ -5,8}
	expected := true
	result := p.IsValid()

	if result != expected {
		t.Errorf("Validity of %s returned %t, not %t with bounds of XMin,XMax,Ymin,YMax = %d,%d,%d,%d",
		p.ToString(), result, expected, XMin, XMax, YMin, YMax)
	}
}

func TestIsValidPoint2D_10(t *testing.T) { //test x is max and y is min
	SetDimensions2D(-5, 10, -4, 8)
	p := Point2D{ 10,-4}
	expected := true
	result := p.IsValid()

	if result != expected {
		t.Errorf("Validity of %s returned %t, not %t with bounds of XMin,XMax,Ymin,YMax = %d,%d,%d,%d",
		p.ToString(), result, expected, XMin, XMax, YMin, YMax)
	}
}

func TestIsValidPoint2D_11(t *testing.T) { //test both are at min
	SetDimensions2D(-5, 10, -4, 8)
	p := Point2D{ -5,-4}
	expected := true
	result := p.IsValid()

	if result != expected {
		t.Errorf("Validity of %s returned %t, not %t with bounds of XMin,XMax,Ymin,YMax = %d,%d,%d,%d",
		p.ToString(), result, expected, XMin, XMax, YMin, YMax)
	}
}

func TestIsValidPoint2D_12(t *testing.T) { //test both are at max
	SetDimensions2D(-5, 10, -4, 8)
	p := Point2D{ 10,8}
	expected := true
	result := p.IsValid()

	if result != expected {
		t.Errorf("Validity of %s returned %t, not %t with bounds of XMin,XMax,Ymin,YMax = %d,%d,%d,%d",
		p.ToString(), result, expected, XMin, XMax, YMin, YMax)
	}
}

func TestIsValidPoint2D_13(t *testing.T) { //test both in middle of range
	SetDimensions2D(-5, 10, -4, 8)
	p := Point2D{ -3,4}
	expected := true
	result := p.IsValid()

	if result != expected {
		t.Errorf("Validity of %s returned %t, not %t with bounds of XMin,XMax,Ymin,YMax = %d,%d,%d,%d",
		p.ToString(), result, expected, XMin, XMax, YMin, YMax)
	}
}
