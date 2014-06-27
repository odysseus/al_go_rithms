package main

import (
  "fmt"
  "math/rand"
  "time"
)

func main() {
  t := time.Now()

  fmt.Println("Hello, world!")

  x := randomArray(1000000, 1000000)
  QuickSort(x)
  fmt.Println(x[0:20])

  fmt.Println(compareSort("Quick", "Merge", 1000000, 1))

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

func BubbleSort(a []int) {
  done := false
  for !done {
    done = true
    for i:=1; i<len(a); i++ {
      if a[i-1] > a[i] {
        exch(a, i-1, i)
        done = false
      }
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

// MergeSort and helper methods
func MergeSort(a []int) {
  aux := make([]int, len(a), len(a))
  mergeRecur(a, aux, 0, len(a)-1)
}

func MergeBUSort(a []int) {
  aux := make([]int, len(a), len(a))
  for sz:=1; sz<len(a); sz*=2 {
    for lo:=0; lo<len(a)-sz; lo+=sz*2 {
      mid := lo+sz-1
      if mid > 0 && a[mid-1] <= a[mid] {
        merge(a, aux, lo, mid, min(lo+sz*2-1, len(a)-1))
      }
    }
  }
}

func mergeRecur(a, aux []int, lo, hi int) {
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
    mergeRecur(a, aux, lo, mid)
    mergeRecur(a, aux, mid + 1, hi)
    merge(a, aux, lo, mid, hi)
  }
}

func merge(a, aux []int, lo, mid, hi int) {
  i := lo
  j := mid + 1
  for k:=lo; k<=hi; k++ {
    aux[k] = a[k]
  }
  for k:=lo; k<=hi; k++ {
    if i > mid {
      a[k] = aux[j]
      j++
    } else if j > hi {
      a[k] = aux[i]
      i++
    } else if less(aux[j], aux[i]) {
      a[k] = aux[j]
      j++
    } else {
      a[k] = aux[i]
      i++
    }
  }
}

func QuickSort(a []int) {
  var quickRecur func([]int, int, int)
  quickRecur = func(a []int, lo, hi int) {
    if lo >= hi { return }
    if (hi-lo) == 1 {
      if a[lo] > a[hi] {
        exch(a, lo, hi)
        return
      }
    }
    lp := lo+1
    rp := hi
    pivot := (lo + hi) / 2
    exch(a, lo, pivot)
    for rp > lp {
      for lp <= hi && a[lp] <= a[lo] {
        lp++
      }
      for rp >= lo && a[rp] > a[lo] {
        rp--
      }
      if lp < rp {
        exch(a, lp, rp)
      }
    }
    exch(a, rp, lo)
    quickRecur(a, lo, rp-1)
    quickRecur(a, rp+1, hi)
  }
  quickRecur(a, 0, len(a)-1)
}

// Helper Functions used by many of the algorithms
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

// Testing Helper Functions
// Creating random numbers and arrays to sort
func randInRange(maxRange int) int {
  return int(rand.Float32() * float32(maxRange))
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
  } else if alg == "Bubble" {
    t := time.Now()
    BubbleSort(rarr)
    return time.Since(t).Seconds()
  } else if alg == "Quick" {
    t := time.Now()
    QuickSort(rarr)
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
  result := fmt.Sprintf("For %d random ints:\n%s is %0.3f times faster than %s", size, faster, ratio, slower)
  return result
}

