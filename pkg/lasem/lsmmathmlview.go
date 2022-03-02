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

// glib.Type values for lsmmathmlview.go.
var GTypeMathMLView = externglib.Type(C.lsm_mathml_view_get_type())

func init() {
	externglib.RegisterGValueMarshalers([]externglib.TypeMarshaler{
		{T: GTypeMathMLView, F: marshalMathMLView},
	})
}

// MathMLViewOverrider contains methods that are overridable.
type MathMLViewOverrider interface {
}

type MathMLView struct {
	_ [0]func() // equal guard
	*externglib.Object
}

var (
	_ externglib.Objector = (*MathMLView)(nil)
)

func classInitMathMLViewer(gclassPtr, data C.gpointer) {
	C.g_type_class_add_private(gclassPtr, C.gsize(unsafe.Sizeof(uintptr(0))))

	goffset := C.g_type_class_get_instance_private_offset(gclassPtr)
	*(*C.gpointer)(unsafe.Add(unsafe.Pointer(gclassPtr), goffset)) = data

}

func wrapMathMLView(obj *externglib.Object) *MathMLView {
	return &MathMLView{
		Object: obj,
	}
}

func marshalMathMLView(p uintptr) (interface{}, error) {
	return wrapMathMLView(externglib.ValueFromNative(unsafe.Pointer(p)).Object()), nil
}

// The function takes the following parameters:
//
// The function returns the following values:
//
func NewMathMLView(document *MathMLDocument) *MathMLView {
	var _arg1 *C.LsmMathmlDocument // out
	var _cret *C.LsmMathmlView     // in

	_arg1 = (*C.LsmMathmlDocument)(unsafe.Pointer(externglib.InternObject(document).Native()))

	_cret = C.lsm_mathml_view_new(_arg1)
	runtime.KeepAlive(document)

	var _mathmlView *MathMLView // out

	_mathmlView = wrapMathMLView(externglib.AssumeOwnership(unsafe.Pointer(_cret)))

	return _mathmlView
}

// The function takes the following parameters:
//
//    - style
//    - ascent
//    - descent
//
func (view *MathMLView) FontMetrics(style *MathMLElementStyle, ascent, descent *float64) {
	var _arg0 *C.LsmMathmlView         // out
	var _arg1 *C.LsmMathmlElementStyle // out
	var _arg2 *C.double                // out
	var _arg3 *C.double                // out

	_arg0 = (*C.LsmMathmlView)(unsafe.Pointer(externglib.InternObject(view).Native()))
	_arg1 = (*C.LsmMathmlElementStyle)(gextras.StructNative(unsafe.Pointer(style)))
	_arg2 = (*C.double)(unsafe.Pointer(ascent))
	_arg3 = (*C.double)(unsafe.Pointer(descent))

	C.lsm_mathml_view_get_font_metrics(_arg0, _arg1, _arg2, _arg3)
	runtime.KeepAlive(view)
	runtime.KeepAlive(style)
	runtime.KeepAlive(ascent)
	runtime.KeepAlive(descent)
}

// The function takes the following parameters:
//
//    - style
//    - text
//
// The function returns the following values:
//
func (view *MathMLView) OperatorSlant(style *MathMLElementStyle, text string) float64 {
	var _arg0 *C.LsmMathmlView         // out
	var _arg1 *C.LsmMathmlElementStyle // out
	var _arg2 *C.char                  // out
	var _cret C.double                 // in

	_arg0 = (*C.LsmMathmlView)(unsafe.Pointer(externglib.InternObject(view).Native()))
	_arg1 = (*C.LsmMathmlElementStyle)(gextras.StructNative(unsafe.Pointer(style)))
	_arg2 = (*C.char)(unsafe.Pointer(C.CString(text)))
	defer C.free(unsafe.Pointer(_arg2))

	_cret = C.lsm_mathml_view_get_operator_slant(_arg0, _arg1, _arg2)
	runtime.KeepAlive(view)
	runtime.KeepAlive(style)
	runtime.KeepAlive(text)

	var _gdouble float64 // out

	_gdouble = float64(_cret)

	return _gdouble
}

// The function takes the following parameters:
//
// The function returns the following values:
//
func (view *MathMLView) MeasureAxisOffset(mathSize float64) float64 {
	var _arg0 *C.LsmMathmlView // out
	var _arg1 C.double         // out
	var _cret C.double         // in

	_arg0 = (*C.LsmMathmlView)(unsafe.Pointer(externglib.InternObject(view).Native()))
	_arg1 = C.double(mathSize)

	_cret = C.lsm_mathml_view_measure_axis_offset(_arg0, _arg1)
	runtime.KeepAlive(view)
	runtime.KeepAlive(mathSize)

	var _gdouble float64 // out

	_gdouble = float64(_cret)

	return _gdouble
}

