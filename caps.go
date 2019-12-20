
package gst


/*
#cgo pkg-config: gstreamer-1.0 gstreamer-base-1.0 gstreamer-app-1.0 gstreamer-plugins-base-1.0 gstreamer-video-1.0 gstreamer-audio-1.0 gstreamer-plugins-bad-1.0
#include "gst.h"
*/
import "C"


import (
	"unsafe"
	"runtime"
)



type Caps struct {
	caps *C.GstCaps
}

func CapsFromString(caps string) (gstCaps *Caps) {
	c := (*C.gchar)(unsafe.Pointer(C.CString(caps)))
	defer C.g_free(C.gpointer(unsafe.Pointer(c)))
	CCaps := C.gst_caps_from_string(c)
	gstCaps = &Caps{
		caps: CCaps,
	}

	runtime.SetFinalizer(gstCaps, func(gstCaps *Caps) {
		C.gst_caps_unref(gstCaps.caps)
	})

	return
}

func (c *Caps) ToString() (str string)  {
	CStr := C.gst_caps_to_string(c.caps)
	defer C.g_free(C.gpointer(unsafe.Pointer(CStr)))
	str = C.GoString((*C.char)(unsafe.Pointer(CStr)))

	return
}