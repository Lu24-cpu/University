package macchinetta;

import java.util.List;

import customException.InsufficentChangeException;
import customException.InsufficentValueException;
import customException.InsufficentcoinsException;
import customException.InvalidImportoException;
import customException.InvalidResultException;
import customException.TotalvalueException;

/**
 * {@code Change} è una classe astratta che implementa l'interfaccia del resto ({@link StrategiaResto})
 * 
 */
public abstract class Change implements StrategiaResto {

    /**
     * {@code Change} è il costruttore della classe astratta che calcola il resto in formato {@code importo}
     */
    public Change() {}

    @Override
    public Aggregato Resto(Aggregato cassa, Importo resto) throws InsufficentChangeException, InsufficentValueException, InvalidImportoException, InvalidResultException {
        Aggregato change = new Aggregato();
        Aggregato copy = new Aggregato();
        List<Moneta> register = getOrder();

        try {
            copy.Insert(cassa);
            if (resto.compareTo(cassa.getTotalImporto())>0) throw new InsufficentValueException("quantità insufficente di monete");

            for (Moneta value : register) {
                if (resto.compareTo(value.getValue()) >= 0) {
                    int quantity = resto.Div(value.getValue());
                    
                    if (quantity > copy.getQuantity(value)) quantity = copy.getQuantity(value);

                    resto = resto.Sub(value.getValue().Mul(quantity));
                    
                    copy.Remove(value, quantity);
                    change.Insert(value, quantity);
                }
            }
            
            if (resto.getTotalCents() != 0) {
                if (cassa.getTotalImporto().getTotalCents() > resto.getTotalCents()) {
                    throw new InsufficentChangeException("Resto non disponibile");
                } else {
                    throw new InsufficentValueException("Importo nella cassa insufficente");
                }
            }

            cassa.Empty();
            cassa.Insert(copy);

        } catch(InsufficentcoinsException e) {
            throw new InsufficentChangeException("quantità di moneta insufficente");
        } catch(TotalvalueException e) {
            throw new InsufficentValueException("quantità di moneta insufficente");
        }
    
        return change;
    }

    /**
     * {@code getOrder} restituisce la mappa dell'{@code aggregato} cassa nell'ordine selezionato dalla strategia
     * 
     * @param cassa è l'{@code aggregato} di cui viene restituita la mappa
     * @return la mappa della {@code cassa} in ordine basato sulla strategia
     */
    public abstract List<Moneta> getOrder();

}
