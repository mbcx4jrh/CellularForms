package main

type mesh struct {
	triangles [][3]int
	vertices  []Vector3
}

func isocahedron() mesh {

	a := 0.8506507174597755
	b := 0.5257312591858783

	vertices := []Vector3{
		Vector3{-a, -b, 0}, Vector3{-a, b, 0}, Vector3{-b, 0, -a}, Vector3{-b, 0, a},
		Vector3{0, -a, -b}, Vector3{0, -a, b}, Vector3{0, a, -b}, Vector3{0, a, b},
		Vector3{b, 0, -a}, Vector3{b, 0, a}, Vector3{a, -b, 0}, Vector3{a, b, 0},
	}

	triangles := [][3]int{
		{0, 3, 1}, {1, 3, 7}, {2, 0, 1}, {2, 1, 6},
		{4, 0, 2}, {4, 5, 0}, {5, 3, 0}, {6, 1, 7},
		{6, 7, 11}, {7, 3, 9}, {8, 2, 6}, {8, 4, 2},
		{8, 6, 11}, {8, 10, 4}, {8, 11, 10}, {9, 3, 5},
		{10, 5, 4}, {10, 9, 5}, {11, 7, 9}, {11, 9, 10},
	}

	return mesh{triangles, vertices}
}
