#include <stdio.h>

// gcc test.c

int main() {
    int x = 1337;
    
    for (;;) {
        printf("%p = %d\n", &x, x);
        sleep(1);
    }

    return 0;
}