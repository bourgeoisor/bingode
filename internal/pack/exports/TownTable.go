// Code generated by the FlatBuffers compiler. DO NOT EDIT.

package exports

import (
	flatbuffers "github.com/google/flatbuffers/go"
)

type TownTable struct {
	_tab flatbuffers.Table
}

func GetRootAsTownTable(buf []byte, offset flatbuffers.UOffsetT) *TownTable {
	n := flatbuffers.GetUOffsetT(buf[offset:])
	x := &TownTable{}
	x.Init(buf, n+offset)
	return x
}

func FinishTownTableBuffer(builder *flatbuffers.Builder, offset flatbuffers.UOffsetT) {
	builder.Finish(offset)
}

func GetSizePrefixedRootAsTownTable(buf []byte, offset flatbuffers.UOffsetT) *TownTable {
	n := flatbuffers.GetUOffsetT(buf[offset+flatbuffers.SizeUint32:])
	x := &TownTable{}
	x.Init(buf, n+offset+flatbuffers.SizeUint32)
	return x
}

func FinishSizePrefixedTownTableBuffer(builder *flatbuffers.Builder, offset flatbuffers.UOffsetT) {
	builder.FinishSizePrefixed(offset)
}

func (rcv *TownTable) Init(buf []byte, i flatbuffers.UOffsetT) {
	rcv._tab.Bytes = buf
	rcv._tab.Pos = i
}

func (rcv *TownTable) Table() flatbuffers.Table {
	return rcv._tab
}

func (rcv *TownTable) Towns(obj *Town, j int) bool {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(4))
	if o != 0 {
		x := rcv._tab.Vector(o)
		x += flatbuffers.UOffsetT(j) * 4
		x = rcv._tab.Indirect(x)
		obj.Init(rcv._tab.Bytes, x)
		return true
	}
	return false
}

func (rcv *TownTable) TownsByKey(obj *Town, key uint32) bool {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(4))
	if o != 0 {
		x := rcv._tab.Vector(o)
		return obj.LookupByKey(key, x, rcv._tab.Bytes)
	}
	return false
}

func (rcv *TownTable) TownsLength() int {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(4))
	if o != 0 {
		return rcv._tab.VectorLen(o)
	}
	return 0
}

func TownTableStart(builder *flatbuffers.Builder) {
	builder.StartObject(1)
}
func TownTableAddTowns(builder *flatbuffers.Builder, towns flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(0, flatbuffers.UOffsetT(towns), 0)
}
func TownTableStartTownsVector(builder *flatbuffers.Builder, numElems int) flatbuffers.UOffsetT {
	return builder.StartVector(4, numElems, 4)
}
func TownTableEnd(builder *flatbuffers.Builder) flatbuffers.UOffsetT {
	return builder.EndObject()
}
type Town struct {
	_tab flatbuffers.Table
}

func GetRootAsTown(buf []byte, offset flatbuffers.UOffsetT) *Town {
	n := flatbuffers.GetUOffsetT(buf[offset:])
	x := &Town{}
	x.Init(buf, n+offset)
	return x
}

func FinishTownBuffer(builder *flatbuffers.Builder, offset flatbuffers.UOffsetT) {
	builder.Finish(offset)
}

func GetSizePrefixedRootAsTown(buf []byte, offset flatbuffers.UOffsetT) *Town {
	n := flatbuffers.GetUOffsetT(buf[offset+flatbuffers.SizeUint32:])
	x := &Town{}
	x.Init(buf, n+offset+flatbuffers.SizeUint32)
	return x
}

func FinishSizePrefixedTownBuffer(builder *flatbuffers.Builder, offset flatbuffers.UOffsetT) {
	builder.FinishSizePrefixed(offset)
}

func (rcv *Town) Init(buf []byte, i flatbuffers.UOffsetT) {
	rcv._tab.Bytes = buf
	rcv._tab.Pos = i
}

func (rcv *Town) Table() flatbuffers.Table {
	return rcv._tab
}

func (rcv *Town) Id() uint32 {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(4))
	if o != 0 {
		return rcv._tab.GetUint32(o + rcv._tab.Pos)
	}
	return 0
}

func (rcv *Town) MutateId(n uint32) bool {
	return rcv._tab.MutateUint32Slot(4, n)
}

func TownKeyCompare(o1, o2 flatbuffers.UOffsetT, buf []byte) bool {
	obj1 := &Town{}
	obj2 := &Town{}
	obj1.Init(buf, flatbuffers.UOffsetT(len(buf))-o1)
	obj2.Init(buf, flatbuffers.UOffsetT(len(buf))-o2)
	return obj1.Id() < obj2.Id()
}

func (rcv *Town) LookupByKey(key uint32, vectorLocation flatbuffers.UOffsetT, buf []byte) bool {
	span := flatbuffers.GetUOffsetT(buf[vectorLocation-4:])
	start := flatbuffers.UOffsetT(0)
	for span != 0 {
		middle := span / 2
		tableOffset := flatbuffers.GetIndirectOffset(buf, vectorLocation+4*(start+middle))
		obj := &Town{}
		obj.Init(buf, tableOffset)
		val := obj.Id()
		comp := 0
		if val > key {
			comp = 1
		} else if val < key {
			comp = -1
		}
		if comp > 0 {
			span = middle
		} else if comp < 0 {
			middle += 1
			start += middle
			span -= middle
		} else {
			rcv.Init(buf, tableOffset)
			return true
		}
	}
	return false
}

func (rcv *Town) NameEn() []byte {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(6))
	if o != 0 {
		return rcv._tab.ByteVector(o + rcv._tab.Pos)
	}
	return nil
}

func (rcv *Town) NameFr() []byte {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(8))
	if o != 0 {
		return rcv._tab.ByteVector(o + rcv._tab.Pos)
	}
	return nil
}

func (rcv *Town) NameDe() []byte {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(10))
	if o != 0 {
		return rcv._tab.ByteVector(o + rcv._tab.Pos)
	}
	return nil
}

func (rcv *Town) NameJa() []byte {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(12))
	if o != 0 {
		return rcv._tab.ByteVector(o + rcv._tab.Pos)
	}
	return nil
}

func TownStart(builder *flatbuffers.Builder) {
	builder.StartObject(5)
}
func TownAddId(builder *flatbuffers.Builder, id uint32) {
	builder.PrependUint32Slot(0, id, 0)
}
func TownAddNameEn(builder *flatbuffers.Builder, nameEn flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(1, flatbuffers.UOffsetT(nameEn), 0)
}
func TownAddNameFr(builder *flatbuffers.Builder, nameFr flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(2, flatbuffers.UOffsetT(nameFr), 0)
}
func TownAddNameDe(builder *flatbuffers.Builder, nameDe flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(3, flatbuffers.UOffsetT(nameDe), 0)
}
func TownAddNameJa(builder *flatbuffers.Builder, nameJa flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(4, flatbuffers.UOffsetT(nameJa), 0)
}
func TownEnd(builder *flatbuffers.Builder) flatbuffers.UOffsetT {
	return builder.EndObject()
}
