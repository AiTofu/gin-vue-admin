package main

import (
	"encoding/json"
	"sync"
	"testing"
)

type Student2 struct {
	Name   string
	Age    int32
	Remark [4096]byte
}

var buf, _ = json.Marshal(Student2{Name: "Geektutu", Age: 25})

func unmarsh() {
	stu := &Student2{}
	json.Unmarshal(buf, stu)
}

var studentPool = sync.Pool{
	New: func() interface{} {
		return new(Student2)
	},
}

func BenchmarkUnmarshal(b *testing.B) {
	for n := 0; n < b.N; n++ {
		stu := &Student2{}
		json.Unmarshal(buf, stu)
	}
}

func BenchmarkUnmarshalWithPool(b *testing.B) {
	for n := 0; n < b.N; n++ {
		stu := studentPool.Get().(*Student2)
		json.Unmarshal(buf, stu)
		studentPool.Put(stu)
	}
}
