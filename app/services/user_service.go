package services

import (
	"errors"
	"golang_line/app/repositories"
)

type userService struct {
}

func UserService() *userService {
	return &userService{}
}

func (h *userService) CheckUserExist(userID string) (bool, error) {
	if messages, err := repositories.MessageRepository().GetMessagesbyUserID(userID); err != nil {
		return false, err
	} else {
		if len(messages) != 0 {
			return true, nil
		} else {
			return false, errors.New("UserID 不存在")
		}
	}
}
