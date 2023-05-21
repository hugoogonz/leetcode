# 1. Two Sum

`Easy`

Given an array of integers `nums` and an integer `target`, return indices of the two numbers such that they add up to `target`.

You may assume that each input would have exactly one solution, and you may not use the same element twice.

You can return the answer in any order.
 
**Example 1:**

```
Input: nums = [2,7,11,15], target = 9
Output: [0,1]
Explanation: Because nums[0] + nums[1] == 9, we return [0, 1].
```

**Example 2:**
```
Input: nums = [3,2,4], target = 6
Output: [1,2]
```

**Example 3:**

```
Input: nums = [3,3], target = 6
Output: [0,1]
```

**Constraints:**

```
2 <= nums.length <= 104
-109 <= nums[i] <= 109
-109 <= target <= 109
```

**Solution:**

```go
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
```

Complejidad temporal de O(n^2):

```go
func twoSum(nums []int, target int) []int {
	for i, value := range nums {
		for j := i + 1; j < len(nums); j++ {
			if value + nums[j] == target {
				return []int{i, j}
			}
		}
	}
	return []int{}
}
```