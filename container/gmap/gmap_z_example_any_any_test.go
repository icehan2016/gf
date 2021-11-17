// Copyright GoFrame Author(https://goframe.org). All Rights Reserved.
//
// This Source Code Form is subject to the terms of the MIT License.
// If a copy of the MIT was not distributed with gm file,
// You can obtain one at https://github.com/gogf/gf.

package gmap_test

import (
	"fmt"

	"github.com/gogf/gf/v2/container/gmap"
	"github.com/gogf/gf/v2/frame/g"
)

func ExampleAnyAnyMap_Iterator() {
	m := gmap.New()
	for i := 0; i < 10; i++ {
		m.Set(i, i*2)
	}

	var totalKey, totalValue int
	m.Iterator(func(k interface{}, v interface{}) bool {
		totalKey += k.(int)
		totalValue += v.(int)

		return totalKey < 10
	})

	fmt.Println("totalKey:", totalKey)
	fmt.Println("totalValue:", totalValue)

	// May Output:
	// totalKey: 11
	// totalValue: 22
}

func ExampleAnyAnyMap_Clone() {
	m := gmap.New()

	m.Set("key1", "val1")
	fmt.Println(m)

	n := m.Clone()
	fmt.Println(n)

	// Output:
	// {"key1":"val1"}
	// {"key1":"val1"}
}

func ExampleAnyAnyMap_Map() {
	// non concurrent-safety, a pointer to the underlying data
	m1 := gmap.New()
	m1.Set("key1", "val1")
	fmt.Println("m1:", m1)

	n1 := m1.Map()
	fmt.Println("before n1:", n1)
	m1.Set("key1", "val2")
	fmt.Println("after n1:", n1)

	// concurrent-safety, copy of underlying data
	m2 := gmap.New(true)
	m2.Set("key1", "val1")
	fmt.Println("m1:", m2)

	n2 := m2.Map()
	fmt.Println("before n2:", n2)
	m2.Set("key1", "val2")
	fmt.Println("after n2:", n2)

	// Output:
	// m1: {"key1":"val1"}
	// before n1: map[key1:val1]
	// after n1: map[key1:val2]
	// m1: {"key1":"val1"}
	// before n2: map[key1:val1]
	// after n2: map[key1:val1]
}

func ExampleAnyAnyMap_MapCopy() {
	m := gmap.New()

	m.Set("key1", "val1")
	m.Set("key2", "val2")
	fmt.Println(m)

	n := m.MapCopy()
	fmt.Println(n)

	// Output:
	// {"key1":"val1","key2":"val2"}
	// map[key1:val1 key2:val2]
}

func ExampleAnyAnyMap_MapStrAny() {
	m := gmap.New()
	m.Set(1001, "val1")
	m.Set(1002, "val2")

	n := m.MapStrAny()
	fmt.Println(n)

	// Output:
	// map[1001:val1 1002:val2]
}

func ExampleAnyAnyMap_FilterEmpty() {
	m := gmap.NewFrom(g.MapAnyAny{
		"k1": "",
		"k2": nil,
		"k3": 0,
		"k4": 1,
	})
	m.FilterEmpty()
	fmt.Println(m.Map())

	// May Output:
	// map[k4:1]
}

func ExampleAnyAnyMap_FilterNil() {
	m := gmap.NewFrom(g.MapAnyAny{
		"k1": "",
		"k2": nil,
		"k3": 0,
		"k4": 1,
	})
	m.FilterNil()
	fmt.Println(m.Map())

	// May Output:
	// map[k1: k3:0 k4:1]
}

func ExampleAnyAnyMap_Set() {
	m := gmap.New()

	m.Set("key1", "val1")
	fmt.Println(m)

	// Output:
	// {"key1":"val1"}
}

func ExampleAnyAnyMap_Sets() {
	m := gmap.New()

	addMap := make(map[interface{}]interface{})
	addMap["key1"] = "val1"
	addMap["key2"] = "val2"
	addMap["key3"] = "val3"

	m.Sets(addMap)
	fmt.Println(m)

	// Output:
	// {"key1":"val1","key2":"val2","key3":"val3"}
}

func ExampleAnyAnyMap_Search() {
	m := gmap.New()

	m.Set("key1", "val1")

	value, found := m.Search("key1")
	if found {
		fmt.Println("find key1 value:", value)
	}

	value, found = m.Search("key2")
	if !found {
		fmt.Println("key2 not find")
	}

	// Output:
	// find key1 value: val1
	// key2 not find
}

func ExampleAnyAnyMap_Get() {
	m := gmap.New()

	m.Set("key1", "val1")

	fmt.Println("key1 value:", m.Get("key1"))
	fmt.Println("key2 value:", m.Get("key2"))

	// Output:
	// key1 value: val1
	// key2 value: <nil>
}

func ExampleAnyAnyMap_Pop() {
	var m gmap.Map
	m.Sets(g.MapAnyAny{
		"k1": "v1",
		"k2": "v2",
		"k3": "v3",
		"k4": "v4",
	})

	fmt.Println(m.Pop())

	// May Output:
	// k1 v1
}

func ExampleAnyAnyMap_Pops() {
	var m gmap.Map
	m.Sets(g.MapAnyAny{
		"k1": "v1",
		"k2": "v2",
		"k3": "v3",
		"k4": "v4",
	})
	fmt.Println(m.Pops(-1))
	fmt.Println("size:", m.Size())

	m.Sets(g.MapAnyAny{
		"k1": "v1",
		"k2": "v2",
		"k3": "v3",
		"k4": "v4",
	})
	fmt.Println(m.Pops(2))
	fmt.Println("size:", m.Size())

	// May Output:
	// map[k1:v1 k2:v2 k3:v3 k4:v4]
	// size: 0
	// map[k1:v1 k2:v2]
	// size: 2
}

