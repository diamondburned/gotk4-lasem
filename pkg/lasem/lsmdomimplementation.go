// Code generated by girgen. DO NOT EDIT.

package lasem

import (
	"runtime"
	"unsafe"

	externglib "github.com/diamondburned/gotk4/pkg/core/glib"
)

// #include <stdlib.h>
import "C"

func DOMImplementationCleanup() {
	C.lsm_dom_implementation_cleanup()
}

// The function takes the following parameters:
//
//    - namespaceUri: namespace uri.
//    - qualifiedName: qualified name.
//
// The function returns the following values:
//
//    - domDocument: new DomDocument.
//
func DOMImplementationCreateDocument(namespaceUri, qualifiedName string) DOMDocumenter {
	var _arg1 *C.char           // out
	var _arg2 *C.char           // out
	var _cret *C.LsmDomDocument // in

	_arg1 = (*C.char)(unsafe.Pointer(C.CString(namespaceUri)))
	defer C.free(unsafe.Pointer(_arg1))
	_arg2 = (*C.char)(unsafe.Pointer(C.CString(qualifiedName)))
	defer C.free(unsafe.Pointer(_arg2))

	_cret = C.lsm_dom_implementation_create_document(_arg1, _arg2)
	runtime.KeepAlive(namespaceUri)
	runtime.KeepAlive(qualifiedName)

	var _domDocument DOMDocumenter // out

	{
		objptr := unsafe.Pointer(_cret)
		if objptr == nil {
			panic("object of type lasem.DOMDocumenter is nil")
		}

		object := externglib.AssumeOwnership(objptr)
		casted := object.WalkCast(func(obj externglib.Objector) bool {
			_, ok := obj.(DOMDocumenter)
			return ok
		})
		rv, ok := casted.(DOMDocumenter)
		if !ok {
			panic("no marshaler for " + object.TypeFromInstance().String() + " matching lasem.DOMDocumenter")
		}
		_domDocument = rv
	}

	return _domDocument
}
