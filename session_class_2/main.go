package main

import "fmt"

func main() {
	var kalimat = "selamat malam"

	for i := 0; i < len(kalimat); i++ {
		fmt.Println(string(kalimat[i]))
	}

	jumlahHuruf := map[rune]int{}

	for _, huruf := range kalimat {
		jumlahHuruf[huruf]++
	}

	fmt.Print("map[")
	for key, value := range jumlahHuruf {
		if key == ' ' {
			fmt.Printf("' ': %d ", value)
		} else {
			fmt.Printf("%c:%d ", key, value)
		}
	}
	fmt.Println("]")
}
