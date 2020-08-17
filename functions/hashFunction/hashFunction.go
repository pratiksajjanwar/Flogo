package stringFunctions

import (
	"github.com/project-flogo/core/data"
	"github.com/project-flogo/core/data/expression/function"
	"crypto/sha256"
	"encoding"
	"fmt"
	"bytes"
	"log"
	"hash"
)

func init() {
	_ = function.Register(&hashFunction{})
}

type hashFunction struct {
}

func (s *hashFunction) Name() string {
	return "hashFunction"
}

func (s *hashFunction) Sig() (paramTypes []data.Type, isVariadic bool) {
	return []data.Type{data.TypeString,data.TypeString}, true
}

func (s *hashFunction) Eval(params ...interface{}) (interface{}, error) {	
	example1 := params[0].(string)
	example2 := params[1].(string)
	if example1 == "" || example2 == "" {
		return "", fmt.Errorf("Parameters cannot be empty")
	}
	var firstHash hash.Hash
	firstHash = sha256.New()
	firstHash.Write([]byte(example1))
	var marshaler encoding.BinaryMarshaler
	var ok bool
	marshaler, ok = firstHash.(encoding.BinaryMarshaler)
	if !ok {
	log.Fatal("first Hash is not generated by encoding.BinaryMarshaler")
	}
	var data []byte
	var err error
	data, err = marshaler.MarshalBinary()
	if err != nil {
	log.Fatal("failure to create first Hash:", err)
	}
	var secondHash hash.Hash
	secondHash = sha256.New()
	var unmarshaler encoding.BinaryUnmarshaler
	unmarshaler, ok = secondHash.(encoding.BinaryUnmarshaler)
	if !ok {
	log.Fatal("second Hash is not generated by encoding.BinaryUnmarshaler")
	}
	if err := unmarshaler.UnmarshalBinary(data); err != nil {
	log.Fatal("failure to create hash:", err)
	}
	firstHash.Write([]byte(example2))
	secondHash.Write([]byte(example2))
	fmt.Printf("%x\n", firstHash.Sum(nil))
	return firstHash.Sum(nil),nil
}
