/*
Copyright AppsCode Inc. and Contributors

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

// Code generated by Kubeform. DO NOT EDIT.

package v1alpha1

import (
	"fmt"

	base "kubeform.dev/apimachinery/api/v1alpha1"

	"k8s.io/apimachinery/pkg/runtime"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/webhook"
)

func (r *ConfigurationRecorderStatus_) SetupWebhookWithManager(mgr ctrl.Manager) error {
	return ctrl.NewWebhookManagedBy(mgr).
		For(r).
		Complete()
}

//+kubebuilder:webhook:verbs=create;update;delete,path=/validate-config-aws-kubeform-com-v1alpha1-configurationrecorderstatus_,mutating=false,failurePolicy=fail,groups=config.aws.kubeform.com,resources=configurationrecorderstatus_configurationrecorderstatuses,versions=v1alpha1,name=configurationrecorderstatus_.config.aws.kubeform.io,sideEffects=None,admissionReviewVersions=v1

var _ webhook.Validator = &ConfigurationRecorderStatus_{}

// ValidateCreate implements webhook.Validator so a webhook will be registered for the type
func (r *ConfigurationRecorderStatus_) ValidateCreate() error {
	return nil
}

// ValidateUpdate implements webhook.Validator so a webhook will be registered for the type
func (r *ConfigurationRecorderStatus_) ValidateUpdate(old runtime.Object) error {
	return nil
}

// ValidateDelete implements webhook.Validator so a webhook will be registered for the type
func (r *ConfigurationRecorderStatus_) ValidateDelete() error {
	if r.Spec.TerminationPolicy == base.TerminationPolicyDoNotTerminate {
		return fmt.Errorf(`configurationrecorderstatus_ "%v/%v" can't be terminated. To delete, change spec.terminationPolicy to Delete`, r.Namespace, r.Name)
	}
	return nil
}
