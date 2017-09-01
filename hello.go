package main

import (
	"fmt"
	"time"
	"math/rand"
	"math"
	"math/cmplx"
)

func reverse(s string) string {
	r := []rune(s)
	for i, j := 0, len(r)-1; i < len(r)/2; i, j = i+1, j-1 {
		r[i], r[j] = r[j], r[i]
	}
	return string(r)
}

func swap(x, y string) (string, string) {
	return y, x // go can return "implicit" tuples
}

// x and y are returned 
func split(sum int) (x, y int) {
	x = sum * 4 / 9
	y = sum - x

	y = 10
	return
}

var foo, bar bool
var baz bool = true

const Pi2 = 3.14

var (
	ToBe bool = false
	MaxInt uint64 = 1<<64 - 1
	z complex128 = cmplx.Sqrt(-5 + 12i)
	num rune = 123
	//num2 rune = rune(123.45)
	fnum float64 = 123.45
	fnum2 float64 = 123
	fnum3 float64 = float64(123)
	Big = 2 << 30
	Small = Big >> 29
)


// var nums [rune]

func main() {

	fmt.Println("hello")
	fmt.Println(time.Now())
	fmt.Println("favorite number ", rand.Intn(100)) // 100 is not seed; max num
	fmt.Println(math.Sqrt(7))

	fmt.Println(math.Pi)
	fmt.Println("swap")
	fmt.Println(swap("foo", "bar"))
	fmt.Println("split")
	fmt.Println(split(17))
	
	
	fmt.Printf(reverse("Hello, world!"))
	fmt.Printf("\n")
	// fmt.Printf(string(sum))

	fmt.Println(foo, bar, baz)

	fmt.Println(ToBe, MaxInt, z)

	fmt.Println(num)
	// fmt.Println(num2)
	fmt.Println(fnum)
	fmt.Println(fnum2)
	fmt.Println(fnum3)
	fmt.Println(Pi2)
	fmt.Println(Big)
	fmt.Println(Small)


	sum := 0
	for i := 0; i < 10; i++ {
		sum += i;
	}

	fmt.Println(sum)

	sum = 0
	for sum < 1000 {
		sum += sum
	}
	fmt.Println(sum)
	
}
