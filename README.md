# Article Rest API

## Introduction

This application is a rest API implementation of an Article service written in Go. It allows for the following core functionality:

- Get an article by ID
- Create an article
- Get tag/article information based on a given tag and date

## Installation

This application runs on Go v1.14. Please make sure you have a compatible version on your machine by running `go version`, or by installing the latest version [here](https://golang.org/doc/install).

## Setup

To start the app, run the command:

```
go run cmd/main.go
```

Alternatively, you can build the app into a binary first and then execute the binary by running the following:

```
go build -o bin/article-rest-api cmd/main.go
./bin/article-rest-api
```

## Development

To run tests, simply run:

```
go test ./...
```

## Solution

I created the solution with the intent of building a simple service in a structure that could be easily extended or modified. This is due to scalability; if ever we needed to change an implementation or piece of infrastructure, we shouldn't have to spend time changing every single part of the service.

### Structure

I structured the service in a way that allows us to draw clear boundaries between the different layers of our application. These layers are:

- Domain
  - Contains business domain models and logic
  - Should not be concerned about/affected by how the application logic is implemented
- Service
  - The application layer, contains application logic
  - Should not be concerned by presentation layer or persistence layer
- Handler
  - The presentation layer, contains logic on how we present data to the outside world
  - In this implementation we are using a REST handler
- Repository
  - The data persistence layer, contains logic on retrieving/mutating data
  - Should be easily swappable/extendable, not tied together with service or domain logic
  - In this implementation we are using an local JSON file as a database

With this sort of structure, it hopefully shouldn't be too much work to change the data persistence implementation from an in memory store to a SQL database, or change the presentation later from REST to GraphQL (although I haven't tried...)

I've also tried to separate the representation of the Article model in the different packages (i.e Article as a domain model, Article as a JSON response etc). This may initially lead to a duplication of the Article struct in multiple packages, but I believe it's worth it as the service gets bigger. While we could try to have everything fit in the one domain struct (eg. having multiple json and validation tags on our domain model), this creates a coupling between the domain model and other layers. If we wanted to change what we send as a JSON response, or change what exactly we store and retrieve from our data layer, we could simply change it in the handler or the repository layer, instead of modifying our domain model and hoping it doesn't affect any other part of the codebase.

### Language

I chose Go because that's my preferred language when it comes to service/API based projects, as well as it being the backend language I probably spend the most time on at my job. I also chose it because it seems to be the preferred language for this submission.

Go also tends to lend itself well to the layered architecture with its use of interfaces, which allows for flexible dependency injection which is conducive to an extendable architecture.

### Internal Packages

#### Main

The main package exists as the entryway to the application. It's where all the setup happens (logging, server setup, config) as well as where we tie all of our internal packages together through dependency injection (passing repos to services etc)

#### Domain

Contains the Article model in a form that is specific to the business domain it belongs to. Currently there is not enough information for me to include anything more than a simple model, however business logic such as validation or other specific actions can belong here, as well as of course other related business models.

#### Service

Contains application logic, such as getting articles, creating them and getting tag information. The application logic for the most part was pretty straightforward, only the tag endpoints requires some small finessing of data.

#### LocalJSON

This was the chosen data persistence implentation for this submission. I decided not to do a proper SQL based solution as I thought it would be too much for this exercise. The way it works is we simply have a JSON file of articles as our data store, which we query and write to and from. Definitely not the most robust solution, but can definitely be swapped out for something more holistic like Postgres.

#### Rest

Contains handler logic for wiring up REST routes to service layer logic. This is basically where all our routes get mounted. Any sort of authentication middleware can also be set here, although I opted to not include any as I thought it was out of scope for this exercise.

#### Internal

The only file in this internal package is the `errors.go` file which contains application based error handling, which I'll go into more in the error handling section below.

### Error Handling

Coming up with a decent error handling solution is probably what took the most time, and I ended up using a solution I found from this example microservice [here](https://github.com/MarioCarrion/todo-api-microservice-example/blob/main/internal/errors.go).

The reason behind this is that we are able to pass around application specific error codes, while also building context through wrapped error messages as our errors are passed around the application. This means that we can use the application error codes to do programmatically peform actions based on specific errors (which helps in setting correct status codes), while also providing detailed readable error messages with context to the end user.

### Testing

The way I approached testing was through simple unit assertion tests and mocking dependencies. The layered structure coupled with Go's interfaces allows for easy mock testing.

## Assumptions

Below are a list of assumptions made as I built the submission:

- Article IDs are strings and not ints
- Tags aren't set in a data store and can be anything sent from the API
- All fields from the example article are required to create an article
- Dates must be the same as the format specified, and may not always be formatted correctly when coming from the client
- The higher the ID, the later the article wsa submitted (in regards to getting the 10 last articles for the tag endpoint)
- Authentication isn't in the scope of this exercise
