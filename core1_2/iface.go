package core1_2

import (
	"github.com/vkngwrapper/core/common"
	"github.com/vkngwrapper/core/core1_0"
	"github.com/vkngwrapper/core/core1_1"
	"github.com/vkngwrapper/core/driver"
	"time"
)

//go:generate mockgen -source ./iface.go -destination ../mocks/core1_2_mocks.go -package mocks -mock_names CommandBuffer=CommandBuffer1_2,CommandPool=CommandPool1_2,Device=Device1_2,DescriptorUpdateTemplate=DescriptorUpdateTemplate1_2,Instance=Instance1_2,PhysicalDevice=PhysicalDevice1_2,Buffer=Buffer1_2,BufferView=BufferView1_2,DescriptorPool=DescriptorPool1_2,DescriptorSet=DescriptorSet1_2,DescriptorSetLayout=DescriptorSetLayout1_2,DeviceMemory=DeviceMemory1_2,Event=Event1_2,Fence=Fence1_2,Framebuffer=Framebuffer1_2,Image=Image1_2,ImageView=ImageView1_2,Pipeline=Pipeline1_2,PipelineCache=PipelineCache1_2,PipelineLayout=PipelineLayout1_2,QueryPool=QueryPool1_2,Queue=Queue1_2,RenderPass=RenderPass1_2,Sampler=Sampler1_2,Semaphore=Semaphore1_2,ShaderModule=ShaderModule1_2,InstanceScopedPhysicalDevice=InstanceScopedPhysicalDevice1_2,SamplerYcbcrConversion=SamplerYcbcrConversion1_2,DescriptorUpdateTemplate=DescriptorUpdateTemplate1_2

// Buffer represents a linear array of data, which is used for various purposes by binding it
// to a graphics or compute pipeline.
//
// This interface includes all commands included in Vulkan 1.2.
//
// https://www.khronos.org/registry/vulkan/specs/1.3-extensions/man/html/VkBuffer.html
type Buffer interface {
	core1_1.Buffer
}

// BufferView represents a contiguous range of a buffer and a specific format to be used to
// interpret the data.
//
// This interface includes all commands included in Vulkan 1.2.
//
//https://www.khronos.org/registry/vulkan/specs/1.3-extensions/man/html/VkBufferView.html
type BufferView interface {
	core1_1.BufferView
}

// CommandBuffer is an object used to record commands which can be subsequently submitted to
// a device queue for execution.
//
// This interface includes all commands included in Vulkan 1.2.
//
// https://www.khronos.org/registry/vulkan/specs/1.3-extensions/man/html/VkCommandBuffer.html
type CommandBuffer interface {
	core1_1.CommandBuffer

	// CmdBeginRenderPass2 begins a new RenderPass
	//
	// renderPassBegin - Specifies the RenderPass to begin an instance of, and the Framebuffer the instance
	// uses
	//
	// subpassBegin - Contains information about the subpass which is about to begin rendering
	//
	// https://www.khronos.org/registry/vulkan/specs/1.3-extensions/man/html/vkCmdBeginRenderPass2.html
	CmdBeginRenderPass2(renderPassBegin core1_0.RenderPassBeginInfo, subpassBegin SubpassBeginInfo) error
	// CmdEndRenderPass2 ends the current RenderPass
	//
	// subpassEnd - Contains information about how the previous subpass will be ended
	//
	// https://www.khronos.org/registry/vulkan/specs/1.3-extensions/man/html/vkCmdEndRenderPass2.html
	CmdEndRenderPass2(subpassEnd SubpassEndInfo) error
	// CmdNextSubpass2 transitions to the next subpass of a RenderPass
	//
	// subpassBegin - Contains information about the subpass which is about to begin rendering.
	//
	// subpassEnd - Contains information about how the previous subpass will be ended.
	//
	// https://www.khronos.org/registry/vulkan/specs/1.3-extensions/man/html/vkCmdNextSubpass2.html
	CmdNextSubpass2(subpassBegin SubpassBeginInfo, subpassEnd SubpassEndInfo) error
	// CmdDrawIndexedIndirectCount draws with indirect parameters, indexed vertices, and draw count
	//
	// buffer - The Buffer containing draw parameters
	//
	// offset - The byte offset into buffer where parameters begin
	//
	// countBuffer - The Buffer containing the draw count
	//
	// countBufferOffset - The byte offset into countBuffer where the draw count begins
	//
	// maxDrawCount - Specifies the maximum number of draws that will be executed.
	//
	// stride - The byte stride between successive sets of draw parameters
	//
	// https://www.khronos.org/registry/vulkan/specs/1.3-extensions/man/html/vkCmdDrawIndexedIndirectCount.html
	CmdDrawIndexedIndirectCount(buffer core1_0.Buffer, offset uint64, countBuffer core1_0.Buffer, countBufferOffset uint64, maxDrawCount, stride int)
	// CmdDrawIndirectCount draws primitives with indirect parameters and draw count
	//
	// buffer - The Buffer containing draw parameters
	//
	// offset - The byte offset into buffer where parameters begin
	//
	// countBuffer - The Buffer containing the draw count
	//
	// countBufferOffset - The byte offset into countBuffer where the draw count begins
	//
	// maxDrawCount - Specifies the maximum number of draws that will be executed.
	//
	// stride - The byte stride between successive sets of draw parameters
	//
	// https://www.khronos.org/registry/vulkan/specs/1.3-extensions/man/html/vkCmdDrawIndirectCount.html
	CmdDrawIndirectCount(buffer core1_0.Buffer, offset uint64, countBuffer core1_0.Buffer, countBufferOffset uint64, maxDrawCount, stride int)
}

