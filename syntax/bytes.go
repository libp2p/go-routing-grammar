package syntax

import (
	"bytes"
	"encoding/hex"
	"fmt"
	"io"

	"github.com/ipld/go-ipld-prime"
	xipld "github.com/libp2p/go-routing-language/syntax/ipld"
)

type Bytes struct {
	Bytes []byte
}

func (b Bytes) WritePretty(w io.Writer) error {
	_, err := fmt.Fprintf(w, "0x%s", hex.EncodeToString(b.Bytes)) // TODO: We can do better. E.g. wrap on 80-column boundary.
	return err
}

func IsEqualBytes(x, y Bytes) bool {
	return bytes.Equal(x.Bytes, y.Bytes)
}

// ToIPLD converts xr.Node into its corresponding IPLD Node type
func (b Bytes) ToIPLD() (ipld.Node, error) {
	return xipld.Type.Bytes_IPLD.FromBytes(b.Bytes)
}

// toNode_IPLD convert into IPLD Node of dynamic type NODE_IPLD
func (b Bytes) toNode_IPLD() (ipld.Node, error) {
	t := xipld.Type.Node_IPLD.NewBuilder()
	ma, err := t.BeginMap(-1)
	if err != nil {
		return nil, err
	}
	asm, err := ma.AssembleEntry("Bytes_IPLD")
	if err != nil {
		return nil, err
	}
	err = asm.AssignBytes(b.Bytes)
	if err != nil {
		return nil, err
	}
	if err := ma.Finish(); err != nil {
		return nil, err
	}
	return t.Build(), nil
}
