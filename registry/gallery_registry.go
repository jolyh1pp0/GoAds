package registry

import (
	"GoAds/interface/controller"
	ir "GoAds/interface/repository"
	"GoAds/usecase/interfactor"
	ur "GoAds/usecase/repository"
)

func (r *registry) NewGalleryController() controller.GalleryController {
	return controller.NewGalleryController(r.NewGalleryInterfactor())
}

func (r *registry) NewGalleryInterfactor() interfactor.GalleryInterfactor {
	return interfactor.NewGalleryInterfactor(r.NewGalleryRepository())
}

func (r *registry) NewGalleryRepository() ur.GalleryRepository {
	return ir.NewGalleryRepository(r.db)
}