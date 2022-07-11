package registry

import (
	"GoAds/interface/controller"
	ir "GoAds/interface/repository"
	"GoAds/usecase/interfactor"
	ur "GoAds/usecase/repository"
)

func (r *registry) NewUserController() controller.UserController {
	return controller.NewUserController(r.NewUserInterfactor())
}

func (r *registry) NewUserInterfactor() interfactor.UserInterfactor {
	return interfactor.NewUserInterfactor(r.NewUserRepository())
}

func (r *registry) NewUserRepository() ur.UserRepository {
	return ir.NewUserRepository(r.db)
}
