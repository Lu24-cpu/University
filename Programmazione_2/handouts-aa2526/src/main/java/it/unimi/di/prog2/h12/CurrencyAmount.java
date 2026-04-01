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

import java.util.Objects;

/**
 * Immutable class that represents a (positive or negative) amount in a currency that has an integer
 * part and cents.
 */
public class CurrencyAmount {

  /** A constant representing the zero amount. */
  public static final CurrencyAmount ZERO = new CurrencyAmount(0);

  /** The amount expressed in cents. */
  private final int cents;

  /*
   * RI:
   *
   * AF:
   *
   *  - cents represents the amount in cents, for example 3.24 is represented as 324.
   *
   */

  /**
   * Private constructor.
   *
   * @param cents the amount in cents.
   */
  private CurrencyAmount(int cents) {
    this.cents = cents;
  }

  /**
   * Produce a positive amount.
   *
   * @param units whole units of the amount.
   * @param cents cents of the amount.
   * @return a positive {@link CurrencyAmount} representing the specified amount.
   * @throws IllegalArgumentException if {@code cents} is not between 0 and 99 or {@code units} is
   *     negative.
   */
  public static CurrencyAmount positive(int units, int cents) {
    if (units < 0) throw new IllegalArgumentException("Units must be non negative");
    if (cents < 0 || cents > 99)
      throw new IllegalArgumentException("Cents must be between 0 and 99");
    return new CurrencyAmount(units * 100 + cents);
  }

  /**
   * Produce a negative amount.
   *
   * @param units whole units of the amount.
   * @param cents cents of the amount.
   * @return a negative {@link CurrencyAmount} representing the specified amount.
   * @throws IllegalArgumentException if {@code cents} is not between 0 and 99 or {@code units} is
   *     negative.
   */
  public static CurrencyAmount negative(int units, int cents) {
    if (units < 0) throw new IllegalArgumentException("Units must be non negative");
    if (cents < 0 || cents > 99)
      throw new IllegalArgumentException("Cents must be between 0 and 99");
    return new CurrencyAmount(-(units * 100 + cents));
  }

  /**
   * Parses an amount from a string.
   *
   * <p>The string must have the format <samp>U.C</samp> or <samp>-U.C</samp>, where <samp>U</samp>
   * and <samp>C</samp> are two integers corresponding to the units and cents and <samp>C</samp> has
   * exactly two digits.
   *
   * @param amount the amount.
   * @return the parsed {@link CurrencyAmount}.
   * @throws NullPointerException if {@code amount} is {@code null}.
   * @throws IllegalArgumentException if the string is not in the correct format.
   */
  public static CurrencyAmount parseCurrencyAmount(String amount) {
    if (Objects.requireNonNull(amount, "The amount string must not be null").isEmpty())
      throw new IllegalArgumentException("The amount string must not be empty");
    char sign = amount.charAt(0);
    if (sign == '-') amount = amount.substring(1);
    String[] parts = amount.split("\\.");
    if (parts.length != 2)
      throw new IllegalArgumentException("The amount string must contain a single dot");
    if (parts[1].length() != 2)
      throw new IllegalArgumentException("The part after the dot must have exactly two digits");
    int units, cents;
    try {
      units = Integer.parseInt(parts[0]);
    } catch (NumberFormatException e) {
      throw new IllegalArgumentException("The part before the dot is not an integer");
    }
    try {
      cents = Integer.parseInt(parts[1]);
    } catch (NumberFormatException e) {
      throw new IllegalArgumentException("The part after the dot is not an integer");
    }
    if (units < 0) throw new IllegalArgumentException("Units must be non negative");
    if (cents < 0 || cents > 99)
      throw new IllegalArgumentException("Cents must be between 0 and 99");
    return new CurrencyAmount((sign == '-' ? -1 : 1) * (units * 100 + cents));
  }

  /**
   * Returns the sign of the amount.
   *
   * @return 1 if the amount is positive, 0 if it is zero, -1 otherwise.
   */
  public int sign() {
    if (cents == 0) return 0;
    if (cents < 0) return -1;
    return 1;
  }

  /**
   * Returns the cents of the amount.
   *
   * @return the cents of the amount, it is always non negative.
   */
  public int cents() {
    return Math.abs(cents % 100);
  }

  /**
   * Returns the integer part of the amount.
   *
   * @return the integer part of the amount, it is always non negative.
   */
  public int units() {
    return Math.abs(cents / 100);
  }

  /**
   * Returns a new amount representing the sum of this amount and another.
   *
   * @param other the other amount.
   * @return the sum of the two amounts.
   * @throws NullPointerException if {@code altro} is {@code null}.
   */
  public CurrencyAmount somma(CurrencyAmount other) {
    return new CurrencyAmount(
        cents + Objects.requireNonNull(other, "Other amount must not be null").cents);
  }

  /**
   * Returns a new amount representing the difference between this amount and another.
   *
   * @param other the other amount.
   * @return the difference between the two amounts.
   * @throws IllegalArgumentException if {@code other} is greater than this amount.
   * @throws NullPointerException if {@code other} is {@code null}.
   */
  public CurrencyAmount sottrai(CurrencyAmount other) {
    return new CurrencyAmount(
        cents - Objects.requireNonNull(other, "Other amount must not be null").cents);
  }

  @Override
  public String toString() {
    return (cents < 0 ? "-" : "")
        + String.format("%d.%02d", Math.abs(cents / 100), Math.abs(cents % 100));
  }

  @Override
  public boolean equals(Object obj) {
    if (this == obj) return true;
    if (!(obj instanceof CurrencyAmount other)) return false;
    return cents == other.cents;
  }

  @Override
  public int hashCode() {
    return Integer.hashCode(cents);
  }
}
