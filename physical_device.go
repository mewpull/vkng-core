package core

/*
#include <stdlib.h>
#include "../vulkan/vulkan.h"
*/
import "C"
import (
	"github.com/CannibalVox/VKng"
	"github.com/CannibalVox/cgoalloc"
	"github.com/google/uuid"
	"unsafe"
)

type PhysicalDeviceHandle C.VkPhysicalDevice
type PhysicalDevice struct {
	handle C.VkPhysicalDevice
}

func (d *PhysicalDevice) Handle() PhysicalDeviceHandle {
	return PhysicalDeviceHandle(d.handle)
}

func (d *PhysicalDevice) QueueFamilyProperties(allocator cgoalloc.Allocator) ([]*VKng.QueueFamily, error) {
	count := (*C.uint32_t)(allocator.Malloc(int(unsafe.Sizeof(C.uint32_t(0)))))
	defer allocator.Free(unsafe.Pointer(count))

	C.vkGetPhysicalDeviceQueueFamilyProperties(d.handle, count, nil)

	if *count == 0 {
		return nil, nil
	}

	goCount := int(*count)

	allocatedHandles := allocator.Malloc(goCount * int(unsafe.Sizeof(C.VkQueueFamilyProperties{})))
	defer allocator.Free(allocatedHandles)
	familyProperties := ([]C.VkQueueFamilyProperties)(unsafe.Slice((*C.VkQueueFamilyProperties)(allocatedHandles), int(*count)))

	C.vkGetPhysicalDeviceQueueFamilyProperties(d.handle, count, (*C.VkQueueFamilyProperties)(allocatedHandles))

	var queueFamilies []*VKng.QueueFamily
	for i := 0; i < goCount; i++ {
		queueFamilies = append(queueFamilies, &VKng.QueueFamily{
			Flags:              VKng.QueueFlags(familyProperties[i].queueFlags),
			QueueCount:         uint32(familyProperties[i].queueCount),
			TimestampValidBits: uint32(familyProperties[i].timestampValidBits),
			MinImageTransferGranularity: VKng.Extent3D{
				Width:  uint32(familyProperties[i].minImageTransferGranularity.width),
				Height: uint32(familyProperties[i].minImageTransferGranularity.height),
				Depth:  uint32(familyProperties[i].minImageTransferGranularity.depth),
			},
		})
	}

	return queueFamilies, nil
}

func (d *PhysicalDevice) CreateDevice(allocator cgoalloc.Allocator, options *DeviceOptions) (*Device, VKng.Result, error) {
	arena := cgoalloc.CreateArenaAllocator(allocator)
	defer arena.FreeAll()

	createInfo, err := options.AllocForC(arena)
	if err != nil {
		return nil, VKng.VKErrorUnknown, err
	}

	var deviceHandle C.VkDevice
	res := VKng.Result(C.vkCreateDevice(d.handle, (*C.VkDeviceCreateInfo)(createInfo), nil, &deviceHandle))
	err = res.ToError()
	if err != nil {
		return nil, res, err
	}

	return &Device{handle: deviceHandle}, res, nil
}

func (d *PhysicalDevice) Properties(allocator cgoalloc.Allocator) (*VKng.PhysicalDeviceProperties, error) {
	properties := (*C.VkPhysicalDeviceProperties)(allocator.Malloc(int(unsafe.Sizeof([1]C.VkPhysicalDeviceProperties{}))))
	defer allocator.Free(unsafe.Pointer(properties))

	C.vkGetPhysicalDeviceProperties(d.handle, properties)

	return createPhysicalDeviceProperties(properties)
}

func (d *PhysicalDevice) Features(allocator cgoalloc.Allocator) (*VKng.PhysicalDeviceFeatures, error) {
	features := (*C.VkPhysicalDeviceFeatures)(allocator.Malloc(int(unsafe.Sizeof([1]C.VkPhysicalDeviceFeatures{}))))
	defer allocator.Free(unsafe.Pointer(features))

	C.vkGetPhysicalDeviceFeatures(d.handle, features)

	return createPhysicalDeviceFeatures(features), nil
}

