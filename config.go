package main

import (
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
		"Xor",
	})
	strcases.SetPascalWords(map[string]string{
		"Mathml": "MathML",
		"MaX":    "Max",
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
	"lsmdom.h",
	"lsmdomdocument.h",
	"lsmdomdocumentfragment.h",
	"lsmdomnamednodemap.h",

	"lsmmathmlactionelement.h",
	"lsmmathmlaligngroupelement.h",
	"lsmmathmlalignmarkelement.h",
	"lsmmathmlattributes.h",
	"lsmmathmldocument.h",
	"lsmmathmlelement.h",
	"lsmmathmlencloseelement.h",
	"lsmmathmlenums.h",
	"lsmmathmlerrorelement.h",
	"lsmmathmlfencedelement.h",
	"lsmmathmlfractionelement.h",
	"lsmmathmlglyphtableams.h",
	"lsmmathmlitexelement.h",
	"lsmmathmllayoututils.h",
	"lsmmathmlmathelement.h",
	"lsmmathmloperatordictionary.h",
	"lsmmathmloperatorelement.h",
	"lsmmathmlpaddedelement.h",
	"lsmmathmlphantomelement.h",
	"lsmmathmlpresentationcontainer.h",
	"lsmmathmlpresentationtoken.h",
	"lsmmathmlradicalelement.h",
	"lsmmathmlrowelement.h",
	"lsmmathmlscriptelement.h",
	"lsmmathmlsemanticselement.h",
	"lsmmathmlspaceelement.h",
	"lsmmathmlstringelement.h",
	"lsmmathmlstyleelement.h",
	"lsmmathmlstyle.h",
	"lsmmathmltablecellelement.h",
	"lsmmathmltableelement.h",
	"lsmmathmltablerowelement.h",
	"lsmmathmltraits.h",
	"lsmmathmlunderoverelement.h",
	"lsmmathmlutils.h",
	"lsmmathmlview.h",

	"lsmsvgaelement.h",
	"lsmsvgattributes.h",
	"lsmsvgcircleelement.h",
	"lsmsvgclippathelement.h",
	"lsmsvgcolors.h",
	"lsmsvgdefselement.h",
	"lsmsvgdocument.h",
	"lsmsvgelement.h",
	"lsmsvgellipseelement.h",
	"lsmsvgenums.h",
	"lsmsvgfilterblend.h",
	"lsmsvgfiltercolormatrix.h",
	"lsmsvgfiltercomposite.h",
	"lsmsvgfilterconvolvematrix.h",
	"lsmsvgfilterdisplacementmap.h",
	"lsmsvgfilterelement.h",
	"lsmsvgfilterflood.h",
	"lsmsvgfiltergaussianblur.h",
	"lsmsvgfilterimage.h",
	"lsmsvgfiltermerge.h",
	"lsmsvgfiltermergenode.h",
	"lsmsvgfiltermorphology.h",
	"lsmsvgfilteroffset.h",
	"lsmsvgfilterprimitive.h",
	"lsmsvgfilterspecularlighting.h",
	"lsmsvgfiltersurface.h",
	"lsmsvgfiltertile.h",
	"lsmsvgfilterturbulence.h",
	"lsmsvggelement.h",
	"lsmsvggradientelement.h",
	"lsmsvgimageelement.h",
	"lsmsvglength.h",
	"lsmsvglineargradientelement.h",
	"lsmsvglineelement.h",
	"lsmsvgmarkerelement.h",
	"lsmsvgmaskelement.h",
	"lsmsvgmatrix.h",
	"lsmsvgpathelement.h",
	"lsmsvgpatternelement.h",
	"lsmsvgpolygonelement.h",
	"lsmsvgpolylineelement.h",
	"lsmsvgradialgradientelement.h",
	"lsmsvgrectelement.h",
	"lsmsvgstopelement.h",
	"lsmsvgstyle.h",
	"lsmsvgsvgelement.h",
	"lsmsvgswitchelement.h",
	"lsmsvgsymbolelement.h",
	"lsmsvgtextelement.h",
	"lsmsvgtraits.h",
	"lsmsvgtransformable.h",
	"lsmsvgtspanelement.h",
	"lsmsvguseelement.h",
	"lsmsvgview.h",
}

var postprocessors = map[string][]girgen.Postprocessor{}

var filters = []types.FilterMatcher{
	// These are not found for some reason.
	types.AbsoluteFilter("C.lsm_debug_level_get_type"),
	types.AbsoluteFilter("C.lsm_dom_node_type_get_type"),
	types.AbsoluteFilter("C.lsm_mathml_column_align_get_type"),
	types.AbsoluteFilter("C.lsm_mathml_css_type_get_type"),
	types.AbsoluteFilter("C.lsm_mathml_display_get_type"),
	types.AbsoluteFilter("C.lsm_mathml_font_get_type"),
	types.AbsoluteFilter("C.lsm_mathml_font_style_get_type"),
	types.AbsoluteFilter("C.lsm_mathml_font_weight_get_type"),
	types.AbsoluteFilter("C.lsm_mathml_form_get_type"),
	types.AbsoluteFilter("C.lsm_mathml_glyph_flags_get_type"),
	types.AbsoluteFilter("C.lsm_mathml_line_get_type"),
	types.AbsoluteFilter("C.lsm_mathml_linebreak_get_type"),
	types.AbsoluteFilter("C.lsm_mathml_mode_get_type"),
	types.AbsoluteFilter("C.lsm_mathml_notation_get_type"),
	types.AbsoluteFilter("C.lsm_mathml_presentation_token_type_get_type"),
	types.AbsoluteFilter("C.lsm_mathml_radical_element_type_get_type"),
	types.AbsoluteFilter("C.lsm_mathml_row_align_get_type"),
	types.AbsoluteFilter("C.lsm_mathml_script_element_type_get_type"),
	types.AbsoluteFilter("C.lsm_mathml_script_level_sign_get_type"),
	types.AbsoluteFilter("C.lsm_mathml_space_name_get_type"),
	types.AbsoluteFilter("C.lsm_mathml_table_row_element_type_get_type"),
	types.AbsoluteFilter("C.lsm_mathml_under_over_element_type_get_type"),
	types.AbsoluteFilter("C.lsm_mathml_unit_get_type"),
	types.AbsoluteFilter("C.lsm_mathml_variant_get_type"),
	types.AbsoluteFilter("C.lsm_svg_align_get_type"),
	types.AbsoluteFilter("C.lsm_svg_angle_type_get_type"),
	types.AbsoluteFilter("C.lsm_svg_blending_mode_get_type"),
	types.AbsoluteFilter("C.lsm_svg_channel_selector_get_type"),
	types.AbsoluteFilter("C.lsm_svg_color_filter_type_get_type"),
	types.AbsoluteFilter("C.lsm_svg_comp_op_get_type"),
	types.AbsoluteFilter("C.lsm_svg_display_get_type"),
	types.AbsoluteFilter("C.lsm_svg_edge_mode_get_type"),
	types.AbsoluteFilter("C.lsm_svg_element_category_get_type"),
	types.AbsoluteFilter("C.lsm_svg_enable_background_get_type"),
	types.AbsoluteFilter("C.lsm_svg_fill_rule_get_type"),
	types.AbsoluteFilter("C.lsm_svg_filter_input_get_type"),
	types.AbsoluteFilter("C.lsm_svg_font_stretch_get_type"),
	types.AbsoluteFilter("C.lsm_svg_font_style_get_type"),
	types.AbsoluteFilter("C.lsm_svg_font_weight_get_type"),
	types.AbsoluteFilter("C.lsm_svg_length_direction_get_type"),
	types.AbsoluteFilter("C.lsm_svg_length_type_get_type"),
	types.AbsoluteFilter("C.lsm_svg_line_cap_get_type"),
	types.AbsoluteFilter("C.lsm_svg_line_join_get_type"),
	types.AbsoluteFilter("C.lsm_svg_marker_units_get_type"),
	types.AbsoluteFilter("C.lsm_svg_matrix_flags_get_type"),
	types.AbsoluteFilter("C.lsm_svg_meet_or_slice_get_type"),
	types.AbsoluteFilter("C.lsm_svg_morphology_operator_get_type"),
	types.AbsoluteFilter("C.lsm_svg_overflow_get_type"),
	types.AbsoluteFilter("C.lsm_svg_paint_type_get_type"),
	types.AbsoluteFilter("C.lsm_svg_pattern_units_get_type"),
	types.AbsoluteFilter("C.lsm_svg_spread_method_get_type"),
	types.AbsoluteFilter("C.lsm_svg_stitch_tiles_get_type"),
	types.AbsoluteFilter("C.lsm_svg_text_anchor_get_type"),
	types.AbsoluteFilter("C.lsm_svg_transform_type_get_type"),
	types.AbsoluteFilter("C.lsm_svg_turbulence_type_get_type"),
	types.AbsoluteFilter("C.lsm_svg_use_element_flags_get_type"),
	types.AbsoluteFilter("C.lsm_svg_view_surface_type_get_type"),
	types.AbsoluteFilter("C.lsm_svg_visibility_get_type"),
	types.AbsoluteFilter("C.lsm_svg_writing_mode_get_type"),
	// These aren't actually implemented.
	types.AbsoluteFilter("C.lsm_mathml_view_draw_root"),
	types.AbsoluteFilter("C.lsm_svg_dash_array_copy"),
}

var preprocessors = []types.Preprocessor{
	types.PreprocessorFunc(func(repos gir.Repositories) {
		repo := repos.FromGIRFile("Lasem-0.7.gir")
		for _, incl := range lasemIncludes {
			repo.CIncludes = append(repo.CIncludes, gir.CInclude{Name: incl})
		}
	}),
	types.PreprocessorFunc(func(repos gir.Repositories) {
		repo := repos.FromGIRFile("Lasem-0.7.gir")
		namespace := repo.Namespaces[0]

		for _, class := range namespace.Classes {
			// Fix constructors returning nil without Nullable.
			for i := range class.Constructors {
				fixGError(&class.Constructors[i].CallableAttrs)
			}
		}
	}),
	modifyBufferInsert("Lasem-0.DomDocument.new_from_memory", "buffer"),
	modifyBufferInsert("Lasem-0.DomDocument.append_from_memory", "buffer"),
	modifyBufferInsert("Lasem-0.itex_to_mathml", "itex"),
}

func fixGError(callable *gir.CallableAttrs) {
	if callable.Throws && callable.ReturnValue != nil {
		if types.AnyTypeIsPtr(callable.ReturnValue.AnyType) {
			// Return type is a pointer, so it should be nullable.
			callable.ReturnValue.Nullable = true
			callable.ReturnValue.AllowNone = true
		}
	}
}

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
