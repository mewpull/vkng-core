package core1_0

/*
#include <stdlib.h>
#include "vulkan/vulkan.h"
*/
import "C"
import (
	"github.com/CannibalVox/VKng/core/common"
	"github.com/CannibalVox/cgoparam"
	"unsafe"
)

type ClearRect struct {
	Rect           common.Rect2D
	BaseArrayLayer int
	LayerCount     int
}

func (r ClearRect) PopulateCPointer(allocator *cgoparam.Allocator, preallocatedPointer unsafe.Pointer) (unsafe.Pointer, error) {
	if preallocatedPointer == unsafe.Pointer(nil) {
		preallocatedPointer = allocator.Malloc(C.sizeof_struct_VkClearRect)
	}

	clearRect := (*C.VkClearRect)(preallocatedPointer)
	clearRect.baseArrayLayer = C.uint32_t(r.BaseArrayLayer)
	clearRect.layerCount = C.uint32_t(r.LayerCount)
	clearRect.rect.extent.width = C.uint32_t(r.Rect.Extent.Width)
	clearRect.rect.extent.height = C.uint32_t(r.Rect.Extent.Height)
	clearRect.rect.offset.x = C.int32_t(r.Rect.Offset.X)
	clearRect.rect.offset.y = C.int32_t(r.Rect.Offset.Y)

	return preallocatedPointer, nil
}
