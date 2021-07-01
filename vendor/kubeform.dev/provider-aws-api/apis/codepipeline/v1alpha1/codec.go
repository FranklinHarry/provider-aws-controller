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
	"unsafe"

	jsoniter "github.com/json-iterator/go"
	"github.com/modern-go/reflect2"
)

func GetEncoder() map[string]jsoniter.ValEncoder {
	return map[string]jsoniter.ValEncoder{
		jsoniter.MustGetKind(reflect2.TypeOf(CodepipelineSpecArtifactStoreEncryptionKey{}).Type1()): CodepipelineSpecArtifactStoreEncryptionKeyCodec{},
		jsoniter.MustGetKind(reflect2.TypeOf(WebhookSpecAuthenticationConfiguration{}).Type1()):     WebhookSpecAuthenticationConfigurationCodec{},
	}
}

func GetDecoder() map[string]jsoniter.ValDecoder {
	return map[string]jsoniter.ValDecoder{
		jsoniter.MustGetKind(reflect2.TypeOf(CodepipelineSpecArtifactStoreEncryptionKey{}).Type1()): CodepipelineSpecArtifactStoreEncryptionKeyCodec{},
		jsoniter.MustGetKind(reflect2.TypeOf(WebhookSpecAuthenticationConfiguration{}).Type1()):     WebhookSpecAuthenticationConfigurationCodec{},
	}
}

func getEncodersWithout(typ string) map[string]jsoniter.ValEncoder {
	origMap := GetEncoder()
	delete(origMap, typ)
	return origMap
}

func getDecodersWithout(typ string) map[string]jsoniter.ValDecoder {
	origMap := GetDecoder()
	delete(origMap, typ)
	return origMap
}

// +k8s:deepcopy-gen=false
type CodepipelineSpecArtifactStoreEncryptionKeyCodec struct {
}

func (CodepipelineSpecArtifactStoreEncryptionKeyCodec) IsEmpty(ptr unsafe.Pointer) bool {
	return (*CodepipelineSpecArtifactStoreEncryptionKey)(ptr) == nil
}

func (CodepipelineSpecArtifactStoreEncryptionKeyCodec) Encode(ptr unsafe.Pointer, stream *jsoniter.Stream) {
	obj := (*CodepipelineSpecArtifactStoreEncryptionKey)(ptr)
	var objs []CodepipelineSpecArtifactStoreEncryptionKey
	if obj != nil {
		objs = []CodepipelineSpecArtifactStoreEncryptionKey{*obj}
	}

	jsonit := jsoniter.Config{
		EscapeHTML:             true,
		SortMapKeys:            true,
		ValidateJsonRawMessage: true,
		TagKey:                 "tf",
		TypeEncoders:           getEncodersWithout(jsoniter.MustGetKind(reflect2.TypeOf(CodepipelineSpecArtifactStoreEncryptionKey{}).Type1())),
	}.Froze()

	byt, _ := jsonit.Marshal(objs)

	stream.Write(byt)
}

func (CodepipelineSpecArtifactStoreEncryptionKeyCodec) Decode(ptr unsafe.Pointer, iter *jsoniter.Iterator) {
	switch iter.WhatIsNext() {
	case jsoniter.NilValue:
		iter.Skip()
		*(*CodepipelineSpecArtifactStoreEncryptionKey)(ptr) = CodepipelineSpecArtifactStoreEncryptionKey{}
		return
	case jsoniter.ArrayValue:
		objsByte := iter.SkipAndReturnBytes()
		if len(objsByte) > 0 {
			var objs []CodepipelineSpecArtifactStoreEncryptionKey

			jsonit := jsoniter.Config{
				EscapeHTML:             true,
				SortMapKeys:            true,
				ValidateJsonRawMessage: true,
				TagKey:                 "tf",
				TypeDecoders:           getDecodersWithout(jsoniter.MustGetKind(reflect2.TypeOf(CodepipelineSpecArtifactStoreEncryptionKey{}).Type1())),
			}.Froze()
			jsonit.Unmarshal(objsByte, &objs)

			if len(objs) > 0 {
				*(*CodepipelineSpecArtifactStoreEncryptionKey)(ptr) = objs[0]
			} else {
				*(*CodepipelineSpecArtifactStoreEncryptionKey)(ptr) = CodepipelineSpecArtifactStoreEncryptionKey{}
			}
		} else {
			*(*CodepipelineSpecArtifactStoreEncryptionKey)(ptr) = CodepipelineSpecArtifactStoreEncryptionKey{}
		}
	case jsoniter.ObjectValue:
		objByte := iter.SkipAndReturnBytes()
		if len(objByte) > 0 {
			var obj CodepipelineSpecArtifactStoreEncryptionKey

			jsonit := jsoniter.Config{
				EscapeHTML:             true,
				SortMapKeys:            true,
				ValidateJsonRawMessage: true,
				TagKey:                 "tf",
				TypeDecoders:           getDecodersWithout(jsoniter.MustGetKind(reflect2.TypeOf(CodepipelineSpecArtifactStoreEncryptionKey{}).Type1())),
			}.Froze()
			jsonit.Unmarshal(objByte, &obj)

			*(*CodepipelineSpecArtifactStoreEncryptionKey)(ptr) = obj
		} else {
			*(*CodepipelineSpecArtifactStoreEncryptionKey)(ptr) = CodepipelineSpecArtifactStoreEncryptionKey{}
		}
	default:
		iter.ReportError("decode CodepipelineSpecArtifactStoreEncryptionKey", "unexpected JSON type")
	}
}

