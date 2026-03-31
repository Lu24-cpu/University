#include <stdio.h>
#include <time.h>

#define DIM 100

int Random(){
    srand((unsigned)time(NULL));

    return rand() % DIM + 1;
}

int main(){
    int x = Random(), y;

    y = x;
    printf("Dimensione triangolo: %d\n", x);
    do{
        for (int i=x; i>0; i--){
            if (i==x || i == 1 || x == y){
                printf("*");
            } else {
                printf(" ");
            }
        }
        printf("\n");
        x--;
    }while (x>0);
}