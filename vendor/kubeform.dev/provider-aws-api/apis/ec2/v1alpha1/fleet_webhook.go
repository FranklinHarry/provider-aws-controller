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
	"strings"

	base "kubeform.dev/apimachinery/api/v1alpha1"
	"kubeform.dev/apimachinery/pkg/util"

	jsoniter "github.com/json-iterator/go"
	"k8s.io/apimachinery/pkg/runtime"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/webhook"
)

func (r *Fleet) SetupWebhookWithManager(mgr ctrl.Manager) error {
	return ctrl.NewWebhookManagedBy(mgr).
		For(r).
		Complete()
}

//+kubebuilder:webhook:verbs=create;update;delete,path=/validate-ec2-aws-kubeform-com-v1alpha1-fleet,mutating=false,failurePolicy=fail,groups=ec2.aws.kubeform.com,resources=fleet,versions=v1alpha1,name=fleet.ec2.aws.kubeform.io,sideEffects=None,admissionReviewVersions=v1

var _ webhook.Validator = &Fleet{}

var fleetForceNewList = map[string]bool{
	"/launch_template_config/*/launch_template_specification/*/launch_template_id":       true,
	"/launch_template_config/*/launch_template_specification/*/launch_template_name":     true,
	"/launch_template_config/*/launch_template_specification/*/version":                  true,
	"/launch_template_config/*/override/*/availability_zone":                             true,
	"/launch_template_config/*/override/*/instance_type":                                 true,
	"/launch_template_config/*/override/*/max_price":                                     true,
	"/launch_template_config/*/override/*/priority":                                      true,
	"/launch_template_config/*/override/*/subnet_id":                                     true,
	"/launch_template_config/*/override/*/weighted_capacity":                             true,
	"/on_demand_options/*/allocation_strategy":                                           true,
	"/replace_unhealthy_instances":                                                       true,
	"/spot_options/*/allocation_strategy":                                                true,
	"/spot_options/*/instance_interruption_behavior":                                     true,
	"/spot_options/*/instance_pools_to_use_count":                                        true,
	"/spot_options/*/maintenance_strategies/*/capacity_rebalance/*/replacement_strategy": true,
	"/target_capacity_specification/*/default_target_capacity_type":                      true,
	"/target_capacity_specification/*/on_demand_target_capacity":                         true,
	"/target_capacity_specification/*/spot_target_capacity":                              true,
	"/terminate_instances_with_expiration":                                               true,
	"/type":                                                                              true,
}

// ValidateCreate implements webhook.Validator so a webhook will be registered for the type
func (r *Fleet) ValidateCreate() error {
	return nil
}

// ValidateUpdate implements webhook.Validator so a webhook will be registered for the type
func (r *Fleet) ValidateUpdate(old runtime.Object) error {
	if r.Spec.Resource.ID == "" {
		return nil
	}
	newObj := r.Spec.Resource
	res := old.(*Fleet)
	oldObj := res.Spec.Resource

	jsnitr := jsoniter.Config{
		EscapeHTML:             true,
		SortMapKeys:            true,
		TagKey:                 "tf",
		ValidateJsonRawMessage: true,
		TypeEncoders:           GetEncoder(),
		TypeDecoders:           GetDecoder(),
	}.Froze()

	byteNew, err := jsnitr.Marshal(newObj)
	if err != nil {
		return err
	}
	tempNew := make(map[string]interface{})
	err = jsnitr.Unmarshal(byteNew, &tempNew)
	if err != nil {
		return err
	}

	byteOld, err := jsnitr.Marshal(oldObj)
	if err != nil {
		return err
	}
	tempOld := make(map[string]interface{})
	err = jsnitr.Unmarshal(byteOld, &tempOld)
	if err != nil {
		return err
	}

	for key := range fleetForceNewList {
		keySplit := strings.Split(key, "/*")
		length := len(keySplit)
		checkIfAnyDif := false
		util.CheckIfAnyDifference("", keySplit, 0, length, &checkIfAnyDif, tempOld, tempOld, tempNew)
		util.CheckIfAnyDifference("", keySplit, 0, length, &checkIfAnyDif, tempNew, tempOld, tempNew)

		if checkIfAnyDif && r.Spec.UpdatePolicy == base.UpdatePolicyDoNotDestroy {
			return fmt.Errorf(`fleet "%v/%v" immutable field can't be updated. To update, change spec.updatePolicy to Destroy`, r.Namespace, r.Name)
		}
	}
	return nil
}

// ValidateDelete implements webhook.Validator so a webhook will be registered for the type
func (r *Fleet) ValidateDelete() error {
	if r.Spec.TerminationPolicy == base.TerminationPolicyDoNotTerminate {
		return fmt.Errorf(`fleet "%v/%v" can't be terminated. To delete, change spec.terminationPolicy to Delete`, r.Namespace, r.Name)
	}
	return nil
}
