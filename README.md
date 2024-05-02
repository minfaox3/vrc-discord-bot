# vrc-discord-bot

![](VRC-SOffFriend.gif)

## Description
This application is a bot that allows you to view your VRChat friends' status and profiles as cards on the discord.  
Planning to add more content to cards, world and group related features in the future.

このアプリケーションはディスコード上でVRChatのフレンドのステータスやプロフィールをカードで確認できるbotです。  
今後カード内容の充実,ワールドとグループ関連の機能も追加する予定です。  
[README日本語(Japanese)](README(Japanese).md)←

## Getting Started
### How to startup
0. Download this repository.
1. Prepare the .env file directly under the repository. The contents to be included are as follows
```
BOT_TOKEN=Token(Discord bot)
NAME=username(VRChat)
PASS=password(VRChat)
USER_AGENT=application name/version email-address
```
`Example of USER_AGENT`：`vrc-discord-bot/1.0 mail@example.com`
2. Run the application (`go run main.go`) directly under the repository.
3. A login attempt will be made, requiring the entry of a code if two-factor authentication is enabled.
```
Get environment data...
Try login to VRChat...
Enter 2FA code:
```
4. If the login is successful, the following will be displayed. Startup is complete.
```
Successfully logged in
Bot is now running.  Press CTRL-C to exit.
```

### How to use
There are currently three commands supported.  
※Planning to change the bot to be called with a slash command instead of a prefix.
* `!vrcbot ShowCurrentUser`
    * Display information about yourself.
      ![](VRC.gif)
* `!vrcbot ShowOnlineFriend`
    * Select the name of the friend whose online friends you want to see more information about, and the cards of that friend will be displayed.
* `!vrcbot ShowOfflineFriend`
    * Select the name of the friend whose offline friends you want to see more information about, and the cards of that friend will be displayed.

## Plan
* Addition of content to be displayed on the card
* Add support for two-factor authentication other than email
* Changed to call bots with slash commands instead of prefixes
* Add　various APIs for VRChat
* Add feature related to World
* Add feature related to Group

## LICENSE
[MIT](LICENSE)

## Notice
Any country, region, organization, or individual must comply with the laws.