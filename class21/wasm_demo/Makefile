.PHONY: clean, goexec

test.wasm:
	env GOOS=js GOARCH=wasm go build -o test.wasm .

startup:
	goexec 'http.ListenAndServe(`:8080`, http.FileServer(http.Dir(`.`)))'

clean:
	- rm test.wasm

goexec:
	go get -u github.com/shurcooL/goexec