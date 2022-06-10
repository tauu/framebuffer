//go:build ignore
// +build ignore

package framebuffer

/*
#include <linux/fb.h>
*/
import "C"

const FBIOGET_FSCREENINFO = C.FBIOGET_FSCREENINFO
const FBIOGET_VSCREENINFO = C.FBIOGET_VSCREENINFO
const FBIOPUT_VSCREENINFO = C.FBIOPUT_VSCREENINFO
const FBIOPAN_DISPLAY = C.FBIOPAN_DISPLAY
