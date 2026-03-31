#include <stdio.h>
#include <stdlib.h>
#include <string.h>
#include <ctype.h>

#define DIM 20

main(){
    typedef struct{
        char specie[DIM];
        char razza[DIM];
        int zampe;
        char nome_animale[DIM];
        char cognome[DIM];
    }veterinario;

    int i, num, max, point;
    veterinario animali[DIM];

    printf("Quanti animali sono arrivati in giornata: ");
    scanf("%d", &num);
    getchar();
    for(i=0; i<num; i++){
        printf("\n");
        printf("Animale %d", i+1);
        printf("\nInserire la specie: ");
        gets(animali[i].specie);
        printf("\nInserire la razza: ");
        gets(animali[i].razza);
        printf("\nInserire il numero di zampe: ");
        scanf("%d", &animali[i].zampe);
        getchar();
        printf("\nInserire il nome dell'animale: ");
        gets(animali[i].nome_animale);
        printf("\nInserire il nome del proprietario: ");
        gets(animali[i].cognome);
        if(i==0){
            max=animali[i].zampe;
            point=i;
        }
        if(max<animali[i].zampe){
            max=animali[i].zampe;
            point=i;
        }
    }
    printf("\n\nL'animale che aveva piu' zampe, tra quelli visitati era:\n");
    printf("\nUn %s %s con %d zampe e il suo nome era: %s\n\nIl cognome del proprietario era: %s", animali[point].specie, animali[point].razza, animali[point].zampe, animali[point].nome_animale, animali[point].cognome);
}
