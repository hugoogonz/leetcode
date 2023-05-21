package main

import (
	"fmt"
)

func isPalindrome(num int) bool {
	// Se declara una variable reversedNum para almacenar el número revertido mientras se recorre x
	var reversedNum int
	// Primero, se realiza una verificación inicial para determinar si el número num es negativo o si termina en cero pero no es igual a cero.
	// En ambos casos, el número no puede ser un palíndromo, por lo que se devuelve false inmediatamente.
	// Esto se hace para evitar casos como -121 o 120 que no pueden ser palíndromos.
	if num < 0 || (num%10 == 0 && num != 0) {
		return false
	}
	// Se inicia un bucle for que se ejecuta mientras num sea mayor que reversedNum.
	// Esto se hace porque solo necesitamos revertir la mitad de los dígitos del número para verificar si es un palíndromo.
	for num > reversedNum {
		// Dentro del bucle, se actualiza reversedNum multiplicándolo por 10 y sumando el último dígito de num.
		// Esto se logra utilizando el operador % para obtener el último dígito y el operador / para descartar el último dígito.
		reversedNum = reversedNum*10 + num%10
		num = num / 10
	}

	return num == reversedNum || num == reversedNum/10
}

// func isPalindrome(x int) bool {
// 	str := strconv.Itoa(x)
// 	start := 0
// 	end := len(str) - 1

// 	for start <= end {
// 		if str[start] != str[end] {
// 			return false
// 		}
// 		start++
// 		end--
// 	}

// 	return true
// }

func main() {
	fmt.Println(isPalindrome(121))
	// isPalindrome(121)
}
