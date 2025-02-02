// Code generated by the FlatBuffers compiler. DO NOT EDIT.

package exports

import (
	flatbuffers "github.com/google/flatbuffers/go"
)

type DeityTable struct {
	_tab flatbuffers.Table
}

func GetRootAsDeityTable(buf []byte, offset flatbuffers.UOffsetT) *DeityTable {
	n := flatbuffers.GetUOffsetT(buf[offset:])
	x := &DeityTable{}
	x.Init(buf, n+offset)
	return x
}

func FinishDeityTableBuffer(builder *flatbuffers.Builder, offset flatbuffers.UOffsetT) {
	builder.Finish(offset)
}

func GetSizePrefixedRootAsDeityTable(buf []byte, offset flatbuffers.UOffsetT) *DeityTable {
	n := flatbuffers.GetUOffsetT(buf[offset+flatbuffers.SizeUint32:])
	x := &DeityTable{}
	x.Init(buf, n+offset+flatbuffers.SizeUint32)
	return x
}

func FinishSizePrefixedDeityTableBuffer(builder *flatbuffers.Builder, offset flatbuffers.UOffsetT) {
	builder.FinishSizePrefixed(offset)
}

func (rcv *DeityTable) Init(buf []byte, i flatbuffers.UOffsetT) {
	rcv._tab.Bytes = buf
	rcv._tab.Pos = i
}

func (rcv *DeityTable) Table() flatbuffers.Table {
	return rcv._tab
}

func (rcv *DeityTable) Deities(obj *Deity, j int) bool {
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

func (rcv *DeityTable) DeitiesByKey(obj *Deity, key uint32) bool {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(4))
	if o != 0 {
		x := rcv._tab.Vector(o)
		return obj.LookupByKey(key, x, rcv._tab.Bytes)
	}
	return false
}

func (rcv *DeityTable) DeitiesLength() int {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(4))
	if o != 0 {
		return rcv._tab.VectorLen(o)
	}
	return 0
}

func DeityTableStart(builder *flatbuffers.Builder) {
	builder.StartObject(1)
}
func DeityTableAddDeities(builder *flatbuffers.Builder, deities flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(0, flatbuffers.UOffsetT(deities), 0)
}
func DeityTableStartDeitiesVector(builder *flatbuffers.Builder, numElems int) flatbuffers.UOffsetT {
	return builder.StartVector(4, numElems, 4)
}
func DeityTableEnd(builder *flatbuffers.Builder) flatbuffers.UOffsetT {
	return builder.EndObject()
}
type Deity struct {
	_tab flatbuffers.Table
}

func GetRootAsDeity(buf []byte, offset flatbuffers.UOffsetT) *Deity {
	n := flatbuffers.GetUOffsetT(buf[offset:])
	x := &Deity{}
	x.Init(buf, n+offset)
	return x
}

func FinishDeityBuffer(builder *flatbuffers.Builder, offset flatbuffers.UOffsetT) {
	builder.Finish(offset)
}

func GetSizePrefixedRootAsDeity(buf []byte, offset flatbuffers.UOffsetT) *Deity {
	n := flatbuffers.GetUOffsetT(buf[offset+flatbuffers.SizeUint32:])
	x := &Deity{}
	x.Init(buf, n+offset+flatbuffers.SizeUint32)
	return x
}

func FinishSizePrefixedDeityBuffer(builder *flatbuffers.Builder, offset flatbuffers.UOffsetT) {
	builder.FinishSizePrefixed(offset)
}

func (rcv *Deity) Init(buf []byte, i flatbuffers.UOffsetT) {
	rcv._tab.Bytes = buf
	rcv._tab.Pos = i
}

func (rcv *Deity) Table() flatbuffers.Table {
	return rcv._tab
}

func (rcv *Deity) Id() uint32 {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(4))
	if o != 0 {
		return rcv._tab.GetUint32(o + rcv._tab.Pos)
	}
	return 0
}

func (rcv *Deity) MutateId(n uint32) bool {
	return rcv._tab.MutateUint32Slot(4, n)
}

func DeityKeyCompare(o1, o2 flatbuffers.UOffsetT, buf []byte) bool {
	obj1 := &Deity{}
	obj2 := &Deity{}
	obj1.Init(buf, flatbuffers.UOffsetT(len(buf))-o1)
	obj2.Init(buf, flatbuffers.UOffsetT(len(buf))-o2)
	return obj1.Id() < obj2.Id()
}

func (rcv *Deity) LookupByKey(key uint32, vectorLocation flatbuffers.UOffsetT, buf []byte) bool {
	span := flatbuffers.GetUOffsetT(buf[vectorLocation-4:])
	start := flatbuffers.UOffsetT(0)
	for span != 0 {
		middle := span / 2
		tableOffset := flatbuffers.GetIndirectOffset(buf, vectorLocation+4*(start+middle))
		obj := &Deity{}
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

func (rcv *Deity) NameEn() []byte {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(6))
	if o != 0 {
		return rcv._tab.ByteVector(o + rcv._tab.Pos)
	}
	return nil
}

func (rcv *Deity) NameFr() []byte {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(8))
	if o != 0 {
		return rcv._tab.ByteVector(o + rcv._tab.Pos)
	}
	return nil
}

func (rcv *Deity) NameDe() []byte {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(10))
	if o != 0 {
		return rcv._tab.ByteVector(o + rcv._tab.Pos)
	}
	return nil
}

func (rcv *Deity) NameJa() []byte {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(12))
	if o != 0 {
		return rcv._tab.ByteVector(o + rcv._tab.Pos)
	}
	return nil
}

func DeityStart(builder *flatbuffers.Builder) {
	builder.StartObject(5)
}
func DeityAddId(builder *flatbuffers.Builder, id uint32) {
	builder.PrependUint32Slot(0, id, 0)
}
func DeityAddNameEn(builder *flatbuffers.Builder, nameEn flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(1, flatbuffers.UOffsetT(nameEn), 0)
}
func DeityAddNameFr(builder *flatbuffers.Builder, nameFr flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(2, flatbuffers.UOffsetT(nameFr), 0)
}
func DeityAddNameDe(builder *flatbuffers.Builder, nameDe flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(3, flatbuffers.UOffsetT(nameDe), 0)
}
func DeityAddNameJa(builder *flatbuffers.Builder, nameJa flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(4, flatbuffers.UOffsetT(nameJa), 0)
}
func DeityEnd(builder *flatbuffers.Builder) flatbuffers.UOffsetT {
	return builder.EndObject()
}
