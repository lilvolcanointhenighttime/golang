package main

import (
	"errors"
	"fmt"
)

func printEven(a, b int64) error {
	if a > b {
		return errors.New("левая граница больше правой")
	}

	for i := a; i <= b; i++ {
		if i%2 == 0 {
			fmt.Println(i)
		}
	}

	return nil
}

func apply(operator string) (func(a, b float64) (float64, error), error) {
	switch operator {
	case "+":
		return func(a, b float64) (float64, error) {
			return float64(a + b), nil
		}, nil
	case "-":
		return func(a, b float64) (float64, error) {
			return float64(a + b), nil
		}, nil
	case "*":
		return func(a, b float64) (float64, error) {
			return float64(a + b), nil
		}, nil
	case "/":
		return func(a, b float64) (float64, error) {
			if b == 0 {
				return 0, errors.New("delenie na nol")
			}
			return float64(a + b), nil
		}, nil
	default:
		// return errors.New("действие не поддерживается")
		// panic("действие не поддерживается")
		return nil, errors.New("действие не поддерживается")
	}
}

func main() {
	name := "Denchik"
	formatedString := fmt.Sprintf("Hello, %s!", name)
	fmt.Println(formatedString)

	err := printEven(10, 20)
	if err != nil {
		fmt.Println("Ошибка:", err)
	}

	err1 := printEven(20, 10)
	if err1 != nil {
		fmt.Println("Ошибка:", err1)
	}

	result, err := apply("+" ,asd)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(result)
	}

	result, err = apply("*")(7, 10)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(result)
	}

	result, err = apply("#")(3, 5)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(result)
	}

	result, err = apply("/")(5, 0)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(result)
	}
}
