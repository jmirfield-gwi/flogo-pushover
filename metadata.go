package pushover

import "github.com/project-flogo/core/data/coerce"

type Settings struct {
	AppToken   string `md:"appToken,required"`
	GroupToken string `md:"groupToken,required"`
	Active     bool   `md:"active,required"`
}

type Input struct {
	Message string `md:"message,required"`
}

func (r *Input) FromMap(values map[string]interface{}) error {
	strVal, _ := coerce.ToString(values["message"])
	r.Message = strVal
	return nil
}

func (r *Input) ToMap() map[string]interface{} {
	return map[string]interface{}{
		"message": r.Message,
	}
}

type Output struct {
	Status int `md:"status,required"`
}

func (o *Output) FromMap(values map[string]interface{}) error {
	intVal, _ := coerce.ToInt(values["status"])
	o.Status = intVal
	return nil
}

func (o *Output) ToMap() map[string]interface{} {
	return map[string]interface{}{
		"status": o.Status,
	}
}
