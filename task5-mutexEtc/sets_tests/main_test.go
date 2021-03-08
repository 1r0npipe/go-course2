//Протестируйте производительность операций чтения и записи
//на множестве действительных чисел, безопасность которого
//обеспечивается sync.Mutex и sync.RWMutex для разных вариантов
//использования: 10% запись, 90% чтение; 50% запись, 50% чтение; 90% запись, 10% чтение
package main

import (
	"testing"
)

const (
	base = 1000
	// k10  = int(base * 0.1)
	// k90  = base - k10
	// k50  = base / 2
)

// BenchmarkSetAddM doing the benchmart for ADD in Mutex version
func BenchmarkSetAddMr10rw90(b *testing.B) {
	var set = NewSetM()
	var counter uint = 0
	b.Run("", func(b *testing.B) {
		b.SetParallelism(base)
		b.RunParallel(func(pb *testing.PB) {
			for pb.Next() {
				if counter%10 == 0 {
					set.HasM(1)
					counter++
					continue
				}
				set.AddM(1)
				counter++
			}
		})
	})
}

// BenchmarkSetHasRWM doing the benchmart for has in RWMutex version
func BenchmarkSetHasMr50w50(b *testing.B) {
	var set = NewSetM()
	var counter uint = 0
	b.Run("", func(b *testing.B) {
		b.SetParallelism(base)
		b.RunParallel(func(pb *testing.PB) {
			for pb.Next() {
				if counter%5 == 0 {
					set.HasM(1)
					counter++
					continue
				}
				set.AddM(1)
				counter++
			}
		})
	})
}

// BenchmarkSetAddRWMw10 write skew 10 for RWMutex
func BenchmarkSetAddMw10r90(b *testing.B) {
	var set = NewSetM()
	var counter uint = 0
	b.Run("", func(b *testing.B) {
		b.SetParallelism(base)
		b.RunParallel(func(pb *testing.PB) {
			for pb.Next() {
				if counter%10 == 0 {
					set.AddM(1)
					counter++
					continue
				}
				set.HasM(1)
				counter++
			}
		})
	})
}
