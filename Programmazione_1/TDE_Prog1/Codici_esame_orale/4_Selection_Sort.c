#include <stdio.h>
#include <stdlib.h>
#include <time.h>

#define DIM 5
#define MAX 10

main(){
    int array[DIM], i, j, k, cambio;

    srand(time(NULL));

    for(i=0; i<DIM; i++){
        array[i]=rand() %MAX;
    }

    printf("Array: ");

    for(i=0; i<DIM; i++){
        printf("%d ", array[i]);
    }

    for(i=0; i<DIM; i++){
        for(j=0; i<DIM-1; j++){
            printf("\n");
            for(k=0; k<DIM; k++){
                printf("%d ", array[k]);
            }
            if(array[j]>array[j+1]){
                cambio=array[j];
                array[j]=array[j+1];
                array[j+1]=cambio;
                break;
            }
        }
    }
}
