package commands

import (
	"github.com/celestix/gotgproto/dispatcher"
)

func (m *command) LoadStart(dispatcher dispatcher.Dispatcher) {
	log := m.log.Named("start")
	defer log.Sugar().Info("Loaded")
	// dispatcher.AddHandler(handlers.NewCommand("start", start))
}

// func start(ctx *ext.Context, u *ext.Update) error {
// 	chatId := u.EffectiveChat().GetID()
// 	peerChatId := ctx.PeerStorage.GetPeerById(chatId)
// 	if peerChatId.Type != int(storage.TypeUser) {
// 		return dispatcher.EndGroups
// 	}
// 	if len(config.ValueOf.AllowedUsers) != 0 && !utils.Contains(config.ValueOf.AllowedUsers, chatId) {
// 		ctx.Reply(u, "You are not allowed to use this bot.", nil)
// 		return dispatcher.EndGroups
// 	}
// 	ctx.Reply(u, "Hi, send me any file to get a direct streamble link to that file.", nil)
// 	return dispatcher.EndGroups
// }
