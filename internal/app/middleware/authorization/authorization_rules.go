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
	// ics_file_manager
	//
	"/ics_file_manager.Service/CreateICSFile": {
		Public: struct{}{},
	},
	"/ics_file_manager.Service/DeleteICSFile": {
		Public: struct{}{},
	},
	"/ics_file_manager.Service/GetICSFile": {
		Public: struct{}{},
	},
	"/ics_file_manager.Service/GetICSFileByGameID": {
		Public: struct{}{},
	},
	"/ics_file_manager.Service/ListICSFiles": {
		Public: struct{}{},
	},
	//
	// users
	//
	"/users.CroupierService/GetLotteryStatus": {
		Public: struct{}{},
	},
	"/users.CroupierService/RegisterForLottery": {
		User: struct{}{},
	},
	"/users.PhotographerService/AddGamePhotos": {
		Management: struct{}{},
	},
	"/users.PhotographerService/GetGamesWithPhotos": {
		Public: struct{}{},
	},
	"/users.PhotographerService/GetPhotosByGameID": {
		Public: struct{}{},
	},
	"/users.RegistratorService/AddGame": {
		Management: struct{}{},
	},
	"/users.RegistratorService/AddGames": {
		Public: struct{}{},
	},
	"/users.RegistratorService/CreateUser": {
		Public: struct{}{},
	},
	"/users.RegistratorService/DeleteGame": {
		Management: struct{}{},
	},
	"/users.RegistratorService/GetGameByID": {
		Public: struct{}{},
	},
	"/users.RegistratorService/GetGames": {
		Public: struct{}{},
	},
	"/users.RegistratorService/GetLeagueByID": {
		Public: struct{}{},
	},
	"/users.RegistratorService/GetPlaceByID": {
		Public: struct{}{},
	},
	"/users.RegistratorService/GetPlaceByNameAndAddress": {
		Public: struct{}{},
	},
	"/users.RegistratorService/GetPlayersByGameID": {
		Public: struct{}{},
	},
	"/users.RegistratorService/GetRegisteredGames": {
		Public: struct{}{},
	},
	"/users.RegistratorService/GetUser": {
		User: struct{}{},
	},
	"/users.RegistratorService/GetUserByID": {
		Public: struct{}{},
	},
	"/users.RegistratorService/GetUserByTelegramID": {
		Public: struct{}{},
	},
	"/users.RegistratorService/GetUserGames": {
		User: struct{}{},
	},
	"/users.RegistratorService/RegisterGame": {
		User: struct{}{},
	},
	"/users.RegistratorService/RegisterPlayer": {
		User: struct{}{},
	},
	"/users.RegistratorService/UnregisterGame": {
		User: struct{}{},
	},
	"/users.RegistratorService/UnregisterPlayer": {
		User: struct{}{},
	},
	"/users.RegistratorService/UpdateUserEmail": {
		User: struct{}{},
	},
	"/users.RegistratorService/UpdateUserName": {
		User: struct{}{},
	},
	"/users.RegistratorService/UpdateUserPhone": {
		User: struct{}{},
	},
	"/users.RegistratorService/UpdateUserState": {
		User: struct{}{},
	},
	"/users.RegistratorService/UpdatePayment": {
		User: struct{}{},
	},
}
