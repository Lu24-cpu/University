package customException;

/**
 * Eccezione personalizzata per segnalare errori relativi alla gestione
 * delle monete.
 * <p>
 * Viene lanciata quando si verifica una condizione non valida durante
 * operazioni che coinvolgono monete, come valori non ammessi o monete
 * non supportate dal sistema.
 * </p>
 */
public class MonetaException extends Exception {

    /**
     * Identificatore univoco della versione della classe serializzata.
     * <p>
     * Permette di controllare la compatibilità tra versioni differenti
     * della classe durante la deserializzazione.
     */
    private static final long serialVersionUID = 1L;

    /**
     * Costruisce una nuova {@code MonetaException} con un messaggio
     * specificato.
     *
     * @param errore messaggio che descrive la causa dell'eccezione
     */
    public MonetaException(String errore) {
        super(errore);
    }
}
