package main

import (
	"fmt"
	"os"
	"sync"
)

func writeFile(filename string, numbers []int, wg *sync.WaitGroup) {
	defer wg.Done()

	file, err := os.Create(filename)
	if err != nil {
		fmt.Println("create error:", err)
		return
	}
	defer file.Close()

	for _, num := range numbers {
		fmt.Fprintf(file, "%d ", num)
	}
}

func main() {
	var wg sync.WaitGroup


	text1 := make([]int, 0)

	
	for i := 0; i <= 100; i += 5 {
		text1 = append(text1, i)
	}


	text2 := make([]int, 0)
	for i := 0; i <= 100; i += 3 {
		text2 = append(text2, i)
	}

	
	wg.Add(2)
	go writeFile("file1.txt", text1, &wg)


	go writeFile("file2.txt", text2, &wg)

	
	wg.Wait()

	
	text3 := make([]int, 0)
	for i := 1; i <= 100; i++ {
		if i%15 == 0 {
			text3 = append(text3, i)
		}
	}

	wg.Add(1)
	go writeFile("file3.txt", text3, &wg)

	
	wg.Wait()
	fmt.Println("write all.")
}
