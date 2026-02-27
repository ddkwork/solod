#include "main.h"

int main(void) {
    so_Slice nums = so_make_slice(so_int, 3, 3);
    so_int n = so_index(nums, so_int, 1);
    so_index(nums, so_int, 1) = 42;
    so_int l1 = so_len(nums);
    so_int c1 = so_cap(nums);
    nums = so_make_slice(so_int, 0, 3);
    nums = so_append(nums, so_int, 1);
    nums = so_append(nums, so_int, 2, 3);
    so_int l2 = so_len(nums);
    so_int c2 = so_cap(nums);
    (void)n;
    (void)l1;
    (void)c1;
    (void)l2;
    (void)c2;
}
