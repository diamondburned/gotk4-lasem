// Code generated by girgen. DO NOT EDIT.

package lasem

import (
	"runtime"
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

// glib.Type values for lsmmathmlmathelement.go.
var GTypeMathMLMathElement = externglib.Type(C.lsm_mathml_math_element_get_type())

func init() {
	externglib.RegisterGValueMarshalers([]externglib.TypeMarshaler{
		{T: GTypeMathMLMathElement, F: marshalMathMLMathElement},
	})
}

const MATHML_FONT_DOUBLE_STRUCK = "msbm10"
const MATHML_FONT_FRAKTUR = "Serif"
const MATHML_FONT_MONOSPACE = "Monospace"
const MATHML_FONT_SANS = "Sans"
const MATHML_FONT_SCRIPT = "cmsy10"
const MATHML_FONT_SERIF = "Serif"

// MathMLMathElementOverrider contains methods that are overridable.
type MathMLMathElementOverrider interface {
}

type MathMLMathElement struct {
	_ [0]func() // equal guard
	MathMLElement
}

var (
	_ MathMLElementer = (*MathMLMathElement)(nil)
)

func classInitMathMLMathElementer(gclassPtr, data C.gpointer) {
	C.g_type_class_add_private(gclassPtr, C.gsize(unsafe.Sizeof(uintptr(0))))

	goffset := C.g_type_class_get_instance_private_offset(gclassPtr)
	*(*C.gpointer)(unsafe.Add(unsafe.Pointer(gclassPtr), goffset)) = data

}

func wrapMathMLMathElement(obj *externglib.Object) *MathMLMathElement {
	return &MathMLMathElement{
		MathMLElement: MathMLElement{
			DOMElement: DOMElement{
				DOMNode: DOMNode{
					Object: obj,
				},
			},
		},
	}
}

func marshalMathMLMathElement(p uintptr) (interface{}, error) {
	return wrapMathMLMathElement(externglib.ValueFromNative(unsafe.Pointer(p)).Object()), nil
}

// The function returns the following values:
//
func NewMathMLMathElement() *MathMLMathElement {
	var _cret *C.LsmDomNode // in

	_cret = C.lsm_mathml_math_element_new()

	var _mathmlMathElement *MathMLMathElement // out

	_mathmlMathElement = wrapMathMLMathElement(externglib.AssumeOwnership(unsafe.Pointer(_cret)))

	return _mathmlMathElement
}

// The function returns the following values:
//
func (mathElement *MathMLMathElement) DefaultStyle() *MathMLStyle {
	var _arg0 *C.LsmMathmlMathElement // out
	var _cret *C.LsmMathmlStyle       // in

	_arg0 = (*C.LsmMathmlMathElement)(unsafe.Pointer(externglib.InternObject(mathElement).Native()))

	_cret = C.lsm_mathml_math_element_get_default_style(_arg0)
	runtime.KeepAlive(mathElement)

	var _mathmlStyle *MathMLStyle // out

	_mathmlStyle = (*MathMLStyle)(gextras.NewStructNative(unsafe.Pointer(_cret)))
	runtime.SetFinalizer(
		gextras.StructIntern(unsafe.Pointer(_mathmlStyle)),
		func(intern *struct{ C unsafe.Pointer }) {
			C.lsm_mathml_style_free((*C.LsmMathmlStyle)(intern.C))
		},
	)

	return _mathmlStyle
}

// The function takes the following parameters:
//
//    - view
//    - bbox
//
func (mathElement *MathMLMathElement) Layout(view *MathMLView, bbox *MathMLBbox) {
	var _arg0 *C.LsmMathmlMathElement // out
	var _arg1 *C.LsmMathmlView        // out
	var _arg2 *C.LsmMathmlBbox        // out

	_arg0 = (*C.LsmMathmlMathElement)(unsafe.Pointer(externglib.InternObject(mathElement).Native()))
	_arg1 = (*C.LsmMathmlView)(unsafe.Pointer(externglib.InternObject(view).Native()))
	_arg2 = (*C.LsmMathmlBbox)(gextras.StructNative(unsafe.Pointer(bbox)))

	C.lsm_mathml_math_element_layout(_arg0, _arg1, _arg2)
	runtime.KeepAlive(mathElement)
	runtime.KeepAlive(view)
	runtime.KeepAlive(bbox)
}

// The function takes the following parameters:
//
// The function returns the following values:
//
func (mathElement *MathMLMathElement) Measure(view *MathMLView) *MathMLBbox {
	var _arg0 *C.LsmMathmlMathElement // out
	var _arg1 *C.LsmMathmlView        // out
	var _cret *C.LsmMathmlBbox        // in

	_arg0 = (*C.LsmMathmlMathElement)(unsafe.Pointer(externglib.InternObject(mathElement).Native()))
	_arg1 = (*C.LsmMathmlView)(unsafe.Pointer(externglib.InternObject(view).Native()))

	_cret = C.lsm_mathml_math_element_measure(_arg0, _arg1)
	runtime.KeepAlive(mathElement)
	runtime.KeepAlive(view)

	var _mathmlBbox *MathMLBbox // out

	_mathmlBbox = (*MathMLBbox)(gextras.NewStructNative(unsafe.Pointer(_cret)))

	return _mathmlBbox
}

// The function takes the following parameters:
//
func (mathElement *MathMLMathElement) Render(view *MathMLView) {
	var _arg0 *C.LsmMathmlMathElement // out
	var _arg1 *C.LsmMathmlView        // out

	_arg0 = (*C.LsmMathmlMathElement)(unsafe.Pointer(externglib.InternObject(mathElement).Native()))
	_arg1 = (*C.LsmMathmlView)(unsafe.Pointer(externglib.InternObject(view).Native()))

	C.lsm_mathml_math_element_render(_arg0, _arg1)
	runtime.KeepAlive(mathElement)
	runtime.KeepAlive(view)
}

// The function takes the following parameters:
//
func (mathElement *MathMLMathElement) SetDefaultStyle(style *MathMLStyle) {
	var _arg0 *C.LsmMathmlMathElement // out
	var _arg1 *C.LsmMathmlStyle       // out

	_arg0 = (*C.LsmMathmlMathElement)(unsafe.Pointer(externglib.InternObject(mathElement).Native()))
	_arg1 = (*C.LsmMathmlStyle)(gextras.StructNative(unsafe.Pointer(style)))

	C.lsm_mathml_math_element_set_default_style(_arg0, _arg1)
	runtime.KeepAlive(mathElement)
	runtime.KeepAlive(style)
}

func (mathElement *MathMLMathElement) Update() {
	var _arg0 *C.LsmMathmlMathElement // out

	_arg0 = (*C.LsmMathmlMathElement)(unsafe.Pointer(externglib.InternObject(mathElement).Native()))

	C.lsm_mathml_math_element_update(_arg0)
	runtime.KeepAlive(mathElement)
}
