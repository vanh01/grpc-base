package consumer

import (
	"fmt"

	grpcbase "github.com/vanh01/grpc-base"
	"github.com/vanh01/grpc-base/domain"
)

func Init() {
	fmt.Println("Init register type for comsumers")
	grpcbase.RegisterGenericType(A1[domain.A1Request, domain.A1Result]{})
	grpcbase.RegisterGenericType(A2[domain.A2Request, domain.A2Result]{})
	// fmt.Println(grpcbase.TypeRegistry)
}
