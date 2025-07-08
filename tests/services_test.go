package tests

import (
	"testing"

	"perfect-numbers-api/internal/services"

	"github.com/stretchr/testify/assert"
)

func TestPerfectNumberService_IsPerfectNumber(t *testing.T) {
	service := services.NewPerfectNumberService()

	tests := []struct {
		name     string
		number   int
		expected bool
	}{
		{"Número 6 é perfeito", 6, true},
		{"Número 28 é perfeito", 28, true},
		{"Número 496 é perfeito", 496, true},
		{"Número 8128 é perfeito", 8128, true},
		{"Número 1 não é perfeito", 1, false},
		{"Número 2 não é perfeito", 2, false},
		{"Número 12 não é perfeito", 12, false},
		{"Número 100 não é perfeito", 100, false},
		{"Número negativo não é perfeito", -1, false},
		{"Número zero não é perfeito", 0, false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := service.IsPerfectNumber(tt.number)
			assert.Equal(t, tt.expected, result, "IsPerfectNumber(%d) = %v, esperado %v", tt.number, result, tt.expected)
		})
	}
}

func TestPerfectNumberService_FindPerfectNumbers(t *testing.T) {
	service := services.NewPerfectNumberService()

	tests := []struct {
		name     string
		start    int
		end      int
		expected []int
	}{
		{
			name:     "Range 1-10 deve encontrar apenas 6",
			start:    1,
			end:      10,
			expected: []int{6},
		},
		{
			name:     "Range 1-100 deve encontrar 6 e 28",
			start:    1,
			end:      100,
			expected: []int{6, 28},
		},
		{
			name:     "Range 1-10000 deve encontrar 6, 28, 496, 8128",
			start:    1,
			end:      10000,
			expected: []int{6, 28, 496, 8128},
		},
		{
			name:     "Range 7-27 não deve encontrar números perfeitos",
			start:    7,
			end:      27,
			expected: nil,
		},
		{
			name:     "Range 28-28 deve encontrar apenas 28",
			start:    28,
			end:      28,
			expected: []int{28},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			response := service.FindPerfectNumbers(tt.start, tt.end)
			assert.Equal(t, tt.expected, response.PerfectNumbers, "FindPerfectNumbers(%d, %d) = %v, esperado %v", tt.start, tt.end, response.PerfectNumbers, tt.expected)
			assert.Equal(t, len(tt.expected), response.Count, "Count deve ser igual ao número de elementos encontrados")
			assert.NotEmpty(t, response.ProcessingTime, "ProcessingTime não deve estar vazio")
			assert.NotZero(t, response.Timestamp, "Timestamp não deve estar vazio")
		})
	}
}

func BenchmarkIsPerfectNumber(b *testing.B) {
	service := services.NewPerfectNumberService()

	b.Run("Número pequeno (6)", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			service.IsPerfectNumber(6)
		}
	})

	b.Run("Número médio (8128)", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			service.IsPerfectNumber(8128)
		}
	})

	b.Run("Número grande (33550336)", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			service.IsPerfectNumber(33550336)
		}
	})
}
