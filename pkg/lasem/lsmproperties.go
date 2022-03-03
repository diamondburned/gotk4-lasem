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

// glib.Type values for lsmproperties.go.
var GTypePropertyManager = externglib.Type(C.lsm_property_manager_get_type())

func init() {
	externglib.RegisterGValueMarshalers([]externglib.TypeMarshaler{
		{T: GTypePropertyManager, F: marshalPropertyManager},
	})
}

// Property: instance of this type is always passed by reference.
type Property struct {
	*property
}

// property is the struct that's finalized.
type property struct {
	native *C.LsmProperty
}

func (p *Property) ID() uint16 {
	var v uint16 // out
	v = uint16(p.native.id)
	return v
}

func (p *Property) Flags() uint16 {
	var v uint16 // out
	v = uint16(p.native.flags)
	return v
}

func (p *Property) Value() string {
	var v string // out
	v = C.GoString((*C.gchar)(unsafe.Pointer(p.native.value)))
	return v
}

// PropertyBag: instance of this type is always passed by reference.
type PropertyBag struct {
	*propertyBag
}

// propertyBag is the struct that's finalized.
type propertyBag struct {
	native *C.LsmPropertyBag
}

func (bag *PropertyBag) Init() {
	var _arg0 *C.LsmPropertyBag // out

	_arg0 = (*C.LsmPropertyBag)(gextras.StructNative(unsafe.Pointer(bag)))

	C.lsm_property_bag_init(_arg0)
	runtime.KeepAlive(bag)
}

// PropertyInfos: instance of this type is always passed by reference.
type PropertyInfos struct {
	*propertyInfos
}

// propertyInfos is the struct that's finalized.
type propertyInfos struct {
	native *C.LsmPropertyInfos
}

// PropertyManager: instance of this type is always passed by reference.
type PropertyManager struct {
	*propertyManager
}

// propertyManager is the struct that's finalized.
type propertyManager struct {
	native *C.LsmPropertyManager
}

func marshalPropertyManager(p uintptr) (interface{}, error) {
	b := externglib.ValueFromNative(unsafe.Pointer(p)).Boxed()
	return &PropertyManager{&propertyManager{(*C.LsmPropertyManager)(b)}}, nil
}

// NewPropertyManager constructs a struct PropertyManager.
func NewPropertyManager(nProperties uint, propertyInfos *PropertyInfos) *PropertyManager {
	var _arg1 C.uint                // out
	var _arg2 *C.LsmPropertyInfos   // out
	var _cret *C.LsmPropertyManager // in

	_arg1 = C.uint(nProperties)
	_arg2 = (*C.LsmPropertyInfos)(gextras.StructNative(unsafe.Pointer(propertyInfos)))

	_cret = C.lsm_property_manager_new(_arg1, _arg2)
	runtime.KeepAlive(nProperties)
	runtime.KeepAlive(propertyInfos)

	var _propertyManager *PropertyManager // out

	_propertyManager = (*PropertyManager)(gextras.NewStructNative(unsafe.Pointer(_cret)))
	runtime.SetFinalizer(
		gextras.StructIntern(unsafe.Pointer(_propertyManager)),
		func(intern *struct{ C unsafe.Pointer }) {
			C.lsm_property_manager_unref((*C.LsmPropertyManager)(intern.C))
		},
	)

	return _propertyManager
}

// The function takes the following parameters:
//
func (manager *PropertyManager) CleanProperties(propertyBag *PropertyBag) {
	var _arg0 *C.LsmPropertyManager // out
	var _arg1 *C.LsmPropertyBag     // out

	_arg0 = (*C.LsmPropertyManager)(gextras.StructNative(unsafe.Pointer(manager)))
	_arg1 = (*C.LsmPropertyBag)(gextras.StructNative(unsafe.Pointer(propertyBag)))

	C.lsm_property_manager_clean_properties(_arg0, _arg1)
	runtime.KeepAlive(manager)
	runtime.KeepAlive(propertyBag)
}

// The function takes the following parameters:
//
//    - propertyBag
//    - name
//
// The function returns the following values:
//
func (manager *PropertyManager) Property(propertyBag *PropertyBag, name string) string {
	var _arg0 *C.LsmPropertyManager // out
	var _arg1 *C.LsmPropertyBag     // out
	var _arg2 *C.char               // out
	var _cret *C.char               // in

	_arg0 = (*C.LsmPropertyManager)(gextras.StructNative(unsafe.Pointer(manager)))
	_arg1 = (*C.LsmPropertyBag)(gextras.StructNative(unsafe.Pointer(propertyBag)))
	_arg2 = (*C.char)(unsafe.Pointer(C.CString(name)))
	defer C.free(unsafe.Pointer(_arg2))

	_cret = C.lsm_property_manager_get_property(_arg0, _arg1, _arg2)
	runtime.KeepAlive(manager)
	runtime.KeepAlive(propertyBag)
	runtime.KeepAlive(name)

	var _utf8 string // out

	_utf8 = C.GoString((*C.gchar)(unsafe.Pointer(_cret)))

	return _utf8
}

// The function takes the following parameters:
//
// The function returns the following values:
//
func (propertyManager *PropertyManager) Serialize(propertyBag *PropertyBag) string {
	var _arg0 *C.LsmPropertyManager // out
	var _arg1 *C.LsmPropertyBag     // out
	var _cret *C.char               // in

	_arg0 = (*C.LsmPropertyManager)(gextras.StructNative(unsafe.Pointer(propertyManager)))
	_arg1 = (*C.LsmPropertyBag)(gextras.StructNative(unsafe.Pointer(propertyBag)))

	_cret = C.lsm_property_manager_serialize(_arg0, _arg1)
	runtime.KeepAlive(propertyManager)
	runtime.KeepAlive(propertyBag)

	var _utf8 string // out

	_utf8 = C.GoString((*C.gchar)(unsafe.Pointer(_cret)))
	defer C.free(unsafe.Pointer(_cret))

	return _utf8
}

// The function takes the following parameters:
//
//    - propertyBag
//    - name
//    - value
//
// The function returns the following values:
//
func (manager *PropertyManager) SetProperty(propertyBag *PropertyBag, name string, value string) bool {
	var _arg0 *C.LsmPropertyManager // out
	var _arg1 *C.LsmPropertyBag     // out
	var _arg2 *C.char               // out
	var _arg3 *C.char               // out
	var _cret C.gboolean            // in

	_arg0 = (*C.LsmPropertyManager)(gextras.StructNative(unsafe.Pointer(manager)))
	_arg1 = (*C.LsmPropertyBag)(gextras.StructNative(unsafe.Pointer(propertyBag)))
	_arg2 = (*C.char)(unsafe.Pointer(C.CString(name)))
	defer C.free(unsafe.Pointer(_arg2))
	_arg3 = (*C.char)(unsafe.Pointer(C.CString(value)))
	defer C.free(unsafe.Pointer(_arg3))

	_cret = C.lsm_property_manager_set_property(_arg0, _arg1, _arg2, _arg3)
	runtime.KeepAlive(manager)
	runtime.KeepAlive(propertyBag)
	runtime.KeepAlive(name)
	runtime.KeepAlive(value)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}
