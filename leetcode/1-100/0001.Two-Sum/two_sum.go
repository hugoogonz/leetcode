package main

import "fmt"

// Esta implementación tiene una complejidad temporal de O(n), donde n es la longitud del arreglo nums, ya que el tiempo de búsqueda y de inserción en el map se realiza en tiempo constante en promedio.
func twoSum(nums []int, target int) []int {
	// Se crea un map "newMap" para almacenar los valores y sus índices correspondientes.
	newMap := make(map[int]int)
	// Se recorre el arreglo nums utilizando un bucle for junto con el índice "i" y el valor "value" en cada iteración.
	for i, value := range nums {
		// Se calcula el complemento "c" restando target - value
		c := target - value
		// se verifica si el complemento "c" esta presente en el map "m"
		if index, ok := newMap[c]; ok {
			// Si el complemento "c" está presente en el map, se devuelve un arreglo con los índices index e "i", ya que estos dos números suman el objetivo.
			return []int{index, i}
		}
		// Si el complemento "c" no está presente en el map, se agrega el valor actual (value) al map "m" con su índice "i", para que pueda ser utilizado como complemento en futuras iteraciones.
		newMap[value] = i
	}
	// Si no se encuentra ninguna solución después de recorrer todo el arreglo, se devuelve un arreglo vacío
	return []int{}
}

func main() {
	fmt.Println(twoSum([]int{3, 2, 4}, 6))
	fmt.Println(twoSum([]int{2, 7, 11, 15}, 9))
	fmt.Println(twoSum([]int{7, 6, 5, 3, 2, 1, 4, 9, 10}, 17))
}
