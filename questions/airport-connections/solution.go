package main

import (
	"fmt"
)

func solution(airports []string, routes [][]string, startingAirport string) int {
	airportsSet := make(map[string]int)
	for i, a := range airports {
		airportsSet[a] = i
	}
	routesBool := make([][]bool, len(airports))
	for i := range airports {
		routesBool[i] = make([]bool, len(airports))
	}
	for _, r := range routes {
		a1, a2 := airportsSet[r[0]], airportsSet[r[1]]
		routesBool[a1][a2] = true
	}
	for i, r := range routesBool {
		for j, a := range r {
			
		}
	}
}

func main() {
	fmt.Println(solution([]string{}, [][]string{}, ""))
}
