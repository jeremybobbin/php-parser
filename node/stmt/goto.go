package stmt

import (
	"github.com/z7zmey/php-parser/node"
)

type Goto struct {
	attributes map[string]interface{}
	position   *node.Position
	Label      node.Node
}

func NewGoto(Label node.Node) node.Node {
	return &Goto{
		map[string]interface{}{},
		nil,
		Label,
	}
}

func (n Goto) Attributes() map[string]interface{} {
	return n.attributes
}

func (n Goto) Position() *node.Position {
	return n.position
}

func (n Goto) SetPosition(p *node.Position) node.Node {
	n.position = p
	return n
}

func (n Goto) Walk(v node.Visitor) {
	if v.EnterNode(n) == false {
		return
	}

	if n.Label != nil {
		vv := v.GetChildrenVisitor("Label")
		n.Label.Walk(vv)
	}

	v.LeaveNode(n)
}
