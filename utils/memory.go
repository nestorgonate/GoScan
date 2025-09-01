package utils

import (
	"fmt"
	"runtime"
)

func (r *Utils) bytesToMegabytes(b uint64) uint64 {
	return b / 1024 / 1024
}

func (r *Utils) LogMemoryUsage() {
	var m runtime.MemStats
	runtime.ReadMemStats(&m)

	fmt.Printf("\n===== MEMORY USAGE =====\n")
	fmt.Printf("Alloc      : %6d MiB (heap usado actualmente)\n", r.bytesToMegabytes(m.Alloc))
	fmt.Printf("TotalAlloc : %6d MiB (heap asignado desde que se inicio el programa Golang)\n", r.bytesToMegabytes(m.TotalAlloc))
	fmt.Printf("Sys        : %6d MiB (memoria pedida al SO)\n", r.bytesToMegabytes(m.Sys))
	fmt.Printf("NumGC      : %6d (ciclos de Garbage Collector)\n", m.NumGC)
	fmt.Printf("========================\n")
}