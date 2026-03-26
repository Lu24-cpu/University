package macchinetta;

import customException.*;

import java.util.Map;

import clients.Parser;

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
        Map<Moneta, Integer> register = getOrder(cassa);

        for (Map.Entry<Moneta, Integer> value : register.entrySet()) {
            if (resto.compareTo(value.getKey().getValue()) >= 0) {
                int i = 0;
                StringBuilder calcresto = new StringBuilder();
                calcresto.append((float) value.getKey().getValue().getTotalCents() / 100);

                try {
                    Importo inizialvalue = Parser.parseImporto(calcresto.toString());
                    Importo calcoloresto = Parser.parseImporto("0");

                    while (calcoloresto.compareTo(resto) <= 0) {
                        if (i == value.getValue()) break;
                        i++;
                        calcoloresto = calcoloresto.Add(inizialvalue);
                    }

                    if (calcoloresto.compareTo(resto) > 0) {
                        i--;
                        calcoloresto = calcoloresto.Sub(inizialvalue);
                    }

                    resto = resto.Sub(calcoloresto);

                    if (i != 0) {
                        change.Insert(value.getKey(), i);
                    }
                } catch (InvalidImportoException | InvalidResultException e) {
                    if (e.getMessage().contains("valido")) throw new InsufficentChangeException("Resto non disponibile");
                    else throw new InsufficentValueException("Importo nella cassa insufficente");
                }
            }
        }

        if (resto.getTotalCents() != 0) {
            if (cassa.getTotalImporto().getTotalCents() > resto.getTotalCents()) {
                for (Map.Entry<Moneta, Integer> coin : change.getAggregato().entrySet()) {
                    cassa.Insert(coin.getKey(), coin.getValue());
                }

                throw new InsufficentChangeException("Resto non disponibile");
            } else {
                for (Map.Entry<Moneta, Integer> coin : change.getAggregato().entrySet()) {
                    cassa.Insert(coin.getKey(), coin.getValue());
                }

                throw new InsufficentValueException("Importo nella cassa insufficente");
            }

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
