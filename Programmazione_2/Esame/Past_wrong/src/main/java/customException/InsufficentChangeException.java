package customException;

/**
 * Eccezione personalizzata per segnalare l'impossibilità di fornire il resto
 * richiesto.
 * <p>
 * Viene lanciata quando il sistema non dispone di monete o combinazioni
 * sufficienti per restituire il resto necessario al completamento
 * dell'operazione.
 * </p>
 */
public class InsufficentChangeException extends Exception {

    /**
     * Identificatore univoco della versione della classe serializzata.
     * <p>
     * Permette di controllare la compatibilità tra versioni differenti
     * della classe durante la deserializzazione.
     */
    private static final long serialVersionUID = 1L;

    /**
     * Costruisce una nuova {@code InsufficentChangeException} con un messaggio
     * specificato.
     *
     * @param errore messaggio che descrive la causa dell'eccezione
     */
    public InsufficentChangeException(String errore) {
        super(errore);
    }
}
