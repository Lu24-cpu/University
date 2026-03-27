package macchinetta;

import java.util.Map;

import clients.Parser;
import customException.InsufficentChangeException;
import customException.InsufficentValueException;
import customException.InvalidImportoException;
import customException.InvalidResultException;
import customException.MonetaException;

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

        while (k <= j) {

            if (k != j) {
                if (resto.compareTo(tagli[k].getValue()) > 0) {
                    int quantity = resto.Div(tagli[k].getValue());
                    resto = resto.Sub(tagli[k].getValue().Mul(quantity));
                    copy.Remove(tagli[k], quantity);
                    change.Insert(tagli[k], quantity);
                }

                if (resto.compareTo(tagli[j].getValue()) > 0) {
                    int quantity = resto.Div(tagli[j].getValue());
                    resto = resto.Sub(tagli[j].getValue().Mul(quantity));
                    copy.Remove(tagli[j], quantity);
                    change.Insert(tagli[j], quantity);
                }

                k++;
                j--;
            } else {
                if (resto.compareTo(tagli[k].getValue()) > 0) {
                    int quantity = resto.Div(tagli[j].getValue());
                    resto = resto.Sub(tagli[j].getValue().Mul(quantity));
                    copy.Remove(tagli[j], quantity);
                    change.Insert(tagli[j], quantity);
                }

                break;
            }
        }

        if (resto.getTotalCents() != 0) {
            if (cassa.getTotalImporto().getTotalCents() > resto.getTotalCents()) {
                for (Map.Entry<Moneta, Integer> coin : change.getAggregato().entrySet()) {
                    try {
                        cassa.Insert(coin.getKey(), coin.getValue());
                    } catch (InvalidImportoException | InvalidResultException e) {
                        System.out.println("Occurred this error: " + e.getMessage());
                    }
                }

                throw new InsufficentChangeException("Resto non disponibile");
            } else {
                for (Map.Entry<Moneta, Integer> coin : change.getAggregato().entrySet()) {
                    try {
                        cassa.Insert(coin.getKey(), coin.getValue());
                    } catch (InvalidImportoException | InvalidResultException e) {
                        System.out.println("Occurred this error: " + e.getMessage());
                    }
                }

                throw new InsufficentValueException("Importo nella cassa insufficente");
            }

        }

        return change;
    }

    private Aggregato calcolo(Aggregato copia, Importo resto) {
        Aggregato change = new Aggregato();
        Moneta[] tagli = Moneta.values();
        
    }
}
