package macchinetta;

import java.util.ArrayList;
import java.util.Collections;
import java.util.List;

/**
 * {@code RestoAlto} è una classe che estende quella di change e gestisce il resto partendo dalle {@link Moneta} con valore più alto a quella più bassa 
 */
public class RestoAlto extends Change {

    /**
     * {@code RestoAlto} è il costruttore della strategia di resto dall'alto
     * Data la cassa ritorna la mappa ordinata in modo decrescente
     */
    public RestoAlto() {}

    @Override
    public ArrayList<Moneta> getOrder() {
        ArrayList<Moneta> values = new ArrayList<>(List.of(Moneta.values()));
        Collections.reverse(values);
        return values;
    }
    
}
