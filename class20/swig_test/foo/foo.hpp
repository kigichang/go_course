#pragma once

extern double Foo;

class cxxFoo {
public:
    int a;
    cxxFoo(int _a): a(_a) {};
    ~cxxFoo() {};
    void Bar();
};

extern int gcd(int x, int y);
extern double add(double x);