#pragma once

#ifdef __cplusplus
extern "C" {
#endif

typedef void* Foo;
typedef void* Test;

Foo FooInit(void);
void FooFree(Foo);
void FooBar(Foo);

Test TestNew(const int argc, const char **argv, int *err);
void TestFree(Test);

#ifdef __cplusplus
}
#endif