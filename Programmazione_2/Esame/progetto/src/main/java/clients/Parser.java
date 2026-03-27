package clients;

import java.math.BigDecimal;
import java.util.ArrayList;

import customException.EmptyRailException;
import customException.InsufficentChangeException;
import customException.InsufficentValueException;
import customException.InsufficentcoinsException;
import customException.InvalidImportoException;
import customException.InvalidResultException;
import customException.MonetaException;
import customException.SlotException;
import customException.TagliaException;
import customException.TotalvalueException;
import macchinetta.Aggregato;
import macchinetta.Binario;
import macchinetta.DistributoreAutomatico;
import macchinetta.Importo;
import macchinetta.Moneta;
import macchinetta.Prodotto;
import macchinetta.RestoAlto;
import macchinetta.RestoBasso;
import macchinetta.RestoMedio;
import macchinetta.StrategiaResto;
import macchinetta.Taglia;

/**
 * Questa classe, chiamata {@code Parser}, viene usata per convertire stringhe in oggetti del tipo:
 * <ol>
 *      <li>{@code Moneta} che è la rappresentazione delle monete possibili</li>
 *      <li>{@code Aggregato} che è un insieme di monete con le loro quantità</li>
 *      <li>{@code Importo} che è un valore in centesimi e unità salvate in variabili intere</li>
 *      <li>{@code Binario} che è un oggetto composto da prodotto, taglia, quantià e capacità</li>
 *      <li>{@code Prodotto} che è un oggetto che rappresenta un elemento composto da nome, taglie e importo</li>
 * </ol>
 */
public class Parser {

    /**
     * È il costruttore della classe parser
     */
    private Parser() {
    }

    /**
     * {@code parseMoneta} converte una stringa data in input ({@code value}) in un qualcosa del tipo {@code Moneta}
     * 
     * @param value è la stringa rappresentativa della moneta da convertire
     * @return l'oggetto {@code Moneta} rappresentativo della stringa data in input
     * @throws MonetaException se la stringa data in input non è una moneta valida
     */
    public static Moneta parseMoneta(String value) throws MonetaException {
        BigDecimal centsvalue = (new BigDecimal(value)).multiply(BigDecimal.valueOf(100));

        if (centsvalue.scale() > 2) {
            throw new MonetaException("Invalid-input");
        }

        int importo = centsvalue.intValueExact();

        try {
            return Moneta.moneta(new Importo(importo));
        } catch (MonetaException e) {
            throw new MonetaException(e.getMessage());
        }
    }

    /**
     * {@code parseImporto} converte la stringa data in input ({@code input}) in un oggetto del tipo {@code Importo}
     * 
     * @param input è la stringa rappresentativa dell'importo da convertire in {@code Importo}
     * @return l'oggetto {@code Importo} che rappresenta il valore dato in input
     * @throws InvalidImportoException se è stato inserito un valore in input errato
     */
    public static Importo parseImporto(String input) throws InvalidImportoException {
        input = input.replaceAll("\\s+", " ");

        if (input.matches(".*[a-zA-Z].*")) throw new InvalidImportoException("invalid-result");

        BigDecimal centsvalue = (new BigDecimal(input)).multiply(BigDecimal.valueOf(100));

        if (centsvalue.scale() > 2) {
            throw new InvalidImportoException("invalid-result");
        }

        int importo = centsvalue.intValueExact();
        return new Importo(importo);
    }

    /**
     * {@code parseProdotto} converte una stringa data in input ({@code input}) in un oggetto di tipo {@code Prodotto}
     * 
     * @param input è la stringa rappresentativa del prodotto
     * @return l'oggetto di tipo {@code Prodotto} rappresentativo della stringa {@code input}
     * @throws TagliaException se la taglia indicata nella stringa {@code input} non è una tra le seguenti: S, M, L o XL
     * @throws InvalidImportoException se l'importo indicato nella stringa {@code input} non è nel formato corretto
     */
    public static Prodotto parseProdotto(String input) throws TagliaException, InvalidImportoException {
        input = input.replaceAll("\\s+", "");
        String[] specifiche = input.split("\\|");

        return new Prodotto(specifiche[0], Taglia.taglia(specifiche[2]), parseImporto(specifiche[1]));
    }

    /**
     * {@code parseAggregato} converte una stringa data in input ({@code input}) in una tupla nel formato:
     *  {@code (Moneta, quantità)} e lo aggiunge all'aggregato cassa dato in input, ritornando cassa
     * 
     * @param cassa è l'aggregato a cui aggiungere la moneta con la quantità data dalla stringa {@code input}
     * @param input è la strnga rappresentativa della moneta con la sua quantità
     * @return l'aggregato cassa aggiornato con la nuova moneta e la sua quantità
     * @throws InvalidImportoException se l'inserimento della moneta nella cassa non va a buon fine
     * @throws InvalidResultException se il risultato dell'operazione non è valido
     * @throws MonetaException se la moneta inserita non è valida
     */
    public static Aggregato parseAggregato(Aggregato cassa, String input) throws TotalvalueException, InsufficentcoinsException, InvalidImportoException, InvalidResultException, MonetaException {
        input = input.replaceAll("\\s+", " ");
        String[] specifiche = input.split(", ");
        String[] first = specifiche[0].split(" ");

        if (first[0].equals("+")) cassa.Insert(Moneta.moneta(parseImporto(first[first.length - 1])), Integer.parseInt(first[1]));
        else if (first[0].equals("-")) cassa.Remove(Moneta.moneta(parseImporto(first[first.length - 1])), Integer.parseInt(first[1]));
        else cassa.Insert(Moneta.moneta(parseImporto(first[first.length - 1])), Integer.parseInt(first[0]));

        for (int i = 1; i < specifiche.length; i++) {
            String[] single = specifiche[i].split(" ");
            if (first[0].equals("+") && !first[0].equals("-")) cassa.Insert(Moneta.moneta(parseImporto(single[2])), Integer.parseInt(single[0]));
            else cassa.Remove(Moneta.moneta(parseImporto(single[2])), Integer.parseInt(single[0]));
        }

        return cassa;
    }

