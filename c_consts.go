//go:build linux && cgo
// +build linux,cgo

package framebuffer

/*
#include <linux/fb.h>
*/
import "C"

const FBIOGET_FSCREENINFO = C.FBIOGET_FSCREENINFO
const FBIOGET_VSCREENINFO = C.FBIOGET_VSCREENINFO
const FBIOPAN_DISPLAY = C.FBIOPAN_DISPLAY
