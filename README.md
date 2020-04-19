Earth DateTime to Mars Time
=====================

This project outputs a Coordinated Mars Time and Mars Sol Date for the given UTC

install
```bash
~$ go get -u github.com/wesleyvicthor/marsoldate
```
to execute the service run
```
~$ marsmission 2019-12-27T15:22:22Z
```

it expects a datetime [RFC3339](https://tools.ietf.org/html/rfc3339) format in order to compute a valid response.

for tests execution
```
~$ go test
```

The algorithm has been ported from [James Tauber](http://jtauber.github.io/mars-clock/) and partially validated against mars24 from nasa.

