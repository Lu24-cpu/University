package macchinetta;

import java.util.Map;

/**
 * {@code RestoBasso} è la classe che implementa la strategia del resto che predilige l'uso delle monete di piccolo taglio per calcolare il resto
 */
public class RestoBasso extends Change {

    /**
     * {@code RestoBasso} è il costruttore della strategia di resto dal basso
     * Data la cassa ritorna la mappa ordinata in modo crescente
     */
    public RestoBasso() {}

    @Override
    public Map<Moneta, Integer> getOrder(Aggregato cassa) {
        return cassa.getAggregato();
    }

}
