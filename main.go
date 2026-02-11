package main

import (
	"errors"
	"fmt"
	"math/rand"
	"sync"
	"time"
)

const (
	SIZE   = 100_000_000
	CHUNKS = 8
)

// generateRandomElements generates random elements.
func generateRandomElements(size int) []int {
	if size <= 0 {
		return []int{}
	}
	elements := make([]int, size)
	for i := range elements {
		elements[i] = rand.Int()
	}
	return elements
}

// maximum returns the maximum number of elements.
func maximum(data []int) (max int, err error) {
	if len(data) == 0 {
		return 0, errors.New("слайс пуст")
	}

	max = data[0]
	for _, num := range data[1:] {
		if num > max {
			max = num
		}
	}
	return max, nil
}

// maxChunks returns the maximum number of elements in a chunks.
func maxChunks(data []int) (int, error) {
	if len(data) == 0 {
		return 0, errors.New("слайс пуст")
	}

	if len(data) < CHUNKS {
		return maximum(data)
	}

	chunkSize := len(data) / CHUNKS
	maxSlice := make([]int, CHUNKS)

	var wg sync.WaitGroup
	wg.Add(CHUNKS)
	for i := 0; i < CHUNKS; i++ {
		start := i * chunkSize
		end := start + chunkSize
		if i == CHUNKS-1 {
			end = len(data)
		}
		go func(index int, slice []int) {
			defer wg.Done()
			maxVal, _ := maximum(slice)
			maxSlice[index] = maxVal
		}(i, data[start:end])
	}
	wg.Wait()
	maxVal, err := maximum(maxSlice)
	if err != nil {
		return 0, fmt.Errorf("ошибка: %w", err)
	}
	return maxVal, nil
}

func main() {
	fmt.Printf("Генерируем %d целых чисел\n", SIZE)
	data := generateRandomElements(SIZE)

	fmt.Println("Ищем максимальное значение в один поток")
	start := time.Now()
	maxVal, err := maximum(data)
	if err != nil {
		fmt.Println("Ошибка:", err)
		return
	}
	elapsed := time.Since(start).Milliseconds()

	fmt.Printf("Максимальное значение элемента: %d\nВремя поиска: %d ms\n", maxVal, elapsed)

	fmt.Printf("Ищем максимальное значение в %d потоков\n", CHUNKS)
	start = time.Now()
	maxVal, err = maxChunks(data)
	if err != nil {
		fmt.Println("Ошибка:", err)
		return
	}
	elapsed = time.Since(start).Milliseconds()
	fmt.Printf("Максимальное значение элемента: %d\nВремя поиска: %d ms\n", maxVal, elapsed)
}
