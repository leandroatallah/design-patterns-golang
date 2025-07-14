package main

import (
	"fmt"
	"strings"
)

type ProfileBuilderFacade struct {
	builder *UserProfileBuilder
}

func (f *ProfileBuilderFacade) CreateUser(firstName, lastName string) interface{} {
	if f.builder == nil {
		f.builder = NewUserProfileBuilder()
	}

	f.builder.Reset()
	f.builder.SetFirstName(firstName)
	f.builder.SetLastName(lastName)
	f.builder.SetEmail(strings.ToLower(fmt.Sprintf("%s%s@mail.com", firstName, lastName)))
	return f.builder.Result()
}

func main() {
	profileBuilderFacade := &ProfileBuilderFacade{}
	user := profileBuilderFacade.CreateUser("Robert", "Martin")

	fmt.Println(user)
}
