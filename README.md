# Simple Rate Limit

### Introduction
The project is used to implement simple rate-limit feature.

Used technologies:
- `echo` as web service framework to handle http request.
- `redis` as cache database to calculate request limit.
- `viper` as config component.
- `logrus` as log component.
- `dig` as dependency injection component.
- `dockertest` and `httpexpect` as test utility for integration testing.

The reason for using redis is that it allows me to use `incr` to make a counter easily.

### Rules
- Server can only accept 60 requests per minute each IP (use `X-Forwarded-For`).
- It will response the text of N if N <= 60.
- It will response the text `Error` if N > 60.

### Demo
The project is hosted in [Heroku](https://young-wave-60838.herokuapp.com/) and uses `Heroku Redis` as database.

Sending requests by `curl -X POST https://young-wave-60838.herokuapp.com/post` allows you to test this function easily.

### Test (For Mac)
Required: Docker Desktop, Go.

Execute `go test -run=. -v ./src/tests` to see the test result.

The detail are listed below:
1. Use `dockertest` to run docker `redis` container as database.
2. Start running http server for integration testing.
3. Use `httpexpect` to check responses if correct.
    - First, call `/post` 60 times, it must be responsed the text of requested count with http status `200`.
    - Then, call `/post` once, it must be responsed the text `Error` with http status `429`.
4. Stop running http server and purge `redis` container.
