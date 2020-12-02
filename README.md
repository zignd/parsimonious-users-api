# parsimonious-users-api

A users API exposing endpoints to search for users based on its name and username. Some users have higher priority over others; therefore, to identify those users and prioritize them, we also consider two "relevance" lists during the search, which will affect the final result.

The API itself has the following endpoints:

### `GET /users/name/:name?page=:page&pageSize=:pageSize`

Retrieves users whose names partially match the provided value for `name`.

Parameters:
  * `name`: Part of the name of the user to search for
  * `page`: the current page of the results (default: 1)
  * `pageSize`: the maximum number of users to return in a single request (default and maximum: 15)

### `GET /users/username/:username?page=:page&pageSize=:pageSize`

Retrieves users whose usernames partially match the provided value for `username`.

Parameters:
  * `username`: Part of the username of the user to search for
  * `page`: the current page of the results (default: 1)
  * `pageSize`: the maximum number of users to return in a single request (default and maximum: 15)

### `GET /health-check`

A simple health check endpoint to check whether or not the connection to the database is still available.

## Running the application locally

The application is an HTTP server written in Go and uses a PostgreSQL database for the user's data. It's entirely automated to run itself and its dependencies in containers; therefore, you only need to have Docker and make installed.

You can start the application using `make up`, and you can shut it down using `make down`.

During the start-up (`make up`), the users' data will be downloaded from an external URL, then the download file is extracted, and its CSV content will be imported to the database. I've noticed that the step of importing the data takes around 3~5 minutes. Please be patient at this point. After that, the application is started. It will be listening by default to port 3000, mapped to the same port in your machine.

You can experiment with it by sending HTTP requests to it. Here are a few examples which will return some data to get you started:

```
curl http://localhost:3000/users/name/Ta\?page\=1\&pageSize\=15
```

```
curl http://localhost:3000/users/username/ba\?page\=1\&pageSize\=15
```

```
curl http://localhost:3000/health-check
````

## Tests

The application tests were automated to run inside and outside containers. The idea was to use the container-based version inside a CI pipeline and locally for development.

To run the tests inside containers, you can use `make run-test`. And to run locally (this one requires Go installed), you can use `run-test-local`.

You can clean up the containers created using `make down` as well.

## Regarding the project name

Parsimonious is nothing but a randomly chosen word to uniquely name this repository, but the meaning according online dictionaries is related to "unwillingness to spend". According to Merriam-Webster:

> 1 : exhibiting or marked by parsimony especially : frugal to the point of stinginess.
> 2 : sparing, restrained.