// +k8s:deepcopy-gen=false
type WebhookSpecAuthenticationConfigurationCodec struct {
}

func (WebhookSpecAuthenticationConfigurationCodec) IsEmpty(ptr unsafe.Pointer) bool {
	return (*WebhookSpecAuthenticationConfiguration)(ptr) == nil
}

func (WebhookSpecAuthenticationConfigurationCodec) Encode(ptr unsafe.Pointer, stream *jsoniter.Stream) {
	obj := (*WebhookSpecAuthenticationConfiguration)(ptr)
	var objs []WebhookSpecAuthenticationConfiguration
	if obj != nil {
		objs = []WebhookSpecAuthenticationConfiguration{*obj}
	}

	jsonit := jsoniter.Config{
		EscapeHTML:             true,
		SortMapKeys:            true,
		ValidateJsonRawMessage: true,
		TagKey:                 "tf",
		TypeEncoders:           getEncodersWithout(jsoniter.MustGetKind(reflect2.TypeOf(WebhookSpecAuthenticationConfiguration{}).Type1())),
	}.Froze()

	byt, _ := jsonit.Marshal(objs)

	stream.Write(byt)
}

func (WebhookSpecAuthenticationConfigurationCodec) Decode(ptr unsafe.Pointer, iter *jsoniter.Iterator) {
	switch iter.WhatIsNext() {
	case jsoniter.NilValue:
		iter.Skip()
		*(*WebhookSpecAuthenticationConfiguration)(ptr) = WebhookSpecAuthenticationConfiguration{}
		return
	case jsoniter.ArrayValue:
		objsByte := iter.SkipAndReturnBytes()
		if len(objsByte) > 0 {
			var objs []WebhookSpecAuthenticationConfiguration

			jsonit := jsoniter.Config{
				EscapeHTML:             true,
				SortMapKeys:            true,
				ValidateJsonRawMessage: true,
				TagKey:                 "tf",
				TypeDecoders:           getDecodersWithout(jsoniter.MustGetKind(reflect2.TypeOf(WebhookSpecAuthenticationConfiguration{}).Type1())),
			}.Froze()
			jsonit.Unmarshal(objsByte, &objs)

			if len(objs) > 0 {
				*(*WebhookSpecAuthenticationConfiguration)(ptr) = objs[0]
			} else {
				*(*WebhookSpecAuthenticationConfiguration)(ptr) = WebhookSpecAuthenticationConfiguration{}
			}
		} else {
			*(*WebhookSpecAuthenticationConfiguration)(ptr) = WebhookSpecAuthenticationConfiguration{}
		}
	case jsoniter.ObjectValue:
		objByte := iter.SkipAndReturnBytes()
		if len(objByte) > 0 {
			var obj WebhookSpecAuthenticationConfiguration

			jsonit := jsoniter.Config{
				EscapeHTML:             true,
				SortMapKeys:            true,
				ValidateJsonRawMessage: true,
				TagKey:                 "tf",
				TypeDecoders:           getDecodersWithout(jsoniter.MustGetKind(reflect2.TypeOf(WebhookSpecAuthenticationConfiguration{}).Type1())),
			}.Froze()
			jsonit.Unmarshal(objByte, &obj)

			*(*WebhookSpecAuthenticationConfiguration)(ptr) = obj
		} else {
			*(*WebhookSpecAuthenticationConfiguration)(ptr) = WebhookSpecAuthenticationConfiguration{}
		}
	default:
		iter.ReportError("decode WebhookSpecAuthenticationConfiguration", "unexpected JSON type")
	}
}
