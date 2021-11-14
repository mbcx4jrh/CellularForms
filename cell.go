package main

type cell struct {
	index    int
	position Vector3
	normal   Vector3
	links    []cell
}

func (c *cell) computeNormal() {
	sum := Zero()
	for i, n := range c.links {
		a := NewSubtract(&n.position, &c.position)
		b := NewSubtract(&c.links[(i+1)%len(c.links)].position, &c.position)
		p := Cross(&a, &b)
		sum.Add(&p)
	}
	sum.Normalise()

	if Dot(&c.normal, &sum) < 0 {
		sum.Negate()
	}

	c.normal = sum
}
