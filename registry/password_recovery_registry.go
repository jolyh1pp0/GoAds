package registry

import (
	"GoAds/interface/controller"
	ir "GoAds/interface/repository"
	"GoAds/usecase/interfactor"
	ur "GoAds/usecase/repository"
)

func (r *registry) NewPasswordRecoveryController() controller.PasswordRecoveryController {
	return controller.NewPasswordRecoveryController(r.NewPasswordRecoveryInterfactor())
}

func (r *registry) NewPasswordRecoveryInterfactor() interfactor.PasswordRecoveryInterfactor {
	return interfactor.NewPasswordRecoveryInterfactor(r.NewPasswordRecoveryRepository())
}

func (r *registry) NewPasswordRecoveryRepository() ur.PasswordRecoveryRepository {
	return ir.NewPasswordRecoveryRepository(r.db)
}