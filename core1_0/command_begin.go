package core1_0

/*
#include <stdlib.h>
#include "vulkan/vulkan.h"
*/
import "C"
import (
	"github.com/CannibalVox/cgoparam"
	"github.com/vkngwrapper/core/common"
	"unsafe"
)

const (
	// CommandBufferUsageOneTimeSubmit specifies that each recording of the CommandBuffer will only
	// be submitted once, and the CommandBuffer will be reset and recorded again between each submission
	//
	// https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkCommandBufferUsageFlagBits.html
	CommandBufferUsageOneTimeSubmit CommandBufferUsageFlags = C.VK_COMMAND_BUFFER_USAGE_ONE_TIME_SUBMIT_BIT
	// CommandBufferUsageRenderPassContinue specifies that a secondary CommandBuffer is considered to
	// be entirely inside a RenderPass.
	//
	// https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkCommandBufferUsageFlagBits.html
	CommandBufferUsageRenderPassContinue CommandBufferUsageFlags = C.VK_COMMAND_BUFFER_USAGE_RENDER_PASS_CONTINUE_BIT
	// CommandBufferUsageSimultaneousUse specifies that a CommandBuffer can be resubmitted to a Queue
	// while it is in the pending state, and recorded into multiple primary CommandBuffer objects
	//
	// https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkCommandBufferUsageFlagBits.html
	CommandBufferUsageSimultaneousUse CommandBufferUsageFlags = C.VK_COMMAND_BUFFER_USAGE_SIMULTANEOUS_USE_BIT

	// QueryControlPrecise specifies the precision of occlusion queries
	//
	// https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkQueryControlFlagBits.html
	QueryControlPrecise QueryControlFlags = C.VK_QUERY_CONTROL_PRECISE_BIT

	// QueryPipelineStatisticInputAssemblyVertices specifies that queries managed by the pool
	// will count the number of vertices processed by the input assembly stage.
	//
	// https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkQueryPipelineStatisticFlagBits.html
	QueryPipelineStatisticInputAssemblyVertices QueryPipelineStatisticFlags = C.VK_QUERY_PIPELINE_STATISTIC_INPUT_ASSEMBLY_VERTICES_BIT
	// QueryPipelineStatisticInputAssemblyPrimitives specifies that queries managed by the pool
	// will count the number of primitives processed by the input assembly stage.
	//
	// https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkQueryPipelineStatisticFlagBits.html
	QueryPipelineStatisticInputAssemblyPrimitives QueryPipelineStatisticFlags = C.VK_QUERY_PIPELINE_STATISTIC_INPUT_ASSEMBLY_PRIMITIVES_BIT
	// QueryPipelineStatisticVertexShaderInvocations specifies that queries managed by the pool
	// will count the number of vertex shader invocations.
	//
	// https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkQueryPipelineStatisticFlagBits.html
	QueryPipelineStatisticVertexShaderInvocations QueryPipelineStatisticFlags = C.VK_QUERY_PIPELINE_STATISTIC_VERTEX_SHADER_INVOCATIONS_BIT
	// QueryPipelineStatisticGeometryShaderInvocations specifies that queries managed by the pool
	// will count the number of geometry shader invocations.
	//
	// https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkQueryPipelineStatisticFlagBits.html
	QueryPipelineStatisticGeometryShaderInvocations QueryPipelineStatisticFlags = C.VK_QUERY_PIPELINE_STATISTIC_GEOMETRY_SHADER_INVOCATIONS_BIT
	// QueryPipelineStatisticGeometryShaderPrimitives specifies that queries managed by the pool will
	// count the number of primitives generated by geometry shader invocations.
	//
	// https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkQueryPipelineStatisticFlagBits.html
	QueryPipelineStatisticGeometryShaderPrimitives QueryPipelineStatisticFlags = C.VK_QUERY_PIPELINE_STATISTIC_GEOMETRY_SHADER_PRIMITIVES_BIT
	// QueryPipelineStatisticClippingInvocations specifies that queries managed by the pool will
	// count the number of primitives processed by the primitive clipping stage of the pipeline.
	//
	// https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkQueryPipelineStatisticFlagBits.html
	QueryPipelineStatisticClippingInvocations QueryPipelineStatisticFlags = C.VK_QUERY_PIPELINE_STATISTIC_CLIPPING_INVOCATIONS_BIT
	// QueryPipelineStatisticClippingPrimitives specifies that the queries managed by the pool
	// will count the number of primitives output by the primitive clipping stage of the pipeline.
	//
	// https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkQueryPipelineStatisticFlagBits.html
	QueryPipelineStatisticClippingPrimitives QueryPipelineStatisticFlags = C.VK_QUERY_PIPELINE_STATISTIC_CLIPPING_PRIMITIVES_BIT
	// QueryPipelineStatisticFragmentShaderInvocations specifies that the queries managed by the
	// pool will count the number of fragment shader invocations.
	//
	// https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkQueryPipelineStatisticFlagBits.html
	QueryPipelineStatisticFragmentShaderInvocations QueryPipelineStatisticFlags = C.VK_QUERY_PIPELINE_STATISTIC_FRAGMENT_SHADER_INVOCATIONS_BIT
	// QueryPipelineStatisticTessellationControlShaderPatches specifies that the queries managed by
	// the pool will count the number of patches processed by the tessellation control shader.
	//
	// https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkQueryPipelineStatisticFlagBits.html
	QueryPipelineStatisticTessellationControlShaderPatches QueryPipelineStatisticFlags = C.VK_QUERY_PIPELINE_STATISTIC_TESSELLATION_CONTROL_SHADER_PATCHES_BIT
	// QueryPipelineStatisticTessellationEvaluationShaderInvocations specifies that the queries managed
	// by the pool will count the number of invocations of the tessellation evaluation shader.
	//
	// https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkQueryPipelineStatisticFlagBits.html
	QueryPipelineStatisticTessellationEvaluationShaderInvocations QueryPipelineStatisticFlags = C.VK_QUERY_PIPELINE_STATISTIC_TESSELLATION_EVALUATION_SHADER_INVOCATIONS_BIT
	// QueryPipelineStatisticComputeShaderInvocations specifies that queries managed by the pool will
	// count the number of compute shader invocations.
	//
	// https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkQueryPipelineStatisticFlagBits.html
	QueryPipelineStatisticComputeShaderInvocations QueryPipelineStatisticFlags = C.VK_QUERY_PIPELINE_STATISTIC_COMPUTE_SHADER_INVOCATIONS_BIT
)

