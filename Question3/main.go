package main

import "fmt"

type SeatMap map[int]string

func main() {
	seatTable := SeatMap{
		1: "Abbot",
		2: "Doris",
		3: "Emerson",
		4: "Green",
		5: "Jeames",
	}

	result := swapSeats(seatTable)

	printMap(result)
}

func swapSeats(seats SeatMap) SeatMap {
	result := make(SeatMap)

	for i := 1; i <= len(seats); i += 2 {
		// Check if there are at least two more students
		if i+1 <= len(seats) {
			result[i] = seats[i+1]
			result[i+1] = seats[i]
		} else {
			result[i] = seats[i]
		}
	}

	return result
}

// printMap prints the SeatMap
func printMap(seats SeatMap) {
	fmt.Println("| id | student |")

	for id, student := range seats {
		fmt.Printf("| %2d | %-7s |\n", id, student)
	}
}
