package customException;

/**
 * Eccezione personalizzata che viene sollevata quando il valore totale di
 * un'operazione non soddisfa determinate condizioni predefinite.
 * 
 */
public class TotalvalueException extends Exception {

    /**
     * Identificatore univoco della versione della classe serializzata.
     * <p>
     * Permette di controllare la compatibilità tra versioni differenti
     * della classe durante la deserializzazione.
     */
    private static final long serialVersionUID = 1L;

    /**
     * Costruisce una nuova {@code TotalvalueException} con un messaggio
     * specificato.
     *
     * @param errore messaggio che descrive la causa dell'eccezione
     */
    public TotalvalueException(String errore) {
        super(errore);
    }
}
