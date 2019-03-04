.PHONY: clean
  
TARGET=gocpp

$(TARGET): libfoo.a
	go build -x -work .

libfoo.a: foo.o cfoo.o
	ar r $@ $^

%.o: %.cpp
	clang -O2 -o $@ -c $^

clean:
	rm -f *.o *.so *.a $(TARGET)