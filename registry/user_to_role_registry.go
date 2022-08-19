package registry

import (
	"GoAds/interface/controller"
	ir "GoAds/interface/repository"
	"GoAds/usecase/interfactor"
	ur "GoAds/usecase/repository"
)

func (r *registry) NewUserToRoleController() controller.UserToRoleController {
	return controller.NewUserToRoleController(r.NewUserToRoleInterfactor())
}

func (r *registry) NewUserToRoleInterfactor() interfactor.UserToRoleInterfactor {
	return interfactor.NewUserToRoleInterfactor(r.NewUserToRoleRepository())
}

func (r *registry) NewUserToRoleRepository() ur.UserToRoleRepository {
	return ir.NewUserToRoleRepository(r.db)
}
