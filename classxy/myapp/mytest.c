#include "stdio.h"
#include "libexportc.h"

int main() {
    printf("This is a test\n");

    GoString str = {"Hi JXES", 7};
    Hello(str);

    printf("sum: %d+%d=%lld", 3, 5, MySum(3, 5));

}