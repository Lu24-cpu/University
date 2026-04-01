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

package it.unimi.di.prog2.h06;

import java.util.ArrayList;
import java.util.List;
import java.util.Scanner;

/** Classe client per {@link ListUtils}. */
public class ListUtilsClient {

  /** . */
  private ListUtilsClient() {}

  /**
   * Legge dal flusso di ingresso una lista di stringhe in ordine lessicografico e stampa l'indice
   * della stringa passata come argomento sulla linea di comando, oppure -1 se tale stringa non
   * compare nella lista.
   *
   * @param args il primo elemento dell'array è la stringa da cercare (ci deve essere almeno un
   *     argomento).
   */
  public static void main(String[] args) {
    List<String> stringhe = new ArrayList<>();
    try (Scanner sc = new Scanner(System.in)) {
      while (sc.hasNext()) stringhe.add(sc.next());
    }
    System.out.println(ListUtils.binarySearch(stringhe, args[0]));
  }
}
