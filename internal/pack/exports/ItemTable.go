// Code generated by the FlatBuffers compiler. DO NOT EDIT.

package exports

import (
	flatbuffers "github.com/google/flatbuffers/go"
)

type ItemTable struct {
	_tab flatbuffers.Table
}

func GetRootAsItemTable(buf []byte, offset flatbuffers.UOffsetT) *ItemTable {
	n := flatbuffers.GetUOffsetT(buf[offset:])
	x := &ItemTable{}
	x.Init(buf, n+offset)
	return x
}

func FinishItemTableBuffer(builder *flatbuffers.Builder, offset flatbuffers.UOffsetT) {
	builder.Finish(offset)
}

func GetSizePrefixedRootAsItemTable(buf []byte, offset flatbuffers.UOffsetT) *ItemTable {
	n := flatbuffers.GetUOffsetT(buf[offset+flatbuffers.SizeUint32:])
	x := &ItemTable{}
	x.Init(buf, n+offset+flatbuffers.SizeUint32)
	return x
}

func FinishSizePrefixedItemTableBuffer(builder *flatbuffers.Builder, offset flatbuffers.UOffsetT) {
	builder.FinishSizePrefixed(offset)
}

func (rcv *ItemTable) Init(buf []byte, i flatbuffers.UOffsetT) {
	rcv._tab.Bytes = buf
	rcv._tab.Pos = i
}

func (rcv *ItemTable) Table() flatbuffers.Table {
	return rcv._tab
}

func (rcv *ItemTable) Items(obj *Item, j int) bool {
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

func (rcv *ItemTable) ItemsByKey(obj *Item, key uint32) bool {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(4))
	if o != 0 {
		x := rcv._tab.Vector(o)
		return obj.LookupByKey(key, x, rcv._tab.Bytes)
	}
	return false
}

func (rcv *ItemTable) ItemsLength() int {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(4))
	if o != 0 {
		return rcv._tab.VectorLen(o)
	}
	return 0
}

func ItemTableStart(builder *flatbuffers.Builder) {
	builder.StartObject(1)
}
func ItemTableAddItems(builder *flatbuffers.Builder, items flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(0, flatbuffers.UOffsetT(items), 0)
}
func ItemTableStartItemsVector(builder *flatbuffers.Builder, numElems int) flatbuffers.UOffsetT {
	return builder.StartVector(4, numElems, 4)
}
func ItemTableEnd(builder *flatbuffers.Builder) flatbuffers.UOffsetT {
	return builder.EndObject()
}
type Item struct {
	_tab flatbuffers.Table
}

func GetRootAsItem(buf []byte, offset flatbuffers.UOffsetT) *Item {
	n := flatbuffers.GetUOffsetT(buf[offset:])
	x := &Item{}
	x.Init(buf, n+offset)
	return x
}

func FinishItemBuffer(builder *flatbuffers.Builder, offset flatbuffers.UOffsetT) {
	builder.Finish(offset)
}

func GetSizePrefixedRootAsItem(buf []byte, offset flatbuffers.UOffsetT) *Item {
	n := flatbuffers.GetUOffsetT(buf[offset+flatbuffers.SizeUint32:])
	x := &Item{}
	x.Init(buf, n+offset+flatbuffers.SizeUint32)
	return x
}

func FinishSizePrefixedItemBuffer(builder *flatbuffers.Builder, offset flatbuffers.UOffsetT) {
	builder.FinishSizePrefixed(offset)
}

func (rcv *Item) Init(buf []byte, i flatbuffers.UOffsetT) {
	rcv._tab.Bytes = buf
	rcv._tab.Pos = i
}

func (rcv *Item) Table() flatbuffers.Table {
	return rcv._tab
}

func (rcv *Item) Id() uint32 {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(4))
	if o != 0 {
		return rcv._tab.GetUint32(o + rcv._tab.Pos)
	}
	return 0
}

func (rcv *Item) MutateId(n uint32) bool {
	return rcv._tab.MutateUint32Slot(4, n)
}

func ItemKeyCompare(o1, o2 flatbuffers.UOffsetT, buf []byte) bool {
	obj1 := &Item{}
	obj2 := &Item{}
	obj1.Init(buf, flatbuffers.UOffsetT(len(buf))-o1)
	obj2.Init(buf, flatbuffers.UOffsetT(len(buf))-o2)
	return obj1.Id() < obj2.Id()
}

func (rcv *Item) LookupByKey(key uint32, vectorLocation flatbuffers.UOffsetT, buf []byte) bool {
	span := flatbuffers.GetUOffsetT(buf[vectorLocation-4:])
	start := flatbuffers.UOffsetT(0)
	for span != 0 {
		middle := span / 2
		tableOffset := flatbuffers.GetIndirectOffset(buf, vectorLocation+4*(start+middle))
		obj := &Item{}
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

func (rcv *Item) NameEn() []byte {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(6))
	if o != 0 {
		return rcv._tab.ByteVector(o + rcv._tab.Pos)
	}
	return nil
}

func (rcv *Item) NameFr() []byte {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(8))
	if o != 0 {
		return rcv._tab.ByteVector(o + rcv._tab.Pos)
	}
	return nil
}

func (rcv *Item) NameDe() []byte {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(10))
	if o != 0 {
		return rcv._tab.ByteVector(o + rcv._tab.Pos)
	}
	return nil
}

func (rcv *Item) NameJa() []byte {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(12))
	if o != 0 {
		return rcv._tab.ByteVector(o + rcv._tab.Pos)
	}
	return nil
}

func ItemStart(builder *flatbuffers.Builder) {
	builder.StartObject(5)
}
func ItemAddId(builder *flatbuffers.Builder, id uint32) {
	builder.PrependUint32Slot(0, id, 0)
}
func ItemAddNameEn(builder *flatbuffers.Builder, nameEn flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(1, flatbuffers.UOffsetT(nameEn), 0)
}
func ItemAddNameFr(builder *flatbuffers.Builder, nameFr flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(2, flatbuffers.UOffsetT(nameFr), 0)
}
func ItemAddNameDe(builder *flatbuffers.Builder, nameDe flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(3, flatbuffers.UOffsetT(nameDe), 0)
}
func ItemAddNameJa(builder *flatbuffers.Builder, nameJa flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(4, flatbuffers.UOffsetT(nameJa), 0)
}
func ItemEnd(builder *flatbuffers.Builder) flatbuffers.UOffsetT {
	return builder.EndObject()
}
