package main

import (
	"encoding/json"
	"errors"
	"fmt"
)

type HelloWorld struct {
	Hello string
}

type CreateUserRequest struct {
	Email                string `json:"email"`
	Password             string `json:"password"`
	PasswordConfirmation string `json:"password_confirmation"`
}

// validation errors
var (
	ErrEmailRequired                = errors.New("email is required")                             // когда поле email не заполнено
	ErrPasswordRequired             = errors.New("password is required")                          // когда поле password не заполнено
	ErrPasswordConfirmationRequired = errors.New("password confirmation is required")             // когда поле password_confirmation не заполнено
	ErrPasswordDoesNotMatch         = errors.New("password does not match with the confirmation") // когда поля password и password_confirmation не совпадают
)

func DecodeAndValidateRequest(requestBody []byte) (CreateUserRequest, error) {
	userRequest := CreateUserRequest{}

	err := json.Unmarshal(requestBody, &userRequest)
	if err != nil {
		return CreateUserRequest{}, err
	}

	if userRequest.Email == "" {
		return CreateUserRequest{}, ErrEmailRequired
	}
	if userRequest.Password == "" {
		return CreateUserRequest{}, ErrPasswordRequired
	}
	if userRequest.PasswordConfirmation == "" {
		return CreateUserRequest{}, ErrPasswordConfirmationRequired
	}
	if userRequest.Password != userRequest.PasswordConfirmation {
		return CreateUserRequest{}, ErrPasswordDoesNotMatch
	}

	return userRequest, nil
}

func main() {
	hw := HelloWorld{}
	// первым аргументом передается JSON-строка в виде слайса байт. Вторым аргументом указатель на структуру, в которую нужно декодировать результат.
	err := json.Unmarshal([]byte("{\"hello\":\"world\"}"), &hw)
	fmt.Printf("error: %s, struct: %+v\n", err, hw) // error: %!s(<nil>), struct: {Hello:world}

	fmt.Println(DecodeAndValidateRequest([]byte("{\"email\":\"\",\"password\":\"test\",\"password_confirmation\":\"test\"}")))                               // CreateUserRequest{}, "email is required"
	fmt.Println(DecodeAndValidateRequest([]byte("{\"email\":\"test\",\"password\":\"\",\"password_confirmation\":\"test\"}")))                               // CreateUserRequest{}, "password is required"
	fmt.Println(DecodeAndValidateRequest([]byte("{\"email\":\"test\",\"password\":\"test\",\"password_confirmation\":\"\"}")))                               // CreateUserRequest{}, "password confirmation is required"
	fmt.Println(DecodeAndValidateRequest([]byte("{\"email\":\"test\",\"password\":\"test\",\"password_confirmation\":\"test2\"}")))                          // CreateUserRequest{}, "password does not match with the confirmation"
	fmt.Println(DecodeAndValidateRequest([]byte("{\"email\":\"email@test.com\",\"password\":\"passwordtest\",\"password_confirmation\":\"passwordtest\"}"))) // CreateUserRequest{Email:"email@test.com", Password:"passwordtest", PasswordConfirmation:"passwordtest"}, nil
}
