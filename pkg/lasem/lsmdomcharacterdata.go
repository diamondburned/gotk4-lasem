// Code generated by girgen. DO NOT EDIT.

package lasem

import (
	"runtime"
	"unsafe"

	externglib "github.com/diamondburned/gotk4/pkg/core/glib"
)

// #include <stdlib.h>
// #include <glib-object.h>
// #include <lsm.h>
// #include <lsmdom.h>
// #include <lsmdomdocument.h>
// #include <lsmdomdocumentfragment.h>
// #include <lsmdomnamednodemap.h>
import "C"

// glib.Type values for lsmdomcharacterdata.go.
var GTypeDOMCharacterData = externglib.Type(C.lsm_dom_character_data_get_type())

func init() {
	externglib.RegisterGValueMarshalers([]externglib.TypeMarshaler{
		{T: GTypeDOMCharacterData, F: marshalDOMCharacterData},
	})
}

// DOMCharacterDataOverrider contains methods that are overridable.
type DOMCharacterDataOverrider interface {
}

type DOMCharacterData struct {
	_ [0]func() // equal guard
	DOMNode
}

var (
	_ DOMNoder = (*DOMCharacterData)(nil)
)

// DOMCharacterDatar describes types inherited from class DOMCharacterData.
//
// To get the original type, the caller must assert this to an interface or
// another type.
type DOMCharacterDatar interface {
	externglib.Objector
	baseDOMCharacterData() *DOMCharacterData
}

var _ DOMCharacterDatar = (*DOMCharacterData)(nil)

func classInitDOMCharacterDatar(gclassPtr, data C.gpointer) {
	C.g_type_class_add_private(gclassPtr, C.gsize(unsafe.Sizeof(uintptr(0))))

	goffset := C.g_type_class_get_instance_private_offset(gclassPtr)
	*(*C.gpointer)(unsafe.Add(unsafe.Pointer(gclassPtr), goffset)) = data

}

func wrapDOMCharacterData(obj *externglib.Object) *DOMCharacterData {
	return &DOMCharacterData{
		DOMNode: DOMNode{
			Object: obj,
		},
	}
}

func marshalDOMCharacterData(p uintptr) (interface{}, error) {
	return wrapDOMCharacterData(externglib.ValueFromNative(unsafe.Pointer(p)).Object()), nil
}

func (self *DOMCharacterData) baseDOMCharacterData() *DOMCharacterData {
	return self
}

// BaseDOMCharacterData returns the underlying base object.
func BaseDOMCharacterData(obj DOMCharacterDatar) *DOMCharacterData {
	return obj.baseDOMCharacterData()
}

// Data: get the character node data.
//
// The function returns the following values:
//
//    - utf8: character node data.
//
func (self *DOMCharacterData) Data() string {
	var _arg0 *C.LsmDomCharacterData // out
	var _cret *C.char                // in

	_arg0 = (*C.LsmDomCharacterData)(unsafe.Pointer(externglib.InternObject(self).Native()))

	_cret = C.lsm_dom_character_data_get_data(_arg0)
	runtime.KeepAlive(self)

	var _utf8 string // out

	_utf8 = C.GoString((*C.gchar)(unsafe.Pointer(_cret)))

	return _utf8
}

// SetData: set the character node data.
//
// The function takes the following parameters:
//
//    - value: new node data.
//
func (self *DOMCharacterData) SetData(value string) {
	var _arg0 *C.LsmDomCharacterData // out
	var _arg1 *C.char                // out

	_arg0 = (*C.LsmDomCharacterData)(unsafe.Pointer(externglib.InternObject(self).Native()))
	_arg1 = (*C.char)(unsafe.Pointer(C.CString(value)))
	defer C.free(unsafe.Pointer(_arg1))

	C.lsm_dom_character_data_set_data(_arg0, _arg1)
	runtime.KeepAlive(self)
	runtime.KeepAlive(value)
}
