package customException;

/**
 * Eccezione personalizzata per segnalare che una guida, binario o contenitore
 * risulta vuoto.
 * <p>
 * Viene lanciata quando si tenta di eseguire un'operazione su una struttura
 * che non contiene elementi, rendendo impossibile completare l'operazione
 * richiesta.
 * </p>
 */
public class EmptyRailException extends Exception {

    /**
     * Identificatore univoco della versione della classe serializzata.
     * <p>
     * Permette di controllare la compatibilità tra versioni differenti
     * della classe durante la deserializzazione.
     */
    private static final long serialVersionUID = 1L;

    /**
     * Costruisce una nuova {@code EmptyRailException} con un messaggio
     * specificato.
     *
     * @param errore messaggio che descrive la causa dell'eccezione
     */
    public EmptyRailException(String errore) {
        super(errore);
    }
}
