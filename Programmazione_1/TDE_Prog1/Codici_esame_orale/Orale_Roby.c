#include <stdio.h>
#include <stdlib.h>
#include <ctype.h>
#include <string.h>

#define DIM 100

main(){
    int sel = 0, count=0, len;
    char parole[5][DIM], parola;

    do{
        printf("Inserire 0 per inserire una serie di parole oppure 1 per terminare: ");
        scanf("%d", &sel);
        gets();

        switch (sel){
            case 0:
                for (int i=0; i<5; i++){
                    printf("Inserire l'%desima parola: ", i+1);
                    gets(parole[i]);
                }
                for (int i=0; i<=4; i++){
                    if (parole[i][0]>parole[i+1][0]) {
                        if (strlen(parole[i])>strlen(parole[i+1])){
                            len = strlen(parole[i]);
                        } else {
                            len = strlen(parole[i+1]);
                        }
                        for (int j=0; j<len; j++){
                            parola=parole[i][j];
                            parole[i][j]=parole[i+1][j];
                            parole[i+1][j]=parola;
                        }
                    }
                    if (parole[i][0]==parole[i+1][0]){
                        if (strlen(parole[i])>strlen(parole[i+1])){
                            len = strlen(parole[i]);
                        } else {
                            len = strlen(parole[i+1]);
                        }
                        for (int j=0; j<len; j++){
                            if (parole[i][j]>parole[i+1][j]){
                                count = 1;
                                break;
                            }
                        }
                        if (count == 1){
                            for (int j=0; j<strlen(parole[i]); j++){
                                parola=parole[i][j];
                                parole[i][j]=parole[i+1][j];
                                parole[i+1][j]=parola;
                            }
                        }
                    }
                }
                
                for (int i=0; i<5; i++){
                    for (int k=0; k<5; k++){
                        for (int j=0; j<strlen(parole[k]); j++){
                            printf("%c", parole[k][j]);
                        }
                        printf(" ");
                    }

                    printf("\n");
                }

                break;
            case 1:
                break;
            default:
                printf("Inserimento errato ritentaren\n");
        }
    }while (sel != 1);
}