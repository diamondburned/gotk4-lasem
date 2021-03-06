// Code generated by girgen. DO NOT EDIT.

package lasem

import (
	"runtime"
	"unsafe"

	"github.com/diamondburned/gotk4/pkg/core/gbox"
	"github.com/diamondburned/gotk4/pkg/core/gerror"
	externglib "github.com/diamondburned/gotk4/pkg/core/glib"
	"github.com/diamondburned/gotk4/pkg/gio/v2"
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
// extern LsmDomNodeType _gotk4_lasem0_DomNodeClass_get_node_type(LsmDomNode*);
// extern char* _gotk4_lasem0_DomNodeClass_get_node_name(LsmDomNode*);
// extern char* _gotk4_lasem0_DomNodeClass_get_node_value(LsmDomNode*);
// extern gboolean _gotk4_lasem0_DomNodeClass_can_append_child(LsmDomNode*, LsmDomNode*);
// extern gboolean _gotk4_lasem0_DomNodeClass_child_changed(LsmDomNode*, LsmDomNode*);
// extern void _gotk4_lasem0_DomNodeClass_changed(LsmDomNode*);
// extern void _gotk4_lasem0_DomNodeClass_post_new_child(LsmDomNode*, LsmDomNode*);
// extern void _gotk4_lasem0_DomNodeClass_pre_remove_child(LsmDomNode*, LsmDomNode*);
// extern void _gotk4_lasem0_DomNodeClass_set_node_value(LsmDomNode*, char*);
// extern void _gotk4_lasem0_DomNodeClass_write_to_stream(LsmDomNode*, GOutputStream*, GError**);
import "C"

// glib.Type values for lsmdomnode.go.
var GTypeDOMNode = externglib.Type(C.lsm_dom_node_get_type())

func init() {
	externglib.RegisterGValueMarshalers([]externglib.TypeMarshaler{
		{T: GTypeDOMNode, F: marshalDOMNode},
	})
}

// DOMNodeOverrider contains methods that are overridable.
type DOMNodeOverrider interface {
	// The function takes the following parameters:
	//
	// The function returns the following values:
	//
	CanAppendChild(newChild DOMNoder) bool
	Changed()
	// The function takes the following parameters:
	//
	// The function returns the following values:
	//
	ChildChanged(child DOMNoder) bool
	// NodeName gets the node name.
	//
	// The function returns the following values:
	//
	//    - utf8: node name.
	//
	NodeName() string
	// The function returns the following values:
	//
	NodeType() DOMNodeType
	// NodeValue gets the node value.
	//
	// The function returns the following values:
	//
	//    - utf8: node value.
	//
	NodeValue() string
	// The function takes the following parameters:
	//
	PostNewChild(child DOMNoder)
	// The function takes the following parameters:
	//
	PreRemoveChild(child DOMNoder)
	// The function takes the following parameters:
	//
	SetNodeValue(newValue string)
	// The function takes the following parameters:
	//
	WriteToStream(stream gio.OutputStreamer) error
}

type DOMNode struct {
	_ [0]func() // equal guard
	*externglib.Object
}

var (
	_ externglib.Objector = (*DOMNode)(nil)
)

// DOMNoder describes types inherited from class DOMNode.
//
// To get the original type, the caller must assert this to an interface or
// another type.
type DOMNoder interface {
	externglib.Objector
	baseDOMNode() *DOMNode
}

var _ DOMNoder = (*DOMNode)(nil)

func classInitDOMNoder(gclassPtr, data C.gpointer) {
	C.g_type_class_add_private(gclassPtr, C.gsize(unsafe.Sizeof(uintptr(0))))

	goffset := C.g_type_class_get_instance_private_offset(gclassPtr)
	*(*C.gpointer)(unsafe.Add(unsafe.Pointer(gclassPtr), goffset)) = data

	goval := gbox.Get(uintptr(data))
	pclass := (*C.LsmDomNodeClass)(unsafe.Pointer(gclassPtr))
	// gclass := (*C.GTypeClass)(unsafe.Pointer(gclassPtr))
	// pclass := (*C.LsmDomNodeClass)(unsafe.Pointer(C.g_type_class_peek_parent(gclass)))

	if _, ok := goval.(interface{ CanAppendChild(newChild DOMNoder) bool }); ok {
		pclass.can_append_child = (*[0]byte)(C._gotk4_lasem0_DomNodeClass_can_append_child)
	}

	if _, ok := goval.(interface{ Changed() }); ok {
		pclass.changed = (*[0]byte)(C._gotk4_lasem0_DomNodeClass_changed)
	}

	if _, ok := goval.(interface{ ChildChanged(child DOMNoder) bool }); ok {
		pclass.child_changed = (*[0]byte)(C._gotk4_lasem0_DomNodeClass_child_changed)
	}

	if _, ok := goval.(interface{ NodeName() string }); ok {
		pclass.get_node_name = (*[0]byte)(C._gotk4_lasem0_DomNodeClass_get_node_name)
	}

	if _, ok := goval.(interface{ NodeType() DOMNodeType }); ok {
		pclass.get_node_type = (*[0]byte)(C._gotk4_lasem0_DomNodeClass_get_node_type)
	}

	if _, ok := goval.(interface{ NodeValue() string }); ok {
		pclass.get_node_value = (*[0]byte)(C._gotk4_lasem0_DomNodeClass_get_node_value)
	}

	if _, ok := goval.(interface{ PostNewChild(child DOMNoder) }); ok {
		pclass.post_new_child = (*[0]byte)(C._gotk4_lasem0_DomNodeClass_post_new_child)
	}

	if _, ok := goval.(interface{ PreRemoveChild(child DOMNoder) }); ok {
		pclass.pre_remove_child = (*[0]byte)(C._gotk4_lasem0_DomNodeClass_pre_remove_child)
	}

	if _, ok := goval.(interface{ SetNodeValue(newValue string) }); ok {
		pclass.set_node_value = (*[0]byte)(C._gotk4_lasem0_DomNodeClass_set_node_value)
	}

	if _, ok := goval.(interface {
		WriteToStream(stream gio.OutputStreamer) error
	}); ok {
		pclass.write_to_stream = (*[0]byte)(C._gotk4_lasem0_DomNodeClass_write_to_stream)
	}
}

//export _gotk4_lasem0_DomNodeClass_can_append_child
func _gotk4_lasem0_DomNodeClass_can_append_child(arg0 *C.LsmDomNode, arg1 *C.LsmDomNode) (cret C.gboolean) {
	goval := externglib.GoPrivateFromObject(unsafe.Pointer(arg0))
	iface := goval.(interface{ CanAppendChild(newChild DOMNoder) bool })

	var _newChild DOMNoder // out

	{
		objptr := unsafe.Pointer(arg1)
		if objptr == nil {
			panic("object of type lasem.DOMNoder is nil")
		}

		object := externglib.Take(objptr)
		casted := object.WalkCast(func(obj externglib.Objector) bool {
			_, ok := obj.(DOMNoder)
			return ok
		})
		rv, ok := casted.(DOMNoder)
		if !ok {
			panic("no marshaler for " + object.TypeFromInstance().String() + " matching lasem.DOMNoder")
		}
		_newChild = rv
	}

	ok := iface.CanAppendChild(_newChild)

	if ok {
		cret = C.TRUE
	}

	return cret
}

//export _gotk4_lasem0_DomNodeClass_changed
func _gotk4_lasem0_DomNodeClass_changed(arg0 *C.LsmDomNode) {
	goval := externglib.GoPrivateFromObject(unsafe.Pointer(arg0))
	iface := goval.(interface{ Changed() })

	iface.Changed()
}

//export _gotk4_lasem0_DomNodeClass_child_changed
func _gotk4_lasem0_DomNodeClass_child_changed(arg0 *C.LsmDomNode, arg1 *C.LsmDomNode) (cret C.gboolean) {
	goval := externglib.GoPrivateFromObject(unsafe.Pointer(arg0))
	iface := goval.(interface{ ChildChanged(child DOMNoder) bool })

	var _child DOMNoder // out

	{
		objptr := unsafe.Pointer(arg1)
		if objptr == nil {
			panic("object of type lasem.DOMNoder is nil")
		}

		object := externglib.Take(objptr)
		casted := object.WalkCast(func(obj externglib.Objector) bool {
			_, ok := obj.(DOMNoder)
			return ok
		})
		rv, ok := casted.(DOMNoder)
		if !ok {
			panic("no marshaler for " + object.TypeFromInstance().String() + " matching lasem.DOMNoder")
		}
		_child = rv
	}

	ok := iface.ChildChanged(_child)

	if ok {
		cret = C.TRUE
	}

	return cret
}

//export _gotk4_lasem0_DomNodeClass_get_node_name
func _gotk4_lasem0_DomNodeClass_get_node_name(arg0 *C.LsmDomNode) (cret *C.char) {
	goval := externglib.GoPrivateFromObject(unsafe.Pointer(arg0))
	iface := goval.(interface{ NodeName() string })

	utf8 := iface.NodeName()

	cret = (*C.char)(unsafe.Pointer(C.CString(utf8)))
	defer C.free(unsafe.Pointer(cret))

	return cret
}

//export _gotk4_lasem0_DomNodeClass_get_node_type
func _gotk4_lasem0_DomNodeClass_get_node_type(arg0 *C.LsmDomNode) (cret C.LsmDomNodeType) {
	goval := externglib.GoPrivateFromObject(unsafe.Pointer(arg0))
	iface := goval.(interface{ NodeType() DOMNodeType })

	domNodeType := iface.NodeType()

	cret = C.LsmDomNodeType(domNodeType)

	return cret
}

//export _gotk4_lasem0_DomNodeClass_get_node_value
func _gotk4_lasem0_DomNodeClass_get_node_value(arg0 *C.LsmDomNode) (cret *C.char) {
	goval := externglib.GoPrivateFromObject(unsafe.Pointer(arg0))
	iface := goval.(interface{ NodeValue() string })

	utf8 := iface.NodeValue()

	cret = (*C.char)(unsafe.Pointer(C.CString(utf8)))
	defer C.free(unsafe.Pointer(cret))

	return cret
}

//export _gotk4_lasem0_DomNodeClass_post_new_child
func _gotk4_lasem0_DomNodeClass_post_new_child(arg0 *C.LsmDomNode, arg1 *C.LsmDomNode) {
	goval := externglib.GoPrivateFromObject(unsafe.Pointer(arg0))
	iface := goval.(interface{ PostNewChild(child DOMNoder) })

	var _child DOMNoder // out

	{
		objptr := unsafe.Pointer(arg1)
		if objptr == nil {
			panic("object of type lasem.DOMNoder is nil")
		}

		object := externglib.Take(objptr)
		casted := object.WalkCast(func(obj externglib.Objector) bool {
			_, ok := obj.(DOMNoder)
			return ok
		})
		rv, ok := casted.(DOMNoder)
		if !ok {
			panic("no marshaler for " + object.TypeFromInstance().String() + " matching lasem.DOMNoder")
		}
		_child = rv
	}

	iface.PostNewChild(_child)
}

//export _gotk4_lasem0_DomNodeClass_pre_remove_child
func _gotk4_lasem0_DomNodeClass_pre_remove_child(arg0 *C.LsmDomNode, arg1 *C.LsmDomNode) {
	goval := externglib.GoPrivateFromObject(unsafe.Pointer(arg0))
	iface := goval.(interface{ PreRemoveChild(child DOMNoder) })

	var _child DOMNoder // out

	{
		objptr := unsafe.Pointer(arg1)
		if objptr == nil {
			panic("object of type lasem.DOMNoder is nil")
		}

		object := externglib.Take(objptr)
		casted := object.WalkCast(func(obj externglib.Objector) bool {
			_, ok := obj.(DOMNoder)
			return ok
		})
		rv, ok := casted.(DOMNoder)
		if !ok {
			panic("no marshaler for " + object.TypeFromInstance().String() + " matching lasem.DOMNoder")
		}
		_child = rv
	}

	iface.PreRemoveChild(_child)
}

//export _gotk4_lasem0_DomNodeClass_set_node_value
func _gotk4_lasem0_DomNodeClass_set_node_value(arg0 *C.LsmDomNode, arg1 *C.char) {
	goval := externglib.GoPrivateFromObject(unsafe.Pointer(arg0))
	iface := goval.(interface{ SetNodeValue(newValue string) })

	var _newValue string // out

	_newValue = C.GoString((*C.gchar)(unsafe.Pointer(arg1)))

	iface.SetNodeValue(_newValue)
}

//export _gotk4_lasem0_DomNodeClass_write_to_stream
func _gotk4_lasem0_DomNodeClass_write_to_stream(arg0 *C.LsmDomNode, arg1 *C.GOutputStream, _cerr **C.GError) {
	goval := externglib.GoPrivateFromObject(unsafe.Pointer(arg0))
	iface := goval.(interface {
		WriteToStream(stream gio.OutputStreamer) error
	})

	var _stream gio.OutputStreamer // out

	{
		objptr := unsafe.Pointer(arg1)
		if objptr == nil {
			panic("object of type gio.OutputStreamer is nil")
		}

		object := externglib.Take(objptr)
		casted := object.WalkCast(func(obj externglib.Objector) bool {
			_, ok := obj.(gio.OutputStreamer)
			return ok
		})
		rv, ok := casted.(gio.OutputStreamer)
		if !ok {
			panic("no marshaler for " + object.TypeFromInstance().String() + " matching gio.OutputStreamer")
		}
		_stream = rv
	}

	_goerr := iface.WriteToStream(_stream)

	if _goerr != nil && _cerr != nil {
		*_cerr = (*C.GError)(gerror.New(_goerr))
	}
}

func wrapDOMNode(obj *externglib.Object) *DOMNode {
	return &DOMNode{
		Object: obj,
	}
}

func marshalDOMNode(p uintptr) (interface{}, error) {
	return wrapDOMNode(externglib.ValueFromNative(unsafe.Pointer(p)).Object()), nil
}

func (self *DOMNode) baseDOMNode() *DOMNode {
	return self
}

// BaseDOMNode returns the underlying base object.
func BaseDOMNode(obj DOMNoder) *DOMNode {
	return obj.baseDOMNode()
}

// AppendChild adds the node new_child to the end of the list of children of
// this node. If the new_child is already in the tree, it is first removed.
//
// The function takes the following parameters:
//
//    - newChild: node to append.
//
// The function returns the following values:
//
//    - domNode: added node.
//
func (self *DOMNode) AppendChild(newChild DOMNoder) DOMNoder {
	var _arg0 *C.LsmDomNode // out
	var _arg1 *C.LsmDomNode // out
	var _cret *C.LsmDomNode // in

	_arg0 = (*C.LsmDomNode)(unsafe.Pointer(externglib.InternObject(self).Native()))
	_arg1 = (*C.LsmDomNode)(unsafe.Pointer(externglib.InternObject(newChild).Native()))
	C.g_object_ref(C.gpointer(externglib.InternObject(newChild).Native()))

	_cret = C.lsm_dom_node_append_child(_arg0, _arg1)
	runtime.KeepAlive(self)
	runtime.KeepAlive(newChild)

	var _domNode DOMNoder // out

	{
		objptr := unsafe.Pointer(_cret)
		if objptr == nil {
			panic("object of type lasem.DOMNoder is nil")
		}

		object := externglib.Take(objptr)
		casted := object.WalkCast(func(obj externglib.Objector) bool {
			_, ok := obj.(DOMNoder)
			return ok
		})
		rv, ok := casted.(DOMNoder)
		if !ok {
			panic("no marshaler for " + object.TypeFromInstance().String() + " matching lasem.DOMNoder")
		}
		_domNode = rv
	}

	return _domNode
}

func (self *DOMNode) Changed() {
	var _arg0 *C.LsmDomNode // out

	_arg0 = (*C.LsmDomNode)(unsafe.Pointer(externglib.InternObject(self).Native()))

	C.lsm_dom_node_changed(_arg0)
	runtime.KeepAlive(self)
}

// The function returns the following values:
//
//    - domNodeList: node child list.
//
func (self *DOMNode) ChildNodes() DOMNodeLister {
	var _arg0 *C.LsmDomNode     // out
	var _cret *C.LsmDomNodeList // in

	_arg0 = (*C.LsmDomNode)(unsafe.Pointer(externglib.InternObject(self).Native()))

	_cret = C.lsm_dom_node_get_child_nodes(_arg0)
	runtime.KeepAlive(self)

	var _domNodeList DOMNodeLister // out

	{
		objptr := unsafe.Pointer(_cret)
		if objptr == nil {
			panic("object of type lasem.DOMNodeLister is nil")
		}

		object := externglib.Take(objptr)
		casted := object.WalkCast(func(obj externglib.Objector) bool {
			_, ok := obj.(DOMNodeLister)
			return ok
		})
		rv, ok := casted.(DOMNodeLister)
		if !ok {
			panic("no marshaler for " + object.TypeFromInstance().String() + " matching lasem.DOMNodeLister")
		}
		_domNodeList = rv
	}

	return _domNodeList
}

// The function returns the following values:
//
//    - domNode: node first child.
//
func (self *DOMNode) FirstChild() DOMNoder {
	var _arg0 *C.LsmDomNode // out
	var _cret *C.LsmDomNode // in

	_arg0 = (*C.LsmDomNode)(unsafe.Pointer(externglib.InternObject(self).Native()))

	_cret = C.lsm_dom_node_get_first_child(_arg0)
	runtime.KeepAlive(self)

	var _domNode DOMNoder // out

	{
		objptr := unsafe.Pointer(_cret)
		if objptr == nil {
			panic("object of type lasem.DOMNoder is nil")
		}

		object := externglib.Take(objptr)
		casted := object.WalkCast(func(obj externglib.Objector) bool {
			_, ok := obj.(DOMNoder)
			return ok
		})
		rv, ok := casted.(DOMNoder)
		if !ok {
			panic("no marshaler for " + object.TypeFromInstance().String() + " matching lasem.DOMNoder")
		}
		_domNode = rv
	}

	return _domNode
}

// The function returns the following values:
//
//    - domNode: node last child.
//
func (self *DOMNode) LastChild() DOMNoder {
	var _arg0 *C.LsmDomNode // out
	var _cret *C.LsmDomNode // in

	_arg0 = (*C.LsmDomNode)(unsafe.Pointer(externglib.InternObject(self).Native()))

	_cret = C.lsm_dom_node_get_last_child(_arg0)
	runtime.KeepAlive(self)

	var _domNode DOMNoder // out

	{
		objptr := unsafe.Pointer(_cret)
		if objptr == nil {
			panic("object of type lasem.DOMNoder is nil")
		}

		object := externglib.Take(objptr)
		casted := object.WalkCast(func(obj externglib.Objector) bool {
			_, ok := obj.(DOMNoder)
			return ok
		})
		rv, ok := casted.(DOMNoder)
		if !ok {
			panic("no marshaler for " + object.TypeFromInstance().String() + " matching lasem.DOMNoder")
		}
		_domNode = rv
	}

	return _domNode
}

// The function returns the following values:
//
//    - domNode: node next sibling.
//
func (self *DOMNode) NextSibling() DOMNoder {
	var _arg0 *C.LsmDomNode // out
	var _cret *C.LsmDomNode // in

	_arg0 = (*C.LsmDomNode)(unsafe.Pointer(externglib.InternObject(self).Native()))

	_cret = C.lsm_dom_node_get_next_sibling(_arg0)
	runtime.KeepAlive(self)

	var _domNode DOMNoder // out

	{
		objptr := unsafe.Pointer(_cret)
		if objptr == nil {
			panic("object of type lasem.DOMNoder is nil")
		}

		object := externglib.Take(objptr)
		casted := object.WalkCast(func(obj externglib.Objector) bool {
			_, ok := obj.(DOMNoder)
			return ok
		})
		rv, ok := casted.(DOMNoder)
		if !ok {
			panic("no marshaler for " + object.TypeFromInstance().String() + " matching lasem.DOMNoder")
		}
		_domNode = rv
	}

	return _domNode
}

// NodeName gets the node name.
//
// The function returns the following values:
//
//    - utf8: node name.
//
func (self *DOMNode) NodeName() string {
	var _arg0 *C.LsmDomNode // out
	var _cret *C.char       // in

	_arg0 = (*C.LsmDomNode)(unsafe.Pointer(externglib.InternObject(self).Native()))

	_cret = C.lsm_dom_node_get_node_name(_arg0)
	runtime.KeepAlive(self)

	var _utf8 string // out

	_utf8 = C.GoString((*C.gchar)(unsafe.Pointer(_cret)))

	return _utf8
}

// The function returns the following values:
//
func (self *DOMNode) NodeType() DOMNodeType {
	var _arg0 *C.LsmDomNode    // out
	var _cret C.LsmDomNodeType // in

	_arg0 = (*C.LsmDomNode)(unsafe.Pointer(externglib.InternObject(self).Native()))

	_cret = C.lsm_dom_node_get_node_type(_arg0)
	runtime.KeepAlive(self)

	var _domNodeType DOMNodeType // out

	_domNodeType = DOMNodeType(_cret)

	return _domNodeType
}

// NodeValue gets the node value.
//
// The function returns the following values:
//
//    - utf8: node value.
//
func (self *DOMNode) NodeValue() string {
	var _arg0 *C.LsmDomNode // out
	var _cret *C.char       // in

	_arg0 = (*C.LsmDomNode)(unsafe.Pointer(externglib.InternObject(self).Native()))

	_cret = C.lsm_dom_node_get_node_value(_arg0)
	runtime.KeepAlive(self)

	var _utf8 string // out

	_utf8 = C.GoString((*C.gchar)(unsafe.Pointer(_cret)))

	return _utf8
}

// The function returns the following values:
//
//    - domDocument: node owner document.
//
func (self *DOMNode) OwnerDocument() DOMDocumenter {
	var _arg0 *C.LsmDomNode     // out
	var _cret *C.LsmDomDocument // in

	_arg0 = (*C.LsmDomNode)(unsafe.Pointer(externglib.InternObject(self).Native()))

	_cret = C.lsm_dom_node_get_owner_document(_arg0)
	runtime.KeepAlive(self)

	var _domDocument DOMDocumenter // out

	{
		objptr := unsafe.Pointer(_cret)
		if objptr == nil {
			panic("object of type lasem.DOMDocumenter is nil")
		}

		object := externglib.Take(objptr)
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

// The function returns the following values:
//
//    - domNode: node parent.
//
func (self *DOMNode) ParentNode() DOMNoder {
	var _arg0 *C.LsmDomNode // out
	var _cret *C.LsmDomNode // in

	_arg0 = (*C.LsmDomNode)(unsafe.Pointer(externglib.InternObject(self).Native()))

	_cret = C.lsm_dom_node_get_parent_node(_arg0)
	runtime.KeepAlive(self)

	var _domNode DOMNoder // out

	{
		objptr := unsafe.Pointer(_cret)
		if objptr == nil {
			panic("object of type lasem.DOMNoder is nil")
		}

		object := externglib.Take(objptr)
		casted := object.WalkCast(func(obj externglib.Objector) bool {
			_, ok := obj.(DOMNoder)
			return ok
		})
		rv, ok := casted.(DOMNoder)
		if !ok {
			panic("no marshaler for " + object.TypeFromInstance().String() + " matching lasem.DOMNoder")
		}
		_domNode = rv
	}

	return _domNode
}

// The function returns the following values:
//
//    - domNode: node previous sibling.
//
func (self *DOMNode) PreviousSibling() DOMNoder {
	var _arg0 *C.LsmDomNode // out
	var _cret *C.LsmDomNode // in

	_arg0 = (*C.LsmDomNode)(unsafe.Pointer(externglib.InternObject(self).Native()))

	_cret = C.lsm_dom_node_get_previous_sibling(_arg0)
	runtime.KeepAlive(self)

	var _domNode DOMNoder // out

	{
		objptr := unsafe.Pointer(_cret)
		if objptr == nil {
			panic("object of type lasem.DOMNoder is nil")
		}

		object := externglib.Take(objptr)
		casted := object.WalkCast(func(obj externglib.Objector) bool {
			_, ok := obj.(DOMNoder)
			return ok
		})
		rv, ok := casted.(DOMNoder)
		if !ok {
			panic("no marshaler for " + object.TypeFromInstance().String() + " matching lasem.DOMNoder")
		}
		_domNode = rv
	}

	return _domNode
}

// The function returns the following values:
//
func (self *DOMNode) HasChildNodes() bool {
	var _arg0 *C.LsmDomNode // out
	var _cret C.gboolean    // in

	_arg0 = (*C.LsmDomNode)(unsafe.Pointer(externglib.InternObject(self).Native()))

	_cret = C.lsm_dom_node_has_child_nodes(_arg0)
	runtime.KeepAlive(self)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// InsertBefore inserts the node new_child before the existing child node
// ref_child. If ref_child is null, insert new_child at the end of the list of
// children. If the new_child is already in the tree, it is first removed.
//
// The function takes the following parameters:
//
//    - newChild: node to insert.
//    - refChild: reference node, i.e., the node before which the new node must
//      be inserted.
//
// The function returns the following values:
//
//    - domNode: inserted node.
//
func (self *DOMNode) InsertBefore(newChild, refChild DOMNoder) DOMNoder {
	var _arg0 *C.LsmDomNode // out
	var _arg1 *C.LsmDomNode // out
	var _arg2 *C.LsmDomNode // out
	var _cret *C.LsmDomNode // in

	_arg0 = (*C.LsmDomNode)(unsafe.Pointer(externglib.InternObject(self).Native()))
	_arg1 = (*C.LsmDomNode)(unsafe.Pointer(externglib.InternObject(newChild).Native()))
	C.g_object_ref(C.gpointer(externglib.InternObject(newChild).Native()))
	_arg2 = (*C.LsmDomNode)(unsafe.Pointer(externglib.InternObject(refChild).Native()))

	_cret = C.lsm_dom_node_insert_before(_arg0, _arg1, _arg2)
	runtime.KeepAlive(self)
	runtime.KeepAlive(newChild)
	runtime.KeepAlive(refChild)

	var _domNode DOMNoder // out

	{
		objptr := unsafe.Pointer(_cret)
		if objptr == nil {
			panic("object of type lasem.DOMNoder is nil")
		}

		object := externglib.Take(objptr)
		casted := object.WalkCast(func(obj externglib.Objector) bool {
			_, ok := obj.(DOMNoder)
			return ok
		})
		rv, ok := casted.(DOMNoder)
		if !ok {
			panic("no marshaler for " + object.TypeFromInstance().String() + " matching lasem.DOMNoder")
		}
		_domNode = rv
	}

	return _domNode
}

// RemoveChild removes the child node indicated by old_child from the list of
// children, and returns it.
//
// The function takes the following parameters:
//
//    - oldChild: node to remove.
//
// The function returns the following values:
//
//    - domNode: removed node.
//
func (self *DOMNode) RemoveChild(oldChild DOMNoder) DOMNoder {
	var _arg0 *C.LsmDomNode // out
	var _arg1 *C.LsmDomNode // out
	var _cret *C.LsmDomNode // in

	_arg0 = (*C.LsmDomNode)(unsafe.Pointer(externglib.InternObject(self).Native()))
	_arg1 = (*C.LsmDomNode)(unsafe.Pointer(externglib.InternObject(oldChild).Native()))

	_cret = C.lsm_dom_node_remove_child(_arg0, _arg1)
	runtime.KeepAlive(self)
	runtime.KeepAlive(oldChild)

	var _domNode DOMNoder // out

	{
		objptr := unsafe.Pointer(_cret)
		if objptr == nil {
			panic("object of type lasem.DOMNoder is nil")
		}

		object := externglib.AssumeOwnership(objptr)
		casted := object.WalkCast(func(obj externglib.Objector) bool {
			_, ok := obj.(DOMNoder)
			return ok
		})
		rv, ok := casted.(DOMNoder)
		if !ok {
			panic("no marshaler for " + object.TypeFromInstance().String() + " matching lasem.DOMNoder")
		}
		_domNode = rv
	}

	return _domNode
}

// ReplaceChild replaces the child node old_child with new_child in the list of
// children, and returns the old_child node. If the new_child is already in the
// tree, it is first removed.
//
// The function takes the following parameters:
//
//    - newChild: replacement node.
//    - oldChild: node to replace.
//
// The function returns the following values:
//
//    - domNode: replaced node.
//
func (self *DOMNode) ReplaceChild(newChild, oldChild DOMNoder) DOMNoder {
	var _arg0 *C.LsmDomNode // out
	var _arg1 *C.LsmDomNode // out
	var _arg2 *C.LsmDomNode // out
	var _cret *C.LsmDomNode // in

	_arg0 = (*C.LsmDomNode)(unsafe.Pointer(externglib.InternObject(self).Native()))
	_arg1 = (*C.LsmDomNode)(unsafe.Pointer(externglib.InternObject(newChild).Native()))
	C.g_object_ref(C.gpointer(externglib.InternObject(newChild).Native()))
	_arg2 = (*C.LsmDomNode)(unsafe.Pointer(externglib.InternObject(oldChild).Native()))

	_cret = C.lsm_dom_node_replace_child(_arg0, _arg1, _arg2)
	runtime.KeepAlive(self)
	runtime.KeepAlive(newChild)
	runtime.KeepAlive(oldChild)

	var _domNode DOMNoder // out

	{
		objptr := unsafe.Pointer(_cret)
		if objptr == nil {
			panic("object of type lasem.DOMNoder is nil")
		}

		object := externglib.AssumeOwnership(objptr)
		casted := object.WalkCast(func(obj externglib.Objector) bool {
			_, ok := obj.(DOMNoder)
			return ok
		})
		rv, ok := casted.(DOMNoder)
		if !ok {
			panic("no marshaler for " + object.TypeFromInstance().String() + " matching lasem.DOMNoder")
		}
		_domNode = rv
	}

	return _domNode
}

// The function takes the following parameters:
//
func (self *DOMNode) SetNodeValue(newValue string) {
	var _arg0 *C.LsmDomNode // out
	var _arg1 *C.char       // out

	_arg0 = (*C.LsmDomNode)(unsafe.Pointer(externglib.InternObject(self).Native()))
	_arg1 = (*C.char)(unsafe.Pointer(C.CString(newValue)))
	defer C.free(unsafe.Pointer(_arg1))

	C.lsm_dom_node_set_node_value(_arg0, _arg1)
	runtime.KeepAlive(self)
	runtime.KeepAlive(newValue)
}

// The function takes the following parameters:
//
func (self *DOMNode) WriteToStream(stream gio.OutputStreamer) error {
	var _arg0 *C.LsmDomNode    // out
	var _arg1 *C.GOutputStream // out
	var _cerr *C.GError        // in

	_arg0 = (*C.LsmDomNode)(unsafe.Pointer(externglib.InternObject(self).Native()))
	_arg1 = (*C.GOutputStream)(unsafe.Pointer(externglib.InternObject(stream).Native()))

	C.lsm_dom_node_write_to_stream(_arg0, _arg1, &_cerr)
	runtime.KeepAlive(self)
	runtime.KeepAlive(stream)

	var _goerr error // out

	if _cerr != nil {
		_goerr = gerror.Take(unsafe.Pointer(_cerr))
	}

	return _goerr
}
