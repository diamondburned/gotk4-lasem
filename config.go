package main

import (
	"log"
	"strings"

	"github.com/diamondburned/gotk4/gir"
	"github.com/diamondburned/gotk4/gir/cmd/gir_generate/gendata"
	"github.com/diamondburned/gotk4/gir/girgen"
	"github.com/diamondburned/gotk4/gir/girgen/strcases"
	"github.com/diamondburned/gotk4/gir/girgen/types"
)

const (
	gotk4Module = "github.com/diamondburned/gotk4/pkg"
	thisModule  = "github.com/diamondburned/gotk4-lasem/pkg"
)

func init() {
	strcases.AddPascalSpecials([]string{
		"Dom",
	})
	strcases.SetPascalWords(map[string]string{
		"Mathml": "MathML",
	})
}

var externPkgs = []gendata.Package{
	{PkgName: "gobject-introspection-1.0", Namespaces: []string{
		"GLib-2",
		"GObject-2",
		"Gio-2",
		"cairo-1",
	}},
	{PkgName: "gdk-pixbuf-2.0", Namespaces: nil},
	{PkgName: "pango", Namespaces: []string{
		"Pango-1",
		"PangoCairo-1",
	}},
}

var packages = []gendata.Package{
	{PkgName: "lasem-0.7", Namespaces: nil},
}

var pkgGenerated = []string{
	"lasem",
}

var pkgExceptions = []string{
	"go.mod",
	"go.sum",
	"LICENSE",
}

var lasemIncludes = []string{
	// Already added.
	// "lsm.h",
	"lsmdom.h",
	"lsmdomdocument.h",
	"lsmdomdocumentfragment.h",
	"lsmdomnamednodemap.h",
}

var preprocessors = []types.Preprocessor{
	types.PreprocessorFunc(func(repos gir.Repositories) {
		repo := repos.FromGIRFile("Lasem-0.7.gir")
		if repo == nil {
			log.Panicf("Lasem GIR not found")
		}

		for _, incl := range lasemIncludes {
			repo.CIncludes = append(repo.CIncludes, gir.CInclude{Name: incl})
		}
	}),
	modifyBufferInsert("Lasem-0.DomDocument.new_from_memory", "buffer"),
	modifyBufferInsert("Lasem-0.DomDocument.append_from_memory", "buffer"),
	modifyBufferInsert("Lasem-0.itex_to_mathml", "itex"),
}

var postprocessors = map[string][]girgen.Postprocessor{}

var filters = []types.FilterMatcher{}

// Taken and modified from gotk4.
func modifyBufferInsert(girType, param string) types.Preprocessor {
	return types.ModifyCallable(girType, func(c *gir.CallableAttrs) {
		p := types.FindParameter(c, param)
		if p == nil {
			return
		}

		lenIx := findTextLenParam(c.Parameters.Parameters)
		if lenIx == -1 {
			return
		}

		p.Type = nil
		p.Array = &gir.Array{
			CType:          "const char*",
			Type:           &gir.Type{Name: "gchar"},
			Length:         &lenIx,
			ZeroTerminated: new(bool), // false
		}
	})
}

func findTextLenParam(params []gir.Parameter) int {
	const doc = "size of"

	for i, param := range params {
		if param.Name == "size" || (param.Doc != nil && strings.Contains(param.Doc.String, doc)) {
			return i
		}
	}

	return -1
}
