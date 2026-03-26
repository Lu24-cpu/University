package customException;

/**
 * Eccezione personalizzata per segnalare errori relativi alla gestione degli
 * slot.
 * <p>
 * Questa eccezione viene lanciata quando si verifica una condizione non valida
 * durante operazioni che coinvolgono slot, come accessi non consentiti,
 * valori errati o slot non disponibili.
 *
 */
public class SlotException extends Exception {

    /**
     * Identificatore univoco della versione della classe serializzata.
     * <p>
     * Permette di controllare la compatibilità tra versioni differenti
     * della classe durante la deserializzazione.
     */
    private static final long serialVersionUID = 1L;

    /**
     * Costruisce una nuova {@code SlotException} con un messaggio
     * specificato.
     *
     * @param errore messaggio che descrive la causa dell'eccezione
     */
    public SlotException(String errore) {
        super(errore);
    }
}
