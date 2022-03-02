// Code generated by girgen. DO NOT EDIT.

package lasem

import (
	"unsafe"

	"github.com/diamondburned/gotk4/pkg/core/gextras"
)

// #include <stdlib.h>
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

// SVGAngle: instance of this type is always passed by reference.
type SVGAngle struct {
	*svgAngle
}

// svgAngle is the struct that's finalized.
type svgAngle struct {
	native *C.LsmSvgAngle
}

func (s *SVGAngle) Type() SVGAngleType {
	var v SVGAngleType // out
	v = SVGAngleType(s.native._type)
	return v
}

func (s *SVGAngle) Angle() float64 {
	var v float64 // out
	v = float64(s.native.angle)
	return v
}

// SVGColor: instance of this type is always passed by reference.
type SVGColor struct {
	*svgColor
}

// svgColor is the struct that's finalized.
type svgColor struct {
	native *C.LsmSvgColor
}

// NewSVGColor creates a new SVGColor instance from the given
// fields.
func NewSVGColor(red, green, blue float64) SVGColor {
	var f0 C.double // out
	f0 = C.double(red)
	var f1 C.double // out
	f1 = C.double(green)
	var f2 C.double // out
	f2 = C.double(blue)

	v := C.LsmSvgColor{
		red:   f0,
		green: f1,
		blue:  f2,
	}

	return *(*SVGColor)(gextras.NewStructNative(unsafe.Pointer(&v)))
}

func (s *SVGColor) Red() float64 {
	var v float64 // out
	v = float64(s.native.red)
	return v
}

func (s *SVGColor) Green() float64 {
	var v float64 // out
	v = float64(s.native.green)
	return v
}

func (s *SVGColor) Blue() float64 {
	var v float64 // out
	v = float64(s.native.blue)
	return v
}

// SVGDashArray: instance of this type is always passed by reference.
type SVGDashArray struct {
	*svgDashArray
}

// svgDashArray is the struct that's finalized.
type svgDashArray struct {
	native *C.LsmSvgDashArray
}

func (s *SVGDashArray) NDashes() uint {
	var v uint // out
	v = uint(s.native.n_dashes)
	return v
}

func (s *SVGDashArray) Dashes() *SVGLength {
	var v *SVGLength // out
	v = (*SVGLength)(gextras.NewStructNative(unsafe.Pointer(s.native.dashes)))
	return v
}

// SVGOneOrTwoDouble: instance of this type is always passed by reference.
type SVGOneOrTwoDouble struct {
	*svgOneOrTwoDouble
}

// svgOneOrTwoDouble is the struct that's finalized.
type svgOneOrTwoDouble struct {
	native *C.LsmSvgOneOrTwoDouble
}

// NewSVGOneOrTwoDouble creates a new SVGOneOrTwoDouble instance from the given
// fields.
func NewSVGOneOrTwoDouble(a, b float64) SVGOneOrTwoDouble {
	var f0 C.double // out
	f0 = C.double(a)
	var f1 C.double // out
	f1 = C.double(b)

	v := C.LsmSvgOneOrTwoDouble{
		a: f0,
		b: f1,
	}

	return *(*SVGOneOrTwoDouble)(gextras.NewStructNative(unsafe.Pointer(&v)))
}

func (s *SVGOneOrTwoDouble) A() float64 {
	var v float64 // out
	v = float64(s.native.a)
	return v
}

func (s *SVGOneOrTwoDouble) B() float64 {
	var v float64 // out
	v = float64(s.native.b)
	return v
}

// SVGOneOrTwoInteger: instance of this type is always passed by reference.
type SVGOneOrTwoInteger struct {
	*svgOneOrTwoInteger
}

// svgOneOrTwoInteger is the struct that's finalized.
type svgOneOrTwoInteger struct {
	native *C.LsmSvgOneOrTwoInteger
}

// NewSVGOneOrTwoInteger creates a new SVGOneOrTwoInteger instance from the given
// fields.
func NewSVGOneOrTwoInteger(a, b int) SVGOneOrTwoInteger {
	var f0 C.int // out
	f0 = C.int(a)
	var f1 C.int // out
	f1 = C.int(b)

	v := C.LsmSvgOneOrTwoInteger{
		a: f0,
		b: f1,
	}

	return *(*SVGOneOrTwoInteger)(gextras.NewStructNative(unsafe.Pointer(&v)))
}

func (s *SVGOneOrTwoInteger) A() int {
	var v int // out
	v = int(s.native.a)
	return v
}

func (s *SVGOneOrTwoInteger) B() int {
	var v int // out
	v = int(s.native.b)
	return v
}

// SVGPaint: instance of this type is always passed by reference.
type SVGPaint struct {
	*svgPaint
}

// svgPaint is the struct that's finalized.
type svgPaint struct {
	native *C.LsmSvgPaint
}

func (s *SVGPaint) Type() SVGPaintType {
	var v SVGPaintType // out
	v = SVGPaintType(s.native._type)
	return v
}

func (s *SVGPaint) URL() string {
	var v string // out
	v = C.GoString((*C.gchar)(unsafe.Pointer(s.native.url)))
	return v
}

func (s *SVGPaint) Color() *SVGColor {
	var v *SVGColor // out
	v = (*SVGColor)(gextras.NewStructNative(unsafe.Pointer((&s.native.color))))
	return v
}

// SVGPreserveAspectRatio: instance of this type is always passed by reference.
type SVGPreserveAspectRatio struct {
	*svgPreserveAspectRatio
}

// svgPreserveAspectRatio is the struct that's finalized.
type svgPreserveAspectRatio struct {
	native *C.LsmSvgPreserveAspectRatio
}

func (s *SVGPreserveAspectRatio) Defer() bool {
	var v bool // out
	if s.native._defer != 0 {
		v = true
	}
	return v
}

func (s *SVGPreserveAspectRatio) Align() SVGAlign {
	var v SVGAlign // out
	v = SVGAlign(s.native.align)
	return v
}

func (s *SVGPreserveAspectRatio) MeetOrSlice() SVGMeetOrSlice {
	var v SVGMeetOrSlice // out
	v = SVGMeetOrSlice(s.native.meet_or_slice)
	return v
}

// SVGVector: instance of this type is always passed by reference.
type SVGVector struct {
	*svgVector
}

// svgVector is the struct that's finalized.
type svgVector struct {
	native *C.LsmSvgVector
}

func (s *SVGVector) NValues() uint {
	var v uint // out
	v = uint(s.native.n_values)
	return v
}

func (s *SVGVector) Values() *float64 {
	var v *float64 // out
	v = (*float64)(unsafe.Pointer(s.native.values))
	return v
}