// CommandPool is an opaque object that CommandBuffer memory is allocated from
//
// This interface includes all commands included in Vulkan 1.2.
//
// https://www.khronos.org/registry/vulkan/specs/1.3-extensions/man/html/VkCommandPool.html
type CommandPool interface {
	core1_1.CommandPool
}

// DescriptorPool maintains a pool of descriptors, from which DescriptorSet objects are allocated.
//
// This interface includes all commands included in Vulkan 1.2.
//
// https://www.khronos.org/registry/vulkan/specs/1.3-extensions/man/html/VkDescriptorPool.html
type DescriptorPool interface {
	core1_1.DescriptorPool
}

// DescriptorSet is an opaque object allocated from a DescriptorPool
//
// This interface includes all commands included in Vulkan 1.2.
//
// https://www.khronos.org/registry/vulkan/specs/1.3-extensions/man/html/vkResetDescriptorPool.html
type DescriptorSet interface {
	core1_1.DescriptorSet
}

// DescriptorSetLayout is a group of zero or more descriptor bindings definitions.
//
// This interface includes all commands included in Vulkan 1.2.
//
// https://www.khronos.org/registry/vulkan/specs/1.3-extensions/man/html/VkDescriptorSetLayout.html
type DescriptorSetLayout interface {
	core1_1.DescriptorSetLayout
}

// Device represents a logical device on the host
//
// This interface includes all commands included in Vulkan 1.2.
//
// https://www.khronos.org/registry/vulkan/specs/1.3-extensions/man/html/VkDevice.html
type Device interface {
	core1_1.Device

	// CreateRenderPass2 creates a new RenderPass object
	//
	// allocator - Controls host memory allocation behavior
	//
	// options - Describes the parameters of the RenderPass
	//
	// https://www.khronos.org/registry/vulkan/specs/1.3-extensions/man/html/vkCreateRenderPass2.html
	CreateRenderPass2(allocator *driver.AllocationCallbacks, options RenderPassCreateInfo2) (core1_0.RenderPass, common.VkResult, error)

	// GetBufferDeviceAddress queries an address of a Buffer
	//
	// o - Specifies the Buffer to retrieve an address for
	//
	// https://www.khronos.org/registry/vulkan/specs/1.3-extensions/man/html/vkGetBufferDeviceAddress.html
	GetBufferDeviceAddress(o BufferDeviceAddressInfo) (uint64, error)
	// GetBufferOpaqueCaptureAddress queries an opaque capture address of a Buffer
	//
	// o - Specifies the Buffer to retrieve an address for
	//
	// https://www.khronos.org/registry/vulkan/specs/1.3-extensions/man/html/vkGetBufferOpaqueCaptureAddress.html
	GetBufferOpaqueCaptureAddress(o BufferDeviceAddressInfo) (uint64, error)
	// GetDeviceMemoryOpaqueCaptureAddress queries an opaque capture address of a DeviceMemory object
	//
	// o - Specifies the DeviceMemory object to retrieve an address for
	//
	// https://www.khronos.org/registry/vulkan/specs/1.3-extensions/man/html/vkGetDeviceMemoryOpaqueCaptureAddress.html
	GetDeviceMemoryOpaqueCaptureAddress(o DeviceMemoryOpaqueCaptureAddressInfo) (uint64, error)

	// SignalSemaphore signals a timeline Semaphore on the host
	//
	// o - Contains information about the signal operation
	//
	// https://www.khronos.org/registry/vulkan/specs/1.3-extensions/man/html/vkSignalSemaphore.html
	SignalSemaphore(o SemaphoreSignalInfo) (common.VkResult, error)
	// WaitSemaphores waits for timeline Semaphore objects on the host
	//
	// timeout - How long to wait before returning VKTimeout. May be common.NoTimeout to wait indefinitely.
	// The timeout is adjusted to the closest value allowed by the implementation timeout accuracy,
	// which may be substantially longer than the requested timeout.
	//
	// o - Contains information about the wait condition
	//
	// https://www.khronos.org/registry/vulkan/specs/1.3-extensions/man/html/vkWaitSemaphores.html
	WaitSemaphores(timeout time.Duration, o SemaphoreWaitInfo) (common.VkResult, error)
}

