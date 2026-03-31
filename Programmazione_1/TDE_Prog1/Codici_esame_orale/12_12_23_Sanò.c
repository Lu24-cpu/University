#include <stdio.h>
#include <ctype.h>
#include <string.h>

#define DIM 20
#define PROF 5

main(){
    typedef struct {
        char materia[DIM];
        char cognome_prof[DIM];
        struct alunni{
            char cognome[DIM];
            float voto;
            int classe;
            char sezione;
        }studenti[30];
    }insegnante;

    insegnante prof[PROF];
    int i, num, j;

    for(i=0; i<PROF; i++){
        printf("Inserire il cognome del professore: ");
        gets(prof[i].cognome_prof);

        printf("Inserire la materia che insegna: ");
        gets(prof[i].materia);

        printf("Inserire il numero di alunni (tra 2 e 30): ");
        scanf("%d", &num);
        getchar();
        while(num<2||num>30){
            printf("Reinserire: ");
            scanf("%d", &num);
            getchar();
        }

        for(j=0; j<num; j++){
            printf("\nInserire il cognome dello studente: ");
            gets(prof[i].studenti[j].cognome);

            printf("\nInserire la classe dello studente (1, 2 o 3): ");
            scanf("%d", &prof[i].studenti[j].classe);
            getchar();
            while(prof[i].studenti[j].classe!=1 && prof[i].studenti[j].classe!=2 && prof[i].studenti[j].classe!=3){
                printf("\nReinserire: ");
                scanf("%d", &prof[i].studenti[j].classe);
                getchar();
            }

            printf("Inserire la sezione dello studente (A o B): ");
            prof[i].studenti[j].sezione=getchar();
            while(prof[i].studenti[j].sezione!='A' && prof[i].studenti[j].sezione!='B'){
                printf("\nReinserire: ");
                prof[i].studenti[j].sezione=getchar();
            }

            printf("\nInserire il voto dello studente (tra 4 e 10): ");
            scanf("%f", &prof[i].studenti[j].voto);
            getchar();
            while(prof[i].studenti[j].voto<4 || prof[i].studenti[j].voto>10){
                printf("Reinserire: ");
                scanf("%f", &prof[i].studenti[j].voto);
                getchar();
            }
        }

        printf("\n\n");
    }

    for(i=0; i<PROF; i++){
        printf("\n\nCognome insegnante: %s\n", prof[i].cognome_prof);
        printf("Materia che insegna: %s\n", prof[i].materia);
        printf("Alunni:\n");

        for(j=0; j<num; j++){
            printf("\n%d Alunno:\n", j+1);
            printf("Cognome: %s\n", prof[i].studenti[j].cognome);
            printf("Classe: %d%c\n", prof[i].studenti[j].classe, prof[i].studenti[j].sezione);
            printf("Voto: %.2f\n", prof[i].studenti[j].voto);
        }
        printf("\n");
    }
}
