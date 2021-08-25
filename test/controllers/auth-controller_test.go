// +build e2e

package test

import (
	"finance-manager/src/domain/models"
	"finance-manager/test"
	"fmt"
	"github.com/stretchr/testify/assert"
	"gopkg.in/resty.v1"
	"testing"
)

func TestLoginEndpoint(t *testing.T) {
	fmt.Println("Running tests to the login endpoint.")

	client := resty.New()

	body, err := client.JSONMarshal(models.User{Username: "edinhodograu", Password: "senhaforte"})
	if err != nil {
		t.Fail()
	}

	r, err := client.R().SetBody(body).Post(fmt.Sprintf("%s/login", test.BASE_URL))
	if err != nil {
		t.Fail()
	}

	assert.Equal(t, 200, r.StatusCode())
}

func TestSignUpEndpoint(t *testing.T) {
	fmt.Println("Running tests to the signup endpoint.")

	client := resty.New()

	body, err := client.JSONMarshal(models.User{Username: "eddietest", Password: "strongpassword"})
	if err != nil {
		t.Fail()
	}

	r, err := client.R().SetBody(body).Post(fmt.Sprintf("%s/sign-up", test.BASE_URL))
	if err != nil {
		t.Fail()
	}

	assert.Equal(t, 200, r.StatusCode())
}