// DeviceMemory represents a block of memory on the device
//
// This interface includes all commands included in Vulkan 1.2.
//
// https://www.khronos.org/registry/vulkan/specs/1.3-extensions/man/html/VkDeviceMemory.html
type DeviceMemory interface {
	core1_1.DeviceMemory
}

// DescriptorUpdateTemplate specifies a mapping from descriptor update information in host memory to
// descriptors in a DescriptorSet
//
// This interface includes all commands included in Vulkan 1.2.
//
// https://www.khronos.org/registry/vulkan/specs/1.3-extensions/man/html/VkDescriptorUpdateTemplate.html
type DescriptorUpdateTemplate interface {
	core1_1.DescriptorUpdateTemplate
}

// Event is a synchronization primitive that can be used to insert fine-grained dependencies between
// commands submitted to the same queue, or between the host and a queue.
//
// This interface includes all commands included in Vulkan 1.2.
//
// https://www.khronos.org/registry/vulkan/specs/1.3-extensions/man/html/VkEvent.html
type Event interface {
	core1_1.Event
}

// Fence is a synchronization primitive that can be used to insert a dependency from a queue to
// the host.
//
// This interface includes all commands included in Vulkan 1.2.
//
// https://www.khronos.org/registry/vulkan/specs/1.3-extensions/man/html/VkFence.html
type Fence interface {
	core1_1.Fence
}

// Framebuffer represents a collection of specific memory attachments that a RenderPass uses
//
// This interface includes all commands included in Vulkan 1.2.
//
// https://www.khronos.org/registry/vulkan/specs/1.3-extensions/man/html/VkFramebuffer.html
type Framebuffer interface {
	core1_1.Framebuffer
}

// Image represents multidimensional arrays of data which can be used for various purposes.
//
// This interface includes all commands included in Vulkan 1.2.
//
// https://www.khronos.org/registry/vulkan/specs/1.3-extensions/man/html/VkImage.html
type Image interface {
	core1_1.Image
}

// ImageView represents contiguous ranges of Image subresources and contains additional metadata
//
// This interface includes all commands included in Vulkan 1.2.
//
// https://www.khronos.org/registry/vulkan/specs/1.3-extensions/man/html/VkImageView.html
type ImageView interface {
	core1_1.ImageView
}

// Instance stores per-application state for Vulkan
//
// This interface includes all commands included in Vulkan 1.2.
//
// https://www.khronos.org/registry/vulkan/specs/1.3-extensions/man/html/VkInstance.html
type Instance interface {
	core1_1.Instance
}

// InstanceScopedPhysicalDevice represents the instance-scoped functionality of a single complete
// implementation of Vulkan available to the host, of which there are a finite number.
//
// This interface includes all instance-scoped commands included in Vulkan 1.2.
//
// PhysicalDevice objects are unusual in that they exist between the Instance and (logical) Device level.
// As a result, PhysicalDevice is the only object that can be extended by both Instance and Device
// extensions. Consequently, there are some unusual cases in which a higher core version may be available
// for some PhysicalDevice functionality but not others. In order to represent this, physical devices
// are split into two objects at core1.1+, the PhysicalDevice and the "instance-scoped" PhysicalDevice.
//
// The InstanceScopedPhysicalDevice is usually available at the same core versions as PhysicalDevice, but
// in rare cases, a higher core version may be available.
//
// https://www.khronos.org/registry/vulkan/specs/1.3-extensions/man/html/VkPhysicalDevice.html
type InstanceScopedPhysicalDevice interface {
	core1_1.InstanceScopedPhysicalDevice
}

