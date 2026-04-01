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

import java.util.Objects;
import java.util.function.Predicate;

/** A predicate checking whether strings end with a given suffix. */
class StringEndsWith implements Predicate<String> {

  /** The suffix to check for. */
  private final String suffix;

  /*
   * RI:
   *
   *  - suffix is not null
   *
   * AF:
   *
   *  - represents the predicate P(s) = (s != null && s.endsWith(suffix))
   *
   */

  /**
   * Creates a new predicate checking for the given suffix.
   *
   * @param suffix the suffix to check for.
   * @throws NullPointerException if suffix is null.
   */
  public StringEndsWith(String suffix) {
    this.suffix = Objects.requireNonNull(suffix);
  }

  @Override
  public boolean test(String s) {
    return s != null && s.endsWith(suffix);
  }
}
