//go:build ignore
#include "os.h"

os_File os_Stdin = {0};
os_File os_Stdout = {0};
os_File os_Stderr = {0};

static void __attribute__((constructor)) os_init() {
    os_Stdin = (os_File){.fd = stdin};
    os_Stdout = (os_File){.fd = stdout};
    os_Stderr = (os_File){.fd = stderr};
}
