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

package it.unimi.di.prog2.h17;

import java.util.Scanner;

/** A class to test {@link ArrayIntSet} and {@link ListIntSet}. */
public class IntSetClient {

  /** . */
  private IntSetClient() {}

  /**
   * Tests some methods of {@link ArrayIntSet} and {@link ListIntSet}.
   *
   * <p>This method reads integers from standard input and inserts them into a set, then it emits
   * the set size and elements in the standard output. The first parameter determines the
   * implementation to use: {@code a} for {@link ArrayIntSet}, anything else for {@link ListIntSet}.
   *
   * @param args if the set should be ordered.
   */
  public static void main(String[] args) {
    final char impl = args.length > 0 && args[0].length() == 1 ? args[0].charAt(0) : 'l';
    AbstractIntSet set = impl == 'a' ? new ArrayIntSet() : new ListIntSet();
    try (Scanner sc = new Scanner(System.in)) {
      while (sc.hasNextInt()) set.insert(sc.nextInt());
    }
    System.out.println("Size: " + set.size());
    if (set.size() > 0) {
      System.out.println("Elements:");
      for (int x : set) System.out.println(x);
    }
  }
}
