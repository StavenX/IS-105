package algorithms

import ("math/rand"
"testing"
"time")

func init() {
	seed := time.Now().Unix()
	rand.Seed(seed)
}
func Bubble_sort_modified(list []int) {
	n := len(list) -1
	for {
		if n == 0 {
			break
		}
		for i := 0; i < len(list) - 1; i++ {
			if list[i] > list[i+1] {
				list[i], list[i+1] = list[i+1], list[i]
			}
		}
		n -= 1
	}
}
func perm(n int) (out []int) {
	for _, v := range rand.Perm(n) {
		out = append(out, v)
	}
	return
}



func BenchmarkMBSort100(b *testing.B) {
	benchmarkBSortModified(100, b)
}

func BenchmarkMBSort1000(b *testing.B) {
	benchmarkBSortModified(1000, b)
}

func BenchmarkMBSort10000(b *testing.B) {
	benchmarkBSortModified(10000, b)
}

func benchmarkBSortModified(i int, b *testing.B){
	for j := 0; j < b.N; j++ {
		b.StopTimer()
		values := perm(i)
		b.StartTimer()
		Bubble_sort_modified(values)
	}
}