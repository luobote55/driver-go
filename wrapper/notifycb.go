package wrapper

/*
#include <stdio.h>
#include <stdlib.h>
#include <string.h>
#include <taos.h>
*/
import "C"
import (
	"unsafe"

	"github.com/luobote55/driver-go/v3/common"
	"github.com/luobote55/driver-go/v3/wrapper/cgo"
)

//export NotifyCallback
func NotifyCallback(p unsafe.Pointer, ext unsafe.Pointer, notifyType C.int) {
	defer func() {
		recover()
	}()
	if int(notifyType) == common.TAOS_NOTIFY_PASSVER {
		version := int32(*(*C.int32_t)(ext))
		c := (*(*cgo.Handle)(p)).Value().(chan int32)
		c <- version
	}
}
