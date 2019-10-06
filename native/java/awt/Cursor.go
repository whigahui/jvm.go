package awt

import (
	"github.com/zxh0/jvm.go/rtda"
	"github.com/zxh0/jvm.go/rtda/heap"
)

func init() {
}

func _cursor(method func(frame *rtda.Frame), name, desc string) {
	heap.RegisterNativeMethod("java/awt/Cursor", name, desc, method)
}