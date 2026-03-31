#include <stdio.h>
#include <stdlib.h>
#include <time.h>

#define DIM 5
#define MAX 10

main(){
    int i, flag=1, k, a[DIM], stop=DIM-1, cambio;

    srand(time(NULL));

    for(i=0; i<DIM; i++){
        a[i]=rand() % MAX;
    }

    printf("Array: ");

    for(i=0; i<DIM; i++){
        printf("%d ", a[i]);
    }
    printf("\n");

    while(flag==1){
        flag=0;

        for(i=0; i<stop; i++){
            if(a[i]>a[i+1]){
                cambio=a[i];
                a[i]=a[i+1];
                a[i+1]=cambio;
                flag=1;

                for(k=0; k<DIM; k++){
                    printf("%d ", a[k]);
                }
                printf("\n");
            }
        }
    }
}
