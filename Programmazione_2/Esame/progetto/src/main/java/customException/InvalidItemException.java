package customException;

/**
 * Eccezione personalizzata per segnalare elementi non validi.
 * <p>
 * Viene lanciata quando un oggetto, elemento o dato fornito
 * non rispetta i requisiti richiesti oppure non è supportato
 * dal sistema.
 * </p>
 */
public class InvalidItemException extends Exception {

    /**
     * Identificatore univoco della versione della classe serializzata.
     * <p>
     * Permette di controllare la compatibilità tra versioni differenti
     * della classe durante la deserializzazione.
     */
    private static final long serialVersionUID = 1L;

    /**
     * Costruisce una nuova {@code InvalidItemException} con un messaggio
     * specificato.
     *
     * @param errore messaggio che descrive la causa dell'eccezione
     */
    public InvalidItemException(String errore) {
        super(errore);
    }
}
