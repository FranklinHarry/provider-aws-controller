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
		jsoniter.MustGetKind(reflect2.TypeOf(PeeringConnectionSpecAccepter{}).Type1()):          PeeringConnectionSpecAccepterCodec{},
		jsoniter.MustGetKind(reflect2.TypeOf(PeeringConnectionSpecRequester{}).Type1()):         PeeringConnectionSpecRequesterCodec{},
		jsoniter.MustGetKind(reflect2.TypeOf(PeeringConnectionAccepterSpecAccepter{}).Type1()):  PeeringConnectionAccepterSpecAccepterCodec{},
		jsoniter.MustGetKind(reflect2.TypeOf(PeeringConnectionAccepterSpecRequester{}).Type1()): PeeringConnectionAccepterSpecRequesterCodec{},
		jsoniter.MustGetKind(reflect2.TypeOf(PeeringConnectionOptionsSpecAccepter{}).Type1()):   PeeringConnectionOptionsSpecAccepterCodec{},
		jsoniter.MustGetKind(reflect2.TypeOf(PeeringConnectionOptionsSpecRequester{}).Type1()):  PeeringConnectionOptionsSpecRequesterCodec{},
	}
}

func GetDecoder() map[string]jsoniter.ValDecoder {
	return map[string]jsoniter.ValDecoder{
		jsoniter.MustGetKind(reflect2.TypeOf(PeeringConnectionSpecAccepter{}).Type1()):          PeeringConnectionSpecAccepterCodec{},
		jsoniter.MustGetKind(reflect2.TypeOf(PeeringConnectionSpecRequester{}).Type1()):         PeeringConnectionSpecRequesterCodec{},
		jsoniter.MustGetKind(reflect2.TypeOf(PeeringConnectionAccepterSpecAccepter{}).Type1()):  PeeringConnectionAccepterSpecAccepterCodec{},
		jsoniter.MustGetKind(reflect2.TypeOf(PeeringConnectionAccepterSpecRequester{}).Type1()): PeeringConnectionAccepterSpecRequesterCodec{},
		jsoniter.MustGetKind(reflect2.TypeOf(PeeringConnectionOptionsSpecAccepter{}).Type1()):   PeeringConnectionOptionsSpecAccepterCodec{},
		jsoniter.MustGetKind(reflect2.TypeOf(PeeringConnectionOptionsSpecRequester{}).Type1()):  PeeringConnectionOptionsSpecRequesterCodec{},
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
type PeeringConnectionSpecAccepterCodec struct {
}

func (PeeringConnectionSpecAccepterCodec) IsEmpty(ptr unsafe.Pointer) bool {
	return (*PeeringConnectionSpecAccepter)(ptr) == nil
}

func (PeeringConnectionSpecAccepterCodec) Encode(ptr unsafe.Pointer, stream *jsoniter.Stream) {
	obj := (*PeeringConnectionSpecAccepter)(ptr)
	var objs []PeeringConnectionSpecAccepter
	if obj != nil {
		objs = []PeeringConnectionSpecAccepter{*obj}
	}

	jsonit := jsoniter.Config{
		EscapeHTML:             true,
		SortMapKeys:            true,
		ValidateJsonRawMessage: true,
		TagKey:                 "tf",
		TypeEncoders:           getEncodersWithout(jsoniter.MustGetKind(reflect2.TypeOf(PeeringConnectionSpecAccepter{}).Type1())),
	}.Froze()

	byt, _ := jsonit.Marshal(objs)

	stream.Write(byt)
}

func (PeeringConnectionSpecAccepterCodec) Decode(ptr unsafe.Pointer, iter *jsoniter.Iterator) {
	switch iter.WhatIsNext() {
	case jsoniter.NilValue:
		iter.Skip()
		*(*PeeringConnectionSpecAccepter)(ptr) = PeeringConnectionSpecAccepter{}
		return
	case jsoniter.ArrayValue:
		objsByte := iter.SkipAndReturnBytes()
		if len(objsByte) > 0 {
			var objs []PeeringConnectionSpecAccepter

			jsonit := jsoniter.Config{
				EscapeHTML:             true,
				SortMapKeys:            true,
				ValidateJsonRawMessage: true,
				TagKey:                 "tf",
				TypeDecoders:           getDecodersWithout(jsoniter.MustGetKind(reflect2.TypeOf(PeeringConnectionSpecAccepter{}).Type1())),
			}.Froze()
			jsonit.Unmarshal(objsByte, &objs)

			if len(objs) > 0 {
				*(*PeeringConnectionSpecAccepter)(ptr) = objs[0]
			} else {
				*(*PeeringConnectionSpecAccepter)(ptr) = PeeringConnectionSpecAccepter{}
			}
		} else {
			*(*PeeringConnectionSpecAccepter)(ptr) = PeeringConnectionSpecAccepter{}
		}
	case jsoniter.ObjectValue:
		objByte := iter.SkipAndReturnBytes()
		if len(objByte) > 0 {
			var obj PeeringConnectionSpecAccepter

			jsonit := jsoniter.Config{
				EscapeHTML:             true,
				SortMapKeys:            true,
				ValidateJsonRawMessage: true,
				TagKey:                 "tf",
				TypeDecoders:           getDecodersWithout(jsoniter.MustGetKind(reflect2.TypeOf(PeeringConnectionSpecAccepter{}).Type1())),
			}.Froze()
			jsonit.Unmarshal(objByte, &obj)

			*(*PeeringConnectionSpecAccepter)(ptr) = obj
		} else {
			*(*PeeringConnectionSpecAccepter)(ptr) = PeeringConnectionSpecAccepter{}
		}
	default:
		iter.ReportError("decode PeeringConnectionSpecAccepter", "unexpected JSON type")
	}
}

// +k8s:deepcopy-gen=false
type PeeringConnectionSpecRequesterCodec struct {
}

func (PeeringConnectionSpecRequesterCodec) IsEmpty(ptr unsafe.Pointer) bool {
	return (*PeeringConnectionSpecRequester)(ptr) == nil
}

func (PeeringConnectionSpecRequesterCodec) Encode(ptr unsafe.Pointer, stream *jsoniter.Stream) {
	obj := (*PeeringConnectionSpecRequester)(ptr)
	var objs []PeeringConnectionSpecRequester
	if obj != nil {
		objs = []PeeringConnectionSpecRequester{*obj}
	}

	jsonit := jsoniter.Config{
		EscapeHTML:             true,
		SortMapKeys:            true,
		ValidateJsonRawMessage: true,
		TagKey:                 "tf",
		TypeEncoders:           getEncodersWithout(jsoniter.MustGetKind(reflect2.TypeOf(PeeringConnectionSpecRequester{}).Type1())),
	}.Froze()

	byt, _ := jsonit.Marshal(objs)

	stream.Write(byt)
}

func (PeeringConnectionSpecRequesterCodec) Decode(ptr unsafe.Pointer, iter *jsoniter.Iterator) {
	switch iter.WhatIsNext() {
	case jsoniter.NilValue:
		iter.Skip()
		*(*PeeringConnectionSpecRequester)(ptr) = PeeringConnectionSpecRequester{}
		return
	case jsoniter.ArrayValue:
		objsByte := iter.SkipAndReturnBytes()
		if len(objsByte) > 0 {
			var objs []PeeringConnectionSpecRequester

			jsonit := jsoniter.Config{
				EscapeHTML:             true,
				SortMapKeys:            true,
				ValidateJsonRawMessage: true,
				TagKey:                 "tf",
				TypeDecoders:           getDecodersWithout(jsoniter.MustGetKind(reflect2.TypeOf(PeeringConnectionSpecRequester{}).Type1())),
			}.Froze()
			jsonit.Unmarshal(objsByte, &objs)

			if len(objs) > 0 {
				*(*PeeringConnectionSpecRequester)(ptr) = objs[0]
			} else {
				*(*PeeringConnectionSpecRequester)(ptr) = PeeringConnectionSpecRequester{}
			}
		} else {
			*(*PeeringConnectionSpecRequester)(ptr) = PeeringConnectionSpecRequester{}
		}
	case jsoniter.ObjectValue:
		objByte := iter.SkipAndReturnBytes()
		if len(objByte) > 0 {
			var obj PeeringConnectionSpecRequester

			jsonit := jsoniter.Config{
				EscapeHTML:             true,
				SortMapKeys:            true,
				ValidateJsonRawMessage: true,
				TagKey:                 "tf",
				TypeDecoders:           getDecodersWithout(jsoniter.MustGetKind(reflect2.TypeOf(PeeringConnectionSpecRequester{}).Type1())),
			}.Froze()
			jsonit.Unmarshal(objByte, &obj)

			*(*PeeringConnectionSpecRequester)(ptr) = obj
		} else {
			*(*PeeringConnectionSpecRequester)(ptr) = PeeringConnectionSpecRequester{}
		}
	default:
		iter.ReportError("decode PeeringConnectionSpecRequester", "unexpected JSON type")
	}
}

// +k8s:deepcopy-gen=false
type PeeringConnectionAccepterSpecAccepterCodec struct {
}

func (PeeringConnectionAccepterSpecAccepterCodec) IsEmpty(ptr unsafe.Pointer) bool {
	return (*PeeringConnectionAccepterSpecAccepter)(ptr) == nil
}

func (PeeringConnectionAccepterSpecAccepterCodec) Encode(ptr unsafe.Pointer, stream *jsoniter.Stream) {
	obj := (*PeeringConnectionAccepterSpecAccepter)(ptr)
	var objs []PeeringConnectionAccepterSpecAccepter
	if obj != nil {
		objs = []PeeringConnectionAccepterSpecAccepter{*obj}
	}

	jsonit := jsoniter.Config{
		EscapeHTML:             true,
		SortMapKeys:            true,
		ValidateJsonRawMessage: true,
		TagKey:                 "tf",
		TypeEncoders:           getEncodersWithout(jsoniter.MustGetKind(reflect2.TypeOf(PeeringConnectionAccepterSpecAccepter{}).Type1())),
	}.Froze()

	byt, _ := jsonit.Marshal(objs)

	stream.Write(byt)
}

func (PeeringConnectionAccepterSpecAccepterCodec) Decode(ptr unsafe.Pointer, iter *jsoniter.Iterator) {
	switch iter.WhatIsNext() {
	case jsoniter.NilValue:
		iter.Skip()
		*(*PeeringConnectionAccepterSpecAccepter)(ptr) = PeeringConnectionAccepterSpecAccepter{}
		return
	case jsoniter.ArrayValue:
		objsByte := iter.SkipAndReturnBytes()
		if len(objsByte) > 0 {
			var objs []PeeringConnectionAccepterSpecAccepter

			jsonit := jsoniter.Config{
				EscapeHTML:             true,
				SortMapKeys:            true,
				ValidateJsonRawMessage: true,
				TagKey:                 "tf",
				TypeDecoders:           getDecodersWithout(jsoniter.MustGetKind(reflect2.TypeOf(PeeringConnectionAccepterSpecAccepter{}).Type1())),
			}.Froze()
			jsonit.Unmarshal(objsByte, &objs)

			if len(objs) > 0 {
				*(*PeeringConnectionAccepterSpecAccepter)(ptr) = objs[0]
			} else {
				*(*PeeringConnectionAccepterSpecAccepter)(ptr) = PeeringConnectionAccepterSpecAccepter{}
			}
		} else {
			*(*PeeringConnectionAccepterSpecAccepter)(ptr) = PeeringConnectionAccepterSpecAccepter{}
		}
	case jsoniter.ObjectValue:
		objByte := iter.SkipAndReturnBytes()
		if len(objByte) > 0 {
			var obj PeeringConnectionAccepterSpecAccepter

			jsonit := jsoniter.Config{
				EscapeHTML:             true,
				SortMapKeys:            true,
				ValidateJsonRawMessage: true,
				TagKey:                 "tf",
				TypeDecoders:           getDecodersWithout(jsoniter.MustGetKind(reflect2.TypeOf(PeeringConnectionAccepterSpecAccepter{}).Type1())),
			}.Froze()
			jsonit.Unmarshal(objByte, &obj)

			*(*PeeringConnectionAccepterSpecAccepter)(ptr) = obj
		} else {
			*(*PeeringConnectionAccepterSpecAccepter)(ptr) = PeeringConnectionAccepterSpecAccepter{}
		}
	default:
		iter.ReportError("decode PeeringConnectionAccepterSpecAccepter", "unexpected JSON type")
	}
}

// +k8s:deepcopy-gen=false
type PeeringConnectionAccepterSpecRequesterCodec struct {
}

func (PeeringConnectionAccepterSpecRequesterCodec) IsEmpty(ptr unsafe.Pointer) bool {
	return (*PeeringConnectionAccepterSpecRequester)(ptr) == nil
}

func (PeeringConnectionAccepterSpecRequesterCodec) Encode(ptr unsafe.Pointer, stream *jsoniter.Stream) {
	obj := (*PeeringConnectionAccepterSpecRequester)(ptr)
	var objs []PeeringConnectionAccepterSpecRequester
	if obj != nil {
		objs = []PeeringConnectionAccepterSpecRequester{*obj}
	}

	jsonit := jsoniter.Config{
		EscapeHTML:             true,
		SortMapKeys:            true,
		ValidateJsonRawMessage: true,
		TagKey:                 "tf",
		TypeEncoders:           getEncodersWithout(jsoniter.MustGetKind(reflect2.TypeOf(PeeringConnectionAccepterSpecRequester{}).Type1())),
	}.Froze()

	byt, _ := jsonit.Marshal(objs)

	stream.Write(byt)
}

func (PeeringConnectionAccepterSpecRequesterCodec) Decode(ptr unsafe.Pointer, iter *jsoniter.Iterator) {
	switch iter.WhatIsNext() {
	case jsoniter.NilValue:
		iter.Skip()
		*(*PeeringConnectionAccepterSpecRequester)(ptr) = PeeringConnectionAccepterSpecRequester{}
		return
	case jsoniter.ArrayValue:
		objsByte := iter.SkipAndReturnBytes()
		if len(objsByte) > 0 {
			var objs []PeeringConnectionAccepterSpecRequester

			jsonit := jsoniter.Config{
				EscapeHTML:             true,
				SortMapKeys:            true,
				ValidateJsonRawMessage: true,
				TagKey:                 "tf",
				TypeDecoders:           getDecodersWithout(jsoniter.MustGetKind(reflect2.TypeOf(PeeringConnectionAccepterSpecRequester{}).Type1())),
			}.Froze()
			jsonit.Unmarshal(objsByte, &objs)

			if len(objs) > 0 {
				*(*PeeringConnectionAccepterSpecRequester)(ptr) = objs[0]
			} else {
				*(*PeeringConnectionAccepterSpecRequester)(ptr) = PeeringConnectionAccepterSpecRequester{}
			}
		} else {
			*(*PeeringConnectionAccepterSpecRequester)(ptr) = PeeringConnectionAccepterSpecRequester{}
		}
	case jsoniter.ObjectValue:
		objByte := iter.SkipAndReturnBytes()
		if len(objByte) > 0 {
			var obj PeeringConnectionAccepterSpecRequester

			jsonit := jsoniter.Config{
				EscapeHTML:             true,
				SortMapKeys:            true,
				ValidateJsonRawMessage: true,
				TagKey:                 "tf",
				TypeDecoders:           getDecodersWithout(jsoniter.MustGetKind(reflect2.TypeOf(PeeringConnectionAccepterSpecRequester{}).Type1())),
			}.Froze()
			jsonit.Unmarshal(objByte, &obj)

			*(*PeeringConnectionAccepterSpecRequester)(ptr) = obj
		} else {
			*(*PeeringConnectionAccepterSpecRequester)(ptr) = PeeringConnectionAccepterSpecRequester{}
		}
	default:
		iter.ReportError("decode PeeringConnectionAccepterSpecRequester", "unexpected JSON type")
	}
}

// +k8s:deepcopy-gen=false
type PeeringConnectionOptionsSpecAccepterCodec struct {
}

func (PeeringConnectionOptionsSpecAccepterCodec) IsEmpty(ptr unsafe.Pointer) bool {
	return (*PeeringConnectionOptionsSpecAccepter)(ptr) == nil
}

func (PeeringConnectionOptionsSpecAccepterCodec) Encode(ptr unsafe.Pointer, stream *jsoniter.Stream) {
	obj := (*PeeringConnectionOptionsSpecAccepter)(ptr)
	var objs []PeeringConnectionOptionsSpecAccepter
	if obj != nil {
		objs = []PeeringConnectionOptionsSpecAccepter{*obj}
	}

	jsonit := jsoniter.Config{
		EscapeHTML:             true,
		SortMapKeys:            true,
		ValidateJsonRawMessage: true,
		TagKey:                 "tf",
		TypeEncoders:           getEncodersWithout(jsoniter.MustGetKind(reflect2.TypeOf(PeeringConnectionOptionsSpecAccepter{}).Type1())),
	}.Froze()

	byt, _ := jsonit.Marshal(objs)

	stream.Write(byt)
}

func (PeeringConnectionOptionsSpecAccepterCodec) Decode(ptr unsafe.Pointer, iter *jsoniter.Iterator) {
	switch iter.WhatIsNext() {
	case jsoniter.NilValue:
		iter.Skip()
		*(*PeeringConnectionOptionsSpecAccepter)(ptr) = PeeringConnectionOptionsSpecAccepter{}
		return
	case jsoniter.ArrayValue:
		objsByte := iter.SkipAndReturnBytes()
		if len(objsByte) > 0 {
			var objs []PeeringConnectionOptionsSpecAccepter

			jsonit := jsoniter.Config{
				EscapeHTML:             true,
				SortMapKeys:            true,
				ValidateJsonRawMessage: true,
				TagKey:                 "tf",
				TypeDecoders:           getDecodersWithout(jsoniter.MustGetKind(reflect2.TypeOf(PeeringConnectionOptionsSpecAccepter{}).Type1())),
			}.Froze()
			jsonit.Unmarshal(objsByte, &objs)

			if len(objs) > 0 {
				*(*PeeringConnectionOptionsSpecAccepter)(ptr) = objs[0]
			} else {
				*(*PeeringConnectionOptionsSpecAccepter)(ptr) = PeeringConnectionOptionsSpecAccepter{}
			}
		} else {
			*(*PeeringConnectionOptionsSpecAccepter)(ptr) = PeeringConnectionOptionsSpecAccepter{}
		}
	case jsoniter.ObjectValue:
		objByte := iter.SkipAndReturnBytes()
		if len(objByte) > 0 {
			var obj PeeringConnectionOptionsSpecAccepter

			jsonit := jsoniter.Config{
				EscapeHTML:             true,
				SortMapKeys:            true,
				ValidateJsonRawMessage: true,
				TagKey:                 "tf",
				TypeDecoders:           getDecodersWithout(jsoniter.MustGetKind(reflect2.TypeOf(PeeringConnectionOptionsSpecAccepter{}).Type1())),
			}.Froze()
			jsonit.Unmarshal(objByte, &obj)

			*(*PeeringConnectionOptionsSpecAccepter)(ptr) = obj
		} else {
			*(*PeeringConnectionOptionsSpecAccepter)(ptr) = PeeringConnectionOptionsSpecAccepter{}
		}
	default:
		iter.ReportError("decode PeeringConnectionOptionsSpecAccepter", "unexpected JSON type")
	}
}

// +k8s:deepcopy-gen=false
type PeeringConnectionOptionsSpecRequesterCodec struct {
}

func (PeeringConnectionOptionsSpecRequesterCodec) IsEmpty(ptr unsafe.Pointer) bool {
	return (*PeeringConnectionOptionsSpecRequester)(ptr) == nil
}

func (PeeringConnectionOptionsSpecRequesterCodec) Encode(ptr unsafe.Pointer, stream *jsoniter.Stream) {
	obj := (*PeeringConnectionOptionsSpecRequester)(ptr)
	var objs []PeeringConnectionOptionsSpecRequester
	if obj != nil {
		objs = []PeeringConnectionOptionsSpecRequester{*obj}
	}

	jsonit := jsoniter.Config{
		EscapeHTML:             true,
		SortMapKeys:            true,
		ValidateJsonRawMessage: true,
		TagKey:                 "tf",
		TypeEncoders:           getEncodersWithout(jsoniter.MustGetKind(reflect2.TypeOf(PeeringConnectionOptionsSpecRequester{}).Type1())),
	}.Froze()

	byt, _ := jsonit.Marshal(objs)

	stream.Write(byt)
}

func (PeeringConnectionOptionsSpecRequesterCodec) Decode(ptr unsafe.Pointer, iter *jsoniter.Iterator) {
	switch iter.WhatIsNext() {
	case jsoniter.NilValue:
		iter.Skip()
		*(*PeeringConnectionOptionsSpecRequester)(ptr) = PeeringConnectionOptionsSpecRequester{}
		return
	case jsoniter.ArrayValue:
		objsByte := iter.SkipAndReturnBytes()
		if len(objsByte) > 0 {
			var objs []PeeringConnectionOptionsSpecRequester

			jsonit := jsoniter.Config{
				EscapeHTML:             true,
				SortMapKeys:            true,
				ValidateJsonRawMessage: true,
				TagKey:                 "tf",
				TypeDecoders:           getDecodersWithout(jsoniter.MustGetKind(reflect2.TypeOf(PeeringConnectionOptionsSpecRequester{}).Type1())),
			}.Froze()
			jsonit.Unmarshal(objsByte, &objs)

			if len(objs) > 0 {
				*(*PeeringConnectionOptionsSpecRequester)(ptr) = objs[0]
			} else {
				*(*PeeringConnectionOptionsSpecRequester)(ptr) = PeeringConnectionOptionsSpecRequester{}
			}
		} else {
			*(*PeeringConnectionOptionsSpecRequester)(ptr) = PeeringConnectionOptionsSpecRequester{}
		}
	case jsoniter.ObjectValue:
		objByte := iter.SkipAndReturnBytes()
		if len(objByte) > 0 {
			var obj PeeringConnectionOptionsSpecRequester

			jsonit := jsoniter.Config{
				EscapeHTML:             true,
				SortMapKeys:            true,
				ValidateJsonRawMessage: true,
				TagKey:                 "tf",
				TypeDecoders:           getDecodersWithout(jsoniter.MustGetKind(reflect2.TypeOf(PeeringConnectionOptionsSpecRequester{}).Type1())),
			}.Froze()
			jsonit.Unmarshal(objByte, &obj)

			*(*PeeringConnectionOptionsSpecRequester)(ptr) = obj
		} else {
			*(*PeeringConnectionOptionsSpecRequester)(ptr) = PeeringConnectionOptionsSpecRequester{}
		}
	default:
		iter.ReportError("decode PeeringConnectionOptionsSpecRequester", "unexpected JSON type")
	}
}
