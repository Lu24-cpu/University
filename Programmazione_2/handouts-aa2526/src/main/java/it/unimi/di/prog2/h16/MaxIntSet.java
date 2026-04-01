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

package it.unimi.di.prog2.h16;

import it.unimi.di.prog2.h13.ArrayIntSet;
import java.util.NoSuchElementException;

/**
 * Example of {@code MaxIntSet} taken from section 7.4 of the textbook by Liskov <em>et al.</em>.
 *
 * <p><b>Note</b>: this class extends {@link it.unimi.di.prog2.h13.ArrayIntSet}.
 */
public class MaxIntSet extends ArrayIntSet {

  /** The biggest element, if set is not empty. */
  private int biggest;

  // RI: size() == 0 or isIn(biggest) and for every x isIn(x) implies biggest >= x.
  // AF: coincides with that of IntSet

  /** Construct an empty {@code MaxIntSet}. */
  public MaxIntSet() {
    super();
  }

  @Override
  public void insert(final int x) {
    if (size() == 0 || x > biggest) biggest = x;
    super.insert(x);
  }

  @Override
  public void remove(final int x) {
    super.remove(x);
    if (size() == 0 || x != biggest)
      return; // observe that if x > biggest it was not actually in this, so we don't need to
    // update biggest
    biggest = Integer.MIN_VALUE;
    for (int z : this) if (z > biggest) biggest = z;
  }

  /**
   * Returns the maximum value in the set, or raises {@link IllegalStateException} otherwise.
   *
   * @return the maximum value in the set.
   * @throws NoSuchElementException if the set is empty.
   */
  public int max() throws NoSuchElementException {
    if (size() == 0) throw new NoSuchElementException("An empty set does not have a maximum");
    return biggest;
  }

  @Override
  public String toString() {
    return super.toString() + (size() > 0 ? ", max = " + biggest : "");
  }
}
