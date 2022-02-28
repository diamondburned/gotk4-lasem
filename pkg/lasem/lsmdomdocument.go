// Code generated by girgen. DO NOT EDIT.

package lasem

import (
	"runtime"
	"unsafe"

	"github.com/diamondburned/gotk4/pkg/core/gbox"
	externglib "github.com/diamondburned/gotk4/pkg/core/glib"
)

// #include <stdlib.h>
// #include <glib-object.h>
// extern LsmDomElement* _gotk4_lasem0_DomDocumentClass_create_element(LsmDomDocument*, char*);
// extern LsmDomElement* _gotk4_lasem0_DomDocumentClass_get_document_element(LsmDomDocument*);
// extern LsmDomText* _gotk4_lasem0_DomDocumentClass_create_text_node(LsmDomDocument*, char*);
// extern LsmDomView* _gotk4_lasem0_DomDocumentClass_create_view(LsmDomDocument*);
import "C"

// glib.Type values for lsmdomdocument.go.
var GTypeDOMDocument = externglib.Type(C.lsm_dom_document_get_type())

func init() {
	externglib.RegisterGValueMarshalers([]externglib.TypeMarshaler{
		{T: GTypeDOMDocument, F: marshalDOMDocument},
	})
}

// DOMDocumentOverrider contains methods that are overridable.
type DOMDocumentOverrider interface {
	// The function takes the following parameters:
	//
	//    - tagName: name of the element to create.
	//
	// The function returns the following values:
	//
	//    - domElement: newly created DomElement.
	//
	CreateElement(tagName string) DOMElementer
	// The function takes the following parameters:
	//
	//    - data: content of the text node.
	//
	// The function returns the following values:
	//
	//    - domText: newly created DomText.
	//
	CreateTextNode(data string) *DOMText
	// The function returns the following values:
	//
	//    - domView: new DomView.
	//
	CreateView() DOMViewer
	// The function returns the following values:
	//
	//    - domElement: document element.
	//
	DocumentElement() DOMElementer
}

type DOMDocument struct {
	_ [0]func() // equal guard
	DOMNode
}

var (
	_ DOMNoder = (*DOMDocument)(nil)
)

// DOMDocumenter describes types inherited from class DOMDocument.
//
// To get the original type, the caller must assert this to an interface or
// another type.
type DOMDocumenter interface {
	externglib.Objector
	baseDOMDocument() *DOMDocument
}

var _ DOMDocumenter = (*DOMDocument)(nil)

func classInitDOMDocumenter(gclassPtr, data C.gpointer) {
	C.g_type_class_add_private(gclassPtr, C.gsize(unsafe.Sizeof(uintptr(0))))

	goffset := C.g_type_class_get_instance_private_offset(gclassPtr)
	*(*C.gpointer)(unsafe.Add(unsafe.Pointer(gclassPtr), goffset)) = data

	goval := gbox.Get(uintptr(data))
	pclass := (*C.LsmDomDocumentClass)(unsafe.Pointer(gclassPtr))
	// gclass := (*C.GTypeClass)(unsafe.Pointer(gclassPtr))
	// pclass := (*C.LsmDomDocumentClass)(unsafe.Pointer(C.g_type_class_peek_parent(gclass)))

	if _, ok := goval.(interface {
		CreateElement(tagName string) DOMElementer
	}); ok {
		pclass.create_element = (*[0]byte)(C._gotk4_lasem0_DomDocumentClass_create_element)
	}

	if _, ok := goval.(interface{ CreateTextNode(data string) *DOMText }); ok {
		pclass.create_text_node = (*[0]byte)(C._gotk4_lasem0_DomDocumentClass_create_text_node)
	}

	if _, ok := goval.(interface{ CreateView() DOMViewer }); ok {
		pclass.create_view = (*[0]byte)(C._gotk4_lasem0_DomDocumentClass_create_view)
	}

	if _, ok := goval.(interface{ DocumentElement() DOMElementer }); ok {
		pclass.get_document_element = (*[0]byte)(C._gotk4_lasem0_DomDocumentClass_get_document_element)
	}
}

