package macchinetta;

import customException.InvalidImportoException;
import customException.InvalidResultException;
import customException.TotalvalueException;

/**
 * {@code StrategiaResto} è un intergaffia che rappresenta il contratto per la gestione dei resti di un {@code distributore automatico}
 */
public interface StrategiaResto {

    /**
     * {@code Resto} è un metodo che presi in input due {@code aggregati}, la cassa e uno dove depositare il resto e l'{@code importo} del resto
     * e restituisce {@code aggregato} corrispondente al resto effettivo
     * 
     * @param cassa è l'{@code aggregato} corrispondente alla cassa del {@code distributore}
     * @param resto è l'{@code importo} che rappresenta il resto che deve essere calcolato
     * @return l'{@code aggregato} corrispondente al resto calcolato, in base alla strategia utilizzata
     * @throws TotalvalueException se non c'è abbastanza {@code Moneta} di un taglio in cassa
     * @throws InvalidImportoException se l'{@code importo} calcolato non è valido
     * @throws InvalidResultException se non c'è abbastanza {@code resto} in {@code cassa}
     */
    public Aggregato Resto(Aggregato cassa, Importo resto) throws TotalvalueException, InvalidResultException, InvalidImportoException;

}
