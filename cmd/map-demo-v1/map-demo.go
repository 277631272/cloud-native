package main

import (
	"fmt"
	"github.com/pkg/errors"
	"math/rand"
	"reflect"
	"strings"
	"sync"
	"time"
)

var clusterFlag sync.Map

func main() {
	// sync.Map
	{
		key := "cluster-01"
		t, ok := clusterFlag.Load(key)
		if ok {
			fmt.Printf("key1: %s exist, t: %v\n", key, t)
		} else {
			fmt.Printf("key1: %v not exist, t: %v\n", key, t)
			if t == nil {
				fmt.Printf("t is nil: %v\n", t)
			}
		}

		var value time.Time = time.Now()
		clusterFlag.Store(key, value)
		t, ok = clusterFlag.Load(key)
		if ok {
			fmt.Printf("key2: %s exist, t: %v, diff: %v\n", key, t, time.Since(t.(time.Time)))
		} else {
			fmt.Printf("key2: %v not exist\n", key)
		}
	}

	// time
	{
		rand.Seed(time.Now().Unix())
		t1 := time.Duration(rand.Int63n(30))
		randTime := t1 * time.Minute
		fmt.Printf("t1: %d, randTime: %d\n", t1, randTime)
	}

	// slice
	{
		foo := make([]int, 5)
		foo[3] = 42
		foo[4] = 100

		bar := foo[1:4]
		bar[1] = 99

		for i, v := range foo {
			fmt.Printf("foo[%d] = %d\n", i, v)
		}
		fmt.Printf("foo.len: %d, cap: %d\n", len(foo), cap(foo))

		for i, v := range bar {
			fmt.Printf("bar[%d] = %d\n", i, v)
		}
		fmt.Printf("bar.len: %d, cap: %d\n", len(bar), cap(bar))
	}
	{
		a := make([]int, 32)
		b := a[1:16]
		a = append(a, 1)

		for i, v := range a {
			fmt.Printf("a[%d] = %d\n", i, v)
		}
		fmt.Printf("a.len: %d, cap: %d\n", len(a), cap(a))

		for i, v := range b {
			fmt.Printf("b[%d] = %d\n", i, v)
		}
		fmt.Printf("b.len: %d, cap: %d\n", len(b), cap(b))
	}

	// DeepEquel
	{
		m1 := map[string]string{"one": "a", "two": "b"}
		m2 := map[string]string{"two": "b", "one": "a1"}
		fmt.Println("m1 == m2 :", reflect.DeepEqual(m1, m2))

		s1 := []int{1, 2, 3}
		s2 := []int{1, 2, 3}
		fmt.Println("s1 == s2 : ", reflect.DeepEqual(s1, s2))
	}

	// interface
	{
		var _ Shape = (*Square)(nil)
	}

	// error
	{
		err := f2()
		if err != nil {
			fmt.Println(err)
		}
	}

	// map rbac
	{
		rbac := "true"
		productName := "Grafana"
		if rbac == "true" && AllowRbac(productName) {
			fmt.Println("open rbac, productName:", productName)
		} else {
			fmt.Println("close rbac, productName:", productName)
		}
	}

}

const (
	CloudRun = "cloudrun"
	Tps      = "prometheus"
	Tps2     = "grafana"
	Tme      = "tem"
)

func AllowRbac(productName string) bool {
	fmt.Println("productName1:", productName)
	productName = strings.ToLower(productName)
	fmt.Println("productName2:", productName)

	if productName == CloudRun || productName == Tps || productName == Tps2 || productName == Tme {
		return false
	}
	return true
}

func f1() error {
	return errors.New("f1v() invalid string")
}
func f2() error {
	if err := f1(); err != nil {
		return errors.Wrapf(err, "f2v(), %d", 123)
	}
	return nil
}

type Shape interface {
	Sides() int
	Area() int
}

type Square struct {
	len int
}

func (s *Square) Sides() int {
	return 8
}

func (s *Square) Area() int {
	panic("implement me")
}

func (s *Square) Side() int {
	return 4
}
