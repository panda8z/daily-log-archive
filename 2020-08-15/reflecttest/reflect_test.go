package main

import (
	"reflect"
	"testing"
)

type data struct {
	Hp int
}

func BenchmarkNativeAssign(b *testing.B) {
	v := data{Hp: 2}

	b.StopTimer()

	b.ResetTimer()

	b.StartTimer()

	for i := 0; i < b.N; i++ {
		v.Hp = 3
	}
}

func BenchmarkReflectAssign(b *testing.B) {
	v := data{Hp: 2}

	vv := reflect.ValueOf(&v).Elem()

	f := vv.FieldByName("Hp")

	b.StopTimer()
	b.ResetTimer()
	b.StartTimer()

	for i := 0; i < b.N; i++ {
		f.SetInt(3)
	}
}

func BenchmarkReflectFindFieldAndAssign(b *testing.B) {
	v := data{Hp: 2}

	vv := reflect.ValueOf(&v).Elem()
	b.StopTimer()
	b.ResetTimer()
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		vv.FieldByName("Hp").SetInt(3)
	}
}

func foo(v int) int {
	return v
}

func BenchmarkNativeCall(b *testing.B) {
	for i := 0; i < b.N; i++ {
		foo(0)
	}
}

func BenchmarkRelectCall(b *testing.B) {
	v := reflect.ValueOf(foo)

	b.StopTimer()
	b.ResetTimer()
	b.StartTimer()

	for i := 0; i < b.N; i++ {
		v.Call([]reflect.Value{reflect.ValueOf(2)})
	}
}
