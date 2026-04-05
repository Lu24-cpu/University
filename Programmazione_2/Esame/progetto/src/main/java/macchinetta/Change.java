package macchinetta;

import java.util.List;

import customException.*;

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
    public Aggregato Resto(Aggregato cassa, Importo resto) throws InsufficentChangeException, InsufficentValueException, InvalidResultException, InvalidImportoException {
        Aggregato change = new Aggregato();
        Aggregato copy = new Aggregato();
        List<Moneta> register = getOrder();

        try {
            copy.Insert(cassa);
            if (resto.compareTo(cassa.getTotalImporto())>0) throw new InvalidResultException("quantità insufficente di monete");

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
                    throw new InvalidResultException("Importo nella cassa insufficente");
                }
            }

            cassa.Empty();
            cassa.Insert(copy);
            return change;
        } catch(InsufficentcoinsException e) {
            throw new InsufficentChangeException("quantità di moneta insufficente");
        } catch(TotalvalueException e) {
            throw new InsufficentValueException("quantità di moneta insufficente");
        }
    }

    /**
     * {@code getOrder} restituisce una lista delle {@code monete} in base ad un ordine specificato
     * 
     * @return la lista delle {@code monete} in ordine basato sulla strategia
     */
    public abstract List<Moneta> getOrder();

}
