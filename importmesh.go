package main

import "fmt"

func importMesh(m mesh) []cell {

	//create cells
	cells := make([]cell, 0, len(m.vertices))
	for _, v := range m.vertices {
		index := len(cells)
		cells = append(cells, cell{index, v, v, v, []cell{}})
	}

	//initialise mapping for triangles on a vertex
	vertexTriangles := make([][]triangle, len(m.vertices))
	for i := range vertexTriangles {
		vertexTriangles[i] = make([]triangle, 0)
	}

	//map triangles to vertex
	for _, i := range m.triangles {
		t := triangle{i[0], i[1], i[2]}
		//fmt.Printf("adding triangle %d %d %d\n", t.a, t.b, t.c)
		vertexTriangles[t.a] = append(vertexTriangles[t.a], t)
		vertexTriangles[t.b] = append(vertexTriangles[t.b], t)
		vertexTriangles[t.c] = append(vertexTriangles[t.c], t)
	}

	//sort triangles at each vertex to ccw
	for v, verTris := range vertexTriangles {
		for i := 1; i < len(verTris); i++ {
			prev := verTris[i-1].vertexBefore(v)
			for j := 1; j < len(verTris); j++ {
				if verTris[j].vertexAfter(v) == prev {
					temp := verTris[i]
					verTris[i] = verTris[j]
					verTris[j] = temp
					break
				}
			}
		}

		//create links for each cell
		for _, t := range verTris {
			cell := cells[t.vertexAfter(v)]
			//if !contains(cells[v].links, cell) {
			cells[v].links = append(cells[v].links, cell)
			//}
		}
	}

	return cells
}

//efficient only for small slices
//may not be needed but retained in case
// func contains(cells []cell, c cell) bool {
// 	for _, i := range cells {
// 		if i.index == c.index {
// 			return true
// 		}
// 	}
// 	return false
// }

type triangle struct {
	a, b, c int
}

func (t triangle) vertexBefore(p int) int {
	if p == t.a {
		return t.c
	}
	if p == t.b {
		return t.a
	}
	if p == t.c {
		return t.b
	}
	fmt.Println("Vertex not found in vertexBefore!")
	return t.a
}

func (t triangle) vertexAfter(p int) int {
	if p == t.a {
		return t.b
	}
	if p == t.b {
		return t.c
	}
	if p == t.c {
		return t.a
	}
	fmt.Println("Vertex not found in vertexAfter!")
	return t.a
}
