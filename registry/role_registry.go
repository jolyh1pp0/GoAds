package registry

import (
	"GoAds/interface/controller"
	ir "GoAds/interface/repository"
	"GoAds/usecase/interfactor"
	ur "GoAds/usecase/repository"
)

func (r *registry) NewRoleController() controller.RoleController {
	return controller.NewRoleController(r.NewRoleInterfactor())
}

func (r *registry) NewRoleInterfactor() interfactor.RoleInterfactor {
	return interfactor.NewRoleInterfactor(r.NewRoleRepository())
}

func (r *registry) NewRoleRepository() ur.RoleRepository {
	return ir.NewRoleRepository(r.db)
}
