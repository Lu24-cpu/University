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

package it.unimi.di.prog2.s06;

import java.util.Scanner;

/** Client for the {@link MathFunctions} class. */
public class MathFunctionsClient {

  /** . */
  private MathFunctionsClient() {}

  /**
   * Reads a sequence of floating-point numbers from standard input, computes their square root
   * (when possible), and prints {@code true} or {@code false} depending on whether the result is
   * correct within {@code 10^-3} or not.
   *
   * @param args the command line arguments (ignored).
   */
  public static void main(String[] args) {
    try (Scanner s = new Scanner(System.in)) {
      while (s.hasNextDouble()) {
        double x = s.nextDouble();
        double y = MathFunctions.sqrt(x);
        System.out.println(Math.abs(y * y - x) < 0.001);
      }
    }
  }
}