func (d *PhysicalDevice) AvailableExtensions(allocator cgoalloc.Allocator) (map[string]*VKng.ExtensionProperties, VKng.Result, error) {
	extensionCountPtr := allocator.Malloc(int(unsafe.Sizeof(C.uint32_t(0))))
	defer allocator.Free(extensionCountPtr)

	extensionCount := (*C.uint32_t)(extensionCountPtr)

	res := VKng.Result(C.vkEnumerateDeviceExtensionProperties(d.handle, nil, extensionCount, nil))
	err := res.ToError()

	if err != nil || *extensionCount == 0 {
		return nil, res, err
	}

	extensionTotal := int(*extensionCount)
	extensionsPtr := allocator.Malloc(extensionTotal * int(unsafe.Sizeof([1]C.VkExtensionProperties{})))
	defer allocator.Free(extensionsPtr)

	typedExtensionsPtr := (*C.VkExtensionProperties)(extensionsPtr)
	res = VKng.Result(C.vkEnumerateDeviceExtensionProperties(d.handle, nil, extensionCount, typedExtensionsPtr))
	err = res.ToError()
	if err != nil {
		return nil, res, err
	}

	retVal := make(map[string]*VKng.ExtensionProperties)
	extensionSlice := ([]C.VkExtensionProperties)(unsafe.Slice(typedExtensionsPtr, extensionTotal))

	for i := 0; i < extensionTotal; i++ {
		extension := extensionSlice[i]

		outExtension := &VKng.ExtensionProperties{
			ExtensionName: C.GoString((*C.char)(&extension.extensionName[0])),
			SpecVersion:   VKng.Version(extension.specVersion),
		}

		existingExtension, ok := retVal[outExtension.ExtensionName]
		if ok && existingExtension.SpecVersion >= outExtension.SpecVersion {
			continue
		}
		retVal[outExtension.ExtensionName] = outExtension
	}

	return retVal, res, nil
}

