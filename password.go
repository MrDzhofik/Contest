// Программист Василий работает в ИТ-компании. Он забыл пароль к своему рабочему ноутбуку и теперь боится,
// что попросту не успеет сделать порученные ему задачи.

// В компании действуют строгие требования к пароля на рабочие ноутбуки: для каждого сотрудника определен набор символов,
// только из которых должен состоять пароль, причём каждый символ из набора должен встречаться хотя бы один раз.
// Василий помнит этот набор. Также Василий помнит, что длина его пароля не превосходит kk символов.

// С помощью небольших усилий ему удалось восстановить, какие клавиши он нажимал на клавиатуре за последнее время.
// Теперь у него в распоряжении есть последовательность символов, в которой может оказаться его пароль.
//  Помогите Василию восстановить свой пароль или определите, что восстановить его уже невозможно!

// Формат входных данных

// В первой строке ввода дана последовательность длины nn (1≤n≤2×105)(1≤n≤2×105) из строчных латинских букв — последовательность символов, которые нажимал Василий за последнее время.

// Во второй строке дан набор символов — требования к паролю, а в третьей — число kk (1≤k≤2×105)(1≤k≤2×105), максимальная длина пароля.

// Формат выходных данных

// Выведите возможный пароль от ноутбука, удовлетворяющий указанным условиям. Если вариантов пароля несколько, выберите тот, который начинается в последовательности из первой строки правее (позже) других, а среди всех с одинаковым с ним началом — самый длинный.

// Если восстановить пароль не удастся, выведите «−1−1» (без кавычек).
// Примеры данных
// Пример 1
// abacaba
// abc
// 4
// caba
// Пример 2
// abacaba
// abc
// 3
// cab

package main

import (
	"fmt"
)

func main() {
	var sequence, charset string
	var k int

	fmt.Scan(&sequence)
	fmt.Scan(&charset)
	fmt.Scan(&k)

	required := make(map[byte]int)
	for i := 0; i < len(charset); i++ {
		required[charset[i]]++
	}

	requiredUnique := len(required)

	window := make(map[byte]int)
	left := 0
	foundUnique := 0

	bestLeft, bestRight := -1, -1

	for right := 0; right < len(sequence); right++ {
		char := sequence[right]

		if _, ok := required[char]; ok {
			window[char]++
			if window[char] == required[char] {
				foundUnique++
			}
		}

		for foundUnique == requiredUnique {
			if right-left+1 <= k {
				if bestLeft == -1 || left > bestLeft || (left == bestLeft && right-left > bestRight-bestLeft) {
					bestLeft = left
					bestRight = right
				}
			}

			leftChar := sequence[left]
			if _, ok := required[leftChar]; ok {
				if window[leftChar] == required[leftChar] {
					foundUnique--
				}
				window[leftChar]--
			}
			left++
		}
	}

	if bestLeft == -1 {
		fmt.Println(bestLeft)
	} else {
		fmt.Println(sequence[bestLeft : bestLeft+k])
	}
}
