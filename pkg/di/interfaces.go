package di

import "14-TestingAPI/internal/user"

type IStatRepository interface {
	AddClick(LinkId uint)
}

type IUserRepository interface {
	Create(user *user.User) (*user.User, error)
	FindByEmail(email string) (*user.User, error)
}
