package core

/*
#include <stdlib.h>
#include "vulkan/vulkan.h"
*/
import "C"
import (
	"github.com/CannibalVox/VKng/core/common"
	driver3 "github.com/CannibalVox/VKng/core/driver"
	"github.com/CannibalVox/cgoparam"
	"unsafe"
)

type vulkanEvent struct {
	handle driver3.VkEvent
	device driver3.VkDevice
	driver driver3.Driver
}

func (e *vulkanEvent) Handle() driver3.VkEvent {
	return e.handle
}

func (e *vulkanEvent) Destroy(callbacks *AllocationCallbacks) {
	e.driver.VkDestroyEvent(e.device, e.handle, callbacks.Handle())
}

func (e *vulkanEvent) Set() (common.VkResult, error) {
	return e.driver.VkSetEvent(e.device, e.handle)
}

func (e *vulkanEvent) Reset() (common.VkResult, error) {
	return e.driver.VkResetEvent(e.device, e.handle)
}

func (e *vulkanEvent) Status() (common.VkResult, error) {
	return e.driver.VkGetEventStatus(e.device, e.handle)
}

type EventFlags int32

const (
	EventDeviceOnlyKHR EventFlags = C.VK_EVENT_CREATE_DEVICE_ONLY_BIT_KHR
)

var eventFlagsToString = map[EventFlags]string{
	EventDeviceOnlyKHR: "Device Only (Khronos Extension)",
}

func (f EventFlags) String() string {
	return common.FlagsToString(f, eventFlagsToString)
}

type EventOptions struct {
	Flags EventFlags

	common.HaveNext
}

func (o *EventOptions) AllocForC(allocator *cgoparam.Allocator, next unsafe.Pointer) (unsafe.Pointer, error) {
	createInfo := (*C.VkEventCreateInfo)(allocator.Malloc(C.sizeof_struct_VkEventCreateInfo))
	createInfo.sType = C.VK_STRUCTURE_TYPE_EVENT_CREATE_INFO
	createInfo.flags = C.VkEventCreateFlags(o.Flags)
	createInfo.pNext = next

	return unsafe.Pointer(createInfo), nil
}