// The function takes the following parameters:
//
//    - style
//    - notation
//    - stretchBbox
//    - bbox
//    - xChildOffset
//
func (view *MathMLView) MeasureNotation(style *MathMLElementStyle, notation MathMLNotation, stretchBbox, bbox *MathMLBbox, xChildOffset *float64) {
	var _arg0 *C.LsmMathmlView         // out
	var _arg1 *C.LsmMathmlElementStyle // out
	var _arg2 C.LsmMathmlNotation      // out
	var _arg3 *C.LsmMathmlBbox         // out
	var _arg4 *C.LsmMathmlBbox         // out
	var _arg5 *C.double                // out

	_arg0 = (*C.LsmMathmlView)(unsafe.Pointer(externglib.InternObject(view).Native()))
	_arg1 = (*C.LsmMathmlElementStyle)(gextras.StructNative(unsafe.Pointer(style)))
	_arg2 = C.LsmMathmlNotation(notation)
	_arg3 = (*C.LsmMathmlBbox)(gextras.StructNative(unsafe.Pointer(stretchBbox)))
	_arg4 = (*C.LsmMathmlBbox)(gextras.StructNative(unsafe.Pointer(bbox)))
	_arg5 = (*C.double)(unsafe.Pointer(xChildOffset))

	C.lsm_mathml_view_measure_notation(_arg0, _arg1, _arg2, _arg3, _arg4, _arg5)
	runtime.KeepAlive(view)
	runtime.KeepAlive(style)
	runtime.KeepAlive(notation)
	runtime.KeepAlive(stretchBbox)
	runtime.KeepAlive(bbox)
	runtime.KeepAlive(xChildOffset)
}

// The function takes the following parameters:
//
//    - style
//    - text
//    - large
//    - symmetric
//    - axisOffset
//    - minSize
//    - maxSize
//    - stretchBbox
//    - bbox
//
func (view *MathMLView) MeasureOperator(style *MathMLElementStyle, text string, large, symmetric bool, axisOffset, minSize, maxSize float64, stretchBbox, bbox *MathMLBbox) {
	var _arg0 *C.LsmMathmlView         // out
	var _arg1 *C.LsmMathmlElementStyle // out
	var _arg2 *C.char                  // out
	var _arg3 C.gboolean               // out
	var _arg4 C.gboolean               // out
	var _arg5 C.double                 // out
	var _arg6 C.double                 // out
	var _arg7 C.double                 // out
	var _arg8 *C.LsmMathmlBbox         // out
	var _arg9 *C.LsmMathmlBbox         // out

	_arg0 = (*C.LsmMathmlView)(unsafe.Pointer(externglib.InternObject(view).Native()))
	_arg1 = (*C.LsmMathmlElementStyle)(gextras.StructNative(unsafe.Pointer(style)))
	_arg2 = (*C.char)(unsafe.Pointer(C.CString(text)))
	defer C.free(unsafe.Pointer(_arg2))
	if large {
		_arg3 = C.TRUE
	}
	if symmetric {
		_arg4 = C.TRUE
	}
	_arg5 = C.double(axisOffset)
	_arg6 = C.double(minSize)
	_arg7 = C.double(maxSize)
	_arg8 = (*C.LsmMathmlBbox)(gextras.StructNative(unsafe.Pointer(stretchBbox)))
	_arg9 = (*C.LsmMathmlBbox)(gextras.StructNative(unsafe.Pointer(bbox)))

	C.lsm_mathml_view_measure_operator(_arg0, _arg1, _arg2, _arg3, _arg4, _arg5, _arg6, _arg7, _arg8, _arg9)
	runtime.KeepAlive(view)
	runtime.KeepAlive(style)
	runtime.KeepAlive(text)
	runtime.KeepAlive(large)
	runtime.KeepAlive(symmetric)
	runtime.KeepAlive(axisOffset)
	runtime.KeepAlive(minSize)
	runtime.KeepAlive(maxSize)
	runtime.KeepAlive(stretchBbox)
	runtime.KeepAlive(bbox)
}

