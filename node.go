package tree

type Nodable[T any, ID comparable] interface {
	GetID() ID
	GetChilds() map[ID]Nodable[T, ID]
	GetValue() T
}

type MutableNode[T any, ID comparable] interface {
	Nodable[T, ID]
	AddNode(node ...Nodable[T, ID])
	DeleteNode(node ...string)
}

type Node[T any, ID comparable] struct {
	ID    ID                    `yaml:"id" json:"id"`
	Nodes map[ID]Nodable[T, ID] `yaml:"nodes" json:"nodes"`
	Value T
}

func (bn Node[T, ID]) GetID() ID {
	return bn.ID
}

func (bn Node[T, ID]) GetChilds() map[ID]Nodable[T, ID] {
	return bn.Nodes
}

func (bn Node[T, ID]) GetValue() T {
	return bn.Value
}

func (bn *Node[T, ID]) AddNode(node ...Nodable[T, ID]) {
	for _, n := range node {
		bn.Nodes[n.GetID()] = n
	}
}

func (bn *Node[T, ID]) DeleteNode(node ...ID) {
	for _, n := range node {
		delete(bn.Nodes, n)
	}
}

func (bn *Node[T, ID]) Walk(alg Searchable[T, ID], cb SearchableCallBack[T, ID]) {
	alg.Walk(bn, cb)
}

func NewNode[T any, ID comparable](id ID, value T, childs ...Nodable[T, ID]) Nodable[T, ID] {
	childsMap := make(map[ID]Nodable[T, ID], len(childs))
	for _, c := range childs {
		childsMap[c.GetID()] = c
	}
	return Node[T, ID]{
		ID:    id,
		Value: value,
		Nodes: childsMap,
	}
}
