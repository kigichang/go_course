#include <iostream>
#include "foo.hpp"
  
double Foo = 3.0;
  
void cxxFoo::Bar(void) {
    std::cout << this->a<<std::endl;
}
  
int gcd(int x, int y) {
    int g;
    g = y;
  
    while (x > 0) {
        g = x;
        x = y % x;
        y = g;
    }
  
    return g;
}
  
double add(double x) {
    return x + Foo;
}