// The function takes the following parameters:
//
//    - style
//    - stretchBbox
//    - bbox
//    - xOffset
//    - yOffset
//
func (view *MathMLView) MeasureRadical(style *MathMLElementStyle, stretchBbox, bbox *MathMLBbox, xOffset, yOffset *float64) {
	var _arg0 *C.LsmMathmlView         // out
	var _arg1 *C.LsmMathmlElementStyle // out
	var _arg2 *C.LsmMathmlBbox         // out
	var _arg3 *C.LsmMathmlBbox         // out
	var _arg4 *C.double                // out
	var _arg5 *C.double                // out

	_arg0 = (*C.LsmMathmlView)(unsafe.Pointer(externglib.InternObject(view).Native()))
	_arg1 = (*C.LsmMathmlElementStyle)(gextras.StructNative(unsafe.Pointer(style)))
	_arg2 = (*C.LsmMathmlBbox)(gextras.StructNative(unsafe.Pointer(stretchBbox)))
	_arg3 = (*C.LsmMathmlBbox)(gextras.StructNative(unsafe.Pointer(bbox)))
	_arg4 = (*C.double)(unsafe.Pointer(xOffset))
	_arg5 = (*C.double)(unsafe.Pointer(yOffset))

	C.lsm_mathml_view_measure_radical(_arg0, _arg1, _arg2, _arg3, _arg4, _arg5)
	runtime.KeepAlive(view)
	runtime.KeepAlive(style)
	runtime.KeepAlive(stretchBbox)
	runtime.KeepAlive(bbox)
	runtime.KeepAlive(xOffset)
	runtime.KeepAlive(yOffset)
}

// The function takes the following parameters:
//
//    - style
//    - text
//    - bbox
//
func (view *MathMLView) MeasureText(style *MathMLElementStyle, text string, bbox *MathMLBbox) {
	var _arg0 *C.LsmMathmlView         // out
	var _arg1 *C.LsmMathmlElementStyle // out
	var _arg2 *C.char                  // out
	var _arg3 *C.LsmMathmlBbox         // out

	_arg0 = (*C.LsmMathmlView)(unsafe.Pointer(externglib.InternObject(view).Native()))
	_arg1 = (*C.LsmMathmlElementStyle)(gextras.StructNative(unsafe.Pointer(style)))
	_arg2 = (*C.char)(unsafe.Pointer(C.CString(text)))
	defer C.free(unsafe.Pointer(_arg2))
	_arg3 = (*C.LsmMathmlBbox)(gextras.StructNative(unsafe.Pointer(bbox)))

	C.lsm_mathml_view_measure_text(_arg0, _arg1, _arg2, _arg3)
	runtime.KeepAlive(view)
	runtime.KeepAlive(style)
	runtime.KeepAlive(text)
	runtime.KeepAlive(bbox)
}

// The function takes the following parameters:
//
//    - style
//    - x
//    - y
//    - bbox
//
func (view *MathMLView) ShowBackground(style *MathMLElementStyle, x, y float64, bbox *MathMLBbox) {
	var _arg0 *C.LsmMathmlView         // out
	var _arg1 *C.LsmMathmlElementStyle // out
	var _arg2 C.double                 // out
	var _arg3 C.double                 // out
	var _arg4 *C.LsmMathmlBbox         // out

	_arg0 = (*C.LsmMathmlView)(unsafe.Pointer(externglib.InternObject(view).Native()))
	_arg1 = (*C.LsmMathmlElementStyle)(gextras.StructNative(unsafe.Pointer(style)))
	_arg2 = C.double(x)
	_arg3 = C.double(y)
	_arg4 = (*C.LsmMathmlBbox)(gextras.StructNative(unsafe.Pointer(bbox)))

	C.lsm_mathml_view_show_background(_arg0, _arg1, _arg2, _arg3, _arg4)
	runtime.KeepAlive(view)
	runtime.KeepAlive(style)
	runtime.KeepAlive(x)
	runtime.KeepAlive(y)
	runtime.KeepAlive(bbox)
}

// The function takes the following parameters:
//
//    - x
//    - y
//    - bbox
//
func (view *MathMLView) ShowBbox(x, y float64, bbox *MathMLBbox) {
	var _arg0 *C.LsmMathmlView // out
	var _arg1 C.double         // out
	var _arg2 C.double         // out
	var _arg3 *C.LsmMathmlBbox // out

	_arg0 = (*C.LsmMathmlView)(unsafe.Pointer(externglib.InternObject(view).Native()))
	_arg1 = C.double(x)
	_arg2 = C.double(y)
	_arg3 = (*C.LsmMathmlBbox)(gextras.StructNative(unsafe.Pointer(bbox)))

	C.lsm_mathml_view_show_bbox(_arg0, _arg1, _arg2, _arg3)
	runtime.KeepAlive(view)
	runtime.KeepAlive(x)
	runtime.KeepAlive(y)
	runtime.KeepAlive(bbox)
}

