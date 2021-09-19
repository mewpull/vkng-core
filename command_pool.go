package core

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

type vulkanCommandPool struct {
	driver Driver
	handle VkCommandPool
	device VkDevice
}

func (p *vulkanCommandPool) Handle() VkCommandPool {
	return p.handle
}

func (p *vulkanCommandPool) Destroy() error {
	return p.driver.VkDestroyCommandPool(p.device, p.handle, nil)
}

func (p *vulkanCommandPool) FreeCommandBuffers(buffers []CommandBuffer) error {
	allocator := cgoparam.GetAlloc()
	defer cgoparam.ReturnAlloc(allocator)

	bufferCount := len(buffers)
	if bufferCount == 0 {
		return nil
	}

	destroyPtr := allocator.Malloc(bufferCount * int(unsafe.Sizeof([1]C.VkCommandBuffer{})))
	destroySlice := ([]VkCommandBuffer)(unsafe.Slice((*VkCommandBuffer)(destroyPtr), bufferCount))
	for i := 0; i < bufferCount; i++ {
		destroySlice[i] = buffers[i].Handle()
	}

	return p.driver.VkFreeCommandBuffers(p.device, p.handle, Uint32(bufferCount), (*VkCommandBuffer)(destroyPtr))
}

func (p *vulkanCommandPool) AllocateCommandBuffers(o *CommandBufferOptions) ([]CommandBuffer, VkResult, error) {
	arena := cgoparam.GetAlloc()
	defer cgoparam.ReturnAlloc(arena)

	o.commandPool = p

	createInfo, err := common.AllocOptions(arena, o)
	if err != nil {
		return nil, VKErrorUnknown, err
	}

	commandBufferPtr := (*VkCommandBuffer)(arena.Malloc(o.BufferCount * int(unsafe.Sizeof([1]VkCommandBuffer{}))))

	res, err := p.driver.VkAllocateCommandBuffers(p.device, (*VkCommandBufferAllocateInfo)(createInfo), commandBufferPtr)
	err = res.ToError()
	if err != nil {
		return nil, res, err
	}

	commandBufferArray := ([]VkCommandBuffer)(unsafe.Slice(commandBufferPtr, o.BufferCount))
	var result []CommandBuffer
	for i := 0; i < o.BufferCount; i++ {
		result = append(result, &vulkanCommandBuffer{driver: p.driver, pool: p.handle, device: p.device, handle: commandBufferArray[i]})
	}

	return result, res, nil
}
