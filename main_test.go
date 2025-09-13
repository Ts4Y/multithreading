package main

// Пишите тесты в этом файле

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestGenerateRandomElements(t *testing.T) {
	// Тест на генерацию массива заданного размера
	size := 1000
	assert.Len(t, generateRandomElements(size), size)
	assert.Len(t, generateRandomElements(size), size)
	require.Len(t, generateRandomElements(0), 0)
	assert.Len(t, generateRandomElements(-10), 0)
	size = 1_000_000
	assert.Len(t, generateRandomElements(size), size)

}

func TestMaximum(t *testing.T) {
	arr := []int{1, 3, 2, 5, 4}
	max := maximum(arr)
	assert.Equal(t, 5, max, "максимум должен быть 5")

	emptyArr := []int{}
	max = maximum(emptyArr)
	require.Equal(t, 0, max)

	singleElementArr := []int{42}
	max = maximum(singleElementArr)
	assert.Equal(t, 42, max)

	negativeArr := []int{-1, -3, -2, -5, -4}
	max = maximum(negativeArr)
	assert.Equal(t, -1, max)

	largeArr := make([]int, 1000000)
	for i := 0; i < 1000000; i++ {
		largeArr[i] = i
	}
	assert.Equal(t, 999999, maximum(largeArr))
}
func TestMaxChunks(t *testing.T) {
	arr := []int{1, 3, 2, 5, 4, 8, 7, 6}
	max := maxChunks(arr)
	assert.Equal(t, 8, max)

	emptyArr := []int{}
	max = maxChunks(emptyArr)
	require.Equal(t, 0, max)

	singleElementArr := []int{42}
	max = maxChunks(singleElementArr)
	assert.Equal(t, 42, max)

	negativeArr := []int{-1, -3, -2, -5, -4}
	max = maxChunks(negativeArr)
	assert.Equal(t, -1, max)

	largeArr := make([]int, 1000000)
	for i := 0; i < 1000000; i++ {
		largeArr[i] = i
	}

	max = maxChunks(largeArr)
	assert.Equal(t, 999999, max)
}
