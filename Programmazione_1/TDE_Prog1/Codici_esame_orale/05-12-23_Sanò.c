#include <stdio.h>
#include <string.h>
#include <ctype.h>

#define DIM 100

struct stadio{
    int settore;
    int anello;
    char posto[10];
    char colore;
    char cognome[20];
};

main(){
    struct stadio biglietti[DIM];
    int num, i, j=0;

    printf("All the seats are free (100).\n\n");
    for(i=0; i+j<DIM; i++){
        printf("Insert the number of tikets that you want to buy: ");
        scanf("%d", &num);
        getchar();
        while(j<=num-1){
            printf("Enter the sector where you want to buy the seat (1 or 2): ");
            scanf("%d", &biglietti[i+j].settore);
            getchar();
            while(biglietti[i+j].settore!=1 && biglietti[i+j].settore!=2){
                printf("Retry: ");
                scanf("%d", &biglietti[i+j].settore);
                getchar();
            }

            printf("Enter the ring where you want to buy the seat (1 or 2): ");
            scanf("%d", &biglietti[i+j].anello);
            while(biglietti[i+j].anello!=1 && biglietti[i+j].anello!=2){
                printf("Retry: ");
                scanf("%d", &biglietti[i+j].anello);
                getchar();
            }

            switch(biglietti[i+j].anello){
                case 1:
                    if(biglietti[i+j].settore==1)
                        printf("The seats that you can buy are from A1 to A10 and from B1 to B15\n ");
                    else
                        printf("The seats that you can buy are from C1 to C10 and from D1 to D15\n ");
                    break;
                case 2:
                    if(biglietti[i+j].settore==1)
                        printf("The seats that you can buy are from A10 to A20 and from B15 to B30\n ");
                    else
                        printf("The seats that you can buy are from C10 to C20 and from D15 to D30\n ");
                    break;
            }

            printf("\nInsert the seat: ");
            gets(biglietti[i+j].posto);

            printf("Choose the color of the seat (B, N, R): ");
            scanf("%c", &biglietti[i+j].colore);

            j++;
        }
    }
}
