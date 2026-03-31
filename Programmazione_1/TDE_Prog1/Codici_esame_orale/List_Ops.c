#include <stdio.h>
#include <stdlib.h>
#include <string.h>

#define DIM 100

typedef struct {
    char Nome[DIM];
    struct Lista *Next;
}Lista;

/*void Append(Lista *Linked, Lista *Nodo) {
    Lista *t = Linked;
    printf("Entra append ok\n");
    
    do {
        t = t->Next;
        printf("Passaggio Nodi ok\n, prossimo: %p\n", t->Next);
    }while(t->Next == NULL);

    printf("Aggiunta nodo\n");
    t->Next = &Nodo;
    printf("Aggiunto\n");
}*/

void Append(Lista **Linked, Lista *Nodo) {
    printf("Entra append ok\n");
    
    if (*Linked == NULL) {  // Se la lista è vuota, il nuovo nodo diventa la testa
        *Linked = Nodo;
    } else {
        Lista *t = *Linked;
        while (t->Next != NULL) {
            *t = t->Next;
            printf("Passaggio Nodi ok\n, prossimo: %p\n", t->Next);
        }

        printf("Aggiunta nodo\n");
        t->Next = Nodo; // Aggiungi il nodo alla fine della lista
        printf("Aggiunto\n");
    }

}

void Concatenate(Lista *Link1, Lista *Link2) {
    Lista *t = Link1;
    
    do {
        t = t->Next;
    }while(t->Next != NULL);

    t->Next = Link2;
}

void Filter(Lista *Principale, char Item[DIM], Lista *Reciver) {
    Lista *t = Principale, *Items = Reciver;
    int counter=0;

    do {
        if (strstr(t->Nome, Item)!=NULL){
            if (counter==0)
                Items=t;
            else
                Items->Next = t;
            counter++;
        }
        t = t->Next;
    }while(t->Next != NULL);

}

int Lenth(Lista *Link) {
    Lista *t = Link;
    int count =0;

    do {
        t = t->Next;
        count++;
    }while(t->Next != NULL);

    return count;
}

int main(){
    Lista *Lincata1 = NULL, *Lincata2 = NULL, *Ricerca = NULL;
    char Name[DIM], Item[DIM];
    int count = 0;

    do {
        Lista *Nodo = (Lista*)malloc(sizeof(Lista));
        Nodo->Next = NULL;

        printf("Inserire il nome del Nodo:\n");
        gets(&Name);


        for (int i=0; i<strlen(Name); i++){
            Nodo->Nome[i] = Name[i];
        }
        printf("Copia nodo ok\n");

        Append(&Lincata1, Nodo);
        printf("Aggiunta el lista\n");

    }while(Name== "End");

    count = 0;

    do {
        Lista *Nodo = (Lista*)malloc(sizeof(Lista));
        Nodo->Next = NULL;

        printf("Inserire il nome del Nodo: ");
        gets(&Name);


        for (int i=0; i<strlen(Name); i++){
            Nodo->Nome[i] = Name[i];
        }

        Append(Lincata2, Nodo);
        
    }while(Name== "End");

    Lista *Stampa1 = Lincata1, *Stampa2 = Lincata2, *RicercaSt = Ricerca;

    printf("Prima lista:\n");
    while(Stampa1!=NULL){
        printf("%s\n", Stampa1->Nome);
        Stampa1 = Stampa1->Next;
    }
    printf("Lunghezza: %d\n", Lenth(Lincata1));


    printf("\nSeconda lista:\n");
    while(Stampa2!=NULL){
        printf("%s\n", Stampa2->Nome);
        Stampa2 = Stampa2->Next;
    }
    printf("Lunghezza: %d\n", Lenth(Lincata2));

    Concatenate(Lincata1, Lincata2);
    printf("\nLunghezza concatenare: %d\n", Lenth(Lincata1));

    printf("Inserire l'elemento da cercare: ");
    gets(&Item);

    Filter(Lincata1, Item, Ricerca);

    printf("\nRicerca elementi:\n");
    while(RicercaSt!=NULL){
        printf("%s\n", RicercaSt->Nome);
        RicercaSt = RicercaSt->Next;
    }
    printf("Lunghezza: %d\n", Lenth(Ricerca));
}