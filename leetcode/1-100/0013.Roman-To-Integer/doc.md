## Veamos un ejemplo más detallado para ilustrar el funcionamiento del código:

```go
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

	result := 0
	prevValue := 0

	for i := len(s) - 1; i >= 0; i-- {
		symbol := s[i]
		value := romanNumber[symbol]

		if value < prevValue {
			result -= value
		} else {
			result += value
		}

		prevValue = value
	}

	return result
}

func main() {
	fmt.Println(romanToInt("XIV"))         // Output: 14

}
```

Supongamos que tenemos la cadena romana `"XIV"`:

Iniciamos con la cadena `"XIV"` como entrada para la función `romanToInt`.

El mapa `romanNumber` se crea con los valores correspondientes para cada símbolo romano.

La variable `result se inicializa en 0` y la variable `prevValue también se inicializa en 0`.

Comenzamos el bucle for que recorre la cadena en orden inverso.

En la primera iteración, `el índice i es 2`. Accedemos al símbolo romano en la posición 2, que es 'V'. Obtenemos su valor correspondiente del mapa romanNumber, que es 5.

Al ser la primera iteración, no hay un valor previo con el que comparar, por lo que simplemente sumamos el valor actual al resultado. Ahora, `result es 5` y `prevValue se actualiza a 5`.

En la siguiente iteración, `el índice i es 1`. Accedemos al símbolo romano en la posición 1, que es 'I'. Obtenemos su valor correspondiente del mapa romanNumber, que es 1.

Comparamos el valor actual (1) con el valor previo (5). Como el valor actual es menor que el valor previo, restamos el valor actual al resultado. Ahora, result es 4 y prevValue se mantiene como 5.

En la última iteración, el índice i es 0. Accedemos al símbolo romano en la posición 0, que es 'X'. Obtenemos su valor correspondiente del mapa romanNumber, que es 10.

Comparamos el valor actual (10) con el valor previo (1). Como el valor actual es mayor que el valor previo, simplemente sumamos el valor actual al resultado. Ahora, result es 14 y prevValue se actualiza a 10.

El bucle for ha terminado de recorrer la cadena y ahora se devuelve el resultado final, que es 14.