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
		jsoniter.MustGetKind(reflect2.TypeOf(ClusterSpecLogging{}).Type1()):      ClusterSpecLoggingCodec{},
		jsoniter.MustGetKind(reflect2.TypeOf(ClusterSpecSnapshotCopy{}).Type1()): ClusterSpecSnapshotCopyCodec{},
	}
}

func GetDecoder() map[string]jsoniter.ValDecoder {
	return map[string]jsoniter.ValDecoder{
		jsoniter.MustGetKind(reflect2.TypeOf(ClusterSpecLogging{}).Type1()):      ClusterSpecLoggingCodec{},
		jsoniter.MustGetKind(reflect2.TypeOf(ClusterSpecSnapshotCopy{}).Type1()): ClusterSpecSnapshotCopyCodec{},
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
type ClusterSpecLoggingCodec struct {
}

func (ClusterSpecLoggingCodec) IsEmpty(ptr unsafe.Pointer) bool {
	return (*ClusterSpecLogging)(ptr) == nil
}

func (ClusterSpecLoggingCodec) Encode(ptr unsafe.Pointer, stream *jsoniter.Stream) {
	obj := (*ClusterSpecLogging)(ptr)
	var objs []ClusterSpecLogging
	if obj != nil {
		objs = []ClusterSpecLogging{*obj}
	}

	jsonit := jsoniter.Config{
		EscapeHTML:             true,
		SortMapKeys:            true,
		ValidateJsonRawMessage: true,
		TagKey:                 "tf",
		TypeEncoders:           getEncodersWithout(jsoniter.MustGetKind(reflect2.TypeOf(ClusterSpecLogging{}).Type1())),
	}.Froze()

	byt, _ := jsonit.Marshal(objs)

	stream.Write(byt)
}

func (ClusterSpecLoggingCodec) Decode(ptr unsafe.Pointer, iter *jsoniter.Iterator) {
	switch iter.WhatIsNext() {
	case jsoniter.NilValue:
		iter.Skip()
		*(*ClusterSpecLogging)(ptr) = ClusterSpecLogging{}
		return
	case jsoniter.ArrayValue:
		objsByte := iter.SkipAndReturnBytes()
		if len(objsByte) > 0 {
			var objs []ClusterSpecLogging

			jsonit := jsoniter.Config{
				EscapeHTML:             true,
				SortMapKeys:            true,
				ValidateJsonRawMessage: true,
				TagKey:                 "tf",
				TypeDecoders:           getDecodersWithout(jsoniter.MustGetKind(reflect2.TypeOf(ClusterSpecLogging{}).Type1())),
			}.Froze()
			jsonit.Unmarshal(objsByte, &objs)

			if len(objs) > 0 {
				*(*ClusterSpecLogging)(ptr) = objs[0]
			} else {
				*(*ClusterSpecLogging)(ptr) = ClusterSpecLogging{}
			}
		} else {
			*(*ClusterSpecLogging)(ptr) = ClusterSpecLogging{}
		}
	case jsoniter.ObjectValue:
		objByte := iter.SkipAndReturnBytes()
		if len(objByte) > 0 {
			var obj ClusterSpecLogging

			jsonit := jsoniter.Config{
				EscapeHTML:             true,
				SortMapKeys:            true,
				ValidateJsonRawMessage: true,
				TagKey:                 "tf",
				TypeDecoders:           getDecodersWithout(jsoniter.MustGetKind(reflect2.TypeOf(ClusterSpecLogging{}).Type1())),
			}.Froze()
			jsonit.Unmarshal(objByte, &obj)

			*(*ClusterSpecLogging)(ptr) = obj
		} else {
			*(*ClusterSpecLogging)(ptr) = ClusterSpecLogging{}
		}
	default:
		iter.ReportError("decode ClusterSpecLogging", "unexpected JSON type")
	}
}

// +k8s:deepcopy-gen=false
type ClusterSpecSnapshotCopyCodec struct {
}

func (ClusterSpecSnapshotCopyCodec) IsEmpty(ptr unsafe.Pointer) bool {
	return (*ClusterSpecSnapshotCopy)(ptr) == nil
}

func (ClusterSpecSnapshotCopyCodec) Encode(ptr unsafe.Pointer, stream *jsoniter.Stream) {
	obj := (*ClusterSpecSnapshotCopy)(ptr)
	var objs []ClusterSpecSnapshotCopy
	if obj != nil {
		objs = []ClusterSpecSnapshotCopy{*obj}
	}

	jsonit := jsoniter.Config{
		EscapeHTML:             true,
		SortMapKeys:            true,
		ValidateJsonRawMessage: true,
		TagKey:                 "tf",
		TypeEncoders:           getEncodersWithout(jsoniter.MustGetKind(reflect2.TypeOf(ClusterSpecSnapshotCopy{}).Type1())),
	}.Froze()

	byt, _ := jsonit.Marshal(objs)

	stream.Write(byt)
}

func (ClusterSpecSnapshotCopyCodec) Decode(ptr unsafe.Pointer, iter *jsoniter.Iterator) {
	switch iter.WhatIsNext() {
	case jsoniter.NilValue:
		iter.Skip()
		*(*ClusterSpecSnapshotCopy)(ptr) = ClusterSpecSnapshotCopy{}
		return
	case jsoniter.ArrayValue:
		objsByte := iter.SkipAndReturnBytes()
		if len(objsByte) > 0 {
			var objs []ClusterSpecSnapshotCopy

			jsonit := jsoniter.Config{
				EscapeHTML:             true,
				SortMapKeys:            true,
				ValidateJsonRawMessage: true,
				TagKey:                 "tf",
				TypeDecoders:           getDecodersWithout(jsoniter.MustGetKind(reflect2.TypeOf(ClusterSpecSnapshotCopy{}).Type1())),
			}.Froze()
			jsonit.Unmarshal(objsByte, &objs)

			if len(objs) > 0 {
				*(*ClusterSpecSnapshotCopy)(ptr) = objs[0]
			} else {
				*(*ClusterSpecSnapshotCopy)(ptr) = ClusterSpecSnapshotCopy{}
			}
		} else {
			*(*ClusterSpecSnapshotCopy)(ptr) = ClusterSpecSnapshotCopy{}
		}
	case jsoniter.ObjectValue:
		objByte := iter.SkipAndReturnBytes()
		if len(objByte) > 0 {
			var obj ClusterSpecSnapshotCopy

			jsonit := jsoniter.Config{
				EscapeHTML:             true,
				SortMapKeys:            true,
				ValidateJsonRawMessage: true,
				TagKey:                 "tf",
				TypeDecoders:           getDecodersWithout(jsoniter.MustGetKind(reflect2.TypeOf(ClusterSpecSnapshotCopy{}).Type1())),
			}.Froze()
			jsonit.Unmarshal(objByte, &obj)

			*(*ClusterSpecSnapshotCopy)(ptr) = obj
		} else {
			*(*ClusterSpecSnapshotCopy)(ptr) = ClusterSpecSnapshotCopy{}
		}
	default:
		iter.ReportError("decode ClusterSpecSnapshotCopy", "unexpected JSON type")
	}
}
