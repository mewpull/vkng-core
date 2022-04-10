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

const (
	FenceCreateSignaled common.FenceCreateFlags = C.VK_FENCE_CREATE_SIGNALED_BIT
)

func init() {
	FenceCreateSignaled.Register("Signaled")
}

type FenceOptions struct {
	Flags common.FenceCreateFlags

	common.HaveNext
}

func (o FenceOptions) PopulateCPointer(allocator *cgoparam.Allocator, preallocatedPointer unsafe.Pointer, next unsafe.Pointer) (unsafe.Pointer, error) {
	if preallocatedPointer == unsafe.Pointer(nil) {
		preallocatedPointer = allocator.Malloc(C.sizeof_struct_VkFenceCreateInfo)
	}
	createInfo := (*C.VkFenceCreateInfo)(preallocatedPointer)
	createInfo.sType = C.VK_STRUCTURE_TYPE_FENCE_CREATE_INFO
	createInfo.flags = C.VkFenceCreateFlags(o.Flags)
	createInfo.pNext = next

	return unsafe.Pointer(createInfo), nil
}

func (o FenceOptions) PopulateOutData(cDataPointer unsafe.Pointer, helpers ...any) (next unsafe.Pointer, err error) {
	createInfo := (*C.VkFenceCreateInfo)(cDataPointer)
	return createInfo.pNext, nil
}
