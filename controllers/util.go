/*
Copyright AppsCode Inc. and Contributors

Licensed under the AppsCode Free Trial License 1.0.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    https://github.com/appscode/licenses/raw/1.0.0/AppsCode-Free-Trial-1.0.0.md

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

// Code generated by Kubeform. DO NOT EDIT.

package controllers

import (
	"bytes"
	"context"
	"encoding/gob"
	"encoding/json"
	"fmt"
	"reflect"
	"strings"

	"github.com/fatih/structs"
	"github.com/hashicorp/go-cty/cty"
	"github.com/hashicorp/go-cty/cty/msgpack"
	"github.com/hashicorp/terraform-plugin-go/tfprotov5"
	tfschema "github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
	"github.com/imdario/mergo"
	jsoniter "github.com/json-iterator/go"
	errors2 "github.com/pkg/errors"
	corev1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/api/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/apimachinery/pkg/types"
	kmapi "kmodules.xyz/client-go/api/v1"
	"kmodules.xyz/client-go/meta"
	aws "kubeform.dev/provider-aws-api/api/provider"
	"sigs.k8s.io/cli-utils/pkg/kstatus/status"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

const (
	KFCFinalizer       = "aws.kubeform.com/finalizer"
	UnknownIdValue     = "4856ec62-a372-11eb-bcbc-0242ac130002"
	UpdateNotSupported = "doesn't support update"
)

func StartProcess(rClient client.Client, provider *tfschema.Provider, ctx context.Context, res *tfschema.Resource, gv schema.GroupVersion, unstructuredObj *unstructured.Unstructured, tName string, jsonit jsoniter.API) error {
	err := initialUpdateStatus(rClient, ctx, gv, unstructuredObj, nil, true)
	if err != nil {
		return err
	}

	err = reconcile(rClient, provider, ctx, res, gv, unstructuredObj, tName, jsonit)
	if err != nil {
		err2 := initialUpdateStatus(rClient, ctx, gv, unstructuredObj, err, false)
		if err2 != nil {
			return err2
		}
		return err
	}

	return finalUpdateStatus(rClient, ctx, gv, unstructuredObj)
}

func reconcile(rClient client.Client, provider *tfschema.Provider, ctx context.Context, res *tfschema.Resource, gv schema.GroupVersion, unstructuredObj *unstructured.Unstructured, tName string, jsonit jsoniter.API) error {
	server := tfschema.NewGRPCProviderServer(provider)

	// Get RawSpec (including sensitive data)
	rawSpec, err := getSpecWithSensitiveData(gv, rClient, ctx, unstructuredObj, jsonit)
	if err != nil {
		return err
	}

	// Get RawStatus (including sensitive data)
	rawStatus, err := getStatusWithSensitiveData(gv, rClient, ctx, unstructuredObj, jsonit)
	if err != nil {
		return err
	}

	// Get object ID
	_, found, err := unstructured.NestedString(unstructuredObj.Object, "spec", "resource", "id")
	if err != nil {
		return err
	}

	// Set ProviderMeta
	err = setProviderMeta(rClient, provider, ctx, unstructuredObj, server, jsonit)
	if err != nil {
		return err
	}

	// validation check
	if rawSpec["id"] == nil {
		rawSpec["id"] = UnknownIdValue
	}
	rawSpecCty := HCL2ValueFromConfigValue(rawSpec)
	initialState, err := msgpack.Marshal(rawSpecCty, res.CoreConfigSchema().ImpliedType())
	if err != nil {
		return err
	}
	req := &tfprotov5.ValidateResourceTypeConfigRequest{
		TypeName: tName,
		Config: &tfprotov5.DynamicValue{
			MsgPack: initialState,
		},
	}
	valid, err := server.ValidateResourceTypeConfig(ctx, req)
	if err != nil {
		return err
	}
	if len(valid.Diagnostics) > 0 {
		err = diagToError(valid.Diagnostics)
		if err != nil {
			return err
		}
	}

	if hasFinalizer(unstructuredObj.GetFinalizers(), KFCFinalizer) {
		if unstructuredObj.GetDeletionTimestamp() != nil {
			err := updateStatus(rClient, ctx, unstructuredObj, status.TerminatingStatus)
			if err != nil {
				return err
			}
			// if not found then also delete
			err = destroyTheObject(rawStatus, res, server, tName)
			if err != nil && !strings.Contains(err.Error(), "[404] Not found") {
				return err
			}
			return removeFinalizer(ctx, rClient, unstructuredObj, KFCFinalizer)
		}
	} else {
		err := addFinalizer(ctx, rClient, unstructuredObj, KFCFinalizer)
		if err != nil {
			return err
		}
	}

	if !found {
		err := updateStatus(rClient, ctx, unstructuredObj, status.InProgressStatus)
		if err != nil {
			return err
		}
		newStateVal, intrfc, err := createTheObject(rawSpec, res, server, tName)
		if err != nil {
			return err
		}
		err = updateStatus(rClient, ctx, unstructuredObj, status.CurrentStatus)
		if err != nil {
			return err
		}
		err = updateStateField(rClient, ctx, intrfc.Raw, gv, unstructuredObj, jsonit)
		if err != nil {
			return err
		}

		// set the id value in unstructuredObj object
		err = unstructured.SetNestedField(unstructuredObj.Object, newStateVal.GetAttr("id").AsString(), "spec", "resource", "id")
		if err != nil {
			return err
		}

		// apply the update of the object
		if err = rClient.Update(ctx, unstructuredObj); err != nil {
			return err
		}
		return nil
	}

	combineRaw, err := getCombineRawAndDeepCopyRawStatus(rawStatus, rawSpec)
	if err != nil {
		return err
	}

	requireNew, priorState, planResp, plannedState, err := checkRequireNewOrNot(combineRaw, rawStatus, res, server, tName)
	if err != nil {
		return err
	}

	if requireNew { // Resources is needed to be destroyed because one of the field needs the resource to be replaced
		err = updateStatus(rClient, ctx, unstructuredObj, status.TerminatingStatus)
		if err != nil {
			return err
		}
		err := destroyTheObject(rawStatus, res, server, tName)
		if err != nil {
			return err
		}

		err = updateStatus(rClient, ctx, unstructuredObj, status.InProgressStatus)
		if err != nil {
			return err
		}
		newStateVal, intrfc, err := createTheObject(rawSpec, res, server, tName)
		if err != nil {
			return err
		}
		err = updateStatus(rClient, ctx, unstructuredObj, status.CurrentStatus)
		if err != nil {
			return err
		}
		err = updateStateField(rClient, ctx, intrfc.Raw, gv, unstructuredObj, jsonit)
		if err != nil {
			return err
		}

		// set the id value in unstructuredObj object
		err = unstructured.SetNestedField(unstructuredObj.Object, newStateVal.GetAttr("id").AsString(), "spec", "resource", "id")
		if err != nil {
			return err
		}

		// apply the update of the object
		if err = rClient.Update(ctx, unstructuredObj); err != nil {
			return err
		}
		return nil
	}

	newStateVal, updated, err := updateTheObject(priorState, plannedState, planResp, server, res, tName)
	if err != nil {
		return err
	}

	if updated {
		//set the id value in unstructuredObj object
		err = unstructured.SetNestedField(unstructuredObj.Object, newStateVal.GetAttr("id").AsString(), "spec", "resource", "id")
		if err != nil {
			return err
		}

		// apply the update of the object
		if err = rClient.Update(ctx, unstructuredObj); err != nil {
			return err
		}

		intrfc := terraform.NewResourceConfigShimmed(newStateVal, res.CoreConfigSchema())

		err = updateStateField(rClient, ctx, intrfc.Raw, gv, unstructuredObj, jsonit)
		if err != nil {
			return err
		}
	}

	return nil
}

func initialUpdateStatus(rClient client.Client, ctx context.Context, gv schema.GroupVersion, obj *unstructured.Unstructured, er error, flag bool) error {
	objGen, _, err := unstructured.NestedInt64(obj.Object, "metadata", "generation")
	if err != nil {
		return err
	}

	var newCondi []kmapi.Condition
	phase := status.InProgressStatus
	if flag {
		newCondi = kmapi.SetCondition(newCondi, kmapi.NewCondition("Reconciling", "Kubeform is currently reconciling "+obj.GetKind()+" resource", objGen))
	} else {
		newCondi = kmapi.SetCondition(newCondi, kmapi.NewCondition("Stalled", er.Error(), objGen))
		phase = status.FailedStatus
	}

	err = setNestedFieldNoCopy(obj.Object, newCondi, "status", "conditions")
	if err = rClient.Status().Update(ctx, obj); err != nil {
		return err
	}

	return updateStatus(rClient, ctx, obj, phase)
}

func finalUpdateStatus(rClient client.Client, ctx context.Context, gv schema.GroupVersion, obj *unstructured.Unstructured) error {
	var newCondi []kmapi.Condition
	err := setNestedFieldNoCopy(obj.Object, newCondi, "status", "conditions")
	if err = rClient.Status().Update(ctx, obj); err != nil {
		return err
	}
	err = updateStatus(rClient, ctx, obj, status.CurrentStatus)
	if err != nil {
		return err
	}
	return nil
}

func updateStatus(rClient client.Client, ctx context.Context, obj *unstructured.Unstructured, phase status.Status) error {
	obsGen, _, err := unstructured.NestedInt64(obj.Object, "metadata", "generation")
	if err != nil {
		return err
	}
	err = unstructured.SetNestedField(obj.Object, obsGen, "status", "observedGeneration")
	if err != nil {
		return err
	}
	err = setNestedFieldNoCopy(obj.Object, phase, "status", "phase")
	if err != nil {
		return err
	}

	// apply the status update of the object
	if err = rClient.Status().Update(ctx, obj); err != nil {
		return err
	}
	return nil
}

func setNestedFieldNoCopy(obj map[string]interface{}, value interface{}, fields ...string) error {
	m := obj

	for i, field := range fields[:len(fields)-1] {
		if val, ok := m[field]; ok {
			if valMap, ok := val.(map[string]interface{}); ok {
				m = valMap
			} else {
				return fmt.Errorf("value cannot be set because %v is not a map[string]interface{}", jsonPath(fields[:i+1]))
			}
		} else {
			newVal := make(map[string]interface{})
			m[field] = newVal
			m = newVal
		}
	}
	m[fields[len(fields)-1]] = value
	return nil
}

func jsonPath(fields []string) string {
	return "." + strings.Join(fields, ".")
}

func diagToError(d []*tfprotov5.Diagnostic) error {
	var err error
	var flag bool
	for idx, key := range d {
		if key.Severity.String() == "WARNING" || key.Summary == "Invalid or unknown key" || key.Summary == UpdateNotSupported {
			continue
		}
		if !flag {
			err = errors2.New("")
			flag = true
		}
		err = errors2.Wrapf(err, "%s %d: %s", "Error", idx, key.Summary)
	}
	return err
}

func getCombineRawAndDeepCopyRawStatus(rawStatus map[string]interface{}, rawSpec map[string]interface{}) (map[string]interface{}, error) {
	var buf bytes.Buffer
	enc := gob.NewEncoder(&buf)
	dec := gob.NewDecoder(&buf)
	err := enc.Encode(rawSpec)
	if err != nil {
		return nil, err
	}

	var copyrawSpec map[string]interface{}
	err = dec.Decode(&copyrawSpec)
	if err != nil {
		return nil, err
	}
	if err := mergo.Merge(&copyrawSpec, rawStatus); err != nil {
		return nil, err
	}

	return copyrawSpec, nil
}

func updateTheObject(priorState []byte, plannedState []byte, planResp *tfprotov5.PlanResourceChangeResponse, server *tfschema.GRPCProviderServer, res *tfschema.Resource, tName string) (cty.Value, bool, error) {
	applyReq := &tfprotov5.ApplyResourceChangeRequest{
		TypeName: tName,
		PriorState: &tfprotov5.DynamicValue{
			MsgPack: priorState,
		},
		PlannedState: &tfprotov5.DynamicValue{
			MsgPack: planResp.PlannedState.MsgPack,
		},
		Config: &tfprotov5.DynamicValue{
			MsgPack: plannedState,
		},
		PlannedPrivate: planResp.PlannedPrivate,
	}

	applyResp, err := server.ApplyResourceChange(context.Background(), applyReq)
	if err != nil {
		return cty.Value{}, false, err
	}

	if !isUpdateSupported(applyResp.Diagnostics) {
		return cty.Value{}, false, err
	}

	err = diagToError(applyResp.Diagnostics)
	if err != nil {
		return cty.Value{}, false, err
	}

	schma := res.CoreConfigSchema()
	newStateVal, err := msgpack.Unmarshal(applyResp.NewState.MsgPack, schma.ImpliedType())
	if err != nil {
		return cty.Value{}, false, err
	}

	return newStateVal, true, nil
}

func checkRequireNewOrNot(combineRaw map[string]interface{}, copyrawStatus map[string]interface{}, res *tfschema.Resource, server *tfschema.GRPCProviderServer, tName string) (bool, []byte, *tfprotov5.PlanResourceChangeResponse, []byte, error) {
	stateVal := HCL2ValueFromConfigValue(copyrawStatus)
	proposedPlanVal := HCL2ValueFromConfigValue(combineRaw)

	schma := res.CoreConfigSchema()
	priorState, err := msgpack.Marshal(stateVal, schma.ImpliedType())
	if err != nil {
		return false, nil, nil, nil, err
	}

	plannedState, err := msgpack.Marshal(proposedPlanVal, schma.ImpliedType())
	if err != nil {
		return false, nil, nil, nil, err
	}

	planReq := &tfprotov5.PlanResourceChangeRequest{
		TypeName: tName,
		PriorState: &tfprotov5.DynamicValue{
			MsgPack: priorState,
		},
		ProposedNewState: &tfprotov5.DynamicValue{
			MsgPack: plannedState,
		},
	}

	planResp, err := server.PlanResourceChange(context.Background(), planReq)
	if err != nil {
		return false, nil, nil, nil, err
	}
	if len(planResp.Diagnostics) > 0 {
		return false, nil, nil, nil, diagToError(planResp.Diagnostics)
	}

	var requireNew bool
	if len(planResp.RequiresReplace) > 0 {
		requireNew = true
	}

	return requireNew, priorState, planResp, plannedState, nil
}

func createTheObject(rawSpec map[string]interface{}, res *tfschema.Resource, server *tfschema.GRPCProviderServer, tName string) (cty.Value, *terraform.ResourceConfig, error) {
	rawSpec["id"] = UnknownIdValue
	stateVal := HCL2ValueFromConfigValue(rawSpec)

	schma := res.CoreConfigSchema()
	priorState, err := msgpack.Marshal(cty.NullVal(schma.ImpliedType()), schma.ImpliedType())
	if err != nil {
		return cty.Value{}, nil, err
	}
	plannedState, err := msgpack.Marshal(stateVal, schma.ImpliedType())
	if err != nil {
		return cty.Value{}, nil, err
	}

	planReq := &tfprotov5.PlanResourceChangeRequest{
		TypeName: tName,
		PriorState: &tfprotov5.DynamicValue{
			MsgPack: priorState,
		},
		ProposedNewState: &tfprotov5.DynamicValue{
			MsgPack: plannedState,
		},
	}

	planResp, err := server.PlanResourceChange(context.Background(), planReq)
	if err != nil {
		return cty.Value{}, nil, err
	}

	if len(planResp.Diagnostics) > 0 {
		return cty.Value{}, nil, diagToError(planResp.Diagnostics)
	}

	applyReq := &tfprotov5.ApplyResourceChangeRequest{
		TypeName: tName,
		PriorState: &tfprotov5.DynamicValue{
			MsgPack: priorState,
		},
		PlannedState: &tfprotov5.DynamicValue{
			MsgPack: planResp.PlannedState.MsgPack,
		},
		Config: &tfprotov5.DynamicValue{
			MsgPack: plannedState,
		},
		PlannedPrivate: planResp.PlannedPrivate,
	}

	applyResp, err := server.ApplyResourceChange(context.Background(), applyReq)
	if err != nil {
		return cty.Value{}, nil, err
	}
	if len(applyResp.Diagnostics) > 0 {
		return cty.Value{}, nil, diagToError(applyResp.Diagnostics)
	}

	newStateVal, err := msgpack.Unmarshal(applyResp.NewState.MsgPack, schma.ImpliedType())
	if err != nil {
		return cty.Value{}, nil, err
	}
	intrfc := terraform.NewResourceConfigShimmed(newStateVal, res.CoreConfigSchema())

	return newStateVal, intrfc, nil
}

func destroyTheObject(rawStatus map[string]interface{}, res *tfschema.Resource, server *tfschema.GRPCProviderServer, tName string) error {
	stateVal := HCL2ValueFromConfigValue(rawStatus)
	schma := res.CoreConfigSchema()
	priorState, err := msgpack.Marshal(stateVal, schma.ImpliedType())
	if err != nil {
		return err
	}

	plannedState, err := msgpack.Marshal(cty.NullVal(schma.ImpliedType()), schma.ImpliedType())
	if err != nil {
		return err
	}

	planReq := &tfprotov5.PlanResourceChangeRequest{
		TypeName: tName,
		PriorState: &tfprotov5.DynamicValue{
			MsgPack: priorState,
		},
		ProposedNewState: &tfprotov5.DynamicValue{
			MsgPack: plannedState,
		},
	}
	planResp, err := server.PlanResourceChange(context.Background(), planReq)
	if err != nil {
		return err
	}
	if len(planResp.Diagnostics) > 0 {
		return diagToError(planResp.Diagnostics)
	}

	applyReq := &tfprotov5.ApplyResourceChangeRequest{
		TypeName: tName,
		PriorState: &tfprotov5.DynamicValue{
			MsgPack: priorState,
		},
		PlannedState: &tfprotov5.DynamicValue{
			MsgPack: planResp.PlannedState.MsgPack,
		},
		Config: &tfprotov5.DynamicValue{
			MsgPack: plannedState,
		},
		PlannedPrivate: planResp.PlannedPrivate,
	}

	applyResp, err := server.ApplyResourceChange(context.Background(), applyReq)
	if err != nil {
		return err
	}
	if len(applyResp.Diagnostics) > 0 {
		return diagToError(applyResp.Diagnostics)
	}
	return nil
}

func GetJSONItr(typeEncoder map[string]jsoniter.ValEncoder, typeDecoder map[string]jsoniter.ValDecoder) jsoniter.API {
	return jsoniter.Config{
		EscapeHTML:             true,
		SortMapKeys:            true,
		ValidateJsonRawMessage: true,
		TagKey:                 "tf",
		TypeDecoders:           typeDecoder,
		TypeEncoders:           typeEncoder,
	}.Froze()
}

func getStatusWithSensitiveData(gv schema.GroupVersion, rClient client.Client, ctx context.Context, obj *unstructured.Unstructured, jsonit jsoniter.API) (map[string]interface{}, error) {
	data, err := meta.MarshalToJson(obj, gv)
	if err != nil {
		return nil, err
	}

	typedObj, err := meta.UnmarshalFromJSON(data, gv)
	if err != nil {
		return nil, err
	}

	typedStruct := structs.New(typedObj)
	status := reflect.ValueOf(typedStruct.Field("Spec").Field("KubeformOutput").Value())
	statusType := reflect.TypeOf(typedStruct.Field("Spec").Field("KubeformOutput").Value())
	statusValue := reflect.New(statusType)
	statusValue.Elem().Set(status)

	secretRef, _, err := unstructured.NestedFieldNoCopy(obj.Object, "spec", "secretRef")
	if err != nil {
		return nil, err
	}

	secretData := make(map[string]interface{})
	if secretRef != nil {
		secretName := typedStruct.Field("Spec").Field("SecretRef").Field("Name").Value()

		if secretName != nil {
			var secret corev1.Secret
			req := types.NamespacedName{
				Namespace: obj.GetNamespace(),
				Name:      secretName.(string),
			}
			if err := rClient.Get(ctx, req, &secret); err != nil {
				return nil, err
			}

			if _, ok := secret.Data["output"]; ok {
				err = json.Unmarshal(secret.Data["output"], &secretData)
				if err != nil {
					return nil, err
				}
			}
		}
	}

	str, err := jsonit.Marshal(statusValue.Interface())
	if err != nil {
		return nil, err
	}
	rawStatus := make(map[string]interface{})
	err = json.Unmarshal(str, &rawStatus)
	if err != nil {
		return nil, err
	}

	if err := mergo.Merge(&rawStatus, secretData); err != nil {
		return nil, err
	}

	return rawStatus, nil
}

func getSpecWithSensitiveData(gv schema.GroupVersion, rClient client.Client, ctx context.Context, obj *unstructured.Unstructured, jsonit jsoniter.API) (map[string]interface{}, error) {
	data, err := meta.MarshalToJson(obj, gv)
	if err != nil {
		return nil, err
	}

	typedObj, err := meta.UnmarshalFromJSON(data, gv)
	if err != nil {
		return nil, err
	}

	typedStruct := structs.New(typedObj)
	spec := reflect.ValueOf(typedStruct.Field("Spec").Field("Resource").Value())
	specType := reflect.TypeOf(typedStruct.Field("Spec").Field("Resource").Value())
	specValue := reflect.New(specType)
	specValue.Elem().Set(spec)

	secretRef, _, err := unstructured.NestedFieldNoCopy(obj.Object, "spec", "secretRef")
	if err != nil {
		return nil, err
	}

	secretData := make(map[string]interface{})
	if secretRef != nil {
		secretName := typedStruct.Field("Spec").Field("SecretRef").Field("Name").Value()

		if secretName != nil {
			var secret corev1.Secret
			req := types.NamespacedName{
				Namespace: obj.GetNamespace(),
				Name:      secretName.(string),
			}
			if err := rClient.Get(ctx, req, &secret); err != nil {
				return nil, err
			}

			if _, ok := secret.Data["input"]; ok {
				err = json.Unmarshal(secret.Data["input"], &secretData)
				if err != nil {
					return nil, err
				}
			}
		}
	}

	str, err := jsonit.Marshal(specValue.Interface())
	if err != nil {
		return nil, err
	}
	rawSpec := make(map[string]interface{})
	err = json.Unmarshal(str, &rawSpec)
	if err != nil {
		return nil, err
	}

	if err := mergo.Merge(&rawSpec, secretData); err != nil {
		return nil, err
	}

	return rawSpec, nil
}

func getProviderSecretData(rClient client.Client, ctx context.Context, obj *unstructured.Unstructured) (map[string][]byte, error) {
	providerRef, _, err := unstructured.NestedFieldNoCopy(obj.Object, "spec", "providerRef", "name")
	if err != nil {
		return nil, err
	}
	configData := make(map[string][]byte)
	if providerRef != nil {
		var secret corev1.Secret
		req := types.NamespacedName{
			Namespace: obj.GetNamespace(),
			Name:      providerRef.(string),
		}
		if err := rClient.Get(ctx, req, &secret); err != nil {
			return nil, err
		}

		configData = secret.Data
	}
	return configData, nil
}

func hasFinalizer(finalizers []string, finalizer string) bool {
	for _, f := range finalizers {
		if f == finalizer {
			return true
		}
	}
	return false
}

func removeFinalizer(ctx context.Context, rClient client.Client, u *unstructured.Unstructured, finalizer string) error {
	finalizers := u.GetFinalizers()
	for i, v := range finalizers {
		if v == finalizer {
			finalizers = append(finalizers[:i], finalizers[i+1:]...)
			break
		}
	}
	err := unstructured.SetNestedStringSlice(u.Object, finalizers, "metadata", "finalizers")
	if err != nil {
		return err
	}

	err = rClient.Update(ctx, u)
	return err
}

func addFinalizer(ctx context.Context, rClient client.Client, u *unstructured.Unstructured, finalizer string) error {
	finalizers := u.GetFinalizers()
	for _, v := range finalizers {
		if v == finalizer {
			return nil
		}
	}
	finalizers = append(finalizers, finalizer)
	err := unstructured.SetNestedStringSlice(u.Object, finalizers, "metadata", "finalizers")
	if err != nil {
		return err
	}
	err = rClient.Update(ctx, u)
	return err
}

func updateStateField(rClient client.Client, ctx context.Context, intrfc map[string]interface{}, gv schema.GroupVersion, obj *unstructured.Unstructured, jsonit jsoniter.API) error {
	data, err := meta.MarshalToJson(obj, gv)
	if err != nil {
		return err
	}

	typedObj, err := meta.UnmarshalFromJSON(data, gv)
	if err != nil {
		return err
	}

	var raw []byte
	jsonByte, err := json.Marshal(intrfc)
	if err != nil {
		return err
	}

	raw = append(raw, []byte(`{"spec":{ "resource":`)...)
	raw = append(raw, jsonByte...)
	raw = append(raw, []byte(`}}`)...)

	err = jsonit.Unmarshal(raw, &typedObj)
	if err != nil {
		return err
	}

	s := structs.New(typedObj)

	secretData, err := processSensitiveFields(reflect.TypeOf(s.Field("Spec").Field("Resource").Value()), reflect.ValueOf(s.Field("Spec").Field("Resource").Value()))
	if err != nil {
		return err
	}

	if len(secretData) != 0 {
		var secretName string

		secretRef, _, err := unstructured.NestedFieldNoCopy(obj.Object, "spec", "secretRef")
		if err != nil {
			return err
		}

		if secretRef != nil {
			secretName = s.Field("Spec").Field("SecretRef").Field("Name").Value().(string)
		} else {
			secretName = obj.GetName() + "-" + "sensitive"
		}

		var secret corev1.Secret
		req := types.NamespacedName{
			Namespace: obj.GetNamespace(),
			Name:      secretName,
		}
		tr := true
		if err := rClient.Get(ctx, req, &secret); err != nil {
			if errors.ReasonForError(err) == metav1.StatusReasonNotFound {
				err = rClient.Create(ctx, &corev1.Secret{
					ObjectMeta: metav1.ObjectMeta{
						Name:      secretName,
						Namespace: obj.GetNamespace(),
						OwnerReferences: []metav1.OwnerReference{
							{
								APIVersion: obj.GetAPIVersion(),
								Kind:       obj.GetKind(),
								Name:       obj.GetName(),
								Controller: &tr,
								UID:        obj.GetUID(),
							},
						},
					},
				})
				if err != nil {
					return err
				}
			}
			output := s.Field("Spec").Field("Resource").Value()
			specByte, err := json.Marshal(output)
			if err != nil {
				return err
			}
			var specMap map[string]interface{}
			err = json.Unmarshal(specByte, &specMap)
			if err != nil {
				return err
			}
			err = unstructured.SetNestedField(obj.Object, specMap, "spec", "kubeformOutput")
			if err != nil {
				return err
			}
			if err = rClient.Update(ctx, obj); err != nil {
				return err
			}
			return err
		}
		secret.OwnerReferences = []metav1.OwnerReference{
			{
				APIVersion: obj.GetAPIVersion(),
				Kind:       obj.GetKind(),
				Name:       obj.GetName(),
				Controller: &tr,
				UID:        obj.GetUID(),
			},
		}
		if secret.Data == nil {
			secret.Data = make(map[string][]byte)
		}

		secretByte, err := json.Marshal(secretData)
		if err != nil {
			return err
		}
		secret.Data["output"] = secretByte

		// apply the update of the object
		if err = rClient.Update(ctx, &secret); err != nil {
			return err
		}
	}

	output := s.Field("Spec").Field("Resource").Value()
	specByte, err := json.Marshal(output)
	if err != nil {
		return err
	}

	var specMap map[string]interface{}
	err = json.Unmarshal(specByte, &specMap)
	if err != nil {
		return err
	}

	err = unstructured.SetNestedField(obj.Object, specMap, "spec", "kubeformOutput")
	if err != nil {
		return err
	}

	if err = rClient.Update(ctx, obj); err != nil {
		return err
	}
	return nil
}

func processSensitiveFields(r reflect.Type, v reflect.Value) (map[string]interface{}, error) {
	var err error
	out := make(map[string]interface{})
	n := r.NumField()
	for i := 0; i < n; i++ {
		field := r.Field(i)
		value := v.Field(i)
		tag := strings.ReplaceAll(field.Tag.Get("tf"), ",omitempty", "")
		if tag == "-" {
			continue
		}

		if field.Tag.Get("sensitive") == "true" && value.Kind() == reflect.Ptr && value.Elem().Kind() == reflect.String && value.Elem().String() != "" {
			out[tag] = value.Elem().String()
		} else if field.Tag.Get("sensitive") == "true" && value.Kind() == reflect.Map && value.Interface().(map[string]string) != nil && len(value.Interface().(map[string]string)) != 0 {
			secretJson, err := json.Marshal(value.Interface())
			if err != nil {
				return nil, err
			} else {
				out[tag] = string(secretJson)
			}
		}
		if value.Kind() == reflect.Struct {
			out[tag], err = processSensitiveFields(value.Type(), value)
			if err != nil {
				return nil, err
			}
		}
		if value.Kind() == reflect.Ptr && value.Elem().Kind() == reflect.Struct {
			out[tag], err = processSensitiveFields(value.Elem().Type(), value.Elem())
			if err != nil {
				return nil, err
			}
		}

		if value.Kind() == reflect.Slice {
			n := value.Len()
			tempMap := make([]map[string]interface{}, n)
			for i := 0; i < n; i++ {
				if value.Index(i).Kind() == reflect.Struct {
					tempMap[i], err = processSensitiveFields(value.Index(i).Type(), value.Index(i))
					if err != nil {
						return nil, err
					}
				}
			}
			out[tag] = tempMap
		}
	}
	return out, nil
}

func setProviderMeta(rClient client.Client, provider *tfschema.Provider, ctx context.Context, unstructuredObj *unstructured.Unstructured, server *tfschema.GRPCProviderServer, jsonit jsoniter.API) error {
	jsonit = GetJSONItr(aws.GetEncoder(), aws.GetDecoder())
	providerSecretData, err := getProviderSecretData(rClient, ctx, unstructuredObj)
	if err != nil {
		return err
	}
	providerSchema, err := provider.GetSchema(&terraform.ProviderSchemaRequest{})
	if err != nil {
		return err
	}

	if providerSchema.Provider == nil {
		return fmt.Errorf("missing provider schema")
	}

	providerSpec := &aws.AwsSpec{}
	err = jsonit.Unmarshal(providerSecretData["provider"], providerSpec)
	if err != nil {
		return err
	}

	providerDataByte, err := jsonit.Marshal(providerSpec)
	if err != nil {
		return err
	}

	mapData := make(map[string]interface{})
	err = jsonit.Unmarshal(providerDataByte, &mapData)
	if err != nil {
		return err
	}

	configRaw := HCL2ValueFromConfigValue(mapData)
	configPlan, err := msgpack.Marshal(configRaw, providerSchema.Provider.ImpliedType())
	if err != nil {
		return err
	}

	prepareConfigReq := &tfprotov5.PrepareProviderConfigRequest{
		Config: &tfprotov5.DynamicValue{
			MsgPack: configPlan,
		},
	}
	prepareConfigResp, err := server.PrepareProviderConfig(ctx, prepareConfigReq)
	if err != nil {
		return err
	}

	configReq := &tfprotov5.ConfigureProviderRequest{
		Config: prepareConfigResp.PreparedConfig,
	}
	configResp, err := server.ConfigureProvider(ctx, configReq)
	if err != nil {
		return err
	}
	if len(configResp.Diagnostics) > 0 {
		err = diagToError(configResp.Diagnostics)
		if err != nil {
			return err
		}
	}

	return nil
}

func HCL2ValueFromConfigValue(v interface{}) cty.Value {
	if v == nil || v == UnknownIdValue {
		return cty.NullVal(cty.DynamicPseudoType)
	}

	switch tv := v.(type) {
	case bool:
		return cty.BoolVal(tv)
	case string:
		return cty.StringVal(tv)
	case int:
		return cty.NumberIntVal(int64(tv))
	case float64:
		return cty.NumberFloatVal(tv)
	case []interface{}:
		vals := make([]cty.Value, len(tv))
		for i, ev := range tv {
			vals[i] = HCL2ValueFromConfigValue(ev)
		}
		return cty.TupleVal(vals)
	case map[string]interface{}:
		vals := map[string]cty.Value{}
		for k, ev := range tv {
			vals[k] = HCL2ValueFromConfigValue(ev)
		}
		return cty.ObjectVal(vals)
	case map[string]string:
		vals := map[string]cty.Value{}
		for k, ev := range tv {
			vals[k] = HCL2ValueFromConfigValue(ev)
		}
		return cty.ObjectVal(vals)
	default:
		// HCL/HIL should never generate anything that isn't caught by
		// the above, so if we get here something has gone very wrong.
		panic(fmt.Errorf("can't convert %#v to cty.Value", v))
	}
}

func isUpdateSupported(diags []*tfprotov5.Diagnostic) bool {
	for _, diag := range diags {
		if diag.Summary == UpdateNotSupported {
			return false
		}
	}

	return true
}

//func convertCamelToSnakeCase(in map[string]interface{}) map[string]interface{} {
//	out := make(map[string]interface{})
//	for key, val := range in {
//		switch val.(type) {
//		case map[string]interface{}:
//			out[key] = convertCamelToSnakeCase(val.(map[string]interface{}))
//		case []map[string]interface{}:
//			tempMap := make([]map[string]interface{}, len(val.([]map[string]interface{})))
//			for i := range val.([]map[string]interface{}) {
//				tempMap[i] = convertCamelToSnakeCase(val.([]map[string]interface{})[i])
//			}
//			out[key] = tempMap
//		default:
//			out[key] = val
//		}
//	}
//	return out
//}
