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
		jsoniter.MustGetKind(reflect2.TypeOf(ClientVPNEndpointSpecConnectionLogOptions{}).Type1()):                  ClientVPNEndpointSpecConnectionLogOptionsCodec{},
		jsoniter.MustGetKind(reflect2.TypeOf(FleetSpecLaunchTemplateConfig{}).Type1()):                              FleetSpecLaunchTemplateConfigCodec{},
		jsoniter.MustGetKind(reflect2.TypeOf(FleetSpecLaunchTemplateConfigLaunchTemplateSpecification{}).Type1()):   FleetSpecLaunchTemplateConfigLaunchTemplateSpecificationCodec{},
		jsoniter.MustGetKind(reflect2.TypeOf(FleetSpecOnDemandOptions{}).Type1()):                                   FleetSpecOnDemandOptionsCodec{},
		jsoniter.MustGetKind(reflect2.TypeOf(FleetSpecSpotOptions{}).Type1()):                                       FleetSpecSpotOptionsCodec{},
		jsoniter.MustGetKind(reflect2.TypeOf(FleetSpecSpotOptionsMaintenanceStrategies{}).Type1()):                  FleetSpecSpotOptionsMaintenanceStrategiesCodec{},
		jsoniter.MustGetKind(reflect2.TypeOf(FleetSpecSpotOptionsMaintenanceStrategiesCapacityRebalance{}).Type1()): FleetSpecSpotOptionsMaintenanceStrategiesCapacityRebalanceCodec{},
		jsoniter.MustGetKind(reflect2.TypeOf(FleetSpecTargetCapacitySpecification{}).Type1()):                       FleetSpecTargetCapacitySpecificationCodec{},
		jsoniter.MustGetKind(reflect2.TypeOf(TrafficMirrorFilterRuleSpecDestinationPortRange{}).Type1()):            TrafficMirrorFilterRuleSpecDestinationPortRangeCodec{},
		jsoniter.MustGetKind(reflect2.TypeOf(TrafficMirrorFilterRuleSpecSourcePortRange{}).Type1()):                 TrafficMirrorFilterRuleSpecSourcePortRangeCodec{},
	}
}

