package repository

import (
	"GoAds/domain"
	"GoAds/domain/model"
	"GoAds/usecase/repository"
	"github.com/jinzhu/gorm"
	"strconv"
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

func (ar *authorizationRepository) UserExists(email string) (string, string, error) {
	user := model.User{}
	err := ar.db.Model(&user).Select("*").Where("email = ?", email).Find(&user).Error

	if err != nil {
		return "", "", err
	}

	return user.Password, user.ID, nil
}

func (ar *authorizationRepository) GetUserRoles(userID string) ([]int, error) {
	result, _ := ar.db.Raw("select role_id FROM \"user_to_roles\"  WHERE (user_id = ?)\n", userID).Rows()
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

func (ar *authorizationRepository) Login(u []*model.User) ([]*model.User, error) {
	return nil, nil
}
