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

/** Classe di metodi di utilità per array. */
public class ArrayUtils {

  /** . */
  private ArrayUtils() {}

  /**
   * Finds the index (or insertion point) of an integer in an array of integers in increasing order.
   *
   * <ul>
   *   <li><b>Requires:</b> {@code haystack} is not {@code null} and is sorted in increasing order.
   *   <li><b>Modifies:</b> nothing.
   *   <li><b>Effects:</b> if {@code haystack} contains {@code needle}, returns the index of the
   *       first occurrence of {@code needle} in {@code haystack}; otherwise, returns {@code
   *       -(insertion_point) - 1} where {@code insertion_point} is the index of the first integer
   *       greater than {@code needle}; note that this implies that the return value is non-negative
   *       iff {@code haystack} contains {@code needle}.
   * </ul>
   */
  @SuppressWarnings("doclint:missing") // this is because we are still Liskov style
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
}
