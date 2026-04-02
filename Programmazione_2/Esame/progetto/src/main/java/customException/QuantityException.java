package customException;

/**
 * Eccezione personalizzata per segnalare errori relativi al superamento
 * della capacità consentita.
 * <p>
 * Viene lanciata quando una struttura, contenitore o sistema raggiunge
 * o supera il limite massimo di elementi o risorse gestibili.
 * </p>
 */
public class QuantityException extends Exception {

    /**
     * Identificatore univoco della versione della classe serializzata.
     * <p>
     * Permette di controllare la compatibilità tra versioni differenti
     * della classe durante la deserializzazione.
     */
    private static final long serialVersionUID = 1L;

    /**
     * Costruisce una nuova {@code QuantityException} con un messaggio specificato.
     *
     * @param errore messaggio che descrive la causa dell'eccezione
     */
    public QuantityException(String errore) {
        super(errore);
    }
}
