#include "foo.hpp"
#include "foo.h"

Foo FooInit() {
    cxxFoo *ret = new cxxFoo(1);
    return (void *)ret;
}

void FooFree(Foo f) {
    cxxFoo *foo = (cxxFoo *)f;
    delete foo;
}

void FooBar(Foo f) {
    cxxFoo *foo = (cxxFoo *)f;
    foo->Bar();
}

Test TestNew(const int argc, const char **argv, int *err) {
    cxxTest *ret = new cxxTest(argc, argv, err);
    return (void *)ret;
}

void TestFree(Test t) {
    cxxTest *tt = (cxxTest *)t;
    delete tt;
}