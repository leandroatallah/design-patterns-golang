package main

import "fmt"

type ProfileBuilder interface {
	SetFirstName(text string)
	SetLastName(text string)
	SetEmail(text string)
	Reset()
	Result() interface{}
}

type UserProfile struct {
	firstName string
	lastName  string
	email     string
}
type AdminProfile struct {
	firstName   string
	lastName    string
	email       string
	isSuperUser bool
}

type UserProfileBuilder struct {
	userProfile UserProfile
}

func (b *UserProfileBuilder) Reset() {
	b.userProfile.firstName = ""
	b.userProfile.lastName = ""
	b.userProfile.email = ""
}
func (b *UserProfileBuilder) SetFirstName(text string) {
	b.userProfile.firstName = text
}
func (b *UserProfileBuilder) SetLastName(text string) {
	b.userProfile.lastName = text
}
func (b *UserProfileBuilder) SetEmail(text string) {
	b.userProfile.email = text
}

func (b *UserProfileBuilder) Result() (result interface{}) {
	defer b.Reset()
	result = b.userProfile
	return
}

type AdminProfileBuilder struct {
	adminProfile AdminProfile
}

func (b *AdminProfileBuilder) Reset() {
	b.adminProfile.firstName = ""
	b.adminProfile.lastName = ""
	b.adminProfile.email = ""
	b.adminProfile.isSuperUser = false
}
func (b *AdminProfileBuilder) SetFirstName(text string) {
	b.adminProfile.firstName = text
}
func (b *AdminProfileBuilder) SetLastName(text string) {
	b.adminProfile.lastName = text
}
func (b *AdminProfileBuilder) SetEmail(text string) {
	b.adminProfile.email = text
}

func (b *AdminProfileBuilder) Result() (result interface{}) {
	defer b.Reset()
	result = b.adminProfile
	return
}

func (b *AdminProfileBuilder) setSuperUser() {
	b.adminProfile.isSuperUser = true
}

type ProfileDirector struct{}

func (pd *ProfileDirector) createProfile(builder ProfileBuilder) {
	builder.Reset()
	builder.SetFirstName("Leandro")
	builder.SetLastName("Atallah")

	if admin, ok := builder.(*AdminProfileBuilder); ok {
		builder.SetEmail("admin@mail.com")
		admin.setSuperUser()
	} else {
		builder.SetEmail("leandro@mail.com")
	}
}

func NewUserProfileBuilder() *UserProfileBuilder {
	return &UserProfileBuilder{}
}

func NewAdminProfileBuilder() *AdminProfileBuilder {
	return &AdminProfileBuilder{}
}

func main() {
	director := ProfileDirector{}

	userBuilder := NewUserProfileBuilder()
	director.createProfile(userBuilder)
	userResult := userBuilder.Result()
	fmt.Println(userResult)

	adminBuilder := NewAdminProfileBuilder()
	director.createProfile(adminBuilder)
	adminResult := adminBuilder.Result()
	fmt.Println(adminResult)
}
