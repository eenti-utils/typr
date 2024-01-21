package typr

// an envelope holding two values, each respectively of specified type
type Pair[X, Y any] struct {
	X X
	Y Y
}

// returns a copy of the the Pair
func (p Pair[X, Y]) Value() (r Pair[X, Y]) {
	r = Pair[X, Y]{X: p.X, Y: p.Y}
	return
}

// an envelope holding three values, each respectively of specified type
type Trio[X, Y, Z any] struct {
	X X
	Y Y
	Z Z
}

// returns a copy of the Trio
func (t Trio[X, Y, Z]) Value() (r Trio[X, Y, Z]) {
	r = Trio[X, Y, Z]{X: t.X, Y: t.Y, Z: t.Z}
	return
}
