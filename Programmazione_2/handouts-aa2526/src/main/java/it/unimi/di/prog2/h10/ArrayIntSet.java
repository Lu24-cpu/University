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

package it.unimi.di.prog2.h10;

import java.util.Arrays;
import java.util.List;
import java.util.NoSuchElementException;

/**
 * {@code ArrayIntSet}s are mutable, unbounded sets of integers.
 *
 * <p>A typical ArrayIntSet is \( S = \{x_1, \ldots, x_n \} \).
 */
public class ArrayIntSet {

  // Fields

  /** The initial capacity of the internal array. */
  private static final int INITIAL_CAPACITY = 16;

  /** The {@link List} containing this set elements. */
  private int[] els;

  /** The number of elements in this set. */
  private int size;

  /*
   * RI:
   *
   * - els != null
   * - els.length >= size
   * - size >= 0
   * - for all 0 <= i != j < size, els[i] != els[j]
   *
   * AF:
   *
   *  - represents the set S = { els[0], els[1], ..., els[size-1] }
   *
   */

  // Constructors

  /**
   * Initializes this set to be empty.
   *
   * <p>Builds the set \( S = \varnothing \).
   */
  public ArrayIntSet() {
    els = new int[INITIAL_CAPACITY];
    size = 0;
  }

  // Methods

  /**
   * Looks for a given element in this set.
   *
   * @param x the element to look for.
   * @return an index {@code i} such that {@code els[i] == x} if the element belongs to this set, or
   *     -1
   */
  private int indexOf(int x) {
    for (int i = 0; i < size; i++) {
      if (els[i] == x) return i;
    }
    return -1;
  }

  /**
   * Adds the given element to this set.
   *
   * <p>This method modifies the object, that is: \( S' = S \cup \{ x \} \).
   *
   * @param x the element to be added.
   */
  public void insert(int x) {
    if (indexOf(x) != -1) return;
    if (size == els.length) els = Arrays.copyOf(els, els.length * 2);
    els[size++] = x;
  }

  /**
   * Removes the given element from this set.
   *
   * <p>This method modifies the object, that is: \( S' = S \setminus \{ x \} \).
   *
   * @param x the element to be removed.
   */
  public void remove(int x) {
    int i = indexOf(x);
    if (i == -1) return;
    els[i] = els[--size];
  }

  /**
   * Tells if the given element is in this set.
   *
   * <p>Answers the question \( x\in S \).
   *
   * @param x the element to look for.
   * @return whether the given element belongs to this set, or not.
   */
  public boolean isIn(int x) {
    return indexOf(x) != -1;
  }

  /**
   * Returns the cardinality of this set.
   *
   * <p>Responds with \( |S| \).
   *
   * @return the size of this set.
   */
  public int size() {
    return size;
  }

  /**
   * Returns an element from this set.
   *
   * @return an arbitrary element from this set.
   * @throws NoSuchElementException if this set is empty.
   */
  public int choose() throws NoSuchElementException {
    if (size == 0) throw new NoSuchElementException("Can't choose from an empty set");
    return els[0];
  }
}
