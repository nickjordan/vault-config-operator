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
	"errors"
	"reflect"

	vaultutils "github.com/redhat-cop/vault-config-operator/api/v1alpha1/utils"
	"k8s.io/apimachinery/pkg/runtime"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/controller/controllerutil"
	logf "sigs.k8s.io/controller-runtime/pkg/log"
	"sigs.k8s.io/controller-runtime/pkg/webhook"
)

// log is for logging in this package.
var secretenginemountlog = logf.Log.WithName("secretenginemount-resource")

func (r *SecretEngineMount) SetupWebhookWithManager(mgr ctrl.Manager) error {
	return ctrl.NewWebhookManagedBy(mgr).
		For(r).
		Complete()
}

// EDIT THIS FILE!  THIS IS SCAFFOLDING FOR YOU TO OWN!

//+kubebuilder:webhook:path=/mutate-redhatcop-redhat-io-v1alpha1-secretenginemount,mutating=true,failurePolicy=fail,sideEffects=None,groups=redhatcop.redhat.io,resources=secretenginemounts,verbs=create,versions=v1alpha1,name=msecretenginemount.kb.io,admissionReviewVersions={v1,v1beta1}

var _ webhook.Defaulter = &SecretEngineMount{}

// Default implements webhook.Defaulter so a webhook will be registered for the type
func (r *SecretEngineMount) Default() {
	authenginemountlog.Info("default", "name", r.Name)
	if !controllerutil.ContainsFinalizer(r, vaultutils.GetFinalizer(r)) {
		controllerutil.AddFinalizer(r, vaultutils.GetFinalizer(r))
	}
}

// TODO(user): change verbs to "verbs=create;update;delete" if you want to enable deletion validation.
//+kubebuilder:webhook:path=/validate-redhatcop-redhat-io-v1alpha1-secretenginemount,mutating=false,failurePolicy=fail,sideEffects=None,groups=redhatcop.redhat.io,resources=secretenginemounts,verbs=update,versions=v1alpha1,name=vsecretenginemount.kb.io,admissionReviewVersions={v1,v1beta1}

var _ webhook.Validator = &SecretEngineMount{}

// ValidateCreate implements webhook.Validator so a webhook will be registered for the type
func (r *SecretEngineMount) ValidateCreate() error {
	secretenginemountlog.Info("validate create", "name", r.Name)

	// TODO(user): fill in your validation logic upon object creation.
	return nil
}

// ValidateUpdate implements webhook.Validator so a webhook will be registered for the type
func (r *SecretEngineMount) ValidateUpdate(old runtime.Object) error {
	secretenginemountlog.Info("validate update", "name", r.Name)

	// the path cannot be updated
	if r.Spec.Path != old.(*SecretEngineMount).Spec.Path {
		return errors.New("spec.path cannot be updated")
	}
	// only mount config can be modified
	oldMount := old.(*SecretEngineMount).Spec.Mount
	newMount := r.Spec.Mount
	oldMount.Config = MountConfig{}
	newMount.Config = MountConfig{}
	if !reflect.DeepEqual(oldMount, newMount) {
		return errors.New("only .spec.config can be modified")
	}
	return nil
}

// ValidateDelete implements webhook.Validator so a webhook will be registered for the type
func (r *SecretEngineMount) ValidateDelete() error {
	secretenginemountlog.Info("validate delete", "name", r.Name)

	// TODO(user): fill in your validation logic upon object deletion.
	return nil
}
