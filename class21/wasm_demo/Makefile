.PHONY: clean

test.wasm:
	env GOOS=js GOARCH=wasm go build -o test.wasm .
	goexec 'http.ListenAndServe(`:8080`, http.FileServer(http.Dir(`.`)))'

clean:
	- rm test.wasm