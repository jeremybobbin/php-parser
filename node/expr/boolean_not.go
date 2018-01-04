package expr

import (
	"github.com/z7zmey/php-parser/node"
)

type BooleanNot struct {
	attributes map[string]interface{}
	position   *node.Position
	expr       node.Node
}

func NewBooleanNot(Expression node.Node) node.Node {
	return &BooleanNot{
		map[string]interface{}{},
		nil,
		Expression,
	}
}

func (n BooleanNot) Attributes() map[string]interface{} {
	return n.attributes
}

func (n BooleanNot) Position() *node.Position {
	return n.position
}

func (n BooleanNot) SetPosition(p *node.Position) node.Node {
	n.position = p
	return n
}

func (n BooleanNot) Walk(v node.Visitor) {
	if v.EnterNode(n) == false {
		return
	}

	if n.expr != nil {
		vv := v.GetChildrenVisitor("expr")
		n.expr.Walk(vv)
	}

	v.LeaveNode(n)
}
