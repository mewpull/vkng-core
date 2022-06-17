package core1_2

import (
	"github.com/CannibalVox/VKng/core/common"
	"github.com/CannibalVox/VKng/core/core1_0"
	"github.com/CannibalVox/VKng/core/core1_1"
	"github.com/CannibalVox/VKng/core/core1_2"
	"github.com/CannibalVox/VKng/core/driver"
)

type VulkanDescriptorSetLayout struct {
	core1_1.DescriptorSetLayout
}

func PromoteDescriptorSetLayout(layout core1_0.DescriptorSetLayout) core1_2.DescriptorSetLayout {
	if !layout.APIVersion().IsAtLeast(common.Vulkan1_2) {
		return nil
	}

	return layout.Driver().ObjectStore().GetOrCreate(
		driver.VulkanHandle(layout.Handle()),
		driver.Core1_2,
		func() any {
			return &VulkanDescriptorSetLayout{core1_1.PromoteDescriptorSetLayout(layout)}
		}).(core1_2.DescriptorSetLayout)
}
