package registry

import (
	"GoAds/interface/controller"
	ip "GoAds/interface/presenter"
	ir "GoAds/interface/repository"
	"GoAds/usecase/interfactor"
	up "GoAds/usecase/presenter"
	ur "GoAds/usecase/repository"
)

func (r *registry) NewAdvertisementController() controller.AdvertisementController {
	return controller.NewAdvertisementController(r.NewAdvertisementInterfactor())
}

func (r *registry) NewAdvertisementInterfactor() interfactor.AdvertisementInterfactor {
	return interfactor.NewAdvertisementInterfactor(r.NewAdvertisementRepository(), r.NewAdvertisementPresenter())
}

func (r *registry) NewAdvertisementRepository() ur.AdvertisementRepository {
	return ir.NewAdvertisementRepository(r.db)
}

func (r *registry) NewAdvertisementPresenter() up.AdvertisementPresenter {
	return ip.NewAdvertisementPresenter()
}
