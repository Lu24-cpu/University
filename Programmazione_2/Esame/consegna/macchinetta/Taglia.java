package macchinetta;

import customException.TagliaException;

/**
 * {@code Taglia} è una classe enum che rappresenta le taglie possibili di un {@code Prodotto} o {@code Binario}
 */
public enum Taglia {

    /** {@code Small} è la taglia più piccola */
    Small("S"),
    /** {@code Medium} è la taglia media */
    Medium("M"),
    /** {@code Large} è la taglia larga */
    Large("L"),
    /** {@code ExtraLarge} è la raglia più grande */
    ExtraLarge("XL");

    /** È la stringa rappresentativa della taglia */
    private final String size;

    /*
     * Invariante di rappresentazione (RI):
     *      - La taglia non può essere un intero
     *      - La taglia non può essere diversa da S, M, L o XL
     * 
     * Funzione di astrazione (AF):
     *      - Presa una stringa restituisce l'oggetto taglia corrispondente
     * 
     */

    /**
     * {@code Taglia} è il costruttore del tipo taglia
     * 
     * @param taglia è la stringa rappresentativa della {@code taglia}
     */
    private Taglia(String taglia) {
        size = taglia;
    }

    /**
     * {@code getTaglia} restituisce una delle variabili del tipo {@code Taglia} dalla stringa data in input
     * 
     * @param size è la stringa rappresentativa della {@code taglia}
     * @return l'oggetto del tipo taglia corrispondente alla stringa in input
     * @throws TagliaException se la stringa data in input non è conforme ad una delle taglie possibili
     */
    public static Taglia taglia(String size) throws TagliaException {
        for(Taglia taglia : Taglia.values()) {
            if (size.equals(taglia.size)) return taglia;
        }

        throw new TagliaException("invalid-size");
    }

    @Override
    public String toString() {
        return size;
    }
}
