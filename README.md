# Status Service

A simple Go program, that serves a web server that always returns the specified status. Configuration is done exclusively with environment variables.

## Configuration
* `PORT`, defaults to `80`.
* `HOST`, defaults to `0.0.0.0`.
* `STATUS_CODE`, defaults to `200`, must be between 200-299, 400-499, or 500-599.
* `STATUS_MESSAGE`, defaults to the default message for the code or empty if it doesn't have one.

## Makefile
Run `make` to run the tests, and build the normal and minified binary.

## Dockerfile
The included Dockerfile builds a very tiny container with just the binary in it (size is currently ~1.7MB).
