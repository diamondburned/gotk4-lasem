package main

import (
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
	{PkgName: "lasem-0.4", Namespaces: nil},
}

var pkgGenerated = []string{
	"lasem",
}

var pkgExceptions = []string{
	"go.mod",
	"go.sum",
	"LICENSE",
}

var postprocessors = map[string][]girgen.Postprocessor{}

var filters = []types.FilterMatcher{}
