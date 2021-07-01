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
		jsoniter.MustGetKind(reflect2.TypeOf(InstanceSpecCapacityReservationSpecification{}).Type1()):                          InstanceSpecCapacityReservationSpecificationCodec{},
		jsoniter.MustGetKind(reflect2.TypeOf(InstanceSpecCapacityReservationSpecificationCapacityReservationTarget{}).Type1()): InstanceSpecCapacityReservationSpecificationCapacityReservationTargetCodec{},
		jsoniter.MustGetKind(reflect2.TypeOf(InstanceSpecCreditSpecification{}).Type1()):                                       InstanceSpecCreditSpecificationCodec{},
		jsoniter.MustGetKind(reflect2.TypeOf(InstanceSpecEnclaveOptions{}).Type1()):                                            InstanceSpecEnclaveOptionsCodec{},
		jsoniter.MustGetKind(reflect2.TypeOf(InstanceSpecMetadataOptions{}).Type1()):                                           InstanceSpecMetadataOptionsCodec{},
		jsoniter.MustGetKind(reflect2.TypeOf(InstanceSpecRootBlockDevice{}).Type1()):                                           InstanceSpecRootBlockDeviceCodec{},
	}
}

func GetDecoder() map[string]jsoniter.ValDecoder {
	return map[string]jsoniter.ValDecoder{
		jsoniter.MustGetKind(reflect2.TypeOf(InstanceSpecCapacityReservationSpecification{}).Type1()):                          InstanceSpecCapacityReservationSpecificationCodec{},
		jsoniter.MustGetKind(reflect2.TypeOf(InstanceSpecCapacityReservationSpecificationCapacityReservationTarget{}).Type1()): InstanceSpecCapacityReservationSpecificationCapacityReservationTargetCodec{},
		jsoniter.MustGetKind(reflect2.TypeOf(InstanceSpecCreditSpecification{}).Type1()):                                       InstanceSpecCreditSpecificationCodec{},
		jsoniter.MustGetKind(reflect2.TypeOf(InstanceSpecEnclaveOptions{}).Type1()):                                            InstanceSpecEnclaveOptionsCodec{},
		jsoniter.MustGetKind(reflect2.TypeOf(InstanceSpecMetadataOptions{}).Type1()):                                           InstanceSpecMetadataOptionsCodec{},
		jsoniter.MustGetKind(reflect2.TypeOf(InstanceSpecRootBlockDevice{}).Type1()):                                           InstanceSpecRootBlockDeviceCodec{},
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
type InstanceSpecCapacityReservationSpecificationCodec struct {
}

func (InstanceSpecCapacityReservationSpecificationCodec) IsEmpty(ptr unsafe.Pointer) bool {
	return (*InstanceSpecCapacityReservationSpecification)(ptr) == nil
}

func (InstanceSpecCapacityReservationSpecificationCodec) Encode(ptr unsafe.Pointer, stream *jsoniter.Stream) {
	obj := (*InstanceSpecCapacityReservationSpecification)(ptr)
	var objs []InstanceSpecCapacityReservationSpecification
	if obj != nil {
		objs = []InstanceSpecCapacityReservationSpecification{*obj}
	}

	jsonit := jsoniter.Config{
		EscapeHTML:             true,
		SortMapKeys:            true,
		ValidateJsonRawMessage: true,
		TagKey:                 "tf",
		TypeEncoders:           getEncodersWithout(jsoniter.MustGetKind(reflect2.TypeOf(InstanceSpecCapacityReservationSpecification{}).Type1())),
	}.Froze()

	byt, _ := jsonit.Marshal(objs)

	stream.Write(byt)
}

func (InstanceSpecCapacityReservationSpecificationCodec) Decode(ptr unsafe.Pointer, iter *jsoniter.Iterator) {
	switch iter.WhatIsNext() {
	case jsoniter.NilValue:
		iter.Skip()
		*(*InstanceSpecCapacityReservationSpecification)(ptr) = InstanceSpecCapacityReservationSpecification{}
		return
	case jsoniter.ArrayValue:
		objsByte := iter.SkipAndReturnBytes()
		if len(objsByte) > 0 {
			var objs []InstanceSpecCapacityReservationSpecification

			jsonit := jsoniter.Config{
				EscapeHTML:             true,
				SortMapKeys:            true,
				ValidateJsonRawMessage: true,
				TagKey:                 "tf",
				TypeDecoders:           getDecodersWithout(jsoniter.MustGetKind(reflect2.TypeOf(InstanceSpecCapacityReservationSpecification{}).Type1())),
			}.Froze()
			jsonit.Unmarshal(objsByte, &objs)

			if len(objs) > 0 {
				*(*InstanceSpecCapacityReservationSpecification)(ptr) = objs[0]
			} else {
				*(*InstanceSpecCapacityReservationSpecification)(ptr) = InstanceSpecCapacityReservationSpecification{}
			}
		} else {
			*(*InstanceSpecCapacityReservationSpecification)(ptr) = InstanceSpecCapacityReservationSpecification{}
		}
	case jsoniter.ObjectValue:
		objByte := iter.SkipAndReturnBytes()
		if len(objByte) > 0 {
			var obj InstanceSpecCapacityReservationSpecification

			jsonit := jsoniter.Config{
				EscapeHTML:             true,
				SortMapKeys:            true,
				ValidateJsonRawMessage: true,
				TagKey:                 "tf",
				TypeDecoders:           getDecodersWithout(jsoniter.MustGetKind(reflect2.TypeOf(InstanceSpecCapacityReservationSpecification{}).Type1())),
			}.Froze()
			jsonit.Unmarshal(objByte, &obj)

			*(*InstanceSpecCapacityReservationSpecification)(ptr) = obj
		} else {
			*(*InstanceSpecCapacityReservationSpecification)(ptr) = InstanceSpecCapacityReservationSpecification{}
		}
	default:
		iter.ReportError("decode InstanceSpecCapacityReservationSpecification", "unexpected JSON type")
	}
}

// +k8s:deepcopy-gen=false
type InstanceSpecCapacityReservationSpecificationCapacityReservationTargetCodec struct {
}

func (InstanceSpecCapacityReservationSpecificationCapacityReservationTargetCodec) IsEmpty(ptr unsafe.Pointer) bool {
	return (*InstanceSpecCapacityReservationSpecificationCapacityReservationTarget)(ptr) == nil
}

func (InstanceSpecCapacityReservationSpecificationCapacityReservationTargetCodec) Encode(ptr unsafe.Pointer, stream *jsoniter.Stream) {
	obj := (*InstanceSpecCapacityReservationSpecificationCapacityReservationTarget)(ptr)
	var objs []InstanceSpecCapacityReservationSpecificationCapacityReservationTarget
	if obj != nil {
		objs = []InstanceSpecCapacityReservationSpecificationCapacityReservationTarget{*obj}
	}

	jsonit := jsoniter.Config{
		EscapeHTML:             true,
		SortMapKeys:            true,
		ValidateJsonRawMessage: true,
		TagKey:                 "tf",
		TypeEncoders:           getEncodersWithout(jsoniter.MustGetKind(reflect2.TypeOf(InstanceSpecCapacityReservationSpecificationCapacityReservationTarget{}).Type1())),
	}.Froze()

	byt, _ := jsonit.Marshal(objs)

	stream.Write(byt)
}

func (InstanceSpecCapacityReservationSpecificationCapacityReservationTargetCodec) Decode(ptr unsafe.Pointer, iter *jsoniter.Iterator) {
	switch iter.WhatIsNext() {
	case jsoniter.NilValue:
		iter.Skip()
		*(*InstanceSpecCapacityReservationSpecificationCapacityReservationTarget)(ptr) = InstanceSpecCapacityReservationSpecificationCapacityReservationTarget{}
		return
	case jsoniter.ArrayValue:
		objsByte := iter.SkipAndReturnBytes()
		if len(objsByte) > 0 {
			var objs []InstanceSpecCapacityReservationSpecificationCapacityReservationTarget

			jsonit := jsoniter.Config{
				EscapeHTML:             true,
				SortMapKeys:            true,
				ValidateJsonRawMessage: true,
				TagKey:                 "tf",
				TypeDecoders:           getDecodersWithout(jsoniter.MustGetKind(reflect2.TypeOf(InstanceSpecCapacityReservationSpecificationCapacityReservationTarget{}).Type1())),
			}.Froze()
			jsonit.Unmarshal(objsByte, &objs)

			if len(objs) > 0 {
				*(*InstanceSpecCapacityReservationSpecificationCapacityReservationTarget)(ptr) = objs[0]
			} else {
				*(*InstanceSpecCapacityReservationSpecificationCapacityReservationTarget)(ptr) = InstanceSpecCapacityReservationSpecificationCapacityReservationTarget{}
			}
		} else {
			*(*InstanceSpecCapacityReservationSpecificationCapacityReservationTarget)(ptr) = InstanceSpecCapacityReservationSpecificationCapacityReservationTarget{}
		}
	case jsoniter.ObjectValue:
		objByte := iter.SkipAndReturnBytes()
		if len(objByte) > 0 {
			var obj InstanceSpecCapacityReservationSpecificationCapacityReservationTarget

			jsonit := jsoniter.Config{
				EscapeHTML:             true,
				SortMapKeys:            true,
				ValidateJsonRawMessage: true,
				TagKey:                 "tf",
				TypeDecoders:           getDecodersWithout(jsoniter.MustGetKind(reflect2.TypeOf(InstanceSpecCapacityReservationSpecificationCapacityReservationTarget{}).Type1())),
			}.Froze()
			jsonit.Unmarshal(objByte, &obj)

			*(*InstanceSpecCapacityReservationSpecificationCapacityReservationTarget)(ptr) = obj
		} else {
			*(*InstanceSpecCapacityReservationSpecificationCapacityReservationTarget)(ptr) = InstanceSpecCapacityReservationSpecificationCapacityReservationTarget{}
		}
	default:
		iter.ReportError("decode InstanceSpecCapacityReservationSpecificationCapacityReservationTarget", "unexpected JSON type")
	}
}

// +k8s:deepcopy-gen=false
type InstanceSpecCreditSpecificationCodec struct {
}

func (InstanceSpecCreditSpecificationCodec) IsEmpty(ptr unsafe.Pointer) bool {
	return (*InstanceSpecCreditSpecification)(ptr) == nil
}

func (InstanceSpecCreditSpecificationCodec) Encode(ptr unsafe.Pointer, stream *jsoniter.Stream) {
	obj := (*InstanceSpecCreditSpecification)(ptr)
	var objs []InstanceSpecCreditSpecification
	if obj != nil {
		objs = []InstanceSpecCreditSpecification{*obj}
	}

	jsonit := jsoniter.Config{
		EscapeHTML:             true,
		SortMapKeys:            true,
		ValidateJsonRawMessage: true,
		TagKey:                 "tf",
		TypeEncoders:           getEncodersWithout(jsoniter.MustGetKind(reflect2.TypeOf(InstanceSpecCreditSpecification{}).Type1())),
	}.Froze()

	byt, _ := jsonit.Marshal(objs)

	stream.Write(byt)
}

func (InstanceSpecCreditSpecificationCodec) Decode(ptr unsafe.Pointer, iter *jsoniter.Iterator) {
	switch iter.WhatIsNext() {
	case jsoniter.NilValue:
		iter.Skip()
		*(*InstanceSpecCreditSpecification)(ptr) = InstanceSpecCreditSpecification{}
		return
	case jsoniter.ArrayValue:
		objsByte := iter.SkipAndReturnBytes()
		if len(objsByte) > 0 {
			var objs []InstanceSpecCreditSpecification

			jsonit := jsoniter.Config{
				EscapeHTML:             true,
				SortMapKeys:            true,
				ValidateJsonRawMessage: true,
				TagKey:                 "tf",
				TypeDecoders:           getDecodersWithout(jsoniter.MustGetKind(reflect2.TypeOf(InstanceSpecCreditSpecification{}).Type1())),
			}.Froze()
			jsonit.Unmarshal(objsByte, &objs)

			if len(objs) > 0 {
				*(*InstanceSpecCreditSpecification)(ptr) = objs[0]
			} else {
				*(*InstanceSpecCreditSpecification)(ptr) = InstanceSpecCreditSpecification{}
			}
		} else {
			*(*InstanceSpecCreditSpecification)(ptr) = InstanceSpecCreditSpecification{}
		}
	case jsoniter.ObjectValue:
		objByte := iter.SkipAndReturnBytes()
		if len(objByte) > 0 {
			var obj InstanceSpecCreditSpecification

			jsonit := jsoniter.Config{
				EscapeHTML:             true,
				SortMapKeys:            true,
				ValidateJsonRawMessage: true,
				TagKey:                 "tf",
				TypeDecoders:           getDecodersWithout(jsoniter.MustGetKind(reflect2.TypeOf(InstanceSpecCreditSpecification{}).Type1())),
			}.Froze()
			jsonit.Unmarshal(objByte, &obj)

			*(*InstanceSpecCreditSpecification)(ptr) = obj
		} else {
			*(*InstanceSpecCreditSpecification)(ptr) = InstanceSpecCreditSpecification{}
		}
	default:
		iter.ReportError("decode InstanceSpecCreditSpecification", "unexpected JSON type")
	}
}

// +k8s:deepcopy-gen=false
type InstanceSpecEnclaveOptionsCodec struct {
}

func (InstanceSpecEnclaveOptionsCodec) IsEmpty(ptr unsafe.Pointer) bool {
	return (*InstanceSpecEnclaveOptions)(ptr) == nil
}

func (InstanceSpecEnclaveOptionsCodec) Encode(ptr unsafe.Pointer, stream *jsoniter.Stream) {
	obj := (*InstanceSpecEnclaveOptions)(ptr)
	var objs []InstanceSpecEnclaveOptions
	if obj != nil {
		objs = []InstanceSpecEnclaveOptions{*obj}
	}

	jsonit := jsoniter.Config{
		EscapeHTML:             true,
		SortMapKeys:            true,
		ValidateJsonRawMessage: true,
		TagKey:                 "tf",
		TypeEncoders:           getEncodersWithout(jsoniter.MustGetKind(reflect2.TypeOf(InstanceSpecEnclaveOptions{}).Type1())),
	}.Froze()

	byt, _ := jsonit.Marshal(objs)

	stream.Write(byt)
}

func (InstanceSpecEnclaveOptionsCodec) Decode(ptr unsafe.Pointer, iter *jsoniter.Iterator) {
	switch iter.WhatIsNext() {
	case jsoniter.NilValue:
		iter.Skip()
		*(*InstanceSpecEnclaveOptions)(ptr) = InstanceSpecEnclaveOptions{}
		return
	case jsoniter.ArrayValue:
		objsByte := iter.SkipAndReturnBytes()
		if len(objsByte) > 0 {
			var objs []InstanceSpecEnclaveOptions

			jsonit := jsoniter.Config{
				EscapeHTML:             true,
				SortMapKeys:            true,
				ValidateJsonRawMessage: true,
				TagKey:                 "tf",
				TypeDecoders:           getDecodersWithout(jsoniter.MustGetKind(reflect2.TypeOf(InstanceSpecEnclaveOptions{}).Type1())),
			}.Froze()
			jsonit.Unmarshal(objsByte, &objs)

			if len(objs) > 0 {
				*(*InstanceSpecEnclaveOptions)(ptr) = objs[0]
			} else {
				*(*InstanceSpecEnclaveOptions)(ptr) = InstanceSpecEnclaveOptions{}
			}
		} else {
			*(*InstanceSpecEnclaveOptions)(ptr) = InstanceSpecEnclaveOptions{}
		}
	case jsoniter.ObjectValue:
		objByte := iter.SkipAndReturnBytes()
		if len(objByte) > 0 {
			var obj InstanceSpecEnclaveOptions

			jsonit := jsoniter.Config{
				EscapeHTML:             true,
				SortMapKeys:            true,
				ValidateJsonRawMessage: true,
				TagKey:                 "tf",
				TypeDecoders:           getDecodersWithout(jsoniter.MustGetKind(reflect2.TypeOf(InstanceSpecEnclaveOptions{}).Type1())),
			}.Froze()
			jsonit.Unmarshal(objByte, &obj)

			*(*InstanceSpecEnclaveOptions)(ptr) = obj
		} else {
			*(*InstanceSpecEnclaveOptions)(ptr) = InstanceSpecEnclaveOptions{}
		}
	default:
		iter.ReportError("decode InstanceSpecEnclaveOptions", "unexpected JSON type")
	}
}

// +k8s:deepcopy-gen=false
type InstanceSpecMetadataOptionsCodec struct {
}

func (InstanceSpecMetadataOptionsCodec) IsEmpty(ptr unsafe.Pointer) bool {
	return (*InstanceSpecMetadataOptions)(ptr) == nil
}

func (InstanceSpecMetadataOptionsCodec) Encode(ptr unsafe.Pointer, stream *jsoniter.Stream) {
	obj := (*InstanceSpecMetadataOptions)(ptr)
	var objs []InstanceSpecMetadataOptions
	if obj != nil {
		objs = []InstanceSpecMetadataOptions{*obj}
	}

	jsonit := jsoniter.Config{
		EscapeHTML:             true,
		SortMapKeys:            true,
		ValidateJsonRawMessage: true,
		TagKey:                 "tf",
		TypeEncoders:           getEncodersWithout(jsoniter.MustGetKind(reflect2.TypeOf(InstanceSpecMetadataOptions{}).Type1())),
	}.Froze()

	byt, _ := jsonit.Marshal(objs)

	stream.Write(byt)
}

func (InstanceSpecMetadataOptionsCodec) Decode(ptr unsafe.Pointer, iter *jsoniter.Iterator) {
	switch iter.WhatIsNext() {
	case jsoniter.NilValue:
		iter.Skip()
		*(*InstanceSpecMetadataOptions)(ptr) = InstanceSpecMetadataOptions{}
		return
	case jsoniter.ArrayValue:
		objsByte := iter.SkipAndReturnBytes()
		if len(objsByte) > 0 {
			var objs []InstanceSpecMetadataOptions

			jsonit := jsoniter.Config{
				EscapeHTML:             true,
				SortMapKeys:            true,
				ValidateJsonRawMessage: true,
				TagKey:                 "tf",
				TypeDecoders:           getDecodersWithout(jsoniter.MustGetKind(reflect2.TypeOf(InstanceSpecMetadataOptions{}).Type1())),
			}.Froze()
			jsonit.Unmarshal(objsByte, &objs)

			if len(objs) > 0 {
				*(*InstanceSpecMetadataOptions)(ptr) = objs[0]
			} else {
				*(*InstanceSpecMetadataOptions)(ptr) = InstanceSpecMetadataOptions{}
			}
		} else {
			*(*InstanceSpecMetadataOptions)(ptr) = InstanceSpecMetadataOptions{}
		}
	case jsoniter.ObjectValue:
		objByte := iter.SkipAndReturnBytes()
		if len(objByte) > 0 {
			var obj InstanceSpecMetadataOptions

			jsonit := jsoniter.Config{
				EscapeHTML:             true,
				SortMapKeys:            true,
				ValidateJsonRawMessage: true,
				TagKey:                 "tf",
				TypeDecoders:           getDecodersWithout(jsoniter.MustGetKind(reflect2.TypeOf(InstanceSpecMetadataOptions{}).Type1())),
			}.Froze()
			jsonit.Unmarshal(objByte, &obj)

			*(*InstanceSpecMetadataOptions)(ptr) = obj
		} else {
			*(*InstanceSpecMetadataOptions)(ptr) = InstanceSpecMetadataOptions{}
		}
	default:
		iter.ReportError("decode InstanceSpecMetadataOptions", "unexpected JSON type")
	}
}

// +k8s:deepcopy-gen=false
type InstanceSpecRootBlockDeviceCodec struct {
}

func (InstanceSpecRootBlockDeviceCodec) IsEmpty(ptr unsafe.Pointer) bool {
	return (*InstanceSpecRootBlockDevice)(ptr) == nil
}

func (InstanceSpecRootBlockDeviceCodec) Encode(ptr unsafe.Pointer, stream *jsoniter.Stream) {
	obj := (*InstanceSpecRootBlockDevice)(ptr)
	var objs []InstanceSpecRootBlockDevice
	if obj != nil {
		objs = []InstanceSpecRootBlockDevice{*obj}
	}

	jsonit := jsoniter.Config{
		EscapeHTML:             true,
		SortMapKeys:            true,
		ValidateJsonRawMessage: true,
		TagKey:                 "tf",
		TypeEncoders:           getEncodersWithout(jsoniter.MustGetKind(reflect2.TypeOf(InstanceSpecRootBlockDevice{}).Type1())),
	}.Froze()

	byt, _ := jsonit.Marshal(objs)

	stream.Write(byt)
}

func (InstanceSpecRootBlockDeviceCodec) Decode(ptr unsafe.Pointer, iter *jsoniter.Iterator) {
	switch iter.WhatIsNext() {
	case jsoniter.NilValue:
		iter.Skip()
		*(*InstanceSpecRootBlockDevice)(ptr) = InstanceSpecRootBlockDevice{}
		return
	case jsoniter.ArrayValue:
		objsByte := iter.SkipAndReturnBytes()
		if len(objsByte) > 0 {
			var objs []InstanceSpecRootBlockDevice

			jsonit := jsoniter.Config{
				EscapeHTML:             true,
				SortMapKeys:            true,
				ValidateJsonRawMessage: true,
				TagKey:                 "tf",
				TypeDecoders:           getDecodersWithout(jsoniter.MustGetKind(reflect2.TypeOf(InstanceSpecRootBlockDevice{}).Type1())),
			}.Froze()
			jsonit.Unmarshal(objsByte, &objs)

			if len(objs) > 0 {
				*(*InstanceSpecRootBlockDevice)(ptr) = objs[0]
			} else {
				*(*InstanceSpecRootBlockDevice)(ptr) = InstanceSpecRootBlockDevice{}
			}
		} else {
			*(*InstanceSpecRootBlockDevice)(ptr) = InstanceSpecRootBlockDevice{}
		}
	case jsoniter.ObjectValue:
		objByte := iter.SkipAndReturnBytes()
		if len(objByte) > 0 {
			var obj InstanceSpecRootBlockDevice

			jsonit := jsoniter.Config{
				EscapeHTML:             true,
				SortMapKeys:            true,
				ValidateJsonRawMessage: true,
				TagKey:                 "tf",
				TypeDecoders:           getDecodersWithout(jsoniter.MustGetKind(reflect2.TypeOf(InstanceSpecRootBlockDevice{}).Type1())),
			}.Froze()
			jsonit.Unmarshal(objByte, &obj)

			*(*InstanceSpecRootBlockDevice)(ptr) = obj
		} else {
			*(*InstanceSpecRootBlockDevice)(ptr) = InstanceSpecRootBlockDevice{}
		}
	default:
		iter.ReportError("decode InstanceSpecRootBlockDevice", "unexpected JSON type")
	}
}
