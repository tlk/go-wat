# Surprising properties of golang

```
$ go run github.com/tlk/go-wat@latest
the last element of b is: 456  (expected 456)
the last element of c is: 999  (expected 999)
the last element of y is: 999  (expected 456)  wat?!
the last element of z is: 999  (expected 999)
```

Alternatively,
```
$ git clone git@github.com:tlk/go-wat.git
$ cd go-wat
$ go run main.go
```

### Acknowledgements
https://gist.github.com/erikcorry/5bcd95b916527969a4210a3aa32bd7a0
