package gomcts

import (
	"math"
	"sort"
)

type Node struct {
	move         *Move
	parent       *Node
	state        GameState
	wins         uint64
	visits       uint64
	untriedMoves []*Move
	children     ChildNodes
}

type ChildNodes []*Node

func (children ChildNodes) Len() int {
	return len(children)
}

func (children ChildNodes) Less(i, j int) bool {
	return children[i].UctValue() < children[j].UctValue()
}

func (children ChildNodes) Swap(i, j int) {
	children[i], children[j] = children[j], children[i]
}

func NewNode(move *Move, parent *Node, state GameState) *Node {
	untriedMoves := state.GetMoves()
	children := []*Node{}
	return &Node{move, parent, state, 0, 0, untriedMoves, children}
}

func (n Node) AddChild(move *Move, state GameState) *Node {
	for i := 0; i < len(n.untriedMoves); i++ {
		if *n.untriedMoves[i] == *move {
			n.untriedMoves, n.untriedMoves[len(n.untriedMoves)-1] =
				append(n.untriedMoves[:i], n.untriedMoves[i+1:]...), nil
		}
	}
	child := NewNode(move, &n, state)
	n.children = append(n.children, child)
	return child
}

func (n Node) UctValue() float64 {
	return float64(n.wins)/float64(n.visits) + math.Sqrt(2*math.Log(float64(n.parent.visits))/float64(n.visits))
}

func (n Node) SelectChild() *Node {
	sort.Sort(n.children)
	return n.children[len(n.children)-1]
}
