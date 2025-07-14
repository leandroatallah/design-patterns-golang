package main

type ProfileBuilder interface {
	SetFirstName(text string)
	SetLastName(text string)
	SetEmail(text string)
	Reset()
	Result() interface{}
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

func NewUserProfileBuilder() *UserProfileBuilder {
	return &UserProfileBuilder{}
}

func NewAdminProfileBuilder() *AdminProfileBuilder {
	return &AdminProfileBuilder{}
}
