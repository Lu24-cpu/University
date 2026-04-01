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

import java.util.ArrayList;
import java.util.Collections;
import java.util.Iterator;
import java.util.List;
import java.util.NoSuchElementException;

/**
 * {@code ListIntSet}s are mutable, unbounded sets of integers.
 *
 * <p>A typical ListIntSet is \( S = \{x_1, \ldots, x_n \} \).
 */
public class ListIntSet extends AbstractIntSet {

  // Fields

  /** The {@link List} containing this set elements. */
  private final List<Integer> els;

  /*
   * RI:
   *
   * - els != null
   * - els do not contain null elements
   * - for all 0 <= i != j < els.size(), els.get(i) != els.get(j)
   *
   * AF:
   *
   * - represents the set S = { els.get(0), els.get(1), ..., els.get(els.size()-1) }
   *
   */

  // Constructors

  /**
   * Initializes this set to be empty.
   *
   * <p>Builds the set \( S = \varnothing \).
   */
  public ListIntSet() {
    super();
    els = new ArrayList<>();
  }

  // Methods

  /**
   * Looks for a given element in this set.
   *
   * @param x the element to look for.
   * @return the index where {@code x} appears in {@code els} if the element belongs to this set, or
   *     -1
   */
  private int getIndex(int x) {
    return els.indexOf(x);
  }

  /**
   * Adds the given element to this set.
   *
   * <p>This method modifies the object, that is: \( S' = S \cup \{ x \} \).
   *
   * @param x the element to be added.
   */
  public void insert(int x) {
    if (getIndex(x) < 0) {
      els.add(x);
      size++;
    }
  }

  /**
   * Removes the given element from this set.
   *
   * <p>This method modifies the object, that is: \( S' = S \setminus \{ x \} \).
   *
   * @param x the element to be removed.
   */
  public void remove(int x) {
    int i = getIndex(x);
    if (i < 0) return;
    int last = els.size() - 1;
    els.set(i, els.get(last));
    els.remove(last);
    size--;
  }

  /**
   * Returns an element from this set.
   *
   * @return an arbitrary element from this set.
   * @throws NoSuchElementException if this set is empty.
   */
  public int choose() throws NoSuchElementException {
    if (els.isEmpty()) throw new NoSuchElementException("Can't choose from an empty set");
    return els.get(els.size() - 1);
  }

  @Override
  public Iterator<Integer> iterator() {
    return Collections.unmodifiableCollection(els).iterator();
  }
}
