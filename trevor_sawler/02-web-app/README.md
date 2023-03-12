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

# Testing POST Requests

- Install Postgres with Docker
- Connect our web app to Postgres
- Add authentication to our application
- Write some authentication middleware & test it
- Write tests for authentication to our web application


