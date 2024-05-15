package infra

import (
	"errors"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"gym-management/src/components/auth/core/domain"
	"gym-management/src/lib"
	"gym-management/src/lib/persistence/psql/gorm/customer_types"
	"gym-management/src/lib/persistence/psql/gorm/models"
	"gym-management/src/lib/primitives/application_specific"
)

type GormUserRepository struct {
}

func (g *GormUserRepository) FindByUsername(username domain.Username, session *application_specific.Session) (*domain.User, *application_specific.ApplicationException) {
	db := lib.GormDB(session)

	var usernameModel models.Username
	result := db.First(&usernameModel, "username = ?", username)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return nil, application_specific.NewNotFoundException("USERS.NOT_FOUND", "User not found", map[string]string{
				"username": string(username),
			})
		}
		return nil, application_specific.NewUnknownException("FAILED_TO_FIND_USER", "Failed to find user", map[string]string{
			"error": result.Error.Error(),
		})
	}

	var user models.User
	result = db.Clauses(clause.Locking{Strength: "UPDATE"}).Preload("Usernames").First(&user, "id = ?", usernameModel.UserId)
	if result.Error != nil {
		return nil, application_specific.NewUnknownException("FAILED_TO_FIND_USER", "Failed to find user", map[string]string{
			"error": result.Error.Error(),
		})
	}

	domainUser := toDomain(&user)

	return domainUser, nil
}

func (g *GormUserRepository) UsernameUsed(username domain.Username, session *application_specific.Session) (bool, *application_specific.ApplicationException) {
	db := lib.GormDB(session)

	var usernameModel models.Username
	result := db.First(&usernameModel, "username = ?", username)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return false, nil
		}
		return false, application_specific.NewUnknownException("FAILED_TO_CHECK_USERNAME", "Failed to check username", map[string]string{
			"error": result.Error.Error(),
		})
	}

	return true, nil
}

func (g *GormUserRepository) FindByID(id string, session *application_specific.Session) (*domain.User, *application_specific.ApplicationException) {
	db := lib.GormDB(session)

	var user models.User
	result := db.Clauses(clause.Locking{Strength: "UPDATE"}).Preload("Usernames").First(&user, id)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return nil, application_specific.NewNotFoundException("USERS.NOT_FOUND", "User not found", map[string]string{
				"id": id,
			})
		}
		return nil, application_specific.NewUnknownException("FAILED_TO_FIND_USER", "Failed to find user", map[string]string{
			"error": result.Error.Error(),
		})
	}

	domainUser := toDomain(&user)

	return domainUser, nil
}

func (g *GormUserRepository) Create(user *domain.User, session *application_specific.Session) *application_specific.ApplicationException {
	db := lib.GormDB(session)

	model := toDBModel(user)

	result := db.Create(model)
	if result.Error != nil {
		return application_specific.NewUnknownException("FAILED_TO_CREATE_USER", "Failed to save user", map[string]string{
			"error": result.Error.Error(),
		})
	}

	return nil
}

func (g *GormUserRepository) Update(user *domain.User, session *application_specific.Session) *application_specific.ApplicationException {
	db := lib.GormDB(session)

	model := toDBModel(user)

	result := db.Save(model)
	if result.Error != nil {
		return application_specific.NewUnknownException("FAILED_TO_UPDATE_USER", "Failed to update user", map[string]string{
			"error": result.Error.Error(),
		})
	}

	return nil
}

func toDomain(dbModel *models.User) *domain.User {
	usernames := make([]string, 0, len(dbModel.Usernames))

	for _, username := range dbModel.Usernames {
		usernames = append(usernames, username.Username)
	}

	profile := application_specific.UserProfile{
		FirstName:        dbModel.Profile["first_name"].(string),
		LastName:         dbModel.Profile["last_name"].(string),
		Phone:            dbModel.Profile["phone"].(string),
		Email:            dbModel.Profile["email"].(string),
		OwnedGyms:        dbModel.Profile["owned_gyms"].([]string),
		EnabledOwnedGyms: dbModel.Profile["enabled_owned_gyms"].([]string),
	}

	state := domain.UserState{
		Id:         dbModel.Id,
		Usernames:  usernames,
		Password:   dbModel.Password,
		Role:       dbModel.Role,
		Profile:    profile,
		Restricted: dbModel.Restricted,
		LastLogin:  dbModel.LastLogin,
		DeletedAt:  dbModel.DeletedAt,
	}

	return domain.UserFromState(state)
}

func toDBModel(domainModel *domain.User) *models.User {
	state := domainModel.State()

	usernames := make([]models.Username, 0, len(state.Usernames))
	for _, username := range state.Usernames {
		usernames = append(usernames, models.Username{
			UserId:   state.Id,
			Username: username,
		})
	}

	user := &models.User{
		Id:        state.Id,
		Usernames: usernames,
		Password:  state.Password,
		Role:      state.Role,
		Profile: customer_types.JSONB{
			"first_name":         state.Profile.FirstName,
			"last_name":          state.Profile.LastName,
			"phone":              state.Profile.Phone,
			"email":              state.Profile.Email,
			"owned_gyms":         state.Profile.OwnedGyms,
			"enabled_owned_gyms": state.Profile.EnabledOwnedGyms,
		},
		Restricted: state.Restricted,
		LastLogin:  state.LastLogin,
		DeletedAt:  state.DeletedAt,
	}

	return user
}
