package conf

type BookDesc struct {
	BookName	string
	BookNumber	string
	BookState	bool
	BookIntro	string
}

type ReqPayload struct {
	BookName	string
	BookNumber	string
}
// func init() {
// 	ActionType := actionType{1,2,3,4}
// }