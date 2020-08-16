package assign

import (
	"github.com/i582/php2go/src/php/freefloating"
	"github.com/i582/php2go/src/php/node"
	"github.com/i582/php2go/src/php/position"
	"github.com/i582/php2go/src/php/walker"
)

// BitwiseXor node
type BitwiseXor struct {
	FreeFloating freefloating.Collection
	Position     *position.Position
	Variable     node.Node
	Expression   node.Node
}

// NewBitwiseXor node constructor
func NewBitwiseXor(Variable node.Node, Expression node.Node) *BitwiseXor {
	return &BitwiseXor{
		FreeFloating: nil,
		Variable:     Variable,
		Expression:   Expression,
	}
}

// SetPosition sets node position
func (n *BitwiseXor) SetPosition(p *position.Position) {
	n.Position = p
}

// GetPosition returns node positions
func (n *BitwiseXor) GetPosition() *position.Position {
	return n.Position
}

func (n *BitwiseXor) GetFreeFloating() *freefloating.Collection {
	return &n.FreeFloating
}

// Attributes returns node attributes as map
func (n *BitwiseXor) Attributes() map[string]interface{} {
	return nil
}

// Walk traverses nodes
// Walk is invoked recursively until v.EnterNode returns true
func (n *BitwiseXor) Walk(v walker.Visitor) {
	if v.EnterNode(n) == false {
		return
	}

	if n.Variable != nil {
		v.EnterChildNode("Variable", n)
		n.Variable.Walk(v)
		v.LeaveChildNode("Variable", n)
	}

	if n.Expression != nil {
		v.EnterChildNode("Expression", n)
		n.Expression.Walk(v)
		v.LeaveChildNode("Expression", n)
	}

	v.LeaveNode(n)
}
