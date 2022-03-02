// Code generated by girgen. DO NOT EDIT.

package lasem

import (
	"runtime"
	"unsafe"

	"github.com/diamondburned/gotk4/pkg/core/gerror"
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

// glib.Type values for lsmmathmldocument.go.
var GTypeMathMLDocument = externglib.Type(C.lsm_mathml_document_get_type())

func init() {
	externglib.RegisterGValueMarshalers([]externglib.TypeMarshaler{
		{T: GTypeMathMLDocument, F: marshalMathMLDocument},
	})
}

// MathMLDocumentOverrider contains methods that are overridable.
type MathMLDocumentOverrider interface {
}

type MathMLDocument struct {
	_ [0]func() // equal guard
	*externglib.Object
}

var (
	_ externglib.Objector = (*MathMLDocument)(nil)
)

func classInitMathMLDocumenter(gclassPtr, data C.gpointer) {
	C.g_type_class_add_private(gclassPtr, C.gsize(unsafe.Sizeof(uintptr(0))))

	goffset := C.g_type_class_get_instance_private_offset(gclassPtr)
	*(*C.gpointer)(unsafe.Add(unsafe.Pointer(gclassPtr), goffset)) = data

}

func wrapMathMLDocument(obj *externglib.Object) *MathMLDocument {
	return &MathMLDocument{
		Object: obj,
	}
}

func marshalMathMLDocument(p uintptr) (interface{}, error) {
	return wrapMathMLDocument(externglib.ValueFromNative(unsafe.Pointer(p)).Object()), nil
}

// The function takes the following parameters:
//
//    - itex
//    - size
//
// The function returns the following values:
//
func NewMathMLDocumentFromItex(itex string, size int) (*MathMLDocument, error) {
	var _arg1 *C.char              // out
	var _arg2 C.gssize             // out
	var _cret *C.LsmMathmlDocument // in
	var _cerr *C.GError            // in

	_arg1 = (*C.char)(unsafe.Pointer(C.CString(itex)))
	defer C.free(unsafe.Pointer(_arg1))
	_arg2 = C.gssize(size)

	_cret = C.lsm_mathml_document_new_from_itex(_arg1, _arg2, &_cerr)
	runtime.KeepAlive(itex)
	runtime.KeepAlive(size)

	var _mathmlDocument *MathMLDocument // out
	var _goerr error                    // out

	_mathmlDocument = wrapMathMLDocument(externglib.AssumeOwnership(unsafe.Pointer(_cret)))
	if _cerr != nil {
		_goerr = gerror.Take(unsafe.Pointer(_cerr))
	}

	return _mathmlDocument, _goerr
}

// The function takes the following parameters:
//
// The function returns the following values:
//
func NewMathMLDocumentFromItexPath(url string) (*MathMLDocument, error) {
	var _arg1 *C.char              // out
	var _cret *C.LsmMathmlDocument // in
	var _cerr *C.GError            // in

	_arg1 = (*C.char)(unsafe.Pointer(C.CString(url)))
	defer C.free(unsafe.Pointer(_arg1))

	_cret = C.lsm_mathml_document_new_from_itex_path(_arg1, &_cerr)
	runtime.KeepAlive(url)

	var _mathmlDocument *MathMLDocument // out
	var _goerr error                    // out

	_mathmlDocument = wrapMathMLDocument(externglib.AssumeOwnership(unsafe.Pointer(_cret)))
	if _cerr != nil {
		_goerr = gerror.Take(unsafe.Pointer(_cerr))
	}

	return _mathmlDocument, _goerr
}

// The function takes the following parameters:
//
// The function returns the following values:
//
func NewMathMLDocumentFromItexURL(url string) (*MathMLDocument, error) {
	var _arg1 *C.char              // out
	var _cret *C.LsmMathmlDocument // in
	var _cerr *C.GError            // in

	_arg1 = (*C.char)(unsafe.Pointer(C.CString(url)))
	defer C.free(unsafe.Pointer(_arg1))

	_cret = C.lsm_mathml_document_new_from_itex_url(_arg1, &_cerr)
	runtime.KeepAlive(url)

	var _mathmlDocument *MathMLDocument // out
	var _goerr error                    // out

	_mathmlDocument = wrapMathMLDocument(externglib.AssumeOwnership(unsafe.Pointer(_cret)))
	if _cerr != nil {
		_goerr = gerror.Take(unsafe.Pointer(_cerr))
	}

	return _mathmlDocument, _goerr
}