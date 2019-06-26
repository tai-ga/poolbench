package main

import (
	"fmt"
	"math"
	"sync"
	"testing"

	"github.com/klauspost/cpuid"
)

var counts = []int{}
var formats string
var digit = 8

func init() {
	/*
	   Print basic CPU information:
	*/
	fmt.Println("Name:", cpuid.CPU.BrandName)
	fmt.Println("PhysicalCores:", cpuid.CPU.PhysicalCores)
	fmt.Println("ThreadsPerCore:", cpuid.CPU.ThreadsPerCore)
	fmt.Println("LogicalCores:", cpuid.CPU.LogicalCores)
	fmt.Println("Family", cpuid.CPU.Family, "Model:", cpuid.CPU.Model)
	fmt.Println("Features:", cpuid.CPU.Features)
	fmt.Println("Cacheline bytes:", cpuid.CPU.CacheLine)
	fmt.Println("L1 Data Cache:", cpuid.CPU.Cache.L1D, "bytes")
	fmt.Println("L1 Instruction Cache:", cpuid.CPU.Cache.L1D, "bytes")
	fmt.Println("L2 Cache:", cpuid.CPU.Cache.L2, "bytes")
	fmt.Println("L3 Cache:", cpuid.CPU.Cache.L3, "bytes")

	/*
	   Test if we have a specific feature:
	*/
	if cpuid.CPU.SSE() {
		fmt.Println("We have Streaming SIMD Extensions")
	}
	fmt.Println("---")
}

func init() {
	for i := 1; i <= int(math.Pow10(digit)); i *= 10 {
		counts = append(counts, i)
	}
	formats = "%" + fmt.Sprintf("%dd", digit+1)
}

func BenchmarkWithoutPool(b *testing.B) {
	for _, c := range counts {
		b.Run(fmt.Sprintf(formats, c), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				_ = make([]byte, c)
			}
		})
	}
}

func BenchmarkWithPool(b *testing.B) {
	for _, c := range counts {
		b.Run(fmt.Sprintf(formats, c), func(b *testing.B) {
			pool := sync.Pool{
				New: func() interface{} {
					return make([]byte, c)
				},
			}
			for i := 0; i < b.N; i++ {
				p := pool.Get().([]byte)
				pool.Put(p)
			}
		})
	}
}