func createPhysicalDeviceFeatures(f *C.VkPhysicalDeviceFeatures) *VKng.PhysicalDeviceFeatures {
	return &VKng.PhysicalDeviceFeatures{
		RobustBufferAccess:                      f.robustBufferAccess != C.VK_FALSE,
		FullDrawIndexUint32:                     f.fullDrawIndexUint32 != C.VK_FALSE,
		ImageCubeArray:                          f.imageCubeArray != C.VK_FALSE,
		IndependentBlend:                        f.independentBlend != C.VK_FALSE,
		GeometryShader:                          f.geometryShader != C.VK_FALSE,
		TessellationShader:                      f.tessellationShader != C.VK_FALSE,
		SampleRateShading:                       f.sampleRateShading != C.VK_FALSE,
		DualSrcBlend:                            f.dualSrcBlend != C.VK_FALSE,
		LogicOp:                                 f.logicOp != C.VK_FALSE,
		MultiDrawIndirect:                       f.multiDrawIndirect != C.VK_FALSE,
		DrawIndirectFirstInstance:               f.drawIndirectFirstInstance != C.VK_FALSE,
		DepthClamp:                              f.depthClamp != C.VK_FALSE,
		DepthBiasClamp:                          f.depthBiasClamp != C.VK_FALSE,
		FillModeNonSolid:                        f.fillModeNonSolid != C.VK_FALSE,
		DepthBounds:                             f.depthBounds != C.VK_FALSE,
		WideLines:                               f.wideLines != C.VK_FALSE,
		LargePoints:                             f.largePoints != C.VK_FALSE,
		AlphaToOne:                              f.alphaToOne != C.VK_FALSE,
		MultiViewport:                           f.multiViewport != C.VK_FALSE,
		SamplerAnisotropy:                       f.samplerAnisotropy != C.VK_FALSE,
		TextureCompressionEtc2:                  f.textureCompressionETC2 != C.VK_FALSE,
		TextureCompressionAstcLdc:               f.textureCompressionASTC_LDR != C.VK_FALSE,
		TextureCompressionBc:                    f.textureCompressionBC != C.VK_FALSE,
		OcclusionQueryPrecise:                   f.occlusionQueryPrecise != C.VK_FALSE,
		PipelineStatisticsQuery:                 f.pipelineStatisticsQuery != C.VK_FALSE,
		VertexPipelineStoresAndAtomics:          f.vertexPipelineStoresAndAtomics != C.VK_FALSE,
		FragmentStoresAndAtomics:                f.fragmentStoresAndAtomics != C.VK_FALSE,
		ShaderTessellationAndGeometryPointSize:  f.shaderTessellationAndGeometryPointSize != C.VK_FALSE,
		ShaderImageGatherExtended:               f.shaderImageGatherExtended != C.VK_FALSE,
		ShaderStorageImageExtendedFormats:       f.shaderStorageImageExtendedFormats != C.VK_FALSE,
		ShaderStorageImageMultisample:           f.shaderStorageImageMultisample != C.VK_FALSE,
		ShaderStorageImageReadWithoutFormat:     f.shaderStorageImageReadWithoutFormat != C.VK_FALSE,
		ShaderStorageImageWriteWithoutFormat:    f.shaderStorageImageWriteWithoutFormat != C.VK_FALSE,
		ShaderUniformBufferArrayDynamicIndexing: f.shaderUniformBufferArrayDynamicIndexing != C.VK_FALSE,
		ShaderSampledImageArrayDynamicIndexing:  f.shaderSampledImageArrayDynamicIndexing != C.VK_FALSE,
		ShaderStorageBufferArrayDynamicIndexing: f.shaderStorageBufferArrayDynamicIndexing != C.VK_FALSE,
		ShaderStorageImageArrayDynamicIndexing:  f.shaderStorageImageArrayDynamicIndexing != C.VK_FALSE,
		ShaderClipDistance:                      f.shaderClipDistance != C.VK_FALSE,
		ShaderCullDistance:                      f.shaderCullDistance != C.VK_FALSE,
		ShaderFloat64:                           f.shaderFloat64 != C.VK_FALSE,
		ShaderInt64:                             f.shaderInt64 != C.VK_FALSE,
		ShaderInt16:                             f.shaderInt16 != C.VK_FALSE,
		ShaderResourceResidency:                 f.shaderResourceResidency != C.VK_FALSE,
		ShaderResourceMinLod:                    f.shaderResourceMinLod != C.VK_FALSE,
		SparseBinding:                           f.sparseBinding != C.VK_FALSE,
		SparseResidencyBuffer:                   f.sparseResidencyBuffer != C.VK_FALSE,
		SparseResidencyImage2D:                  f.sparseResidencyImage2D != C.VK_FALSE,
		SparseResidencyImage3D:                  f.sparseResidencyImage3D != C.VK_FALSE,
		SparseResidency2Samples:                 f.sparseResidency2Samples != C.VK_FALSE,
		SparseResidency4Samples:                 f.sparseResidency4Samples != C.VK_FALSE,
		SparseResidency8Samples:                 f.sparseResidency8Samples != C.VK_FALSE,
		SparseResidency16Samples:                f.sparseResidency16Samples != C.VK_FALSE,
		SparseResidencyAliased:                  f.sparseResidencyAliased != C.VK_FALSE,
		VariableMultisampleRate:                 f.variableMultisampleRate != C.VK_FALSE,
		InheritedQueries:                        f.inheritedQueries != C.VK_FALSE,
	}
}

