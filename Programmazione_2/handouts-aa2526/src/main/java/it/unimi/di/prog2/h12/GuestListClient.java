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
import java.util.List;

/** Client code for {@link GuestList}. */
public class GuestListClient {

  /** . */
  private GuestListClient() {}

  /**
   * Main method.
   *
   * <p>Performs a set of operations, to show some potential issues with the {@link GuestList}
   * implementation.
   *
   * @param args command line arguments (ignored).
   */
  public static void main(String[] args) {

    List<String> invitees = new ArrayList<>();
    invitees.add("Brown");
    invitees.add("Black");

    GuestList party = new GuestList(invitees, 3);

    invitees.add(
        "White"); // In absence of defensive copy in the constructor, this would modify party!

    List<String> guests = party.guests();

    try {
      guests.add("Blue"); // In absence of unmodifiable view in guests(), this would modify party!
    } catch (UnsupportedOperationException e) {
      System.out.println(
          "Caught expected UnsupportedOperationException when trying to modify guests list");
    }

    System.out.println("Invitees list: " + invitees);
    System.out.println("Party guests: " + party);
    System.out.println("Guests list: " + guests);
  }
}
