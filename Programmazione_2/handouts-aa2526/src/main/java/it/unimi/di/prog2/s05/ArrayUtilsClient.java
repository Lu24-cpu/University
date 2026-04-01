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

package it.unimi.di.prog2.s05;

import java.util.Scanner;

/** Client for the {@link ArrayUtils} class. */
public class ArrayUtilsClient {

  /** . */
  private ArrayUtilsClient() {}

  /** The fixed size of the input and array. */
  private static final int SIZE = 10;

  /**
   * Reads 10 integers in increasing order from the standard input and prints the index of the
   * integer passed as argument on the command line if found.
   *
   * @param args the value to search for.
   */
  public static void main(String[] args) {

    int[] array = new int[SIZE];
    try (Scanner scanner = new Scanner(System.in)) {
      for (int i = 0; i < SIZE; i++) array[i] = scanner.nextInt();
    }
    int index = ArrayUtils.binarySearch(array, Integer.parseInt(args[0]));
    if (index >= 0) System.out.println(index);
  }
}
