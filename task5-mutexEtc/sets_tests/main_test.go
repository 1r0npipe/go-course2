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
	k10  = int(base * 0.1)
	k90  = base - k10
	k50  = base / 2
)

// BenchmarkSetAddM doing the benchmart for ADD in Mutex version
func BenchmarkSetAddMw10(b *testing.B) {
	var set = NewSetM()

	b.Run("", func(b *testing.B) {
		b.SetParallelism(k10)
		b.RunParallel(func(pb *testing.PB) {
			for pb.Next() {
				set.AddM(1)
			}
		})
	})
}

// BenchmarkSetHasRWM doing the benchmart for has in RWMutex version
func BenchmarkSetHasMr90(b *testing.B) {
	var set = NewSetM()
	b.Run("", func(b *testing.B) {
		b.SetParallelism(k90)
		b.RunParallel(func(pb *testing.PB) {
			for pb.Next() {
				set.HasM(1)
			}
		})
	})
}

// BenchmarkSetAddRWMw10 write skew 10 for RWMutex
func BenchmarkSetAddRWMw10(b *testing.B) {
	var set = NewSetRWM()

	b.Run("", func(b *testing.B) {
		b.SetParallelism(k10)
		b.RunParallel(func(pb *testing.PB) {
			for pb.Next() {
				set.AddRWM(1)
			}
		})
	})
}

// BenchmarkSetHasRWM doing the benchmart for has in RWMutex version
func BenchmarkSetHasRWMr90(b *testing.B) {
	var set = NewSetRWM()
	b.Run("", func(b *testing.B) {
		b.SetParallelism(k90)
		b.RunParallel(func(pb *testing.PB) {
			for pb.Next() {
				set.HasRWM(1)
			}
		})
	})
}

//======================================================
// section to 50/50
//
// BenchmarkSetAddM doing the benchmart for ADD in Mutex version
func BenchmarkSetAddMw50(b *testing.B) {
	var set = NewSetM()

	b.Run("", func(b *testing.B) {
		b.SetParallelism(k50)
		b.RunParallel(func(pb *testing.PB) {
			for pb.Next() {
				set.AddM(1)
			}
		})
	})
}

// BenchmarkSetHasRWM doing the benchmart for has in RWMutex version
func BenchmarkSetHasMr50(b *testing.B) {
	var set = NewSetM()
	b.Run("", func(b *testing.B) {
		b.SetParallelism(k50)
		b.RunParallel(func(pb *testing.PB) {
			for pb.Next() {
				set.HasM(1)
			}
		})
	})
}

// BenchmarkSetAddRWMw50 write 50 skew for RWMutex
func BenchmarkSetAddRWMw50(b *testing.B) {
	var set = NewSetRWM()

	b.Run("", func(b *testing.B) {
		b.SetParallelism(k10)
		b.RunParallel(func(pb *testing.PB) {
			for pb.Next() {
				set.AddRWM(1)
			}
		})
	})
}

// BenchmarkSetHasRWM doing the benchmart for has in RWMutex version
func BenchmarkSetHasRWMr50(b *testing.B) {
	var set = NewSetRWM()
	b.Run("", func(b *testing.B) {
		b.SetParallelism(k50)
		b.RunParallel(func(pb *testing.PB) {
			for pb.Next() {
				set.HasRWM(1)
			}
		})
	})
}

//======================================================
// section to 90/10 (WRITE/READ)
//
// BenchmarkSetAddM doing the benchmart for ADD in Mutex version
func BenchmarkSetAddMw90(b *testing.B) {
	var set = NewSetM()

	b.Run("", func(b *testing.B) {
		b.SetParallelism(k90)
		b.RunParallel(func(pb *testing.PB) {
			for pb.Next() {
				set.AddM(1)
			}
		})
	})
}

// BenchmarkSetHasRWM doing the benchmart for has in RWMutex version
func BenchmarkSetHasMr10(b *testing.B) {
	var set = NewSetM()
	b.Run("", func(b *testing.B) {
		b.SetParallelism(k10)
		b.RunParallel(func(pb *testing.PB) {
			for pb.Next() {
				set.HasM(1)
			}
		})
	})
}

// BenchmarkSetAddRWMw50 write 50 skew for RWMutex
func BenchmarkSetAddRWMw90(b *testing.B) {
	var set = NewSetRWM()

	b.Run("", func(b *testing.B) {
		b.SetParallelism(k90)
		b.RunParallel(func(pb *testing.PB) {
			for pb.Next() {
				set.AddRWM(1)
			}
		})
	})
}

// BenchmarkSetHasRWM doing the benchmart for has in RWMutex version
func BenchmarkSetHasRWMr10(b *testing.B) {
	var set = NewSetRWM()
	b.Run("", func(b *testing.B) {
		b.SetParallelism(k10)
		b.RunParallel(func(pb *testing.PB) {
			for pb.Next() {
				set.HasRWM(1)
			}
		})
	})
}
