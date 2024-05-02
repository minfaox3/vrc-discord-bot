package vrchatapi

type Need2FA struct {
	RequiresTwoFactorAuth []string `json:"requiresTwoFactorAuth"`
}

type Code2FA struct {
	Code string `json:"code"`
}

type SteamDetails struct {
	//TODO
}

type GoogleDetails struct {
	//TODO
}

type Presence struct {
	ID                  string   `json:"id"`
	Platform            string   `json:"platform"`
	Status              string   `json:"status"`
	World               string   `json:"world"`
	Instance            string   `json:"instance"`
	InstanceType        string   `json:"instanceType"`
	TravelingToWorld    string   `json:"travelingToWorld"`
	TravelingToInstance string   `json:"travelingToInstance"`
	Groups              []string `json:"groups"`
}
