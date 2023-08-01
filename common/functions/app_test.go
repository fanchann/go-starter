package functions_test

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"

	functions_mocks "github.com/fanchann/go-starter/common/functions/mock"
)

var (
	funcMock = functions_mocks.IAppPathGenerator{mock.Mock{}}
	errorMsg = errors.New("failed generate structure")
)

func TestCreatePathAppFailed(t *testing.T) {
	funcMock.Mock.On("CreateAppPath", "").Return(errorMsg)
	err := funcMock.CreateAppPath("")

	assert.EqualError(t, err, errorMsg.Error())
}

func TestCreateAppSuccess(t *testing.T) {
	funcMock.Mock.On("CreateAppPath", "$HOME/app").Return(nil)
	err := funcMock.CreateAppPath("$HOME/app")

	assert.Nil(t, err)
}

func TestGenerateCodeSuccess(t *testing.T) {
	funcMock.Mock.On("GenerateAppCode", "HelloWorld").Return(nil)
	err := funcMock.GenerateAppCode("HelloWorld")

	assert.Nil(t, err)
}

func TestGenerateCodeFailed(t *testing.T) {
	funcMock.Mock.On("GenerateAppCode", "").Return(errorMsg)

	err := funcMock.GenerateAppCode("")

	assert.NotNil(t, err)
	assert.EqualError(t, err, errorMsg.Error())
}
