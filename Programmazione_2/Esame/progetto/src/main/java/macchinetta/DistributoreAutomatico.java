package macchinetta;


import java.util.List;

import customException.CapacityException;
import customException.EmptyRailException;
import customException.InsufficentChangeException;
import customException.InsufficentValueException;
import customException.InvalidImportoException;
import customException.InvalidItemException;
import customException.InvalidResultException;
import customException.SlotException;
import customException.TagliaException;

/**
 * {@code DistributoreAutomatico} è la classe che gestisce tutte le possibili funzioni di un effettivo distributore automatico
 */
public class DistributoreAutomatico {

    /** {@code rails} è una lista dei {@code binari} del distributore */
    private final List<Binario> rails;
    /** {@code strategy} è la strategia di calcolo del {@code resto} del distributore */
    private final StrategiaResto strategy;
    /** {@code exchange} è l'{@code aggregato} che rappresenta la cassa del distributore */
    private final Aggregato cashier;

    /*
     * Invariante di Rappresentazione (RI):
     *      - La strategia deve essere una tra Media, High o Low
     *      - Exchange deve essere un aggregato valido
     * 
     * Funzione di Astrazione (AF):
     *      - Un distributore automatico è un insieme di funzioni che permettono di caricare/scaricare un binario o restituire cosa c'è sui binari
     * 
     */

    /**
     * {@code DistributoreAutomatico} è il costruttore del distributore.
     * Tale costruttore inizializza la lista dei binari con quella data in input, come per strategia e cassa
     * 
     * @param cassa è l'{@code aggregato} rappresentativo della cassa del distributore
     * @param strategia è la {@code strategia del resto} che il distributore utilizzerà
     * @param binari è la lista dei binari disponibili nel distributore
     */
    public DistributoreAutomatico(Aggregato cassa, StrategiaResto strategia, List<Binario> binari) {
        rails = binari;
        strategy = strategia;
        cashier = cassa;
    }

    /**
     * {@code caricaBianrio} è un metodo che permette di caricare un prodotto su un binario, data la stringa del prodotto in input
     * 
     * @param product è il {@code prodotto} che si vuole caricare nel {@code distributore}
     * @param quantity è la quantità di prodotto che si vuole caricare
     * @return la stringa rappresentativa della quantità di {@code prodotto} che non è stata inserita
     * @throws SlotException se nel binario scelto non può essere caricato tale prodotto
     */
    public int caricaBinario(Prodotto product, int quantity) throws SlotException {
        for (Binario rail : rails) {
            try {
                int finale = rail.getCapacity() - rail.getQuantity();

                if(quantity <= finale) {
                    rail.uploadRail(product, quantity);
                    quantity = 0;
                    break;
                } else {
                    rail.uploadRail(product, finale);
                    quantity -= finale;
                }
            } catch (CapacityException | InvalidItemException | TagliaException e) {}
        }
        
        return quantity;
    }

    /**
     * {@code scaricaBinario} è un metodo che permette di rimuovere un singolo {@code prodotto} da un binario, diminuentdone la quantità di 1.
     * Il {@code binario} scelto e l'{@code importo} inserito vengono passati in input sotto forma di stringa e viene poi gestita la divisione della stringa
     * 
     * @param rail è l'indice del {@code binario} da cui prelevare il {@code prodotto}
     * @param importo è l'{@code importo} inserito nel distributore
     * @return la rappresentaizone in stringa dell'{@code aggregato} del {@code resto} dovuto
     * @throws SlotException se il binario scelto non esiste
     * @throws EmptyRailException se il binario è vuoto
     * @throws InsufficentChangeException se non è stato possbile calcolare il {@code resto} per mancanza di {@code monete}
     * @throws InsufficentValueException se non è stato possbile calcolare il {@code resto} per {@code importo} insufficente in cassa
     */
    public Aggregato scaricaBinario(int rail, Importo importo) throws SlotException, EmptyRailException, InsufficentChangeException, InsufficentValueException {
        if (rail > this.rails.size()) throw new SlotException("Il binario non esiste");
        if (this.rails.get(rail).getProduct() == null || this.rails.get(rail).getQuantity() == 0) throw new EmptyRailException("Il Binario è vuoto");
        //Riprendere da qui
        try {
            Prodotto prodotto = this.rails.get(rail).getProduct();
            Importo resto = importo.Sub(prodotto.value());
            Aggregato changeAgg = strategy.Resto(cashier, resto);
            
            rails.get(rail).unloadRail(1);
            return changeAgg;
        } catch (InsufficentChangeException e) {
            throw new InsufficentChangeException("resto non riuscito");
        } catch (InsufficentValueException | CapacityException | InvalidImportoException | InvalidResultException | NumberFormatException e) {
            throw new InsufficentValueException("resto non disponibile");
        }
    }

    /**
     * {@code getAggregato} ritorna l'{@code aggregato} di cassa del {@code distributore}
     * 
     * @return l'{@code aggregato} della cassa
     */
    public Aggregato getAggregato() {
        return cashier;
    }

    @Override
    public String toString() {
        StringBuilder menu = new StringBuilder();

        for (int i = 0; i < this.rails.size(); i++) {
            if (!this.rails.get(i).getProduct().equals(new Prodotto("", rails.get(i).getProduct().size(), new Importo(0)))) {
                String[] prodotto = rails.get(i).getProduct().toString().replace("<", "").replace(">", "").split("\\, ");
                
                try {
                    menu.append("? ").append(i).append(" | ").append(prodotto[0]).append(" | ").append(prodotto[1]).append("\n");
                } catch (Exception e) {
                    System.out.println("Errore prodotto");
                }
            }
        }

        menu.setLength(menu.length() - 1);

        return menu.toString();
    }
}
