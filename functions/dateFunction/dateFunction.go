package DateFunction

import (
  "fmt"
  "time"
)

type dateFunction struct {
}

func init() {
	_ = function.Register(&dateFunction{})
}

func (s *dateFunction) Name() string {
	return "date"
}

func (s *hashFunction) Sig() (paramTypes []data.Type, isVariadic bool) {
	return []data.Type{data.TypeString}, false
}

func (s *dateFunction) Eval(params ...interface{}) (interface{}, error) {	
	
	current_time := time.Now().Local()
	
	return current_time,nil
}