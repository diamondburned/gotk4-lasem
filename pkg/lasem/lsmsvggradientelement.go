// Code generated by girgen. DO NOT EDIT.

package lasem

import (
	"unsafe"

	"github.com/diamondburned/gotk4/pkg/core/gextras"
	externglib "github.com/diamondburned/gotk4/pkg/core/glib"
)

// #include <stdlib.h>
// #include <glib-object.h>
// #include <lsm.h>
// #include <lsmdom.h>
// #include <lsmdomdocument.h>
// #include <lsmdomdocumentfragment.h>
// #include <lsmdomnamednodemap.h>
// #include <lsmmathmlactionelement.h>
// #include <lsmmathmlaligngroupelement.h>
// #include <lsmmathmlalignmarkelement.h>
// #include <lsmmathmlattributes.h>
// #include <lsmmathmldocument.h>
// #include <lsmmathmlelement.h>
// #include <lsmmathmlencloseelement.h>
// #include <lsmmathmlenums.h>
// #include <lsmmathmlerrorelement.h>
// #include <lsmmathmlfencedelement.h>
// #include <lsmmathmlfractionelement.h>
// #include <lsmmathmlglyphtableams.h>
// #include <lsmmathmlitexelement.h>
// #include <lsmmathmllayoututils.h>
// #include <lsmmathmlmathelement.h>
// #include <lsmmathmloperatordictionary.h>
// #include <lsmmathmloperatorelement.h>
// #include <lsmmathmlpaddedelement.h>
// #include <lsmmathmlphantomelement.h>
// #include <lsmmathmlpresentationcontainer.h>
// #include <lsmmathmlpresentationtoken.h>
// #include <lsmmathmlradicalelement.h>
// #include <lsmmathmlrowelement.h>
// #include <lsmmathmlscriptelement.h>
// #include <lsmmathmlsemanticselement.h>
// #include <lsmmathmlspaceelement.h>
// #include <lsmmathmlstringelement.h>
// #include <lsmmathmlstyle.h>
// #include <lsmmathmlstyleelement.h>
// #include <lsmmathmltablecellelement.h>
// #include <lsmmathmltableelement.h>
// #include <lsmmathmltablerowelement.h>
// #include <lsmmathmltraits.h>
// #include <lsmmathmlunderoverelement.h>
// #include <lsmmathmlutils.h>
// #include <lsmmathmlview.h>
// #include <lsmsvgaelement.h>
// #include <lsmsvgattributes.h>
// #include <lsmsvgcircleelement.h>
// #include <lsmsvgclippathelement.h>
// #include <lsmsvgcolors.h>
// #include <lsmsvgdefselement.h>
// #include <lsmsvgdocument.h>
// #include <lsmsvgelement.h>
// #include <lsmsvgellipseelement.h>
// #include <lsmsvgenums.h>
// #include <lsmsvgfilterblend.h>
// #include <lsmsvgfiltercolormatrix.h>
// #include <lsmsvgfiltercomposite.h>
// #include <lsmsvgfilterconvolvematrix.h>
// #include <lsmsvgfilterdisplacementmap.h>
// #include <lsmsvgfilterelement.h>
// #include <lsmsvgfilterflood.h>
// #include <lsmsvgfiltergaussianblur.h>
// #include <lsmsvgfilterimage.h>
// #include <lsmsvgfiltermerge.h>
// #include <lsmsvgfiltermergenode.h>
// #include <lsmsvgfiltermorphology.h>
// #include <lsmsvgfilteroffset.h>
// #include <lsmsvgfilterprimitive.h>
// #include <lsmsvgfilterspecularlighting.h>
// #include <lsmsvgfiltersurface.h>
// #include <lsmsvgfiltertile.h>
// #include <lsmsvgfilterturbulence.h>
// #include <lsmsvggelement.h>
// #include <lsmsvggradientelement.h>
// #include <lsmsvgimageelement.h>
// #include <lsmsvglength.h>
// #include <lsmsvglineargradientelement.h>
// #include <lsmsvglineelement.h>
// #include <lsmsvgmarkerelement.h>
// #include <lsmsvgmaskelement.h>
// #include <lsmsvgmatrix.h>
// #include <lsmsvgpathelement.h>
// #include <lsmsvgpatternelement.h>
// #include <lsmsvgpolygonelement.h>
// #include <lsmsvgpolylineelement.h>
// #include <lsmsvgradialgradientelement.h>
// #include <lsmsvgrectelement.h>
// #include <lsmsvgstopelement.h>
// #include <lsmsvgstyle.h>
// #include <lsmsvgsvgelement.h>
// #include <lsmsvgswitchelement.h>
// #include <lsmsvgsymbolelement.h>
// #include <lsmsvgtextelement.h>
// #include <lsmsvgtraits.h>
// #include <lsmsvgtransformable.h>
// #include <lsmsvgtspanelement.h>
// #include <lsmsvguseelement.h>
// #include <lsmsvgview.h>
import "C"

