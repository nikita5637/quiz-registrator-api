package authorization

type roles map[string]struct{}

const (
	// Admin ...
	Admin = "admin"
	// Management ...
	Management = "management"
	// Public ...
	Public = "public"
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
	//
	// photo_manager
	//
	"/photo_manager.Service/AddGamePhotos": {
		Management: struct{}{},
	},
	"/photo_manager.Service/GetGamesWithPhotos": {
		Public: struct{}{},
	},
	"/photo_manager.Service/GetPhotosByGameID": {
		Public: struct{}{},
	},
	//
	// users
	//
	"/registrator.RegistratorService/AddGame": {
		Management: struct{}{},
	},
	"/registrator.RegistratorService/AddGames": {
		Public: struct{}{},
	},
	"/registrator.RegistratorService/CreateUser": {
		Public: struct{}{},
	},
	"/registrator.RegistratorService/DeleteGame": {
		Management: struct{}{},
	},
	"/registrator.RegistratorService/GetGameByID": {
		Public: struct{}{},
	},
	"/registrator.RegistratorService/GetGames": {
		Public: struct{}{},
	},
	"/registrator.RegistratorService/GetLeagueByID": {
		Public: struct{}{},
	},
	"/registrator.RegistratorService/GetPlaceByID": {
		Public: struct{}{},
	},
	"/registrator.RegistratorService/GetPlaceByNameAndAddress": {
		Public: struct{}{},
	},
	"/registrator.RegistratorService/GetPlayersByGameID": {
		Public: struct{}{},
	},
	"/registrator.RegistratorService/GetRegisteredGames": {
		Public: struct{}{},
	},
	"/registrator.RegistratorService/GetUser": {
		User: struct{}{},
	},
	"/registrator.RegistratorService/GetUserByID": {
		Public: struct{}{},
	},
	"/registrator.RegistratorService/GetUserByTelegramID": {
		Public: struct{}{},
	},
	"/registrator.RegistratorService/GetUserGames": {
		User: struct{}{},
	},
	"/registrator.RegistratorService/RegisterGame": {
		User: struct{}{},
	},
	"/registrator.RegistratorService/RegisterPlayer": {
		User: struct{}{},
	},
	"/registrator.RegistratorService/UnregisterGame": {
		User: struct{}{},
	},
	"/registrator.RegistratorService/UnregisterPlayer": {
		User: struct{}{},
	},
	"/registrator.RegistratorService/UpdateUserEmail": {
		User: struct{}{},
	},
	"/registrator.RegistratorService/UpdateUserName": {
		User: struct{}{},
	},
	"/registrator.RegistratorService/UpdateUserPhone": {
		User: struct{}{},
	},
	"/registrator.RegistratorService/UpdateUserState": {
		User: struct{}{},
	},
	"/registrator.RegistratorService/UpdatePayment": {
		User: struct{}{},
	},
}
