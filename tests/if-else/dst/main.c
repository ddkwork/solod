#include "main.h"

int main(void) {
    if (7 % 2 == 0) {
        so_println("%s", "7 is even");
    } else {
        so_println("%s", "7 is odd");
    }
    if (8 % 2 == 0 || 7 % 2 == 0) {
        so_println("%s", "either 8 or 7 are even");
    }
    if (1 == 2 - 1 && (2 == 1 + 1 || 3 == 6 / 2) && !(4 != 2 * 2)) {
        so_println("%s", "all conditions are true");
    }
    if (9 % 3 == 0) {
        so_println("%s", "9 is divisible by 3");
    } else if (9 % 2 == 0) {
        so_println("%s", "9 is divisible by 2");
    } else {
        so_println("%s", "9 is not divisible by 3 or 2");
    }
    {
        so_int num = 9;
        if (num < 0) {
            so_println("%lld %s", num, "is negative");
        } else if (num < 10) {
            so_println("%lld %s", num, "has 1 digit");
        } else {
            so_println("%lld %s", num, "has multiple digits");
        }
    }
}
