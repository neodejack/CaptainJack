# CaptainJack
Captain Jack is a telegram bot built for NTU tennis club members with the
intent to streamline day-to-day club operation, so that it's a smoother experience for the members

#### Hexagonal Architecture
I'm setting up the project with a hexagonal architecture, so that my core functions (domains)
are separated from the not really relevant technical details

![hex_arch](https://user-images.githubusercontent.com/75655215/188204457-90452ce0-921a-46eb-a82d-d7cd844fbeb8.png)


To list down some examples of what i'm trying to achieve with this architecture:

1. Right now my choice of user interface (driving framework) is an interactive telegram bot, in the future 
i can change this to a web app without having to 
rewrite my core and application code again.
2. My choice of database for the first iteration is Postgres.
But because all the data for the club matters are all stored in google sheets, i need to migrate the data and have the python script (most probably a lambda function) to migrate the data into database.

In the future i will fully use postgres and get rid of google sheets
and the cool thing is that i can just "unplug" the adapter of my current google sheets database 
and just "plug in" my PostgreSQL database adapter under the driven framework part.

## Quick Start

Clone the repo of course

And make sure you installed `go` and already set up `GOPATH`

```
go run cmd/main.go
```

### The current progress

Working on the endpoint that give users court allocations

Currently finished the telegram bot ui adapter prototyping


### What's Next
Datebase adapter
application adapter

