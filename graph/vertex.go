package graph

type Vertex[T any] struct {
	key      int
	data     T
	adjacent []*Vertex[T]
}
