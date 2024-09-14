// Любая крупная система обработки больших данных оперирует терабайтами информации.

// Рассмотрим подобную примитивную систему. В нашей системе данные обрабатываются процессами,
// которые выстраиваются в зависимости друг от друга. В системе находится nn процессов, пронумерованных от 11 до nn.
//  Процесс с номером ii сначала ждет завершения всех процессов, от которых он зависит, исполняется за titi​ секунд,
// после чего завершается. Гарантируется, что все процессы завершатся за конечное время, т. е. нет циклических зависимостей
//  между процессами.

// Вам предстоит определить, за какое минимальное время смогут завершиться все процессы.
// Минимальным временем считается то, которое достигается, когда планировщик процессов действует оптимально и
// имеет в распоряжении бесконечное количество вычислительных узлов, на которых, тем не менее, каждый процесс
// завершается строго за указанное ему время.

// Формат входных данных

// В первой строке дано число nn (1≤n≤105)(1≤n≤105) — количество процессов.

// Далее дано nn строк. В ii-й строке первым числом идёт titi (0≤ti≤1012)(0≤ti​≤1012) — время исполнения ii-го процесса в секундах.
// Далее до конца строки идут номера процессов, от которых зависит процесс ii.

// Формат выходных данных

// В единственной строке выведите одно число — минимальное время в секундах, за которое могут исполниться все процессы.
// Примеры данных
// Пример 1
// 5
// 10 2 3 5
// 5 4
// 0
// 4
// 15 3
// 25
// Пример 2
// 6
// 2 2
// 2 3
// 15 4
// 1 5
// 2 6
// 0
// 22

package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	scanner.Scan()
	n, _ := strconv.Atoi(scanner.Text())

	processingTime := make([]int, n+1)
	graph := make([][]int, n+1)
	inDegree := make([]int, n+1)

	for i := 1; i <= n; i++ {
		scanner.Scan()
		line := scanner.Text()
		parts := strings.Split(line, " ")

		t, _ := strconv.Atoi(parts[0])
		processingTime[i] = t

		for _, dep := range parts[1:] {
			if dep != "" {
				d, _ := strconv.Atoi(dep)
				graph[d] = append(graph[d], i)
				inDegree[i]++
			}
		}
	}

	queue := []int{}
	for i := 1; i <= n; i++ {
		if inDegree[i] == 0 {
			queue = append(queue, i)
		}
	}

	finishTime := make([]int, n+1)
	for i := 1; i <= n; i++ {
		finishTime[i] = processingTime[i]
	}

	for len(queue) > 0 {
		process := queue[0]
		queue = queue[1:]

		for _, neighbor := range graph[process] {
			finishTime[neighbor] = max(finishTime[neighbor], finishTime[process]+processingTime[neighbor])
			inDegree[neighbor]--
			if inDegree[neighbor] == 0 {
				queue = append(queue, neighbor)
			}
		}
	}

	minimumCompletionTime := 0
	for i := 1; i <= n; i++ {
		minimumCompletionTime = max(minimumCompletionTime, finishTime[i])
	}

	fmt.Println(minimumCompletionTime)
}
