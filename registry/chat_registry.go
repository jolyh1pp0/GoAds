package registry

import (
	"GoAds/interface/controller"
	ir "GoAds/interface/repository"
	"GoAds/usecase/interfactor"
	ur "GoAds/usecase/repository"
)

func (r *registry) NewChatController() controller.ChatController {
	return controller.NewChatController(r.NewChatInterfactor())
}

func (r *registry) NewChatInterfactor() interfactor.ChatInterfactor {
	return interfactor.NewChatInterfactor(r.NewChatRepository())
}

func (r *registry) NewChatRepository() ur.ChatRepository {
	return ir.NewChatRepository(r.db)
}