package transformer

import (
	"github.com/muhammadisa/go-transformer/example/model"
	pb "github.com/muhammadisa/go-transformer/example/protobuf"
	uuid "github.com/satori/go.uuid"
	"github.com/stretchr/testify/suite"
	"testing"
)

type TransformerTestSuite struct {
	suite.Suite
}

func TestTransformerTestSuite(t *testing.T) {
	suite.Run(t, new(TransformerTestSuite))
}

func (suite *TransformerTestSuite) TestAutoConf() {
	suite.Run("TestAutoConf", func() {
		todoCopyExample := map[string][]interface{}{
			"Id":            {"string", "5b9e1416-1f06-4a61-a30a-0dcff164639b"},
			"Name":          {"string", "Mark Zuck"},
			"Completed":     {"bool", false},
			"NumberCode":    {"int32", int32(129520)},
			"NumberProduct": {"uint64", uint64(25983578228)},
		}

		pPerson := pb.Todo{}
		autoConf(&pPerson, todoCopyExample)

		suite.T().Log(&pPerson)
	})
}

func (suite *TransformerTestSuite) TestAutoConfBool() {
	type DummyStruct struct {
		Content bool
	}

	suite.Run("AutoConf for bool", func() {
		const data = false
		todoCopyExample := map[string][]interface{}{
			"Content": {"bool", data},
		}
		t := DummyStruct{}
		autoConf(&t, todoCopyExample)
		suite.Assert().Equal(t.Content, data)
	})
}

func (suite *TransformerTestSuite) TestAutoConfString() {
	type DummyStruct struct {
		Content string
	}

	suite.Run("AutoConf for string", func() {
		const data = "hello"
		todoCopyExample := map[string][]interface{}{
			"Content": {"string", data},
		}
		t := DummyStruct{}
		autoConf(&t, todoCopyExample)
		suite.Assert().Equal(t.Content, data)
	})
}

func (suite *TransformerTestSuite) TestAutoConfByte() {
	type DummyStruct struct {
		Content []byte
	}

	suite.Run("AutoConf for byte", func() {
		data := []byte{1, 2, 3, 4, 5}
		todoCopyExample := map[string][]interface{}{
			"Content": {"byte", data},
		}
		t := DummyStruct{}
		autoConf(&t, todoCopyExample)
		suite.Assert().Equal(t.Content, data)
	})
}

func (suite *TransformerTestSuite) TestAutoConfFloat32() {
	type DummyStruct struct {
		Content float32
	}

	suite.Run("AutoConf for float32", func() {
		const data = float32(10)
		todoCopyExample := map[string][]interface{}{
			"Content": {"float32", data},
		}
		t := DummyStruct{}
		autoConf(&t, todoCopyExample)
		suite.Assert().Equal(t.Content, data)
	})
}

func (suite *TransformerTestSuite) TestAutoConfFloat64() {
	type DummyStruct struct {
		Content float64
	}

	suite.Run("AutoConf for float64", func() {
		const data = float64(10)
		todoCopyExample := map[string][]interface{}{
			"Content": {"float64", data},
		}
		t := DummyStruct{}
		autoConf(&t, todoCopyExample)
		suite.Assert().Equal(t.Content, data)
	})
}

func (suite *TransformerTestSuite) TestAutoConfInt() {
	type DummyStruct struct {
		Content int
	}

	suite.Run("AutoConf for int", func() {
		const data = int(10)
		todoCopyExample := map[string][]interface{}{
			"Content": {"int", data},
		}
		t := DummyStruct{}
		autoConf(&t, todoCopyExample)
		suite.Assert().Equal(t.Content, data)
	})
}

func (suite *TransformerTestSuite) TestAutoConfInt8() {
	type DummyStruct struct {
		Content int8
	}

	suite.Run("AutoConf for int8", func() {
		const data = int8(10)
		todoCopyExample := map[string][]interface{}{
			"Content": {"int8", data},
		}
		t := DummyStruct{}
		autoConf(&t, todoCopyExample)
		suite.Assert().Equal(t.Content, data)
	})
}

func (suite *TransformerTestSuite) TestAutoConfInt16() {
	type DummyStruct struct {
		Content int16
	}

	suite.Run("AutoConf for int16", func() {
		const data = int16(10)
		todoCopyExample := map[string][]interface{}{
			"Content": {"int16", data},
		}
		t := DummyStruct{}
		autoConf(&t, todoCopyExample)
		suite.Assert().Equal(t.Content, data)
	})
}

