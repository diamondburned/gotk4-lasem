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

// glib.Type values for lsmattributes.go.
var GTypeAttributeManager = externglib.Type(C.lsm_attribute_manager_get_type())

func init() {
	externglib.RegisterGValueMarshalers([]externglib.TypeMarshaler{
		{T: GTypeAttributeManager, F: marshalAttributeManager},
	})
}

// Attribute: instance of this type is always passed by reference.
type Attribute struct {
	*attribute
}

// attribute is the struct that's finalized.
type attribute struct {
	native *C.LsmAttribute
}

func (a *Attribute) Value() string {
	var v string // out
	v = C.GoString((*C.gchar)(unsafe.Pointer(a.native.value)))
	return v
}

// The function returns the following values:
//
func (attribute *Attribute) IsDefined() bool {
	var _arg0 *C.LsmAttribute // out
	var _cret C.gboolean      // in

	_arg0 = (*C.LsmAttribute)(gextras.StructNative(unsafe.Pointer(attribute)))

	_cret = C.lsm_attribute_is_defined(_arg0)
	runtime.KeepAlive(attribute)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// AttributeInfos: instance of this type is always passed by reference.
type AttributeInfos struct {
	*attributeInfos
}

// attributeInfos is the struct that's finalized.
type attributeInfos struct {
	native *C.LsmAttributeInfos
}

// AttributeManager: instance of this type is always passed by reference.
type AttributeManager struct {
	*attributeManager
}

// attributeManager is the struct that's finalized.
type attributeManager struct {
	native *C.LsmAttributeManager
}

func marshalAttributeManager(p uintptr) (interface{}, error) {
	b := externglib.ValueFromNative(unsafe.Pointer(p)).Boxed()
	return &AttributeManager{&attributeManager{(*C.LsmAttributeManager)(b)}}, nil
}

// NewAttributeManager constructs a struct AttributeManager.
func NewAttributeManager(nAttributes uint, attributeInfos *AttributeInfos) *AttributeManager {
	var _arg1 C.uint                 // out
	var _arg2 *C.LsmAttributeInfos   // out
	var _cret *C.LsmAttributeManager // in

	_arg1 = C.uint(nAttributes)
	_arg2 = (*C.LsmAttributeInfos)(gextras.StructNative(unsafe.Pointer(attributeInfos)))

	_cret = C.lsm_attribute_manager_new(_arg1, _arg2)
	runtime.KeepAlive(nAttributes)
	runtime.KeepAlive(attributeInfos)

	var _attributeManager *AttributeManager // out

	_attributeManager = (*AttributeManager)(gextras.NewStructNative(unsafe.Pointer(_cret)))
	runtime.SetFinalizer(
		gextras.StructIntern(unsafe.Pointer(_attributeManager)),
		func(intern *struct{ C unsafe.Pointer }) {
			C.lsm_attribute_manager_unref((*C.LsmAttributeManager)(intern.C))
		},
	)

	return _attributeManager
}

// The function takes the following parameters:
//
//    - nAttributes
//    - attributeInfos
//
func (manager *AttributeManager) AddAttributes(nAttributes uint, attributeInfos *AttributeInfos) {
	var _arg0 *C.LsmAttributeManager // out
	var _arg1 C.uint                 // out
	var _arg2 *C.LsmAttributeInfos   // out

	_arg0 = (*C.LsmAttributeManager)(gextras.StructNative(unsafe.Pointer(manager)))
	_arg1 = C.uint(nAttributes)
	_arg2 = (*C.LsmAttributeInfos)(gextras.StructNative(unsafe.Pointer(attributeInfos)))

	C.lsm_attribute_manager_add_attributes(_arg0, _arg1, _arg2)
	runtime.KeepAlive(manager)
	runtime.KeepAlive(nAttributes)
	runtime.KeepAlive(attributeInfos)
}

// The function returns the following values:
//
func (origin *AttributeManager) Duplicate() *AttributeManager {
	var _arg0 *C.LsmAttributeManager // out
	var _cret *C.LsmAttributeManager // in

	_arg0 = (*C.LsmAttributeManager)(gextras.StructNative(unsafe.Pointer(origin)))

	_cret = C.lsm_attribute_manager_duplicate(_arg0)
	runtime.KeepAlive(origin)

	var _attributeManager *AttributeManager // out

	_attributeManager = (*AttributeManager)(gextras.NewStructNative(unsafe.Pointer(_cret)))
	runtime.SetFinalizer(
		gextras.StructIntern(unsafe.Pointer(_attributeManager)),
		func(intern *struct{ C unsafe.Pointer }) {
			C.lsm_attribute_manager_unref((*C.LsmAttributeManager)(intern.C))
		},
	)

	return _attributeManager
}
