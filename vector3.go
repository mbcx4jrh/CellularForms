package main

type Vector3 struct {
	x, y, z float64
}

func (v1 *Vector3) Add(v2 *Vector3) {
	v1.x += v2.x
	v1.y += v2.y
	v1.z += v2.z
}
