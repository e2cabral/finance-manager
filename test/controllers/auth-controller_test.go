// +build e2e

package test

import (
	"finance-manager/test"
	"fmt"
	"github.com/stretchr/testify/assert"
	"gopkg.in/resty.v1"
	"testing"
)

func TestLoginEndpoint(t *testing.T) {
	fmt.Println("Running tests to the login endpoint.")

	client := resty.New()

	r, err := client.
		R().
		SetBody(`{"username": "edinhodograu", "password": "senhaforte"}`).
		Post(fmt.Sprintf("%s/login", test.BaseUrl))

	assert.NoError(t, err)

	assert.Equal(t, 200, r.StatusCode())
}

func TestSignUpEndpoint(t *testing.T) {
	fmt.Println("Running tests to the signup endpoint.")

	client := resty.New()

	r, err := client.
		R().
		SetBody(`{"username": "eddietest", "password": "strongpassword"}`).
		Post(fmt.Sprintf("%s/sign-up", test.BaseUrl))

	assert.NoError(t, err)

	assert.Equal(t, 200, r.StatusCode())
}
