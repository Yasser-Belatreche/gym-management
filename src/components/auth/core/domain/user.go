package domain

import (
	"gym-management/src/lib/primitives/application_specific"
	"gym-management/src/lib/primitives/generic"
	"time"
)

type User struct {
	id         string
	usernames  []Username
	password   Password
	role       string
	profile    *application_specific.UserProfile
	restricted bool
	createdAt  time.Time
	updatedAt  time.Time
	lastLogin  *time.Time
	createdBy  string
	updatedBy  string
	deletedAt  *time.Time
	deletedBy  *string
}

type UserState struct {
	Id         string
	Usernames  []string
	Password   string
	Role       string
	Profile    application_specific.UserProfile
	Restricted bool
	CreatedAt  time.Time
	LastLogin  *time.Time
	UpdatedAt  time.Time
	CreatedBy  string
	UpdatedBy  string
	DeletedAt  *time.Time
	DeletedBy  *string
}

func CreateUser(id string, usernames []Username, password Password, role string, profile *application_specific.UserProfile, createdBy string) *User {
	return &User{
		id:         id,
		usernames:  usernames,
		password:   password,
		role:       role,
		profile:    profile,
		restricted: false,
		lastLogin:  nil,
		createdAt:  time.Now(),
		updatedAt:  time.Now(),
		createdBy:  createdBy,
		updatedBy:  createdBy,
		deletedAt:  nil,
		deletedBy:  nil,
	}
}

func CreateAdmin(firstName string, lastName string, phone string, email application_specific.Email, password Password, by string) *User {
	return CreateUser(
		generic.GenerateUUID(),
		[]Username{
			UsernameFromEmail(email),
		},
		password,
		RoleAdmin,
		&application_specific.UserProfile{
			FirstName: firstName,
			LastName:  lastName,
			Phone:     phone,
			Email:     email.Value,
			OwnedGyms: nil,
		},
		by,
	)
}

func UserFromState(state UserState) *User {
	usernames := make([]Username, len(state.Usernames))

	for i, username := range state.Usernames {
		usernames[i] = Username(username)
	}

	return &User{
		id:         state.Id,
		usernames:  usernames,
		password:   Password(state.Password),
		role:       state.Role,
		profile:    &state.Profile,
		restricted: state.Restricted,
		createdAt:  state.CreatedAt,
		updatedAt:  state.UpdatedAt,
		lastLogin:  state.LastLogin,
		createdBy:  state.CreatedBy,
		updatedBy:  state.UpdatedBy,
		deletedAt:  state.DeletedAt,
		deletedBy:  state.DeletedBy,
	}
}

func (u *User) ChangeUsernames(usernames []Username) {
	u.usernames = usernames
	u.updatedAt = time.Now()
}

func (u *User) ChangePassword(password Password) {
	u.password = password
	u.updatedAt = time.Now()
}

func (u *User) SetProfile(profile *application_specific.UserProfile) {
	u.profile = profile
	u.updatedAt = time.Now()
}

func (u *User) Restrict() {
	u.restricted = true
	u.updatedAt = time.Now()
}

func (u *User) Unrestrict() {
	u.restricted = false
	u.updatedAt = time.Now()
}

func (u *User) Delete(by string) {
	now := time.Now()
	u.restricted = true
	u.deletedAt = &now
	u.deletedBy = &by
}

func (u *User) Restore() {
	u.restricted = false
	u.deletedAt = nil
	u.deletedBy = nil
}

func (u *User) Login(password string, tokenSecret string) (Token, *application_specific.ApplicationException) {
	if u.IsRestricted() {
		return "", application_specific.NewAuthenticationException("AUTH.USER.RESTRICTED", "User is restricted", map[string]string{
			"id": u.id,
		})
	}

	if u.IsDeleted() {
		return "", application_specific.NewAuthenticationException("AUTH.USER.DELETED", "User is deleted", map[string]string{
			"id": u.id,
		})
	}

	if !u.password.Equals(password) {
		return "", application_specific.NewValidationException("AUTH.LOGIN.INVALID", "Invalid Credentials", map[string]string{})
	}

	token, err := NewToken(
		TokenClaims{UserId: u.id, Role: u.role},
		tokenSecret,
	)
	if err != nil {
		return "", err
	}

	now := time.Now()
	u.lastLogin = &now
	u.updatedAt = now

	return token, nil
}

func (u *User) GetSession() (*application_specific.UserSession, *application_specific.ApplicationException) {
	if u.IsRestricted() {
		return nil, application_specific.NewAuthenticationException("AUTH.USER.RESTRICTED", "User is restricted", map[string]string{
			"id": u.id,
		})
	}

	if u.IsDeleted() {
		return nil, application_specific.NewAuthenticationException("AUTH.USER.DELETED", "User is deleted", map[string]string{
			"id": u.id,
		})
	}

	return application_specific.NewUserSession(u.id, u.role, u.profile), nil
}

func (u *User) IsRestricted() bool {
	return u.restricted
}

func (u *User) IsDeleted() bool {
	return u.deletedAt != nil
}

func (u *User) State() UserState {
	usernames := make([]string, len(u.usernames))

	for i, username := range u.usernames {
		usernames[i] = username.Value()
	}

	return UserState{
		Id:         u.id,
		Usernames:  usernames,
		Password:   u.password.Value(),
		Role:       u.role,
		Profile:    *u.profile,
		Restricted: u.restricted,
		CreatedAt:  u.createdAt,
		UpdatedAt:  u.updatedAt,
		CreatedBy:  u.createdBy,
		UpdatedBy:  u.updatedBy,
		DeletedAt:  u.deletedAt,
		DeletedBy:  u.deletedBy,
	}
}