func (suite *TransformerTestSuite) TestAutoConfInt32() {
	type DummyStruct struct {
		Content int32
	}

	suite.Run("AutoConf for int32", func() {
		const data = int32(10)
		todoCopyExample := map[string][]interface{}{
			"Content": {"int32", data},
		}
		t := DummyStruct{}
		autoConf(&t, todoCopyExample)
		suite.Assert().Equal(t.Content, data)
	})
}

func (suite *TransformerTestSuite) TestAutoConfInt64() {
	type DummyStruct struct {
		Content int64
	}

	suite.Run("AutoConf for int64", func() {
		const data = int64(10)
		todoCopyExample := map[string][]interface{}{
			"Content": {"int64", data},
		}
		t := DummyStruct{}
		autoConf(&t, todoCopyExample)
		suite.Assert().Equal(t.Content, data)
	})
}

func (suite *TransformerTestSuite) TestAutoConfUint() {
	type DummyStruct struct {
		Content uint
	}

	suite.Run("AutoConf for uint", func() {
		const data = uint(10)
		todoCopyExample := map[string][]interface{}{
			"Content": {"uint", data},
		}
		t := DummyStruct{}
		autoConf(&t, todoCopyExample)
		suite.Assert().Equal(t.Content, data)
	})
}

func (suite *TransformerTestSuite) TestAutoConfUint8() {
	type DummyStruct struct {
		Content uint8
	}

	suite.Run("AutoConf for uint8", func() {
		const data = uint8(10)
		todoCopyExample := map[string][]interface{}{
			"Content": {"uint8", data},
		}
		t := DummyStruct{}
		autoConf(&t, todoCopyExample)
		suite.Assert().Equal(t.Content, data)
	})
}

func (suite *TransformerTestSuite) TestAutoConfUint16() {
	type DummyStruct struct {
		Content uint16
	}

	suite.Run("AutoConf for uint16", func() {
		const data = uint16(10)
		todoCopyExample := map[string][]interface{}{
			"Content": {"uint16", data},
		}
		t := DummyStruct{}
		autoConf(&t, todoCopyExample)
		suite.Assert().Equal(t.Content, data)
	})
}

func (suite *TransformerTestSuite) TestAutoConfUint32() {
	type DummyStruct struct {
		Content uint32
	}

	suite.Run("AutoConf for uint32", func() {
		const data = uint32(10)
		todoCopyExample := map[string][]interface{}{
			"Content": {"uint32", data},
		}
		t := DummyStruct{}
		autoConf(&t, todoCopyExample)
		suite.Assert().Equal(t.Content, data)
	})
}

func (suite *TransformerTestSuite) TestAutoConfUint64() {
	type DummyStruct struct {
		Content uint64
	}

	suite.Run("AutoConf for uint64", func() {
		const data = uint64(10)
		todoCopyExample := map[string][]interface{}{
			"Content": {"uint64", data},
		}
		t := DummyStruct{}
		autoConf(&t, todoCopyExample)
		suite.Assert().Equal(t.Content, data)
	})
}

func (suite *TransformerTestSuite) TestAutoConfUintprt() {
	type DummyStruct struct {
		Content uintptr
	}

	suite.Run("AutoConf for uintptr", func() {
		const data = uintptr(10)
		todoCopyExample := map[string][]interface{}{
			"Content": {"uintptr", data},
		}
		t := DummyStruct{}
		autoConf(&t, todoCopyExample)
		suite.Assert().Equal(t.Content, data)
	})
}

func (suite *TransformerTestSuite) TestTransformerToProto() {
	suite.Run("TransformerToProto", func() {
		uid := uuid.NewV4().String()
		person := model.Todo{
			ID:            uid,
			Name:          "Isa",
			Completed:     false,
			NumberCode:    129520,
			NumberProduct: 25983578228,
		}
		codes := []model.Code{{1}, {4}, {9}}

		pPerson := pb.Todo{}

		for _, i := range codes {
			var localPbCode pb.Code
			Transformed{From: &i}.TransformerToProto(&localPbCode)
			pPerson.Codes = append(pPerson.Codes, &localPbCode)
		}
		Transformed{From: &person}.TransformerToProto(&pPerson)

		suite.Assert().Len(pPerson.Codes, 3)
		suite.Assert().Equal(pPerson.Id, uid)
		suite.Assert().Equal(pPerson.Name, "Isa")
		suite.Assert().Equal(pPerson.NumberCode, int32(129520))
		suite.Assert().Equal(pPerson.NumberProduct, uint64(25983578228))
	})
}
