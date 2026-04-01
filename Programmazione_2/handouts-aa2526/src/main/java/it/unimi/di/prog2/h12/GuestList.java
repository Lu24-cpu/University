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

package it.unimi.di.prog2.h12;

import java.util.ArrayList;
import java.util.Collections;
import java.util.List;
import java.util.Objects;

/**
 * Guest list for a private party of prescribed capacity.
 *
 * <p>Guests are identified by their names that are represented by strings, which must not be null
 * or empty; no check on duplicates is performed.
 */
public class GuestList {

  /** The list of guests. */
  private final List<String> guests;

  /** The maximum capacity of the guest list. */
  private final int capacity;

  /*
   * RI:
   *
   *  - guests is not null and does not contain nulls.
   *  - guests does not contain empty strings.
   *  - capacity >= guests.size()
   *  - capacity > 0
   *
   * AF:
   *
   *  - the invitees are the strings in guests
   *
   */

  /**
   * Builds an empty guest list with the given maximum capacity.
   *
   * @param capacity the maximum size of the guest list.
   * @throws IllegalArgumentException if {@code capacity} is not positive.
   */
  public GuestList(int capacity) {
    if (capacity <= 0) throw new IllegalArgumentException("Capacity must be positive");
    this.capacity = capacity;
    this.guests = new ArrayList<>();
  }

  /**
   * Builds a guest list with the given maximum capacity, populated with an initial list of guests.
   *
   * @param guests the initial list of guests.
   * @param capacity the maximum size of the guest list.
   * @throws IllegalArgumentException if {@code capacity} is not positive.
   * @throws IllegalArgumentException if {@code guests.size()} is greater than {@code capacity}.
   * @throws NullPointerException if {@code guests} is null or contains {@code null}.
   */
  public GuestList(List<String> guests, int capacity) {
    if (capacity <= 0) throw new IllegalArgumentException("Capacity must be positive");
    if (Objects.requireNonNull(guests, "Guests list must not be null").size() > capacity)
      throw new IllegalArgumentException("Too many guests");
    for (String name : guests)
      if (Objects.requireNonNull(name, "Guest name must not be null").isEmpty())
        throw new IllegalArgumentException("Guest name must not be empty");
    this.capacity = capacity;
    this.guests = new ArrayList<>(guests);
  }

  /**
   * Returns the list of guests.
   *
   * <p>The returned {@link List} is unmodifiable.
   *
   * <p>As we'll see this method is an horrible idea!
   *
   * @return an unmodifiable view of the guests list.
   */
  public List<String> guests() {
    return Collections.unmodifiableList(guests);
  }

  /**
   * Adds a guest to this guest list.
   *
   * @param guest the guest name.
   * @throws IllegalStateException if the list has reached its maximum capacity.
   * @throws NullPointerException if {@code guest} is {@code null}.
   */
  public void invite(String guest) {
    if (guests.size() >= capacity) throw new IllegalStateException("Capacity reached");
    if (Objects.requireNonNull(guest, "Guest must not be null").isEmpty())
      throw new IllegalArgumentException("Guest name must not be empty");
    guests.add(guest);
  }

  @Override
  public String toString() {
    return guests.toString();
  }
}
