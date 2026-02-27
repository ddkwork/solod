#include "main.h"

static main_Person newPerson(so_String name) {
    main_Person p = {.name = name};
    p.age = 42;
    return p;
}

int main(void) {
    main_Person bob = {so_strlit("Bob"), 20};
    (void)bob;
    main_Person alice = {.name = so_strlit("Alice"), .age = 30};
    (void)alice;
    main_Person fred = {.name = so_strlit("Fred")};
    (void)fred;
    main_Person* ann = &(main_Person){.name = so_strlit("Ann"), .age = 40};
    *ann = newPerson(so_strlit("Jon"));
    (void)ann;
    main_Person sean = {0};
    sean.name = so_strlit("Sean");
    sean.age = 50;
    main_Person* sp = &sean;
    sp->age = 51;
    (void)sean;
    so_auto dog = (struct {
        so_String name;
        bool isGood;
    }){
        .name = so_strlit("Rex"),
        .isGood = true,
    };
    (void)dog;
}
