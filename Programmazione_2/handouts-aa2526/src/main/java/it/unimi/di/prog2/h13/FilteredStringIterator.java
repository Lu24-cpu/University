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
import java.util.NoSuchElementException;
import java.util.function.Predicate;

/** An iterator filtering an underlying iterator using a {@link Predicate}. */
public class FilteredStringIterator implements Iterator<String> {

  /** The original iterator. */
  private final Iterator<String> original;

  /** The predicate used for filtering. */
  private final Predicate<String> predicate;

  /** The next element to be returned, or null if not yet computed. */
  private String next;

  /*
   * RI:
   *
   *  - original, predicate are not null
   *
   * AF:
   *
   *  - the next element to be returned is stored in next, if not null
   *  - the iterator is exhausted if original is exhausted and next is null
   *
   */

  /**
   * Creates a new filtered iterator.
   *
   * @param original the original iterator.
   * @param predicate the predicate used for filtering.
   */
  public FilteredStringIterator(Iterator<String> original, Predicate<String> predicate) {
    this.original = original;
    this.predicate = predicate;
    this.next = null;
  }

  @Override
  public boolean hasNext() {
    if (next != null) return true;
    while (original.hasNext()) {
      next = original.next();
      if (predicate.test(next)) return true;
    }
    return false;
  }

  @Override
  public String next() {
    if (!hasNext()) throw new NoSuchElementException();
    String result = next;
    next = null;
    return result;
  }
}
