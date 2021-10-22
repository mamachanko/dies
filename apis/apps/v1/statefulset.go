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
	diecorev1 "dies.dev/apis/core/v1"
	diemetav1 "dies.dev/apis/meta/v1"
	appsv1 "k8s.io/api/apps/v1"
	corev1 "k8s.io/api/core/v1"
)

// +die:object=true
type _ = appsv1.StatefulSet

// +die
type _ = appsv1.StatefulSetSpec

type statefulSetSpecDieExtension interface {
	SelectorDie(fn func(d diemetav1.LabelSelectorDie)) StatefulSetSpecDie
	TemplateDie(fn func(d diecorev1.PodTemplateSpecDie)) StatefulSetSpecDie
	VolumeClaimTemplatesDie(volumeClaimTemplates ...diecorev1.PersistentVolumeClaimDie) StatefulSetSpecDie
	UpdateStrategyDie(fn func(d StatefulSetUpdateStrategyDie)) StatefulSetSpecDie
}

func (d *statefulSetSpecDie) SelectorDie(fn func(d diemetav1.LabelSelectorDie)) StatefulSetSpecDie {
	return d.DieStamp(func(r *appsv1.StatefulSetSpec) {
		d := diemetav1.LabelSelectorBlank.DieImmutable(false).DieFeedPtr(r.Selector)
		fn(d)
		r.Selector = d.DieReleasePtr()
	})
}

func (d *statefulSetSpecDie) TemplateDie(fn func(d diecorev1.PodTemplateSpecDie)) StatefulSetSpecDie {
	return d.DieStamp(func(r *appsv1.StatefulSetSpec) {
		d := diecorev1.PodTemplateSpecBlank.DieImmutable(false).DieFeed(r.Template)
		fn(d)
		r.Template = d.DieRelease()
	})
}

func (d *statefulSetSpecDie) VolumeClaimTemplatesDie(volumeClaimTemplates ...diecorev1.PersistentVolumeClaimDie) StatefulSetSpecDie {
	return d.DieStamp(func(r *appsv1.StatefulSetSpec) {
		r.VolumeClaimTemplates = make([]corev1.PersistentVolumeClaim, len(volumeClaimTemplates))
		for i, v := range volumeClaimTemplates {
			r.VolumeClaimTemplates[i] = v.DieRelease()
		}
	})
}

func (d *statefulSetSpecDie) UpdateStrategyDie(fn func(d StatefulSetUpdateStrategyDie)) StatefulSetSpecDie {
	return d.DieStamp(func(r *appsv1.StatefulSetSpec) {
		d := StatefulSetUpdateStrategyBlank.DieImmutable(false).DieFeed(r.UpdateStrategy)
		fn(d)
		r.UpdateStrategy = d.DieRelease()
	})
}

// +die
type _ = appsv1.StatefulSetUpdateStrategy

type statefulSetUpdateStrategyDieExtension interface {
	OnDelete() StatefulSetUpdateStrategyDie
	RollingUpdateDie(fn func(d RollingUpdateStatefulSetStrategyDie)) StatefulSetUpdateStrategyDie
}

func (d *statefulSetUpdateStrategyDie) OnDelete() StatefulSetUpdateStrategyDie {
	return d.DieStamp(func(r *appsv1.StatefulSetUpdateStrategy) {
		r.Type = appsv1.OnDeleteStatefulSetStrategyType
		r.RollingUpdate = nil
	})
}

func (d *statefulSetUpdateStrategyDie) RollingUpdateDie(fn func(d RollingUpdateStatefulSetStrategyDie)) StatefulSetUpdateStrategyDie {
	return d.DieStamp(func(r *appsv1.StatefulSetUpdateStrategy) {
		r.Type = appsv1.RollingUpdateStatefulSetStrategyType
		d := RollingUpdateStatefulSetStrategyBlank.DieImmutable(false).DieFeedPtr(r.RollingUpdate)
		fn(d)
		r.RollingUpdate = d.DieReleasePtr()
	})
}

// +die
type _ = appsv1.RollingUpdateStatefulSetStrategy

// +die
type _ = appsv1.StatefulSetStatus

type statefulSetStatusDieExtension interface {
	ConditionsDie(conditions ...diemetav1.ConditionDie) StatefulSetStatusDie
}

func (d *statefulSetStatusDie) ConditionsDie(conditions ...diemetav1.ConditionDie) StatefulSetStatusDie {
	return d.DieStamp(func(r *appsv1.StatefulSetStatus) {
		r.Conditions = make([]appsv1.StatefulSetCondition, len(conditions))
		for i := range conditions {
			c := conditions[i].DieRelease()
			r.Conditions[i] = appsv1.StatefulSetCondition{
				Type:               appsv1.StatefulSetConditionType(c.Type),
				Status:             corev1.ConditionStatus(c.Status),
				Reason:             c.Reason,
				Message:            c.Message,
				LastTransitionTime: c.LastTransitionTime,
			}
		}
	})
}
