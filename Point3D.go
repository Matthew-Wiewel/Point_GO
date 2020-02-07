package point

//Point3D is used to represent points in 3D space
type Point3D struct {
    Point2D
    Z int
}

//ZMin is the minimum z value allowed
var ZMin int
//ZMax is the max z value allowed
var ZMax int

