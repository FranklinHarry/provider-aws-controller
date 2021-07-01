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
		jsoniter.MustGetKind(reflect2.TypeOf(BudgetSpecCostTypes{}).Type1()):                           BudgetSpecCostTypesCodec{},
		jsoniter.MustGetKind(reflect2.TypeOf(BudgetActionSpecActionThreshold{}).Type1()):               BudgetActionSpecActionThresholdCodec{},
		jsoniter.MustGetKind(reflect2.TypeOf(BudgetActionSpecDefinition{}).Type1()):                    BudgetActionSpecDefinitionCodec{},
		jsoniter.MustGetKind(reflect2.TypeOf(BudgetActionSpecDefinitionIamActionDefinition{}).Type1()): BudgetActionSpecDefinitionIamActionDefinitionCodec{},
		jsoniter.MustGetKind(reflect2.TypeOf(BudgetActionSpecDefinitionScpActionDefinition{}).Type1()): BudgetActionSpecDefinitionScpActionDefinitionCodec{},
		jsoniter.MustGetKind(reflect2.TypeOf(BudgetActionSpecDefinitionSsmActionDefinition{}).Type1()): BudgetActionSpecDefinitionSsmActionDefinitionCodec{},
	}
}

func GetDecoder() map[string]jsoniter.ValDecoder {
	return map[string]jsoniter.ValDecoder{
		jsoniter.MustGetKind(reflect2.TypeOf(BudgetSpecCostTypes{}).Type1()):                           BudgetSpecCostTypesCodec{},
		jsoniter.MustGetKind(reflect2.TypeOf(BudgetActionSpecActionThreshold{}).Type1()):               BudgetActionSpecActionThresholdCodec{},
		jsoniter.MustGetKind(reflect2.TypeOf(BudgetActionSpecDefinition{}).Type1()):                    BudgetActionSpecDefinitionCodec{},
		jsoniter.MustGetKind(reflect2.TypeOf(BudgetActionSpecDefinitionIamActionDefinition{}).Type1()): BudgetActionSpecDefinitionIamActionDefinitionCodec{},
		jsoniter.MustGetKind(reflect2.TypeOf(BudgetActionSpecDefinitionScpActionDefinition{}).Type1()): BudgetActionSpecDefinitionScpActionDefinitionCodec{},
		jsoniter.MustGetKind(reflect2.TypeOf(BudgetActionSpecDefinitionSsmActionDefinition{}).Type1()): BudgetActionSpecDefinitionSsmActionDefinitionCodec{},
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
type BudgetSpecCostTypesCodec struct {
}

func (BudgetSpecCostTypesCodec) IsEmpty(ptr unsafe.Pointer) bool {
	return (*BudgetSpecCostTypes)(ptr) == nil
}

func (BudgetSpecCostTypesCodec) Encode(ptr unsafe.Pointer, stream *jsoniter.Stream) {
	obj := (*BudgetSpecCostTypes)(ptr)
	var objs []BudgetSpecCostTypes
	if obj != nil {
		objs = []BudgetSpecCostTypes{*obj}
	}

	jsonit := jsoniter.Config{
		EscapeHTML:             true,
		SortMapKeys:            true,
		ValidateJsonRawMessage: true,
		TagKey:                 "tf",
		TypeEncoders:           getEncodersWithout(jsoniter.MustGetKind(reflect2.TypeOf(BudgetSpecCostTypes{}).Type1())),
	}.Froze()

	byt, _ := jsonit.Marshal(objs)

	stream.Write(byt)
}

func (BudgetSpecCostTypesCodec) Decode(ptr unsafe.Pointer, iter *jsoniter.Iterator) {
	switch iter.WhatIsNext() {
	case jsoniter.NilValue:
		iter.Skip()
		*(*BudgetSpecCostTypes)(ptr) = BudgetSpecCostTypes{}
		return
	case jsoniter.ArrayValue:
		objsByte := iter.SkipAndReturnBytes()
		if len(objsByte) > 0 {
			var objs []BudgetSpecCostTypes

			jsonit := jsoniter.Config{
				EscapeHTML:             true,
				SortMapKeys:            true,
				ValidateJsonRawMessage: true,
				TagKey:                 "tf",
				TypeDecoders:           getDecodersWithout(jsoniter.MustGetKind(reflect2.TypeOf(BudgetSpecCostTypes{}).Type1())),
			}.Froze()
			jsonit.Unmarshal(objsByte, &objs)

			if len(objs) > 0 {
				*(*BudgetSpecCostTypes)(ptr) = objs[0]
			} else {
				*(*BudgetSpecCostTypes)(ptr) = BudgetSpecCostTypes{}
			}
		} else {
			*(*BudgetSpecCostTypes)(ptr) = BudgetSpecCostTypes{}
		}
	case jsoniter.ObjectValue:
		objByte := iter.SkipAndReturnBytes()
		if len(objByte) > 0 {
			var obj BudgetSpecCostTypes

			jsonit := jsoniter.Config{
				EscapeHTML:             true,
				SortMapKeys:            true,
				ValidateJsonRawMessage: true,
				TagKey:                 "tf",
				TypeDecoders:           getDecodersWithout(jsoniter.MustGetKind(reflect2.TypeOf(BudgetSpecCostTypes{}).Type1())),
			}.Froze()
			jsonit.Unmarshal(objByte, &obj)

			*(*BudgetSpecCostTypes)(ptr) = obj
		} else {
			*(*BudgetSpecCostTypes)(ptr) = BudgetSpecCostTypes{}
		}
	default:
		iter.ReportError("decode BudgetSpecCostTypes", "unexpected JSON type")
	}
}

// +k8s:deepcopy-gen=false
type BudgetActionSpecActionThresholdCodec struct {
}

func (BudgetActionSpecActionThresholdCodec) IsEmpty(ptr unsafe.Pointer) bool {
	return (*BudgetActionSpecActionThreshold)(ptr) == nil
}

func (BudgetActionSpecActionThresholdCodec) Encode(ptr unsafe.Pointer, stream *jsoniter.Stream) {
	obj := (*BudgetActionSpecActionThreshold)(ptr)
	var objs []BudgetActionSpecActionThreshold
	if obj != nil {
		objs = []BudgetActionSpecActionThreshold{*obj}
	}

	jsonit := jsoniter.Config{
		EscapeHTML:             true,
		SortMapKeys:            true,
		ValidateJsonRawMessage: true,
		TagKey:                 "tf",
		TypeEncoders:           getEncodersWithout(jsoniter.MustGetKind(reflect2.TypeOf(BudgetActionSpecActionThreshold{}).Type1())),
	}.Froze()

	byt, _ := jsonit.Marshal(objs)

	stream.Write(byt)
}

func (BudgetActionSpecActionThresholdCodec) Decode(ptr unsafe.Pointer, iter *jsoniter.Iterator) {
	switch iter.WhatIsNext() {
	case jsoniter.NilValue:
		iter.Skip()
		*(*BudgetActionSpecActionThreshold)(ptr) = BudgetActionSpecActionThreshold{}
		return
	case jsoniter.ArrayValue:
		objsByte := iter.SkipAndReturnBytes()
		if len(objsByte) > 0 {
			var objs []BudgetActionSpecActionThreshold

			jsonit := jsoniter.Config{
				EscapeHTML:             true,
				SortMapKeys:            true,
				ValidateJsonRawMessage: true,
				TagKey:                 "tf",
				TypeDecoders:           getDecodersWithout(jsoniter.MustGetKind(reflect2.TypeOf(BudgetActionSpecActionThreshold{}).Type1())),
			}.Froze()
			jsonit.Unmarshal(objsByte, &objs)

			if len(objs) > 0 {
				*(*BudgetActionSpecActionThreshold)(ptr) = objs[0]
			} else {
				*(*BudgetActionSpecActionThreshold)(ptr) = BudgetActionSpecActionThreshold{}
			}
		} else {
			*(*BudgetActionSpecActionThreshold)(ptr) = BudgetActionSpecActionThreshold{}
		}
	case jsoniter.ObjectValue:
		objByte := iter.SkipAndReturnBytes()
		if len(objByte) > 0 {
			var obj BudgetActionSpecActionThreshold

			jsonit := jsoniter.Config{
				EscapeHTML:             true,
				SortMapKeys:            true,
				ValidateJsonRawMessage: true,
				TagKey:                 "tf",
				TypeDecoders:           getDecodersWithout(jsoniter.MustGetKind(reflect2.TypeOf(BudgetActionSpecActionThreshold{}).Type1())),
			}.Froze()
			jsonit.Unmarshal(objByte, &obj)

			*(*BudgetActionSpecActionThreshold)(ptr) = obj
		} else {
			*(*BudgetActionSpecActionThreshold)(ptr) = BudgetActionSpecActionThreshold{}
		}
	default:
		iter.ReportError("decode BudgetActionSpecActionThreshold", "unexpected JSON type")
	}
}

// +k8s:deepcopy-gen=false
type BudgetActionSpecDefinitionCodec struct {
}

func (BudgetActionSpecDefinitionCodec) IsEmpty(ptr unsafe.Pointer) bool {
	return (*BudgetActionSpecDefinition)(ptr) == nil
}

func (BudgetActionSpecDefinitionCodec) Encode(ptr unsafe.Pointer, stream *jsoniter.Stream) {
	obj := (*BudgetActionSpecDefinition)(ptr)
	var objs []BudgetActionSpecDefinition
	if obj != nil {
		objs = []BudgetActionSpecDefinition{*obj}
	}

	jsonit := jsoniter.Config{
		EscapeHTML:             true,
		SortMapKeys:            true,
		ValidateJsonRawMessage: true,
		TagKey:                 "tf",
		TypeEncoders:           getEncodersWithout(jsoniter.MustGetKind(reflect2.TypeOf(BudgetActionSpecDefinition{}).Type1())),
	}.Froze()

	byt, _ := jsonit.Marshal(objs)

	stream.Write(byt)
}

func (BudgetActionSpecDefinitionCodec) Decode(ptr unsafe.Pointer, iter *jsoniter.Iterator) {
	switch iter.WhatIsNext() {
	case jsoniter.NilValue:
		iter.Skip()
		*(*BudgetActionSpecDefinition)(ptr) = BudgetActionSpecDefinition{}
		return
	case jsoniter.ArrayValue:
		objsByte := iter.SkipAndReturnBytes()
		if len(objsByte) > 0 {
			var objs []BudgetActionSpecDefinition

			jsonit := jsoniter.Config{
				EscapeHTML:             true,
				SortMapKeys:            true,
				ValidateJsonRawMessage: true,
				TagKey:                 "tf",
				TypeDecoders:           getDecodersWithout(jsoniter.MustGetKind(reflect2.TypeOf(BudgetActionSpecDefinition{}).Type1())),
			}.Froze()
			jsonit.Unmarshal(objsByte, &objs)

			if len(objs) > 0 {
				*(*BudgetActionSpecDefinition)(ptr) = objs[0]
			} else {
				*(*BudgetActionSpecDefinition)(ptr) = BudgetActionSpecDefinition{}
			}
		} else {
			*(*BudgetActionSpecDefinition)(ptr) = BudgetActionSpecDefinition{}
		}
	case jsoniter.ObjectValue:
		objByte := iter.SkipAndReturnBytes()
		if len(objByte) > 0 {
			var obj BudgetActionSpecDefinition

			jsonit := jsoniter.Config{
				EscapeHTML:             true,
				SortMapKeys:            true,
				ValidateJsonRawMessage: true,
				TagKey:                 "tf",
				TypeDecoders:           getDecodersWithout(jsoniter.MustGetKind(reflect2.TypeOf(BudgetActionSpecDefinition{}).Type1())),
			}.Froze()
			jsonit.Unmarshal(objByte, &obj)

			*(*BudgetActionSpecDefinition)(ptr) = obj
		} else {
			*(*BudgetActionSpecDefinition)(ptr) = BudgetActionSpecDefinition{}
		}
	default:
		iter.ReportError("decode BudgetActionSpecDefinition", "unexpected JSON type")
	}
}

// +k8s:deepcopy-gen=false
type BudgetActionSpecDefinitionIamActionDefinitionCodec struct {
}

func (BudgetActionSpecDefinitionIamActionDefinitionCodec) IsEmpty(ptr unsafe.Pointer) bool {
	return (*BudgetActionSpecDefinitionIamActionDefinition)(ptr) == nil
}

func (BudgetActionSpecDefinitionIamActionDefinitionCodec) Encode(ptr unsafe.Pointer, stream *jsoniter.Stream) {
	obj := (*BudgetActionSpecDefinitionIamActionDefinition)(ptr)
	var objs []BudgetActionSpecDefinitionIamActionDefinition
	if obj != nil {
		objs = []BudgetActionSpecDefinitionIamActionDefinition{*obj}
	}

	jsonit := jsoniter.Config{
		EscapeHTML:             true,
		SortMapKeys:            true,
		ValidateJsonRawMessage: true,
		TagKey:                 "tf",
		TypeEncoders:           getEncodersWithout(jsoniter.MustGetKind(reflect2.TypeOf(BudgetActionSpecDefinitionIamActionDefinition{}).Type1())),
	}.Froze()

	byt, _ := jsonit.Marshal(objs)

	stream.Write(byt)
}

func (BudgetActionSpecDefinitionIamActionDefinitionCodec) Decode(ptr unsafe.Pointer, iter *jsoniter.Iterator) {
	switch iter.WhatIsNext() {
	case jsoniter.NilValue:
		iter.Skip()
		*(*BudgetActionSpecDefinitionIamActionDefinition)(ptr) = BudgetActionSpecDefinitionIamActionDefinition{}
		return
	case jsoniter.ArrayValue:
		objsByte := iter.SkipAndReturnBytes()
		if len(objsByte) > 0 {
			var objs []BudgetActionSpecDefinitionIamActionDefinition

			jsonit := jsoniter.Config{
				EscapeHTML:             true,
				SortMapKeys:            true,
				ValidateJsonRawMessage: true,
				TagKey:                 "tf",
				TypeDecoders:           getDecodersWithout(jsoniter.MustGetKind(reflect2.TypeOf(BudgetActionSpecDefinitionIamActionDefinition{}).Type1())),
			}.Froze()
			jsonit.Unmarshal(objsByte, &objs)

			if len(objs) > 0 {
				*(*BudgetActionSpecDefinitionIamActionDefinition)(ptr) = objs[0]
			} else {
				*(*BudgetActionSpecDefinitionIamActionDefinition)(ptr) = BudgetActionSpecDefinitionIamActionDefinition{}
			}
		} else {
			*(*BudgetActionSpecDefinitionIamActionDefinition)(ptr) = BudgetActionSpecDefinitionIamActionDefinition{}
		}
	case jsoniter.ObjectValue:
		objByte := iter.SkipAndReturnBytes()
		if len(objByte) > 0 {
			var obj BudgetActionSpecDefinitionIamActionDefinition

			jsonit := jsoniter.Config{
				EscapeHTML:             true,
				SortMapKeys:            true,
				ValidateJsonRawMessage: true,
				TagKey:                 "tf",
				TypeDecoders:           getDecodersWithout(jsoniter.MustGetKind(reflect2.TypeOf(BudgetActionSpecDefinitionIamActionDefinition{}).Type1())),
			}.Froze()
			jsonit.Unmarshal(objByte, &obj)

			*(*BudgetActionSpecDefinitionIamActionDefinition)(ptr) = obj
		} else {
			*(*BudgetActionSpecDefinitionIamActionDefinition)(ptr) = BudgetActionSpecDefinitionIamActionDefinition{}
		}
	default:
		iter.ReportError("decode BudgetActionSpecDefinitionIamActionDefinition", "unexpected JSON type")
	}
}

// +k8s:deepcopy-gen=false
type BudgetActionSpecDefinitionScpActionDefinitionCodec struct {
}

func (BudgetActionSpecDefinitionScpActionDefinitionCodec) IsEmpty(ptr unsafe.Pointer) bool {
	return (*BudgetActionSpecDefinitionScpActionDefinition)(ptr) == nil
}

func (BudgetActionSpecDefinitionScpActionDefinitionCodec) Encode(ptr unsafe.Pointer, stream *jsoniter.Stream) {
	obj := (*BudgetActionSpecDefinitionScpActionDefinition)(ptr)
	var objs []BudgetActionSpecDefinitionScpActionDefinition
	if obj != nil {
		objs = []BudgetActionSpecDefinitionScpActionDefinition{*obj}
	}

	jsonit := jsoniter.Config{
		EscapeHTML:             true,
		SortMapKeys:            true,
		ValidateJsonRawMessage: true,
		TagKey:                 "tf",
		TypeEncoders:           getEncodersWithout(jsoniter.MustGetKind(reflect2.TypeOf(BudgetActionSpecDefinitionScpActionDefinition{}).Type1())),
	}.Froze()

	byt, _ := jsonit.Marshal(objs)

	stream.Write(byt)
}

func (BudgetActionSpecDefinitionScpActionDefinitionCodec) Decode(ptr unsafe.Pointer, iter *jsoniter.Iterator) {
	switch iter.WhatIsNext() {
	case jsoniter.NilValue:
		iter.Skip()
		*(*BudgetActionSpecDefinitionScpActionDefinition)(ptr) = BudgetActionSpecDefinitionScpActionDefinition{}
		return
	case jsoniter.ArrayValue:
		objsByte := iter.SkipAndReturnBytes()
		if len(objsByte) > 0 {
			var objs []BudgetActionSpecDefinitionScpActionDefinition

			jsonit := jsoniter.Config{
				EscapeHTML:             true,
				SortMapKeys:            true,
				ValidateJsonRawMessage: true,
				TagKey:                 "tf",
				TypeDecoders:           getDecodersWithout(jsoniter.MustGetKind(reflect2.TypeOf(BudgetActionSpecDefinitionScpActionDefinition{}).Type1())),
			}.Froze()
			jsonit.Unmarshal(objsByte, &objs)

			if len(objs) > 0 {
				*(*BudgetActionSpecDefinitionScpActionDefinition)(ptr) = objs[0]
			} else {
				*(*BudgetActionSpecDefinitionScpActionDefinition)(ptr) = BudgetActionSpecDefinitionScpActionDefinition{}
			}
		} else {
			*(*BudgetActionSpecDefinitionScpActionDefinition)(ptr) = BudgetActionSpecDefinitionScpActionDefinition{}
		}
	case jsoniter.ObjectValue:
		objByte := iter.SkipAndReturnBytes()
		if len(objByte) > 0 {
			var obj BudgetActionSpecDefinitionScpActionDefinition

			jsonit := jsoniter.Config{
				EscapeHTML:             true,
				SortMapKeys:            true,
				ValidateJsonRawMessage: true,
				TagKey:                 "tf",
				TypeDecoders:           getDecodersWithout(jsoniter.MustGetKind(reflect2.TypeOf(BudgetActionSpecDefinitionScpActionDefinition{}).Type1())),
			}.Froze()
			jsonit.Unmarshal(objByte, &obj)

			*(*BudgetActionSpecDefinitionScpActionDefinition)(ptr) = obj
		} else {
			*(*BudgetActionSpecDefinitionScpActionDefinition)(ptr) = BudgetActionSpecDefinitionScpActionDefinition{}
		}
	default:
		iter.ReportError("decode BudgetActionSpecDefinitionScpActionDefinition", "unexpected JSON type")
	}
}

// +k8s:deepcopy-gen=false
type BudgetActionSpecDefinitionSsmActionDefinitionCodec struct {
}

func (BudgetActionSpecDefinitionSsmActionDefinitionCodec) IsEmpty(ptr unsafe.Pointer) bool {
	return (*BudgetActionSpecDefinitionSsmActionDefinition)(ptr) == nil
}

func (BudgetActionSpecDefinitionSsmActionDefinitionCodec) Encode(ptr unsafe.Pointer, stream *jsoniter.Stream) {
	obj := (*BudgetActionSpecDefinitionSsmActionDefinition)(ptr)
	var objs []BudgetActionSpecDefinitionSsmActionDefinition
	if obj != nil {
		objs = []BudgetActionSpecDefinitionSsmActionDefinition{*obj}
	}

	jsonit := jsoniter.Config{
		EscapeHTML:             true,
		SortMapKeys:            true,
		ValidateJsonRawMessage: true,
		TagKey:                 "tf",
		TypeEncoders:           getEncodersWithout(jsoniter.MustGetKind(reflect2.TypeOf(BudgetActionSpecDefinitionSsmActionDefinition{}).Type1())),
	}.Froze()

	byt, _ := jsonit.Marshal(objs)

	stream.Write(byt)
}

func (BudgetActionSpecDefinitionSsmActionDefinitionCodec) Decode(ptr unsafe.Pointer, iter *jsoniter.Iterator) {
	switch iter.WhatIsNext() {
	case jsoniter.NilValue:
		iter.Skip()
		*(*BudgetActionSpecDefinitionSsmActionDefinition)(ptr) = BudgetActionSpecDefinitionSsmActionDefinition{}
		return
	case jsoniter.ArrayValue:
		objsByte := iter.SkipAndReturnBytes()
		if len(objsByte) > 0 {
			var objs []BudgetActionSpecDefinitionSsmActionDefinition

			jsonit := jsoniter.Config{
				EscapeHTML:             true,
				SortMapKeys:            true,
				ValidateJsonRawMessage: true,
				TagKey:                 "tf",
				TypeDecoders:           getDecodersWithout(jsoniter.MustGetKind(reflect2.TypeOf(BudgetActionSpecDefinitionSsmActionDefinition{}).Type1())),
			}.Froze()
			jsonit.Unmarshal(objsByte, &objs)

			if len(objs) > 0 {
				*(*BudgetActionSpecDefinitionSsmActionDefinition)(ptr) = objs[0]
			} else {
				*(*BudgetActionSpecDefinitionSsmActionDefinition)(ptr) = BudgetActionSpecDefinitionSsmActionDefinition{}
			}
		} else {
			*(*BudgetActionSpecDefinitionSsmActionDefinition)(ptr) = BudgetActionSpecDefinitionSsmActionDefinition{}
		}
	case jsoniter.ObjectValue:
		objByte := iter.SkipAndReturnBytes()
		if len(objByte) > 0 {
			var obj BudgetActionSpecDefinitionSsmActionDefinition

			jsonit := jsoniter.Config{
				EscapeHTML:             true,
				SortMapKeys:            true,
				ValidateJsonRawMessage: true,
				TagKey:                 "tf",
				TypeDecoders:           getDecodersWithout(jsoniter.MustGetKind(reflect2.TypeOf(BudgetActionSpecDefinitionSsmActionDefinition{}).Type1())),
			}.Froze()
			jsonit.Unmarshal(objByte, &obj)

			*(*BudgetActionSpecDefinitionSsmActionDefinition)(ptr) = obj
		} else {
			*(*BudgetActionSpecDefinitionSsmActionDefinition)(ptr) = BudgetActionSpecDefinitionSsmActionDefinition{}
		}
	default:
		iter.ReportError("decode BudgetActionSpecDefinitionSsmActionDefinition", "unexpected JSON type")
	}
}