func GetDecoder() map[string]jsoniter.ValDecoder {
	return map[string]jsoniter.ValDecoder{
		jsoniter.MustGetKind(reflect2.TypeOf(ClientVPNEndpointSpecConnectionLogOptions{}).Type1()):                  ClientVPNEndpointSpecConnectionLogOptionsCodec{},
		jsoniter.MustGetKind(reflect2.TypeOf(FleetSpecLaunchTemplateConfig{}).Type1()):                              FleetSpecLaunchTemplateConfigCodec{},
		jsoniter.MustGetKind(reflect2.TypeOf(FleetSpecLaunchTemplateConfigLaunchTemplateSpecification{}).Type1()):   FleetSpecLaunchTemplateConfigLaunchTemplateSpecificationCodec{},
		jsoniter.MustGetKind(reflect2.TypeOf(FleetSpecOnDemandOptions{}).Type1()):                                   FleetSpecOnDemandOptionsCodec{},
		jsoniter.MustGetKind(reflect2.TypeOf(FleetSpecSpotOptions{}).Type1()):                                       FleetSpecSpotOptionsCodec{},
		jsoniter.MustGetKind(reflect2.TypeOf(FleetSpecSpotOptionsMaintenanceStrategies{}).Type1()):                  FleetSpecSpotOptionsMaintenanceStrategiesCodec{},
		jsoniter.MustGetKind(reflect2.TypeOf(FleetSpecSpotOptionsMaintenanceStrategiesCapacityRebalance{}).Type1()): FleetSpecSpotOptionsMaintenanceStrategiesCapacityRebalanceCodec{},
		jsoniter.MustGetKind(reflect2.TypeOf(FleetSpecTargetCapacitySpecification{}).Type1()):                       FleetSpecTargetCapacitySpecificationCodec{},
		jsoniter.MustGetKind(reflect2.TypeOf(TrafficMirrorFilterRuleSpecDestinationPortRange{}).Type1()):            TrafficMirrorFilterRuleSpecDestinationPortRangeCodec{},
		jsoniter.MustGetKind(reflect2.TypeOf(TrafficMirrorFilterRuleSpecSourcePortRange{}).Type1()):                 TrafficMirrorFilterRuleSpecSourcePortRangeCodec{},
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
type ClientVPNEndpointSpecConnectionLogOptionsCodec struct {
}

func (ClientVPNEndpointSpecConnectionLogOptionsCodec) IsEmpty(ptr unsafe.Pointer) bool {
	return (*ClientVPNEndpointSpecConnectionLogOptions)(ptr) == nil
}

func (ClientVPNEndpointSpecConnectionLogOptionsCodec) Encode(ptr unsafe.Pointer, stream *jsoniter.Stream) {
	obj := (*ClientVPNEndpointSpecConnectionLogOptions)(ptr)
	var objs []ClientVPNEndpointSpecConnectionLogOptions
	if obj != nil {
		objs = []ClientVPNEndpointSpecConnectionLogOptions{*obj}
	}

	jsonit := jsoniter.Config{
		EscapeHTML:             true,
		SortMapKeys:            true,
		ValidateJsonRawMessage: true,
		TagKey:                 "tf",
		TypeEncoders:           getEncodersWithout(jsoniter.MustGetKind(reflect2.TypeOf(ClientVPNEndpointSpecConnectionLogOptions{}).Type1())),
	}.Froze()

	byt, _ := jsonit.Marshal(objs)

	stream.Write(byt)
}

func (ClientVPNEndpointSpecConnectionLogOptionsCodec) Decode(ptr unsafe.Pointer, iter *jsoniter.Iterator) {
	switch iter.WhatIsNext() {
	case jsoniter.NilValue:
		iter.Skip()
		*(*ClientVPNEndpointSpecConnectionLogOptions)(ptr) = ClientVPNEndpointSpecConnectionLogOptions{}
		return
	case jsoniter.ArrayValue:
		objsByte := iter.SkipAndReturnBytes()
		if len(objsByte) > 0 {
			var objs []ClientVPNEndpointSpecConnectionLogOptions

			jsonit := jsoniter.Config{
				EscapeHTML:             true,
				SortMapKeys:            true,
				ValidateJsonRawMessage: true,
				TagKey:                 "tf",
				TypeDecoders:           getDecodersWithout(jsoniter.MustGetKind(reflect2.TypeOf(ClientVPNEndpointSpecConnectionLogOptions{}).Type1())),
			}.Froze()
			jsonit.Unmarshal(objsByte, &objs)

			if len(objs) > 0 {
				*(*ClientVPNEndpointSpecConnectionLogOptions)(ptr) = objs[0]
			} else {
				*(*ClientVPNEndpointSpecConnectionLogOptions)(ptr) = ClientVPNEndpointSpecConnectionLogOptions{}
			}
		} else {
			*(*ClientVPNEndpointSpecConnectionLogOptions)(ptr) = ClientVPNEndpointSpecConnectionLogOptions{}
		}
	case jsoniter.ObjectValue:
		objByte := iter.SkipAndReturnBytes()
		if len(objByte) > 0 {
			var obj ClientVPNEndpointSpecConnectionLogOptions

			jsonit := jsoniter.Config{
				EscapeHTML:             true,
				SortMapKeys:            true,
				ValidateJsonRawMessage: true,
				TagKey:                 "tf",
				TypeDecoders:           getDecodersWithout(jsoniter.MustGetKind(reflect2.TypeOf(ClientVPNEndpointSpecConnectionLogOptions{}).Type1())),
			}.Froze()
			jsonit.Unmarshal(objByte, &obj)

			*(*ClientVPNEndpointSpecConnectionLogOptions)(ptr) = obj
		} else {
			*(*ClientVPNEndpointSpecConnectionLogOptions)(ptr) = ClientVPNEndpointSpecConnectionLogOptions{}
		}
	default:
		iter.ReportError("decode ClientVPNEndpointSpecConnectionLogOptions", "unexpected JSON type")
	}
}

// +k8s:deepcopy-gen=false
type FleetSpecLaunchTemplateConfigCodec struct {
}

func (FleetSpecLaunchTemplateConfigCodec) IsEmpty(ptr unsafe.Pointer) bool {
	return (*FleetSpecLaunchTemplateConfig)(ptr) == nil
}

func (FleetSpecLaunchTemplateConfigCodec) Encode(ptr unsafe.Pointer, stream *jsoniter.Stream) {
	obj := (*FleetSpecLaunchTemplateConfig)(ptr)
	var objs []FleetSpecLaunchTemplateConfig
	if obj != nil {
		objs = []FleetSpecLaunchTemplateConfig{*obj}
	}

	jsonit := jsoniter.Config{
		EscapeHTML:             true,
		SortMapKeys:            true,
		ValidateJsonRawMessage: true,
		TagKey:                 "tf",
		TypeEncoders:           getEncodersWithout(jsoniter.MustGetKind(reflect2.TypeOf(FleetSpecLaunchTemplateConfig{}).Type1())),
	}.Froze()

	byt, _ := jsonit.Marshal(objs)

	stream.Write(byt)
}

func (FleetSpecLaunchTemplateConfigCodec) Decode(ptr unsafe.Pointer, iter *jsoniter.Iterator) {
	switch iter.WhatIsNext() {
	case jsoniter.NilValue:
		iter.Skip()
		*(*FleetSpecLaunchTemplateConfig)(ptr) = FleetSpecLaunchTemplateConfig{}
		return
	case jsoniter.ArrayValue:
		objsByte := iter.SkipAndReturnBytes()
		if len(objsByte) > 0 {
			var objs []FleetSpecLaunchTemplateConfig

			jsonit := jsoniter.Config{
				EscapeHTML:             true,
				SortMapKeys:            true,
				ValidateJsonRawMessage: true,
				TagKey:                 "tf",
				TypeDecoders:           getDecodersWithout(jsoniter.MustGetKind(reflect2.TypeOf(FleetSpecLaunchTemplateConfig{}).Type1())),
			}.Froze()
			jsonit.Unmarshal(objsByte, &objs)

			if len(objs) > 0 {
				*(*FleetSpecLaunchTemplateConfig)(ptr) = objs[0]
			} else {
				*(*FleetSpecLaunchTemplateConfig)(ptr) = FleetSpecLaunchTemplateConfig{}
			}
		} else {
			*(*FleetSpecLaunchTemplateConfig)(ptr) = FleetSpecLaunchTemplateConfig{}
		}
	case jsoniter.ObjectValue:
		objByte := iter.SkipAndReturnBytes()
		if len(objByte) > 0 {
			var obj FleetSpecLaunchTemplateConfig

			jsonit := jsoniter.Config{
				EscapeHTML:             true,
				SortMapKeys:            true,
				ValidateJsonRawMessage: true,
				TagKey:                 "tf",
				TypeDecoders:           getDecodersWithout(jsoniter.MustGetKind(reflect2.TypeOf(FleetSpecLaunchTemplateConfig{}).Type1())),
			}.Froze()
			jsonit.Unmarshal(objByte, &obj)

			*(*FleetSpecLaunchTemplateConfig)(ptr) = obj
		} else {
			*(*FleetSpecLaunchTemplateConfig)(ptr) = FleetSpecLaunchTemplateConfig{}
		}
	default:
		iter.ReportError("decode FleetSpecLaunchTemplateConfig", "unexpected JSON type")
	}
}

// +k8s:deepcopy-gen=false
type FleetSpecLaunchTemplateConfigLaunchTemplateSpecificationCodec struct {
}

func (FleetSpecLaunchTemplateConfigLaunchTemplateSpecificationCodec) IsEmpty(ptr unsafe.Pointer) bool {
	return (*FleetSpecLaunchTemplateConfigLaunchTemplateSpecification)(ptr) == nil
}

func (FleetSpecLaunchTemplateConfigLaunchTemplateSpecificationCodec) Encode(ptr unsafe.Pointer, stream *jsoniter.Stream) {
	obj := (*FleetSpecLaunchTemplateConfigLaunchTemplateSpecification)(ptr)
	var objs []FleetSpecLaunchTemplateConfigLaunchTemplateSpecification
	if obj != nil {
		objs = []FleetSpecLaunchTemplateConfigLaunchTemplateSpecification{*obj}
	}

	jsonit := jsoniter.Config{
		EscapeHTML:             true,
		SortMapKeys:            true,
		ValidateJsonRawMessage: true,
		TagKey:                 "tf",
		TypeEncoders:           getEncodersWithout(jsoniter.MustGetKind(reflect2.TypeOf(FleetSpecLaunchTemplateConfigLaunchTemplateSpecification{}).Type1())),
	}.Froze()

	byt, _ := jsonit.Marshal(objs)

	stream.Write(byt)
}

func (FleetSpecLaunchTemplateConfigLaunchTemplateSpecificationCodec) Decode(ptr unsafe.Pointer, iter *jsoniter.Iterator) {
	switch iter.WhatIsNext() {
	case jsoniter.NilValue:
		iter.Skip()
		*(*FleetSpecLaunchTemplateConfigLaunchTemplateSpecification)(ptr) = FleetSpecLaunchTemplateConfigLaunchTemplateSpecification{}
		return
	case jsoniter.ArrayValue:
		objsByte := iter.SkipAndReturnBytes()
		if len(objsByte) > 0 {
			var objs []FleetSpecLaunchTemplateConfigLaunchTemplateSpecification

			jsonit := jsoniter.Config{
				EscapeHTML:             true,
				SortMapKeys:            true,
				ValidateJsonRawMessage: true,
				TagKey:                 "tf",
				TypeDecoders:           getDecodersWithout(jsoniter.MustGetKind(reflect2.TypeOf(FleetSpecLaunchTemplateConfigLaunchTemplateSpecification{}).Type1())),
			}.Froze()
			jsonit.Unmarshal(objsByte, &objs)

			if len(objs) > 0 {
				*(*FleetSpecLaunchTemplateConfigLaunchTemplateSpecification)(ptr) = objs[0]
			} else {
				*(*FleetSpecLaunchTemplateConfigLaunchTemplateSpecification)(ptr) = FleetSpecLaunchTemplateConfigLaunchTemplateSpecification{}
			}
		} else {
			*(*FleetSpecLaunchTemplateConfigLaunchTemplateSpecification)(ptr) = FleetSpecLaunchTemplateConfigLaunchTemplateSpecification{}
		}
	case jsoniter.ObjectValue:
		objByte := iter.SkipAndReturnBytes()
		if len(objByte) > 0 {
			var obj FleetSpecLaunchTemplateConfigLaunchTemplateSpecification

			jsonit := jsoniter.Config{
				EscapeHTML:             true,
				SortMapKeys:            true,
				ValidateJsonRawMessage: true,
				TagKey:                 "tf",
				TypeDecoders:           getDecodersWithout(jsoniter.MustGetKind(reflect2.TypeOf(FleetSpecLaunchTemplateConfigLaunchTemplateSpecification{}).Type1())),
			}.Froze()
			jsonit.Unmarshal(objByte, &obj)

			*(*FleetSpecLaunchTemplateConfigLaunchTemplateSpecification)(ptr) = obj
		} else {
			*(*FleetSpecLaunchTemplateConfigLaunchTemplateSpecification)(ptr) = FleetSpecLaunchTemplateConfigLaunchTemplateSpecification{}
		}
	default:
		iter.ReportError("decode FleetSpecLaunchTemplateConfigLaunchTemplateSpecification", "unexpected JSON type")
	}
}

// +k8s:deepcopy-gen=false
type FleetSpecOnDemandOptionsCodec struct {
}

func (FleetSpecOnDemandOptionsCodec) IsEmpty(ptr unsafe.Pointer) bool {
	return (*FleetSpecOnDemandOptions)(ptr) == nil
}

func (FleetSpecOnDemandOptionsCodec) Encode(ptr unsafe.Pointer, stream *jsoniter.Stream) {
	obj := (*FleetSpecOnDemandOptions)(ptr)
	var objs []FleetSpecOnDemandOptions
	if obj != nil {
		objs = []FleetSpecOnDemandOptions{*obj}
	}

	jsonit := jsoniter.Config{
		EscapeHTML:             true,
		SortMapKeys:            true,
		ValidateJsonRawMessage: true,
		TagKey:                 "tf",
		TypeEncoders:           getEncodersWithout(jsoniter.MustGetKind(reflect2.TypeOf(FleetSpecOnDemandOptions{}).Type1())),
	}.Froze()

	byt, _ := jsonit.Marshal(objs)

	stream.Write(byt)
}

func (FleetSpecOnDemandOptionsCodec) Decode(ptr unsafe.Pointer, iter *jsoniter.Iterator) {
	switch iter.WhatIsNext() {
	case jsoniter.NilValue:
		iter.Skip()
		*(*FleetSpecOnDemandOptions)(ptr) = FleetSpecOnDemandOptions{}
		return
	case jsoniter.ArrayValue:
		objsByte := iter.SkipAndReturnBytes()
		if len(objsByte) > 0 {
			var objs []FleetSpecOnDemandOptions

			jsonit := jsoniter.Config{
				EscapeHTML:             true,
				SortMapKeys:            true,
				ValidateJsonRawMessage: true,
				TagKey:                 "tf",
				TypeDecoders:           getDecodersWithout(jsoniter.MustGetKind(reflect2.TypeOf(FleetSpecOnDemandOptions{}).Type1())),
			}.Froze()
			jsonit.Unmarshal(objsByte, &objs)

			if len(objs) > 0 {
				*(*FleetSpecOnDemandOptions)(ptr) = objs[0]
			} else {
				*(*FleetSpecOnDemandOptions)(ptr) = FleetSpecOnDemandOptions{}
			}
		} else {
			*(*FleetSpecOnDemandOptions)(ptr) = FleetSpecOnDemandOptions{}
		}
	case jsoniter.ObjectValue:
		objByte := iter.SkipAndReturnBytes()
		if len(objByte) > 0 {
			var obj FleetSpecOnDemandOptions

			jsonit := jsoniter.Config{
				EscapeHTML:             true,
				SortMapKeys:            true,
				ValidateJsonRawMessage: true,
				TagKey:                 "tf",
				TypeDecoders:           getDecodersWithout(jsoniter.MustGetKind(reflect2.TypeOf(FleetSpecOnDemandOptions{}).Type1())),
			}.Froze()
			jsonit.Unmarshal(objByte, &obj)

			*(*FleetSpecOnDemandOptions)(ptr) = obj
		} else {
			*(*FleetSpecOnDemandOptions)(ptr) = FleetSpecOnDemandOptions{}
		}
	default:
		iter.ReportError("decode FleetSpecOnDemandOptions", "unexpected JSON type")
	}
}

// +k8s:deepcopy-gen=false
type FleetSpecSpotOptionsCodec struct {
}

func (FleetSpecSpotOptionsCodec) IsEmpty(ptr unsafe.Pointer) bool {
	return (*FleetSpecSpotOptions)(ptr) == nil
}

func (FleetSpecSpotOptionsCodec) Encode(ptr unsafe.Pointer, stream *jsoniter.Stream) {
	obj := (*FleetSpecSpotOptions)(ptr)
	var objs []FleetSpecSpotOptions
	if obj != nil {
		objs = []FleetSpecSpotOptions{*obj}
	}

	jsonit := jsoniter.Config{
		EscapeHTML:             true,
		SortMapKeys:            true,
		ValidateJsonRawMessage: true,
		TagKey:                 "tf",
		TypeEncoders:           getEncodersWithout(jsoniter.MustGetKind(reflect2.TypeOf(FleetSpecSpotOptions{}).Type1())),
	}.Froze()

	byt, _ := jsonit.Marshal(objs)

	stream.Write(byt)
}

func (FleetSpecSpotOptionsCodec) Decode(ptr unsafe.Pointer, iter *jsoniter.Iterator) {
	switch iter.WhatIsNext() {
	case jsoniter.NilValue:
		iter.Skip()
		*(*FleetSpecSpotOptions)(ptr) = FleetSpecSpotOptions{}
		return
	case jsoniter.ArrayValue:
		objsByte := iter.SkipAndReturnBytes()
		if len(objsByte) > 0 {
			var objs []FleetSpecSpotOptions

			jsonit := jsoniter.Config{
				EscapeHTML:             true,
				SortMapKeys:            true,
				ValidateJsonRawMessage: true,
				TagKey:                 "tf",
				TypeDecoders:           getDecodersWithout(jsoniter.MustGetKind(reflect2.TypeOf(FleetSpecSpotOptions{}).Type1())),
			}.Froze()
			jsonit.Unmarshal(objsByte, &objs)

			if len(objs) > 0 {
				*(*FleetSpecSpotOptions)(ptr) = objs[0]
			} else {
				*(*FleetSpecSpotOptions)(ptr) = FleetSpecSpotOptions{}
			}
		} else {
			*(*FleetSpecSpotOptions)(ptr) = FleetSpecSpotOptions{}
		}
	case jsoniter.ObjectValue:
		objByte := iter.SkipAndReturnBytes()
		if len(objByte) > 0 {
			var obj FleetSpecSpotOptions

			jsonit := jsoniter.Config{
				EscapeHTML:             true,
				SortMapKeys:            true,
				ValidateJsonRawMessage: true,
				TagKey:                 "tf",
				TypeDecoders:           getDecodersWithout(jsoniter.MustGetKind(reflect2.TypeOf(FleetSpecSpotOptions{}).Type1())),
			}.Froze()
			jsonit.Unmarshal(objByte, &obj)

			*(*FleetSpecSpotOptions)(ptr) = obj
		} else {
			*(*FleetSpecSpotOptions)(ptr) = FleetSpecSpotOptions{}
		}
	default:
		iter.ReportError("decode FleetSpecSpotOptions", "unexpected JSON type")
	}
}

// +k8s:deepcopy-gen=false
type FleetSpecSpotOptionsMaintenanceStrategiesCodec struct {
}

func (FleetSpecSpotOptionsMaintenanceStrategiesCodec) IsEmpty(ptr unsafe.Pointer) bool {
	return (*FleetSpecSpotOptionsMaintenanceStrategies)(ptr) == nil
}

func (FleetSpecSpotOptionsMaintenanceStrategiesCodec) Encode(ptr unsafe.Pointer, stream *jsoniter.Stream) {
	obj := (*FleetSpecSpotOptionsMaintenanceStrategies)(ptr)
	var objs []FleetSpecSpotOptionsMaintenanceStrategies
	if obj != nil {
		objs = []FleetSpecSpotOptionsMaintenanceStrategies{*obj}
	}

	jsonit := jsoniter.Config{
		EscapeHTML:             true,
		SortMapKeys:            true,
		ValidateJsonRawMessage: true,
		TagKey:                 "tf",
		TypeEncoders:           getEncodersWithout(jsoniter.MustGetKind(reflect2.TypeOf(FleetSpecSpotOptionsMaintenanceStrategies{}).Type1())),
	}.Froze()

	byt, _ := jsonit.Marshal(objs)

	stream.Write(byt)
}

func (FleetSpecSpotOptionsMaintenanceStrategiesCodec) Decode(ptr unsafe.Pointer, iter *jsoniter.Iterator) {
	switch iter.WhatIsNext() {
	case jsoniter.NilValue:
		iter.Skip()
		*(*FleetSpecSpotOptionsMaintenanceStrategies)(ptr) = FleetSpecSpotOptionsMaintenanceStrategies{}
		return
	case jsoniter.ArrayValue:
		objsByte := iter.SkipAndReturnBytes()
		if len(objsByte) > 0 {
			var objs []FleetSpecSpotOptionsMaintenanceStrategies

			jsonit := jsoniter.Config{
				EscapeHTML:             true,
				SortMapKeys:            true,
				ValidateJsonRawMessage: true,
				TagKey:                 "tf",
				TypeDecoders:           getDecodersWithout(jsoniter.MustGetKind(reflect2.TypeOf(FleetSpecSpotOptionsMaintenanceStrategies{}).Type1())),
			}.Froze()
			jsonit.Unmarshal(objsByte, &objs)

			if len(objs) > 0 {
				*(*FleetSpecSpotOptionsMaintenanceStrategies)(ptr) = objs[0]
			} else {
				*(*FleetSpecSpotOptionsMaintenanceStrategies)(ptr) = FleetSpecSpotOptionsMaintenanceStrategies{}
			}
		} else {
			*(*FleetSpecSpotOptionsMaintenanceStrategies)(ptr) = FleetSpecSpotOptionsMaintenanceStrategies{}
		}
	case jsoniter.ObjectValue:
		objByte := iter.SkipAndReturnBytes()
		if len(objByte) > 0 {
			var obj FleetSpecSpotOptionsMaintenanceStrategies

			jsonit := jsoniter.Config{
				EscapeHTML:             true,
				SortMapKeys:            true,
				ValidateJsonRawMessage: true,
				TagKey:                 "tf",
				TypeDecoders:           getDecodersWithout(jsoniter.MustGetKind(reflect2.TypeOf(FleetSpecSpotOptionsMaintenanceStrategies{}).Type1())),
			}.Froze()
			jsonit.Unmarshal(objByte, &obj)

			*(*FleetSpecSpotOptionsMaintenanceStrategies)(ptr) = obj
		} else {
			*(*FleetSpecSpotOptionsMaintenanceStrategies)(ptr) = FleetSpecSpotOptionsMaintenanceStrategies{}
		}
	default:
		iter.ReportError("decode FleetSpecSpotOptionsMaintenanceStrategies", "unexpected JSON type")
	}
}

// +k8s:deepcopy-gen=false
type FleetSpecSpotOptionsMaintenanceStrategiesCapacityRebalanceCodec struct {
}

func (FleetSpecSpotOptionsMaintenanceStrategiesCapacityRebalanceCodec) IsEmpty(ptr unsafe.Pointer) bool {
	return (*FleetSpecSpotOptionsMaintenanceStrategiesCapacityRebalance)(ptr) == nil
}

func (FleetSpecSpotOptionsMaintenanceStrategiesCapacityRebalanceCodec) Encode(ptr unsafe.Pointer, stream *jsoniter.Stream) {
	obj := (*FleetSpecSpotOptionsMaintenanceStrategiesCapacityRebalance)(ptr)
	var objs []FleetSpecSpotOptionsMaintenanceStrategiesCapacityRebalance
	if obj != nil {
		objs = []FleetSpecSpotOptionsMaintenanceStrategiesCapacityRebalance{*obj}
	}

	jsonit := jsoniter.Config{
		EscapeHTML:             true,
		SortMapKeys:            true,
		ValidateJsonRawMessage: true,
		TagKey:                 "tf",
		TypeEncoders:           getEncodersWithout(jsoniter.MustGetKind(reflect2.TypeOf(FleetSpecSpotOptionsMaintenanceStrategiesCapacityRebalance{}).Type1())),
	}.Froze()

	byt, _ := jsonit.Marshal(objs)

	stream.Write(byt)
}

func (FleetSpecSpotOptionsMaintenanceStrategiesCapacityRebalanceCodec) Decode(ptr unsafe.Pointer, iter *jsoniter.Iterator) {
	switch iter.WhatIsNext() {
	case jsoniter.NilValue:
		iter.Skip()
		*(*FleetSpecSpotOptionsMaintenanceStrategiesCapacityRebalance)(ptr) = FleetSpecSpotOptionsMaintenanceStrategiesCapacityRebalance{}
		return
	case jsoniter.ArrayValue:
		objsByte := iter.SkipAndReturnBytes()
		if len(objsByte) > 0 {
			var objs []FleetSpecSpotOptionsMaintenanceStrategiesCapacityRebalance

			jsonit := jsoniter.Config{
				EscapeHTML:             true,
				SortMapKeys:            true,
				ValidateJsonRawMessage: true,
				TagKey:                 "tf",
				TypeDecoders:           getDecodersWithout(jsoniter.MustGetKind(reflect2.TypeOf(FleetSpecSpotOptionsMaintenanceStrategiesCapacityRebalance{}).Type1())),
			}.Froze()
			jsonit.Unmarshal(objsByte, &objs)

			if len(objs) > 0 {
				*(*FleetSpecSpotOptionsMaintenanceStrategiesCapacityRebalance)(ptr) = objs[0]
			} else {
				*(*FleetSpecSpotOptionsMaintenanceStrategiesCapacityRebalance)(ptr) = FleetSpecSpotOptionsMaintenanceStrategiesCapacityRebalance{}
			}
		} else {
			*(*FleetSpecSpotOptionsMaintenanceStrategiesCapacityRebalance)(ptr) = FleetSpecSpotOptionsMaintenanceStrategiesCapacityRebalance{}
		}
	case jsoniter.ObjectValue:
		objByte := iter.SkipAndReturnBytes()
		if len(objByte) > 0 {
			var obj FleetSpecSpotOptionsMaintenanceStrategiesCapacityRebalance

			jsonit := jsoniter.Config{
				EscapeHTML:             true,
				SortMapKeys:            true,
				ValidateJsonRawMessage: true,
				TagKey:                 "tf",
				TypeDecoders:           getDecodersWithout(jsoniter.MustGetKind(reflect2.TypeOf(FleetSpecSpotOptionsMaintenanceStrategiesCapacityRebalance{}).Type1())),
			}.Froze()
			jsonit.Unmarshal(objByte, &obj)

			*(*FleetSpecSpotOptionsMaintenanceStrategiesCapacityRebalance)(ptr) = obj
		} else {
			*(*FleetSpecSpotOptionsMaintenanceStrategiesCapacityRebalance)(ptr) = FleetSpecSpotOptionsMaintenanceStrategiesCapacityRebalance{}
		}
	default:
		iter.ReportError("decode FleetSpecSpotOptionsMaintenanceStrategiesCapacityRebalance", "unexpected JSON type")
	}
}

// +k8s:deepcopy-gen=false
type FleetSpecTargetCapacitySpecificationCodec struct {
}

func (FleetSpecTargetCapacitySpecificationCodec) IsEmpty(ptr unsafe.Pointer) bool {
	return (*FleetSpecTargetCapacitySpecification)(ptr) == nil
}

func (FleetSpecTargetCapacitySpecificationCodec) Encode(ptr unsafe.Pointer, stream *jsoniter.Stream) {
	obj := (*FleetSpecTargetCapacitySpecification)(ptr)
	var objs []FleetSpecTargetCapacitySpecification
	if obj != nil {
		objs = []FleetSpecTargetCapacitySpecification{*obj}
	}

	jsonit := jsoniter.Config{
		EscapeHTML:             true,
		SortMapKeys:            true,
		ValidateJsonRawMessage: true,
		TagKey:                 "tf",
		TypeEncoders:           getEncodersWithout(jsoniter.MustGetKind(reflect2.TypeOf(FleetSpecTargetCapacitySpecification{}).Type1())),
	}.Froze()

	byt, _ := jsonit.Marshal(objs)

	stream.Write(byt)
}

func (FleetSpecTargetCapacitySpecificationCodec) Decode(ptr unsafe.Pointer, iter *jsoniter.Iterator) {
	switch iter.WhatIsNext() {
	case jsoniter.NilValue:
		iter.Skip()
		*(*FleetSpecTargetCapacitySpecification)(ptr) = FleetSpecTargetCapacitySpecification{}
		return
	case jsoniter.ArrayValue:
		objsByte := iter.SkipAndReturnBytes()
		if len(objsByte) > 0 {
			var objs []FleetSpecTargetCapacitySpecification

			jsonit := jsoniter.Config{
				EscapeHTML:             true,
				SortMapKeys:            true,
				ValidateJsonRawMessage: true,
				TagKey:                 "tf",
				TypeDecoders:           getDecodersWithout(jsoniter.MustGetKind(reflect2.TypeOf(FleetSpecTargetCapacitySpecification{}).Type1())),
			}.Froze()
			jsonit.Unmarshal(objsByte, &objs)

			if len(objs) > 0 {
				*(*FleetSpecTargetCapacitySpecification)(ptr) = objs[0]
			} else {
				*(*FleetSpecTargetCapacitySpecification)(ptr) = FleetSpecTargetCapacitySpecification{}
			}
		} else {
			*(*FleetSpecTargetCapacitySpecification)(ptr) = FleetSpecTargetCapacitySpecification{}
		}
	case jsoniter.ObjectValue:
		objByte := iter.SkipAndReturnBytes()
		if len(objByte) > 0 {
			var obj FleetSpecTargetCapacitySpecification

			jsonit := jsoniter.Config{
				EscapeHTML:             true,
				SortMapKeys:            true,
				ValidateJsonRawMessage: true,
				TagKey:                 "tf",
				TypeDecoders:           getDecodersWithout(jsoniter.MustGetKind(reflect2.TypeOf(FleetSpecTargetCapacitySpecification{}).Type1())),
			}.Froze()
			jsonit.Unmarshal(objByte, &obj)

			*(*FleetSpecTargetCapacitySpecification)(ptr) = obj
		} else {
			*(*FleetSpecTargetCapacitySpecification)(ptr) = FleetSpecTargetCapacitySpecification{}
		}
	default:
		iter.ReportError("decode FleetSpecTargetCapacitySpecification", "unexpected JSON type")
	}
}

// +k8s:deepcopy-gen=false
type TrafficMirrorFilterRuleSpecDestinationPortRangeCodec struct {
}

func (TrafficMirrorFilterRuleSpecDestinationPortRangeCodec) IsEmpty(ptr unsafe.Pointer) bool {
	return (*TrafficMirrorFilterRuleSpecDestinationPortRange)(ptr) == nil
}

func (TrafficMirrorFilterRuleSpecDestinationPortRangeCodec) Encode(ptr unsafe.Pointer, stream *jsoniter.Stream) {
	obj := (*TrafficMirrorFilterRuleSpecDestinationPortRange)(ptr)
	var objs []TrafficMirrorFilterRuleSpecDestinationPortRange
	if obj != nil {
		objs = []TrafficMirrorFilterRuleSpecDestinationPortRange{*obj}
	}

	jsonit := jsoniter.Config{
		EscapeHTML:             true,
		SortMapKeys:            true,
		ValidateJsonRawMessage: true,
		TagKey:                 "tf",
		TypeEncoders:           getEncodersWithout(jsoniter.MustGetKind(reflect2.TypeOf(TrafficMirrorFilterRuleSpecDestinationPortRange{}).Type1())),
	}.Froze()

	byt, _ := jsonit.Marshal(objs)

	stream.Write(byt)
}

func (TrafficMirrorFilterRuleSpecDestinationPortRangeCodec) Decode(ptr unsafe.Pointer, iter *jsoniter.Iterator) {
	switch iter.WhatIsNext() {
	case jsoniter.NilValue:
		iter.Skip()
		*(*TrafficMirrorFilterRuleSpecDestinationPortRange)(ptr) = TrafficMirrorFilterRuleSpecDestinationPortRange{}
		return
	case jsoniter.ArrayValue:
		objsByte := iter.SkipAndReturnBytes()
		if len(objsByte) > 0 {
			var objs []TrafficMirrorFilterRuleSpecDestinationPortRange

			jsonit := jsoniter.Config{
				EscapeHTML:             true,
				SortMapKeys:            true,
				ValidateJsonRawMessage: true,
				TagKey:                 "tf",
				TypeDecoders:           getDecodersWithout(jsoniter.MustGetKind(reflect2.TypeOf(TrafficMirrorFilterRuleSpecDestinationPortRange{}).Type1())),
			}.Froze()
			jsonit.Unmarshal(objsByte, &objs)

			if len(objs) > 0 {
				*(*TrafficMirrorFilterRuleSpecDestinationPortRange)(ptr) = objs[0]
			} else {
				*(*TrafficMirrorFilterRuleSpecDestinationPortRange)(ptr) = TrafficMirrorFilterRuleSpecDestinationPortRange{}
			}
		} else {
			*(*TrafficMirrorFilterRuleSpecDestinationPortRange)(ptr) = TrafficMirrorFilterRuleSpecDestinationPortRange{}
		}
	case jsoniter.ObjectValue:
		objByte := iter.SkipAndReturnBytes()
		if len(objByte) > 0 {
			var obj TrafficMirrorFilterRuleSpecDestinationPortRange

			jsonit := jsoniter.Config{
				EscapeHTML:             true,
				SortMapKeys:            true,
				ValidateJsonRawMessage: true,
				TagKey:                 "tf",
				TypeDecoders:           getDecodersWithout(jsoniter.MustGetKind(reflect2.TypeOf(TrafficMirrorFilterRuleSpecDestinationPortRange{}).Type1())),
			}.Froze()
			jsonit.Unmarshal(objByte, &obj)

			*(*TrafficMirrorFilterRuleSpecDestinationPortRange)(ptr) = obj
		} else {
			*(*TrafficMirrorFilterRuleSpecDestinationPortRange)(ptr) = TrafficMirrorFilterRuleSpecDestinationPortRange{}
		}
	default:
		iter.ReportError("decode TrafficMirrorFilterRuleSpecDestinationPortRange", "unexpected JSON type")
	}
}

// +k8s:deepcopy-gen=false
type TrafficMirrorFilterRuleSpecSourcePortRangeCodec struct {
}

func (TrafficMirrorFilterRuleSpecSourcePortRangeCodec) IsEmpty(ptr unsafe.Pointer) bool {
	return (*TrafficMirrorFilterRuleSpecSourcePortRange)(ptr) == nil
}

func (TrafficMirrorFilterRuleSpecSourcePortRangeCodec) Encode(ptr unsafe.Pointer, stream *jsoniter.Stream) {
	obj := (*TrafficMirrorFilterRuleSpecSourcePortRange)(ptr)
	var objs []TrafficMirrorFilterRuleSpecSourcePortRange
	if obj != nil {
		objs = []TrafficMirrorFilterRuleSpecSourcePortRange{*obj}
	}

	jsonit := jsoniter.Config{
		EscapeHTML:             true,
		SortMapKeys:            true,
		ValidateJsonRawMessage: true,
		TagKey:                 "tf",
		TypeEncoders:           getEncodersWithout(jsoniter.MustGetKind(reflect2.TypeOf(TrafficMirrorFilterRuleSpecSourcePortRange{}).Type1())),
	}.Froze()

	byt, _ := jsonit.Marshal(objs)

	stream.Write(byt)
}

func (TrafficMirrorFilterRuleSpecSourcePortRangeCodec) Decode(ptr unsafe.Pointer, iter *jsoniter.Iterator) {
	switch iter.WhatIsNext() {
	case jsoniter.NilValue:
		iter.Skip()
		*(*TrafficMirrorFilterRuleSpecSourcePortRange)(ptr) = TrafficMirrorFilterRuleSpecSourcePortRange{}
		return
	case jsoniter.ArrayValue:
		objsByte := iter.SkipAndReturnBytes()
		if len(objsByte) > 0 {
			var objs []TrafficMirrorFilterRuleSpecSourcePortRange

			jsonit := jsoniter.Config{
				EscapeHTML:             true,
				SortMapKeys:            true,
				ValidateJsonRawMessage: true,
				TagKey:                 "tf",
				TypeDecoders:           getDecodersWithout(jsoniter.MustGetKind(reflect2.TypeOf(TrafficMirrorFilterRuleSpecSourcePortRange{}).Type1())),
			}.Froze()
			jsonit.Unmarshal(objsByte, &objs)

			if len(objs) > 0 {
				*(*TrafficMirrorFilterRuleSpecSourcePortRange)(ptr) = objs[0]
			} else {
				*(*TrafficMirrorFilterRuleSpecSourcePortRange)(ptr) = TrafficMirrorFilterRuleSpecSourcePortRange{}
			}
		} else {
			*(*TrafficMirrorFilterRuleSpecSourcePortRange)(ptr) = TrafficMirrorFilterRuleSpecSourcePortRange{}
		}
	case jsoniter.ObjectValue:
		objByte := iter.SkipAndReturnBytes()
		if len(objByte) > 0 {
			var obj TrafficMirrorFilterRuleSpecSourcePortRange

			jsonit := jsoniter.Config{
				EscapeHTML:             true,
				SortMapKeys:            true,
				ValidateJsonRawMessage: true,
				TagKey:                 "tf",
				TypeDecoders:           getDecodersWithout(jsoniter.MustGetKind(reflect2.TypeOf(TrafficMirrorFilterRuleSpecSourcePortRange{}).Type1())),
			}.Froze()
			jsonit.Unmarshal(objByte, &obj)

			*(*TrafficMirrorFilterRuleSpecSourcePortRange)(ptr) = obj
		} else {
			*(*TrafficMirrorFilterRuleSpecSourcePortRange)(ptr) = TrafficMirrorFilterRuleSpecSourcePortRange{}
		}
	default:
		iter.ReportError("decode TrafficMirrorFilterRuleSpecSourcePortRange", "unexpected JSON type")
	}
}
