//go:build integration
// +build integration

package controllers

import (
	"context"
	"reflect"
	"time"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	redhatcopv1alpha1 "github.com/redhat-cop/vault-config-operator/api/v1alpha1"

	corev1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/types"
)

//TODO: Example: https://github.com/kubernetes-sigs/kubebuilder/blob/master/docs/book/src/cronjob-tutorial/testdata/project/controllers/cronjob_controller_test.go
// Define utility constants for object names and testing timeouts/durations and intervals.

var _ = Describe("VaultSecret controller", func() {
	timeout := time.Second * 120
	interval := time.Second * 2
	Context("When creating randomsecret VaultSecret", func() {
		It("Should be Successful when created", func() {
			By("By creating a new VaultSecret")
			ctx := context.Background()

			instance, err := decoder.GetVaultSecretInstance("../test/vaultsecret/vaultsecret-randomsecret.yaml")
			Expect(err).To(BeNil())
			instance.Namespace = vaultTestNamespaceName
			Expect(k8sClient.Create(ctx, instance)).Should(Succeed())

			lookupKey := types.NamespacedName{Name: instance.Name, Namespace: instance.Namespace}
			created := &redhatcopv1alpha1.VaultSecret{}

			// We'll need to retry getting this newly created VaultSecret, given that creation may not immediately happen.
			Eventually(func() bool {
				err := k8sClient.Get(ctx, lookupKey, created)
				if err != nil {
					return false
				}

				for _, condition := range created.Status.Conditions {
					if condition.Type == "ReconcileSuccess" {
						return true
					}
				}

				return false
			}, timeout, interval).Should(BeTrue())

			By("By checking the Secret Exists")

			lookupKey = types.NamespacedName{Name: instance.Spec.TemplatizedK8sSecret.Name, Namespace: instance.Namespace}
			secret := &corev1.Secret{}

			Eventually(func() bool {
				err := k8sClient.Get(ctx, lookupKey, secret)
				if err != nil {
					return false
				}
				return true
			}, timeout, interval).Should(BeTrue())

			kind := reflect.TypeOf(redhatcopv1alpha1.VaultSecret{}).Name()
			Expect(secret.GetObjectMeta().GetOwnerReferences()[0].Kind).Should(Equal(kind))

		})
	})
})
