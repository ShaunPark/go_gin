package types

type InteractiveMessage struct {
	Type    string   `json:"type"`
	User    User     `json:"user"`
	Channel Channel  `json:"channel"`
	Actions []Action `json:"actions"`
}

type Action struct {
	ActionId string `json:"action_id"`
	Value    string `json:"value"`
	ActionTs string `json:"action_ts"`
}

type User struct {
	Id       string `json:"id"`
	UserName string `json:"username"`
	Name     string `json:"name"`
	TeamId   string `json:"team_id"`
}

type Channel struct {
	Id   string `json:"id"`
	Name string `json:"name"`
}
