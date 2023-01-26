package basic_benchmark

import (
	"math/rand"
	"strconv"
	"testing"
)

func BenchmarkAddWithDirectImplementations(b *testing.B) {
	for i := 0; i < b.N; i++ {
		addInts(rand.Int(), rand.Int())
		addFloats(rand.Float64(), rand.Float64())
		addStrings(strconv.Itoa(rand.Int()), strconv.Itoa(rand.Int()))
	}
}

func BenchmarkAddWithTypeCastingUsingSwitch(b *testing.B) {
	for i := 0; i < b.N; i++ {
		addWithTypeCastingUsingSwitch(rand.Int(), rand.Int())
		addWithTypeCastingUsingSwitch(rand.Float64(), rand.Float64())
		addWithTypeCastingUsingSwitch(strconv.Itoa(rand.Int()), strconv.Itoa(rand.Int()))
	}
}

func BenchmarkAddWithReflection(b *testing.B) {
	for i := 0; i < b.N; i++ {
		addWithReflection(rand.Int(), rand.Int())
		addWithReflection(rand.Float64(), rand.Float64())
		addWithReflection(strconv.Itoa(rand.Int()), strconv.Itoa(rand.Int()))
	}
}

func BenchmarkAddUsingGenerics(b *testing.B) {
	for i := 0; i < b.N; i++ {
		addUsingGenerics(rand.Int(), rand.Int())
		addUsingGenerics(rand.Float64(), rand.Float64())
		addUsingGenerics(strconv.Itoa(rand.Int()), strconv.Itoa(rand.Int()))
	}
}

func BenchmarkUsingGenericsWithTypeSet(b *testing.B) {
	for i := 0; i < b.N; i++ {
		addUsingGenericsWithTypeSet(rand.Int(), rand.Int())
		addUsingGenericsWithTypeSet(rand.Float64(), rand.Float64())
		addUsingGenericsWithTypeSet(strconv.Itoa(rand.Int()), strconv.Itoa(rand.Int()))
	}
}
