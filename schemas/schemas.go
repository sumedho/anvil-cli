package schemas

type Configuration struct {
	UserName string `json:"Username"`
	BaseUrl  string `json:"Baseurl"`
}

type Auth struct {
	Username string `json:"username"`
	Apikey   string `json:"apiKey"`
}

type Token struct {
	Token  string `json:"token"`
	Expiry string `json:"expiry"`
}
