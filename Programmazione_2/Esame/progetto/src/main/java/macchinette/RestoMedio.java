package macchinetta;

import java.util.Map;

import customException.*;
import clients.Parser;

/**
 * {@code RestoMedio} permette la gestione del calcolo di un {@code aggregato} per il resto 
 * Il calcolo viene fatto prendendo le {@code monete} una dall'alto e l'altra dal basso finché non considerano la stessa {@code moneta}
 */
public class RestoMedio implements StrategiaResto{

    /*
     * Invariante di Rappresentazione (RI):
     *      - L'importo del resto dato in input non può essere negativo
     *      - L'aggregato cassa data in input non può essere vuota
     *      - L'aggregato del resto (change) deve essere vuoto per calcolare correttamente il resto
     * 
     * Funzione di Astrazione (AF):
     *      - Il resto vine calcolato in base ai tagli possibili e in modo tale da non consumare solo le monete di grosso taglio o piccolo taglio
     * 
     */

    /**
     * {@code RestoMedio} è il costruttore della strategia di {@code Resto} personalizzata
     * Tale strategia cerca di non finire tutte le monete del taglio più grande o del taglio più piccolo
     */
    public RestoMedio(){}

    @Override
    public Aggregato Resto(Aggregato cassa, Importo resto) throws InsufficentChangeException, InsufficentValueException, InvalidImportoException, InvalidResultException {
        Aggregato change = new Aggregato();
        int[] tagli = {200, 100, 50, 20, 10, 5, 2, 1 };
        int k = 0, j = tagli.length;

        while (k <= j) {

            if (k != j) {
                if (resto.getTotalCents() >= tagli[k]) {
                    int i = 1;

                    try {
                        Importo inizialvalue = Parser.parseImporto(Integer.toString(tagli[k]));
                        Importo calcoloresto = Parser.parseImporto("0");

                        Moneta coin = Moneta.moneta(new Importo(tagli[k]));
                        while (calcoloresto.getTotalCents() <= resto.getTotalCents()) {
                            if (i == cassa.getAggregato().get(coin))
                                break;
                            i++;
                            calcoloresto = calcoloresto.Add(inizialvalue);
                        }

                        if (calcoloresto.getTotalCents() <= resto.getTotalCents()) {
                            i--;
                            calcoloresto = calcoloresto.Sub(inizialvalue);
                        }

                        resto = resto.Sub(calcoloresto);

                        if (i != 0) {
                            change.Insert(coin, cassa.getAggregato().get(coin) - i);
                        }
                    } catch (InvalidImportoException | InvalidResultException | MonetaException e) {
                        throw new IllegalArgumentException("Occurred this error: " + e.getMessage());
                    }
                }

                if (resto.getTotalCents() >= tagli[j]) {
                    int i = 1;

                    try {
                        Importo inizialvalue = Parser.parseImporto(Integer.toString(tagli[j]));
                        Importo calcoloresto = Parser.parseImporto("0");
                        Moneta coin = Moneta.moneta(new Importo(tagli[j]));
                        while (calcoloresto.getTotalCents() <= resto.getTotalCents()) {
                            if (i == cassa.getAggregato().get(coin))
                                break;
                            i++;
                            calcoloresto = calcoloresto.Add(inizialvalue);
                        }

                        if (calcoloresto.getTotalCents() <= resto.getTotalCents()) {
                            i--;
                            calcoloresto = calcoloresto.Sub(inizialvalue);
                        }

                        resto.Sub(calcoloresto);

                        if (i != 0) {
                            change.Insert(coin, cassa.getAggregato().get(coin) - i);
                        }
                    } catch (InvalidImportoException | InvalidResultException | MonetaException e) {
                        throw new IllegalArgumentException("Occurred this error:" + e.getMessage());
                    }
                }

                k++;
                j--;
            } else {
                if (resto.getTotalCents() >= tagli[k]) {
                    int i = 1;

                    try {
                        Importo inizialvalue = Parser.parseImporto(Integer.toString(tagli[k]));
                        Importo calcoloresto = Parser.parseImporto("0");

                        Moneta coin = Moneta.moneta(new Importo(tagli[k]));
                        while (calcoloresto.getTotalCents() <= resto.getTotalCents()) {
                            if (i == cassa.getAggregato().get(coin))
                                break;
                            i++;
                            calcoloresto = calcoloresto.Add(inizialvalue);
                        }

                        if (calcoloresto.getTotalCents() <= resto.getTotalCents()) {
                            i--;
                            calcoloresto = calcoloresto.Sub(inizialvalue);
                        }

                        resto.Sub(calcoloresto);

                        if (i != 0) {
                            change.Insert(coin, cassa.getAggregato().get(coin) - i);
                        }
                    } catch (InvalidImportoException | InvalidResultException | MonetaException e) {
                        throw new IllegalArgumentException("Occurred this error:" + e.getMessage());
                    }
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
}
