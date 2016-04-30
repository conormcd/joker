package main

import (
	"bytes"
)

type (
	Seq interface {
		Object
		First() Object
		Rest() Seq
		IsEmpty() bool
		Cons(obj Object) Seq
	}
	Sequenceable interface {
		Seq() Seq
	}
	SeqIterator struct {
		seq Seq
	}
	ConsSeq struct {
		InfoHolder
		first Object
		rest  Seq
	}
	ArraySeq struct {
		InfoHolder
		arr   []Object
		index int
	}
)

func (seq *ArraySeq) Equals(other interface{}) bool {
	if seq == other {
		return true
	}
	switch s := other.(type) {
	case Sequenceable:
		return SeqsEqual(seq, s.Seq())
	default:
		return false
	}
}

func (seq *ArraySeq) ToString(escape bool) string {
	return SeqToString(seq, escape)
}

func (seq *ArraySeq) WithInfo(info *ObjectInfo) Object {
	seq.info = info
	return seq
}

func (seq *ArraySeq) First() Object {
	return seq.arr[seq.index]
}

func (seq *ArraySeq) Rest() Seq {
	if seq.index+1 < len(seq.arr) {
		return &ArraySeq{index: seq.index + 1, arr: seq.arr}
	}
	return EmptyList
}

func (seq *ArraySeq) IsEmpty() bool {
	return seq.index >= len(seq.arr)
}

func (seq *ArraySeq) Cons(obj Object) Seq {
	return &ConsSeq{first: obj, rest: seq}
}

func SeqsEqual(seq1, seq2 Seq) bool {
	iter2 := iter(seq2)
	for iter1 := iter(seq1); iter1.HasNext(); {
		if !iter2.HasNext() || !iter2.Next().Equals(iter1.Next()) {
			return false
		}
	}
	return !iter2.HasNext()
}

func SeqToString(seq Seq, escape bool) string {
	var b bytes.Buffer
	b.WriteRune('(')
	for iter := iter(seq); iter.HasNext(); {
		b.WriteString(iter.Next().ToString(escape))
		if iter.HasNext() {
			b.WriteRune(' ')
		}
	}
	b.WriteRune(')')
	return b.String()
}

func (seq *ConsSeq) Equals(other interface{}) bool {
	if seq == other {
		return true
	}
	switch s := other.(type) {
	case Sequenceable:
		return SeqsEqual(seq, s.Seq())
	default:
		return false
	}
}

func (seq *ConsSeq) ToString(escape bool) string {
	return SeqToString(seq, escape)
}

func (seq *ConsSeq) WithInfo(info *ObjectInfo) Object {
	seq.info = info
	return seq
}

func (seq *ConsSeq) First() Object {
	return seq.first
}

func (seq *ConsSeq) Rest() Seq {
	return seq.rest
}

func (seq *ConsSeq) IsEmpty() bool {
	return false
}

func (seq *ConsSeq) Cons(obj Object) Seq {
	return &ConsSeq{first: obj, rest: seq}
}

func iter(seq Seq) *SeqIterator {
	return &SeqIterator{seq: seq}
}

func (iter *SeqIterator) Next() Object {
	res := iter.seq.First()
	iter.seq = iter.seq.Rest()
	return res
}

func (iter *SeqIterator) HasNext() bool {
	return !iter.seq.IsEmpty()
}

func Second(seq Seq) Object {
	return seq.Rest().First()
}

func Third(seq Seq) Object {
	return seq.Rest().Rest().First()
}

func Forth(seq Seq) Object {
	return seq.Rest().Rest().Rest().First()
}

func ToSlice(seq Seq) []Object {
	res := make([]Object, 0)
	for !seq.IsEmpty() {
		res = append(res, seq.First())
		seq = seq.Rest()
	}
	return res
}
