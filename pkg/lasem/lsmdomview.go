// Code generated by girgen. DO NOT EDIT.

package lasem

import (
	"runtime"
	"unsafe"

	"github.com/diamondburned/gotk4/pkg/cairo"
	"github.com/diamondburned/gotk4/pkg/core/gbox"
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
// extern void _gotk4_lasem0_DomViewClass_measure(LsmDomView*, double*, double*, double*);
// extern void _gotk4_lasem0_DomViewClass_render(LsmDomView*);
// extern void _gotk4_lasem0_DomViewClass_set_debug(LsmDomView*, char*, gboolean);
import "C"

// glib.Type values for lsmdomview.go.
var GTypeDOMView = externglib.Type(C.lsm_dom_view_get_type())

func init() {
	externglib.RegisterGValueMarshalers([]externglib.TypeMarshaler{
		{T: GTypeDOMView, F: marshalDOMView},
	})
}

const DOM_VIEW_DEFAULT_RESOLUTION = 72.000000
const DOM_VIEW_DEFAULT_VIEWBOX_HEIGHT = 200.000000
const DOM_VIEW_DEFAULT_VIEWBOX_WIDTH = 320.000000

// DOMViewOverrider contains methods that are overridable.
type DOMViewOverrider interface {
	// The function takes the following parameters:
	//
	//    - width
	//    - height
	//    - baseline
	//
	Measure(width, height, baseline *float64)
	Render()
	// SetDebug: configure feature debug.
	//
	// The function takes the following parameters:
	//
	//    - feature: name of the feature to debug.
	//    - enable: wether to enable debugging.
	//
	SetDebug(feature string, enable bool)
}

type DOMView struct {
	_ [0]func() // equal guard
	*externglib.Object
}

var (
	_ externglib.Objector = (*DOMView)(nil)
)

// DOMViewer describes types inherited from class DOMView.
//
// To get the original type, the caller must assert this to an interface or
// another type.
type DOMViewer interface {
	externglib.Objector
	baseDOMView() *DOMView
}

var _ DOMViewer = (*DOMView)(nil)

func classInitDOMViewer(gclassPtr, data C.gpointer) {
	C.g_type_class_add_private(gclassPtr, C.gsize(unsafe.Sizeof(uintptr(0))))

	goffset := C.g_type_class_get_instance_private_offset(gclassPtr)
	*(*C.gpointer)(unsafe.Add(unsafe.Pointer(gclassPtr), goffset)) = data

	goval := gbox.Get(uintptr(data))
	pclass := (*C.LsmDomViewClass)(unsafe.Pointer(gclassPtr))
	// gclass := (*C.GTypeClass)(unsafe.Pointer(gclassPtr))
	// pclass := (*C.LsmDomViewClass)(unsafe.Pointer(C.g_type_class_peek_parent(gclass)))

	if _, ok := goval.(interface {
		Measure(width, height, baseline *float64)
	}); ok {
		pclass.measure = (*[0]byte)(C._gotk4_lasem0_DomViewClass_measure)
	}

	if _, ok := goval.(interface{ Render() }); ok {
		pclass.render = (*[0]byte)(C._gotk4_lasem0_DomViewClass_render)
	}

	if _, ok := goval.(interface {
		SetDebug(feature string, enable bool)
	}); ok {
		pclass.set_debug = (*[0]byte)(C._gotk4_lasem0_DomViewClass_set_debug)
	}
}

//export _gotk4_lasem0_DomViewClass_measure
func _gotk4_lasem0_DomViewClass_measure(arg0 *C.LsmDomView, arg1 *C.double, arg2 *C.double, arg3 *C.double) {
	goval := externglib.GoPrivateFromObject(unsafe.Pointer(arg0))
	iface := goval.(interface {
		Measure(width, height, baseline *float64)
	})

	var _width *float64    // out
	var _height *float64   // out
	var _baseline *float64 // out

	_width = (*float64)(unsafe.Pointer(arg1))
	_height = (*float64)(unsafe.Pointer(arg2))
	_baseline = (*float64)(unsafe.Pointer(arg3))

	iface.Measure(_width, _height, _baseline)
}

//export _gotk4_lasem0_DomViewClass_render
func _gotk4_lasem0_DomViewClass_render(arg0 *C.LsmDomView) {
	goval := externglib.GoPrivateFromObject(unsafe.Pointer(arg0))
	iface := goval.(interface{ Render() })

	iface.Render()
}

//export _gotk4_lasem0_DomViewClass_set_debug
func _gotk4_lasem0_DomViewClass_set_debug(arg0 *C.LsmDomView, arg1 *C.char, arg2 C.gboolean) {
	goval := externglib.GoPrivateFromObject(unsafe.Pointer(arg0))
	iface := goval.(interface {
		SetDebug(feature string, enable bool)
	})

	var _feature string // out
	var _enable bool    // out

	_feature = C.GoString((*C.gchar)(unsafe.Pointer(arg1)))
	if arg2 != 0 {
		_enable = true
	}

	iface.SetDebug(_feature, _enable)
}

func wrapDOMView(obj *externglib.Object) *DOMView {
	return &DOMView{
		Object: obj,
	}
}

func marshalDOMView(p uintptr) (interface{}, error) {
	return wrapDOMView(externglib.ValueFromNative(unsafe.Pointer(p)).Object()), nil
}

func (self *DOMView) baseDOMView() *DOMView {
	return self
}

// BaseDOMView returns the underlying base object.
func BaseDOMView(obj DOMViewer) *DOMView {
	return obj.baseDOMView()
}

// The function returns the following values:
//
//    - gdouble: view resolution, in pixel per inch.
//
func (self *DOMView) Resolution() float64 {
	var _arg0 *C.LsmDomView // out
	var _cret C.double      // in

	_arg0 = (*C.LsmDomView)(unsafe.Pointer(externglib.InternObject(self).Native()))

	_cret = C.lsm_dom_view_get_resolution(_arg0)
	runtime.KeepAlive(self)

	var _gdouble float64 // out

	_gdouble = float64(_cret)

	return _gdouble
}

// Size: get the view size and baseline. Baseline is for use of view inside bloc
// of text.
//
// The function returns the following values:
//
//    - width (optional): view width placeholder, in points.
//    - height (optional): view height placeholder, in points.
//    - baseline (optional): view baseline, in points.
//
func (view *DOMView) Size() (width float64, height float64, baseline float64) {
	var _arg0 *C.LsmDomView // out
	var _arg1 C.double      // in
	var _arg2 C.double      // in
	var _arg3 C.double      // in

	_arg0 = (*C.LsmDomView)(unsafe.Pointer(externglib.InternObject(view).Native()))

	C.lsm_dom_view_get_size(_arg0, &_arg1, &_arg2, &_arg3)
	runtime.KeepAlive(view)

	var _width float64    // out
	var _height float64   // out
	var _baseline float64 // out

	_width = float64(_arg1)
	_height = float64(_arg2)
	_baseline = float64(_arg3)

	return _width, _height, _baseline
}

// SizePixels: get the view size and baseline. Baseline is for use of view
// inside bloc of text.
//
// The function returns the following values:
//
//    - width (optional): view width placeholder, in pixels.
//    - height (optional): view height placeholder, in pixels.
//    - baseline (optional): view baseline, in pixels.
//
func (view *DOMView) SizePixels() (width uint, height uint, baseline uint) {
	var _arg0 *C.LsmDomView // out
	var _arg1 C.uint        // in
	var _arg2 C.uint        // in
	var _arg3 C.uint        // in

	_arg0 = (*C.LsmDomView)(unsafe.Pointer(externglib.InternObject(view).Native()))

	C.lsm_dom_view_get_size_pixels(_arg0, &_arg1, &_arg2, &_arg3)
	runtime.KeepAlive(view)

	var _width uint    // out
	var _height uint   // out
	var _baseline uint // out

	_width = uint(_arg1)
	_height = uint(_arg2)
	_baseline = uint(_arg3)

	return _width, _height, _baseline
}

// The function returns the following values:
//
//    - box: viewport size, in points.
//
func (self *DOMView) Viewport() *Box {
	var _arg0 *C.LsmDomView // out
	var _cret C.LsmBox      // in

	_arg0 = (*C.LsmDomView)(unsafe.Pointer(externglib.InternObject(self).Native()))

	_cret = C.lsm_dom_view_get_viewport(_arg0)
	runtime.KeepAlive(self)

	var _box *Box // out

	_box = (*Box)(gextras.NewStructNative(unsafe.Pointer((&_cret))))
	runtime.SetFinalizer(
		gextras.StructIntern(unsafe.Pointer(_box)),
		func(intern *struct{ C unsafe.Pointer }) {
			C.free(intern.C)
		},
	)

	return _box
}

// The function returns the following values:
//
//    - box: viewport size, in pixels.
//
func (self *DOMView) ViewportPixels() *Box {
	var _arg0 *C.LsmDomView // out
	var _cret C.LsmBox      // in

	_arg0 = (*C.LsmDomView)(unsafe.Pointer(externglib.InternObject(self).Native()))

	_cret = C.lsm_dom_view_get_viewport_pixels(_arg0)
	runtime.KeepAlive(self)

	var _box *Box // out

	_box = (*Box)(gextras.NewStructNative(unsafe.Pointer((&_cret))))
	runtime.SetFinalizer(
		gextras.StructIntern(unsafe.Pointer(_box)),
		func(intern *struct{ C unsafe.Pointer }) {
			C.free(intern.C)
		},
	)

	return _box
}

// Render view in the cairo context.
//
// The function takes the following parameters:
//
//    - cairo context.
//    - x posiiton for rendering.
//    - y position for rendering.
//
func (view *DOMView) Render(cairo *cairo.Context, x, y float64) {
	var _arg0 *C.LsmDomView // out
	var _arg1 *C.cairo_t    // out
	var _arg2 C.double      // out
	var _arg3 C.double      // out

	_arg0 = (*C.LsmDomView)(unsafe.Pointer(externglib.InternObject(view).Native()))
	_arg1 = (*C.cairo_t)(unsafe.Pointer(cairo.Native()))
	_arg2 = C.double(x)
	_arg3 = C.double(y)

	C.lsm_dom_view_render(_arg0, _arg1, _arg2, _arg3)
	runtime.KeepAlive(view)
	runtime.KeepAlive(cairo)
	runtime.KeepAlive(x)
	runtime.KeepAlive(y)
}

// SetDebug: configure feature debug.
//
// The function takes the following parameters:
//
//    - feature: name of the feature to debug.
//    - enable: wether to enable debugging.
//
func (view *DOMView) SetDebug(feature string, enable bool) {
	var _arg0 *C.LsmDomView // out
	var _arg1 *C.char       // out
	var _arg2 C.gboolean    // out

	_arg0 = (*C.LsmDomView)(unsafe.Pointer(externglib.InternObject(view).Native()))
	_arg1 = (*C.char)(unsafe.Pointer(C.CString(feature)))
	defer C.free(unsafe.Pointer(_arg1))
	if enable {
		_arg2 = C.TRUE
	}

	C.lsm_dom_view_set_debug(_arg0, _arg1, _arg2)
	runtime.KeepAlive(view)
	runtime.KeepAlive(feature)
	runtime.KeepAlive(enable)
}

// SetDocument: change the document attached to view. The previously attached
// document is unreferenced.
//
// The function takes the following parameters:
//
//    - document: DomDocument.
//
func (view *DOMView) SetDocument(document DOMDocumenter) {
	var _arg0 *C.LsmDomView     // out
	var _arg1 *C.LsmDomDocument // out

	_arg0 = (*C.LsmDomView)(unsafe.Pointer(externglib.InternObject(view).Native()))
	_arg1 = (*C.LsmDomDocument)(unsafe.Pointer(externglib.InternObject(document).Native()))
	C.g_object_ref(C.gpointer(externglib.InternObject(document).Native()))

	C.lsm_dom_view_set_document(_arg0, _arg1)
	runtime.KeepAlive(view)
	runtime.KeepAlive(document)
}

// SetResolution: set the view resolution, in pixel per inch.
//
// The function takes the following parameters:
//
//    - ppi: resolution, in pixel per inch.
//
func (self *DOMView) SetResolution(ppi float64) {
	var _arg0 *C.LsmDomView // out
	var _arg1 C.double      // out

	_arg0 = (*C.LsmDomView)(unsafe.Pointer(externglib.InternObject(self).Native()))
	_arg1 = C.double(ppi)

	C.lsm_dom_view_set_resolution(_arg0, _arg1)
	runtime.KeepAlive(self)
	runtime.KeepAlive(ppi)
}

// SetViewport: set the viewport size.
//
// The function takes the following parameters:
//
//    - viewportPt: viewport size, in points.
//
func (self *DOMView) SetViewport(viewportPt *Box) {
	var _arg0 *C.LsmDomView // out
	var _arg1 *C.LsmBox     // out

	_arg0 = (*C.LsmDomView)(unsafe.Pointer(externglib.InternObject(self).Native()))
	_arg1 = (*C.LsmBox)(gextras.StructNative(unsafe.Pointer(viewportPt)))

	C.lsm_dom_view_set_viewport(_arg0, _arg1)
	runtime.KeepAlive(self)
	runtime.KeepAlive(viewportPt)
}

// SetViewportPixels: set the viewport size.
//
// The function takes the following parameters:
//
//    - viewport size, in pixels.
//
func (self *DOMView) SetViewportPixels(viewport *Box) {
	var _arg0 *C.LsmDomView // out
	var _arg1 *C.LsmBox     // out

	_arg0 = (*C.LsmDomView)(unsafe.Pointer(externglib.InternObject(self).Native()))
	_arg1 = (*C.LsmBox)(gextras.StructNative(unsafe.Pointer(viewport)))

	C.lsm_dom_view_set_viewport_pixels(_arg0, _arg1)
	runtime.KeepAlive(self)
	runtime.KeepAlive(viewport)
}
