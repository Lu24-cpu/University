#include <stdio.h>
#include <stdbool.h>

main(){
    int sumall, allergy[8]={1, 2, 4, 8, 16, 32, 64, 128};
    bool allergie[8] = {false, false, false, false, false, false, false, false};
    char allergia[8][12] = {"eggs", "penats", "shellfish", "strawberries", "tomatoes", "chocolate", "pollen", "cats"};

    printf("Inserire la somma delle proprie allergie: ");
    scanf("%d", &sumall);

    if (sumall>=256){
        sumall -= 256;
    }

    for (int i=7; i>=0; i--){
        if (sumall>=allergy[i] && sumall>0){
            allergie[i]=true;
            sumall-=allergy[i];
        } else {

        }
    }

    for (int i=0; i<8; i++){
        if (allergie[i]==true){
            for (int j=0; j<12; j++){
                printf("%c", allergia[i][j]);
            }
            printf(" (%d)\n", allergy[i]);
        }
    }
}