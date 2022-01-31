/*
Copyright 2022 the original author or authors.

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

package v1beta1_test

import (
	"testing"

	dieflowcontrolv1beta1 "dies.dev/apis/apiserver/flowcontrol/v1beta1"
	diemetav1 "dies.dev/apis/meta/v1"
	"github.com/google/go-cmp/cmp"
	flowcontrolv1beta1 "k8s.io/api/flowcontrol/v1beta1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

func TestPriorityLevelConfiguration(t *testing.T) {
	tests := []struct {
		name     string
		die      *dieflowcontrolv1beta1.PriorityLevelConfigurationDie
		expected flowcontrolv1beta1.PriorityLevelConfiguration
	}{
		{
			name:     "empty",
			die:      dieflowcontrolv1beta1.PriorityLevelConfigurationBlank,
			expected: flowcontrolv1beta1.PriorityLevelConfiguration{},
		},
		{
			name: "object metadata",
			die: dieflowcontrolv1beta1.PriorityLevelConfigurationBlank.
				MetadataDie(func(d *diemetav1.ObjectMetaDie) {
					d.Name("my-name")
				}),
			expected: flowcontrolv1beta1.PriorityLevelConfiguration{
				ObjectMeta: metav1.ObjectMeta{
					Name: "my-name",
				},
			},
		},
	}

	for _, c := range tests {
		t.Run(c.name, func(t *testing.T) {
			actual := c.die.DieRelease()
			if diff := cmp.Diff(c.expected, actual); diff != "" {
				t.Errorf("(-expected, +actual): %s", diff)
			}
		})
	}
}
