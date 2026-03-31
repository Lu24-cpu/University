#include <stdio.h>
#include <stdlib.h>
#include <ctype.h>
#include <string.h>

#define DIM 100

char IsPalindrome(char stringa[DIM], int count) {
    if (count >= strlen(stringa)/2){
        return 0;
    }

    if (stringa[count] == stringa[strlen(stringa)-count]) {
        return IsPalindrome(stringa, count++);
    }
    return 1;
}

main(){
    char stringa[DIM];

    printf("Inserire una parola per controllare se palindroma: ");
    scanf("%s", &stringa);

    if (IsPalindrome(stringa, 0) == 0){
        printf("La stringa è palindroma\n");
    } else {
        printf("La string non è palindroma\n");
    }
}