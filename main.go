package main

import (
	"fmt"

	"github.com/muhammadisa/go-transformer/example/model"
	pb "github.com/muhammadisa/go-transformer/example/protobuf"
	"github.com/muhammadisa/go-transformer/transformer"
)

func sampleStructToProtoc() *pb.Todo {
	todoStruct := model.Todo{
		ID:            "5b9e1416-1f06-4a61-a30a-0dcff164639b",
		Name:          "Isa",
		Completed:     false,
		NumberCode:    129520,
		NumberProduct: 25983578228,
	}
	codes := []model.Code{{1}, {4}, {9}}

	todoProtocBlank := &pb.Todo{}

	transformer.Struct(&todoStruct).Protoc(todoProtocBlank)
	fmt.Println("todoProtocBlank", todoProtocBlank)
	fmt.Println("todoProtocBlank.Codes", todoProtocBlank.Codes)

	for _, c := range codes {
		todoCodeProtocBlank := &pb.Code{}
		transformer.Struct(&c).Protoc(todoCodeProtocBlank)
		todoProtocBlank.Codes = append(todoProtocBlank.Codes, todoCodeProtocBlank)
	}
	fmt.Println("todoProtocBlank", todoProtocBlank)
	fmt.Println("todoProtocBlank.Codes", todoProtocBlank.Codes)

	return todoProtocBlank
}

func sampleProtocToStruct(todoProtoc *pb.Todo) {
	todo := model.Todo{}

	transformer.Protoc(todoProtoc).Struct(&todo)
	fmt.Println("todoStruct", todo)
	fmt.Println("todoStruct.Codes", todo.Codes)

	for _, i := range todoProtoc.Codes {
		var code model.Code
		transformer.Protoc(i).Struct(&code)
		todo.Codes = append(todo.Codes, code)
	}
	fmt.Println("todoStruct", todo)
	fmt.Println("todoStruct.Codes", todo.Codes)
}

func main() {
	todoProtoc := sampleStructToProtoc()
	sampleProtocToStruct(todoProtoc)
}
