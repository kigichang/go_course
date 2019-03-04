#pragma once
  
class cxxFoo {
public:
    int a;
    cxxFoo(int _a): a(_a) {};
    ~cxxFoo() {};
    void Bar();
};

class cxxTest {
public:
    cxxTest(const int argc, const char **argv, int *pnErr);
    ~cxxTest();
};