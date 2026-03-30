package customException;

/**
 * Eccezione personalizzata per segnalare risultati non validi.
 * <p>
 * Viene lanciata quando il risultato di un'operazione non rispetta
 * i vincoli previsti, risulta incoerente oppure non è accettabile
 * nel contesto dell'applicazione.
 * </p>
 */
public class InvalidResultException extends Exception {

    /**
     * Identificatore univoco della versione della classe serializzata.
     * <p>
     * Permette di controllare la compatibilità tra versioni differenti
     * della classe durante la deserializzazione.
     */
    private static final long serialVersionUID = 1L;

    /**
     * Costruisce una nuova {@code InvalidResultException} con un messaggio
     * specificato.
     *
     * @param errore messaggio che descrive la causa dell'eccezione
     */
    public InvalidResultException(String errore) {
        super(errore);
    }
}
