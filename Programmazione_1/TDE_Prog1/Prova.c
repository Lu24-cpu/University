#include <stdio.h>
#include <stdlib.h>

#define DIM 100

int Creamat1(int dim, int mat[DIM][DIM]){
    int i, j;
    for (i=0; i<dim; i++){
        for (j=0; j<dim; j++){
            mat[i][j]= (i+1)*(j+1);
        }
    }
    return mat[DIM][DIM];
}

int Creamat2(int dim, int mat[DIM][DIM]){
    int i, j;
    for (i=dim-1; i>=0; i--){
        for (j=dim-1; j>=0; j--){
            mat[i][j]= (i+1)*(j+1);
        }
    }
    return mat[DIM][DIM];
}

int SommaMat(int mat1[DIM][DIM], int mat2[DIM][DIM], int matsum[DIM][DIM], int i, int j, int dim){
    if (i == dim-1 && j == dim-1){
        matsum[i][j] = mat1[i][j]+mat2[i][j];
        return matsum[DIM][DIM];
    } else if(j == dim-1){
        matsum[i][j] = mat1[i][j]+mat2[i][j];
        return SommaMat(mat1, mat2, matsum, i+1, 0, dim);
    }

    matsum[i][j] = mat1[i][j]+mat2[i][j];
    return SommaMat(mat1, mat2, matsum, i, j+1, dim);
}

int main(){
    int matrice1[DIM][DIM], matrice2[DIM][DIM], matsum[DIM][DIM], dim;

    printf("Inserire la dimensione della matrice (tra 2 e 100): ");
    do{
        scanf("%d", &dim);
    }while (dim<2 || dim>100);

    Creamat1(dim, matrice1);
    Creamat2(dim, matrice2);
    SommaMat(matrice1, matrice2, matsum, 0, 0, dim);

    for (int i=0; i<dim; i++){
        for (int j=0; j<dim; j++){
            printf("%d\t", matsum[i][j]);
        }
        printf("\n");
    }

    return 0;
}
