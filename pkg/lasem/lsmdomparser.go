// Code generated by girgen. DO NOT EDIT.

package lasem

import (
	"runtime"
	"unsafe"

	"github.com/diamondburned/gotk4/pkg/core/gerror"
	externglib "github.com/diamondburned/gotk4/pkg/core/glib"
	"github.com/diamondburned/gotk4/pkg/gio/v2"
)

// #include <stdlib.h>
// #include <lsm.h>
// #include <lsmdom.h>
// #include <lsmdomdocument.h>
// #include <lsmdomdocumentfragment.h>
// #include <lsmdomnamednodemap.h>
import "C"

// The function takes the following parameters:
//
// The function returns the following values:
//
func NewDOMDocumentFromPath(path string) (*DOMDocument, error) {
	var _arg1 *C.char           // out
	var _cret *C.LsmDomDocument // in
	var _cerr *C.GError         // in

	_arg1 = (*C.char)(unsafe.Pointer(C.CString(path)))
	defer C.free(unsafe.Pointer(_arg1))

	_cret = C.lsm_dom_document_new_from_path(_arg1, &_cerr)
	runtime.KeepAlive(path)

	var _domDocument *DOMDocument // out
	var _goerr error              // out

	_domDocument = wrapDOMDocument(externglib.AssumeOwnership(unsafe.Pointer(_cret)))
	if _cerr != nil {
		_goerr = gerror.Take(unsafe.Pointer(_cerr))
	}

	return _domDocument, _goerr
}

// The function takes the following parameters:
//
// The function returns the following values:
//
func NewDOMDocumentFromURL(url string) (*DOMDocument, error) {
	var _arg1 *C.char           // out
	var _cret *C.LsmDomDocument // in
	var _cerr *C.GError         // in

	_arg1 = (*C.char)(unsafe.Pointer(C.CString(url)))
	defer C.free(unsafe.Pointer(_arg1))

	_cret = C.lsm_dom_document_new_from_url(_arg1, &_cerr)
	runtime.KeepAlive(url)

	var _domDocument *DOMDocument // out
	var _goerr error              // out

	_domDocument = wrapDOMDocument(externglib.AssumeOwnership(unsafe.Pointer(_cret)))
	if _cerr != nil {
		_goerr = gerror.Take(unsafe.Pointer(_cerr))
	}

	return _domDocument, _goerr
}

// The function takes the following parameters:
//
func (documennt *DOMDocument) SaveToPath(path string) error {
	var _arg0 *C.LsmDomDocument // out
	var _arg1 *C.char           // out
	var _cerr *C.GError         // in

	_arg0 = (*C.LsmDomDocument)(unsafe.Pointer(externglib.InternObject(documennt).Native()))
	_arg1 = (*C.char)(unsafe.Pointer(C.CString(path)))
	defer C.free(unsafe.Pointer(_arg1))

	C.lsm_dom_document_save_to_path(_arg0, _arg1, &_cerr)
	runtime.KeepAlive(documennt)
	runtime.KeepAlive(path)

	var _goerr error // out

	if _cerr != nil {
		_goerr = gerror.Take(unsafe.Pointer(_cerr))
	}

	return _goerr
}

// The function takes the following parameters:
//
//    - stream to save to.
//
func (document *DOMDocument) SaveToStream(stream gio.OutputStreamer) error {
	var _arg0 *C.LsmDomDocument // out
	var _arg1 *C.GOutputStream  // out
	var _cerr *C.GError         // in

	_arg0 = (*C.LsmDomDocument)(unsafe.Pointer(externglib.InternObject(document).Native()))
	_arg1 = (*C.GOutputStream)(unsafe.Pointer(externglib.InternObject(stream).Native()))

	C.lsm_dom_document_save_to_stream(_arg0, _arg1, &_cerr)
	runtime.KeepAlive(document)
	runtime.KeepAlive(stream)

	var _goerr error // out

	if _cerr != nil {
		_goerr = gerror.Take(unsafe.Pointer(_cerr))
	}

	return _goerr
}

// The function takes the following parameters:
//
func (documennt *DOMDocument) SaveToURL(path string) error {
	var _arg0 *C.LsmDomDocument // out
	var _arg1 *C.char           // out
	var _cerr *C.GError         // in

	_arg0 = (*C.LsmDomDocument)(unsafe.Pointer(externglib.InternObject(documennt).Native()))
	_arg1 = (*C.char)(unsafe.Pointer(C.CString(path)))
	defer C.free(unsafe.Pointer(_arg1))

	C.lsm_dom_document_save_to_url(_arg0, _arg1, &_cerr)
	runtime.KeepAlive(documennt)
	runtime.KeepAlive(path)

	var _goerr error // out

	if _cerr != nil {
		_goerr = gerror.Take(unsafe.Pointer(_cerr))
	}

	return _goerr
}