func createPhysicalDeviceLimits(l *C.VkPhysicalDeviceLimits) *VKng.PhysicalDeviceLimits {
	return &VKng.PhysicalDeviceLimits{
		MaxImageDimension1D:                             uint32(l.maxImageDimension1D),
		MaxImageDimension2D:                             uint32(l.maxImageDimension2D),
		MaxImageDimension3D:                             uint32(l.maxImageDimension3D),
		MaxImageDimensionCube:                           uint32(l.maxImageDimensionCube),
		MaxImageArrayLayers:                             uint32(l.maxImageArrayLayers),
		MaxTexelBufferElements:                          uint32(l.maxTexelBufferElements),
		MaxUniformBufferRange:                           uint32(l.maxUniformBufferRange),
		MaxStorageBufferRange:                           uint32(l.maxStorageBufferRange),
		MaxPushConstantsSize:                            uint32(l.maxPushConstantsSize),
		MaxMemoryAllocationCount:                        uint32(l.maxMemoryAllocationCount),
		MaxSamplerAllocationCount:                       uint32(l.maxSamplerAllocationCount),
		BufferImageGranularity:                          uint64(l.bufferImageGranularity),
		SparseAddressSpaceSize:                          uint64(l.sparseAddressSpaceSize),
		MaxBoundDescriptorSets:                          uint32(l.maxBoundDescriptorSets),
		MaxPerStageDescriptorSamplers:                   uint32(l.maxPerStageDescriptorSamplers),
		MaxPerStageDescriptorUniformBuffers:             uint32(l.maxPerStageDescriptorUniformBuffers),
		MaxPerStageDescriptorStorageBuffers:             uint32(l.maxPerStageDescriptorStorageBuffers),
		MaxPerStageDescriptorSampledImages:              uint32(l.maxPerStageDescriptorSampledImages),
		MaxPerStageDescriptorStorageImages:              uint32(l.maxPerStageDescriptorStorageImages),
		MaxPerStageDescriptorInputAttachments:           uint32(l.maxPerStageDescriptorInputAttachments),
		MaxPerStageResources:                            uint32(l.maxPerStageResources),
		MaxDescriptorSetSamplers:                        uint32(l.maxDescriptorSetSamplers),
		MaxDescriptorSetUniformBuffers:                  uint32(l.maxDescriptorSetUniformBuffers),
		MaxDescriptorSetUniformBuffersDynamic:           uint32(l.maxDescriptorSetUniformBuffersDynamic),
		MaxDescriptorSetStorageBuffers:                  uint32(l.maxDescriptorSetStorageBuffers),
		MaxDescriptorSetStorageBuffersDynamic:           uint32(l.maxDescriptorSetStorageBuffersDynamic),
		MaxDescriptorSetSampledImages:                   uint32(l.maxDescriptorSetSampledImages),
		MaxDescriptorSetStorageImages:                   uint32(l.maxDescriptorSetStorageImages),
		MaxDescriptorSetInputAttachments:                uint32(l.maxDescriptorSetInputAttachments),
		MaxVertexInputAttributes:                        uint32(l.maxVertexInputAttributes),
		MaxVertexInputBindings:                          uint32(l.maxVertexInputBindings),
		MaxVertexInputAttributeOffset:                   uint32(l.maxVertexInputAttributeOffset),
		MaxVertexInputBindingStride:                     uint32(l.maxVertexInputBindingStride),
		MaxVertexOutputComponents:                       uint32(l.maxVertexOutputComponents),
		MaxTessellationGenerationLevel:                  uint32(l.maxTessellationGenerationLevel),
		MaxTessellationPatchSize:                        uint32(l.maxTessellationPatchSize),
		MaxTessellationControlPerVertexInputComponents:  uint32(l.maxTessellationControlPerVertexInputComponents),
		MaxTessellationControlPerVertexOutputComponents: uint32(l.maxTessellationControlPerVertexOutputComponents),
		MaxTessellationControlPerPatchOutputComponents:  uint32(l.maxTessellationControlPerPatchOutputComponents),
		MaxTessellationControlTotalOutputComponents:     uint32(l.maxTessellationControlTotalOutputComponents),
		MaxTessellationEvaluationInputComponents:        uint32(l.maxTessellationEvaluationInputComponents),
		MaxTessellationEvaluationOutputComponents:       uint32(l.maxTessellationEvaluationOutputComponents),
		MaxGeometryShaderInvocations:                    uint32(l.maxGeometryShaderInvocations),
		MaxGeometryInputComponents:                      uint32(l.maxGeometryInputComponents),
		MaxGeometryOutputComponents:                     uint32(l.maxGeometryOutputComponents),
		MaxGeometryOutputVertices:                       uint32(l.maxGeometryOutputVertices),
		MaxGeometryTotalOutputComponents:                uint32(l.maxGeometryTotalOutputComponents),
		MaxFragmentInputComponents:                      uint32(l.maxFragmentInputComponents),
		MaxFragmentOutputAttachments:                    uint32(l.maxFragmentOutputAttachments),
		MaxFragmentDualSrcAttachments:                   uint32(l.maxFragmentDualSrcAttachments),
		MaxFragmentCombinedOutputResources:              uint32(l.maxFragmentCombinedOutputResources),
		MaxComputeSharedMemorySize:                      uint32(l.maxComputeSharedMemorySize),
		MaxComputeWorkGroupInvocations:                  uint32(l.maxComputeWorkGroupInvocations),
		SubPixelPrecisionBits:                           uint32(l.subPixelPrecisionBits),
		SubTexelPrecisionBits:                           uint32(l.subTexelPrecisionBits),
		MipmapPrecisionBits:                             uint32(l.mipmapPrecisionBits),
		MaxDrawIndexedIndexValue:                        uint32(l.maxDrawIndexedIndexValue),
		MaxDrawIndirectCount:                            uint32(l.maxDrawIndirectCount),
		MaxSamplerLodBias:                               float32(l.maxSamplerLodBias),
		MaxSamplerAnisotropy:                            float32(l.maxSamplerAnisotropy),
		MaxViewports:                                    uint32(l.maxViewports),
		ViewportSubPixelBits:                            uint32(l.viewportSubPixelBits),
		MinMemoryMapAlignment:                           uint(l.minMemoryMapAlignment),
		MinTexelBufferOffsetAlignment:                   uint64(l.minTexelBufferOffsetAlignment),
		MinUniformBufferOffsetAlignment:                 uint64(l.minUniformBufferOffsetAlignment),
		MinStorageBufferOffsetAlignment:                 uint64(l.minStorageBufferOffsetAlignment),
		MinTexelOffset:                                  int32(l.minTexelOffset),
		MaxTexelOffset:                                  uint32(l.maxTexelOffset),
		MinTexelGatherOffset:                            int32(l.minTexelGatherOffset),
		MaxTexelGatherOffset:                            uint32(l.maxTexelGatherOffset),
		MinInterpolationOffset:                          float32(l.minInterpolationOffset),
		MaxInterpolationOffset:                          float32(l.maxInterpolationOffset),
		SubPixelInterpolationOffsetBits:                 uint32(l.subPixelInterpolationOffsetBits),
		MaxFramebufferWidth:                             uint32(l.maxFramebufferWidth),
		MaxFramebufferHeight:                            uint32(l.maxFramebufferHeight),
		MaxFramebufferLayers:                            uint32(l.maxFramebufferLayers),
		FramebufferColorSampleCounts:                    VKng.SampleCounts(l.framebufferColorSampleCounts),
		FramebufferDepthSampleCounts:                    VKng.SampleCounts(l.framebufferDepthSampleCounts),
		FramebufferStencilSampleCounts:                  VKng.SampleCounts(l.framebufferStencilSampleCounts),
		FramebufferNoAttachmentsSampleCounts:            VKng.SampleCounts(l.framebufferNoAttachmentsSampleCounts),
		MaxColorAttachments:                             uint32(l.maxColorAttachments),
		SampledImageColorSampleCounts:                   VKng.SampleCounts(l.sampledImageColorSampleCounts),
		SampledImageIntegerSampleCounts:                 VKng.SampleCounts(l.sampledImageIntegerSampleCounts),
		SampledImageDepthSampleCounts:                   VKng.SampleCounts(l.sampledImageDepthSampleCounts),
		SampledImageStencilSampleCounts:                 VKng.SampleCounts(l.sampledImageStencilSampleCounts),
		StorageImageSampleCounts:                        VKng.SampleCounts(l.storageImageSampleCounts),
		MaxSampleMaskWords:                              uint32(l.maxSampleMaskWords),
		TimestampComputeAndGraphics:                     l.timestampComputeAndGraphics != C.VK_FALSE,
		TimestampPeriod:                                 float32(l.timestampPeriod),
		MaxClipDistances:                                uint32(l.maxClipDistances),
		MaxCullDistances:                                uint32(l.maxCullDistances),
		MaxCombinedClipAndCullDistances:                 uint32(l.maxCombinedClipAndCullDistances),
		DiscreteQueuePriorities:                         uint32(l.discreteQueuePriorities),
		PointSizeGranularity:                            float32(l.pointSizeGranularity),
		LineWidthGranularity:                            float32(l.lineWidthGranularity),
		StrictLines:                                     l.strictLines != C.VK_FALSE,
		StandardSampleLocations:                         l.standardSampleLocations != C.VK_FALSE,
		OptimalBufferCopyOffsetAlignment:                uint64(l.optimalBufferCopyOffsetAlignment),
		OptimalBufferCopyRowPitchAlignment:              uint64(l.optimalBufferCopyRowPitchAlignment),
		NonCoherentAtomSize:                             uint64(l.nonCoherentAtomSize),
		MaxComputeWorkGroupCount: [3]uint32{
			uint32(l.maxComputeWorkGroupCount[0]),
			uint32(l.maxComputeWorkGroupCount[1]),
			uint32(l.maxComputeWorkGroupCount[2]),
		},
		MaxComputeWorkGroupSize: [3]uint32{
			uint32(l.maxComputeWorkGroupSize[0]),
			uint32(l.maxComputeWorkGroupSize[1]),
			uint32(l.maxComputeWorkGroupSize[2]),
		},
		MaxViewportDimensions: [2]uint32{
			uint32(l.maxViewportDimensions[0]),
			uint32(l.maxViewportDimensions[1]),
		},
		ViewportBoundsRange: [2]float32{
			float32(l.viewportBoundsRange[0]),
			float32(l.viewportBoundsRange[1]),
		},
		PointSizeRange: [2]float32{
			float32(l.pointSizeRange[0]),
			float32(l.pointSizeRange[1]),
		},
		LineWidthRange: [2]float32{
			float32(l.lineWidthRange[0]),
			float32(l.lineWidthRange[1]),
		},
	}
}

