#include "main.h"
static const so_String s = so_strlit("constant");

static double sin(double x) {
    return x;
}

int main(void) {
    so_println("%s", s.ptr);
    const so_int n = 500000000;
    const double d = 3e20 / n;
    so_println("%f", d);
    so_println("%lld", (int64_t)d);
    so_println("%f", sin(n));
}
