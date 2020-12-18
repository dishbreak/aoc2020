package lib

// Point3D represents an arbitrary point in 3D space
type Point3D struct {
	X, Y, Z int
}

// Add will perform vector addition and return a new Point3D
func (p Point3D) Add(other Point3D) Point3D {
	return Point3D{
		X: p.X + other.X,
		Y: p.Y + other.Y,
		Z: p.Z + other.Z,
	}
}

// Neighbors returns the 26 coordinates surrounding the given point.
func (p Point3D) Neighbors() []Point3D {
	result := make([]Point3D, 26)
	c := 0
	for i := -1; i <= 1; i++ {
		for j := -1; j <= 1; j++ {
			for k := -1; k <= 1; k++ {
				if i == 0 && j == 0 && k == 0 {
					continue
				}
				result[c] = p.Add(Point3D{i, j, k})
				c++
			}
		}
	}
	return result
}
