#include <stdio.h> // standard input and output
#include "hello.h"
// #include <stdlib.h>

int SayHello(char* name) {
    printf("Hello from C, %s!\n", name);
    return 0;
}