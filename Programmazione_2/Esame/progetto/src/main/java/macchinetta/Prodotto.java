package macchinetta;

/**
 * {@code Prodotto} è un record che rappresenta un oggetto con un {@code nome}, 
 * una {@code taglia} e un {@code importo} propri
 * 
 * @param name è il nome del {@code prodotto}
 * @param size è la {@code taglia} del {@code prodotto}
 * @param value è l'{@code importo} del {@code prodotto}
 */
public record Prodotto(String name, Taglia size, Importo value) {

    /*
     * Invariante di Rappresentazione (RI):
     *      - Il nome non può essere null
     *      - L'importo non può essere negativo
     *      - La taglia non può essere diversa da S, M, L, XL
     * 
     * Funzione di Astrazione (AF):
     *      - È una tupla di elementi che rappresentano un prodotto
     * 
     */

    @Override
    public String toString() {
        StringBuilder product = new StringBuilder();

        product.append("<").append(name).append(", ").append(value.toString()).append(", ").append(size).append(">");

        return product.toString();
    }
}
