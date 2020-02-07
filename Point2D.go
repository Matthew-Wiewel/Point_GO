package point

import "math"

//Point2D is a structure for 2D points
type Point2D struct {
	X, Y int
}

//DistanceFrom is a method to get the absolute distance between two points
func (p* Point2D) DistanceFrom(other *Point2D) float64 {
	return math.Sqrt((float64)((p.X - p.X) * (p.X - other.X) + (p.Y - other.Y) * (p.Y - other.Y)))
}