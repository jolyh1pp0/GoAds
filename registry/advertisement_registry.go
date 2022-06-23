package registry

import (
	"GoAds/interface/controller"
	ir "GoAds/interface/repository"
	"GoAds/usecase/interfactor"
	ur "GoAds/usecase/repository"
)

func (r *registry) NewAdvertisementController() controller.AdvertisementController {
	return controller.NewAdvertisementController(r.NewAdvertisementInterfactor())
}

func (r *registry) NewAdvertisementInterfactor() interfactor.AdvertisementInterfactor {
	return interfactor.NewAdvertisementInterfactor(r.NewAdvertisementRepository())
}

func (r *registry) NewAdvertisementRepository() ur.AdvertisementRepository {
	return ir.NewAdvertisementRepository(r.db)
}
