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

package it.unimi.di.prog2.h10;

import java.util.Arrays;
import java.util.Objects;

/**
 * {@code DensePoly}s are immutable polynomials with integer coefficients.
 *
 * <p>A typical {@code Poly} is \( p = c_0 + c_1 x + c_2 x^2 + \cdots + c_n x^n \).
 */
public class DensePoly { // we don't extend Cloneable, see EJ 3.13

  /** The array of coefficients, the {@code coeff[i]} is the coefficient of \( x^i \). */
  private final int[] coefficient;

  /*
   * RI:
   *
   *  - coefficient != null
   *  - if coefficient.length > 0, coefficient[coefficient.length - 1] != 0
   *
   * AF:
   *
   *  - the polynomial 0 is represented by an array of length 0,
   *  - a polynomial of degree n is represented by an array of length n+1,
   *    where coefficient[i] is the coefficient of x^i.
   */

  /** Initializes this to be the zero polynomial, that is \( p = 0 \). */
  public DensePoly() {
    coefficient = new int[0];
  }

  /**
   * Initializes this to be the polynomial \(p = cx^n\).
   *
   * @param c the coefficient.
   * @param n the degree.
   * @throws IllegalArgumentException if {@code n} &lt; 0.
   */
  public DensePoly(int c, int n) throws IllegalArgumentException {
    if (n < 0) throw new IllegalArgumentException("Can't create a monomial with negative exponent");
    if (c == 0) {
      coefficient = new int[0];
    } else {
      coefficient = new int[n + 1];
      coefficient[n] = c;
    }
  }

  /**
   * Initializes a polynomial given an array of coefficients.
   *
   * <p>The coefficient array can have trailing zeros, they will not be considered in the
   * representation of the polynomial.
   *
   * @param coefficient the array of coefficients.
   */
  public DensePoly(int[] coefficient) {
    Objects.requireNonNull(coefficient, "The coefficient array must not be null.");
    if (coefficient.length == 0) {
      this.coefficient = new int[0];
    } else {
      int degree = coefficient.length - 1;
      while (degree >= 0 && coefficient[degree] == 0) degree--;
      this.coefficient = Arrays.copyOf(coefficient, degree + 1);
    }
  }

  /**
   * Returns the degree of this polynomial.
   *
   * <p>The degree is defined as the largest exponent with a non-zero coefficient.
   *
   * @return the largest exponent with a non-zero coefficient.
   * @throws IllegalStateException if this is the zero polynomial.
   */
  public int degree() {
    if (coefficient.length == 0)
      throw new IllegalStateException("The zero polynomial has no degree.");
    return coefficient.length - 1;
  }

  /**
   * Returns the coefficient of the term of given exponent.
   *
   * @param d the exponent of the term to consider.
   * @return the coefficient of the considered term.
   */
  public int coeff(int d) {
    if (d < 0 || d >= coefficient.length) return 0;
    else return coefficient[d];
  }

  /**
   * Checks whether this polynomial is the zero polynomial.
   *
   * @return {@code true} if this polynomial is the zero polynomial, {@code false} otherwise.
   */
  public boolean isZero() {
    return coefficient.length == 0;
  }

  /**
   * Performs polynomial addition.
   *
   * <p>If \( p \) is this polynomial, returns \( p + q \).
   *
   * @param q the polynomial to add to this one.
   * @return the sum among this and the given polynomial.
   * @throws NullPointerException if {@code q} is {@code null}.
   */
  public DensePoly add(DensePoly q) throws NullPointerException {
    Objects.requireNonNull(q, "The polynomial must not be null.");
    if (isZero()) return q;
    if (q.isZero()) return this;
    final DensePoly larger, smaller;
    if (degree() > q.degree()) {
      larger = this;
      smaller = q;
    } else {
      larger = q;
      smaller = this;
    }
    int[] result = new int[larger.degree() + 1];
    int i, j;
    for (i = 0; i <= smaller.degree(); i++)
      result[i] = smaller.coefficient[i] + larger.coefficient[i];
    for (j = i; j <= larger.degree(); j++) result[j] = larger.coefficient[j];
    return new DensePoly(result);
  }

  /**
   * Performs polynomial multiplication.
   *
   * <p>If \( p \) is this polynomial, returns \( p q \).
   *
   * @param q the polynomial to multiply by this one.
   * @return the product among this and the given polynomial.
   * @throws NullPointerException if {@code q} is {@code null}.
   */
  public DensePoly mul(DensePoly q) throws NullPointerException {
    Objects.requireNonNull(q, "The polynomial must not be null.");
    if (isZero() || q.isZero()) return new DensePoly();
    int[] result = new int[degree() + q.degree() + 1];
    for (int i = 0; i <= degree(); i++)
      for (int j = 0; j <= q.degree(); j++)
        result[i + j] = result[i + j] + coefficient[i] * q.coefficient[j];
    return new DensePoly(result);
  }

  /**
   * Performs polynomial subtraction.
   *
   * <p>If \( p \) is this polynomial, returns \( p - q \).
   *
   * @param q the polynomial to subtract from this one.
   * @return the subtraction among this and the given polynomial.
   * @throws NullPointerException if {@code q} is {@code null}.
   */
  public DensePoly sub(DensePoly q) throws NullPointerException {
    Objects.requireNonNull(q, "The polynomial must not be null.");
    if (isZero()) return q.minus();
    if (q.isZero()) return this;
    return add(q.minus());
  }

  /**
   * Returns the negate polynomial.
   *
   * <p>If \( p \) is this polynomial, returns \( -p \).
   *
   * @return this polynomial multiplied by \( -1 \).
   */
  public DensePoly minus() {
    if (isZero()) return this;
    int[] result = new int[coefficient.length];
    for (int i = 0; i <= degree(); i++) result[i] = -coefficient[i];
    return new DensePoly(result);
  }
}
