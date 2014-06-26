package main

import (
  "fmt"
  "math/rand"
  "time"
)

func main() {
  t := time.Now()

  fmt.Println("Hello, world!")

  fmt.Println(compareSort("Shell", "Merge", 100000, 100))

	fmt.Printf("Took: %0.3fs\n", time.Since(t).Seconds())
}

func SelectionSort(a []int) {
  for c:=0; c<len(a); c++ {
    min := a[c]
    mindex := c
    for i:=c; i<len(a); i++ {
      if a[i] < min {
        min = a[i]
        mindex = i
      }
    }
    exch(a, mindex, c)
  }
}

func InsertionSort(a []int) {
  for i:=0; i<len(a); i++ {
    for j:=i; j>0 && less(a[j], a[j-1]); j-- {
      exch(a, j, j-1)
    }
  }
}

func ShellSort(a []int) {
  for h:=len(a)/3; h>=1; h/=3 {
    for i:=h; i<len(a); i++ {
      for j:=i; j>=h && less(a[j], a[j-h]); j-=h {
        exch(a, j, j-h)
      }
    }
  }
}

var Aux []int
func MergeSort(a []int) {
  Aux = make([]int, len(a), len(a))
  sortRecur(a, 0, len(a)-1)
}

func MergeBUSort(a []int) {
  Aux = make([]int, len(a), len(a))
  for sz:=1; sz<len(a); sz*=2 {
    for lo:=0; lo<len(a)-sz; lo+=sz*2 {
      merge(a, lo, lo+sz-1, min(lo+sz*2-1, len(a)-1))
    }
  }
}

func sortRecur(a []int, lo, hi int) {
  if hi <= lo {
    return
  }
  // For small arrays use an insertion sort instead
  if hi - lo <= 10 {
    for i:=lo; i<=hi; i++ {
      for j:=i; j>lo && less(a[j], a[j-1]); j-- {
        exch(a, j, j-1)
      }
    }
  } else {
    mid := lo + (hi - lo) / 2
    sortRecur(a, lo, mid)
    sortRecur(a, mid + 1, hi)
    merge(a, lo, mid, hi)
  }
}

func merge(a []int, lo, mid, hi int) {
  i := lo
  j := mid + 1
  for k:=lo; k<=hi; k++ {
    Aux[k] = a[k]
  }
  for k:=lo; k<=hi; k++ {
    if i > mid {
      a[k] = Aux[j]
      j++
    } else if j > hi {
      a[k] = Aux[i]
      i++
    } else if less(Aux[j], Aux[i]) {
      a[k] = Aux[j]
      j++
    } else {
      a[k] = Aux[i]
      i++
    }
  }
}

func less(a, b int) bool {
  return a < b
}

func min(a, b int) int {
  if a < b {
    return a
  } else {
    return b
  }
}

func exch(a []int, i, j int) {
  a[i], a[j] = a[j], a[i]
}

func randomArray(size, maxRange int) []int {
  rarr := make([]int, size, size)
  for i:=0; i<len(rarr); i++ {
    r := int(rand.Float32() * float32(maxRange))
    rarr[i] = r
  }
  return rarr
}

func sampleSort(alg string, size int) float64 {
  rarr := randomArray(size, size)
  if alg == "Selection" {
  t := time.Now()
    SelectionSort(rarr)
    return time.Since(t).Seconds()
  } else if alg == "Insertion" {
  t := time.Now()
    InsertionSort(rarr)
    return time.Since(t).Seconds()
  } else if alg == "Shell" {
    t := time.Now()
    ShellSort(rarr)
    return time.Since(t).Seconds()
  } else if alg == "Merge" {
    t := time.Now()
    MergeSort(rarr)
    return time.Since(t).Seconds()
  } else if alg == "MergeBU" {
    t := time.Now()
    MergeBUSort(rarr)
    return time.Since(t).Seconds()
  }
  return -1.0
}

func nSampleSort(alg string, size, samples int) float64 {
  var total float64 = 0.0
  for i:=0; i<=samples; i++ {
    total += sampleSort(alg, size)
  }
  total /= float64(samples)
  return total
}

func compareSort(alg1, alg2 string, size, samples int) string {
  t1 := nSampleSort(alg1, size, samples)
  t2 := nSampleSort(alg2, size, samples)
  var faster, slower string
  var ratio float64
  if t1 < t2 {
    faster = alg1
    slower = alg2
    ratio = t2 / t1
  } else {
    faster = alg2
    slower = alg1
    ratio = t1 / t2
  }
  result := fmt.Sprintf("For %d random ints:\n%s is %0.3f faster than %s", size, faster, ratio, slower)
  return result
}

