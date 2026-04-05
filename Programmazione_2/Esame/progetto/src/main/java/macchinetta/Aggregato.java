package macchinetta;

import java.util.Map;
import java.util.TreeMap;

import customException.InsufficentcoinsException;
import customException.InvalidImportoException;
import customException.InvalidResultException;
import customException.TotalvalueException;

/**
 * {@code Aggregato} è la rappresentazione di un portafoglio o una cassa con delle {@link Moneta} al suo interno.
 */
public class Aggregato {

    /** 
     * {@code aggregato} è la variabile rappresentativa dell'insieme di monete con la rispettiva quantità che sono salvate in una mappa che le associa
     * L'implementazione della TreeMap è stata consigliata da chatGPT, permette l'ordinamento del contenuto in ordine crescente tramite l'utilizzo di un
     * albero binario di ricerca autobilanciante
     */
    private final TreeMap<Moneta, Integer> aggregato;

    /*
     * Invariante di Rappresentazione (RI);
     *      - Le monete possibili sono solo quelle della classe enum
     *      - La quantità di ogni moneta non può essere negativa
     *      - Una moneta può avere quantità pari a 0, ma non verrà considerata
     *
     * Funzione di Astrazione (AF):
     *      - L'aggregato è la rappresentazione di un insieme di monete con la rispettiva quantità
     * 
     */

    /**
     * {@code Aggregato} è il costruttore dell'oggetto aggregato, tale inizializza una mappa vuota e l'importo totale a 0
     */
    public Aggregato() {
        aggregato = new TreeMap<>();
    }

    /**
     * {@code Insert} permette l'inserimento di una moneta nell'{@code aggregato} tramite una {@code moneta} e una {@code quantità}
     *
     * @param coin è la {@code moneta} che va inserita nell'{@code aggregato}
     * @param quantity è la quantità della {@code moneta} inserita
     * @throws TotalvalueException se la quantità di monete è negativa
     */
    public void Insert(Moneta coin, int quantity) throws TotalvalueException {
        if (quantity < 0) throw new TotalvalueException("quantità negativa");
        aggregato.put(coin, quantity+aggregato.getOrDefault(coin, 0));
    }

    /**
     * {@code Insert} è una seconda versione dell'inserimento all'interno di un {@code aggregato} delle monete.
     * In questo caso viene passato in input un altro {@code aggregato} e viene inserito in quello corrente
     *
     * @param change è l'{@code aggregato} da reinserire in quello corrente
     * @throws TotalvalueException se la quantità di monete è negativa
     */
    public void Insert(Aggregato change) throws TotalvalueException {
        Moneta[] values = Moneta.values();
        
        for (Moneta value : values) {
            Insert(value, change.getQuantity(value));
        }
    }

    /**
     * {@code Remove} è il metodo che permette di rimuovere una quantità specificata di moneta dall'{@code aggregato}
     * 
     * @param coin è la moneta presa in considerazione
     * @param quantity è la quantità che si vuole togliede dall'aggregato
     * @throws TotalvalueException se l'imprto totale che si vuole togliere è maggiore di quello presente nell'{@code aggregato}
     * @throws InsufficentcoinsException se la quantità di moneta nell'{@code aggregato} non è sufficente
     * @throws InvalidImportoException se il calcolo dell'{@code importo} totale non è corretto
     * @throws InvalidResultException se il risultato del calcolo dell'{@code importo} risulta negativa
     */
    public void Remove(Moneta coin, int quantity) throws TotalvalueException, InsufficentcoinsException, InvalidImportoException, InvalidResultException {
        if(aggregato.containsKey(coin) && aggregato.get(coin) >= quantity) {
            aggregato.put(coin, aggregato.get(coin)-quantity);
        } else {
            if (getTotalImporto().getTotalCents() < quantity * coin.getValue().getTotalCents()) throw new TotalvalueException("Valore complessivo non sufficente");
            throw new InsufficentcoinsException("Quantità di moneta insufficente");
        }
    }

    /**
     * {@code Remove} prende in input un aggregato e rimuove le monete di tale {@code aggregato} da quello corrente
     * 
     * @param change è l'{@code aggregato} che va rimosso da quello corrente
     * @throws TotalvalueException se l'imprto totale che si vuole togliere è maggiore di quello presente nell'{@code aggregato}
     * @throws InsufficentcoinsException se la quantità di moneta nell'{@code aggregato} non è sufficente
     * @throws InvalidImportoException se il calcolo dell'{@code importo} totale non è corretto
     * @throws InvalidResultException se il risultato del calcolo dell'{@code importo} risulta negativa
     */
    public void Remove(Aggregato change) throws TotalvalueException, InsufficentcoinsException, InvalidImportoException, InvalidResultException {
        Moneta[] values = Moneta.values();
        Aggregato copia = new Aggregato();
        
        copia.Insert(this);

        for (Moneta value : values) {
            copia.Remove(value, change.getQuantity(value));
        }

        this.Empty();
        this.Insert(copia);
    }

    /**
     * {@code Empty} è un metodo che svuota completamente un aggregato
     */
    public void Empty() {
        aggregato.clear();
    }

    /**
     * {@code getQuantity} restituisce la quantità di una {@code moneta} indicata
     *
     * @param coin è la {@code moneta} di cui si vuole sapere la quantità
     * @return è la quantità di una {@code moneta} indicata
     */
    public int getQuantity(Moneta coin){
        return aggregato.getOrDefault(coin, 0);
    }

    /**
     * {@code getTotalImporto} resituisce l'{@code importo} totale presente nell'{@code aggregato}
     * 
     * @return l'{@code importo} dell'{@code aggregato}
     * @throws InvalidImportoException se l'{@code importo} calcolato risulta non valido
     * @throws InvalidResultException se il risultato è {@code negativo}
     */
    public Importo getTotalImporto() throws InvalidImportoException, InvalidResultException {
        Importo total = new Importo(0);
    
        for (Map.Entry<Moneta, Integer> value : aggregato.entrySet()) {
            total = total.Add(value.getKey().getValue().Mul(value.getValue()));
        }

        return total;
    }

    @Override
    public String toString() {
        StringBuilder aggr = new StringBuilder();

        aggr.append("<");
        for (Map.Entry<Moneta, Integer> entry : aggregato.entrySet()) {
            if (entry.getValue() > 0) aggr.append(entry.getValue()).append(" x ").append(entry.getKey().toString()).append(", ");
        }

        if (aggr.length() > 1) {
            aggr.setLength(aggr.length() - 2);
        }

        aggr.append(">");

        return aggr.toString();
    }
}
