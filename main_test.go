package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGenerateRandomElementsSizeZero(t *testing.T) {
	result := generateRandomElements(0)
	if len(result) != 0 {
		t.Errorf("GenerateRandomElements: Ошибка, размер слайса 0, полученный %d", len(result))
	}
}

func TestGenerateRandomElementsSizePositive(t *testing.T) {
	size := 10
	if len(generateRandomElements(size)) != size {
		t.Errorf("GenerateRandomElements: Ошибка, размер слайса %d, полученный %d", size, len(generateRandomElements(size)))
	}
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
				assert.Error(t, err, "Ожидалась ошибка: слайс пуст") // сообщение на русском
			} else {
				assert.NoError(t, err, "Не ожидалось ошибки")                       // сообщение на русском
				assert.Equal(t, tt.want, got, "Максимальное значение не совпадает") // тоже на русском
			}
		})
	}
}
