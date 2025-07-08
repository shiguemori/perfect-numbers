package models

import (
	"errors"
)

type PerfectNumberRequest struct {
	Start int `json:"start" binding:"required,min=1"`
	End   int `json:"end" binding:"required,min=1"`
}

func (r *PerfectNumberRequest) Validate() error {
	if r.Start <= 0 {
		return errors.New("start deve ser um número positivo")
	}
	if r.End <= 0 {
		return errors.New("end deve ser um número positivo")
	}
	if r.Start > r.End {
		return errors.New("start deve ser menor ou igual a end")
	}
	if r.End > 1000000 {
		return errors.New("end não pode ser maior que 1.000.000 para evitar timeout")
	}
	return nil
}