// PhysicalDevice represents the device-scoped functionality of a single complete
// implementation of Vulkan available to the host, of which there are a finite number.
//
// This interface includes all commands included in Vulkan 1.2.
//
// PhysicalDevice objects are unusual in that they exist between the Instance and (logical) Device level.
// As a result, PhysicalDevice is the only object that can be extended by both Instance and Device
// extensions. Consequently, there are some unusual cases in which a higher core version may be available
// for some PhysicalDevice functionality but not others. In order to represent this, physical devices
// are split into two objects at core1.1+, the PhysicalDevice and the "instance-scoped" PhysicalDevice.
//
// The InstanceScopedPhysicalDevice is usually available at the same core versions as PhysicalDevice, but
// in rare cases, a higher core version may be available.
//
// https://www.khronos.org/registry/vulkan/specs/1.3-extensions/man/html/VkPhysicalDevice.html
type PhysicalDevice interface {
	core1_1.PhysicalDevice

	// InstanceScopedPhysicalDevice1_2 returns the InstanceScopedPhysicalDevice that represents the
	// instance-scoped portion of this PhysicalDevice object's functionality. Since the instance-scoped
	// support is always equal-to-or-greater-than the device-scoped support, this method will always
	// return a functioning InstanceScopedPhysicalDevice
	InstanceScopedPhysicalDevice1_2() InstanceScopedPhysicalDevice
}

// Pipeline represents compute, ray tracing, and graphics pipelines
//
// This interface includes all commands included in Vulkan 1.2.
//
// https://www.khronos.org/registry/vulkan/specs/1.3-extensions/man/html/VkPipeline.html
type Pipeline interface {
	core1_1.Pipeline
}

// PipelineCache allows the result of Pipeline construction to be reused between Pipeline objects
// and between runs of an application.
//
// This interface includes all commands included in Vulkan 1.2.
//
// https://www.khronos.org/registry/vulkan/specs/1.3-extensions/man/html/VkPipelineCache.html
type PipelineCache interface {
	core1_1.PipelineCache
}

// PipelineLayout provides access to descriptor sets to Pipeline objects by combining zero or more
// descriptor sets and zero or more push constant ranges.
//
// This interface includes all commands included in Vulkan 1.2.
//
// https://www.khronos.org/registry/vulkan/specs/1.3-extensions/man/html/VkPipelineLayout.html
type PipelineLayout interface {
	core1_1.PipelineLayout
}

// QueryPool is a collection of a specific number of queries of a particular type.
//
// This interface includes all commands included in Vulkan 1.2.
//
// https://www.khronos.org/registry/vulkan/specs/1.3-extensions/man/html/VkQueryPool.html
type QueryPool interface {
	core1_1.QueryPool

	// Reset resets queries in this QueryPool
	//
	// firstQuery - The initial query index to reset
	//
	// queryCount - The number of queries to reset
	//
	// https://www.khronos.org/registry/vulkan/specs/1.3-extensions/man/html/vkResetQueryPool.html
	Reset(firstQuery, queryCount int)
}

// Queue represents a Device resource on which work is performed
//
// This interface includes all commands included in Vulkan 1.2.
//
// https://www.khronos.org/registry/vulkan/specs/1.3-extensions/man/html/VkQueue.html
type Queue interface {
	core1_1.Queue
}

// RenderPass represents a collection of attachments, subpasses, and dependencies between the subpasses
// and describes how the attachments are used over the course of the subpasses
//
// This interface includes all commands included in Vulkan 1.2.
//
// https://www.khronos.org/registry/vulkan/specs/1.3-extensions/man/html/VkRenderPass.html
type RenderPass interface {
	core1_1.RenderPass
}

// Sampler represents the state of an Image sampler, which is used by the implementation to read Image data
// and apply filtering and other transformations for the shader.
//
// This interface includes all commands included in Vulkan 1.2.
//
// https://www.khronos.org/registry/vulkan/specs/1.3-extensions/man/html/VkSampler.html
type Sampler interface {
	core1_1.Sampler
}

// SamplerYcbcrConversion is an opaque representation of a device-specific sampler YCbCr conversion
// description.
//
// This interface includes all commands included in Vulkan 1.2.
//
// https://www.khronos.org/registry/vulkan/specs/1.3-extensions/man/html/VkSamplerYcbcrConversion.html
type SamplerYcbcrConversion interface {
	core1_1.SamplerYcbcrConversion
}

// Semaphore is a synchronization primitive that can be used to insert a dependency between Queue operations
// or between a Queue operation and the host.
//
// This interface includes all commands included in Vulkan 1.2.
//
// https://www.khronos.org/registry/vulkan/specs/1.3-extensions/man/html/VkSemaphore.html
type Semaphore interface {
	core1_1.Semaphore

	// CounterValue queries the current state of this timeline Semaphore
	//
	// https://www.khronos.org/registry/vulkan/specs/1.3-extensions/man/html/vkGetSemaphoreCounterValue.html
	CounterValue() (uint64, common.VkResult, error)
}

// ShaderModule objects contain shader code and one or more entry points.
//
// This interface includes all commands included in Vulkan 1.2.
//
// https://www.khronos.org/registry/vulkan/specs/1.3-extensions/man/html/VkShaderModule.html
type ShaderModule interface {
	core1_1.ShaderModule
}
