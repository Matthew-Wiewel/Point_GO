package point

import "math"

//Point2D is a structure for 2D points
type Point2D struct {
	X, Y int
}

//DistanceFrom is a method to get the absolute distance between two points
func (caller *Point2D) DistanceFrom(other *Point2D) float64 {
	return math.Sqrt(float64((caller.X-other.X)*(caller.X-other.X) + (caller.Y-other.Y)*(caller.Y-other.Y)))
}

//DistanceFromOrigin is a specialized version of DistanceFrom to avoid needing an origin Point2D
func (caller *Point2D) DistanceFromOrigin() float64 {
	return math.Sqrt(float64(caller.X*caller.X + caller.Y*caller.Y))
}

//GetQuadrant returns which quadrant on the x/y plane a point lies in
func (caller *Point2D) GetQuadrant() int {
	if caller.X > 0 && caller.Y > 0 { //x and y positive, 1st quadrant
		return 1
	} else if caller.X < 0 && caller.Y > 0 { //x negative and y positive, 2nd qudrant
		return 2
	} else if caller.X < 0 && caller.Y < 0 { //x and y negative, 3rd quadrant
		return 3
	} else if caller.X > 0 && caller.Y < 0 { //a positive and y negative, 4th quadrant
		return 4
	} else { //error, one of the points is on an axis, no quadrant
		return -1
	}
}

//IsOnXaxis returns if the point is on the X axis
func (caller *Point2D) IsOnXaxis() bool {
	return caller.X == 0
}

//IsOnYaxis returns if the point is on the Y axis
func (caller *Point2D) IsOnYaxis() bool {
	return caller.Y == 0
}

//IsOrigin returns if the point is the origin
func (caller *Point2D) IsOrigin() bool {
	return caller.X == 0 && caller.Y == 0
}

//MidpointApprox returns the approximate midpoint between two Point2D objects
//it does so by getting the floating point value and then rounding
func (caller *Point2D) MidpointApprox(other *Point2D) (midpoint Point2D) {
	x, y := caller.MidpointValues(other) //get floating point values first

	//round before casting to ints
	x, y = math.Round(x), math.Round(y)

	//cast x and y as integers and put them into the midpoint
	midpoint.X = int(x)
	midpoint.Y = int(y)
	return midpoint
}

//MidpointValues returns the x and y float64 values of the midpoint between two points
func (caller *Point2D) MidpointValues(other *Point2D) (x, y float64) {
	x = float64(caller.X+other.X) / 2.
	y = float64(caller.Y+other.Y) / 2.
	return x, y
}

/*the following methods are meant for use with a 2D grid such as a 2D array
these methods expect (0,0) to be in the top left of the grid with increasing
x values being to the right, decreasing x values being to the left
increasing y values being down and decreasing y values being up
XMin, XMax, YMin, and YMax should be set prior to calling these methods
*/

//XMin and the rest of these variables are intended for use when using Point2D to
//track locations in a 2D-grid. These values are inclusive
var XMin int //XMin
//XMax lint-satisfy
var XMax int

//YMin lint-satisfy
var YMin int

//YMax lint-satisfy
var YMax int

//SetDimensions2D is a method used to set the minimum and maximum dimesnions along the X/Y plane for Point2D
func SetDimensions2D(xmin, xmax, ymin, ymax int) {
	XMin, XMax, YMin, YMax = xmin, xmax, ymin, ymax
}

//IsValid returns whether or not the calling Point2D is within the bounds of the grid
func (caller *Point2D) isValid() bool {
	return caller.X >= XMin && caller.X <= XMax && caller.Y >= YMin && caller.Y <= YMax
}

//HasRight returns if there is a Point2D to the right of the caller within bounds
func (caller *Point2D) HasRight() bool {
	return caller.X < XMax
}

//HasLeft returns if there a Point2D to the left of the caller within bounds
func (caller *Point2D) HasLeft() bool {
	return caller.X > XMin
}

//HasUp returns if there is a Point2D above the caller within bounds
func (caller *Point2D) HasUp() bool {
	return caller.Y > YMin
}

//HasDown returns if there is a Point2D below the caller within bounds
func (caller *Point2D) HasDown() bool {
	return caller.Y < YMax
}

//these methods return a Point2D in the corresponding direction

//Right return a Point2D to the right
func (caller *Point2D) Right() Point2D {
	return Point2D{caller.X + 1, caller.Y}
}

//Left returns a Point2D to the left
func (caller *Point2D) Left() Point2D {
	return Point2D{caller.X - 1, caller.Y}
}

//Up returns a Point2D above
func (caller *Point2D) Up() Point2D {
	return Point2D{caller.X, caller.Y - 1}
}

//Down returns a Point2D below
func (caller *Point2D) Down() Point2D {
	return Point2D{caller.X, caller.Y + 1}
}
