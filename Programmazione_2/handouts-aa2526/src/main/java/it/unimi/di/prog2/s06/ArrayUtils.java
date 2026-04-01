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

import java.util.List;
import java.util.Objects;

/** Classe di metodi di utilità per array. */
public class ArrayUtils {

  /** . */
  private ArrayUtils() {}

  /**
   * Finds the index (or insertion point) of an integer in an array of integers in increasing order.
   *
   * <p>If the array contains the given integer, returns its index. Otherwise, returns {@code
   * -(insertion_point) - 1} where {@code insertion_point} is the index of the first integer greater
   * than {@code needle}; note that this implies that the return value is non-negative iff the array
   * contains the integer.
   *
   * @see java.util.Collections#binarySearch(List, Object)
   * @param haystack the not {@code null} array of integers in increasing order.
   * @param needle the integer to look for.
   * @return the index of the given integer, or {@code -insertion_point - 1} if none is present.
   * @throws NullPointerException if {@code haystack} is {@code null}.
   */
  static int binarySearch(final int[] haystack, final int needle) {
    Objects.requireNonNull(haystack);
    int lo = 0;
    int hi = haystack.length - 1;
    while (lo <= hi) {
      int mid = lo + (hi - lo) / 2;
      if (needle < haystack[mid]) hi = mid - 1;
      else if (needle > haystack[mid]) lo = mid + 1;
      else return mid;
    }
    return -lo - 1;
  }

  /**
   * Shift elements of {@code array} to the right starting at {@code insertionPoint} (inclusive) and
   * inserts {@code value} at {@code insertionPoint}. The last element is discarded to keep the
   * array size constant. Assumes {@code 0 <= insertionPoint < array.length}.
   *
   * @param array the array where to insert the value.
   * @param insertionPoint the index where to insert the value.
   * @param value the value to insert.
   * @throws NullPointerException if {@code array} is {@code null}.
   * @throws ArrayIndexOutOfBoundsException if {@code insertionPoint} is out of bounds.
   */
  static void insertAt(int[] array, int insertionPoint, int value) {
    Objects.requireNonNull(array);
    if (insertionPoint < 0 || insertionPoint >= array.length) {
      throw new ArrayIndexOutOfBoundsException(insertionPoint);
    }
    for (int i = array.length - 1; i > insertionPoint; i--) array[i] = array[i - 1];
    array[insertionPoint] = value;
  }

  /**
   * Fills the given array with the given value.
   *
   * @param array the array to fill.
   * @param value the value to fill the array with.
   * @throws NullPointerException if {@code array} is {@code null}.
   */
  static void fill(int[] array, int value) {
    Objects.requireNonNull(array);
    for (int i = 0; i < array.length; i++) array[i] = value;
  }

  /**
   * Prints the content of the given array, one element per line.
   *
   * @param array the array to print.
   * @throws NullPointerException if {@code array} is {@code null}.
   */
  static void print(int[] array) {
    Objects.requireNonNull(array);
    for (int i = 0; i < array.length; i++) System.out.println(array[i]);
  }
}
