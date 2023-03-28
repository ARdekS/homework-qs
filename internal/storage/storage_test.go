package storage

import (
	"testing"

	"github.com/ARdekS/homework-qs/mocks"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

func TestMyFunction(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockNode := mocks.NewMockNode(ctrl)

	// настройка поведения мока
	mockNode.EXPECT().Edit(mockNode).Return()

	// тестирование функции с использованием мока
	result, err := myFunction(mockDB, "key")
	assert.Nil(t, err)
	assert.Equal(t, "value", result)
}
