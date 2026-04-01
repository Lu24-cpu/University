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

/** Classe per la verifica dell'uguaglianza tra frazioni. */
public class UguaglianzaFrazioni {
  /** . */
  private UguaglianzaFrazioni() {}

  /**
   * Riceve come parametri sulla linea di comando quattro interi corrispondenti rispettivamente a
   * numeratore e denominatore di due frazioni ed emette nel flusso d'uscita "uguali" se le frazioni
   * sono uguali, oppure "diverse".
   *
   * @param args i quattro interi
   */
  public static void main(String[] args) {
    int a = Integer.parseInt(args[0]);
    int b = Integer.parseInt(args[1]);
    int c = Integer.parseInt(args[2]);
    int d = Integer.parseInt(args[3]);

    /* Per una discussione del problema dell'equivalenza si veda https://stackoverflow.com/a/73394324 */

    System.out.println((long) a * d == (long) b * c ? "uguali" : "diverse");
  }
}
