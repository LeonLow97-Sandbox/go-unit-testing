## Commands

- `go test ./...`
  - run all tests in directory and sub-directories.

# Testing Sessions

- Sessions keep state for web apps that render pages server side.
- Add session support to our web app.
- Write some code that uses sessions.
- Write tests that use sessions.
- Web servers are stateless and sessions give us a means of persisting information between requests.

## Creating a `TestMain` function

- This test function will be ran before all the tests.
- Used to set up the test environment

```go
// Set up Test Environment
func TestMain(m *testing.M) {
	os.Exit(m.Run())
}
```

## Creating a `testdata` folder

- Must be exactly this name for the folder `testdata`.
- This folder will be ignored when building the application. Use for testing data.

## `setup_test.go`

- Setup file for unit tests.
- Defines any necessary setup functions or variables that are used across multiple test files.
- Provides a central place for defining test setup logic that can be reused across multiple test files.
- Helps to reduce code duplication and ensure consistency across tests.

# Testing POST Requests

- Install Postgres with Docker
- Connect our web app to Postgres
- Add authentication to our application
- Write some authentication middleware & test it
- Write tests for authentication to our web application

## Avoid testing cache

- `go test -count=1`

## The Repository Pattern

- Design pattern for managing interactions between an application and a database.
- Creating a set of code that acts as an interface between the application and the database, which can be easily swapped out if needed.
- By creating 2 separate repositories, can write unit tests without actually using a live database.
- Test repository can simply return pre-determined results that simulate database interactions, which allows for faster and more reliable testing.
- Can test our code quickly and reliably, without worrying about the state of the database or any potential issues that might arise.


