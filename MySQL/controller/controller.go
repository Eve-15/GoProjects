package controller

import "github.com/Eve-15/GoProjects/MySQL/model"

type GameController struct {
	Player model.Player
}

type UserController struct {
	UserRepo model.UserRepository
}

func NewGameController(userInfo model.Player) *GameController {
	return &GameController{
		Player: userInfo,
	}
}

func (uc *UserController) CreatUser(user *model.User) error {
	return uc.UserRepo.Create(user)
}

func (uc *UserController) GetUserByID(id string) (*model.User, error) {
	return uc.UserRepo.GetByID(id)
}

func (uc *UserController) UpdateUser(user *model.User) error {
	return uc.UserRepo.Update(user)
}

func (uc *UserController) DeleteUser(id string) error {
	return uc.UserRepo.Delete(id)
}

func (gc *GameController) Xi() {
	gc.Player.Xipai()
}

func (gc *GameController) Pai() {
	gc.Player.Paixu()
}

func (gc *GameController) UserScore() int {
	return gc.Player.Score()
}

func (gc *GameController) Get() []string {
	return gc.Player.GetCard()
}
