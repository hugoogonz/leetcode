# 9. Palindrome Number

`Easy`

Given an integer `x`, return `true` if `x` is a palindrome, and `false` otherwise.
 
**Example 1:**

```
Input: x = 121
Output: true
Explanation: 121 reads as 121 from left to right and from right to left.
```

**Example 2:**
```
Input: x = -121
Output: false
Explanation: From left to right, it reads -121. From right to left, it becomes 121-. Therefore it is not a palindrome.
```

**Example 3:**

```
Input: x = -121
Output: false
Explanation: From left to right, it reads -121. From right to left, it becomes 121-. Therefore it is not a palindrome.
```

**Constraints:**

```
-231 <= x <= 231 - 1
```

**Solution:**

```go
func isPalindrome(x int) bool {
    var reversedNum int
    if x < 0 || (x % 10 == 0 && x != 0) {
        return false
    }
    
    for x > reversedNum {
        reversedNum = reversedNum*10 + x % 10
        x = x / 10
    }

    return x == reversedNum || x == reversedNum / 10
}

```

```
Ejemplo detallado con el número 151:

Comenzamos con x = 151.
En la primera iteración del bucle, reversedNum se actualiza de la siguiente manera:
reversedNum = 0 * 10 + 151 % 10 = 0 + 1 = 1
En este punto, reversedNum contiene el primer dígito del número revertido.

En la segunda iteración del bucle, reversedNum se actualiza de la siguiente manera:
reversedNum = 1 * 10 + 15 % 10 = 10 + 5 = 15
Ahora, reversedNum contiene los dos primeros dígitos del número revertido.

En la tercera iteración del bucle, reversedNum se actualiza de la siguiente manera:
reversedNum = 15 * 10 + 1 % 10 = 150 + 1 = 151
Después de esta iteración, reversedNum contiene el número revertido completo.

El bucle finaliza cuando x se queda solo con el dígito del medio y reversedNum contiene la primera mitad del número revertido.
```