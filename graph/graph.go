package graph

import (
	"errors"
)

func contains[T any](vertices []*Vertex[T], key int) bool {
	for _, v := range vertices {
		if key == v.key {
			return true
		}
	}
	return false
}

type Graph[T any] struct {
	vertices []*Vertex[T]
}

func (graph *Graph[T]) getVertex(key int) *Vertex[T] {
	for i, v := range graph.vertices {
		if v.key == key {
			return graph.vertices[i]
		}
	}
	return nil
}

func (graph *Graph[T]) AddVertex(key int) error {
	if contains(graph.vertices, key) {
		return errors.New("already exists a vertex with this key")
	}
	graph.vertices = append(graph.vertices, &Vertex[T]{key: key})
	return nil
}

func (graph *Graph[T]) AddEdge(from, to int) error {
	fromVertex := graph.getVertex(from)
	toVertex := graph.getVertex(to)

	if fromVertex == nil || toVertex == nil {
		return errors.New("invalid edge")
	} else if contains(fromVertex.adjacent, to) {
		return errors.New("already exists a edge with this key in this vertix")
	}
	fromVertex.adjacent = append(fromVertex.adjacent, toVertex)
	return nil
}
