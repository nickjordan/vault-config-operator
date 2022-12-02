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
	vaultutils "github.com/redhat-cop/vault-config-operator/api/v1alpha1/utils"
	"k8s.io/apimachinery/pkg/runtime"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/controller/controllerutil"
	logf "sigs.k8s.io/controller-runtime/pkg/log"
	"sigs.k8s.io/controller-runtime/pkg/webhook"
)

// log is for logging in this package.
var ldapauthenginegrouplog = logf.Log.WithName("ldapauthenginegroup-resource")

func (r *LDAPAuthEngineGroup) SetupWebhookWithManager(mgr ctrl.Manager) error {
	return ctrl.NewWebhookManagedBy(mgr).
		For(r).
		Complete()
}

//+kubebuilder:webhook:path=/mutate-redhatcop-redhat-io-v1alpha1-ldapauthenginegroup,mutating=true,failurePolicy=fail,sideEffects=None,groups=redhatcop.redhat.io,resources=ldapauthenginegroups,verbs=create,versions=v1alpha1,name=mldapauthenginegroup.kb.io,admissionReviewVersions=v1

var _ webhook.Defaulter = &LDAPAuthEngineGroup{}

// Default implements webhook.Defaulter so a webhook will be registered for the type
func (r *LDAPAuthEngineGroup) Default() {
	ldapauthenginegrouplog.Info("default", "name", r.Name)

	if !controllerutil.ContainsFinalizer(r, vaultutils.GetFinalizer(r)) {
		controllerutil.AddFinalizer(r, vaultutils.GetFinalizer(r))
	}
}

//+kubebuilder:webhook:path=/validate-redhatcop-redhat-io-v1alpha1-ldapauthenginegroup,mutating=false,failurePolicy=fail,sideEffects=None,groups=redhatcop.redhat.io,resources=ldapauthenginegroups,verbs=update,versions=v1alpha1,name=vldapauthenginegroup.kb.io,admissionReviewVersions=v1

var _ webhook.Validator = &LDAPAuthEngineGroup{}

// ValidateCreate implements webhook.Validator so a webhook will be registered for the type
func (r *LDAPAuthEngineGroup) ValidateCreate() error {
	ldapauthenginegrouplog.Info("validate create", "name", r.Name)

	return nil
}

// ValidateUpdate implements webhook.Validator so a webhook will be registered for the type
func (r *LDAPAuthEngineGroup) ValidateUpdate(old runtime.Object) error {
	ldapauthenginegrouplog.Info("validate update", "name", r.Name)

	return nil
}

// ValidateDelete implements webhook.Validator so a webhook will be registered for the type
func (r *LDAPAuthEngineGroup) ValidateDelete() error {
	ldapauthenginegrouplog.Info("validate delete", "name", r.Name)

	return nil
}
