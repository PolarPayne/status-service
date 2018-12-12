# Status Service

A simple Go program, that serves a web server that always returns the specified status. Configuration is done exclusively with environment variables.

## Configuration
Port, host, and status code can also be specified with arguments, but environment will always override those.
* `PORT` (or `-port PORT` on CLI), defaults to `80`.
* `HOST` (or `-host HOST` on CLI), defaults to `0.0.0.0`.
* `STATUS_CODE` (or `-code STATUS_CODE` on CLI), defaults to `200`, must be between 200-299, 400-499, or 500-599.
* `STATUS_MESSAGE`, defaults to the default message for the code or empty if it doesn't have one.

## Makefile
Run `make` to run the tests, and build the normal and minified binary.

## Dockerfile
The included Dockerfile builds a very tiny container with just the binary in it (size is currently ~1.7MB).
