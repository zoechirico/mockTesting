package main

import (
	"errors"
	"fmt"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"testing"
)

func TestSuccessfulCheckoutShoppingCart(t *testing.T) {
	givenNotificationMsg := "Successfully purchased 10 books"
	notifier := new(notifierMock)
	notifier.On("SendMessage", givenNotificationMsg).Return(nil)

	purchaseService := PurchaseService{notifier: notifier}
	err := purchaseService.CheckoutShoppingCart(99, givenNotificationMsg)

	notifier.AssertExpectations(t)
	assert.Nil(t, err)
}

func TestFailedCheckoutShoppingCart(t *testing.T) {
	givenNotificationMsg := "Successfully purchased 10 books"
	expectedErrMsg := "failed to send message"
	notifier := new(notifierMock)
	notifier.On("SendMessage", givenNotificationMsg).Return(errors.New(expectedErrMsg))

	purchaseService := PurchaseService{notifier: notifier}
	err := purchaseService.CheckoutShoppingCart(99, givenNotificationMsg)

	notifier.AssertExpectations(t)
	assert.NotNil(t, err)
	assert.Equal(t, expectedErrMsg, err.Error())
}

func TestSpaz(t *testing.T) {
	givenNotificationMsg := "Successfully purchased 10 books"
	expectedErrMsg := "failed to send message"
	spazN := new(spazMock)
	spazN.On("SendMessage", givenNotificationMsg).Return(expectedErrMsg)
	spazN.On("SendMessage", "two").Return("two", "three")

	r := spazN.SendMessage(givenNotificationMsg)
	fmt.Println("Wow: " + r)
	r = spazN.SendMessage("two")
	fmt.Println("Wow: " + r)

}

type notifierMock struct {
	mock.Mock
}

func (m *notifierMock) SendMessage(message string) error {
	args := m.Called(message)

	return args.Error(0)
}

type spazMock struct {
	mock.Mock
}

func (m *spazMock) SendMessage(message string) string {
	args := m.Called(message)
	fmt.Println(args.Get(0))

	if len(args) > 1 {
		return args.String(1)
	}
	return args.String(0)
}
