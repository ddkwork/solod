#include "main.h"

int main(void) {
    so_println("%s", "golang");
    so_println("%s %lld", "1+1 =", 1 + 1);
    so_println("%s %f", "7.0/3.0 =", 7.0 / 3.0);
    so_println("%s %d", "true && false =", true && false);
    so_println("%s %d", "true || false =", true || false);
    so_println("%s %d", "!true =", !true);
}
