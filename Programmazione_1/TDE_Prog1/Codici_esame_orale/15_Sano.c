#include <stdio.h>
#include <stdlib.h>
#include <time.h>

#define MAX 10
#define DIM 8
main(){
    int array[DIM], i;
    char x;

    srand((unsigned)time(NULL));

    for(i=0; i<MAX; i++){
        array[i]= (rand() % MAX);
        while(array[i]<1){
            array[i]= (rand() % MAX);
        }
    }

    printf("Inserire un carattere speciale, tra |, !, #, *, $, &, ?, @: ");
    scanf("%c", &x);

    switch(x){
        case('|'):
            printf(" | _\n");
            break;
        case('!'):
            printf(" ! _ _\n");
            break;
        case('#'):
            printf(" # _ _ _\n");
            break;
        case('*'):
            printf(" * _ _ _ _\n");
            break;
        case('$'):
            printf(" $ _ _ _ _ _\n");
            break;
        case('&'):
            printf(" & _ _ _ _ _ _\n");
            break;
        case('?'):
            printf(" ? _ _ _ _ _ _ _\n");
            break;
        case('@'):
            printf(" @ _ _ _ _ _ _ _ _\n");
            break;
    }

    printf("   %d", array[0]);

    if('x'!='|'){
        for(i=1; i<DIM; i++){
            printf(" %d", array[i]);
        }
    }
}
