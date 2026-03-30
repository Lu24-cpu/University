package macchinetta;

import customException.CapacityException;
import customException.EmptyRailException;
import customException.InvalidItemException;
import customException.TagliaException;

/**
 * {@code Binario} è la classe che rappresenta un binario di un {@code distibutore automatico}.
 */
public class Binario {

    /** {@code size} è la taglia del {@code binario} */
    private final Taglia size;
    /** {@code product} è il prodotto associato al {@code binario} */
    private Prodotto product;
    /** {@code capacity} è la capacità massima del {@code binario} */
    private final int capacity;
    /** {@code quantity} è la quantità effettiva di prodotto presente nel {@code binario} */
    private int quantity;

    /*
     * Invariante di Rappresentazione (RI):
     *      - La capacità non può essere negativa
     *      - La quantità di prodotto non può essere negativa
     *      - La taglia deve rispettare quelle possibili (S, M, L o XL)
     * 
     * Funzione di Astrazione (AF):
     *      - Un binario è una tupla formata da taglia, capacità, prodotto e quantità
     * 
     */

    /**
     * {@code Bianrio} è il costruttore del tipo {@code binario} e lo costrisce in base alla {@code taglia} e alla {@code capacità} date in input
     * 
     * @param taglia è la taglia del {@code binario}
     * @param capacità è la capatità complessiva del {@code binario}
     */
    public Binario(Taglia taglia, int capacità) {
        size = taglia;
        product = new Prodotto("", taglia, new Importo(0));
        capacity = capacità;
        quantity = 0;
    }

    /**
     * {@code uploadRail} è un metodo che permette di caricare su un binario un {@code prodotto} con una data {@code quantità}
     * 
     * @param prodotto è il {@code prodotto} da inserire nel {@code binario}
     * @param quantità è la quantità di {@code prodotto} da inserire nel {@code binario}
     * @throws TagliaException se la taglia del {@code prodotto} non è minore o uguale a quella del {@code binario}
     * @throws InvalidItemException se il {@code prodotto} non è lo stesso di quello già presente nel {@code binario}
     * @throws CapacityException se la {@code quantità} inserita con quella già presente sfora la {@code capacità} massima
     */
    public void uploadRail(Prodotto prodotto, int quantità) throws TagliaException, InvalidItemException, CapacityException {
        if(size.compareTo(prodotto.size())<0) throw new TagliaException("Taglia non conforme");

        if(!product.equals(prodotto) && !product.equals(new Prodotto("", size, new Importo(0)))) throw new InvalidItemException("item");
        if(quantità + quantity > capacity) throw new CapacityException("capacity");

        if(product.equals(new Prodotto("", size, new Importo(0)))) {
            product = prodotto;
        }

        quantity += quantità;
    }

    /**
     * {@code unloadRail} è un metodo che scarica una data quantità di prodotto da un {@code binario}
     * 
     * @param quantità è la quantità di prodotto che si vuole rimuovede dal {@code binario}
     * @throws CapacityException se si cerca di rimuovere più prodotti dal {@code binario} rispetto a quelli presenti.
     * @throws EmptyRailException se il {@code binario} è vuoto
     */
    public void unloadRail(int quantità) throws CapacityException, EmptyRailException {
        if(quantity == 0) throw new EmptyRailException("Il binario è vuoto");
        if(quantity < quantità) throw new CapacityException("Richiesta di scarico binario eccessivo");
    
        quantity -= quantità;
    }

    /**
     * {@code getProduct} è un metodo che restituisce il prodotto del {@code binario} corrente
     * 
     * @return il {@code prodotto} del {@code binario}
     */
    public Prodotto getProduct() {
        return product;
    }

    /**
     * {@code getQuantity} restituisce la quantità di {@code prodotto} presente nel {@code binario}
     * 
     * @return la quantità di {@code prodotto}
     */
    public int getQuantity(){
        return quantity;
    }

    /**
     * {@code getCapacity} restituisce la {@code capacità} massima del {@code binario}
     * 
     * @return l'intero rappresentativo della {@code capacità}
     */
    public int getCapacity() {
        return capacity;
    }

    @Override
    public String toString() {
        StringBuilder binario = new StringBuilder();

        binario.append("<");

        if (product.equals(new Prodotto("", size, new Importo(0)))) binario.append("-, ");
        else binario.append(product.toString()).append(", ");

        binario.append(this.size).append(", ").append(this.quantity).append(", ").append(this.capacity).append(">");

        return binario.toString();
    }
}
