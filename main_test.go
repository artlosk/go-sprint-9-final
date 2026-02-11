package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGenerateRandomElementsSizeZero(t *testing.T) {
	result := generateRandomElements(0)
	assert.Empty(t, result, "для размера 0 должен возвращаться пустой слайс")
}

func TestGenerateRandomElementsSizePositive(t *testing.T) {
	size := 10
	result := generateRandomElements(size)
	assert.Len(t, result, size, "длина слайса должна совпадать с запрошенным размером")
}

func TestGenerateRandomElementsNegativeSize(t *testing.T) {
	result := generateRandomElements(-5)
	assert.Empty(t, result, "для отрицательного размера должен возвращаться пустой слайс")
}

func TestGenerateRandomElementsSingleElement(t *testing.T) {
	result := generateRandomElements(1)
	assert.Len(t, result, 1, "для размера 1 должен возвращаться слайс из одного элемента")
}

func TestMaximum(t *testing.T) {
	tests := []struct {
		name    string
		input   []int
		want    int
		wantErr bool
	}{
		{
			name:    "пустой слайс",
			input:   []int{},
			wantErr: true,
		},
		{
			name:    "один элемент",
			input:   []int{42},
			want:    42,
			wantErr: false,
		},
		{
			name:    "все положительные числа",
			input:   []int{1, 3, 2, 5, 4},
			want:    5,
			wantErr: false,
		},
		{
			name:    "все отрицательные числа",
			input:   []int{-10, -3, -7, -1},
			want:    -1,
			wantErr: false,
		},
		{
			name:    "положительные и отрицательные числа",
			input:   []int{-2, 0, 5, -1, 3},
			want:    5,
			wantErr: false,
		},
		{
			name:    "повторяющиеся числа",
			input:   []int{2, 7, 7, 3},
			want:    7,
			wantErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := maximum(tt.input)

			if tt.wantErr {
				assert.Error(t, err, "Ожидалась ошибка: слайс пуст")
			} else {
				assert.NoError(t, err, "Не ожидалось ошибки")
				assert.Equal(t, tt.want, got, "Максимальное значение не совпадает")
			}
		})
	}
}

func TestMaxChunksNormalCase(t *testing.T) {
	data := []int{1, 5, 3, 9, 2, 8, 4, 6, 7}
	result, err := maxChunks(data)
	assert.NoError(t, err)
	assert.Equal(t, 9, result, "многопоточный поиск должен находить корректный максимум")
}

func TestMaxChunksEmptySlice(t *testing.T) {
	result, err := maxChunks([]int{})
	assert.Error(t, err, "для пустого слайса должна возвращаться ошибка")
	_ = result
}

func TestMaxChunksSingleElement(t *testing.T) {
	result, err := maxChunks([]int{42})
	assert.NoError(t, err)
	assert.Equal(t, 42, result, "для одного элемента должен возвращаться этот элемент")
}

func TestMaxChunksLessThanChunks(t *testing.T) {
	data := []int{1, 5, 3}
	result, err := maxChunks(data)
	assert.NoError(t, err)
	assert.Equal(t, 5, result, "должен корректно обрабатывать слайс меньше CHUNKS")
}

func TestMaxChunksExactlyChunks(t *testing.T) {
	data := []int{1, 2, 3, 4, 5, 6, 7, 8}
	result, err := maxChunks(data)
	assert.NoError(t, err)
	assert.Equal(t, 8, result, "должен корректно обрабатывать слайс размера CHUNKS")
}

func TestMaxChunksMoreThanChunks(t *testing.T) {
	data := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16}
	result, err := maxChunks(data)
	assert.NoError(t, err)
	assert.Equal(t, 16, result, "должен корректно обрабатывать слайс больше CHUNKS")
}

func TestMaxChunksAllSameElements(t *testing.T) {
	data := []int{7, 7, 7, 7, 7, 7, 7, 7, 7}
	result, err := maxChunks(data)
	assert.NoError(t, err)
	assert.Equal(t, 7, result, "должен корректно обрабатывать одинаковые элементы")
}
