package macchinetta;

import customException.MonetaException;

/**
 * {@code Moneta} è una classe enum che rappresenta tutte le monete possibili in formato {@code Importo}
 */
public enum Moneta {
    
    /** {@code Cent} è la rappresentazione del singolo centesimo */
    Cent(new Importo(1)),
    /** {@code Cents2} è la rappresentazione dei due centesimi */
    Cents2(new Importo(2)),
    /** {@code Cents5} è la rappresentazione dei cinque centesimi */
    Cents5(new Importo(5)),
    /** {@code Cents10} è la rappresentazione dei dieci centesimi */
    Cents10(new Importo(10)),
    /** {@code Cents20} è la rappresentazione dei venti centesimi */
    Cents20(new Importo(20)),
    /** {@code Cents50} è la rappresentazione dei venti centesimi */
    Cents50(new Importo(50)),
    /** {@code Unit} è la rappresentazione della singola unità */
    Unit(new Importo(100)),
    /** {@code Units} è la rappresentazione delle due unità*/
    Units2(new Importo(200));

    /** {@code value} è la rappresentazione della moneta presa in considerazione */
    private final Importo value;

    /*
     * Invariante di Rappresentazione (RI):
     *      - Una moneta non può avere valore negativo
     *      - Una moneta non può avere valore diverso da 1, 2, 5, 10, 20, 50, 100 e 200 cents
     * 
     * Funzione di Astrazione (AF):
     *      - La funzione getMoneta restituisce uno delle possibili monete 
     *      - Una moneta è la rappresentazione di una singola moneta nel formato importo
     * 
     */


    /**
     * {@code Moneta} è il costruttore privato del tipo {@code Moneta}
     * 
     * @param valore è l'oggetto {@code Importo} che rappresenta la moneta
     */
    private Moneta(Importo valore){
        this.value = valore;
    }

    /** 
     * {@code getMoneta} prende in input un {@code Importo} che rappresentano una moneta e restituiscono un tipo {@code Moneta}
     * 
     * @param value è l'{@code importo} rappresentativo della moneta
     * @return la moneta designata in input
     * @throws MonetaException se il {@code valore} dato in input non è una delle monete possibili
     */
    public static Moneta moneta(Importo value) throws MonetaException{
        for (Moneta coin : Moneta.values()) {
            if(coin.value.equals(value)) return coin;
        }

        throw new MonetaException("invalid-amount");
    }

    /**
     * {@code getValue} ritorna l'{@code importo} della moneta corrente
     * 
     * @return il valore della moneta corrente vista in {@code importo}
     */
    public Importo getValue() {
        return value;
    }

    @Override
    public String toString() {
        return value.toString();
    }
}
