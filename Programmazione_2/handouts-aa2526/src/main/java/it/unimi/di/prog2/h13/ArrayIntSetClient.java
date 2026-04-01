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

package it.unimi.di.prog2.h13;

import java.util.Scanner;

/** Classe per il test di {@link ArrayIntSet} {@link IntGenerator}. */
public class ArrayIntSetClient {
  /** . */
  private ArrayIntSetClient() {}

  /**
   * Legge una sequenza di interi dal flusso di ingresso ed emette la loro somma nel flusso
   * d'uscita.
   *
   * @param args non usato.
   */
  public static void main(String[] args) {
    ArrayIntSet set = new ArrayIntSet();

    try (Scanner s = new Scanner(System.in)) {
      while (s.hasNextInt()) set.insert(s.nextInt());
    }

    int tot = 0;
    for (int x : set) tot += x;
    System.out.println(tot);
  }
}
