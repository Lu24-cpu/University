#include <stdio.h>
#include <ctype.h>
#include <string.h>

#define DIM 20
#define NUM 3

main(){
    typedef struct{
        char nome[DIM];
        char cognome[DIM];
        int eta;
        struct macchina{
            int cilindrata1;
            char casa_prod1[DIM];
            char modello1[DIM];
            char colore1;
        }autovetture;
        struct moto{
            int cilindrata2;
            char casa_prod2[DIM];
            char modello2[DIM];
            char colore2;
        }ciclomotori;
    }ragazzi;

    ragazzi giovani[NUM];
    int i, num;

    for(i=0; i<NUM; i++){
        printf("Inserire il proprio nome: ");
        gets(giovani[i].nome);

        printf("\nInserire il proprio cognome: ");
        gets(giovani[i].cognome);

        printf("\nInserire la propria eta': ");
        scanf("%d", &giovani[i].eta);
        getchar();

        if(giovani[i].eta>=16){
            printf("\nHai una macchina o una moto? (Seleziona 1 per la macchina, 2 per la moto e 0 se non si ha nulla): ");
            scanf("%d", &num);
            getchar();

            while(num!=1 && num!=2 && num!=0){
                printf("Errore nell'inserimento, riprova: ");
                scanf("%d", &num);
                getchar();
            }

            switch(num){
                case 1:
                    printf("\n\nInserire la casa prodruttrice della macchina: ");
                    gets(giovani[i].autovetture.casa_prod1);

                    printf("\nInserire il modello della macchina: ");
                    gets(giovani[i].autovetture.modello1);

                    printf("\nInserire la cilindrata della macchina: ");
                    scanf("%d", &giovani[i].autovetture.cilindrata1);

                    printf("\nInserire il colore della macchina: ");
                    gets(giovani[i].autovetture.colore1);

                    break;
                case 2:
                    printf("\n\nInserire la casa prodruttrice della moto: ");
                    gets(giovani[i].ciclomotori.casa_prod2);

                    printf("\nInserire il modello della moto: ");
                    gets(giovani[i].ciclomotori.modello2);

                    printf("\nInserire la cilindrata della moto: ");
                    scanf("%d", &giovani[i].ciclomotori.cilindrata2);
                    getchar();

                    printf("\nInserire il colore della moto: ");
                    gets(giovani[i].ciclomotori.colore2);

                    break;
                case 0:
                    printf("\n\nMi dispiace tu non abbia ancora nessun mezzo\n\n");
                    break;
            }
        }
        printf("\n");
    }

    printf("\n\nI ragazzi che sono passati oggi sono:\n");

    for(i=0; i<10; i++){
        puts(giovani[i].nome);
        printf(" ");
        puts(giovani[i].cognome);
        printf("che ha %d anni. ", giovani[i].eta);
        if(num==1)
            printf("E possiede: %s %s con una cilindrata di: %d e color: %s", giovani[i].autovetture.casa_prod1, giovani[i].autovetture.modello1, giovani[i].autovetture.cilindrata1, giovani[i].autovetture.colore1);
        else
            printf("E possiede: %s %s con una cilindrata di: %d e color: %s", giovani[i].ciclomotori.casa_prod2, giovani[i].ciclomotori.modello2, giovani[i].ciclomotori.cilindrata2, giovani[i].ciclomotori.colore2);
    }
}