//export _gotk4_lasem0_DomDocumentClass_create_element
func _gotk4_lasem0_DomDocumentClass_create_element(arg0 *C.LsmDomDocument, arg1 *C.char) (cret *C.LsmDomElement) {
	goval := externglib.GoPrivateFromObject(unsafe.Pointer(arg0))
	iface := goval.(interface {
		CreateElement(tagName string) DOMElementer
	})

	var _tagName string // out

	_tagName = C.GoString((*C.gchar)(unsafe.Pointer(arg1)))

	domElement := iface.CreateElement(_tagName)

	cret = (*C.LsmDomElement)(unsafe.Pointer(externglib.InternObject(domElement).Native()))
	C.g_object_ref(C.gpointer(externglib.InternObject(domElement).Native()))

	return cret
}

//export _gotk4_lasem0_DomDocumentClass_create_text_node
func _gotk4_lasem0_DomDocumentClass_create_text_node(arg0 *C.LsmDomDocument, arg1 *C.char) (cret *C.LsmDomText) {
	goval := externglib.GoPrivateFromObject(unsafe.Pointer(arg0))
	iface := goval.(interface{ CreateTextNode(data string) *DOMText })

	var _data string // out

	_data = C.GoString((*C.gchar)(unsafe.Pointer(arg1)))

	domText := iface.CreateTextNode(_data)

	cret = (*C.LsmDomText)(unsafe.Pointer(externglib.InternObject(domText).Native()))
	C.g_object_ref(C.gpointer(externglib.InternObject(domText).Native()))

	return cret
}

//export _gotk4_lasem0_DomDocumentClass_create_view
func _gotk4_lasem0_DomDocumentClass_create_view(arg0 *C.LsmDomDocument) (cret *C.LsmDomView) {
	goval := externglib.GoPrivateFromObject(unsafe.Pointer(arg0))
	iface := goval.(interface{ CreateView() DOMViewer })

	domView := iface.CreateView()

	cret = (*C.LsmDomView)(unsafe.Pointer(externglib.InternObject(domView).Native()))
	C.g_object_ref(C.gpointer(externglib.InternObject(domView).Native()))

	return cret
}

//export _gotk4_lasem0_DomDocumentClass_get_document_element
func _gotk4_lasem0_DomDocumentClass_get_document_element(arg0 *C.LsmDomDocument) (cret *C.LsmDomElement) {
	goval := externglib.GoPrivateFromObject(unsafe.Pointer(arg0))
	iface := goval.(interface{ DocumentElement() DOMElementer })

	domElement := iface.DocumentElement()

	cret = (*C.LsmDomElement)(unsafe.Pointer(externglib.InternObject(domElement).Native()))

	return cret
}

func wrapDOMDocument(obj *externglib.Object) *DOMDocument {
	return &DOMDocument{
		DOMNode: DOMNode{
			Object: obj,
		},
	}
}

func marshalDOMDocument(p uintptr) (interface{}, error) {
	return wrapDOMDocument(externglib.ValueFromNative(unsafe.Pointer(p)).Object()), nil
}

func (self *DOMDocument) baseDOMDocument() *DOMDocument {
	return self
}

// BaseDOMDocument returns the underlying base object.
func BaseDOMDocument(obj DOMDocumenter) *DOMDocument {
	return obj.baseDOMDocument()
}

