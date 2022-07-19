#include <stdio.h>
#include <iostream>

// **重點**
// 在C++如果Go編譯(看成C)的話，需要引入.h
extern "C"
{
#include "hello.h"
}

void SayHello(const char *s)
{
    std::cout << "say: " << s << std::endl;
}

void hello()
{
    printf("Hello From C\n");
}
