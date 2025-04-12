package main

import (
	"fmt"
	"sync"
	"time"
)

/**
Grabbed from a SO answer then YOLO converted to go using AI.
Who knows if its performant but it does seem to work.
Gets all permutations of a set of strings brute force using DP.
Nice for verifying results
*/

// "jack", "apple", "maven", "hold",
/*
moon
1
poor
4
solid
280
mark
4312
spark
1491
live
19278
jack
697896
apple
10710432
maven

*/
//var STRINGS = []string{"moon", "poor", "solid", "mark", "spark", "live", "jack", "apple", "maven", "hold"}
var STRINGS = []string{"por", "solid", "mon"}

type Set map[string]struct{}

func NewSet(strings ...string) Set {
	s := Set{}
	for _, str := range strings {
		s[str] = struct{}{}
	}
	return s
}

func (s Set) Add(str string) {
	s[str] = struct{}{}
}

func (s Set) AddAll(other Set) {
	for str := range other {
		s[str] = struct{}{}
	}
}

func MapSet(set Set, f func(string) string) Set {
	result := Set{}
	for str := range set {
		result[f(str)] = struct{}{}
	}
	return result
}

func Shortest(set Set) (Set, int) {
	minLen := -1
	result := Set{}
	for str := range set {
		if minLen == -1 || len(str) < minLen {
			minLen = len(str)
			result = Set{str: {}}
		} else if len(str) == minLen {
			result[str] = struct{}{}
		}
	}
	return result, minLen
}

type ZipCache struct {
	cache map[string]map[string]Set
}

func NewZipCache() *ZipCache {
	return &ZipCache{
		cache: map[string]map[string]Set{},
	}
}

func (zc *ZipCache) Get(str1, str2 string) (Set, bool) {
	if m, ok := zc.cache[str1]; ok {
		if val, ok2 := m[str2]; ok2 {
			return val, true
		}
	}
	return nil, false
}

func (zc *ZipCache) Set(str1, str2 string, zipped Set) {
	if _, ok := zc.cache[str1]; !ok {
		zc.cache[str1] = map[string]Set{}
	}
	zc.cache[str1][str2] = zipped
}

var zipCache = NewZipCache()

func Zip(str1, str2 string) Set {
	if cached, ok := zipCache.Get(str1, str2); ok {
		return cached
	}

	if len(str1) == 0 {
		result := NewSet(str2)
		zipCache.Set(str1, str2, result)
		return result
	}
	if len(str2) == 0 {
		result := NewSet(str1)
		zipCache.Set(str1, str2, result)
		return result
	}
	if str1[0] == str2[0] {
		zipped := Zip(str1[1:], str2[1:])
		result := MapSet(zipped, func(s string) string {
			return string(str1[0]) + s
		})
		zipCache.Set(str1, str2, result)
		return result
	}

	zip1 := Zip(str1[1:], str2)
	zip2 := Zip(str1, str2[1:])
	test1 := MapSet(zip1, func(s string) string {
		return string(str1[0]) + s
	})
	test2 := MapSet(zip2, func(s string) string {
		return string(str2[0]) + s
	})
	test1.AddAll(test2)
	result, _ := Shortest(test1)
	zipCache.Set(str1, str2, result)
	return result
}

func main() {
	cumulative := NewSet("")
	for _, s := range STRINGS {
		fmt.Printf("\n%s\nStarting %s\n", time.Now(), s)
		newCumulative := Set{}
		wg := sync.WaitGroup{}
		for test := range cumulative {
			wg.Add(1)
			test := test
			go func() {
				z := Zip(test, s)
				newCumulative.AddAll(z)
				wg.Done()
			}()
			wg.Wait()
		}
		res, size := Shortest(newCumulative)
		cumulative = res
		fmt.Printf("Completed word %s with %d combinations at length %d\n", s, len(cumulative), size)
	}
	fmt.Println("Final cumulative:", len(cumulative))
	for s := range cumulative {
		fmt.Println("Example resul", s)
		break
	}
}
