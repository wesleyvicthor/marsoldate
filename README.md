Earth DateTime to Mars
=====================

to execute the service run
```
~$ marsmission 2019-12-27T15:22:22Z
```

it expects a datetime [RFC3339](https://tools.ietf.org/html/rfc3339) format in order to compute a valid response.

for tests execution
```
~$ go test
```

Golang was used to achieve the task as it express simplicity and selfcontained dependencies; with a strong portability to a vast
operating systems and platform architectures.

The algorithm was based and ported from James Tauber and partialy validated agains mars24 from nasa.

