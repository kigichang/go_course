.PHONY: clean test

mytest: mygo.a
	gcc -o mytest mygo/mygo.a mytest.c

mygo.a:
	$(MAKE) -C mygo

clean:
	- $(MAKE) clean -C mygo
	- rm mytest mygo.h mygo.a