// The function takes the following parameters:
//
//    - tagName: name of the element to create.
//
// The function returns the following values:
//
//    - domElement: newly created DomElement.
//
func (self *DOMDocument) CreateElement(tagName string) DOMElementer {
	var _arg0 *C.LsmDomDocument // out
	var _arg1 *C.char           // out
	var _cret *C.LsmDomElement  // in

	_arg0 = (*C.LsmDomDocument)(unsafe.Pointer(externglib.InternObject(self).Native()))
	_arg1 = (*C.char)(unsafe.Pointer(C.CString(tagName)))
	defer C.free(unsafe.Pointer(_arg1))

	_cret = C.lsm_dom_document_create_element(_arg0, _arg1)
	runtime.KeepAlive(self)
	runtime.KeepAlive(tagName)

	var _domElement DOMElementer // out

	{
		objptr := unsafe.Pointer(_cret)
		if objptr == nil {
			panic("object of type lasem.DOMElementer is nil")
		}

		object := externglib.AssumeOwnership(objptr)
		casted := object.WalkCast(func(obj externglib.Objector) bool {
			_, ok := obj.(DOMElementer)
			return ok
		})
		rv, ok := casted.(DOMElementer)
		if !ok {
			panic("no marshaler for " + object.TypeFromInstance().String() + " matching lasem.DOMElementer")
		}
		_domElement = rv
	}

	return _domElement
}

// The function takes the following parameters:
//
//    - data: content of the text node.
//
// The function returns the following values:
//
//    - domText: newly created DomText.
//
func (self *DOMDocument) CreateTextNode(data string) *DOMText {
	var _arg0 *C.LsmDomDocument // out
	var _arg1 *C.char           // out
	var _cret *C.LsmDomText     // in

	_arg0 = (*C.LsmDomDocument)(unsafe.Pointer(externglib.InternObject(self).Native()))
	_arg1 = (*C.char)(unsafe.Pointer(C.CString(data)))
	defer C.free(unsafe.Pointer(_arg1))

	_cret = C.lsm_dom_document_create_text_node(_arg0, _arg1)
	runtime.KeepAlive(self)
	runtime.KeepAlive(data)

	var _domText *DOMText // out

	_domText = wrapDOMText(externglib.AssumeOwnership(unsafe.Pointer(_cret)))

	return _domText
}

// The function returns the following values:
//
//    - domView: new DomView.
//
func (self *DOMDocument) CreateView() DOMViewer {
	var _arg0 *C.LsmDomDocument // out
	var _cret *C.LsmDomView     // in

	_arg0 = (*C.LsmDomDocument)(unsafe.Pointer(externglib.InternObject(self).Native()))

	_cret = C.lsm_dom_document_create_view(_arg0)
	runtime.KeepAlive(self)

	var _domView DOMViewer // out

	{
		objptr := unsafe.Pointer(_cret)
		if objptr == nil {
			panic("object of type lasem.DOMViewer is nil")
		}

		object := externglib.AssumeOwnership(objptr)
		casted := object.WalkCast(func(obj externglib.Objector) bool {
			_, ok := obj.(DOMViewer)
			return ok
		})
		rv, ok := casted.(DOMViewer)
		if !ok {
			panic("no marshaler for " + object.TypeFromInstance().String() + " matching lasem.DOMViewer")
		}
		_domView = rv
	}

	return _domView
}

// The function returns the following values:
//
//    - domElement: document element.
//
func (self *DOMDocument) DocumentElement() DOMElementer {
	var _arg0 *C.LsmDomDocument // out
	var _cret *C.LsmDomElement  // in

	_arg0 = (*C.LsmDomDocument)(unsafe.Pointer(externglib.InternObject(self).Native()))

	_cret = C.lsm_dom_document_get_document_element(_arg0)
	runtime.KeepAlive(self)

	var _domElement DOMElementer // out

	{
		objptr := unsafe.Pointer(_cret)
		if objptr == nil {
			panic("object of type lasem.DOMElementer is nil")
		}

		object := externglib.Take(objptr)
		casted := object.WalkCast(func(obj externglib.Objector) bool {
			_, ok := obj.(DOMElementer)
			return ok
		})
		rv, ok := casted.(DOMElementer)
		if !ok {
			panic("no marshaler for " + object.TypeFromInstance().String() + " matching lasem.DOMElementer")
		}
		_domElement = rv
	}

	return _domElement
}

