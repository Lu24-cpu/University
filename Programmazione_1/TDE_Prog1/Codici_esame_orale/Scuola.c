#include <stdio.h>
#include <stdlib.h>
#include <string.h>
#include <ctype.h>

struct Studente {
    char Nome[50];
    int Voti[8][11];
};

struct Anni {
    struct Studente Alunni[30];
};

void CalcoloMedie(struct Anni Scuola[5]) {
    float sum=0;

    for(int i=0; i<5; i++){
        for (int j=0; j<30; j++){
            for (int k=0; k<8; k++){
                for (int n=0; n<11; n++){
                    if (n == 10){
                        sum/=10;
                        Scuola[i].Alunni[j].Voti[k][n] = sum;
                    }
                    sum += Scuola[i].Alunni[j].Voti[k][n];
                }
                sum = 0;
            }
        }
    }
}

int main() {
    char Nome[50];
    int year = 0;
    struct Anni Collegio[5];

    do{
        printf("Inserire gli studenti dell'anno %d (End per terminare se meno di 30 studenti in aula)\n", year);
        for(int student = 0; student<30; student++){
            gets(Nome);

            if (Nome == "End"){
                break;
            }

            strcpy(Collegio[year].Alunni[student].Nome, Nome);

            printf("Inserire i 10 voti per materia di %s\n", Nome);

            for (int i=0; i<8; i++){
                printf("%desima Materia\n", i+1);
                for (int j=0; j<10; j++){
                    scanf("%d", &Collegio[year].Alunni[student].Voti[i][j]);
                    if (i<7 && j<9){
                        gets();
                    }
                }
            }
        }

        year++;
    }while (year != 5);

    CalcoloMedie(Collegio);
    
    for(int i=0; i<5; i++){
        for (int j=0; j<30; j++){
            printf("%s\n", Collegio[i].Alunni[j].Nome);
            for (int k=0; k<8; k++){
                for (int n=0; n<11; n++){
                    printf("%d\t", Collegio[i].Alunni[j].Voti[k][n]);
                }
                printf("\n");
            }
        }
    }
}