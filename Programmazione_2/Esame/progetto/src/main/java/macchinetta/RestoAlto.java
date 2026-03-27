package macchinetta;

import java.util.Map;

/**
 * {@code RestoAlto} è una classe che estende quella di change e gestisce il resto partendo dalle {@code moneta} con valore più alto a quella più bassa 
 */
public class RestoAlto extends Change {

    /**
     * {@code RestoAlto} è il costruttore della strategia di resto dall'alto
     * Data la cassa ritorna la mappa ordinata in modo decrescente
     */
    public RestoAlto() {}

    @Override
    public Map<Moneta, Integer> getOrder(Aggregato cassa) {  
        return cassa.getAggregato().descendingMap(); 
    }
    
}
