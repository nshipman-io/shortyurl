# ShortyUrl
ShortyUrl is built with Go based on the FreeCodeCamp project:
https://www.freecodecamp.org/learn/apis-and-microservices/apis-and-microservices-projects/url-shortener-microservice

## Installation
`go build`

## Usage 
1. `./shortyurl`
2. From the browser visit the url `0.0.0.0:8080` or `127.0.0.1:8080`
3. Input a valid url into the input field. You should receive an ID.
3. Enter your ID in the url `/api/shorturl/{id}`

## TODO

1. Build a user account system, so users can save, view and delete their generated links.
2. Implement hashing algorithm to generate url ids.
3. Setup the project to be "production ready".
4. Deploy to Heroku or GCP.

