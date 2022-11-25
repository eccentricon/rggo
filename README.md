# rggo

Exercises from [_Powerful Command-Line Applications in Go_](https://pragprog.com/titles/rggo/powerful-command-line-applications-in-go/) by Ricardo Gerardi.

## Go references

- [Go user manual][documentation]
- [Using Go Modules]
- [Setting up your workspace] (esp. `go.work` file and multiple modules)

[setting up your workspace]: https://github.com/golang/tools/blob/master/gopls/doc/workspace.md#setting-up-your-workspace
[using go modules]: https://go.dev/blog/using-go-modules
[documentation]: https://go.dev/doc/

### Help

`go help` for top-level help.

- Use `go help <command>` for more information about a command.
- Use `go help <topic>` for more information about that topic.

### Commands

- Modules
  ```
  # Initialize a module (p. 37)
  go mod init pragprog.com/rggo/firstProgram/wc
  ```
- Test
  ```
  # Run test in "local directory mode" (see 'go help test') (p. 43)
  go test -v
  ```
- Build

  ```
  # Create executable in the current directory (p. 43)
  go build

  # Create exe in './bin' subdirectory (see 'go help build')
  go build -o bin
  ```
