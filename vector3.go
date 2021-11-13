package main

import "math"

type Vector3 struct {
	x, y, z float64
}

func Zero() Vector3 {
	return Vector3{0.0, 0.0, 0.0}
}

func (v1 *Vector3) Add(v2 *Vector3) {
	v1.x += v2.x
	v1.y += v2.y
	v1.z += v2.z
}

func (v1 *Vector3) Subtract(v2 *Vector3) {
	v1.x -= v2.x
	v1.y -= v2.y
	v1.z -= v2.z
}

func (v1 *Vector3) Multiply(n float64) {
	v1.x *= n
	v1.y *= n
	v1.z *= n
}

func SqrMagnitude(v1 *Vector3) float64 {
	return v1.x*v1.x + v1.y*v1.y + v1.z*v1.z
}

func Magnitude(v1 *Vector3) float64 {
	return math.Sqrt(SqrMagnitude(v1))
}

func Dot(v1 *Vector3, v2 *Vector3) float64 {
	return v1.x*v2.x + v1.y*v2.y + v1.z*v2.z
}

func (v1 *Vector3) Normalise() {
	m := Magnitude(v1)
	v1.x = v1.x / m
	v1.y = v1.y / m
	v1.z = v1.z / m
}
