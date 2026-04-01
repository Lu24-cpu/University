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

package it.unimi.di.prog2.h03;

import java.util.Scanner;

/** Classe per calcolare la somma degli elementi di un array. */
public class SommaArray {

  /** Costruttore privato per impedire l'istanziazione della classe. */
  private SommaArray() {}

  /**
   * Calcola la somma degli elementi di un array di interi.
   *
   * @param values L'array di interi da sommare.
   * @return La somma degli elementi dell'array.
   */
  static int somma(int[] values) {
    int total = 0;
    for (int v : values) {
      total += v;
    }
    return total;
  }

  /**
   * Legge dal flusso in ingresso al più 10 interi e ne emette la somma nel flusso d'uscita.
   *
   * @param args parametri da linea di comando (non usati).
   */
  public static void main(String[] args) {
    int[] values = new int[10];
    int read = 0;
    try (Scanner scanner = new Scanner(System.in)) {
      while (read < values.length && scanner.hasNextInt()) {
        values[read++] = scanner.nextInt();
      }
    }
    System.out.println(somma(values));
  }
}
