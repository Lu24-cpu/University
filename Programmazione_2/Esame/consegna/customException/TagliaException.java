package customException;

/**
 * Eccezione personalizzata per segnalare errori relativi a dimensioni non
 * valide.
 * <p>
 * Viene lanciata quando una struttura, collezione o valore numerico
 * presenta una dimensione non consentita o non compatibile con
 * l'operazione richiesta.
 * </p>
 */
public class TagliaException extends Exception {

    /**
     * Identificatore univoco della versione della classe serializzata.
     * <p>
     * Permette di controllare la compatibilità tra versioni differenti
     * della classe durante la deserializzazione.
     */
    private static final long serialVersionUID = 1L;

    /**
     * Costruisce una nuova {@code TagliaException} con un messaggio specificato.
     *
     * @param errore messaggio che descrive la causa dell'eccezione
     */
    public TagliaException(String errore) {
        super(errore);
    }
}
