package stmt

import (
	"github.com/z7zmey/php-parser/node"
)

type PropertyList struct {
	attributes map[string]interface{}
	position   *node.Position
	Modifiers  []node.Node
	Properties []node.Node
}

func NewPropertyList(Modifiers []node.Node, Properties []node.Node) node.Node {
	return &PropertyList{
		map[string]interface{}{},
		nil,
		Modifiers,
		Properties,
	}
}

func (n PropertyList) Attributes() map[string]interface{} {
	return n.attributes
}

func (n PropertyList) Position() *node.Position {
	return n.position
}

func (n PropertyList) SetPosition(p *node.Position) node.Node {
	n.position = p
	return n
}

func (n PropertyList) Walk(v node.Visitor) {
	if v.EnterNode(n) == false {
		return
	}

	if n.Modifiers != nil {
		vv := v.GetChildrenVisitor("Modifiers")
		for _, nn := range n.Modifiers {
			nn.Walk(vv)
		}
	}

	if n.Properties != nil {
		vv := v.GetChildrenVisitor("Properties")
		for _, nn := range n.Properties {
			nn.Walk(vv)
		}
	}

	v.LeaveNode(n)
}
