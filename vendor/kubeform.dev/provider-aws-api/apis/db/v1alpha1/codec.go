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
		jsoniter.MustGetKind(reflect2.TypeOf(InstanceSpecRestoreToPointInTime{}).Type1()):                InstanceSpecRestoreToPointInTimeCodec{},
		jsoniter.MustGetKind(reflect2.TypeOf(InstanceSpecS3Import{}).Type1()):                            InstanceSpecS3ImportCodec{},
		jsoniter.MustGetKind(reflect2.TypeOf(ProxyDefaultTargetGroupSpecConnectionPoolConfig{}).Type1()): ProxyDefaultTargetGroupSpecConnectionPoolConfigCodec{},
	}
}

func GetDecoder() map[string]jsoniter.ValDecoder {
	return map[string]jsoniter.ValDecoder{
		jsoniter.MustGetKind(reflect2.TypeOf(InstanceSpecRestoreToPointInTime{}).Type1()):                InstanceSpecRestoreToPointInTimeCodec{},
		jsoniter.MustGetKind(reflect2.TypeOf(InstanceSpecS3Import{}).Type1()):                            InstanceSpecS3ImportCodec{},
		jsoniter.MustGetKind(reflect2.TypeOf(ProxyDefaultTargetGroupSpecConnectionPoolConfig{}).Type1()): ProxyDefaultTargetGroupSpecConnectionPoolConfigCodec{},
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
type InstanceSpecRestoreToPointInTimeCodec struct {
}

func (InstanceSpecRestoreToPointInTimeCodec) IsEmpty(ptr unsafe.Pointer) bool {
	return (*InstanceSpecRestoreToPointInTime)(ptr) == nil
}

func (InstanceSpecRestoreToPointInTimeCodec) Encode(ptr unsafe.Pointer, stream *jsoniter.Stream) {
	obj := (*InstanceSpecRestoreToPointInTime)(ptr)
	var objs []InstanceSpecRestoreToPointInTime
	if obj != nil {
		objs = []InstanceSpecRestoreToPointInTime{*obj}
	}

	jsonit := jsoniter.Config{
		EscapeHTML:             true,
		SortMapKeys:            true,
		ValidateJsonRawMessage: true,
		TagKey:                 "tf",
		TypeEncoders:           getEncodersWithout(jsoniter.MustGetKind(reflect2.TypeOf(InstanceSpecRestoreToPointInTime{}).Type1())),
	}.Froze()

	byt, _ := jsonit.Marshal(objs)

	stream.Write(byt)
}

func (InstanceSpecRestoreToPointInTimeCodec) Decode(ptr unsafe.Pointer, iter *jsoniter.Iterator) {
	switch iter.WhatIsNext() {
	case jsoniter.NilValue:
		iter.Skip()
		*(*InstanceSpecRestoreToPointInTime)(ptr) = InstanceSpecRestoreToPointInTime{}
		return
	case jsoniter.ArrayValue:
		objsByte := iter.SkipAndReturnBytes()
		if len(objsByte) > 0 {
			var objs []InstanceSpecRestoreToPointInTime

			jsonit := jsoniter.Config{
				EscapeHTML:             true,
				SortMapKeys:            true,
				ValidateJsonRawMessage: true,
				TagKey:                 "tf",
				TypeDecoders:           getDecodersWithout(jsoniter.MustGetKind(reflect2.TypeOf(InstanceSpecRestoreToPointInTime{}).Type1())),
			}.Froze()
			jsonit.Unmarshal(objsByte, &objs)

			if len(objs) > 0 {
				*(*InstanceSpecRestoreToPointInTime)(ptr) = objs[0]
			} else {
				*(*InstanceSpecRestoreToPointInTime)(ptr) = InstanceSpecRestoreToPointInTime{}
			}
		} else {
			*(*InstanceSpecRestoreToPointInTime)(ptr) = InstanceSpecRestoreToPointInTime{}
		}
	case jsoniter.ObjectValue:
		objByte := iter.SkipAndReturnBytes()
		if len(objByte) > 0 {
			var obj InstanceSpecRestoreToPointInTime

			jsonit := jsoniter.Config{
				EscapeHTML:             true,
				SortMapKeys:            true,
				ValidateJsonRawMessage: true,
				TagKey:                 "tf",
				TypeDecoders:           getDecodersWithout(jsoniter.MustGetKind(reflect2.TypeOf(InstanceSpecRestoreToPointInTime{}).Type1())),
			}.Froze()
			jsonit.Unmarshal(objByte, &obj)

			*(*InstanceSpecRestoreToPointInTime)(ptr) = obj
		} else {
			*(*InstanceSpecRestoreToPointInTime)(ptr) = InstanceSpecRestoreToPointInTime{}
		}
	default:
		iter.ReportError("decode InstanceSpecRestoreToPointInTime", "unexpected JSON type")
	}
}

// +k8s:deepcopy-gen=false
type InstanceSpecS3ImportCodec struct {
}

func (InstanceSpecS3ImportCodec) IsEmpty(ptr unsafe.Pointer) bool {
	return (*InstanceSpecS3Import)(ptr) == nil
}

func (InstanceSpecS3ImportCodec) Encode(ptr unsafe.Pointer, stream *jsoniter.Stream) {
	obj := (*InstanceSpecS3Import)(ptr)
	var objs []InstanceSpecS3Import
	if obj != nil {
		objs = []InstanceSpecS3Import{*obj}
	}

	jsonit := jsoniter.Config{
		EscapeHTML:             true,
		SortMapKeys:            true,
		ValidateJsonRawMessage: true,
		TagKey:                 "tf",
		TypeEncoders:           getEncodersWithout(jsoniter.MustGetKind(reflect2.TypeOf(InstanceSpecS3Import{}).Type1())),
	}.Froze()

	byt, _ := jsonit.Marshal(objs)

	stream.Write(byt)
}

func (InstanceSpecS3ImportCodec) Decode(ptr unsafe.Pointer, iter *jsoniter.Iterator) {
	switch iter.WhatIsNext() {
	case jsoniter.NilValue:
		iter.Skip()
		*(*InstanceSpecS3Import)(ptr) = InstanceSpecS3Import{}
		return
	case jsoniter.ArrayValue:
		objsByte := iter.SkipAndReturnBytes()
		if len(objsByte) > 0 {
			var objs []InstanceSpecS3Import

			jsonit := jsoniter.Config{
				EscapeHTML:             true,
				SortMapKeys:            true,
				ValidateJsonRawMessage: true,
				TagKey:                 "tf",
				TypeDecoders:           getDecodersWithout(jsoniter.MustGetKind(reflect2.TypeOf(InstanceSpecS3Import{}).Type1())),
			}.Froze()
			jsonit.Unmarshal(objsByte, &objs)

			if len(objs) > 0 {
				*(*InstanceSpecS3Import)(ptr) = objs[0]
			} else {
				*(*InstanceSpecS3Import)(ptr) = InstanceSpecS3Import{}
			}
		} else {
			*(*InstanceSpecS3Import)(ptr) = InstanceSpecS3Import{}
		}
	case jsoniter.ObjectValue:
		objByte := iter.SkipAndReturnBytes()
		if len(objByte) > 0 {
			var obj InstanceSpecS3Import

			jsonit := jsoniter.Config{
				EscapeHTML:             true,
				SortMapKeys:            true,
				ValidateJsonRawMessage: true,
				TagKey:                 "tf",
				TypeDecoders:           getDecodersWithout(jsoniter.MustGetKind(reflect2.TypeOf(InstanceSpecS3Import{}).Type1())),
			}.Froze()
			jsonit.Unmarshal(objByte, &obj)

			*(*InstanceSpecS3Import)(ptr) = obj
		} else {
			*(*InstanceSpecS3Import)(ptr) = InstanceSpecS3Import{}
		}
	default:
		iter.ReportError("decode InstanceSpecS3Import", "unexpected JSON type")
	}
}

// +k8s:deepcopy-gen=false
type ProxyDefaultTargetGroupSpecConnectionPoolConfigCodec struct {
}

func (ProxyDefaultTargetGroupSpecConnectionPoolConfigCodec) IsEmpty(ptr unsafe.Pointer) bool {
	return (*ProxyDefaultTargetGroupSpecConnectionPoolConfig)(ptr) == nil
}

func (ProxyDefaultTargetGroupSpecConnectionPoolConfigCodec) Encode(ptr unsafe.Pointer, stream *jsoniter.Stream) {
	obj := (*ProxyDefaultTargetGroupSpecConnectionPoolConfig)(ptr)
	var objs []ProxyDefaultTargetGroupSpecConnectionPoolConfig
	if obj != nil {
		objs = []ProxyDefaultTargetGroupSpecConnectionPoolConfig{*obj}
	}

	jsonit := jsoniter.Config{
		EscapeHTML:             true,
		SortMapKeys:            true,
		ValidateJsonRawMessage: true,
		TagKey:                 "tf",
		TypeEncoders:           getEncodersWithout(jsoniter.MustGetKind(reflect2.TypeOf(ProxyDefaultTargetGroupSpecConnectionPoolConfig{}).Type1())),
	}.Froze()

	byt, _ := jsonit.Marshal(objs)

	stream.Write(byt)
}

func (ProxyDefaultTargetGroupSpecConnectionPoolConfigCodec) Decode(ptr unsafe.Pointer, iter *jsoniter.Iterator) {
	switch iter.WhatIsNext() {
	case jsoniter.NilValue:
		iter.Skip()
		*(*ProxyDefaultTargetGroupSpecConnectionPoolConfig)(ptr) = ProxyDefaultTargetGroupSpecConnectionPoolConfig{}
		return
	case jsoniter.ArrayValue:
		objsByte := iter.SkipAndReturnBytes()
		if len(objsByte) > 0 {
			var objs []ProxyDefaultTargetGroupSpecConnectionPoolConfig

			jsonit := jsoniter.Config{
				EscapeHTML:             true,
				SortMapKeys:            true,
				ValidateJsonRawMessage: true,
				TagKey:                 "tf",
				TypeDecoders:           getDecodersWithout(jsoniter.MustGetKind(reflect2.TypeOf(ProxyDefaultTargetGroupSpecConnectionPoolConfig{}).Type1())),
			}.Froze()
			jsonit.Unmarshal(objsByte, &objs)

			if len(objs) > 0 {
				*(*ProxyDefaultTargetGroupSpecConnectionPoolConfig)(ptr) = objs[0]
			} else {
				*(*ProxyDefaultTargetGroupSpecConnectionPoolConfig)(ptr) = ProxyDefaultTargetGroupSpecConnectionPoolConfig{}
			}
		} else {
			*(*ProxyDefaultTargetGroupSpecConnectionPoolConfig)(ptr) = ProxyDefaultTargetGroupSpecConnectionPoolConfig{}
		}
	case jsoniter.ObjectValue:
		objByte := iter.SkipAndReturnBytes()
		if len(objByte) > 0 {
			var obj ProxyDefaultTargetGroupSpecConnectionPoolConfig

			jsonit := jsoniter.Config{
				EscapeHTML:             true,
				SortMapKeys:            true,
				ValidateJsonRawMessage: true,
				TagKey:                 "tf",
				TypeDecoders:           getDecodersWithout(jsoniter.MustGetKind(reflect2.TypeOf(ProxyDefaultTargetGroupSpecConnectionPoolConfig{}).Type1())),
			}.Froze()
			jsonit.Unmarshal(objByte, &obj)

			*(*ProxyDefaultTargetGroupSpecConnectionPoolConfig)(ptr) = obj
		} else {
			*(*ProxyDefaultTargetGroupSpecConnectionPoolConfig)(ptr) = ProxyDefaultTargetGroupSpecConnectionPoolConfig{}
		}
	default:
		iter.ReportError("decode ProxyDefaultTargetGroupSpecConnectionPoolConfig", "unexpected JSON type")
	}
}
