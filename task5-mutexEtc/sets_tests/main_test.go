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
				if counter%2 == 0 {
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
// BenchmarkSetAddRWMr10rw90 test
func BenchmarkSetAddRWMr10rw90(b *testing.B) {
	var set = NewSetRWM()
	var counter uint = 0
	b.Run("", func(b *testing.B) {
		b.SetParallelism(base)
		b.RunParallel(func(pb *testing.PB) {
			for pb.Next() {
				if counter%10 == 0 {
					set.HasRWM(1)
					counter++
					continue
				}
				set.AddRWM(1)
				counter++
			}
		})
	})
}

// BenchmarkSetHasRWMr50w50 doing the benchmart for has in RWMutex version
func BenchmarkSetHasRWMr50w50(b *testing.B) {
	var set = NewSetRWM()
	var counter uint = 0
	b.Run("", func(b *testing.B) {
		b.SetParallelism(base)
		b.RunParallel(func(pb *testing.PB) {
			for pb.Next() {
				if counter%2 == 0 {
					set.HasRWM(1)
					counter++
					continue
				}
				set.AddRWM(1)
				counter++
			}
		})
	})
}

// BenchmarkSetAddRWMw10r90 write skew 10 for RWMutex
func BenchmarkSetAddRWMw10r90(b *testing.B) {
	var set = NewSetRWM()
	var counter uint = 0
	b.Run("", func(b *testing.B) {
		b.SetParallelism(base)
		b.RunParallel(func(pb *testing.PB) {
			for pb.Next() {
				if counter%10 == 0 {
					set.AddRWM(1)
					counter++
					continue
				}
				set.HasRWM(1)
				counter++
			}
		})
	})
}
