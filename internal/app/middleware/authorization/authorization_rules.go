package authorization

type roles map[string]struct{}

const (
	// Admin ...
	Admin = "admin"
	// Management ...
	Management = "management"
	// Public ...
	Public = "public"
	// S2S ...
	S2S = "s2s"
	// User ...
	User = "user"
)

var grpcRules = map[string]roles{
	//
	// admin
	//
	"/admin.Service/CreateUserRole": {
		Admin: struct{}{},
	},
	"/admin.Service/DeleteUserRole": {
		Admin: struct{}{},
	},
	"/admin.Service/GetUserRolesByUserID": {
		Admin: struct{}{},
	},
	"/admin.Service/ListUserRoles": {
		Admin: struct{}{},
	},
	//
	// certificate_manager
	//
	"/certificate_manager.Service/CreateCertificate": {
		Management: struct{}{},
	},
	"/certificate_manager.Service/DeleteCertificate": {
		Management: struct{}{},
	},
	"/certificate_manager.Service/ListCertificates": {
		Public: struct{}{},
	},
	"/certificate_manager.Service/PatchCertificate": {
		Management: struct{}{},
	},
	//
	// croupier
	//
	"/croupier.Service/GetLotteryStatus": {
		Public: struct{}{},
	},
	"/croupier.Service/RegisterForLottery": {
		User: struct{}{},
	},
	//
	// game
	//
	"/game.Service/BatchGetGames": {
		Public: struct{}{},
	},
	"/game.Service/CreateGame": {
		Management: struct{}{},
		S2S:        struct{}{},
	},
	"/game.Service/DeleteGame": {
		Management: struct{}{},
	},
	"/game.Service/GetGame": {
		Public: struct{}{},
	},
	"/game.Service/ListGames": {
		Public: struct{}{},
	},
	"/game.Service/PatchGame": {
		Management: struct{}{},
		S2S:        struct{}{},
	},
	"/game.Service/SearchGamesByLeagueID": {
		Public: struct{}{},
		S2S:    struct{}{},
	},
	"/game.Service/SearchPassedAndRegisteredGames": {
		Public: struct{}{},
	},
	"/game.RegistratorService/RegisterGame": {
		User: struct{}{},
	},
	"/game.RegistratorService/UnregisterGame": {
		User: struct{}{},
	},
	"/game.RegistratorService/UpdatePayment": {
		User: struct{}{},
	},
	//
	// game_player
	//
	"/game_player.Service/CreateGamePlayer": {
		User: struct{}{},
	},
	"/game_player.Service/DeleteGamePlayer": {
		User: struct{}{},
	},
	"/game_player.Service/GetGamePlayer": {
		Public: struct{}{},
	},
	"/game_player.Service/GetGamePlayersByGameID": {
		Public: struct{}{},
	},
	"/game_player.Service/GetUserGameIDs": {
		Public: struct{}{},
	},
	"/game_player.Service/PatchGamePlayer": {
		User: struct{}{},
	},
	"/game_player.RegistratorService/RegisterPlayer": {
		User: struct{}{},
	},
	"/game_player.RegistratorService/UnregisterPlayer": {
		User: struct{}{},
	},
	//
	// game_result_manager
	//
	"/game_result_manager.Service/CreateGameResult": {
		Management: struct{}{},
	},
	"/game_result_manager.Service/ListGameResults": {
		Public: struct{}{},
	},
	"/game_result_manager.Service/PatchGameResult": {
		Management: struct{}{},
	},
	"/game_result_manager.Service/SearchGameResultByGameID": {
		Public: struct{}{},
	},
	//
	// league
	//
	"/league.Service/GetLeague": {
		Public: struct{}{},
	},
	//
	// photo_manager
	//
	"/photo_manager.Service/AddGamePhotos": {
		Management: struct{}{},
	},
	"/photo_manager.Service/GetPhotosByGameID": {
		Public: struct{}{},
	},
	//
	// place
	//
	"/place.Service/GetPlace": {
		Public: struct{}{},
	},
	//
	// user_manager
	//
	"/user_manager.Service/CreateUser": {
		Public: struct{}{},
	},
	"/user_manager.Service/GetUser": {
		Public: struct{}{},
	},
	"/user_manager.Service/GetUserByTelegramID": {
		Public: struct{}{},
	},
	"/user_manager.Service/PatchUser": {
		Public: struct{}{},
	},
}
