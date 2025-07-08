package services

import (
	"fmt"
	"time"

	"perfect-numbers-api/internal/models"
)

type PerfectNumberService interface {
	FindPerfectNumbers(start, end int) *models.PerfectNumberResponse
	IsPerfectNumber(n int) bool
}
type perfectNumberService struct{}

func NewPerfectNumberService() PerfectNumberService {
	return &perfectNumberService{}
}

func (s *perfectNumberService) IsPerfectNumber(n int) bool {
	if n <= 1 {
		return false
	}

	sum := 0
	for i := 1; i*i <= n; i++ {
		if n%i == 0 {
			sum += i
			if i != 1 && i*i != n {
				sum += n / i
			}
		}
	}

	return sum == n
}

func (s *perfectNumberService) FindPerfectNumbers(start, end int) *models.PerfectNumberResponse {
	startTime := time.Now()
	var perfectNumbers []int

	for i := start; i <= end; i++ {
		if s.IsPerfectNumber(i) {
			perfectNumbers = append(perfectNumbers, i)
		}
	}

	processingTime := time.Since(startTime)

	return &models.PerfectNumberResponse{
		PerfectNumbers: perfectNumbers,
		Count:          len(perfectNumbers),
		Range:          fmt.Sprintf("%d-%d", start, end),
		ProcessingTime: processingTime.String(),
		Timestamp:      time.Now(),
	}
}