    /**
     * {@code parseBinario} prende in input una stringa rappresentativa della taglia e capacità e li converte in un oggetto {@code Prodotto}
     * 
     * @param capacity è la stringa rappresentativa della capatità totale del binario
     * @param size è la strnga rappresentativa della taglia del binario
     * @return il {@code Binario} che rappresentativo di taglia e capacità date in input
     * @throws TagliaException se la taglia data in input non è valida
     */
    public static Binario parseBinario(String taglia, String capacity) throws TagliaException{
        return new Binario(Taglia.taglia(taglia), Integer.parseInt(capacity));
    }

    /**
     * {@code parseBinari} è un metodo che permette di creare una lista di binari da specifiche date in input
     * 
     * @param specifiche è la stringa rappresentativa delle specifiche dei binari da creare
     * @return un {@code ArrayList} dei binari
     * @throws TagliaException se un binario ha una stringa che non risulta conforme
     */
    public static ArrayList<Binario> parseBinari(String specifiche) throws TagliaException {
        ArrayList<Binario> rails = new ArrayList<>();
        String[] specificheBinari = specifiche.split(", ");

        for (String specifica : specificheBinari) {
            String[] specs = specifica.split("\\|");
            try {
                rails.add(Parser.parseBinario(specs[1], specs[0]));
            } catch (TagliaException e) {
                System.out.println("- size");
            }
        }
        
        return rails;
    }

    /**
     * {@code parseStrategia} è un metodo che converte una stringa in una delle tre {@code strategia} disponibili del distributore
     * 
     * @param strategia è la stringa rappresentativa della strategia scelta per il distributore
     * @return un oggetto {@code StrategiaResto} che identifica la {@code strategia} scelta
     */
    public static StrategiaResto parseStrategia(String strategia) {
        switch (strategia) {
            case "H":
                return new RestoAlto();
            case "L":
                return new RestoBasso();
            default:
                return new RestoMedio();
        }
    }

    /**
     * {@code parseDistributore} permette di gestire gli input di un distributore automatico
     * 
     * @param input è la stringa data in input al distributore
     * @return è la stringa contenente il risultato dell'operazione
     * @throws EmptyRailException se il binario è vuoto
     * @throws InsufficentChangeException se non ci sono abbastanza monete nel distributore
     * @throws InsufficentValueException se non c'è abbastanza valore nella cassa
     * @throws SlotException se è stato selezionato uno slot non disponibile
     */
    public static String parseDistributore(DistributoreAutomatico macchinetta, String input) throws EmptyRailException, InsufficentChangeException, InsufficentValueException, SlotException{
        if (input.contains("+")) {
            String[] elements = input.split(" ! ");
            String[] specifiche = elements[1].split(";");
            StringBuilder prodotto = new StringBuilder();
            String[] q = elements[0].split(" ");
            int quantity = Integer.parseInt(q[1]);

            prodotto.append(specifiche[0]).append("|").append(specifiche[1]).append("|").append(specifiche[2]);

            try {
                return "+ " + macchinetta.caricaBinario(parseProdotto(prodotto.toString()), quantity);
            } catch (TagliaException | InvalidImportoException e) {
                System.out.println(e.getMessage());
            }
        } else if (input.equals("?")) return macchinetta.toString();
        else {
            String[] elements = input.split(" ! ");
            String[] parti = elements[1].split("; ");
            String[] railnumber = elements[0].split(" ");
            int rail = Integer.parseInt(railnumber[1]);
            Importo importo = new Importo(0);

            for (String parte : parti) {
                String[] parts = parte.split(" * ");
                int converted = (new BigDecimal(parts[2])).multiply(BigDecimal.valueOf(100)).intValueExact();
                int multiplier = Integer.parseInt(parts[0]);

                try {
                    importo = importo.Add(Parser.parseImporto(Float.toString((float) (multiplier * converted) / 100)));
                    macchinetta.getAggregato().Insert(Moneta.moneta(Parser.parseImporto(parts[2])), Integer.parseInt(parts[0]));
                } catch (InvalidResultException | InvalidImportoException | MonetaException e) {
                    System.out.println(e.getMessage());
                }
            }

            return "- " + macchinetta.scaricaBinario(rail, importo);
        }

        return "";
    }
}
