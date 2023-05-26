package main

import "fmt"

func romanToInt(s string) int {
	romanNumber := map[byte]int{
		'I': 1,
		'V': 5,
		'X': 10,
		'L': 50,
		'C': 100,
		'D': 500,
		'M': 1000,
	}

	// Esta variable se utilizará para almacenar el resultado final de la conversión de la cadena romana a un número entero.
	result := 0
	// Esta variable se utiliza para realizar la comparación entre el valor actual y el valor previo durante el proceso de conversión
	prevValue := 0

	for i := len(s) - 1; i >= 0; i-- {
		// En cada iteración del bucle, se obtiene el símbolo romano actual utilizando s[i], donde i es el índice actual en la iteración.
		symbol := s[i]
		// Se obtiene el valor correspondiente al símbolo romano actual del mapa romanNumber utilizando romanNumber[symbol].
		value := romanNumber[symbol]

		// Si el valor actual es menor que el valor previo, significa que se trata de un caso especial donde se debe restar el valor actual al resultado.
		// En caso contrario, se suma el valor actual al resultado.
		if value < prevValue {
			result -= value
		} else {
			result += value
		}
		// Después de realizar la operación correspondiente, se actualiza el valor previo (prevValue) con el valor actual (value).
		prevValue = value
	}

	// Retorna el resultado final
	return result
}

func main() {
	fmt.Println(romanToInt("III"))         // Output: 3
	fmt.Println(romanToInt("IV"))          // Output: 4
	fmt.Println(romanToInt("IX"))          // Output: 9
	fmt.Println(romanToInt("LVIII"))       // Output: 58
	fmt.Println(romanToInt("MCMXCIV"))     // Output: 1994
	fmt.Println(romanToInt("MMMMMCMXCIX")) // Output: 5999
}
