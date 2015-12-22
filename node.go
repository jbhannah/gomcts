package gomcts

type Node struct {
	move         *Move
	parent       *Node
	state        GameState
	wins         uint64
	visits       uint64
	untriedMoves []Move
	children     []Node
}

func NewNode(move *Move, parent *Node, state GameState) *Node {
	untriedMoves := state.GetMoves()
	children := []Node{}
	return &Node{move, parent, state, 0, 0, untriedMoves, children}
}
