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
import java.util.List;
import java.util.NoSuchElementException;

/** Generator (in Liskov parlance) of the ints contained in a {@link List}. */
public class IntGenerator implements Iterator<Integer> {

  /** The list elements. */
  private final int[] els;

  /** The number of elements in the set. */
  private final int size;

  /** The position of the next element to return (if {@code idx < els.length}. */
  private int idx;

  /**
   * Builds an iterator (partial constructor, used just in {@link ArrayIntSet}).
   *
   * @param els the list of elements, must not be, or contain, {@code null}.
   * @param size the number of set elements in {@code els}.
   */
  public IntGenerator(int[] els, int size) {
    this.els = els;
    this.size = size;
    this.idx = 0;
  }

  @Override
  public boolean hasNext() {
    return idx < size;
  }

  @Override
  public Integer next() {
    if (!hasNext()) throw new NoSuchElementException();
    return els[idx++];
  }
}
