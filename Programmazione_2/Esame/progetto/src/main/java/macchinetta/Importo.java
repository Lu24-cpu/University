package macchinetta;

import java.util.Objects;

import customException.InvalidImportoException;
import customException.InvalidResultException;

/**
 * {@code Importo} è una classe che rappresenta un valore monetario nel formato {@code unità} e {@code centesimi}
 */
public class Importo implements Comparable<Importo> {
    
    /** {@code units} è la rappresentazione delle unità dell'{@code importo} */
    private final int units;
    /** {@code cents} è la rappresentazione dei centesimi dell'{@code importo} */
    private final int cents;

    /*
     * Invariante di rappresentazione (RI):
     *      - I centesimi devo essere compresi tra 0 e 99
     *      - I centesimi e le unità non possono essere null
     * 
     * Funzione di astrazione (AF):
     *      - Un importo può essere sommato o sottratto ad un altro
     *      - Un importo può essere moltiplicato per un valore intero (senza centesimi)
     *      - La divisione ammete solamente due operandi interi (senza centesimi)
     * 
     */

    /**
     * {@code Importo} è il costruttore dell'oggetto
     * 
     * @param unit è la variabile intera che indica le unità dell'{@code importo}
     * @param cent sono la rappresentazione intera dei centesimi dell'{@code importo}
     * @throws InvalidImportoException se i centesimi sono oltre i 99 o minori di 0 o le unità sono negative
     */
    public Importo(int unit, int cent) throws InvalidImportoException{
        if (cent > 99 || cent < 0 || unit < 0) throw new InvalidImportoException("not valid inputs");
        this.units = unit;
        this.cents = cent;
    }

    /**
     * {@code Importo} è un secondo costruttore di un {@code importo} che prende la visualizzasione in centesimi di tale importo
     * 
     * @param value è la visualizzazione in formato intero dell'{@code importo}
     */
    public Importo(int value) {
        this.cents = value%100;
        this.units = value/100;
    }

    /**
     * {@code getUnits} restituisce la variabile delle unità dell'{@code importo} corrente
     * 
     * @return le unità dell'{@code importo}
     */
    public int getUnits(){
        return units;
    }

    /**
     * {@code getCents} restituisce la variabile dei centesimi dell{@code importo} corrente
     * 
     * @return i centesimi dell'{@code importo}
     */
    public int getCents() {
        return cents;
    }

    /**
     * {@code getTotal} restituisce il valore totale dell'{@code importo} nella visualizzazione in centesimi
     * 
     * @return l'intero rappresentativo del totale dell'{@code importo} corrente
     */
    public int getTotalCents() {
        return units*100 + cents;
    }

    /**
     * {@code Add} è un metodo che fa la somma tra due {@code importi}
     * 
     * @param operando è il secondo {@code importo} della somma
     * @return l'{@code importo} risultate della somma
     * @throws InvalidImportoException se le unità o i centesimi non sono conformi
     * @throws InvalidResultException se il risultato è negativo
     */
    public Importo Add(Importo operando) throws InvalidImportoException, InvalidResultException {
        int result = getTotalCents() + operando.getTotalCents();

        if(result < 0) throw new InvalidResultException("il risultato è negativo");
        return new Importo(result/100, result%100);
    }

    /**
     * {@code Sub} è un metodo che sottrae due {@code importi}
     * 
     * @param operando è il secondo {@code importo} coinvolto nella sottrazione
     * @return l'{@code importo} risultante
     * @throws InvalidImportoException se è stato inserito un importo non conforme
     * @throws InvalidResultException se il risultato è negativo
     */
    public Importo Sub(Importo operando) throws InvalidImportoException, InvalidResultException {
        int result = getTotalCents() - operando.getTotalCents();

        if(getTotalCents() - operando.getTotalCents() < 0) throw new InvalidResultException("Il risultato è negativo");
        
        return new Importo(result);
    }

    /**
     * {@code Mul} è un metodo che calcola la moltiplicazione tra un {@code importo} e un intero salvato nelle {@code unità} di quello in input
     * 
     * @param operando è il moltiplicatore dell'{@code importo} corrente
     * @return l'{@code importo} risultante della moltiplicazione
     * @throws InvalidImportoException se l'{@code operando} passato non è un valore multiplo di 100
     */
    public Importo Mul(int operando) throws InvalidImportoException {
        int result = getTotalCents() * operando;

        if(result%100 != 0) throw new InvalidImportoException("È stato inserito un moltiplicatore non valido");

        return new Importo(result);
    }

    /**
     * {@code Div} è un metodo che divide due importi
     * 
     * @param operando è il divisore dell'operazione
     * @return un intero rappresentate la divisione intera tra i due valori
     * @throws InvalidImportoException se i {@code centesimi} di almeno uno dei due operandi sono diversi da 0
     */
    public int Div(Importo operando) throws InvalidImportoException {
        if (operando.getTotalCents() == 0) throw new InvalidImportoException("invalid-result");

        return getTotalCents()/operando.getTotalCents();
    }

    /** !!! so che non serve questa documentazione ma è per ricordarmi come funziona !!!
     * {@code equals} è un metodo che controlla se due {@code Importi} sono uguali
     * 
     * @param comparate è l'oggetto da comparare all'importo corrente
     * @return {@code true} se sono entrambi {@code importi} e sono uguali, false altrimenti
     */
    @Override
    public boolean equals(Object comparate) {
        if (comparate instanceof Importo value) {
            return cents == value.getCents() && units == value.getUnits();
        }

        return false;
    }

    /** !!! so che non serve questa documentazione ma è per ricordarmi come funziona !!!
     * {@code hashCode} resituisce un valore di hash rappresentante l'{@code importo} corrente
     * 
     * @return il valore hash (int) rappresentante l'{@code importo} corrente
     */
    @Override
    public int hashCode() {
        return Objects.hash(units, cents);
    }

    /** !! so che non serve questa documentazione me la segno per ricordarmi una cosa !!
     * {@code compareTo} ha bisogno che la classe sia inizializzata come implements Comparable &lt;nome_classe&gt;
     * 
     * @param value è l'imporo da comparare
     * @return il risultato del confronto
     */
    @Override
    public int compareTo(Importo value) {
        if(units < value.units || (units == value.units && cents < value.cents)) return -1;
        else if (units > value.units || (units == value.units && cents > value.cents)) return 1;
        else return 0;
    }

    @Override
    public String toString(){
        StringBuilder importo = new StringBuilder();

        if (units > 0) importo.append(units).append(units > 1 ? " units" : units == 1 ? " unit" : "");
        
        if(cents > 0) {
            if (units > 0) importo.append(" ");

            importo.append(cents).append(cents > 1 ? " cents" : cents == 1 ? " cent" : "");
        }

        return importo.toString();
    }
}
