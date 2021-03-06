package DateFunction

import (
	
	"github.com/project-flogo/core/data"
	"github.com/project-flogo/core/data/expression/function"
         "time"
)

func init() {
	_=function.Register(&dateFunction{})
}

type dateFunction struct {
}

func (s *dateFunction) Name() string {
	return "dateFunction"
}

func (s *dateFunction) Sig() (paramTypes []data.Type, isVariadic bool) {
	return []data.Type{data.TypeString}, true
}

func (s *dateFunction) Eval(params ...interface{}) (interface{}, error) {	
	
	current_date := time.Now().Local()
	
	return current_date,nil
}