func init() {
	CommandBufferUsageOneTimeSubmit.Register("One Time Submit")
	CommandBufferUsageRenderPassContinue.Register("Render Pass Continue")
	CommandBufferUsageSimultaneousUse.Register("Simultaneous Use")

	QueryControlPrecise.Register("Precise")

	QueryPipelineStatisticInputAssemblyVertices.Register("Input Assembly Vertices")
	QueryPipelineStatisticInputAssemblyPrimitives.Register("Input Assembly Primitives")
	QueryPipelineStatisticVertexShaderInvocations.Register("Vertex Shader Invocations")
	QueryPipelineStatisticGeometryShaderInvocations.Register("Geometry Shader Invocations")
	QueryPipelineStatisticGeometryShaderPrimitives.Register("Geometry Shader Primitives")
	QueryPipelineStatisticClippingInvocations.Register("Clipping Invocations")
	QueryPipelineStatisticClippingPrimitives.Register("Clipping Primitives")
	QueryPipelineStatisticFragmentShaderInvocations.Register("Fragment Shader Invocations")
	QueryPipelineStatisticTessellationControlShaderPatches.Register("Tessellation Control Shader Patches")
	QueryPipelineStatisticTessellationEvaluationShaderInvocations.Register("Tessellation Evaluation Shader Invocations")
	QueryPipelineStatisticComputeShaderInvocations.Register("Compute Shader Invocations")
}

