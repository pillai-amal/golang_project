package greetings

import (
	"errors"
	"fmt"
	"math/rand"
	"time"
)

//the fuction to return greeting when a for a single user

func Hello(name string) (string, error) {
	if name == "" {
		return "", errors.New("Empty Value")
	}
	message := fmt.Sprintf(randomFormat(), name)
	return message, nil
}

//since overloading is not  feature og GO using a diffrent
//function name to return gerreting for multiple users

func Hellos(names []string) (map[string]string, error) {
	messages := make(map[string]string)
	for _, name := range names {
		message, err := Hello(name)
		if err != nil {
			return nil, err
		}
		messages[name] = message
	}
	return messages, nil
}

//init function is used to intialize the values
//here a random seed is set

func init() {
	rand.Seed(time.Now().UnixNano())
}

func randomFormat() string {
	format := []string{
		"Great to see you %v :)",
		"Hi %v Welcome !",
		"Hail, %v! Well met!",
	}
	return format[rand.Intn(len(format))]
}
