package registry

import (
	"GoAds/interface/controller"
)

func (r *registry) NewChatController() controller.ChatController {
	return controller.NewChatController(r.NewUserInterfactor())
}
