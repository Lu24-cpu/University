package macchinetta;

import java.util.Map;

import customException.InsufficentChangeException;
import customException.InsufficentValueException;
import customException.InsufficentcoinsException;
import customException.InvalidImportoException;
import customException.InvalidResultException;
import customException.TotalvalueException;

/**
 * {@code Change} è una classe astratta che implementa l'interfaccia del resto ({@code StrategiaResto})
 * 
 */
public abstract class Change implements StrategiaResto {

    /*
     * Invariante di Rappresentazione (RI): 
     *      - Il resto non può essere negativo
     * 
     * Funzione di Astrazione (AF):
     *      - La funzione Resto restituisce un aggregato vuoto in quanto il calcolo viene fatto da RestoAlto e RestoBasso
     *      - La funzione Change calcola il resto in base alla moneta e alla quantità passata in input
     * 
     */

    /**
     * {@code Change} è il costruttore della classe astratta che calcola il resto in formato {@code importo}
     */
    public Change() {}

    @Override
    public Aggregato Resto(Aggregato cassa, Importo resto) throws InsufficentChangeException, InsufficentValueException, InvalidImportoException, InvalidResultException {
        Aggregato change = new Aggregato();
        Aggregato copy = new Aggregato();
        Map<Moneta, Integer> register = getOrder(cassa);

        copy.Insert(cassa);

        if (resto.compareTo(cassa.getTotalImporto())>0) throw new InsufficentValueException("quantità insufficente di monete");

        try {
            for (Map.Entry<Moneta, Integer> value : register.entrySet()) {
                if (resto.compareTo(value.getKey().getValue()) >= 0) {
                    int quantity = resto.Div(value.getKey().getValue());
                    resto = resto.Sub(value.getKey().getValue().Mul(quantity));
                    copy.Remove(value.getKey(), quantity);
                    change.Insert(value.getKey(), quantity);
                }
            }
            
            if (resto.getTotalCents() != 0) {
                if (cassa.getTotalImporto().getTotalCents() > resto.getTotalCents()) {
                    throw new InsufficentChangeException("Resto non disponibile");
                } else {
                    throw new InsufficentValueException("Importo nella cassa insufficente");
                }
            }

            cassa.Remove(change);
        } catch (TotalvalueException e) {
            throw new InsufficentValueException("importo insufficente");
        } catch (InsufficentcoinsException e) {
            throw new InsufficentChangeException("quantità di moneta insufficente");
        }
        
        return change;
    }

    /**
     * {@code getOrder} restituisce la mappa dell'{@code aggregato} cassa nell'ordine selezionato dalla strategia
     * 
     * @param cassa è l'{@code aggregato} di cui viene restituita la mappa
     * @return la mappa della {@code cassa} in ordine basato sulla strategia
     */
    public abstract Map<Moneta, Integer> getOrder(Aggregato cassa);

}
