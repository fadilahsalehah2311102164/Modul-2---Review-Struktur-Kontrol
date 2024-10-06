package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	correctOrder := []string{"merah", "kuning", "hijau", "ungu"}
	reader := bufio.NewReader(os.Stdin)
	success := true
	for i := 1; i <= 5; i++ {
		fmt.Printf("Percobaan %d: ", i)
		input, _ := reader.ReadString('\n')
		input = strings.TrimSpace(input)
		colors := strings.Split(input, " ")
		for j := 0; j < 4; j++ {
			if colors[j] != correctOrder[j] {
				success = false
				break
			}
		}
		if !success {
			break
		}
	}
	if success {
		fmt.Println("BERHASIL : true")
	} else {
		fmt.Println("BERHASIL : false")
	}
}
