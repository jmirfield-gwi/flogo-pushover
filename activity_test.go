package pushover

import (
	"log"
	"os"
	"testing"

	"github.com/joho/godotenv"
	"github.com/project-flogo/core/activity"
	"github.com/project-flogo/core/data"
	"github.com/project-flogo/core/data/mapper"
	"github.com/project-flogo/core/data/metadata"
	logger "github.com/project-flogo/core/support/log"
	"github.com/project-flogo/core/support/trace"
	"github.com/stretchr/testify/assert"
)

func init() {
	err := godotenv.Load(".env")

	if err != nil {
		log.Fatal("Error loading .env file")
	}
}

type initContext struct {
	settings map[string]interface{}
}

func newInitContext(values map[string]interface{}) *initContext {
	if values == nil {
		values = make(map[string]interface{})
	}
	return &initContext{
		settings: values,
	}
}

func (i *initContext) Settings() map[string]interface{} {
	return i.settings
}

func (i *initContext) Logger() logger.Logger {
	return logger.RootLogger()
}

func (i *initContext) MapperFactory() mapper.Factory {
	return nil
}

type activityContext struct {
	input  map[string]interface{}
	output map[string]interface{}
}

func newActivityContext(values map[string]interface{}) *activityContext {
	if values == nil {
		values = make(map[string]interface{})
	}
	return &activityContext{
		input:  values,
		output: make(map[string]interface{}),
	}
}

func (a *activityContext) ActivityHost() activity.Host {
	return a
}

func (a *activityContext) Name() string {
	return "test"
}

func (a *activityContext) GetInput(name string) interface{} {
	return a.input[name]
}

func (a *activityContext) SetOutput(name string, value interface{}) error {
	a.output[name] = value
	return nil
}

func (a *activityContext) GetInputObject(input data.StructValue) error {
	return input.FromMap(a.input)
}

func (a *activityContext) SetOutputObject(output data.StructValue) error {
	a.output = output.ToMap()
	return nil
}

func (a *activityContext) GetSharedTempData() map[string]interface{} {
	return nil
}

func (a *activityContext) ID() string {
	return "test"
}
func (a *activityContext) GetTracingContext() trace.TracingContext {
	return nil
}

func (a *activityContext) IOMetadata() *metadata.IOMetadata {
	return nil
}

func (a *activityContext) Reply(replyData map[string]interface{}, err error) {

}

func (a *activityContext) Return(returnData map[string]interface{}, err error) {

}

func (a *activityContext) Scope() data.Scope {
	return nil
}

func (a *activityContext) Logger() logger.Logger {
	return logger.RootLogger()
}

func TestInactivePushover(t *testing.T) {
	act, err := New(newInitContext(map[string]interface{}{
		"appToken":   "test1",
		"groupToken": "test2",
		"active":     false,
	}))
	assert.Nil(t, err)

	ctx := newActivityContext(map[string]interface{}{
		"message": "Hello world",
	})
	_, err = act.Eval(ctx)
	assert.Nil(t, err)
	assert.Equal(t, 204, ctx.output["status"])
}

func TestActivePushover(t *testing.T) {
	act, err := New(newInitContext(map[string]interface{}{
		"appToken":   os.Getenv("APP"),
		"groupToken": os.Getenv("GROUP"),
		"active":     true,
	}))
	assert.Nil(t, err)

	ctx := newActivityContext(map[string]interface{}{
		"message": "Hello world",
	})
	_, err = act.Eval(ctx)
	assert.Nil(t, err)
	assert.Equal(t, 200, ctx.output["status"])
}

func TestActivePushoverEmptyMessage(t *testing.T) {
	act, err := New(newInitContext(map[string]interface{}{
		"appToken":   "test1",
		"groupToken": "test2",
		"active":     true,
	}))
	assert.Nil(t, err)

	ctx := newActivityContext(map[string]interface{}{
		"message": "",
	})
	_, err = act.Eval(ctx)
	assert.Nil(t, err)
	assert.Equal(t, 204, ctx.output["status"])
}

func TestActivePushoverBadTokens(t *testing.T) {
	act, err := New(newInitContext(map[string]interface{}{
		"appToken":   "test1",
		"groupToken": "test2",
		"active":     true,
	}))
	assert.Nil(t, err)

	ctx := newActivityContext(map[string]interface{}{
		"message": "Bad tokens",
	})
	_, err = act.Eval(ctx)
	assert.Nil(t, err)
	assert.Equal(t, 400, ctx.output["status"])
}

func TestMissingAppToken(t *testing.T) {
	_, err := New(newInitContext(map[string]interface{}{
		"groupToken": "test2",
		"active":     true,
	}))
	assert.Error(t, err)
}

func TestMissingGroupToken(t *testing.T) {
	_, err := New(newInitContext(map[string]interface{}{
		"appToken": "test1",
		"active":   true,
	}))
	assert.Error(t, err)
}

func TestMissingActiveField(t *testing.T) {
	_, err := New(newInitContext(map[string]interface{}{
		"appToken":   "test1",
		"groupToken": "test2",
	}))
	assert.Error(t, err)
}

func TestNoActivityContext(t *testing.T) {
	act, err := New(newInitContext(map[string]interface{}{
		"appToken":   os.Getenv("APP"),
		"groupToken": os.Getenv("GROUP"),
		"active":     true,
	}))
	assert.Nil(t, err)

	ctx := newActivityContext(map[string]interface{}{})
	_, err = act.Eval(ctx)
	assert.Nil(t, err)
	assert.Equal(t, 204, ctx.output["status"])
}
