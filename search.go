package tree

import "golang.org/x/exp/maps"

type SearchableCallBack[T any, ID comparable] func(node Nodable[T, ID]) error

type Searchable[T any, ID comparable] interface {
	Walk(root Nodable[T, ID], cb SearchableCallBack[T, ID])
}

type SearchableFunc[T any, ID comparable] func(root Nodable[T, ID], cb SearchableCallBack[T, ID])

func (sf SearchableFunc[T, ID]) Walk(root Nodable[T, ID], cb SearchableCallBack[T, ID]) {
	sf(root, cb)
}

func LevelOrderSearch[T any, ID comparable]() Searchable[T, ID] {
	return SearchableFunc[T, ID](func(root Nodable[T, ID], cb SearchableCallBack[T, ID]) {
		nextNodes := []Nodable[T, ID]{root}
		var currentNode Nodable[T, ID] = nil
		for len(nextNodes) > 0 {
			currentNode = nextNodes[0]
			sliceChilds := maps.Values(currentNode.GetChilds())
			nextNodes = append(nextNodes[1:], sliceChilds...)
			cb(currentNode)
		}
	})
}

func Walker[T any, ID comparable, A ~func() Searchable[T, ID]](root Nodable[T, ID], nextNodesCb SearchableCallBack[T, ID], alg A) {
	alg().Walk(root, nextNodesCb)
}
