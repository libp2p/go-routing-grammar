//go:build ignore

package main

import (
	"fmt"
	"os"

	"github.com/ipld/go-ipld-prime/schema"
	gengo "github.com/ipld/go-ipld-prime/schema/gen/go"
)

func main() {
	ts := schema.TypeSystem{}
	ts.Init()
	adjCfg := &gengo.AdjunctCfg{
		CfgUnionMemlayout: map[schema.TypeName]string{
			"Node_IPLD": "interface",
		},
	}

	pkgName := "ipld"

	// This needs to preclude to have Type Kinds available.
	ts.Accumulate(schema.SpawnString("String"))

	ts.Accumulate(schema.SpawnString("String_IPLD"))
	ts.Accumulate(schema.SpawnBytes("Bytes_IPLD"))
	ts.Accumulate(schema.SpawnInt("Int_IPLD"))
	ts.Accumulate(schema.SpawnFloat("Float_IPLD"))
	ts.Accumulate(schema.SpawnBool("Bool_IPLD"))

	ts.Accumulate(schema.SpawnUnion("Node_IPLD",
		[]schema.TypeName{
			"String_IPLD",
			"Bytes_IPLD",
			"Int_IPLD",
			"Float_IPLD",
			"Bool_IPLD",
			"Dict_IPLD",
			"List_IPLD",
			"Predicate_IPLD",
		},

		// Keys for the JSON representation of types
		schema.SpawnUnionRepresentationKeyed(map[string]schema.TypeName{
			"String":    "String_IPLD",
			"Bytes":     "Bytes_IPLD",
			"Int":       "Int_IPLD",
			"Float":     "Float_IPLD",
			"Bool":      "Bool_IPLD",
			"Dict":      "Dict_IPLD",
			"List":      "List_IPLD",
			"Predicate": "Predicate_IPLD",
		}),
	))

	// DICT
	ts.Accumulate(schema.SpawnStruct("Dict_IPLD",
		[]schema.StructField{
			schema.SpawnStructField("Pairs", "Pairs_IPLD", false, false),
		},
		schema.SpawnStructRepresentationMap(nil),
	))
	ts.Accumulate(schema.SpawnList("Pairs_IPLD",
		"Pair_IPLD", false,
	))
	ts.Accumulate(schema.SpawnStruct("Pair_IPLD",
		[]schema.StructField{
			schema.SpawnStructField("Key", "Node_IPLD", false, false),
			schema.SpawnStructField("Value", "Node_IPLD", false, false),
		},
		schema.SpawnStructRepresentationMap(nil),
	))

	// LIST
	ts.Accumulate(schema.SpawnStruct("List_IPLD",
		[]schema.StructField{
			schema.SpawnStructField("Elements", "Nodes_IPLD", false, false),
		},
		schema.SpawnStructRepresentationMap(nil),
	))
	ts.Accumulate(schema.SpawnList("Nodes_IPLD",
		"Node_IPLD", false,
	))

	// PREDICATE
	ts.Accumulate(schema.SpawnStruct("Predicate_IPLD",
		[]schema.StructField{
			schema.SpawnStructField("Tag", "String", false, false),
			schema.SpawnStructField("Positional", "Nodes_IPLD", false, false),
			schema.SpawnStructField("Named", "Pairs_IPLD", false, false),
		},
		schema.SpawnStructRepresentationMap(nil),
	))

	if errs := ts.ValidateGraph(); errs != nil {
		for _, err := range errs {
			fmt.Printf("- %s\n", err)
		}
		os.Exit(1)
	}

	gengo.Generate(".", pkgName, ts, adjCfg)
}
