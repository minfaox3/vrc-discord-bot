package vrchatapi

import "time"

type FriendUser struct {
	ID                             string    `json:"id"`
	DisplayName                    string    `json:"displayName"`
	BIO                            string    `json:"bio"`
	BIOLinks                       []string  `json:"bioLinks"`
	DeveloperType                  string    `json:"developerType"`
	CurrentAvatarImageUrl          string    `json:"currentAvatarImageUrl"`
	CurrentAvatarThumbnailImageUrl string    `json:"currentAvatarThumbnailImageUrl"`
	CurrentAvatarTags              []string  `json:"currentAvatarTags"`
	UserIcon                       string    `json:"userIcon"`
	ProfilePicOverride             string    `json:"profilePicOverride"`
	ImageURL                       string    `json:"imageUrl"`
	LastLogin                      time.Time `json:"last_login"`
	LastMobile                     string    `json:"last_mobile"`
	Status                         string    `json:"status"`
	StatusDescription              string    `json:"statusDescription"`
	LastPlatform                   string    `json:"last_platform"`
	Location                       string    `json:"location"`
	Tags                           []string  `json:"tags"`
	FriendKey                      string    `json:"friendKey"`
	IsFriend                       bool      `json:"isFriend"`
}

type CurrentUser struct {
	ID                             string        `json:"id"`
	DisplayName                    string        `json:"displayName"`
	UserIcon                       string        `json:"userIcon"`
	BIO                            string        `json:"bio"`
	BIOLinks                       []string      `json:"bioLinks"`
	ProfilePicOverride             string        `json:"profilePicOverride"`
	StatusDescription              string        `json:"statusDescription"`
	Badges                         []string      `json:"badges"`
	UserName                       string        `json:"username"`
	PastDisplayNames               []string      `json:"pastDisplayNames"`
	HasEmail                       bool          `json:"hasEmail"`
	HasPendingEmail                bool          `json:"hasPendingEmail"`
	ObfuscatedEmail                string        `json:"obfuscatedEmail"`
	ObfuscatedPendingEmail         string        `json:"obfuscatedPendingEmail"`
	EmailVerified                  bool          `json:"emailVerified"`
	HasBirthDay                    bool          `json:"hasBirthDay"`
	HideContentFilterSettings      bool          `json:"hideContentFilterSettings"`
	Unsubscribe                    bool          `json:"unsubscribe"`
	StatusHistory                  []string      `json:"statusHistory"`
	StatusFirstTime                bool          `json:"statusFirstTime"`
	Friends                        []string      `json:"friends"`
	FriendGroupNames               []string      `json:"friendGroupNames"`
	UserLanguage                   string        `json:"userLanguage"`
	UserLanguageCode               string        `json:"userLanguageCode"`
	CurrentAvatarImageUrl          string        `json:"currentAvatarImageUrl"`
	CurrentAvatarThumbnailImageUrl string        `json:"currentAvatarThumbnailImageUrl"`
	CurrentAvatarTags              []string      `json:"currentAvatarTags"`
	CurrentAvatar                  string        `json:"currentAvatar"`
	CurrentAvatarAssetUrl          string        `json:"currentAvatarAssetUrl"`
	FallbackAvatar                 string        `json:"fallbackAvatar"`
	AccountDeletionDate            time.Time     `json:"accountDeletionDate"`
	AccountDeletionLog             string        `json:"accountDeletionLog"`
	AcceptedTOSVersion             int           `json:"acceptedTOSVersion"`
	AcceptedPrivacyVersion         int           `json:"acceptedPrivacyVersion"`
	SteamID                        string        `json:"steamId"`
	SteamDetails                   SteamDetails  `json:"steamDetails"`
	GoogleID                       string        `json:"googleId"`
	GoogleDetails                  GoogleDetails `json:"googleDetails"`
	OculusID                       string        `json:"oculusId"`
	PicoID                         string        `json:"picoId"`
	ViveID                         string        `json:"viveId"`
	HasLoggedInFromClient          bool          `json:"hasLoggedInFromClient"`
	HomeLocation                   string        `json:"homeLocation"`
	TwoFactorAuthEnabled           bool          `json:"twoFactorAuthEnabled"`
	TwoFactorAuthEnabledDate       time.Time     `json:"twoFactorAuthEnabledDate"`
	UpdatedAt                      time.Time     `json:"updated_at"`
	State                          string        `json:"state"`
	LastMobile                     string        `json:"last_mobile"`
	Tags                           []string      `json:"tags"`
	DeveloperType                  string        `json:"developerType"`
	LastLogin                      time.Time     `json:"last_login"`
	LastPlatform                   string        `json:"last_platform"`
	AllowAvatarCopying             bool          `json:"allowAvatarCopying"`
	Status                         string        `json:"status"`
	DateJoined                     string        `json:"date_joined"`
	IsFriend                       bool          `json:"isFriend"`
	FriendKey                      string        `json:"friendKey"`
	LastActivity                   time.Time     `json:"last_activity"`
	OnlineFriends                  []string      `json:"onlineFriends"`
	ActiveFriends                  []string      `json:"activeFriends"`
	Presence                       Presence      `json:"presence"`
	OfflineFriends                 []string      `json:"offlineFriends"`
}
