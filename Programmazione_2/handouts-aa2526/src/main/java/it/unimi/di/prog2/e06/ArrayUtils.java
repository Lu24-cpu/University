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

package it.unimi.di.prog2.e06;

import java.util.Collections;
import java.util.List;

/** Classe di metodi di utilità per array. */
@SuppressWarnings(
    "doclint:missing") // this is because comments of methods to implement have yet to be added
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
   * @see Collections#binarySearch(List, Object)
   * @param haystack the not {@code null} array of integers in increasing order.
   * @param needle the integer to look for.
   * @return the index of the given integer, or {@code -insertion_point - 1} if none is present.
   */
  static int binarySearch(final int[] haystack, final int needle) {
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

  /*
   * Specify and implement a method that shifts right all elements of the given array from {@code
   * insertionPoint} (inclusive) to the end of the array, and then inserts {@code value} at {@code
   * insertionPoint}.
   */
  static void insertAt(int[] array, int insertionPoint, int value) {}

  /* Specify and implement a method that fills the given array with the given value. */
  static void fill(int[] array, int value) {}

  /* Specify and implement a method that prints the given array, one element per line. */
  static void print(int[] array) {}
}
