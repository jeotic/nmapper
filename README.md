## NMapper Prototype
This application allows you to view NMap runs by IP Address and uploading new Runs. Note, this is in very early stages, and *many* shortcuts have been taken.

## Note
I did not add support for a production build for the UI. You'll need to use the built in `npm start` (read below)

The UI and API are separate services. 
API: `locahost:8000`
UI: `localhost:3000`

Everything is hard-coded currently
The UI is very ugly and needs a lot of love. 

## File Uploader Notes
I chose XML over the other formats for various reasons:

 - Go has native support for it
 - Easier to tokenize and insert into a DB
 - A well-known standard format for APIs
 - I have the most familiarity with it over the others

I did not Marshal or Unmarshal XML to Structs. I do a pretty raw read to SQL queries to DB. This allows for a small memory footprint, more control, and was quick to implement. 

I used a DB query builder to save time, but can easily be replaced. I implemented Gorm after the importer, so you'll see me grab the underlying DB connection from Gorm. Switching to Gorm would be a valid option down the road

## Why I used Gorm
I was originally going raw SQL, but that was proving to eat up a lot of time and was tedious. ORMs aren't evil, but I usually am wary of them. They do, however, save a tremendous amount of time to get something going

## Why the database is meh
Faster prototyping and not meant for production. I'm not happy with a lot of the column names, nor types (but it's SQLite). It got the job done quickly though

## UI
Simply navigate to `/web` and run `npm start`. This will host an HTTP server on `localhost:3000`


## To Build
You'll need Go 1.13 and GCC in your %PATH. Simply install all the dependencies
```
go get ./...
```
Then run the server
`go run cmd/server/server.go`

You can also run the parser alone
`go run cmd/server/parser.go`

## Things to do

 - Clean up code
 - Make things consistent
 - Make interfaces for returned data in UI. Right now it just uses `any`
 - Tests tests tests
 - Implement Gorm in Parser?
 - Actually have error handling in the endpoints. Right now it just returns ALL errors
 - Prettier status codes and messages
 - Prettify the UI. This needs a lot love
 - Make production worthy
 - Possibly have API server HTML? I prefer separation of services though

## Thoughts
I hope this gives some insight in my capabilities. Unfortunately, I can't dedicate a lot of time to make it production worthy, nor do what I usually do. I can do prototypes, which requires shortcuts, but I prefer a solid foundation to build off of; this isn't a great example of that. Ideally, the UI wouldn't even live in the code-base. 