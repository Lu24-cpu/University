package customException;

/**
 * {@code AggregatoException} rappresenta un'eccezione personalizzata utilizzata
 * per segnalare errori relativi alla gestione di un aggregato di monete.
 * <p>
 * Questa eccezione può essere lanciata nei casi:
 * <ol>
 *     <li>Errori generato cercando di inserire una moneta inesistente nell'aggregato</li>
 *     <li>Incoerenze tra le monete presenti e il totale rappresentato dall'importo</li>
 * </ol>
 */
public class InsufficentcoinsException extends Exception {

    /**
     * Identificatore univoco della versione della classe serializzata.
     * <p>
     * Permette di controllare la compatibilità tra versioni differenti
     * della classe durante la deserializzazione.
     */
    private static final long serialVersionUID = 1L;
    
    /**
     * Costruisce una nuova {@code InsufficentcoinsException} con un messaggio specificato.
     *
     * @param errore messaggio che descrive la causa dell'eccezione
     */
    public InsufficentcoinsException(String errore) {
        super(errore);
    }
}
