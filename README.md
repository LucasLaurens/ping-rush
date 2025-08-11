# ping-rush
Little application to check whether certain web pages are responding and what their status is.

## Description
This package allows you to test whether one or more URLs are available or not (up and down).
Two methods are available, “Ping” and “Pings,” to test one or more URLs and return an error if a problem occurs.
Improvements still need to be made in terms of performance, maintainability, and functionality.
This small package is a kind of skeleton available to provide a basis on which to test URLs and verify the expected status.
Feel free to create issues and related PRs to improve this package, but also to fix certain parts so that it also meets needs that were not initially considered.
The PRs will be analyzed and then merged if relevant to this small project.

## Local
### Init
You must follow the commands bellow to start the project
```go
go init {project_name}
go build
```

### Improvement
You could use the “os” package to define arguments and pass the URLs as arguments with appropriate tests.