// CommandBufferInheritanceInfo specifies CommandBuffer inheritance information
//
// https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkCommandBufferInheritanceInfo.html
type CommandBufferInheritanceInfo struct {
	// Framebuffer refers to the Framebuffer object that the CommandBuffer will be rendering to
	// if it is executed within a RenderPass instance. It can be nil if the Framebuffer is not
	// known.
	Framebuffer Framebuffer
	// RenderPass is a RenderPass object defining which render passes the CommandBuffer will be
	// compatible with and can be executed within
	RenderPass RenderPass
	// Subpass is the index of hte subpass within the RenderPass instance that the CommandBuffer
	// will be executed within
	Subpass int

	// OcclusionQueryEnable specifies whether the CommandBuffer can be executed while an occlusion
	// query is active in the primary CommandBuffer
	OcclusionQueryEnable bool
	// QueryFlags specifies the query flags that can be used by an active occlusion query in the
	// primary CommandBuffer when this secondary CommandBuffer is executed
	QueryFlags QueryControlFlags
	// PipelineStatistics specifies the set of pipeline statistics that can be counted by an
	// active query in the primary CommandBuffer when this secondary CommandBuffer is executed
	PipelineStatistics QueryPipelineStatisticFlags

	common.NextOptions
}

func (o CommandBufferInheritanceInfo) PopulateCPointer(allocator *cgoparam.Allocator, preallocatedPointer unsafe.Pointer, next unsafe.Pointer) (unsafe.Pointer, error) {
	if preallocatedPointer == unsafe.Pointer(nil) {
		preallocatedPointer = allocator.Malloc(C.sizeof_struct_VkCommandBufferInheritanceInfo)
	}
	createInfo := (*C.VkCommandBufferInheritanceInfo)(preallocatedPointer)

	createInfo.sType = C.VK_STRUCTURE_TYPE_COMMAND_BUFFER_INHERITANCE_INFO
	createInfo.pNext = next

	createInfo.renderPass = nil
	createInfo.framebuffer = nil

	if o.Framebuffer != nil {
		createInfo.framebuffer = (C.VkFramebuffer)(unsafe.Pointer(o.Framebuffer.Handle()))
	}

	if o.RenderPass != nil {
		createInfo.renderPass = (C.VkRenderPass)(unsafe.Pointer(o.RenderPass.Handle()))
	}

	createInfo.subpass = C.uint32_t(o.Subpass)
	createInfo.occlusionQueryEnable = C.VK_FALSE

	if o.OcclusionQueryEnable {
		createInfo.occlusionQueryEnable = C.VK_TRUE
	}

	createInfo.queryFlags = (C.VkQueryControlFlags)(o.QueryFlags)
	createInfo.pipelineStatistics = (C.VkQueryPipelineStatisticFlags)(o.PipelineStatistics)

	return unsafe.Pointer(createInfo), nil
}

// CommandBufferBeginInfo specifies a CommandBuffer begin operation
//
// https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkCommandBufferBeginInfo.html
type CommandBufferBeginInfo struct {
	// Flags specifies usage behavior for the CommandBuffer
	Flags CommandBufferUsageFlags
	// InheritanceInfo specifies inheritance from the primary to secondary CommandBuffer. If this
	// structure is used with a primary CommandBuffer, then this field is ignored
	InheritanceInfo *CommandBufferInheritanceInfo

	common.NextOptions
}

func (o CommandBufferBeginInfo) PopulateCPointer(allocator *cgoparam.Allocator, preallocatedPointer unsafe.Pointer, next unsafe.Pointer) (unsafe.Pointer, error) {
	if preallocatedPointer == unsafe.Pointer(nil) {
		preallocatedPointer = allocator.Malloc(C.sizeof_struct_VkCommandBufferBeginInfo)
	}
	createInfo := (*C.VkCommandBufferBeginInfo)(preallocatedPointer)

	createInfo.sType = C.VK_STRUCTURE_TYPE_COMMAND_BUFFER_BEGIN_INFO
	createInfo.flags = C.VkCommandBufferUsageFlags(o.Flags)
	createInfo.pNext = next

	createInfo.pInheritanceInfo = nil

	if o.InheritanceInfo != nil {
		info, err := common.AllocOptions(allocator, *o.InheritanceInfo)
		if err != nil {
			return nil, err
		}
		createInfo.pInheritanceInfo = (*C.VkCommandBufferInheritanceInfo)(info)
	}

	return unsafe.Pointer(createInfo), nil
}
