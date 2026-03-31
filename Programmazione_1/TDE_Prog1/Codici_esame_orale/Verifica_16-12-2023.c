/*Sanò*/
#include <stdio.h>
#include <string.h>
#include <ctype.h>

#define DIM 20

main(){
    typedef struct{
        char squadra[DIM];
        int punti;
        int Goalscored;
        int Goalsubiti;
        struct s_allenatore{
            char cognome[DIM];
            int eta;
        }coach;
    }campionato;

    campionato squadre[5];
    int num, i, maxpunti, maxdif, first=0;

    printf("Quante squadre si vogliono inserire (tra 2 e 5)? ");
    scanf("%d", &num);
    getchar();

    while(num<2 || num>5){
        printf("Errore nell'inserimento del numero di squadre, riprova: ");
        scanf("%d", &num);
        getchar();
    }

    printf("\n\n");

    for(i=0; i<num; i++){
        printf("Inserire il nome della squadra che si vuole inserire: ");
        gets(squadre[i].squadra);

        printf("Inserire il punteggio della squadra: ");
        scanf("%d", &squadre[i].punti);
        getchar();

        printf("Inserire il numero di Goal segnati: ");
        scanf("%d", &squadre[i].Goalscored);
        getchar();

        while(squadre[i].Goalscored<0){
            printf("Errore nell'inserimento del numero di goal segnati, non possono essere negativi, riprova: ");
            scanf("%d", &squadre[i].Goalscored);
            getchar();
        }

        printf("Inserire il numero di Goal subiti: ");
        scanf("%d", &squadre[i].Goalsubiti);
        getchar();

        while(squadre[i].Goalsubiti<0){
            printf("Errore nell'inserimento del numero di goal subiti, non possono essere negativi, riprova: ");
            scanf("%d", &squadre[i].Goalsubiti);
            getchar();
        }

        printf("Inserire il cognome dell'allenatore: ");
        gets(squadre[i].coach.cognome);

        printf("Inserire l'eta' dell'allenatore: ");
        scanf("%d", &squadre[i].coach.eta);
        getchar();

        while(squadre[i].coach.eta<=40 || squadre[i].coach.eta>=65){
            printf("Sei sicuro che sia l'eta' giusta del coach? Riprova: ");
            scanf("%d", &squadre[i].coach.eta);
            getchar();
        }

        printf("\n\n");
    }

    for(i=0; i<num; i++){
        printf("%d. %s\n", i+1, squadre[i].squadra);
        printf("\tLa squadra ha totalizzato %d punti\n", squadre[i].punti);
        printf("\tIl numero di goal segnati e' stato di: %d\n", squadre[i].Goalscored);
        printf("\tMentre il numero di goal subiti e' stato: %d\n", squadre[i].Goalsubiti);
        printf("\tInfine la differenza di Goal segnati e subiti e': %d\n", squadre[i].Goalscored-squadre[i].Goalsubiti);
        printf("\tL'allenatore della squadra e': %s e ha %d anni\n", squadre[i].coach.cognome, squadre[i].coach.eta);

        if(i==0){
            maxpunti=squadre[i].punti;
            maxdif=squadre[i].Goalscored-squadre[i].Goalsubiti;
        }
        if(maxpunti<squadre[i].punti){
            maxpunti=squadre[i].punti;
            first=i;
        }
        if(maxpunti==squadre[i].punti && maxdif<squadre[i].Goalscored-squadre[i].Goalsubiti){
            maxpunti=squadre[i].punti;
            maxdif=squadre[i].Goalscored-squadre[i].Goalsubiti;
            first=i;
        }
    }

    num=squadre[first].Goalscored-squadre[first].Goalsubiti;
    printf("\nLa squadra che ha vinto e': %s con %d punti e una differenza di goal fatti e subiti di: %d\n", squadre[first].squadra[DIM], squadre[first].punti, num);
}