// The function takes the following parameters:
//
//    - style
//    - x
//    - y
//    - width
//    - thickness
//
func (view *MathMLView) ShowFractionLine(style *MathMLElementStyle, x, y, width, thickness float64) {
	var _arg0 *C.LsmMathmlView         // out
	var _arg1 *C.LsmMathmlElementStyle // out
	var _arg2 C.double                 // out
	var _arg3 C.double                 // out
	var _arg4 C.double                 // out
	var _arg5 C.double                 // out

	_arg0 = (*C.LsmMathmlView)(unsafe.Pointer(externglib.InternObject(view).Native()))
	_arg1 = (*C.LsmMathmlElementStyle)(gextras.StructNative(unsafe.Pointer(style)))
	_arg2 = C.double(x)
	_arg3 = C.double(y)
	_arg4 = C.double(width)
	_arg5 = C.double(thickness)

	C.lsm_mathml_view_show_fraction_line(_arg0, _arg1, _arg2, _arg3, _arg4, _arg5)
	runtime.KeepAlive(view)
	runtime.KeepAlive(style)
	runtime.KeepAlive(x)
	runtime.KeepAlive(y)
	runtime.KeepAlive(width)
	runtime.KeepAlive(thickness)
}

// The function takes the following parameters:
//
//    - style
//    - x0
//    - y0
//    - x1
//    - y1
//    - line
//    - lineWidth
//
func (view *MathMLView) ShowLine(style *MathMLElementStyle, x0, y0, x1, y1 float64, line MathMLLine, lineWidth float64) {
	var _arg0 *C.LsmMathmlView         // out
	var _arg1 *C.LsmMathmlElementStyle // out
	var _arg2 C.double                 // out
	var _arg3 C.double                 // out
	var _arg4 C.double                 // out
	var _arg5 C.double                 // out
	var _arg6 C.LsmMathmlLine          // out
	var _arg7 C.double                 // out

	_arg0 = (*C.LsmMathmlView)(unsafe.Pointer(externglib.InternObject(view).Native()))
	_arg1 = (*C.LsmMathmlElementStyle)(gextras.StructNative(unsafe.Pointer(style)))
	_arg2 = C.double(x0)
	_arg3 = C.double(y0)
	_arg4 = C.double(x1)
	_arg5 = C.double(y1)
	_arg6 = C.LsmMathmlLine(line)
	_arg7 = C.double(lineWidth)

	C.lsm_mathml_view_show_line(_arg0, _arg1, _arg2, _arg3, _arg4, _arg5, _arg6, _arg7)
	runtime.KeepAlive(view)
	runtime.KeepAlive(style)
	runtime.KeepAlive(x0)
	runtime.KeepAlive(y0)
	runtime.KeepAlive(x1)
	runtime.KeepAlive(y1)
	runtime.KeepAlive(line)
	runtime.KeepAlive(lineWidth)
}

// The function takes the following parameters:
//
//    - style
//    - notation
//    - x
//    - y
//    - bbox
//    - xChildOffset
//
func (view *MathMLView) ShowNotation(style *MathMLElementStyle, notation MathMLNotation, x, y float64, bbox *MathMLBbox, xChildOffset float64) {
	var _arg0 *C.LsmMathmlView         // out
	var _arg1 *C.LsmMathmlElementStyle // out
	var _arg2 C.LsmMathmlNotation      // out
	var _arg3 C.double                 // out
	var _arg4 C.double                 // out
	var _arg5 *C.LsmMathmlBbox         // out
	var _arg6 C.double                 // out

	_arg0 = (*C.LsmMathmlView)(unsafe.Pointer(externglib.InternObject(view).Native()))
	_arg1 = (*C.LsmMathmlElementStyle)(gextras.StructNative(unsafe.Pointer(style)))
	_arg2 = C.LsmMathmlNotation(notation)
	_arg3 = C.double(x)
	_arg4 = C.double(y)
	_arg5 = (*C.LsmMathmlBbox)(gextras.StructNative(unsafe.Pointer(bbox)))
	_arg6 = C.double(xChildOffset)

	C.lsm_mathml_view_show_notation(_arg0, _arg1, _arg2, _arg3, _arg4, _arg5, _arg6)
	runtime.KeepAlive(view)
	runtime.KeepAlive(style)
	runtime.KeepAlive(notation)
	runtime.KeepAlive(x)
	runtime.KeepAlive(y)
	runtime.KeepAlive(bbox)
	runtime.KeepAlive(xChildOffset)
}

