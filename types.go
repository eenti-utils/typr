package typr

// an envelope holding two values, each respectively of specified type
type Pair[X,Y any] struct {
	X X
	Y Y
}

// an envelope holding three values, each respectively of specified type
type Trio[X,Y,Z any] struct {
	X X
	Y Y
	Z Z
}