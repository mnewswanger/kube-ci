package cmd

type gitlabHookRequestStruct struct {
	Added     string `json:"added"`
	Before    string `json:"before"`
	EventName string `json:"event_name"`
	Ref       string `json:"ref"`
	Project   struct {
		AvatarURL string `json:"avatar_url"`
		Name      string `json:"name"`
		Namespace string `json:"namespace"`
	} `json:"project"`
	UserAvatar string `json:"user_avatar"`
	UserEmail  string `json:"user_email"`
	UserName   string `json:"user_name"`
}
