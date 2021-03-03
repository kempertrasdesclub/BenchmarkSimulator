package data

// Cache (Português): Dado a ser arquivado na cache
//   Caso necessite testar com outro formato de dado, basta mudar aqui.
type Cache struct {
	UserId         string `json:"user_id"`
	Status         string `json:"status"`
	Manual         bool   `json:"manual"`
	LastActivityAt int64  `json:"last_activity_at"`
	ActiveChannel  string `json:"-" db:"-"`
}
