package main

import (
	"fmt"
)

func basicArithmetic(a, b int) (int, int, int, float64, int) {
	add := a + b
	sub := a - b
	mul := a * b
	div := float64(a) / float64(b)
	mod := a % b
	return add, sub, mul, div, mod
}

func compareTwoNumbers(a, b int) (bool, bool, bool, bool, bool, bool) {
	return a > b, a < b, a >= b, a <= b, a == b, a != b
}

func logicOperators(a, b bool) (bool, bool, bool) {
	return a && b, a || b, !a
}

func checkNumber(n int) string {
	if n > 0 {
		return "Positive"
	} else if n < 0 {
		return "Negative"
	}
	return "Zero"
}

func classifyCharacter(c rune) string {
	switch c {
	case 'a', 'e', 'i', 'o', 'u', 'A', 'E', 'I', 'O', 'U':
		return "Vowel"
	default:
		return "Consonant"
	}
}

func grade(score int) string {
	if score >= 90 {
		return "A"
	} else if score >= 80 {
		return "B"
	} else if score >= 70 {
		return "C"
	} else if score >= 60 {
		return "D"
	}
	return "F"
}

func printNaturalNumbers() {
	for i := 1; i <= 10; i++ {
		fmt.Print(i, " ")
	}
	fmt.Println()
}

func factorial(n int) int {
	result := 1
	for i := 1; i <= n; i++ {
		result *= i
	}
	return result
}

func multiplicationTable() {
	for i := 1; i <= 10; i++ {
		for j := 1; j <= 10; j++ {
			fmt.Printf("%d*%d=%d ", i, j, i*j)
		}
		fmt.Println()
	}
}

func checkEvenOdd(n int) string {
	if n%2 == 0 {
		return "Even"
	}
	return "Odd"
}

var globalVar = 10
func localScopeDemo() {
	localVar := 5
	fmt.Println("Global Variable:", globalVar)
	fmt.Println("Local Variable:", localVar)
}

func gcd(a, b int) int {
	if b == 0 {
		return a
	}
	return gcd(b, a%b)
}

func statistics(arr []int) (int, float64, int, int) {
	sum := 0
	min := arr[0]
	max := arr[0]

	for _, value := range arr {
		sum += value
		if value < min {
			min = value
		}
		if value > max {
			max = value
		}
	}

	avg := float64(sum) / float64(len(arr))
	return sum, avg, max, min
}

func reverseString(s string) string {
	runes := []rune(s)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}

func fibonacci(n int) []int {
	if n < 1 {
		return []int{}
	}
	fib := make([]int, n)
	fib[0] = 0
	if n > 1 {
		fib[1] = 1
		for i := 2; i < n; i++ {
			fib[i] = fib[i-1] + fib[i-2]
		}
	}
	return fib
}

func main() {
	add, sub, mul, div, mod := basicArithmetic(3, 5)
	fmt.Printf("Basic Arithmetic - Add: %d, Sub: %d, Mul: %d, Div: %.2f, Mod: %d\n", add, sub, mul, div, mod)

	gt, lt, gte, lte, eq, neq := compareTwoNumbers(3, 5)
	fmt.Printf("Comparison - Greater: %t, Less: %t, Greater or Equal: %t, Less or Equal: %t, Equal: %t, Not Equal: %t\n", gt, lt, gte, lte, eq, neq)

	and, or, not := logicOperators(true, false)
	fmt.Printf("Logic Operators - AND: %t, OR: %t, NOT: %t\n", and, or, not)

	fmt.Println("Check Number:", checkNumber(3))
	fmt.Println("Check Number:", checkNumber(-5))
	fmt.Println("Check Number:", checkNumber(0))

	fmt.Println("Character Classification:", classifyCharacter('a'))
	fmt.Println("Character Classification:", classifyCharacter('b'))

	fmt.Println("Grade:", grade(85))
	fmt.Println("Grade:", grade(92))

	fmt.Println("Natural Numbers:")
	printNaturalNumbers()

	fmt.Println("Factorial of 5:", factorial(5))

	fmt.Println("Multiplication Table:")
	multiplicationTable()

	fmt.Println("Check Even Odd:", checkEvenOdd(4))
	fmt.Println("Check Even Odd:", checkEvenOdd(7))

	fmt.Println("Variable Scope:")
	localScopeDemo()

	fmt.Println("GCD of 18 and 24:", gcd(18, 24))

	arr := []int{1, 2, 3, 4, 5}
	sum, avg, max, min := statistics(arr)
	fmt.Printf("Statistics - Sum: %d, Avg: %.2f, Max: %d, Min: %d\n", sum, avg, max, min)

	fmt.Println("Reversed String:", reverseString("hello"))

	fmt.Println("Fibonacci Series (n=7):", fibonacci(7))
}