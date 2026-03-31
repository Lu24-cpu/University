#include <stdio.h>
#include <stdlib.h>
#include <time.h>

#define MAX 10
#define DIM 5

main(){
    int a[DIM], i, k, j, cambio;

    srand(time(NULL));

    for(i=0; i<DIM; i++){
        a[i]=rand()%MAX;
    }

    printf("Array: ");
    for(j=0; j<DIM; j++){
        printf("%d ", a[j]);
    }
    printf("\n");
    for(i=1; i<DIM; i++){
        k=i-1;

        while(k>=0 && a[k]>a[k+1]){
            cambio=a[k];
            a[k]=a[k+1];
            a[k+1]=cambio;
            k=k-1;

            for(j=0; j<DIM; j++){
                printf("%d ", a[j]);
            }
            printf("\n");
        }
    }
}
