package macchinetta;

import customException.InsufficentChangeException;
import customException.InsufficentValueException;
import customException.InsufficentcoinsException;
import customException.InvalidImportoException;
import customException.InvalidResultException;
import customException.TotalvalueException;

/**
 * {@code RestoMedio} permette la gestione del calcolo di un {@code aggregato} per il resto 
 * Il calcolo viene fatto prendendo le {@code monete} una dall'alto e l'altra dal basso finché non considerano la stessa {@code moneta}
 */
public class RestoMedio implements StrategiaResto{

    /**
     * {@code RestoMedio} è il costruttore della strategia di {@code Resto} personalizzata
     * Tale strategia cerca di non finire tutte le monete del taglio più grande o del taglio più piccolo
     */
    public RestoMedio(){}

    @Override
    public Aggregato Resto(Aggregato cassa, Importo resto) throws InsufficentChangeException, InsufficentValueException, InvalidImportoException, InvalidResultException {
        Aggregato copy = new Aggregato();
        copy.Insert(cassa);

        if (resto.compareTo(cassa.getTotalImporto())>0) throw new InsufficentValueException("quantità insufficente di monete");

        return calcolo(copy, resto);
    }

    private Aggregato calcolo(Aggregato copia, Importo resto) throws InsufficentChangeException, InsufficentValueException, InvalidImportoException, InvalidResultException {
        Aggregato change = new Aggregato();
        Moneta[] tagli = Moneta.values();
        Moneta[] reverse = Moneta.values();

        for (int i =0; i < reverse.length; i++) {
            reverse[i] = reverse[reverse.length - i];
            reverse[reverse.length - i] = tagli[i];
        }

        int i = 0;

        try  {
            while (i<=(tagli.length)/2 && resto.getTotalCents() != 0) {
                if (resto.compareTo(tagli[i].getValue()) > 0) {
                    int quantity = resto.Div(tagli[i].getValue());
                    resto = resto.Sub(tagli[i].getValue().Mul(quantity));
                    copia.Remove(tagli[i], quantity);
                    change.Insert(tagli[i], quantity);
                }

                if (resto.compareTo(reverse[i].getValue()) > 0) {
                    int quantity = resto.Div(reverse[i].getValue());
                    resto = resto.Sub(reverse[i].getValue().Mul(quantity));
                    copia.Remove(reverse[i], quantity);
                    change.Insert(reverse[i], quantity);
                }
            }
        } catch (TotalvalueException e) {
            throw new InsufficentValueException("importo insufficente");
        } catch (InsufficentcoinsException e) {
            throw new InsufficentChangeException("quantità di moneta insufficente");
        }

        return change;
    }
}
