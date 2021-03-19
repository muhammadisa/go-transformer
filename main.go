package main

import (
	"fmt"
	"github.com/muhammadisa/go-transformer/example/model"
	pb "github.com/muhammadisa/go-transformer/example/protobuf"
	"github.com/muhammadisa/go-transformer/transformer"
	uuid "github.com/satori/go.uuid"
)

func main() {
	person := model.Todo{
		ID:            uuid.NewV4().String(),
		Name:          "Isa",
		Completed:     false,
		NumberCode:    129520,
		NumberProduct: 25983578228,
	}
	codes := []model.Code{{1}, {4}, {9}}

	pPerson := pb.Todo{}

	for _, i := range codes {
		var localPbCode pb.Code
		transformer.Transformed{From: &i}.TransformerToProto(&localPbCode)
		pPerson.Codes = append(pPerson.Codes, &localPbCode)
	}
	transformer.Transformed{From: &person}.TransformerToProto(&pPerson)
	fmt.Println(&pPerson)
}
