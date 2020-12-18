package lib

// Point4D represents an arbitrary point in 3D space
type Point4D struct {
	X, Y, Z, W int
}

// Add will perform vector addition and return a new Point4D
func (p Point4D) Add(other Point4D) Point4D {
	return Point4D{
		X: p.X + other.X,
		Y: p.Y + other.Y,
		Z: p.Z + other.Z,
		W: p.W + other.W,
	}
}

// Sub will perform vector subtraction and return a new Point4D with other being
// the right-hand operand.
func (p Point4D) Sub(other Point4D) Point4D {
	return Point4D{
		X: p.X - other.X,
		Y: p.Y - other.Y,
		Z: p.Z - other.Z,
		W: p.W - other.W,
	}
}

// Neighbors returns the 80 coordinates surrounding the given point.
func (p Point4D) Neighbors() []Point4D {
	result := make([]Point4D, 80)
	c := 0
	for i := -1; i <= 1; i++ {
		for j := -1; j <= 1; j++ {
			for k := -1; k <= 1; k++ {
				for l := -1; l <= 1; l++ {
					if i == 0 && j == 0 && k == 0 && l == 0 {
						continue
					}
					result[c] = p.Add(Point4D{i, j, k, l})
					c++
				}
			}
		}
	}
	return result
}
