# CaptainJack
Captain Jack is a telegram bot built for NTU tennis club members with the
intent to streamline day-to-day club operation, so that it's a smoother experience for the members

#### Hexagonal Architecture
I'm setting up the project with a hexagonal architecture, so that my core functions (domains)
are separated from the not really relevant technical details 

To list down some examples of what i'm trying to achieve with this architecture:

1. Right now my choice of user interface (driving framework) is an interactive telegram bot, in the future 
i can change this to a web app without having to 
rewrite my core and application code again.
2. My choice of database for the first iteration is literally just google sheets, cuz that's how our club operates,
and that's where all the data is for now. 
In the future i will user a proper PostgreSQL, 
and the cool thing is that i can just "unplug" the adapter of my current google sheets database 
and just "plug in" my PostgreSQL database adapter under the driven framework part.

## Quick Start

Clone the repo of course

And make sure you installed `go` and already set up `GOPATH`

```
go run cmd/main.go
```

### The current progress

haha just to make sure this architechture works i built a very obnoxious bot that replpy your message with "so what?"
at the end of it

You just freaking hate it right?

I kinda like it ngl

### What's Next
An endpoint to give users the court allocations
1. user are given option to select the date of the allocations that are available
2. user can click the inline option to select the date of allocations that they are interested in
3. user are return with a certain format of the court allocation data 

### Technology that i'm building this project with
1. gRPC