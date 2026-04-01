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

package it.unimi.di.prog2.s03;

import java.util.Scanner;

/** Classe per la somma di importi in centesimi. */
public class SommaCentesimi {

  /** . */
  private SommaCentesimi() {}

  /**
   * Legge dal flusso di ingresso un elenco di importi in euro e centesimi (uno per riga, con la
   * parte decimale separata dalla parte intera da un punto) e ne emette nel flusso d'uscita la
   * somma.
   *
   * @param args gli argomenti (ignorati) della riga di comando
   */
  public static void main(String[] args) {
    int cents = 0;
    try (Scanner sc = new Scanner(System.in)) {
      while (sc.hasNextLine()) {
        String[] parts = sc.nextLine().split("\\.");
        cents += 100 * Integer.parseInt(parts[0]) + Integer.parseInt(parts[1]);
      }
      System.out.println(cents / 100 + "." + (cents % 100 < 10 ? "0" : "") + cents % 100);
    }
  }
}
