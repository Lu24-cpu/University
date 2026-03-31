#include <stdio.h>
#include <stdlib.h>

#define DIM 100

void CreaMat(int mat1[DIM][DIM], int mat2[DIM][DIM], int dim){
    for(int i=0; i<dim; i++){
        for (int j=0; j<dim; j++){
            mat1[i][j]=(i+1)*(j+1);
            mat2[dim-i-1][dim-1-j]=(i+1)*(j+1);
        }
    }
}

int SommaRiga(int riga1[DIM], int riga2[DIM], int riga[DIM], int j, int dim){
    if (j==dim){
        return 0;
    }

    riga[j]= riga1[j]+riga2[j];
    SommaRiga(riga1, riga2, riga, j+1, dim);
}

int SommaMat(int mat1[DIM][DIM], int mat2[DIM][DIM], int mats[DIM][DIM], int i, int dim){
    int riga[DIM];

    if (i==dim){
        return 0;
    }

    SommaRiga(mat1[i], mat2[i], mats[i], 0, dim);
    SommaMat(mat1, mat2, mats, i+1, dim);
}

int main(){
    int matrice1[DIM][DIM], matrice2[DIM][DIM], dim, matsomma[DIM][DIM];

    printf("Inserire la dimensione della matrice (tra 2 e 100): ");
    scanf("%d", &dim);

    while(dim<2 || dim > 100) {
        printf("Reinserire: ");
        scanf("%d", &dim);
    }

    CreaMat(matrice1, matrice2, dim);
    SommaMat(matrice1, matrice2, matsomma, 0, dim);

    for(int i=0; i<dim; i++){
        for (int j=0; j<dim; j++){
            printf("%d\t", matsomma[i][j]);
        }
        printf("\n");
    }

    return 0;
}