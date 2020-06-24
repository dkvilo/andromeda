package vector

// Vector3d Struct
type Vector3d struct {
	X, Y, Z float64
}

func (vec *Vector3d) Xi() int32 {
	return int32(vec.X)
}

func (vec *Vector3d) Yi() int32 {
	return int32(vec.Y)
}

func (vec *Vector3d) Zi() int32 {
	return int32(vec.Z)
}
