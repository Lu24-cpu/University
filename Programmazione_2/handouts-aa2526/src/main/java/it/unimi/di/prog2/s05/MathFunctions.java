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

package it.unimi.di.prog2.s05;

/** Utility class for computing mathematical functions. */
public class MathFunctions {

  /** . */
  private MathFunctions() {}

  /**
   * Extracts (if possible) the square root of the given number.
   *
   * <ul>
   *   <li><b>Requires:</b> {@literal \( x \geq 0 \)}.
   *   <li><b>Modifies:</b> nothing.
   *   <li><b>Effects:</b> returns {@literal \( y : | y^2 - x | < 10^{-3} \)}.
   * </ul>
   */
  @SuppressWarnings("doclint:missing") // this is because we are still Lisokv' style
  public static double sqrt(double x) {
    double low = 0, high = x > 1 ? x : 1, mid = -1;
    while (high - low > .000001) { // this is (10^-3)^2
      mid = (high + low) / 2;
      if (mid * mid - x < 0) low = mid;
      else high = mid;
    }
    return mid;
  }
}
