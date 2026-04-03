package customException;

/**
 * Eccezione personalizzata per segnalare un valore insufficiente.
 * <p>
 * Viene lanciata quando il valore fornito non è sufficiente per
 * completare un'operazione richiesta, come un pagamento o
 * il raggiungimento di una soglia minima.
 * </p>
 */
public class InsufficentValueException extends Exception {

    /**
     * Identificatore univoco della versione della classe serializzata.
     * <p>
     * Permette di controllare la compatibilità tra versioni differenti
     * della classe durante la deserializzazione.
     */
    private static final long serialVersionUID = 1L;

    /**
     * Costruisce una nuova {@code InsufficentValueException} con un messaggio
     * specificato.
     *
     * @param errore messaggio che descrive la causa dell'eccezione
     */
    public InsufficentValueException(String errore) {
        super(errore);
    }
}
