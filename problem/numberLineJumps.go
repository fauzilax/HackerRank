package problem

import "fmt"

func Kangaroo(x1 int32, v1 int32, x2 int32, v2 int32) string {
	kgr1 := x1 + v1
	kgr2 := x2 + v2
	for i := 0; i < 100; i++ {
		if kgr1 == kgr2 {
			return "YES"
		}
		fmt.Println(kgr1, kgr2)
		kgr1 += v1

	}
	kgr1 = x1 + v1

	for i := 0; i < 100; i++ {
		if kgr1 == kgr2 {
			return "YES"
		}
		kgr2 += v2

	}
	return "NO"
}
