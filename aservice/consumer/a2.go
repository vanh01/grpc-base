package consumer

import "github.com/vanh01/grpc-base/domain"

type A2[T domain.A2Request, K domain.A2Result] struct {
	Field1 T
	Field2 K
}

func (a A2[T, K]) Consume(request T) (K, error) {
	r := domain.A2Request(request)

	// TO DO

	return K(domain.A2Result{
		Id:   r.Id,
		Name: "A1 reply: " + r.Name,
	}), nil
}
