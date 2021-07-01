// +build !ignore_autogenerated

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

// Code generated by deepcopy-gen. DO NOT EDIT.

package v1alpha1

import (
	v1 "k8s.io/api/core/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
	apiv1 "kmodules.xyz/client-go/api/v1"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ApiKey) DeepCopyInto(out *ApiKey) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ApiKey.
func (in *ApiKey) DeepCopy() *ApiKey {
	if in == nil {
		return nil
	}
	out := new(ApiKey)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ApiKey) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ApiKeyList) DeepCopyInto(out *ApiKeyList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]ApiKey, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ApiKeyList.
func (in *ApiKeyList) DeepCopy() *ApiKeyList {
	if in == nil {
		return nil
	}
	out := new(ApiKeyList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ApiKeyList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ApiKeySpec) DeepCopyInto(out *ApiKeySpec) {
	*out = *in
	if in.KubeformOutput != nil {
		in, out := &in.KubeformOutput, &out.KubeformOutput
		*out = new(ApiKeySpecResource)
		(*in).DeepCopyInto(*out)
	}
	in.Resource.DeepCopyInto(&out.Resource)
	out.ProviderRef = in.ProviderRef
	if in.SecretRef != nil {
		in, out := &in.SecretRef, &out.SecretRef
		*out = new(v1.LocalObjectReference)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ApiKeySpec.
func (in *ApiKeySpec) DeepCopy() *ApiKeySpec {
	if in == nil {
		return nil
	}
	out := new(ApiKeySpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ApiKeySpecResource) DeepCopyInto(out *ApiKeySpecResource) {
	*out = *in
	if in.ApiID != nil {
		in, out := &in.ApiID, &out.ApiID
		*out = new(string)
		**out = **in
	}
	if in.Description != nil {
		in, out := &in.Description, &out.Description
		*out = new(string)
		**out = **in
	}
	if in.Expires != nil {
		in, out := &in.Expires, &out.Expires
		*out = new(string)
		**out = **in
	}
	if in.Key != nil {
		in, out := &in.Key, &out.Key
		*out = new(string)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ApiKeySpecResource.
func (in *ApiKeySpecResource) DeepCopy() *ApiKeySpecResource {
	if in == nil {
		return nil
	}
	out := new(ApiKeySpecResource)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ApiKeyStatus) DeepCopyInto(out *ApiKeyStatus) {
	*out = *in
	if in.Conditions != nil {
		in, out := &in.Conditions, &out.Conditions
		*out = make([]apiv1.Condition, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ApiKeyStatus.
func (in *ApiKeyStatus) DeepCopy() *ApiKeyStatus {
	if in == nil {
		return nil
	}
	out := new(ApiKeyStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Datasource) DeepCopyInto(out *Datasource) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Datasource.
func (in *Datasource) DeepCopy() *Datasource {
	if in == nil {
		return nil
	}
	out := new(Datasource)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *Datasource) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DatasourceList) DeepCopyInto(out *DatasourceList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]Datasource, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DatasourceList.
func (in *DatasourceList) DeepCopy() *DatasourceList {
	if in == nil {
		return nil
	}
	out := new(DatasourceList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *DatasourceList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DatasourceSpec) DeepCopyInto(out *DatasourceSpec) {
	*out = *in
	if in.KubeformOutput != nil {
		in, out := &in.KubeformOutput, &out.KubeformOutput
		*out = new(DatasourceSpecResource)
		(*in).DeepCopyInto(*out)
	}
	in.Resource.DeepCopyInto(&out.Resource)
	out.ProviderRef = in.ProviderRef
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DatasourceSpec.
func (in *DatasourceSpec) DeepCopy() *DatasourceSpec {
	if in == nil {
		return nil
	}
	out := new(DatasourceSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DatasourceSpecDynamodbConfig) DeepCopyInto(out *DatasourceSpecDynamodbConfig) {
	*out = *in
	if in.Region != nil {
		in, out := &in.Region, &out.Region
		*out = new(string)
		**out = **in
	}
	if in.TableName != nil {
		in, out := &in.TableName, &out.TableName
		*out = new(string)
		**out = **in
	}
	if in.UseCallerCredentials != nil {
		in, out := &in.UseCallerCredentials, &out.UseCallerCredentials
		*out = new(bool)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DatasourceSpecDynamodbConfig.
func (in *DatasourceSpecDynamodbConfig) DeepCopy() *DatasourceSpecDynamodbConfig {
	if in == nil {
		return nil
	}
	out := new(DatasourceSpecDynamodbConfig)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DatasourceSpecElasticsearchConfig) DeepCopyInto(out *DatasourceSpecElasticsearchConfig) {
	*out = *in
	if in.Endpoint != nil {
		in, out := &in.Endpoint, &out.Endpoint
		*out = new(string)
		**out = **in
	}
	if in.Region != nil {
		in, out := &in.Region, &out.Region
		*out = new(string)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DatasourceSpecElasticsearchConfig.
func (in *DatasourceSpecElasticsearchConfig) DeepCopy() *DatasourceSpecElasticsearchConfig {
	if in == nil {
		return nil
	}
	out := new(DatasourceSpecElasticsearchConfig)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DatasourceSpecHttpConfig) DeepCopyInto(out *DatasourceSpecHttpConfig) {
	*out = *in
	if in.Endpoint != nil {
		in, out := &in.Endpoint, &out.Endpoint
		*out = new(string)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DatasourceSpecHttpConfig.
func (in *DatasourceSpecHttpConfig) DeepCopy() *DatasourceSpecHttpConfig {
	if in == nil {
		return nil
	}
	out := new(DatasourceSpecHttpConfig)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DatasourceSpecLambdaConfig) DeepCopyInto(out *DatasourceSpecLambdaConfig) {
	*out = *in
	if in.FunctionArn != nil {
		in, out := &in.FunctionArn, &out.FunctionArn
		*out = new(string)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DatasourceSpecLambdaConfig.
func (in *DatasourceSpecLambdaConfig) DeepCopy() *DatasourceSpecLambdaConfig {
	if in == nil {
		return nil
	}
	out := new(DatasourceSpecLambdaConfig)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DatasourceSpecResource) DeepCopyInto(out *DatasourceSpecResource) {
	*out = *in
	if in.ApiID != nil {
		in, out := &in.ApiID, &out.ApiID
		*out = new(string)
		**out = **in
	}
	if in.Arn != nil {
		in, out := &in.Arn, &out.Arn
		*out = new(string)
		**out = **in
	}
	if in.Description != nil {
		in, out := &in.Description, &out.Description
		*out = new(string)
		**out = **in
	}
	if in.DynamodbConfig != nil {
		in, out := &in.DynamodbConfig, &out.DynamodbConfig
		*out = new(DatasourceSpecDynamodbConfig)
		(*in).DeepCopyInto(*out)
	}
	if in.ElasticsearchConfig != nil {
		in, out := &in.ElasticsearchConfig, &out.ElasticsearchConfig
		*out = new(DatasourceSpecElasticsearchConfig)
		(*in).DeepCopyInto(*out)
	}
	if in.HttpConfig != nil {
		in, out := &in.HttpConfig, &out.HttpConfig
		*out = new(DatasourceSpecHttpConfig)
		(*in).DeepCopyInto(*out)
	}
	if in.LambdaConfig != nil {
		in, out := &in.LambdaConfig, &out.LambdaConfig
		*out = new(DatasourceSpecLambdaConfig)
		(*in).DeepCopyInto(*out)
	}
	if in.Name != nil {
		in, out := &in.Name, &out.Name
		*out = new(string)
		**out = **in
	}
	if in.ServiceRoleArn != nil {
		in, out := &in.ServiceRoleArn, &out.ServiceRoleArn
		*out = new(string)
		**out = **in
	}
	if in.Type != nil {
		in, out := &in.Type, &out.Type
		*out = new(string)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DatasourceSpecResource.
func (in *DatasourceSpecResource) DeepCopy() *DatasourceSpecResource {
	if in == nil {
		return nil
	}
	out := new(DatasourceSpecResource)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DatasourceStatus) DeepCopyInto(out *DatasourceStatus) {
	*out = *in
	if in.Conditions != nil {
		in, out := &in.Conditions, &out.Conditions
		*out = make([]apiv1.Condition, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DatasourceStatus.
func (in *DatasourceStatus) DeepCopy() *DatasourceStatus {
	if in == nil {
		return nil
	}
	out := new(DatasourceStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Function) DeepCopyInto(out *Function) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Function.
func (in *Function) DeepCopy() *Function {
	if in == nil {
		return nil
	}
	out := new(Function)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *Function) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *FunctionList) DeepCopyInto(out *FunctionList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]Function, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new FunctionList.
func (in *FunctionList) DeepCopy() *FunctionList {
	if in == nil {
		return nil
	}
	out := new(FunctionList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *FunctionList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *FunctionSpec) DeepCopyInto(out *FunctionSpec) {
	*out = *in
	if in.KubeformOutput != nil {
		in, out := &in.KubeformOutput, &out.KubeformOutput
		*out = new(FunctionSpecResource)
		(*in).DeepCopyInto(*out)
	}
	in.Resource.DeepCopyInto(&out.Resource)
	out.ProviderRef = in.ProviderRef
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new FunctionSpec.
func (in *FunctionSpec) DeepCopy() *FunctionSpec {
	if in == nil {
		return nil
	}
	out := new(FunctionSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *FunctionSpecResource) DeepCopyInto(out *FunctionSpecResource) {
	*out = *in
	if in.ApiID != nil {
		in, out := &in.ApiID, &out.ApiID
		*out = new(string)
		**out = **in
	}
	if in.Arn != nil {
		in, out := &in.Arn, &out.Arn
		*out = new(string)
		**out = **in
	}
	if in.DataSource != nil {
		in, out := &in.DataSource, &out.DataSource
		*out = new(string)
		**out = **in
	}
	if in.Description != nil {
		in, out := &in.Description, &out.Description
		*out = new(string)
		**out = **in
	}
	if in.FunctionID != nil {
		in, out := &in.FunctionID, &out.FunctionID
		*out = new(string)
		**out = **in
	}
	if in.FunctionVersion != nil {
		in, out := &in.FunctionVersion, &out.FunctionVersion
		*out = new(string)
		**out = **in
	}
	if in.Name != nil {
		in, out := &in.Name, &out.Name
		*out = new(string)
		**out = **in
	}
	if in.RequestMappingTemplate != nil {
		in, out := &in.RequestMappingTemplate, &out.RequestMappingTemplate
		*out = new(string)
		**out = **in
	}
	if in.ResponseMappingTemplate != nil {
		in, out := &in.ResponseMappingTemplate, &out.ResponseMappingTemplate
		*out = new(string)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new FunctionSpecResource.
func (in *FunctionSpecResource) DeepCopy() *FunctionSpecResource {
	if in == nil {
		return nil
	}
	out := new(FunctionSpecResource)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *FunctionStatus) DeepCopyInto(out *FunctionStatus) {
	*out = *in
	if in.Conditions != nil {
		in, out := &in.Conditions, &out.Conditions
		*out = make([]apiv1.Condition, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new FunctionStatus.
func (in *FunctionStatus) DeepCopy() *FunctionStatus {
	if in == nil {
		return nil
	}
	out := new(FunctionStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *GraphqlAPI) DeepCopyInto(out *GraphqlAPI) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new GraphqlAPI.
func (in *GraphqlAPI) DeepCopy() *GraphqlAPI {
	if in == nil {
		return nil
	}
	out := new(GraphqlAPI)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *GraphqlAPI) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *GraphqlAPIList) DeepCopyInto(out *GraphqlAPIList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]GraphqlAPI, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new GraphqlAPIList.
func (in *GraphqlAPIList) DeepCopy() *GraphqlAPIList {
	if in == nil {
		return nil
	}
	out := new(GraphqlAPIList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *GraphqlAPIList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *GraphqlAPISpec) DeepCopyInto(out *GraphqlAPISpec) {
	*out = *in
	if in.KubeformOutput != nil {
		in, out := &in.KubeformOutput, &out.KubeformOutput
		*out = new(GraphqlAPISpecResource)
		(*in).DeepCopyInto(*out)
	}
	in.Resource.DeepCopyInto(&out.Resource)
	out.ProviderRef = in.ProviderRef
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new GraphqlAPISpec.
func (in *GraphqlAPISpec) DeepCopy() *GraphqlAPISpec {
	if in == nil {
		return nil
	}
	out := new(GraphqlAPISpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *GraphqlAPISpecAdditionalAuthenticationProvider) DeepCopyInto(out *GraphqlAPISpecAdditionalAuthenticationProvider) {
	*out = *in
	if in.AuthenticationType != nil {
		in, out := &in.AuthenticationType, &out.AuthenticationType
		*out = new(string)
		**out = **in
	}
	if in.OpenidConnectConfig != nil {
		in, out := &in.OpenidConnectConfig, &out.OpenidConnectConfig
		*out = new(GraphqlAPISpecAdditionalAuthenticationProviderOpenidConnectConfig)
		(*in).DeepCopyInto(*out)
	}
	if in.UserPoolConfig != nil {
		in, out := &in.UserPoolConfig, &out.UserPoolConfig
		*out = new(GraphqlAPISpecAdditionalAuthenticationProviderUserPoolConfig)
		(*in).DeepCopyInto(*out)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new GraphqlAPISpecAdditionalAuthenticationProvider.
func (in *GraphqlAPISpecAdditionalAuthenticationProvider) DeepCopy() *GraphqlAPISpecAdditionalAuthenticationProvider {
	if in == nil {
		return nil
	}
	out := new(GraphqlAPISpecAdditionalAuthenticationProvider)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *GraphqlAPISpecAdditionalAuthenticationProviderOpenidConnectConfig) DeepCopyInto(out *GraphqlAPISpecAdditionalAuthenticationProviderOpenidConnectConfig) {
	*out = *in
	if in.AuthTtl != nil {
		in, out := &in.AuthTtl, &out.AuthTtl
		*out = new(int64)
		**out = **in
	}
	if in.ClientID != nil {
		in, out := &in.ClientID, &out.ClientID
		*out = new(string)
		**out = **in
	}
	if in.IatTtl != nil {
		in, out := &in.IatTtl, &out.IatTtl
		*out = new(int64)
		**out = **in
	}
	if in.Issuer != nil {
		in, out := &in.Issuer, &out.Issuer
		*out = new(string)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new GraphqlAPISpecAdditionalAuthenticationProviderOpenidConnectConfig.
func (in *GraphqlAPISpecAdditionalAuthenticationProviderOpenidConnectConfig) DeepCopy() *GraphqlAPISpecAdditionalAuthenticationProviderOpenidConnectConfig {
	if in == nil {
		return nil
	}
	out := new(GraphqlAPISpecAdditionalAuthenticationProviderOpenidConnectConfig)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *GraphqlAPISpecAdditionalAuthenticationProviderUserPoolConfig) DeepCopyInto(out *GraphqlAPISpecAdditionalAuthenticationProviderUserPoolConfig) {
	*out = *in
	if in.AppIDClientRegex != nil {
		in, out := &in.AppIDClientRegex, &out.AppIDClientRegex
		*out = new(string)
		**out = **in
	}
	if in.AwsRegion != nil {
		in, out := &in.AwsRegion, &out.AwsRegion
		*out = new(string)
		**out = **in
	}
	if in.UserPoolID != nil {
		in, out := &in.UserPoolID, &out.UserPoolID
		*out = new(string)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new GraphqlAPISpecAdditionalAuthenticationProviderUserPoolConfig.
func (in *GraphqlAPISpecAdditionalAuthenticationProviderUserPoolConfig) DeepCopy() *GraphqlAPISpecAdditionalAuthenticationProviderUserPoolConfig {
	if in == nil {
		return nil
	}
	out := new(GraphqlAPISpecAdditionalAuthenticationProviderUserPoolConfig)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *GraphqlAPISpecLogConfig) DeepCopyInto(out *GraphqlAPISpecLogConfig) {
	*out = *in
	if in.CloudwatchLogsRoleArn != nil {
		in, out := &in.CloudwatchLogsRoleArn, &out.CloudwatchLogsRoleArn
		*out = new(string)
		**out = **in
	}
	if in.ExcludeVerboseContent != nil {
		in, out := &in.ExcludeVerboseContent, &out.ExcludeVerboseContent
		*out = new(bool)
		**out = **in
	}
	if in.FieldLogLevel != nil {
		in, out := &in.FieldLogLevel, &out.FieldLogLevel
		*out = new(string)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new GraphqlAPISpecLogConfig.
func (in *GraphqlAPISpecLogConfig) DeepCopy() *GraphqlAPISpecLogConfig {
	if in == nil {
		return nil
	}
	out := new(GraphqlAPISpecLogConfig)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *GraphqlAPISpecOpenidConnectConfig) DeepCopyInto(out *GraphqlAPISpecOpenidConnectConfig) {
	*out = *in
	if in.AuthTtl != nil {
		in, out := &in.AuthTtl, &out.AuthTtl
		*out = new(int64)
		**out = **in
	}
	if in.ClientID != nil {
		in, out := &in.ClientID, &out.ClientID
		*out = new(string)
		**out = **in
	}
	if in.IatTtl != nil {
		in, out := &in.IatTtl, &out.IatTtl
		*out = new(int64)
		**out = **in
	}
	if in.Issuer != nil {
		in, out := &in.Issuer, &out.Issuer
		*out = new(string)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new GraphqlAPISpecOpenidConnectConfig.
func (in *GraphqlAPISpecOpenidConnectConfig) DeepCopy() *GraphqlAPISpecOpenidConnectConfig {
	if in == nil {
		return nil
	}
	out := new(GraphqlAPISpecOpenidConnectConfig)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *GraphqlAPISpecResource) DeepCopyInto(out *GraphqlAPISpecResource) {
	*out = *in
	if in.AdditionalAuthenticationProvider != nil {
		in, out := &in.AdditionalAuthenticationProvider, &out.AdditionalAuthenticationProvider
		*out = make([]GraphqlAPISpecAdditionalAuthenticationProvider, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.Arn != nil {
		in, out := &in.Arn, &out.Arn
		*out = new(string)
		**out = **in
	}
	if in.AuthenticationType != nil {
		in, out := &in.AuthenticationType, &out.AuthenticationType
		*out = new(string)
		**out = **in
	}
	if in.LogConfig != nil {
		in, out := &in.LogConfig, &out.LogConfig
		*out = new(GraphqlAPISpecLogConfig)
		(*in).DeepCopyInto(*out)
	}
	if in.Name != nil {
		in, out := &in.Name, &out.Name
		*out = new(string)
		**out = **in
	}
	if in.OpenidConnectConfig != nil {
		in, out := &in.OpenidConnectConfig, &out.OpenidConnectConfig
		*out = new(GraphqlAPISpecOpenidConnectConfig)
		(*in).DeepCopyInto(*out)
	}
	if in.Schema != nil {
		in, out := &in.Schema, &out.Schema
		*out = new(string)
		**out = **in
	}
	if in.Tags != nil {
		in, out := &in.Tags, &out.Tags
		*out = new(map[string]string)
		if **in != nil {
			in, out := *in, *out
			*out = make(map[string]string, len(*in))
			for key, val := range *in {
				(*out)[key] = val
			}
		}
	}
	if in.TagsAll != nil {
		in, out := &in.TagsAll, &out.TagsAll
		*out = new(map[string]string)
		if **in != nil {
			in, out := *in, *out
			*out = make(map[string]string, len(*in))
			for key, val := range *in {
				(*out)[key] = val
			}
		}
	}
	if in.Uris != nil {
		in, out := &in.Uris, &out.Uris
		*out = new(map[string]string)
		if **in != nil {
			in, out := *in, *out
			*out = make(map[string]string, len(*in))
			for key, val := range *in {
				(*out)[key] = val
			}
		}
	}
	if in.UserPoolConfig != nil {
		in, out := &in.UserPoolConfig, &out.UserPoolConfig
		*out = new(GraphqlAPISpecUserPoolConfig)
		(*in).DeepCopyInto(*out)
	}
	if in.XrayEnabled != nil {
		in, out := &in.XrayEnabled, &out.XrayEnabled
		*out = new(bool)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new GraphqlAPISpecResource.
func (in *GraphqlAPISpecResource) DeepCopy() *GraphqlAPISpecResource {
	if in == nil {
		return nil
	}
	out := new(GraphqlAPISpecResource)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *GraphqlAPISpecUserPoolConfig) DeepCopyInto(out *GraphqlAPISpecUserPoolConfig) {
	*out = *in
	if in.AppIDClientRegex != nil {
		in, out := &in.AppIDClientRegex, &out.AppIDClientRegex
		*out = new(string)
		**out = **in
	}
	if in.AwsRegion != nil {
		in, out := &in.AwsRegion, &out.AwsRegion
		*out = new(string)
		**out = **in
	}
	if in.DefaultAction != nil {
		in, out := &in.DefaultAction, &out.DefaultAction
		*out = new(string)
		**out = **in
	}
	if in.UserPoolID != nil {
		in, out := &in.UserPoolID, &out.UserPoolID
		*out = new(string)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new GraphqlAPISpecUserPoolConfig.
func (in *GraphqlAPISpecUserPoolConfig) DeepCopy() *GraphqlAPISpecUserPoolConfig {
	if in == nil {
		return nil
	}
	out := new(GraphqlAPISpecUserPoolConfig)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *GraphqlAPIStatus) DeepCopyInto(out *GraphqlAPIStatus) {
	*out = *in
	if in.Conditions != nil {
		in, out := &in.Conditions, &out.Conditions
		*out = make([]apiv1.Condition, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new GraphqlAPIStatus.
func (in *GraphqlAPIStatus) DeepCopy() *GraphqlAPIStatus {
	if in == nil {
		return nil
	}
	out := new(GraphqlAPIStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Resolver) DeepCopyInto(out *Resolver) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Resolver.
func (in *Resolver) DeepCopy() *Resolver {
	if in == nil {
		return nil
	}
	out := new(Resolver)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *Resolver) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ResolverList) DeepCopyInto(out *ResolverList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]Resolver, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ResolverList.
func (in *ResolverList) DeepCopy() *ResolverList {
	if in == nil {
		return nil
	}
	out := new(ResolverList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ResolverList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ResolverSpec) DeepCopyInto(out *ResolverSpec) {
	*out = *in
	if in.KubeformOutput != nil {
		in, out := &in.KubeformOutput, &out.KubeformOutput
		*out = new(ResolverSpecResource)
		(*in).DeepCopyInto(*out)
	}
	in.Resource.DeepCopyInto(&out.Resource)
	out.ProviderRef = in.ProviderRef
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ResolverSpec.
func (in *ResolverSpec) DeepCopy() *ResolverSpec {
	if in == nil {
		return nil
	}
	out := new(ResolverSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ResolverSpecCachingConfig) DeepCopyInto(out *ResolverSpecCachingConfig) {
	*out = *in
	if in.CachingKeys != nil {
		in, out := &in.CachingKeys, &out.CachingKeys
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.Ttl != nil {
		in, out := &in.Ttl, &out.Ttl
		*out = new(int64)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ResolverSpecCachingConfig.
func (in *ResolverSpecCachingConfig) DeepCopy() *ResolverSpecCachingConfig {
	if in == nil {
		return nil
	}
	out := new(ResolverSpecCachingConfig)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ResolverSpecPipelineConfig) DeepCopyInto(out *ResolverSpecPipelineConfig) {
	*out = *in
	if in.Functions != nil {
		in, out := &in.Functions, &out.Functions
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ResolverSpecPipelineConfig.
func (in *ResolverSpecPipelineConfig) DeepCopy() *ResolverSpecPipelineConfig {
	if in == nil {
		return nil
	}
	out := new(ResolverSpecPipelineConfig)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ResolverSpecResource) DeepCopyInto(out *ResolverSpecResource) {
	*out = *in
	if in.ApiID != nil {
		in, out := &in.ApiID, &out.ApiID
		*out = new(string)
		**out = **in
	}
	if in.Arn != nil {
		in, out := &in.Arn, &out.Arn
		*out = new(string)
		**out = **in
	}
	if in.CachingConfig != nil {
		in, out := &in.CachingConfig, &out.CachingConfig
		*out = new(ResolverSpecCachingConfig)
		(*in).DeepCopyInto(*out)
	}
	if in.DataSource != nil {
		in, out := &in.DataSource, &out.DataSource
		*out = new(string)
		**out = **in
	}
	if in.Field != nil {
		in, out := &in.Field, &out.Field
		*out = new(string)
		**out = **in
	}
	if in.Kind != nil {
		in, out := &in.Kind, &out.Kind
		*out = new(string)
		**out = **in
	}
	if in.PipelineConfig != nil {
		in, out := &in.PipelineConfig, &out.PipelineConfig
		*out = new(ResolverSpecPipelineConfig)
		(*in).DeepCopyInto(*out)
	}
	if in.RequestTemplate != nil {
		in, out := &in.RequestTemplate, &out.RequestTemplate
		*out = new(string)
		**out = **in
	}
	if in.ResponseTemplate != nil {
		in, out := &in.ResponseTemplate, &out.ResponseTemplate
		*out = new(string)
		**out = **in
	}
	if in.Type != nil {
		in, out := &in.Type, &out.Type
		*out = new(string)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ResolverSpecResource.
func (in *ResolverSpecResource) DeepCopy() *ResolverSpecResource {
	if in == nil {
		return nil
	}
	out := new(ResolverSpecResource)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ResolverStatus) DeepCopyInto(out *ResolverStatus) {
	*out = *in
	if in.Conditions != nil {
		in, out := &in.Conditions, &out.Conditions
		*out = make([]apiv1.Condition, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ResolverStatus.
func (in *ResolverStatus) DeepCopy() *ResolverStatus {
	if in == nil {
		return nil
	}
	out := new(ResolverStatus)
	in.DeepCopyInto(out)
	return out
}
