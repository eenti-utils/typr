package typr


// returns bool true, if the specified element e "qualifies", and bool false, otherwise
type Qualifier[V any] func(e V) (r bool)

// accepts an index and an element of the specified type, and return an Operation value
type Ranger[V any] func(i int,e V) (r Op)

// returns the resolved value (as the specified type), for the specified value
type Resolver[V,R any] func(v V) (r R)

// returns a non-nil error, if the specified element e is considered invalid, and nil, otherwise
type Validator[V any] func(e V) (r error)