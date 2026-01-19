# Contributing to RelayWarden Go SDK

Thank you for your interest in contributing to the RelayWarden Go SDK! This document provides guidelines and instructions for contributing.

## Code of Conduct

This project adheres to a Code of Conduct. By participating, you are expected to uphold this code.

## Getting Started

### Prerequisites

- Go 1.23 or higher
- Git

### Setting Up the Development Environment

1. Fork the repository
2. Clone your fork:
   ```bash
   git clone https://github.com/your-username/relaywarden-go-sdk.git
   cd relaywarden-go-sdk
   ```
3. Download dependencies:
   ```bash
   go mod download
   ```

## Development Workflow

### Making Changes

1. Create a new branch from `main`:
   ```bash
   git checkout -b feature/your-feature-name
   ```
2. Make your changes
3. Ensure code follows Go best practices and conventions
4. Write or update tests
5. Format code:
   ```bash
   gofmt -w .
   ```
6. Run `go vet`:
   ```bash
   go vet ./...
   ```
7. Run tests:
   ```bash
   go test ./...
   go test -v -cover ./...
   ```
8. Commit your changes with clear, descriptive messages

### Coding Standards

- Follow [Effective Go](https://go.dev/doc/effective_go) guidelines
- Use `gofmt` for code formatting
- Run `go vet` to catch common mistakes
- Add comments for exported functions and types
- Keep functions focused and single-purpose
- Use meaningful variable and function names
- Follow Go naming conventions (exported = Capitalized, unexported = lowercase)

### Testing

- Write tests for all new functionality
- Use the `testing` package
- Ensure all tests pass before submitting
- Aim for high test coverage
- Test both success and error cases
- Use table-driven tests when appropriate

### Code Formatting

- Use `gofmt` to format code
- Run `gofmt -w .` before committing
- The project uses standard Go formatting

### Commit Messages

Follow [Conventional Commits](https://www.conventionalcommits.org/) format:

```
type(scope): subject

body (optional)

footer (optional)
```

Types: `feat`, `fix`, `docs`, `style`, `refactor`, `test`, `chore`

Example:
```
feat(messages): add support for message cancellation

Add the ability to cancel pending messages via the API.
This includes retry logic and proper error handling.
```

## Submitting Changes

1. Push your branch to your fork
2. Create a Pull Request targeting the `main` branch
3. Fill out the PR template completely
4. Ensure all CI checks pass
5. Address any review feedback

### Pull Request Checklist

- [ ] Code follows Go best practices
- [ ] Code formatted with `gofmt`
- [ ] `go vet` passes
- [ ] Tests added/updated and passing
- [ ] Documentation updated
- [ ] Commit messages follow Conventional Commits
- [ ] PR description is clear and complete

## Reporting Issues

When reporting bugs or requesting features:

1. Check existing issues to avoid duplicates
2. Use the appropriate issue template
3. Provide clear steps to reproduce (for bugs)
4. Include Go version, SDK version, and environment details
5. Add code examples when relevant

## Documentation

- Update README.md for user-facing changes
- Add comments for exported functions and types
- Keep code examples up to date
- Follow Go documentation conventions

## Questions?

- Open a discussion for questions
- Check existing issues and discussions
- Review the README for common usage patterns

Thank you for contributing! 🎉
