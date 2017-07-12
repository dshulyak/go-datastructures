package ptrie

import "fmt"

const (
	dimension = 2
)

func makeRoot() root {
	r := root{
		nei: makeNode(),
	}
	r.nei.depth = 2
	return r
}

type root struct {
	count int
	nei   *node
}

type node struct {
	neis  [dimension]*node
	val   string
	depth int
}

func makeNode() *node {
	return &node{
		neis: [dimension]*node{},
	}
}

func (r root) insert(val string) root {
	r.count++
	// 3 cases
	// insert in place
	// extend existing structure
	// root overflow
	r.nei = r.nei.insert(r.count, val)
	return r
}

func (r *root) lookup(count int) string {
	return r.nei.lookup(count + 1)
}

func (r root) delete(count int) root {
	count++
	r.nei = r.nei.delete(count)
	r.count--
	nonNil := 0
	var nonNilNode *node
	for _, nei := range r.nei.neis {
		if nei != nil {
			nonNil++
			nonNilNode = nei
		}
	}
	if nonNil == 1 {
		r.nei = nonNilNode
	}
	return r
}

func (r root) listVals() []string {
	vals := []string{}
	vals = append(vals, r.nei.listVals()...)
	return vals
}

func Pow(a, b int) int {
	p := 1
	for b > 0 {
		if b&1 != 0 {
			p *= a
		}
		b >>= 1
		a *= a
	}
	return p
}

func (r *node) insert(count int, val string) *node {
	fmt.Printf("setting val %s with key %d \n", val, count)
	new := *r
	if r.depth == 1 {
		fmt.Println("setting a val")
		new.val = val
		return &new
	}
	capacity := Pow(dimension, new.depth-1)
	if count > capacity {
		fmt.Println("root overflow, adding one more level")
		new.neis[0] = r
		new.depth++
		for i := 1; i < dimension; i++ {
			new.neis[i] = nil
		}
		capacity = Pow(dimension, new.depth-1)
	}
	fmt.Println("new capacity", capacity)
	if count <= capacity {
		pos := (count / new.depth) % dimension
		if new.neis[pos] == nil {
			fmt.Println("creating nei")
			new.neis[pos] = makeNode()
			new.neis[pos].depth = new.depth - 1
		}
		fmt.Println("adding a child with depth", new.neis[pos].depth)
		if count > capacity/2 {
			count -= capacity / 2
		}
		new.neis[pos] = new.neis[pos].insert(count, val)
	}
	return &new
}

func (r node) delete(count int) *node {
	if r.depth == 1 {
		return nil
	}
	pos := (count / r.depth) % dimension
	capacity := Pow(dimension, r.depth-1)
	if count > capacity/2 {
		count -= capacity / 2
	}
	r.neis[pos] = r.neis[pos].delete(count)
	for _, nei := range r.neis {
		if nei != nil {
			return &r
		}
	}
	return nil
}

func (r *node) lookup(count int) string {
	if r.depth == 1 {
		return r.val
	}
	pos := (count / r.depth) % dimension
	capacity := Pow(dimension, r.depth-1)
	if count > capacity/2 {
		count -= capacity / 2
	}
	return r.neis[pos].lookup(count)
}

func (r *node) listVals() []string {
	vals := []string{}
	for _, nei := range r.neis {
		if nei == nil {
			continue
		}
		vals = append(vals, nei.listVals()...)
	}
	if r.val != "" {
		vals = append(vals, r.val)
	}
	return vals
}