// The function takes the following parameters:
//
//    - style
//    - x
//    - y
//    - text
//    - large
//    - stretchBbox
//
func (view *MathMLView) ShowOperator(style *MathMLElementStyle, x, y float64, text string, large bool, stretchBbox *MathMLBbox) {
	var _arg0 *C.LsmMathmlView         // out
	var _arg1 *C.LsmMathmlElementStyle // out
	var _arg2 C.double                 // out
	var _arg3 C.double                 // out
	var _arg4 *C.char                  // out
	var _arg5 C.gboolean               // out
	var _arg6 *C.LsmMathmlBbox         // out

	_arg0 = (*C.LsmMathmlView)(unsafe.Pointer(externglib.InternObject(view).Native()))
	_arg1 = (*C.LsmMathmlElementStyle)(gextras.StructNative(unsafe.Pointer(style)))
	_arg2 = C.double(x)
	_arg3 = C.double(y)
	_arg4 = (*C.char)(unsafe.Pointer(C.CString(text)))
	defer C.free(unsafe.Pointer(_arg4))
	if large {
		_arg5 = C.TRUE
	}
	_arg6 = (*C.LsmMathmlBbox)(gextras.StructNative(unsafe.Pointer(stretchBbox)))

	C.lsm_mathml_view_show_operator(_arg0, _arg1, _arg2, _arg3, _arg4, _arg5, _arg6)
	runtime.KeepAlive(view)
	runtime.KeepAlive(style)
	runtime.KeepAlive(x)
	runtime.KeepAlive(y)
	runtime.KeepAlive(text)
	runtime.KeepAlive(large)
	runtime.KeepAlive(stretchBbox)
}

// The function takes the following parameters:
//
//    - style
//    - x
//    - y
//    - width
//    - stretchBbox
//
func (view *MathMLView) ShowRadical(style *MathMLElementStyle, x, y, width float64, stretchBbox *MathMLBbox) {
	var _arg0 *C.LsmMathmlView         // out
	var _arg1 *C.LsmMathmlElementStyle // out
	var _arg2 C.double                 // out
	var _arg3 C.double                 // out
	var _arg4 C.double                 // out
	var _arg5 *C.LsmMathmlBbox         // out

	_arg0 = (*C.LsmMathmlView)(unsafe.Pointer(externglib.InternObject(view).Native()))
	_arg1 = (*C.LsmMathmlElementStyle)(gextras.StructNative(unsafe.Pointer(style)))
	_arg2 = C.double(x)
	_arg3 = C.double(y)
	_arg4 = C.double(width)
	_arg5 = (*C.LsmMathmlBbox)(gextras.StructNative(unsafe.Pointer(stretchBbox)))

	C.lsm_mathml_view_show_radical(_arg0, _arg1, _arg2, _arg3, _arg4, _arg5)
	runtime.KeepAlive(view)
	runtime.KeepAlive(style)
	runtime.KeepAlive(x)
	runtime.KeepAlive(y)
	runtime.KeepAlive(width)
	runtime.KeepAlive(stretchBbox)
}

// The function takes the following parameters:
//
//    - style
//    - x
//    - y
//    - width
//    - height
//    - line
//    - lineWidth
//
func (view *MathMLView) ShowRectangle(style *MathMLElementStyle, x, y, width, height float64, line MathMLLine, lineWidth float64) {
	var _arg0 *C.LsmMathmlView         // out
	var _arg1 *C.LsmMathmlElementStyle // out
	var _arg2 C.double                 // out
	var _arg3 C.double                 // out
	var _arg4 C.double                 // out
	var _arg5 C.double                 // out
	var _arg6 C.LsmMathmlLine          // out
	var _arg7 C.double                 // out

	_arg0 = (*C.LsmMathmlView)(unsafe.Pointer(externglib.InternObject(view).Native()))
	_arg1 = (*C.LsmMathmlElementStyle)(gextras.StructNative(unsafe.Pointer(style)))
	_arg2 = C.double(x)
	_arg3 = C.double(y)
	_arg4 = C.double(width)
	_arg5 = C.double(height)
	_arg6 = C.LsmMathmlLine(line)
	_arg7 = C.double(lineWidth)

	C.lsm_mathml_view_show_rectangle(_arg0, _arg1, _arg2, _arg3, _arg4, _arg5, _arg6, _arg7)
	runtime.KeepAlive(view)
	runtime.KeepAlive(style)
	runtime.KeepAlive(x)
	runtime.KeepAlive(y)
	runtime.KeepAlive(width)
	runtime.KeepAlive(height)
	runtime.KeepAlive(line)
	runtime.KeepAlive(lineWidth)
}

