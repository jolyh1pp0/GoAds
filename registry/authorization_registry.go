package registry

import (
	"GoAds/interface/controller"
	ir "GoAds/interface/repository"
	"GoAds/usecase/interfactor"
	ur "GoAds/usecase/repository"
)

func (r *registry) NewAuthorizationController() controller.AuthorizationController {
	return controller.NewAuthorizationController(r.NewAuthorizationInterfactor())
}

func (r *registry) NewAuthorizationInterfactor() interfactor.AuthorizationInterfactor {
	return interfactor.NewAuthorizationInterfactor(r.NewAuthorizationRepository())
}

func (r *registry) NewAuthorizationRepository() ur.AuthorizationRepository {
	return ir.NewAuthorizationRepository(r.db)
}
