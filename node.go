package gomcts

type Node struct {
	move         *Move
	parent       *Node
	state        GameState
	wins         uint64
	visits       uint64
	untriedMoves []*Move
	children     []*Node
}

func NewNode(move *Move, parent *Node, state GameState) *Node {
	untriedMoves := state.GetMoves()
	children := []*Node{}
	return &Node{move, parent, state, 0, 0, untriedMoves, children}
}

func (n Node) AddChild(move *Move, state GameState) *Node {
	for i := 0; i < len(n.untriedMoves); i++ {
		if n.untriedMoves[i] == move {
			n.untriedMoves, n.untriedMoves[len(n.untriedMoves)-1] =
				append(n.untriedMoves[:i], n.untriedMoves[i+1:]...), nil
		}
	}
	child := NewNode(move, &n, state)
	n.children = append(n.children, child)
	return child
}