// The function takes the following parameters:
//
//    - id of the element to find.
//
// The function returns the following values:
//
//    - domElement: requested element, NULL if not found.
//
func (self *DOMDocument) ElementByID(id string) DOMElementer {
	var _arg0 *C.LsmDomDocument // out
	var _arg1 *C.char           // out
	var _cret *C.LsmDomElement  // in

	_arg0 = (*C.LsmDomDocument)(unsafe.Pointer(externglib.InternObject(self).Native()))
	_arg1 = (*C.char)(unsafe.Pointer(C.CString(id)))
	defer C.free(unsafe.Pointer(_arg1))

	_cret = C.lsm_dom_document_get_element_by_id(_arg0, _arg1)
	runtime.KeepAlive(self)
	runtime.KeepAlive(id)

	var _domElement DOMElementer // out

	{
		objptr := unsafe.Pointer(_cret)
		if objptr == nil {
			panic("object of type lasem.DOMElementer is nil")
		}

		object := externglib.Take(objptr)
		casted := object.WalkCast(func(obj externglib.Objector) bool {
			_, ok := obj.(DOMElementer)
			return ok
		})
		rv, ok := casted.(DOMElementer)
		if !ok {
			panic("no marshaler for " + object.TypeFromInstance().String() + " matching lasem.DOMElementer")
		}
		_domElement = rv
	}

	return _domElement
}

// The function returns the following values:
//
func (self *DOMDocument) URL() string {
	var _arg0 *C.LsmDomDocument // out
	var _cret *C.char           // in

	_arg0 = (*C.LsmDomDocument)(unsafe.Pointer(externglib.InternObject(self).Native()))

	_cret = C.lsm_dom_document_get_url(_arg0)
	runtime.KeepAlive(self)

	var _utf8 string // out

	_utf8 = C.GoString((*C.gchar)(unsafe.Pointer(_cret)))

	return _utf8
}

// The function takes the following parameters:
//
//    - element
//    - id
//
func (self *DOMDocument) RegisterElement(element DOMElementer, id string) {
	var _arg0 *C.LsmDomDocument // out
	var _arg1 *C.LsmDomElement  // out
	var _arg2 *C.char           // out

	_arg0 = (*C.LsmDomDocument)(unsafe.Pointer(externglib.InternObject(self).Native()))
	_arg1 = (*C.LsmDomElement)(unsafe.Pointer(externglib.InternObject(element).Native()))
	_arg2 = (*C.char)(unsafe.Pointer(C.CString(id)))
	defer C.free(unsafe.Pointer(_arg2))

	C.lsm_dom_document_register_element(_arg0, _arg1, _arg2)
	runtime.KeepAlive(self)
	runtime.KeepAlive(element)
	runtime.KeepAlive(id)
}

// The function takes the following parameters:
//
func (self *DOMDocument) SetPath(path string) {
	var _arg0 *C.LsmDomDocument // out
	var _arg1 *C.char           // out

	_arg0 = (*C.LsmDomDocument)(unsafe.Pointer(externglib.InternObject(self).Native()))
	_arg1 = (*C.char)(unsafe.Pointer(C.CString(path)))
	defer C.free(unsafe.Pointer(_arg1))

	C.lsm_dom_document_set_path(_arg0, _arg1)
	runtime.KeepAlive(self)
	runtime.KeepAlive(path)
}

// The function takes the following parameters:
//
func (self *DOMDocument) SetURL(url string) {
	var _arg0 *C.LsmDomDocument // out
	var _arg1 *C.char           // out

	_arg0 = (*C.LsmDomDocument)(unsafe.Pointer(externglib.InternObject(self).Native()))
	_arg1 = (*C.char)(unsafe.Pointer(C.CString(url)))
	defer C.free(unsafe.Pointer(_arg1))

	C.lsm_dom_document_set_url(_arg0, _arg1)
	runtime.KeepAlive(self)
	runtime.KeepAlive(url)
}
