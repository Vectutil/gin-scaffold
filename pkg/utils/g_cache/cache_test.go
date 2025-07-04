package g_cache

import (
	"fmt"
	"testing"
)

func TestCacheString(t *testing.T) {
	fmt.Println("---------------------string---------------------------")

	cacheString := NewMapCacheString()
	cacheString.Set("1", "1")
	fmt.Println(cacheString.GetValueString("1")) // true 1
	fmt.Println(cacheString.Get("1"))            // 1
	cacheString.Delete("1")
	fmt.Println(cacheString.GetValueString("1")) // false ""
	fmt.Println(cacheString.Get("1"))            // nil

	fmt.Println("----------------------int--------------------------")

	cacheInt := NewMapCacheInt()
	cacheInt.Set(1, 1)
	fmt.Println(cacheInt.GetValueInt(1)) // true 1
	fmt.Println(cacheInt.Get(1))         // 1
	cacheInt.Delete(1)
	fmt.Println(cacheInt.GetValueInt(1)) // false 0
	fmt.Println(cacheInt.Get(1))         // n0

	fmt.Println("----------------------T string--------------------------")

	cacheT := NewMapCache[string, string]()
	cacheT.Set("1", "1")
	fmt.Println(cacheT.GetValue("1")) // true 1
	fmt.Println(cacheT.Get("1"))      // 1
	cacheT.Delete("1")
	fmt.Println(cacheT.GetValue("1")) // false ""
	fmt.Println(cacheT.Get("1"))      // nil

	fmt.Println("----------------------T int--------------------------")
	cacheTInt := NewMapCache[int, int]()
	cacheTInt.Set(1, 1)
	fmt.Println(cacheTInt.GetValue(1)) // true 1
	fmt.Println(cacheTInt.Get(1))      // 1
	cacheTInt.Delete(1)
	fmt.Println(cacheTInt.GetValue(1)) // false ""
	fmt.Println(cacheTInt.Get(1))      // nil
}
