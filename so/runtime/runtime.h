#include <time.h>
#include <stdint.h>

#if defined(so_build_darwin)
#include <stdlib.h>
#elif defined(so_build_linux)
#include <sys/random.h>
#endif

// Seed returns a random 64-bit seed.
static inline uint64_t runtime_Seed(void) {
    uint64_t seed = 0;
#if defined(so_build_darwin)
    arc4random_buf(&seed, sizeof(seed));
#elif defined(so_build_linux)
    if (getrandom(&seed, sizeof(seed), 0) != sizeof(seed)) {
        // Fallback to time-based seed.
        struct timespec ts;
        clock_gettime(CLOCK_MONOTONIC, &ts);
        seed ^= (uint64_t)ts.tv_nsec ^ (uint64_t)ts.tv_sec;
    }
#else
    seed = (uint64_t)time(NULL) ^ (uintptr_t)&seed;
#endif
    return seed;
}

#define runtime_buildVersion so_str(so_version)

#if defined(so_build_darwin)
#define runtime_GOOS so_str("darwin")
#elif defined(so_build_linux)
#define runtime_GOOS so_str("linux")
#elif defined(so_build_windows)
#define runtime_GOOS so_str("windows")
#else
#define runtime_GOOS so_str("unknown")
#endif

#if defined(so_build_amd64)
#define runtime_GOARCH so_str("amd64")
#elif defined(so_build_arm64)
#define runtime_GOARCH so_str("arm64")
#else
#define runtime_GOARCH so_str("unknown")
#endif