// The function takes the following parameters:
//
//    - style
//    - x
//    - y
//    - text
//
func (view *MathMLView) ShowText(style *MathMLElementStyle, x, y float64, text string) {
	var _arg0 *C.LsmMathmlView         // out
	var _arg1 *C.LsmMathmlElementStyle // out
	var _arg2 C.double                 // out
	var _arg3 C.double                 // out
	var _arg4 *C.char                  // out

	_arg0 = (*C.LsmMathmlView)(unsafe.Pointer(externglib.InternObject(view).Native()))
	_arg1 = (*C.LsmMathmlElementStyle)(gextras.StructNative(unsafe.Pointer(style)))
	_arg2 = C.double(x)
	_arg3 = C.double(y)
	_arg4 = (*C.char)(unsafe.Pointer(C.CString(text)))
	defer C.free(unsafe.Pointer(_arg4))

	C.lsm_mathml_view_show_text(_arg0, _arg1, _arg2, _arg3, _arg4)
	runtime.KeepAlive(view)
	runtime.KeepAlive(style)
	runtime.KeepAlive(x)
	runtime.KeepAlive(y)
	runtime.KeepAlive(text)
}

// MathMLGlyph: instance of this type is always passed by reference.
type MathMLGlyph struct {
	*mathMLGlyph
}

// mathMLGlyph is the struct that's finalized.
type mathMLGlyph struct {
	native *C.LsmMathmlGlyph
}

func (m *MathMLGlyph) Font() MathMLFont {
	var v MathMLFont // out
	v = MathMLFont(m.native.font)
	return v
}

func (m *MathMLGlyph) UTF8() [4]byte {
	var v [4]byte // out
	v = *(*[4]byte)(unsafe.Pointer(&m.native.utf8))
	return v
}

// MathMLOperatorGlyph: instance of this type is always passed by reference.
type MathMLOperatorGlyph struct {
	*mathMLOperatorGlyph
}

// mathMLOperatorGlyph is the struct that's finalized.
type mathMLOperatorGlyph struct {
	native *C.LsmMathmlOperatorGlyph
}

func (m *MathMLOperatorGlyph) UTF8() string {
	var v string // out
	v = C.GoString((*C.gchar)(unsafe.Pointer(m.native.utf8)))
	return v
}

func (m *MathMLOperatorGlyph) Flags() MathMLGlyphFlags {
	var v MathMLGlyphFlags // out
	v = MathMLGlyphFlags(m.native.flags)
	return v
}

func (m *MathMLOperatorGlyph) StartGlyph() *MathMLGlyph {
	var v *MathMLGlyph // out
	v = (*MathMLGlyph)(gextras.NewStructNative(unsafe.Pointer((&m.native.start_glyph))))
	return v
}

func (m *MathMLOperatorGlyph) MiddleGlyph() *MathMLGlyph {
	var v *MathMLGlyph // out
	v = (*MathMLGlyph)(gextras.NewStructNative(unsafe.Pointer((&m.native.middle_glyph))))
	return v
}

func (m *MathMLOperatorGlyph) StopGlyph() *MathMLGlyph {
	var v *MathMLGlyph // out
	v = (*MathMLGlyph)(gextras.NewStructNative(unsafe.Pointer((&m.native.stop_glyph))))
	return v
}

func (m *MathMLOperatorGlyph) GlueGlyph() *MathMLGlyph {
	var v *MathMLGlyph // out
	v = (*MathMLGlyph)(gextras.NewStructNative(unsafe.Pointer((&m.native.glue_glyph))))
	return v
}

func (m *MathMLOperatorGlyph) NSizedGlyphs() uint {
	var v uint // out
	v = uint(m.native.n_sized_glyphs)
	return v
}

func (m *MathMLOperatorGlyph) SizedGlyphs() [5]MathMLGlyph {
	var v [5]MathMLGlyph // out
	{
		src := &m.native.sized_glyphs
		for i := 0; i < 5; i++ {
			v[i] = *(*MathMLGlyph)(gextras.NewStructNative(unsafe.Pointer((&src[i]))))
		}
	}
	return v
}