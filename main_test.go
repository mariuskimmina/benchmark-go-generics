package main

import (
	"math/rand"
	"strconv"
	"testing"
)

func BenchmarkAddExplicit(b *testing.B) {
    for i := 0; i < b.N; i++ {
        addInt(rand.Int(), rand.Int())
        addFloat(rand.Float64(), rand.Float64())
        addString(strconv.Itoa(rand.Int()), strconv.Itoa(rand.Int())) 
    }
}

func BenchmarkAddTypeSwitch(b *testing.B) {
    for i := 0; i < b.N; i++ {
        addTypeSwitch(rand.Int(), rand.Int())
        addTypeSwitch(rand.Float64(), rand.Float64())
        addTypeSwitch(strconv.Itoa(rand.Int()), strconv.Itoa(rand.Int())) 
    }
}

func BenchmarkAddGenerics(b *testing.B) {
    for i := 0; i < b.N; i++ {
        addGenerics(rand.Int(), rand.Int())
        addGenerics(rand.Float64(), rand.Float64())
        addGenerics(strconv.Itoa(rand.Int()), strconv.Itoa(rand.Int())) 
    }
}

func BenchmarkAddGenericsTypeSet(b *testing.B) {
    for i := 0; i < b.N; i++ {
        addGenericsTypeSet(rand.Int(), rand.Int())
        addGenericsTypeSet(rand.Float64(), rand.Float64())
        addGenericsTypeSet(strconv.Itoa(rand.Int()), strconv.Itoa(rand.Int())) 
    }
}

func BenchmarkAddReflect(b *testing.B) {
    for i := 0; i < b.N; i++ {
        addReflect(rand.Int(), rand.Int())
        addReflect(rand.Float64(), rand.Float64())
        addReflect(strconv.Itoa(rand.Int()), strconv.Itoa(rand.Int())) 
    }
}