func ExampleAnyAnyMap_GetOrSet() {
	m := gmap.New()
	m.Set("key1", "val1")

	fmt.Println(m.GetOrSet("key1", "NotExistValue"))
	fmt.Println(m.GetOrSet("key2", "val2"))

	// Output:
	// val1
	// val2
}

func ExampleAnyAnyMap_GetOrSetFunc() {
	m := gmap.New()
	m.Set("key1", "val1")

	fmt.Println(m.GetOrSetFunc("key1", func() interface{} {
		return "NotExistValue"
	}))
	fmt.Println(m.GetOrSetFunc("key2", func() interface{} {
		return "NotExistValue"
	}))

	// Output:
	// val1
	// NotExistValue
}

func ExampleAnyAnyMap_GetOrSetFuncLock() {
	m := gmap.New()
	m.Set("key1", "val1")

	fmt.Println(m.GetOrSetFuncLock("key1", func() interface{} {
		return "NotExistValue"
	}))
	fmt.Println(m.GetOrSetFuncLock("key2", func() interface{} {
		return "NotExistValue"
	}))

	// Output:
	// val1
	// NotExistValue
}

func ExampleAnyAnyMap_GetVar() {
	m := gmap.New()
	m.Set("key1", "val1")

	fmt.Println(m.GetVar("key1"))
	fmt.Println(m.GetVar("key2").IsNil())

	// Output:
	// val1
	// true
}

func ExampleAnyAnyMap_GetVarOrSet() {
	m := gmap.New()
	m.Set("key1", "val1")

	fmt.Println(m.GetVarOrSet("key1", "NotExistValue"))
	fmt.Println(m.GetVarOrSet("key2", "val2"))

	// Output:
	// val1
	// val2
}

func ExampleAnyAnyMap_GetVarOrSetFunc() {
	m := gmap.New()
	m.Set("key1", "val1")

	fmt.Println(m.GetVarOrSetFunc("key1", func() interface{} {
		return "NotExistValue"
	}))
	fmt.Println(m.GetVarOrSetFunc("key2", func() interface{} {
		return "NotExistValue"
	}))

	// Output:
	// val1
	// NotExistValue
}

func ExampleAnyAnyMap_GetVarOrSetFuncLock() {
	m := gmap.New()
	m.Set("key1", "val1")

	fmt.Println(m.GetVarOrSetFuncLock("key1", func() interface{} {
		return "NotExistValue"
	}))
	fmt.Println(m.GetVarOrSetFuncLock("key2", func() interface{} {
		return "NotExistValue"
	}))

	// Output:
	// val1
	// NotExistValue
}

func ExampleAnyAnyMap_SetIfNotExist() {
	var m gmap.Map
	fmt.Println(m.SetIfNotExist("k1", "v1"))
	fmt.Println(m.SetIfNotExist("k1", "v1"))
	fmt.Println(m.Map())

	// Output:
	// true
	// false
	// map[k1:v1]
}

func ExampleAnyAnyMap_SetIfNotExistFunc() {

	// Output:
}

func ExampleAnyAnyMap_SetIfNotExistFuncLock() {

	// Output:
}

func ExampleAnyAnyMap_Remove() {

	// Output:
}

func ExampleAnyAnyMap_Removes() {

	// Output:
}

func ExampleAnyAnyMap_Keys() {
	var m gmap.Map
	m.Sets(g.MapAnyAny{
		"k1": "v1",
		"k2": "v2",
		"k3": "v3",
		"k4": "v4",
	})
	fmt.Println(m.Keys())
	fmt.Println(m.Values())

	// May Output:
	// [k1 k2 k3 k4]
	// [v2 v3 v4 v1]
}

func ExampleAnyAnyMap_Values() {
	var m gmap.Map
	m.Sets(g.MapAnyAny{
		"k1": "v1",
		"k2": "v2",
		"k3": "v3",
		"k4": "v4",
	})
	fmt.Println(m.Keys())
	fmt.Println(m.Values())

	// May Output:
	// [k1 k2 k3 k4]
	// [v2 v3 v4 v1]
}

func ExampleAnyAnyMap_Contains() {

	// Output:
}

func ExampleAnyAnyMap_Size() {

	// Output:
}

func ExampleAnyAnyMap_IsEmpty() {

	// Output:
}

func ExampleAnyAnyMap_Clear() {

	// Output:
}

func ExampleAnyAnyMap_Replace() {

	// Output:
}

func ExampleAnyAnyMap_LockFunc() {

	// Output:
}

func ExampleAnyAnyMap_RLockFunc() {

	// Output:
}

func ExampleAnyAnyMap_Flip() {
	var m gmap.Map
	m.Sets(g.MapAnyAny{
		"k1": "v1",
		"k2": "v2",
	})
	m.Flip()
	fmt.Println(m.Map())

	// May Output:
	// map[v1:k1 v2:k2]
}

func ExampleAnyAnyMap_Merge() {
	var m1, m2 gmap.Map
	m1.Set("key1", "val1")
	m2.Set("key2", "val2")
	m1.Merge(&m2)
	fmt.Println(m1.Map())

	// May Output:
	// map[key1:val1 key2:val2]
}

func ExampleAnyAnyMap_String() {

	// Output:
}

func ExampleAnyAnyMap_MarshalJSON() {

	// Output:
}

func ExampleAnyAnyMap_UnmarshalJSON() {

	// Output:
}

func ExampleAnyAnyMap_UnmarshalValue() {

	// Output:
}
