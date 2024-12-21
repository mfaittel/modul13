package main

import "fmt"

func insertionSort(arr []int) {
	for i := 1; i < len(arr); i++ {
		key := arr[i]
		j := i - 1

		for j >= 0 && arr[j] > key {
			arr[j+1] = arr[j]
			j--
		}
		arr[j+1] = key
	}
}

func cekJarakTetap(arr []int) (bool, int) {
	if len(arr) < 2 {
		return true, 0
	}

	jarak := arr[1] - arr[0]

	for i := 2; i < len(arr); i++ {
		if arr[i]-arr[i-1] != jarak {
			return false, 0
		}
	}
	return true, jarak
}

func main() {
	var data []int
	var input int

	fmt.Println("Masukkan bilangan ( - = Selesai ):")

	for {
		fmt.Scan(&input)
		if input < 0 {
			break
		}
		data = append(data, input)
	}

	if len(data) == 0 {
		fmt.Println("Tidak ada data.")
		return
	}

	insertionSort(data)

	fmt.Println("Data setelah diurutkan:")
	for _, v := range data {
		fmt.Printf("%d ", v)
	}
	fmt.Println()

	isJarakTetap, jarak := cekJarakTetap(data)
	if isJarakTetap {
		fmt.Printf("Data berjarak %d\n", jarak)
	} else {
		fmt.Println("Data berjarak tidak tetap")
	}
}
