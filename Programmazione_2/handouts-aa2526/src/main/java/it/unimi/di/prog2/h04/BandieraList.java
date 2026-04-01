/*

Copyright 2025 Massimo Santini

This file is part of "Programmazione 2 @ UniMI" teaching material.

This is free software: you can redistribute it and/or modify
it under the terms of the GNU General Public License as published by
the Free Software Foundation, either version 3 of the License, or
(at your option) any later version.

This material is distributed in the hope that it will be useful,
but WITHOUT ANY WARRANTY; without even the implied warranty of
MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
GNU General Public License for more details.

You should have received a copy of the GNU General Public License
along with this file.  If not, see <https://www.gnu.org/licenses/>.

*/

package it.unimi.di.prog2.h04;

import java.util.ArrayList;
import java.util.List;
import java.util.Scanner;

/**
 * Classe per stampare una sequenza di parole in bandiera destra.
 *
 * <p>Questa classe usa {@link java.util.List} come struttura dati dinamica.
 */
public class BandieraList {

  /** Costruttore privato per impedire l'istanziazione della classe. */
  private BandieraList() {}

  /**
   * Legge una sequenza di parole dal flusso standard di ingresso.
   *
   * @return la lista di parole lette.
   */
  public static List<String> read() {
    List<String> words = new ArrayList<String>();
    try (Scanner sc = new Scanner(System.in)) {
      while (sc.hasNext()) words.add(sc.next());
    }
    return words;
  }

  /**
   * Calcola la lunghezza massima delle parole in una lista.
   *
   * @param words la lista di parole.
   * @return la lunghezza massima delle parole nella lista.
   */
  public static int maxLength(List<String> words) {
    int maxLength = 0;
    for (String word : words) maxLength = Math.max(maxLength, word.length());
    return maxLength;
  }

  /**
   * Legge una sequenza di parole dal flusso standard di ingresso e le stampa in bandiera destra.
   *
   * @param args parametri da linea di comando (non usati).
   */
  public static void main(String[] args) {
    List<String> words = read();
    int maxLength = maxLength(words);
    for (String word : words) {
      System.out.println(String.format("%" + maxLength + "s", word).replace(' ', '#'));
    }
  }
}
