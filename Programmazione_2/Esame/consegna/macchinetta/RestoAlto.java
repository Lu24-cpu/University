package macchinetta;

import java.util.Map;

/**
 * {@code RestoAlto} è una classe che estende quella di change e gestisce il resto partendo dalle {@code moneta} con valore più alto a quella più bassa 
 */
public class RestoAlto extends Change {

    /*
     * Invariante di Rappresentazione (RI): 
     *      - Il resto in Importo non deve essere negativo
     *      - La cassa non deve essere vuota
     *      - L'aggregato change deve essere vuoto
     * 
     * Funzione di Astrazione (AF):
     *      - Viene convertito il resto in aggregato se possibile
     *      - Viene restituito l'aggregato del resto
     * 
     */

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
