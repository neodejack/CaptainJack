# CaptainJack
Captain Jack is a telegram bot built for NTU tennis club members with the
intent to streamline day-to-day club operation, so that it's a smoother experience for the members

### Code Architecture Hexagonal Architecture
I'm setting up the project with a hexagonal architecture, so that my core functions (domains)
are separated from the not really relevant technical details

![hex_arch](https://user-images.githubusercontent.com/75655215/188204457-90452ce0-921a-46eb-a82d-d7cd844fbeb8.png)


To list down some examples of what i'm trying to achieve with this architecture:

1. Right now my choice of user interface (driving framework) is an interactive telegram bot, in the future 
i can change this to a web app without having to 
rewrite my core and application code again.

2. My choice of database for the first iteration is literally just google sheets, cuz that's how our club operates,
and that's where all the data is for now. 
In the future i will user a proper MySQL, 

and the cool thing is that i can just "unplug" the adapter of my current google sheets database 
and just "plug in" my MySQL database adapter under the driven framework part.

## Quick Start

Clone the repo of course

And make sure you installed `go` and already set up `GOPATH`

```
go run cmd/main.go
```

### The current progress

Working on the endpoint that give users court allocations

Currently finished the telegram bot ui adapter prototyping


i also built an inlinekeyboard selection just to see what input i can get from telebot ui

you can activate it by sending `courtallocation` to the bot
### Roadmap
check the Jira board

### CI/CD
Continuous Integration is managed by github actions, it's just basically running go unit tests to see if it passes the important checks

Deployment is managed by Railway.app

main branch is deploy to production environment on Railway, it's tele handle is `@tennis_club_bot`

beta branch is deployed to beta environment on Railway, its tele handle is `@beta_tennis_club_bot`

### Contributing
If you wanna contribute, please be my guest.

But please closely follow the steps stated here:

1. Select a ticket from my jira board that you wanna contribute, assign it to yourself on jira
2. Create a branch from Jira (i have done the integration so that the branch name will be linked to jira ticket)
3. Code it out and test it locally, if you can please write well thought unit tests for your code
4. Create a pull request to `beta` branch (DO NOT go straight to the `main` branch here)



