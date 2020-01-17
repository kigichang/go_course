#include <iostream>
#include "foo.hpp"

void cxxFoo::Bar(void) {
    std::cout << this->a<<std::endl;
}

cxxTest::cxxTest(const int argc, const char **argv, int *err) {

    for(int i = 0; i < argc; i++) {
        std::cout << argv[i] << std::endl;
    }
    *err = 100;
}

cxxTest::~cxxTest() {
    std::cout << "delete cxxTest" << std::endl;
}