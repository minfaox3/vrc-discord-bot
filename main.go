package main

import (
	"fmt"
	"github.com/bwmarrin/discordgo"
	"github.com/joho/godotenv"
	"log"
	"net/http"
	"os"
	"os/signal"
	"strconv"
	"strings"
	"syscall"
	"time"
	"vrc-discord-bot/vrchatapi"
)

// global variables for discordgo
const prefix string = "!vrcbot"
const (
	cid_selection_friend_offline = "selection_friend_offline"
	cid_selection_friend_online  = "selection_friend_online"
)

var (
	token string
	sess  *discordgo.Session
)

// global variables for vrchatapi
var (
	username  string
	password  string
	userAgent string
	cookie    []*http.Cookie
)

// global variables for store data
var (
	currentUser            vrchatapi.CurrentUser
	friendOffline          []vrchatapi.FriendUser
	selectionFriendOffline []discordgo.SelectMenuOption
	friendOnline           []vrchatapi.FriendUser
	selectionFriendOnline  []discordgo.SelectMenuOption
)

func updateMessageComplex(channelID, messageID string, components []discordgo.MessageComponent) {
	message := &discordgo.MessageEdit{
		ID:      messageID,
		Channel: channelID,
		Components: &[]discordgo.MessageComponent{
			discordgo.ActionsRow{
				Components: components,
			},
		},
	}
	_, err := sess.ChannelMessageEditComplex(message)
	if err != nil {
		fmt.Println(err)
		return
	}
}

func showFriendOffline(channelID, messageID string) {
	var err error
	friendOffline, err = vrchatapi.FriendOffline(userAgent, cookie)
	if err != nil {
		return
	}

	selectionFriendOffline = make([]discordgo.SelectMenuOption, 0, len(friendOffline))
	for index, friend := range friendOffline {
		selectionFriendOffline = append(selectionFriendOffline, discordgo.SelectMenuOption{
			Label: friend.DisplayName,
			Value: strconv.Itoa(index),
		})
	}

	components := []discordgo.MessageComponent{
		&discordgo.SelectMenu{
			CustomID: cid_selection_friend_offline,
			Options:  selectionFriendOffline,
		},
	}
	_, err = sess.ChannelMessageEdit(channelID, messageID, "Please select friend")
	if err != nil {
		fmt.Println(err)
		return
	}
	updateMessageComplex(channelID, messageID, components)
}

func showFriendOnline(channelID, messageID string) {
	var err error
	friendOnline, err = vrchatapi.FriendOnline(userAgent, cookie)
	if err != nil {
		return
	}

	selectionFriendOnline = make([]discordgo.SelectMenuOption, 0, len(friendOnline))
	for index, friend := range friendOnline {
		selectionFriendOnline = append(selectionFriendOnline, discordgo.SelectMenuOption{
			Label: friend.DisplayName,
			Value: strconv.Itoa(index),
		})
	}

	components := []discordgo.MessageComponent{
		&discordgo.SelectMenu{
			CustomID: cid_selection_friend_online,
			Options:  selectionFriendOnline,
		},
	}
	_, err = sess.ChannelMessageEdit(channelID, messageID, "Please select friend")
	if err != nil {
		fmt.Println(err)
		return
	}
	updateMessageComplex(channelID, messageID, components)
}

func createFriendUserEmbedMessage(channelID, messageID string, friend vrchatapi.FriendUser) {
	var err error
	fields := []*discordgo.MessageEmbedField{
		{
			Name:  "Status",
			Value: friend.Status,
		},
		{
			Name:  "Location",
			Value: friend.Location,
		},
	}
	color := 0xcccccc
	rank := "VISITOR"
	for _, tag := range friend.Tags {
		//TODO make tag or prefix of tag strings to const variables
		if strings.Contains(tag, "language_") {
			fields = append(fields, &discordgo.MessageEmbedField{
				Name:  "Language",
				Value: strings.ToUpper(tag[len("language_"):]),
			})
		}
		if tag == "system_trust_veteran" {
			color = 0x8143e6
			rank = "Trusted User"
		}
		if tag == "system_trust_trusted" {
			if rank != "Trusted User" {
				color = 0xff7b42
				rank = "Known User"
			}
		}
		if tag == "system_trust_known" {
			if rank != "Trusted User" {
				if rank != "Known User" {
					color = 0x2bcf5c
					rank = "User"
				}
			}
		}
		if tag == "system_trust_basic" {
			if rank != "Trusted User" {
				if rank != "Known User" {
					if rank != "User" {
						color = 0x1674f7
						rank = "New User"
					}
				}
			}
		}
	}
	fields = append(fields, &discordgo.MessageEmbedField{
		Name:  "Rank",
		Value: rank,
	})

	embed := &discordgo.MessageEmbed{
		Color: color,
		Author: &discordgo.MessageEmbedAuthor{
			Name: friend.ID,
		},
		Title: friend.DisplayName,
		Image: &discordgo.MessageEmbedImage{
			URL: friend.CurrentAvatarThumbnailImageUrl,
		},
		Description: friend.BIO,
		Fields:      fields,
		Thumbnail: &discordgo.MessageEmbedThumbnail{
			URL: friend.ImageURL,
		},
		Footer: &discordgo.MessageEmbedFooter{
			Text: "Last login:",
		},
		Timestamp: friend.LastLogin.Format(time.RFC3339),
	}
	_, err = sess.ChannelMessageEditEmbed(channelID, messageID, embed)
	if err != nil {
		fmt.Println(err)
		return
	}
}

