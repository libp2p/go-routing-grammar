package syntax

import (
	"fmt"
	"io"

	"github.com/ipld/go-ipld-prime"
	xipld "github.com/libp2p/go-routing-language/syntax/ipld"
)

// String is a node representing a string literal.
type String struct {
	Value string
}

func (s String) WritePretty(w io.Writer) error {
	_, err := fmt.Fprintf(w, "%q", s.Value)
	return err
}

func IsEqualString(x, y String) bool {
	return x.Value == y.Value
}

// ToIPLD converts xr.Node into its corresponding IPLD Node type
func (s String) ToIPLD() (ipld.Node, error) {
	t := xipld.Type.String_IPLD.NewBuilder()
	err := t.AssignString(s.Value)
	if err != nil {
		return nil, err
	}
	return t.Build(), nil
}

// toNode_IPLD convert into IPLD Node of dynamic type NODE_IPLD
func (s String) toNode_IPLD() (ipld.Node, error) {
	t := xipld.Type.Node_IPLD.NewBuilder()
	ma, err := t.BeginMap(-1)
	if err != nil {
		return nil, err
	}
	asm, err := ma.AssembleEntry("String_IPLD")
	if err != nil {
		return nil, err
	}
	err = asm.AssignString(s.Value)
	if err != nil {
		return nil, err
	}
	if err := ma.Finish(); err != nil {
		return nil, err
	}
	return t.Build(), nil
}
