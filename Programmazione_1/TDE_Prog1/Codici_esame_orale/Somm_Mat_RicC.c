#include <stdio.h>
#include <stdlib.h>

int SommaEl(int riga1[5], int riga2[5], int rigasum[5], int col) {
    if (col == 5){
        return 0;
    }

    rigasum[col] = riga1[col] + riga2[col];
    return SommaEl(riga1, riga2, rigasum, col+1);
}

int Matsomma(int mat1[5][5], int mat2[5][5], int matsum[5][5], int count) {
    if (count == 5){
        return 0;
    }

    int riga1[5];
    SommaEl(mat1[count], mat2[count], riga1, 0);
    for (int i =0; i<5; i++){
        matsum[count][i] = riga1[i];
    }
    count++;
    return Matsomma(mat1, mat2, matsum, count);
}

main() {
    int mat1[5][5] = {{1, 2, 3, 4, 10}, {5, 6, 7, 8, 10}, {9, 1, 2, 3, 10}, {4, 5, 6, 7, 10}, {8, 9, 1, 2, 10}}, mat2[5][5] =  {{8, 9, 1, 2, 10}, {4, 5, 6, 7, 10}, {5, 6, 7, 8, 10}, {9, 1, 2, 3, 10}, {1, 2, 3, 4, 10}}, matsum[5][5];

    Matsomma(mat1, mat2, matsum, 0);

    printf("Matrice somma:\n");
    for (int i=0; i<5; i++){
        for (int j=0; j<5; j++){
            printf("%d ", matsum[i][j]);
        }
        printf("\n");
    }
}