.PHONY: clean

TARGET=cxx_test

$(TARGET): libfoo.a
	go build -x -work .

libfoo.a: foo.o cfoo.o
	ar r $@ $^

%.o: %.cpp
	g++ -O2 -o $@ -c $^

clean:
	rm -f *.o *.so *.a $(TARGET)
	go clean -x -cache