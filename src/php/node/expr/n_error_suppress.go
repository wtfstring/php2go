package expr

import (
	"github.com/i582/php2go/src/php/freefloating"
	"github.com/i582/php2go/src/php/node"
	"github.com/i582/php2go/src/php/position"
	"github.com/i582/php2go/src/php/walker"
)

// ErrorSuppress node
type ErrorSuppress struct {
	FreeFloating freefloating.Collection
	Position     *position.Position
	Expr         node.Node
}

// NewErrorSuppress node constructor
func NewErrorSuppress(Expression node.Node) *ErrorSuppress {
	return &ErrorSuppress{
		FreeFloating: nil,
		Expr:         Expression,
	}
}

// SetPosition sets node position
func (n *ErrorSuppress) SetPosition(p *position.Position) {
	n.Position = p
}

// GetPosition returns node positions
func (n *ErrorSuppress) GetPosition() *position.Position {
	return n.Position
}

func (n *ErrorSuppress) GetFreeFloating() *freefloating.Collection {
	return &n.FreeFloating
}

// Attributes returns node attributes as map
func (n *ErrorSuppress) Attributes() map[string]interface{} {
	return nil
}

// Walk traverses nodes
// Walk is invoked recursively until v.EnterNode returns true
func (n *ErrorSuppress) Walk(v walker.Visitor) {
	if v.EnterNode(n) == false {
		return
	}

	if n.Expr != nil {
		v.EnterChildNode("Expr", n)
		n.Expr.Walk(v)
		v.LeaveChildNode("Expr", n)
	}

	v.LeaveNode(n)
}
