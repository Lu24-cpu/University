#include <stdio.h>
#include <stdlib.h>
#include <string.h>
#include <ctype.h>

#define DIM 200

struct shoes {
    char marca[DIM];
    char modello[DIM];
    char taglia[DIM];
    char colore[DIM];
    int prezzo;
};

int insert(){
    int choice=1;
    struct shoes scarpa;
    FILE *scarpe;

    scarpe = fopen("modelli.txt", "a");

    while(choice!=0){
        printf("\nInserisci il brand della scarpa: ");
        gets(scarpa.marca);

        printf("\nInserisci il modello: ");
        gets(scarpa.modello);

        printf("\nInserisci la taglia: ");
        scanf("%d", &scarpa.taglia);
        getchar();

        printf("\nInserisci il colore dalla scarpa: ");
        gets(scarpa.colore);

        printf("\nInserire il prezzo a cui si vuole vendere le scarpe: ");
        scanf("%d", &scarpa.prezzo);
        getchar();

        fprintf(scarpe, "%s %s %d %s %d\n", scarpa.marca, scarpa.modello, scarpa.taglia, scarpa.colore, scarpa.prezzo);
        printf("\nprocessing...\n\n");

        printf("Se vuoi termire la ricerca inserisci 0: ");
        scanf("%d", &choice);
        getchar();
    }

    fclose(scarpe);
    return 0;
}

int search(){
    int choice2=0, grandezza;
    struct shoes scarpa;
    char seller[DIM];
    FILE *scarpe;

    scarpe = fopen("modelli.txt", "r");

    while(choice2 != 3){
        printf("\nNegozio di scarpe\n");
        printf("1. Cerca il brand\n");
        printf("2. Cerca modello colore e taglia\n");
        printf("3. Esci\n");
        printf("Inerisci una scelta: ");
        scanf("%d", &choice2);
        getchar();

        switch (choice2) {
            case 1:
                printf("Inserisci il brand di una scarpa che vuoi cercare: ");
                gets(scarpa.marca);

                while(fgets(seller, sizeof(seller), scarpe) != NULL) {
                    if(strstr(seller, scarpa.marca) !=NULL)
                        printf("Modello trovato: %s\n", seller);
                    else{
                        printf("Scarpa non trovata");
                        return 0;
                    }
                }

                break;
            case 2:
                printf("Inserisci il modello della scarpa: ");
                gets(scarpa.modello);

                printf("\nInserisci la taglia desiderata: ");
                scanf("%d", &grandezza);
                getchar();

                printf("\nInserisci il colore desiderato: ");
                gets(scarpa.colore);

                sprintf(scarpa.taglia, "%d", grandezza);

                while(fgets(seller, sizeof(seller), scarpe) != NULL) {
                    if(strstr(seller, scarpa.modello) !=NULL && strstr(seller, scarpa.taglia) != NULL && strstr(seller, scarpa.colore) !=NULL)
                        printf("Modello trovato: %s\n", seller);
                    else{
                        printf("\nScarpa non trovata");
                        return 0;
                    }
                }

                break;
            case 3:
                printf("Uscita in corso");
                break;
            default:
                printf("Scelta non valida. Riprova.\n");
        }
    }
    fclose(scarpe);
    return 0;
}

int show(){
    struct shoes scarpa;
    FILE *scarpe;
    char seller[DIM];
    int choice3;

    scarpe = fopen("modelli.txt", "r");

    while(fgets(seller, sizeof(seller), scarpe)!=NULL){
        puts(seller);
    }
    if(fgets(seller, sizeof(seller), scarpe)==NULL)
        printf("No more shoes\n");

    fclose(scarpe);
    return 0;
}

int main() {
    char seller[DIM];
    FILE *scarpe;
    int choice1;
    struct shoes scarpa;

    printf("Benvenuti nel nostro negozio di scarpe\n");

    while(choice1!=4){
        printf("\n\nSeleziona una delle scelte del menu':\n");
        printf("1. Inserisci le scarpe che vuoi vendere:\n");
        printf("2. Cerca uno specifico modello presente sul sito\n");
        printf("3. Visualizza tutti i prodotti sul sito\n");
        printf("4. Uscita dal sito\n");
        printf("\n\nInserisci una scelta: ");
        scanf("%d", &choice1);
        getchar();

        switch(choice1){
            case 1:
                insert();
                break;

            case 2:
                search();
                break;

            case 3:
                show();
                break;

            case 4:
                printf("\n\nGrazie e arrivederci\n\n");
                break;

            default:
                printf("\nNumero inserito non valito");
        }
    }
}