func createSparseProperties(p *C.VkPhysicalDeviceSparseProperties) *VKng.PhysicalDeviceSparseProperties {
	return &VKng.PhysicalDeviceSparseProperties{
		ResidencyStandard2DBlockShape:            p.residencyStandard2DBlockShape != C.VK_FALSE,
		ResidencyStandard2DMultisampleBlockShape: p.residencyStandard2DMultisampleBlockShape != C.VK_FALSE,
		ResidencyStandard3DBlockShape:            p.residencyStandard3DBlockShape != C.VK_FALSE,
		ResidencyNonResidentStrict:               p.residencyNonResidentStrict != C.VK_FALSE,
		ResidencyAlignedMipSize:                  p.residencyAlignedMipSize != C.VK_FALSE,
	}
}

func createPhysicalDeviceProperties(p *C.VkPhysicalDeviceProperties) (*VKng.PhysicalDeviceProperties, error) {
	uuidBytes := C.GoBytes(unsafe.Pointer(&p.pipelineCacheUUID[0]), C.VK_UUID_SIZE)
	uuid, err := uuid.FromBytes(uuidBytes)
	if err != nil {
		return nil, err
	}

	return &VKng.PhysicalDeviceProperties{
		Type: VKng.PhysicalDeviceType(p.deviceType),
		Name: C.GoString((*C.char)(&p.deviceName[0])),

		APIVersion:    VKng.Version(p.apiVersion),
		DriverVersion: VKng.Version(p.driverVersion),

		VendorID: uint32(p.vendorID),
		DeviceID: uint32(p.deviceID),

		PipelineCacheUUID: uuid,
		Limits:            createPhysicalDeviceLimits(&p.limits),
		SparseProperties:  createSparseProperties(&p.sparseProperties),
	}, nil
}
