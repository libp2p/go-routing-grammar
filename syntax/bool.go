package syntax

import (
	"fmt"
	"io"

	"github.com/ipld/go-ipld-prime"
	xipld "github.com/libp2p/go-routing-language/syntax/ipld"
)

type Bool struct {
	Value bool
}

func (b Bool) WritePretty(w io.Writer) (err error) {
	_, err = fmt.Fprintf(w, "%v", b.Value)
	return err
}

func IsEqualBool(x, y Bool) bool {
	return x.Value == y.Value
}

// ToIPLD converts xr.Node into its corresponding IPLD Node type
func (b Bool) ToIPLD() (ipld.Node, error) {
	t := xipld.Type.Bool_IPLD.NewBuilder()
	err := t.AssignBool(b.Value)
	if err != nil {
		return nil, err
	}
	return t.Build(), nil
}

// toNode_IPLD convert into IPLD Node of dynamic type NODE_IPLD
func (b Bool) toNode_IPLD() (ipld.Node, error) {
	t := xipld.Type.Node_IPLD.NewBuilder()
	ma, err := t.BeginMap(-1)
	if err != nil {
		return nil, err
	}
	asm, err := ma.AssembleEntry("Bool_IPLD")
	if err != nil {
		return nil, err
	}
	err = asm.AssignBool(b.Value)
	if err != nil {
		return nil, err
	}
	if err := ma.Finish(); err != nil {
		return nil, err
	}
	return t.Build(), nil
}
