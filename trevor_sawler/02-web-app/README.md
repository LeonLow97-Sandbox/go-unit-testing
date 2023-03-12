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
