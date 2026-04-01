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

import java.util.Iterator;
import java.util.Scanner;

/** A client class for FilteredStringIterator. */
public class FilteredStringIteratorClient {

  /** . */
  private FilteredStringIteratorClient() {}

  /**
   * Reads strings from standard input and emits in the standard output those ending with a given
   * suffix.
   *
   * <p>It expects a single command line argument, the suffix to filter for.
   *
   * @param args the command line arguments, the first of which is the suffix to filter for.
   */
  public static void main(String[] args) {
    StringEndsWith filter = new StringEndsWith(args[0]);
    try (Scanner sc = new Scanner(System.in)) {
      Iterator<String> it = new FilteredStringIterator(sc, filter);
      while (it.hasNext()) System.out.println(it.next());
    }
  }
}
