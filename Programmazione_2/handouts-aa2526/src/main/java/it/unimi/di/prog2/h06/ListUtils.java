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

import java.util.List;

/** Classe di utilità per le liste. */
public class ListUtils {

  /** Costruttore privato che impedisce l'istanziazione. */
  private ListUtils() {}

  /**
   * Cerca una stringa in una lista ordinata lessicograficamente.
   *
   * <p>Se la lista è ordinata e contiene il valore cercato, il metodo restituisce un intero {@code
   * i} tale che {@code haystack.get(i)} è uguale a {@code needle}, se non contiene il valore
   * cercato restituisce {@code -1}. Non muta il contenuto della lista.
   *
   * @param haystack la lista di stringhe, deve essere ordinata lessicograficamente, non può essere
   *     {@code null}.
   * @param needle la stringa da cercare, non può essere {@code null}.
   * @return l'indice corrispondente alla stringa cercata, oppure {@code -1}.
   */
  public static int binarySearch(List<String> haystack, String needle) {
    int lo = 0, high = haystack.size() - 1;
    while (lo <= high) {
      int mid = lo + high >>> 1;
      String midString = haystack.get(mid);
      if (needle.compareTo(midString) < 0) high = mid - 1;
      else if (needle.compareTo(midString) > 0) lo = mid + 1;
      else return mid;
    }
    return -1;
  }
}
