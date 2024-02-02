## Discord-bot


### Description:
 Discord bot that implements new features to your sever


### Usage
Clone the repository:
```
git clone git@github.com:aidostt/discord-bot.git
```

Add this to your system's variables:
```
DISCORD_BOT_TOKEN: ask the owner for the token
OPENWEATHERMAP_API_KEY: 627a1ab7bcc88b03062200758ef06771
```

Run a program:
```
cd cmd/bot/
go run .
```
Invite bot to desired server using this link: 

```
https://discord.com/api/oauth2/authorize?client_id=1202865878354493441&permissions=8&scope=bot
```



### Documentation


Available command: 
```
    -help: see all commands
    -weather <city>: see the weather in desired location.
    -reminder <time> <message>: set reminder for the yourself
    -tictactoe <x> <y> or -tictactoe start: play tictactoe game with bot
```


