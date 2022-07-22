package registry

import (
	"GoAds/interface/controller"
	ir "GoAds/interface/repository"
	"GoAds/usecase/interfactor"
	ur "GoAds/usecase/repository"
)

func (r *registry) NewCommentController() controller.CommentController {
	return controller.NewCommentController(r.NewCommentInterfactor())
}

func (r *registry) NewCommentInterfactor() interfactor.CommentInterfactor {
	return interfactor.NewCommentInterfactor(r.NewCommentRepository())
}

func (r *registry) NewCommentRepository() ur.CommentRepository {
	return ir.NewCommentRepository(r.db)
}
