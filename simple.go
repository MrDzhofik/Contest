// Вы уже наверняка устали читать легенды, придуманные только для того, чтобы удлинить время чтения и
// отдалить вас от настоящего условия задачи.
// В этих легендах какие-то странные люди делают какие-то странные действия, редко относящиеся к реальной жизни.
//  Ну вот какая ещё вечная зима, если грядёт глобальное потепление? Неужели действительно на собеседованиях дают такие задачи?
// Или как можно забыть пароль от рабочего ноутбука?
// Хотя это действительно однажды случилось с автором задачи, остальное совсем не соответствует истине.

// Да и вообще, пора поднять бунт против бесполезных легенд!
// Пусть все условия станут максимально простыми, как, например, числа, имеющие ровно два делителя: единицу и само себя.
//  Все остальные числа, кроме простых и 11, называются составными.
// Вам нужно решить очень понятную задачу: посчитать количество составных чисел от ll до rr,
//  количество делителей которых при этом является простым числом.

// Формат входных данных

// Входная строка содержит два целых числа ll и rr (1≤l≤r≤1014)(1≤l≤r≤1014).

// Формат выходных данных

// Выведите количество таких чисел на отрезке от ll до rr, включительно.
// Примеры данных
// Пример 1
// 1 9
// 2
// Пример 2
// 3 6
// 1
// Пример 3
// 6 9
// 1

package main

import (
	"fmt"
	"math"
)

func isPrime(num int) bool {
	if num <= 1 {
		return false
	}
	if num <= 3 {
		return true
	}
	for i := 2; i*i <= num; i += 1 {
		if num%i == 0 {
			return false
		}
	}
	return true
}

func countDivisors(n int) int {
	count := 0
	for i := 1; i <= int(math.Sqrt(float64(n))); i++ {
		if n%i == 0 {
			if i*i == n {
				count++
			} else {
				count += 2
			}
		}
	}
	return count
}

func countPrimeDivisors(l, r int) int {
	res := 0
	for i := l; i < r; i++ {
		count := countDivisors(i)
		if isPrime(count) {
			res++
		}
	}
	return res
}

func main() {
	var l, r int
	fmt.Scan(&l, &r)

	fmt.Println(countPrimeDivisors(l, r))
}