// glib.Type values for lsmsvggradientelement.go.
var GTypeSVGGradientElement = externglib.Type(C.lsm_svg_gradient_element_get_type())

func init() {
	externglib.RegisterGValueMarshalers([]externglib.TypeMarshaler{
		{T: GTypeSVGGradientElement, F: marshalSVGGradientElement},
	})
}

// SVGGradientElementOverrider contains methods that are overridable.
type SVGGradientElementOverrider interface {
}

type SVGGradientElement struct {
	_ [0]func() // equal guard
	SVGElement
}

var (
	_ SVGElementer = (*SVGGradientElement)(nil)
)

// SVGGradientElementer describes types inherited from class SVGGradientElement.
//
// To get the original type, the caller must assert this to an interface or
// another type.
type SVGGradientElementer interface {
	externglib.Objector
	baseSVGGradientElement() *SVGGradientElement
}

var _ SVGGradientElementer = (*SVGGradientElement)(nil)

func classInitSVGGradientElementer(gclassPtr, data C.gpointer) {
	C.g_type_class_add_private(gclassPtr, C.gsize(unsafe.Sizeof(uintptr(0))))

	goffset := C.g_type_class_get_instance_private_offset(gclassPtr)
	*(*C.gpointer)(unsafe.Add(unsafe.Pointer(gclassPtr), goffset)) = data

}

func wrapSVGGradientElement(obj *externglib.Object) *SVGGradientElement {
	return &SVGGradientElement{
		SVGElement: SVGElement{
			Object: obj,
		},
	}
}

func marshalSVGGradientElement(p uintptr) (interface{}, error) {
	return wrapSVGGradientElement(externglib.ValueFromNative(unsafe.Pointer(p)).Object()), nil
}

func (v *SVGGradientElement) baseSVGGradientElement() *SVGGradientElement {
	return v
}

// BaseSVGGradientElement returns the underlying base object.
func BaseSVGGradientElement(obj SVGGradientElementer) *SVGGradientElement {
	return obj.baseSVGGradientElement()
}

// SVGGradientElementAttributes: instance of this type is always passed by
// reference.
type SVGGradientElementAttributes struct {
	*svgGradientElementAttributes
}

// svgGradientElementAttributes is the struct that's finalized.
type svgGradientElementAttributes struct {
	native *C.LsmSvgGradientElementAttributes
}

func (s *SVGGradientElementAttributes) Transform() *SVGMatrix {
	var v *SVGMatrix // out
	v = (*SVGMatrix)(gextras.NewStructNative(unsafe.Pointer((&s.native.transform))))
	return v
}

func (s *SVGGradientElementAttributes) Units() SVGPatternUnits {
	var v SVGPatternUnits // out
	v = SVGPatternUnits(s.native.units)
	return v
}

func (s *SVGGradientElementAttributes) SpreadMethod() SVGSpreadMethod {
	var v SVGSpreadMethod // out
	v = SVGSpreadMethod(s.native.spread_method)
	return v
}