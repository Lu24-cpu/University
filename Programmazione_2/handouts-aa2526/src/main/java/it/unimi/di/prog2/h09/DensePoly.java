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

package it.unimi.di.prog2.h09;

import java.util.Objects;

/**
 * {@code DensePoly}s are immutable polynomials with integer coefficients.
 *
 * <p>A typical {@code Poly} is \( p = c_0 + c_1 x + c_2 x^2 + \cdots + c_n x^n \).
 */
public class DensePoly { // we don't extend Cloneable, see EJ 3.13

  /** The array of coefficients, the {@code coeff[i]} is the coefficient of \( x^i \). */
  private final int[] coefficient;

  /** Initializes this to be the zero polynomial, that is \( p = 0 \). */
  public DensePoly() {
    coefficient = new int[1];
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
    if (c == 0) n = 0;
    coefficient = new int[n + 1];
    coefficient[n] = c;
  }

  /**
   * Initializes a polynomial of given degree (with all coefficients equal to 0).
   *
   * @param n the degree, must be non-negative.
   */
  private DensePoly(int n) {
    coefficient = new int[n + 1];
  }

  /**
   * Returns the degree of this polynomial.
   *
   * <p>The degree is defined as the largest exponent with a non-zero coefficient, except for the
   * zero polynomial, whose degree is defined to be 0.
   *
   * @return the largest exponent with a non-zero coefficient; returns 0 if this is the zero {@code
   *     Poly}.
   */
  public int degree() {
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
    return coefficient.length == 1 && coefficient[0] == 0;
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
    int resultDegree = larger.degree();
    if (degree() == q.degree()) {
      for (int k = degree(); k > 0; k--)
        if (coefficient[k] + q.coefficient[k] != 0) break;
        else resultDegree--;
    }
    DensePoly result = new DensePoly(resultDegree); // get a new Poly
    int i;
    for (i = 0; i <= smaller.degree() && i <= resultDegree; i++)
      result.coefficient[i] = smaller.coefficient[i] + larger.coefficient[i];
    for (int j = i; j <= resultDegree; j++) result.coefficient[j] = larger.coefficient[j];
    return result;
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
    DensePoly r = new DensePoly(degree() + q.degree());
    for (int i = 0; i <= degree(); i++)
      for (int j = 0; j <= q.degree(); j++)
        r.coefficient[i + j] = r.coefficient[i + j] + coefficient[i] * q.coefficient[j];
    return r;
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
    DensePoly r = new DensePoly(degree());
    for (int i = 0; i <= degree(); i++) r.coefficient[i] = -coefficient[i];
    return r;
  }
}
