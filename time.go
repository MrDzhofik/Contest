// Миша участвует в специальном CTF-соревновании в составе команды, который проходит в формате 2424-часового хакатона.
// Хакатон длится целые сутки после его начала. Если хакатон начинается в 18:00:0018:00:00 одного дня,
// то последняя секунда, в которую можно сдать решение, будет 17:59:5917:59:59 следующего дня.

// Цель в CTF-соревновании — взломать наибольшее количество серверов с наименьшим штрафом.
// Каждый сервер имеет свой идентификатор — одну заглавную латинскую букву.
// Команды взламывают сервера независимо, и взломы одной команды никак не влияют на другие команды.

// Если команда взламывает сервер, её счет увеличивается на один — а к штрафному времени прибавляется время в минутах,
//  округленное вниз, которое прошло от начала соревнования до времени взлома.
// Если перед удачной попыткой взлома одного сервера команда совершает одну или несколько неудачных попыток взлома этого же сервера
// — то к штрафному времени прибавляется по двадцать минут за каждую такую неудачную попытку.
// При этом, если сервер в итоге не был взломан, штрафное время не начисляется.
// В ходе соревнования команды могут делать «PINGPING» запросы к серверу, которые никак не учитываются при подсчете результатов,
//  и за них не предусмотрено начисление штрафного времени.

// Побеждает та команда, которая взламывает наибольшее количество серверов, а если таких несколько,
// то команда с наименьшим штрафным временем. В начальный момент времени команды не взломали ни одного сервера
//  и имеют штрафное время, равное нулю.

// Напишите программу, которая выводит результаты хакатона.

// Формат входных данных

// В первой строке дано время начала хакатона в формате hh:mm:sshh:mm:ss, где даны соответственно часы, минуты и секунды соответственно
// (0≤hh≤23,0≤mm,ss≤59)(0≤hh≤23,0≤mm,ss≤59).

// Во второй строке дано целое число nn (1≤n≤1000)(1≤n≤1000) — количество запросов к серверам за весь хакатон.

// Далее следуют nn строк с описаниями. В начале каждой строки записано название команды в двойных кавычках.
//  Название может состоять из строчных и заглавных латинских букв, пробелов и цифр от 11 до 99.
// Название команды не пустое и не превосходит 255255 символов. После через пробел дано время запроса
// в аналогичном времени начала хакатона формате.

// Далее через пробел идет одна заглавная латинская буква — идентификатор сервера.
//  Далее указан результат запроса команды к серверу: ACCESSEDACCESSED — сервер взломан; DENIEDDENIED — неудачная попытка взлома;
//  FORBIDENFORBIDEN — неудачная попытка взлома; PONGPONG — ответ на запрос «PINGPING».

// Формат выходных данных

// Вывод должен содержать итоговую таблицу результатов — по строке на каждую команду.
// Строки должны быть отсортированы по результату (количество взломанных серверов и штрафное время),
//  а если у нескольких команд результаты равны, то порядок команд определяется лексикографически меньшим названием команды.

// Каждая строка должна начинаться с места команды в итоговом зачете. Место команды — это k+1k+1, где kk — число команд,
// имеющих строго лучший результат. Далее через пробел идет название команды в двойных кавычках,
//  а за ним через пробел два числа — количество взломанных серверов и штрафное время.
// Примеры данных
// Пример 1
// 00:00:00
// 5
// "VK" 00:10:21 A FORBIDEN
// "T" 00:00:23 A DENIED
// "T" 00:20:23 A ACCESSED
// "VK" 00:30:23 A ACCESSED
// "YA" 00:40:23 B ACCESSED
// 1 "T" 1 40
// 1 "YA" 1 40
// 3 "VK" 1 50
// Пример 2
// 01:00:00
// 3
// "Team1" 01:10:00 A FORBIDEN
// "Team1" 01:20:00 A ACCESSED
// "Team2" 01:40:00 B ACCESSED
// 1 "Team1" 1 40
// 1 "Team2" 1 40
// Пример 3
// 23:00:00
// 2
// "Team1" 23:59:59 A PONG
// "Team1" 00:00:00 A ACCESSED
// 1 "Team1" 1 60

package main

import (
	"fmt"
	"math"
	"sort"
	"strconv"
	"strings"
)

type Ans struct {
	teamName string
	servers  int
	penalty  int
}

func timeToSeconds(t string) int {
	parts := strings.Split(t, ":")
	h, _ := strconv.Atoi(parts[0])
	m, _ := strconv.Atoi(parts[1])
	s, _ := strconv.Atoi(parts[2])
	return h*3600 + m*60 + s
}

func secondsToMinutes(seconds int) int {
	return seconds / 60
}

func main() {
	var startTime string
	var n int

	fmt.Scan(&startTime)
	startSeconds := timeToSeconds(startTime)

	fmt.Scan(&n)

	results := make(map[string]map[rune]map[string]int)

	attempts := make(map[string]map[rune]int)

	for i := 0; i < n; i++ {
		var teamName, timestamp, serverID, result string
		fmt.Scan(&teamName, &timestamp, &serverID, &result)
		teamName = teamName[1 : len(teamName)-1]
		timeSeconds := timeToSeconds(timestamp)
		elapsed := timeSeconds - startSeconds
		if _, ok := results[teamName]; !ok {
			results[teamName] = make(map[rune]map[string]int)
			attempts[teamName] = make(map[rune]int)
		}
		if _, ok := results[teamName][rune(serverID[0])]; !ok {
			results[teamName][rune(serverID[0])] = map[string]int{
				"firstAccessTime": math.MaxInt64,
				"successful":      0,
				"totalPenalty":    0,
			}
		}

		if result == "ACCESSED" {
			if elapsed < results[teamName][rune(serverID[0])]["firstAccessTime"] {
				results[teamName][rune(serverID[0])]["firstAccessTime"] = elapsed
				results[teamName][rune(serverID[0])]["successful"] = 1
				results[teamName][rune(serverID[0])]["totalPenalty"] = secondsToMinutes(elapsed) + 20*attempts[teamName][rune(serverID[0])]
			}
		} else if result == "DENIED" || result == "FORBIDEN" {
			attempts[teamName][rune(serverID[0])]++
		}
	}

	var finalResults []Ans

	for teamName, serversMap := range results {
		serverCount := 0
		totalPenalty := 0
		for _, data := range serversMap {
			if data["successful"] > 0 {
				serverCount++
				totalPenalty += data["totalPenalty"]
			}
		}
		finalResults = append(finalResults, struct {
			teamName string
			servers  int
			penalty  int
		}{teamName, serverCount, totalPenalty})
	}

	sort.Slice(finalResults, func(i, j int) bool {
		if finalResults[i].servers != finalResults[j].servers {
			return finalResults[i].servers > finalResults[j].servers
		}
		if finalResults[i].penalty != finalResults[j].penalty {
			return finalResults[i].penalty < finalResults[j].penalty
		}
		return finalResults[i].teamName < finalResults[j].teamName
	})

	for rank, result := range finalResults {
		fmt.Printf("%d \"%s\" %d %d\n", rank+1, result.teamName, result.servers, result.penalty)
	}
}
