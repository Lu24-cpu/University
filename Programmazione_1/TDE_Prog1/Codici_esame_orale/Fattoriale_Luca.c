#include <stdio.h>

int Fattoriale(int n) {
    if (n <=1) {
        return 1;
    }
    return n*Fattoriale(n-1);
}

int main() {
    int num;

    printf("Inserire un valore (minore di 17) di cui fare il fattoriale: ");
    scanf("%d", &num);

    while (num>=17){
        printf("Reinserire un numero minore di 17: ");
        scanf("%d", &num);
    }
    printf("Il fattoriale di %d è: %d\n", num, Fattoriale(num));
}