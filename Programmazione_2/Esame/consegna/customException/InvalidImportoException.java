package customException;

/**
 * Eccezione personalizzata per segnalare errori relativi a importi non validi.
 * <p>
 * Viene lanciata quando un valore monetario risulta negativo, nullo
 * quando non consentito, oppure non compatibile con i vincoli
 * dell'operazione richiesta.
 * </p>
 */
public class InvalidImportoException extends Exception {

    /**
     * Identificatore univoco della versione della classe serializzata.
     * <p>
     * Permette di controllare la compatibilità tra versioni differenti
     * della classe durante la deserializzazione.
     */
    private static final long serialVersionUID = 1L;

    /**
     * Costruisce una nuova {@code InvalidImportoException} con un messaggio
     * specificato.
     *
     * @param errore messaggio che descrive la causa dell'eccezione
     */
    public InvalidImportoException(String errore) {
        super(errore);
    }
}
