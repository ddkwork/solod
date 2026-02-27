#include "main.h"

int main(void) {
    so_Slice nums = {(so_int[5]){1, 2, 3, 4, 5}, 5, 5};
    so_Slice s1 = so_slice(nums, so_int, 0, nums.len);
    so_index(s1, so_int, 1) = 200;
    (void)s1;
    so_Slice s2 = so_slice(nums, so_int, 2, nums.len);
    (void)s2;
    so_Slice s3 = so_slice(nums, so_int, 0, 3);
    (void)s3;
    so_Slice s4 = so_slice(nums, so_int, 1, 4);
    (void)s4;
    so_int n = so_copy(s4, s1, so_int);
    (void)n;
    so_Slice strSlice = {(so_String[3]){so_strlit("a"), so_strlit("b"), so_strlit("c")}, 3, 3};
    so_int sLen = so_len(strSlice);
    (void)sLen;
    so_Slice twoD = {(so_Slice[2]){{(so_int[3]){1, 2, 3}, 3, 3}, {(so_int[3]){4, 5, 6}, 3, 3}}, 2, 2};
    so_int x = so_index(so_index(twoD, so_Slice, 0), so_int, 1);
    (void)x;
}
