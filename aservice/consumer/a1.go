package consumer

import (
	"github.com/vanh01/grpc-base/domain"
)

type A1[T domain.A1Request, K domain.A1Result] struct {
	Field1 T
	Field2 K
}

func (a A1[T, K]) Consume(request domain.A1Request) (domain.A1Result, error) {
	r := request

	// TO DO

	return domain.A1Result{
		Id:   r.Id,
		Name: "A1 reply: " + r.Name,
	}, nil
}
