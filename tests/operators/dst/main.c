#include "main.h"

int main(void) {
    {
        so_int a = 11, b = 22, c = 33;
        so_int d = b / a + (a - c) * a + c % b;
        d += 10;
        d -= 10;
        d *= 10;
        d /= 2;
        d %= 5;
        d++;
        d--;
        (void)d;
    }
    {
        double x = 1.1, y = 2.2, z = 3.3;
        double f = x / y + (y - z) * x;
        f += 1.0;
        f -= 1.0;
        f *= 2.0;
        f /= 2.0;
        f++;
        f--;
        (void)f;
    }
    {
        so_int b1 = 0b1010, b2 = 0b1100;
        so_int b3 = (b1 | b2) & (b1 & b2) | (b1 ^ b2);
        b3 = b3 << 2;
        b3 = b3 >> 1;
        b3 = b3 & ~b1;
        (void)b3;
        so_int b4 = 0b1010;
        b4 |= 0b1100;
        b4 &= 0b1100;
        b4 ^= 0b1100;
        (void)b4;
    }
    {
        bool a = true, b = false, c = true;
        bool d = (a && b) || (b || c) && !a;
        (void)d;
        so_int x = 10, y = 20, z = 30;
        bool e1 = (x < y) && (y > z) || (x == z);
        (void)e1;
        bool e2 = (x <= y) && (y >= z) || (x != z);
        (void)e2;
    }
}
