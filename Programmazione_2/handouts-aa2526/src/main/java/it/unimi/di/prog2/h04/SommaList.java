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
 * Classe per calcolare la somma degli elementi di una lista.
 *
 * <p>In questa classe viene esplicitato il boxing e l'unboxing degli interi.
 */
public class SommaList {

  /** Costruttore privato per impedire l'istanziazione della classe. */
  private SommaList() {}

  /**
   * Calcola la somma degli elementi di un array di interi.
   *
   * @param values L'array di interi da sommare.
   * @return La somma degli elementi dell'array.
   */
  static int somma(List<Integer> values) {
    int total = 0;
    for (Integer v : values) {
      total += v.intValue();
    }
    return total;
  }

  /**
   * Legge dal flusso in ingresso al più 10 interi e ne emette la somma nel flusso d'uscita.
   *
   * @param args parametri da linea di comando (non usati).
   */
  public static void main(String[] args) {
    List<Integer> values = new ArrayList<>();
    try (Scanner scanner = new Scanner(System.in)) {
      while (scanner.hasNextInt()) values.add(Integer.valueOf(scanner.nextInt()));
    }
    System.out.println(somma(values));
  }
}