func showCurrentUser(channelID string) {
	var err error
	var result bool
	result, currentUser, cookie, err = vrchatapi.CurrentUserData(username, password, userAgent, cookie)
	if err != nil {
		return
	}
	if result == false {
		return
	}
	createCurrentUserEmbedMessage(channelID)
}

func createCurrentUserEmbedMessage(channelID string) {
	var err error
	fields := []*discordgo.MessageEmbedField{
		{
			Name:  "Status",
			Value: currentUser.Status,
		},
		{
			Name:  "Home Location",
			Value: currentUser.HomeLocation,
		},
	}
	color := 0xcccccc
	rank := "VISITOR"
	for _, tag := range currentUser.Tags {
		//TODO make tag or prefix of tag strings to const variables
		if strings.Contains(tag, "language_") {
			fields = append(fields, &discordgo.MessageEmbedField{
				Name:  "Language",
				Value: strings.ToUpper(tag[len("language_"):]),
			})
		}
		if tag == "system_trust_veteran" {
			color = 0x8143e6
			rank = "Trusted User"
		}
		if tag == "system_trust_trusted" {
			if rank != "Trusted User" {
				color = 0xff7b42
				rank = "Known User"
			}
		}
		if tag == "system_trust_known" {
			if rank != "Trusted User" {
				if rank != "Known User" {
					color = 0x2bcf5c
					rank = "User"
				}
			}
		}
		if tag == "system_trust_basic" {
			if rank != "Trusted User" {
				if rank != "Known User" {
					if rank != "User" {
						color = 0x1674f7
						rank = "New User"
					}
				}
			}
		}
	}
	fields = append(fields, &discordgo.MessageEmbedField{
		Name:  "RANK",
		Value: rank,
	})

	embed := &discordgo.MessageEmbed{
		Color: color,
		Author: &discordgo.MessageEmbedAuthor{
			Name: currentUser.ID,
		},
		Title: currentUser.DisplayName,
		Image: &discordgo.MessageEmbedImage{
			URL: currentUser.CurrentAvatarThumbnailImageUrl,
		},
		Description: currentUser.BIO,
		Fields:      fields,
		Thumbnail: &discordgo.MessageEmbedThumbnail{
			URL: currentUser.CurrentAvatarImageUrl,
		},
		Footer: &discordgo.MessageEmbedFooter{
			Text: "Last login:",
		},
		Timestamp: currentUser.LastLogin.Format(time.RFC3339),
	}
	_, err = sess.ChannelMessageSendEmbed(channelID, embed)
	if err != nil {
		fmt.Println(err)
		return
	}
}

func handleMessageReaction(s *discordgo.Session, r *discordgo.MessageReactionAdd) {
	//For delete waste message
	//if r.Emoji.Name == "❌" {
	//	err := s.ChannelMessageDelete(r.ChannelID, r.MessageID)
	//	if err != nil {
	//		fmt.Println(err)
	//		return
	//	}
	//}
}

func handleMessage(s *discordgo.Session, m *discordgo.MessageCreate) {
	//do not react to own message
	if m.Author.ID == s.State.User.ID {
		return
	}

	//TODO change to slash command
	args := strings.Split(m.Content, " ")
	if args[0] != prefix {
		return
	}

	if args[1] == "ShowCurrentUser" {
		go showCurrentUser(m.ChannelID)
		return
	}

	if args[1] == "ShowOfflineFriend" {
		message := &discordgo.MessageSend{
			Content: "Please wait for get friend data",
			Components: []discordgo.MessageComponent{
				discordgo.ActionsRow{
					Components: []discordgo.MessageComponent{
						&discordgo.SelectMenu{
							Disabled: true,
							CustomID: cid_selection_friend_offline,
							Options: []discordgo.SelectMenuOption{
								{
									Label: "Wait",
									Value: "0",
								},
							},
						},
					},
				},
			},
		}
		tempMessage, err := s.ChannelMessageSendComplex(m.ChannelID, message)
		if err != nil {
			log.Println(err)
		}
		go showFriendOffline(tempMessage.ChannelID, tempMessage.ID)
		return
	}

	if args[1] == "ShowOnlineFriend" {
		message := &discordgo.MessageSend{
			Content: "Please wait for get friend data",
			Components: []discordgo.MessageComponent{
				discordgo.ActionsRow{
					Components: []discordgo.MessageComponent{
						&discordgo.SelectMenu{
							Disabled: true,
							CustomID: cid_selection_friend_online,
							Options: []discordgo.SelectMenuOption{
								{
									Label: "Wait",
									Value: "0",
								},
							},
						},
					},
				},
			},
		}
		tempMessage, err := s.ChannelMessageSendComplex(m.ChannelID, message)
		if err != nil {
			log.Println(err)
		}
		go showFriendOnline(tempMessage.ChannelID, tempMessage.ID)
		return
	}
}

