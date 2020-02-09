package point

import (
	"fmt"
	"math"
)

//Point3D is used to represent points in 3D space
type Point3D struct {
	Point2D
	Z int
}

//ToString shows how a Point3D shold be printed, returns string in (x,y,z) format
func (caller *Point3D) ToString() string {
	return fmt.Sprintf("(%d,%d,%d)", caller.X, caller.Y, caller.Z)
}

//DistanceFrom calculates distance between two Point3D structs
func (caller *Point3D) DistanceFrom(other *Point3D) float64 {
	return math.Sqrt(float64((caller.X-other.X)*(caller.X-other.X) + (caller.Y-other.Y)*
		(caller.Y-other.Y) + (caller.Z-other.Z)*(caller.Z-other.Z)))
}

//DistranceFromOrigin calculates special case of a Point3D's distance from (0,0,0)
func (caller *Point3D) DistranceFromOrigin() float64 {
	return math.Sqrt(float64(caller.X*caller.X + caller.Y*caller.Y + caller.Z*caller.Z))
}

//TODO : create 3D version of GetQuadrant

//IsOnZaxis returns if the point lies on the z-axis
func (caller *Point3D) IsOnZaxis() bool {
	return caller.Z == 0
}

//IsOrigin returns if the point is the origin
func (caller *Point3D) IsOrigin() bool {
	return caller.Point2D.IsOrigin() && caller.Z == 0
}

//MidpointValues is a method to get the actual midpoint values between two Point3Ds
func (caller *Point3D) MidpointValues(other *Point3D) (x, y, z float64) {
	x, y = caller.Point2D.MidpointValues(&other.Point2D)
	z = float64(caller.Z+other.Z) / 2.
	return x, y, z
}

//MidpointApprox is a method to get the approximate midpoint between two Point3D point
func (caller *Point3D) MidpointApprox(other *Point3D) (midpoint Point3D) {
	x, y, z := caller.MidpointValues(other)                     //get midpoint values
	x, y, z = math.Round(x), math.Round(y), math.Round(z)       //round prior to casting
	midpoint.X, midpoint.Y, midpoint.Z = int(x), int(y), int(z) //cast float values to integers and assign to the midpoint
	return midpoint
}

//TODO add a slope method that returns the vector

/*As in Point2D, the methods below are intended for use incases where Point3D is
being used to keep track of locations in a bounded coordinate grid such as a 3D array*/

//ZMin is the minimum z value allowed, inclusive
var ZMin int

//ZMax is the max z value allowed, inclusive
var ZMax int

//SetDimensions3D sets the the dimensions for the x/y/z planes
func SetDimensions3D(xmin, xmax, ymin, ymax, zmin, zmax int) {
	ZMin, ZMax = zmin, zmax
	SetDimensions2D(xmin, xmax, ymin, ymax)
}

//IsValid returns whether the Point3D is within the dimension bounds
func (caller *Point3D) IsValid() bool {
	return caller.Point2D.IsValid() && caller.Z >= ZMin && caller.Z <= ZMax
}

//HasIn return whether or not a Point3D has a Point3D in from it in the z direction
//in is defined as decreasing z values
func (caller *Point3D) HasIn() bool {
	return caller.Z > ZMin
}

//HasOut returns whether or not a Point3D has a Point3D out from it in the z direction
//out is defined as increasing z values
func (caller *Point3D) HasOut() bool {
	return caller.Z < ZMax
}

//the following return a Point3D in the specified direction

//Right returns the Point3D to the right of the caller
//right is defined as increasing x values
func (caller *Point3D) Right() (right Point3D) {
	right.Point2D = caller.Point2D.Right()
	right.Z = caller.Z
	return right
}

//Left returns the Point3D to the left of the caller
//left is defined as decreasing x values
func (caller *Point3D) Left() (left Point3D) {
	left.Point2D = caller.Point2D.Left()
	left.Z = caller.Z
	return left
}

//Up returns a Point3D up from the calller
//up is defined as decreasing y values
func (caller *Point3D) Up() (up Point3D) {
	up.Point2D = caller.Point2D.Up()
	up.Z = caller.Z
	return up
}

//Down returns a Point3D down from the caller
//down is defined as increasing y values
func (caller *Point3D) Down() (down Point3D) {
	down.Point2D = caller.Point2D.Down()
	down.Z = caller.Z
	return down
}

//In returns a Point3D in from the caller
//in is defined as decreasing z values
func (caller *Point3D) In() (in Point3D) {
	in.Point2D = caller.Point2D
	in.Z = caller.Z - 1
	return in
}

//Out returns a Point3D out from the caller
//out is defined as increasing z values
func (caller *Point3D) Out() (out Point3D) {
	out.Point2D = caller.Point2D
	out.Z = caller.Z + 1
	return out
}
