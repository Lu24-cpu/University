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

/** Utility class for computing mathematical functions. */
public class MathFunctions {

  /** . */
  private MathFunctions() {}

  /**
   * Extracts (if possible) the square root of the given number.
   *
   * <p>Partial function which, when the argument is non-negative, returns an approximation
   * {@literal \( y \)} of the square root in the sense that {@literal \( |y^2 - x| < 10^{-3} \)}.
   *
   * @param x the number whose square root is to be extracted; must be non-negative.
   * @return an approximation of the square root.
   * @throws IllegalArgumentException if {@code x} is negative.
   */
  public static double sqrt(double x) {
    if (x < 0) throw new IllegalArgumentException("Negative argument: " + x);
    if (x == 0) return 0;
    double guess = x / 2;
    while (Math.abs(guess * guess - x) >= 0.00001) guess = (guess + x / guess) / 2;
    return guess;
  }
}
