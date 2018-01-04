package name

import (
	"github.com/z7zmey/php-parser/node"
)

type FullyQualified struct {
	Name
}

func NewFullyQualified(Parts []node.Node) node.Node {
	return &FullyQualified{
		Name{
			map[string]interface{}{},
			nil,
			Parts,
		},
	}
}

func (n FullyQualified) Attributes() map[string]interface{} {
	return n.attributes
}
