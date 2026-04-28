#include "main.h"
#include <math.h>

// -- Implementation --

int main(void) {
    {
        // Typed C expression.
        double nan = NAN;
        if (nan == nan) {
            so_panic("nan == nan");
        }
        double x = sqrt(49);
        if (x != 7) {
            so_panic("x != 7");
        }
    }
    {
        // Raw C block.
        so_int b = 0;
        int a = 7;
        b = a * a;
        b = sqrt(b);
        if (b != 7) {
            so_panic("b != 7");
        }
    }
}
