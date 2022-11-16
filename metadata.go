package pushover

import "github.com/project-flogo/core/data/coerce"

type Settings struct {
	AppToken   string `md:"App Token,required"`
	GroupToken string `md:"Group Token,required"`
}

type Input struct {
	Message string `md:"Message,required"`
}

func (r *Input) FromMap(values map[string]interface{}) error {
	strVal, _ := coerce.ToString(values["Message"])
	r.Message = strVal
	return nil
}

func (r *Input) ToMap() map[string]interface{} {
	return map[string]interface{}{
		"Message": r.Message,
	}
}

type Output struct {
	Status int `md:"Status,required"`
}

func (o *Output) FromMap(values map[string]interface{}) error {
	intVal, _ := coerce.ToInt(values["Status"])
	o.Status = intVal
	return nil
}

func (o *Output) ToMap() map[string]interface{} {
	return map[string]interface{}{
		"Status": o.Status,
	}
}
