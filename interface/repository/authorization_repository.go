package repository

import (
	"GoAds/domain"
	"GoAds/domain/model"
	"GoAds/usecase/repository"
	"github.com/jinzhu/gorm"
	"strconv"
	"time"
)

type authorizationRepository struct {
	db *gorm.DB
}

func NewAuthorizationRepository(db *gorm.DB) repository.AuthorizationRepository {
	return &authorizationRepository{db}
}

func (ar *authorizationRepository) Create(u *model.User) (*model.User, error) {
	err := ar.db.Model(&u).Create(u).Error

	if err != nil {
		if err.Error() == domain.ErrUserAlreadyWithEmail {
			return nil, domain.ErrUserEmailAlreadyExists
		} else if err.Error() == domain.ErrUserAlreadyWithPhone {
			return nil, domain.ErrUserPhoneAlreadyExists
		}
		return nil, domain.ErrUserInternalServerError
	}
	return u, nil
}

func (ar *authorizationRepository) CreateSession(s *model.Session) (*model.Session, error) {
	err := ar.db.Model(&s).Create(s).Error

	if err != nil {
		return nil, err
	}

	return s, nil
}

func (ar *authorizationRepository) UserExists(email string) (string, string, error) {
	user := model.User{}
	err := ar.db.Model(&user).Select("*").Where("email = ?", email).Find(&user).Error

	if err != nil {
		return "", "", err
	}

	return user.Password, user.ID, nil
}

func (ar *authorizationRepository) GetUserRoles(userID string) ([]int, error) {
	result, _ := ar.db.Raw("SELECT role_id FROM \"user_to_roles\"  WHERE (user_id = ?)\n", userID).Rows()
	if result.Err() != nil {
		return nil, result.Err()
	}

	var rolesID []int
	for result.Next() {
		var role string
		err := result.Scan(&role)
		if err != nil {
			return nil, err
		}
		roleID, _ := strconv.Atoi(role)
		rolesID = append(rolesID, roleID)
	}

	return rolesID, nil
}

func (ar *authorizationRepository) GetRefreshTokenUUIDFromTable(uuid string) (string, error) {
	var s model.Session
	err := ar.db.Model(&s).Select("refresh_token_uuid").Where("id = ?", uuid).Find(&s)
	if err.Error != nil {
		return "", err.Error
	}

	return s.RefreshTokenUUID, nil
}

func (ar *authorizationRepository) GetSessionUUID(userID string) (string, error) {
	var s model.Session
	err := ar.db.Model(&s).Where("user_id = ?", userID).Find(&s)
	if err.Error != nil {
		return "", err.Error
	}

	return s.ID, nil
}

func (ar *authorizationRepository) GetSessionExpiration(sessionUUID string) (time.Time, error) {
	var s model.Session
	err := ar.db.Model(&s).Where("id = ?", sessionUUID).Find(&s)
	if err.Error != nil {
		return time.Time{}, err.Error
	}

	return s.ExpiresAt, nil
}

func (ar *authorizationRepository) UpdateSession(sessionUUID string, s *model.Session) error {
	err := ar.db.Model(&s).Where("id = ?", sessionUUID).Update(s)
	if err.Error != nil {
		return err.Error
	}

	return nil
}

func (ar *authorizationRepository) Logout(sessionUUID string) error {
	var s model.Session

	err := ar.db.Model(&s).Where("id = ?", sessionUUID).Delete(s)
	if err.Error != nil {
		return err.Error
	}

	return nil
}

func (ar *authorizationRepository) CreateUserToRole(userRole model.UserRole) error {
	err := ar.db.Model(&userRole).Create(&userRole).Error

	if err != nil {
		return err
	}

	return nil
}
