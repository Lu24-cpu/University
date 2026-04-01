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

package it.unimi.di.prog2.s06;

import java.util.Scanner;

/** Client code for {@link ArrayUtils}. */
public class ArrayUtilsClient {

  /** . */
  private ArrayUtilsClient() {}

  /**
   * Receives a size as command line argument, then initializes an array of that size filled with
   * {@link Integer#MAX_VALUE}. Then reads integers from standard input and inserts them in the
   * array if not already present, keeping the array sorted and of the same size (discarding the
   * last element at every insertion). Finally, prints the array.
   *
   * @param args the size of the array.
   */
  public static void main(String[] args) {

    int size = Integer.parseInt(args[0]);
    int[] array = new int[size];
    ArrayUtils.fill(array, Integer.MAX_VALUE);

    try (Scanner scanner = new Scanner(System.in)) {
      while (scanner.hasNextInt()) {
        int value = scanner.nextInt();
        int index = ArrayUtils.binarySearch(array, value);
        if (index < 0) { // value not found
          int insertionPoint = -index - 1;
          if (insertionPoint < size) ArrayUtils.insertAt(array, insertionPoint, value);
        }
      }
    }

    ArrayUtils.print(array);
  }
}
