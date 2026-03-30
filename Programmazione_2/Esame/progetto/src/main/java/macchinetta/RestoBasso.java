package macchinetta;

import java.util.ArrayList;
import java.util.List;

/**
 * {@code RestoBasso} è la classe che implementa la strategia del resto che predilige l'uso delle {@link Moneta} di piccolo taglio per calcolare il resto
 */
public class RestoBasso extends Change {

    /**
     * {@code RestoBasso} è il costruttore della strategia di resto dal basso
     * Data la cassa ritorna la mappa ordinata in modo crescente
     */
    public RestoBasso() {}

    @Override
    public List<Moneta> getOrder() {
        return new ArrayList<Moneta>(List.of(Moneta.values()));
    }

}