func handleInteraction(s *discordgo.Session, i *discordgo.InteractionCreate) {
	if i.Type == discordgo.InteractionMessageComponent {
		//TODO
		err := s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
			/*Type: discordgo.InteractionResponseChannelMessageWithSource,
			Data: &discordgo.InteractionResponseData{
				Content: "THX! -> " + i.MessageComponentData().CustomID + " x " + i.MessageComponentData().Values[0],
			},*/
			Type: discordgo.InteractionResponseDeferredMessageUpdate,
			Data: &discordgo.InteractionResponseData{
				Content: "OK",
			},
		})
		if err != nil {
			log.Println(err)
		}

		index, err := strconv.Atoi(i.MessageComponentData().Values[0])
		if err != nil {
			//TODO output error on discord
			return
		}

		_, err = s.ChannelMessageEdit(i.ChannelID, i.Message.ID, "")
		if err != nil {
			fmt.Println(err)
		}

		if i.MessageComponentData().CustomID == cid_selection_friend_online {
			createFriendUserEmbedMessage(i.ChannelID, i.Message.ID, friendOnline[index])
		}

		if i.MessageComponentData().CustomID == cid_selection_friend_offline {
			createFriendUserEmbedMessage(i.ChannelID, i.Message.ID, friendOffline[index])
		}

		//TODO make expire interactable components
		message := &discordgo.MessageEdit{
			ID:      i.Message.ID,
			Channel: i.ChannelID,
			Components: &[]discordgo.MessageComponent{
				discordgo.ActionsRow{
					Components: []discordgo.MessageComponent{
						&discordgo.SelectMenu{
							Disabled: true,
							CustomID: "custom_selection",
							Options: []discordgo.SelectMenuOption{
								{
									Label: "Wait",
									Value: "0",
								},
							},
						},
					},
				},
			},
		}
		_, err = s.ChannelMessageEditComplex(message)
		if err != nil {
			fmt.Println(err)
		}
	}
}

func main() {
	fmt.Println("Get environment data...")
	err := godotenv.Load()
	if err != nil {
		log.Fatal(err)
		return
	}

	token = os.Getenv("BOT_TOKEN")
	userAgent = vrchatapi.CheckUserAgent(os.Getenv("USER_AGENT"))
	username = os.Getenv("NAME")
	password = os.Getenv("PASS")

	fmt.Println("Try login to VRChat...")

	loginResult, _, _, tempCookie, err := vrchatapi.Login(username, password, userAgent)
	if err != nil {
		log.Fatal(err)
	}
	cookie = tempCookie

	switch loginResult {
	case vrchatapi.Failed:
		log.Fatal(err)
	case vrchatapi.Succeeded:
		fmt.Println("Successfully logged in")
	case vrchatapi.Wait2FA:
		//TODO need to check 2fa type
		var code string
		fmt.Print("Enter 2FA code:")
		fmt.Scan(&code)
		result, tempCookie, err := vrchatapi.Mail2FA(code, userAgent, cookie)
		if err != nil {
			log.Fatal(err)
		}
		if !result {
			log.Fatal(err)
		}
		cookie = tempCookie
		fmt.Println("Successfully logged in")
	}

	defer func(userAgent string, cookie []*http.Cookie) {
		_, _, err := vrchatapi.Logout(userAgent, cookie)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println("Successfully logged out")
	}(userAgent, cookie)

	sess, err = discordgo.New("Bot " + os.Getenv("BOT_TOKEN"))
	if err != nil {
		log.Fatal(err)
	}

	sess.AddHandler(handleMessage)
	sess.AddHandler(handleInteraction)
	sess.AddHandler(handleMessageReaction)

	sess.Identify.Intents = discordgo.IntentsAllWithoutPrivileged

	err = sess.Open()
	if err != nil {
		log.Fatal(err)
	}
	defer sess.Close()

	fmt.Println("Bot is now running.  Press CTRL-C to exit.")

	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt)
	<-sc
}
