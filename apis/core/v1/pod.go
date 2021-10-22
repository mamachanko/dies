/*
Copyright 2021 the original author or authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package v1

import (
	diemetav1 "dies.dev/apis/meta/v1"
	corev1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/api/resource"
)

// +die:object=true
type _ = corev1.Pod

// +die
type _ = corev1.PodSpec

type podSpecDieExtension interface {
	VolumeDie(name string, fn func(d VolumeDie)) PodSpecDie
	InitContainerDie(name string, fn func(d ContainerDie)) PodSpecDie
	ContainerDie(name string, fn func(d ContainerDie)) PodSpecDie
	SecurityContextDie(fn func(d PodSecurityContextDie)) PodSpecDie
	TolerationDie(key string, fn func(d TolerationDie)) PodSpecDie
	HostAliasesDie(hosts ...HostAliasDie) PodSpecDie
	DNSConfigDie(fn func(d PodDNSConfigDie)) PodSpecDie
	ReadinessGatesDie(gates ...PodReadinessGateDie) PodSpecDie
	AddOverhead(name corev1.ResourceName, quantity resource.Quantity) PodSpecDie
	AddOverheadString(name corev1.ResourceName, quantity string) PodSpecDie
	TopologySpreadConstraintDie(topologyKey string, fn func(d TopologySpreadConstraintDie)) PodSpecDie
}

func (d *podSpecDie) VolumeDie(name string, fn func(d VolumeDie)) PodSpecDie {
	return d.DieStamp(func(r *corev1.PodSpec) {
		for i := range r.Volumes {
			if name == r.Volumes[i].Name {
				d := VolumeBlank.DieImmutable(false).DieFeed(r.Volumes[i])
				fn(d)
				r.Volumes[i] = d.DieRelease()
				return
			}
		}

		d := VolumeBlank.DieImmutable(false).DieFeed(corev1.Volume{Name: name})
		fn(d)
		r.Volumes = append(r.Volumes, d.DieRelease())
	})
}

func (d *podSpecDie) InitContainerDie(name string, fn func(d ContainerDie)) PodSpecDie {
	return d.DieStamp(func(r *corev1.PodSpec) {
		for i := range r.InitContainers {
			if name == r.InitContainers[i].Name {
				d := ContainerBlank.DieImmutable(false).DieFeed(r.InitContainers[i])
				fn(d)
				r.InitContainers[i] = d.DieRelease()
				return
			}
		}

		d := ContainerBlank.DieImmutable(false).DieFeed(corev1.Container{Name: name})
		fn(d)
		r.InitContainers = append(r.InitContainers, d.DieRelease())
	})
}

func (d *podSpecDie) ContainerDie(name string, fn func(d ContainerDie)) PodSpecDie {
	return d.DieStamp(func(r *corev1.PodSpec) {
		for i := range r.Containers {
			if name == r.Containers[i].Name {
				d := ContainerBlank.DieImmutable(false).DieFeed(r.Containers[i])
				fn(d)
				r.Containers[i] = d.DieRelease()
				return
			}
		}

		d := ContainerBlank.DieImmutable(false).DieFeed(corev1.Container{Name: name})
		fn(d)
		r.Containers = append(r.Containers, d.DieRelease())
	})
}

func (d *podSpecDie) SecurityContextDie(fn func(d PodSecurityContextDie)) PodSpecDie {
	return d.DieStamp(func(r *corev1.PodSpec) {
		d := PodSecurityContextBlank.DieImmutable(false).DieFeedPtr(r.SecurityContext)
		fn(d)
		r.SecurityContext = d.DieReleasePtr()
	})
}

func (d *podSpecDie) TolerationDie(key string, fn func(d TolerationDie)) PodSpecDie {
	return d.DieStamp(func(r *corev1.PodSpec) {
		for i := range r.Tolerations {
			if key == r.Tolerations[i].Key {
				d := TolerationBlank.DieImmutable(false).DieFeed(r.Tolerations[i])
				fn(d)
				r.Tolerations[i] = d.DieRelease()
				return
			}
		}

		d := TolerationBlank.DieImmutable(false).DieFeed(corev1.Toleration{Key: key})
		fn(d)
		r.Tolerations = append(r.Tolerations, d.DieRelease())
	})
}

func (d *podSpecDie) HostAliasesDie(hosts ...HostAliasDie) PodSpecDie {
	return d.DieStamp(func(r *corev1.PodSpec) {
		r.HostAliases = make([]corev1.HostAlias, len(hosts))
		for i := range hosts {
			r.HostAliases[i] = hosts[i].DieRelease()
		}
	})
}

func (d *podSpecDie) DNSConfigDie(fn func(d PodDNSConfigDie)) PodSpecDie {
	return d.DieStamp(func(r *corev1.PodSpec) {
		d := PodDNSConfigBlank.DieImmutable(false).DieFeedPtr(r.DNSConfig)
		fn(d)
		r.DNSConfig = d.DieReleasePtr()
	})
}

func (d *podSpecDie) ReadinessGatesDie(gates ...PodReadinessGateDie) PodSpecDie {
	return d.DieStamp(func(r *corev1.PodSpec) {
		r.ReadinessGates = make([]corev1.PodReadinessGate, len(gates))
		for i := range gates {
			r.ReadinessGates[i] = gates[i].DieRelease()
		}
	})
}

func (d *podSpecDie) AddOverhead(name corev1.ResourceName, quantity resource.Quantity) PodSpecDie {
	return d.DieStamp(func(r *corev1.PodSpec) {
		r.Overhead[name] = quantity
	})
}

func (d *podSpecDie) AddOverheadString(name corev1.ResourceName, quantity string) PodSpecDie {
	return d.AddOverhead(name, resource.MustParse(quantity))
}

func (d *podSpecDie) TopologySpreadConstraintDie(topologyKey string, fn func(d TopologySpreadConstraintDie)) PodSpecDie {
	return d.DieStamp(func(r *corev1.PodSpec) {
		for i := range r.TopologySpreadConstraints {
			if topologyKey == r.TopologySpreadConstraints[i].TopologyKey {
				d := TopologySpreadConstraintBlank.DieImmutable(false).DieFeed(r.TopologySpreadConstraints[i])
				fn(d)
				r.TopologySpreadConstraints[i] = d.DieRelease()
				return
			}
		}

		d := TopologySpreadConstraintBlank.DieImmutable(false).DieFeed(corev1.TopologySpreadConstraint{TopologyKey: topologyKey})
		fn(d)
		r.TopologySpreadConstraints = append(r.TopologySpreadConstraints, d.DieRelease())
	})
}

// +die
type _ = corev1.PodSecurityContext

type podSecurityContextDieExtension interface {
	SELinuxOptionsDie(fn func(d SELinuxOptionsDie)) PodSecurityContextDie
	WindowsOptionsDie(fn func(d WindowsSecurityContextOptionsDie)) PodSecurityContextDie
	SysctlsDie(sysctls ...SysctlDie) PodSecurityContextDie
	SeccompProfileDie(fn func(d SeccompProfileDie)) PodSecurityContextDie
}

func (d *podSecurityContextDie) SELinuxOptionsDie(fn func(d SELinuxOptionsDie)) PodSecurityContextDie {
	return d.DieStamp(func(r *corev1.PodSecurityContext) {
		d := SELinuxOptionsBlank.DieImmutable(false).DieFeedPtr(r.SELinuxOptions)
		fn(d)
		r.SELinuxOptions = d.DieReleasePtr()
	})
}

func (d *podSecurityContextDie) WindowsOptionsDie(fn func(d WindowsSecurityContextOptionsDie)) PodSecurityContextDie {
	return d.DieStamp(func(r *corev1.PodSecurityContext) {
		d := WindowsSecurityContextOptionsBlank.DieImmutable(false).DieFeedPtr(r.WindowsOptions)
		fn(d)
		r.WindowsOptions = d.DieReleasePtr()
	})
}

func (d *podSecurityContextDie) SysctlsDie(sysctls ...SysctlDie) PodSecurityContextDie {
	return d.DieStamp(func(r *corev1.PodSecurityContext) {
		r.Sysctls = make([]corev1.Sysctl, len(sysctls))
		for i := range sysctls {
			r.Sysctls[i] = sysctls[i].DieRelease()
		}
	})
}

func (d *podSecurityContextDie) SeccompProfileDie(fn func(d SeccompProfileDie)) PodSecurityContextDie {
	return d.DieStamp(func(r *corev1.PodSecurityContext) {
		d := SeccompProfileBlank.DieImmutable(false).DieFeedPtr(r.SeccompProfile)
		fn(d)
		r.SeccompProfile = d.DieReleasePtr()
	})
}

// +die
type _ = corev1.Sysctl

// +die
type _ = corev1.Toleration

// +die
type _ = corev1.HostAlias

// +die
type _ = corev1.PodDNSConfig

type podDNSConfigDieExtension interface {
	OptionsDie(options ...PodDNSConfigOptionDie) PodDNSConfigDie
}

func (d *podDNSConfigDie) OptionsDie(options ...PodDNSConfigOptionDie) PodDNSConfigDie {
	return d.DieStamp(func(r *corev1.PodDNSConfig) {
		r.Options = make([]corev1.PodDNSConfigOption, len(options))
		for i := range options {
			r.Options[i] = options[i].DieRelease()
		}
	})
}

// +die
type _ = corev1.PodDNSConfigOption

// +die
type _ = corev1.PodReadinessGate

// +die
type _ = corev1.TopologySpreadConstraint

type topologySpreadConstraintDieExtension interface {
	LabelSelectorDie(fn func(d diemetav1.LabelSelectorDie)) TopologySpreadConstraintDie
}

func (d *topologySpreadConstraintDie) LabelSelectorDie(fn func(d diemetav1.LabelSelectorDie)) TopologySpreadConstraintDie {
	return d.DieStamp(func(r *corev1.TopologySpreadConstraint) {
		d := diemetav1.LabelSelectorBlank.DieImmutable(false).DieFeedPtr(r.LabelSelector)
		fn(d)
		r.LabelSelector = d.DieReleasePtr()
	})
}

// +die
type _ = corev1.PodStatus

type podStatusDieExtension interface {
	ConditionsDie(conditions ...diemetav1.ConditionDie) PodStatusDie
	InitContainerStatusDie(name string, fn func(d ContainerStatusDie)) PodStatusDie
	ContainerStatusDie(name string, fn func(d ContainerStatusDie)) PodStatusDie
	EphemeralContainerStatusDie(name string, fn func(d ContainerStatusDie)) PodStatusDie
}

func (d *podStatusDie) ConditionsDie(conditions ...diemetav1.ConditionDie) PodStatusDie {
	return d.DieStamp(func(r *corev1.PodStatus) {
		r.Conditions = make([]corev1.PodCondition, len(conditions))
		for i := range conditions {
			c := conditions[i].DieRelease()
			r.Conditions[i] = corev1.PodCondition{
				Type:               corev1.PodConditionType(c.Type),
				Status:             corev1.ConditionStatus(c.Status),
				Reason:             c.Reason,
				Message:            c.Message,
				LastTransitionTime: c.LastTransitionTime,
			}
		}
	})
}

func (d *podStatusDie) InitContainerStatusDie(name string, fn func(d ContainerStatusDie)) PodStatusDie {
	return d.DieStamp(func(r *corev1.PodStatus) {
		for i := range r.InitContainerStatuses {
			if name == r.InitContainerStatuses[i].Name {
				d := ContainerStatusBlank.DieImmutable(false).DieFeed(r.InitContainerStatuses[i])
				fn(d)
				r.InitContainerStatuses[i] = d.DieRelease()
				return
			}
		}

		d := ContainerStatusBlank.DieImmutable(false).DieFeed(corev1.ContainerStatus{Name: name})
		fn(d)
		r.InitContainerStatuses = append(r.InitContainerStatuses, d.DieRelease())
	})
}

func (d *podStatusDie) ContainerStatusDie(name string, fn func(d ContainerStatusDie)) PodStatusDie {
	return d.DieStamp(func(r *corev1.PodStatus) {
		for i := range r.ContainerStatuses {
			if name == r.ContainerStatuses[i].Name {
				d := ContainerStatusBlank.DieImmutable(false).DieFeed(r.ContainerStatuses[i])
				fn(d)
				r.ContainerStatuses[i] = d.DieRelease()
				return
			}
		}

		d := ContainerStatusBlank.DieImmutable(false).DieFeed(corev1.ContainerStatus{Name: name})
		fn(d)
		r.ContainerStatuses = append(r.ContainerStatuses, d.DieRelease())
	})
}

func (d *podStatusDie) EphemeralContainerStatusDie(name string, fn func(d ContainerStatusDie)) PodStatusDie {
	return d.DieStamp(func(r *corev1.PodStatus) {
		for i := range r.EphemeralContainerStatuses {
			if name == r.EphemeralContainerStatuses[i].Name {
				d := ContainerStatusBlank.DieImmutable(false).DieFeed(r.EphemeralContainerStatuses[i])
				fn(d)
				r.EphemeralContainerStatuses[i] = d.DieRelease()
				return
			}
		}

		d := ContainerStatusBlank.DieImmutable(false).DieFeed(corev1.ContainerStatus{Name: name})
		fn(d)
		r.EphemeralContainerStatuses = append(r.EphemeralContainerStatuses, d.DieRelease())
	})
}

// +die
type _ = corev1.PodTemplateSpec

type podTemplateSpecDieExtension interface {
	MetadataDie(fn func(d diemetav1.ObjectMetaDie)) PodTemplateSpecDie
	SpecDie(fn func(d PodSpecDie)) PodTemplateSpecDie
}

func (d *podTemplateSpecDie) MetadataDie(fn func(d diemetav1.ObjectMetaDie)) PodTemplateSpecDie {
	return d.DieStamp(func(r *corev1.PodTemplateSpec) {
		d := diemetav1.ObjectMetaBlank.DieImmutable(false).DieFeed(r.ObjectMeta)
		fn(d)
		r.ObjectMeta = d.DieRelease()
	})
}

func (d *podTemplateSpecDie) SpecDie(fn func(d PodSpecDie)) PodTemplateSpecDie {
	return d.DieStamp(func(r *corev1.PodTemplateSpec) {
		d := PodSpecBlank.DieImmutable(false).DieFeed(r.Spec)
		fn(d)
		r.Spec = d.DieRelease()
	})
}
