#include <stdio.h>
#include <ctype.h>
#include <string.h>
#include <math.h>

#define DIM 100

int Conversione(char inserimento[DIM], int len, int i){
    int num = 0;
    char numeri[10] = {'0', '1', '2', '3', '4', '5', '6', '7', '8', '9'};
    if (i == 0){
        for (int j=0; j < 10; j++){
            if (inserimento[0]==numeri[j]) {
                return j*pow(10, (len));
            }
        }
    }

    for (int j=0; j<10; j++){
        if (inserimento[len-i]==numeri[j]){
            num += j*pow(10, i);
            char numero[i];

            for (int k=0; k<i-1; k++){
                numero[k]=inserimento[k];
            }
            num+=Conversione(numero, len, i-1);
            return num; 
        }
    }
}

int main(){
    char inserimento[DIM];

    printf("Inserire un numero da al massimo 100 cifre: ");
    scanf("%s", inserimento);

    printf("Il numero inzialmente inserito (stringa) era: %s, il numero convertito in intero è: %d\n", inserimento, Conversione(inserimento, strlen(inserimento), 0));
}