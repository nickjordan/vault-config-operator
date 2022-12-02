/*
Copyright 2021.

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

package v1alpha1

import (
	"context"
	"reflect"

	"github.com/redhat-cop/operator-utils/pkg/util/apis"
	vaultutils "github.com/redhat-cop/vault-config-operator/api/v1alpha1/utils"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

// QuaySecretEngineStaticRoleSpec defines the desired state of QuaySecretEngineStaticRole
type QuaySecretEngineStaticRoleSpec struct {
	// Authentication is the kube auth configuration to be used to execute this request
	// +kubebuilder:validation:Required
	Authentication vaultutils.KubeAuthConfiguration `json:"authentication,omitempty"`

	// Path at which to make the configuration.
	// The final path will be {[spec.authentication.namespace]}/{spec.path}/static-roles/{metadata.name}.
	// The authentication role must have the following capabilities = [ "create", "read", "update", "delete"] on that path.
	// +kubebuilder:validation:Required
	Path vaultutils.Path `json:"path,omitempty"`

	QuayBaseRole `json:",inline"`
}

var _ vaultutils.VaultObject = &QuaySecretEngineStaticRole{}

func (q *QuaySecretEngineStaticRole) GetPath() string {
	return string(q.Spec.Path) + "/" + "static-roles" + "/" + q.Name
}
func (q *QuaySecretEngineStaticRole) GetPayload() map[string]interface{} {
	return q.Spec.toMap()
}
func (q *QuaySecretEngineStaticRole) IsEquivalentToDesiredState(payload map[string]interface{}) bool {
	desiredState := q.Spec.QuayBaseRole.toMap()
	return reflect.DeepEqual(desiredState, payload)
}

func (q *QuaySecretEngineStaticRole) IsInitialized() bool {
	return true
}

func (q *QuaySecretEngineStaticRole) PrepareInternalValues(context context.Context, object client.Object) error {
	return nil
}

func (q *QuaySecretEngineStaticRole) IsValid() (bool, error) {
	return true, nil
}

func (r *QuayBaseRole) toMap() map[string]interface{} {
	payload := map[string]interface{}{}
	payload["namespace_type"] = r.NamespaceType
	payload["namespace_name"] = r.NamespaceName
	payload["create_repositories"] = r.CreateRepositories
	if r.DefaultPermission != nil {
		payload["default_permission"] = r.DefaultPermission
	}
	if r.Teams != nil {
		setMapJson(payload, "teams", r.Teams)

	}
	if r.Repositories != nil {
		setMapJson(payload, "repositories", r.Repositories)
	}
	return payload
}

// QuaySecretEngineStaticRoleStatus defines the observed state of QuaySecretEngineStaticRole
type QuaySecretEngineStaticRoleStatus struct {
	// +patchMergeKey=type
	// +patchStrategy=merge
	// +listType=map
	// +listMapKey=type
	Conditions []metav1.Condition `json:"conditions,omitempty" patchStrategy:"merge" patchMergeKey:"type"`
}

var _ apis.ConditionsAware = &QuaySecretEngineStaticRole{}

func (q *QuaySecretEngineStaticRole) GetConditions() []metav1.Condition {
	return q.Status.Conditions
}

func (q *QuaySecretEngineStaticRole) SetConditions(conditions []metav1.Condition) {
	q.Status.Conditions = conditions
}

//+kubebuilder:object:root=true
//+kubebuilder:subresource:status

// QuaySecretEngineStaticRole is the Schema for the quaysecretenginestaticroles API
type QuaySecretEngineStaticRole struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   QuaySecretEngineStaticRoleSpec   `json:"spec,omitempty"`
	Status QuaySecretEngineStaticRoleStatus `json:"status,omitempty"`
}

//+kubebuilder:object:root=true

// QuaySecretEngineStaticRoleList contains a list of QuaySecretEngineStaticRole
type QuaySecretEngineStaticRoleList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []QuaySecretEngineStaticRole `json:"items"`
}

func init() {
	SchemeBuilder.Register(&QuaySecretEngineStaticRole{}, &QuaySecretEngineStaticRoleList{})
}

func (d *QuaySecretEngineStaticRole) GetKubeAuthConfiguration() *vaultutils.KubeAuthConfiguration {
	return &d.Spec.Authentication
}
