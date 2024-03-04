package tree

import "fmt"

type NodeIndex map[string]string

type IndexableNode[T any, N any, ID comparable] interface {
	BuildIndex(root Nodable[N, ID])
	Query(query T) (Nodable[N, ID], error)
}

type FlatIndex[T any, ID comparable] struct {
	index map[ID]Nodable[T, ID]
}

func (fi *FlatIndex[T, ID]) BuildIndex(root Nodable[T, ID]) {
	Walker(root, func(node Nodable[T, ID]) error {
		fi.index[node.GetID()] = node
		return nil
	}, LevelOrderSearch)
}

func (fi FlatIndex[T, ID]) Query(query ID) (Nodable[T, ID], error) {
	n, ok := fi.index[query]
	if !ok {
		return nil, fmt.Errorf("node not found")
	}
	return n, nil
}
