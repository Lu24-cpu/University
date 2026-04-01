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

package it.unimi.di.prog2.h11;

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

  // Prefer static factory methods to constructors (EJ 2.1, 2.3)

  /** The (only instance) of the zero polynomial. */
  public static final DensePoly ZERO = new DensePoly(new int[0]);

  /**
   * Initializes the polynomial given an array of coefficients.
   *
   * <p>This constructor is partial, it's up to the factory methods to ensure that the
   * representation invariant is preserved.
   *
   * @param coefficient the array of coefficients (must satisfy the representation invariant).
   */
  private DensePoly(int[] coefficient) {
    assert coefficient != null : "The coefficient array must not be null.";
    assert coefficient.length == 0 || coefficient[coefficient.length - 1] != 0
        : "The last value of coefficient array, if present, must not be zero.";
    this.coefficient = coefficient.clone();
  }

  /**
   * Copy constructor.
   *
   * @param poly the polynomial to copy.
   * @throws NullPointerException if {@code poly} is {@code null}.
   */
  public DensePoly(DensePoly poly) throws NullPointerException {
    Objects.requireNonNull(poly, "The polynomial must not be null.");
    this.coefficient = poly.coefficient.clone();
  }

  /**
   * Initializes a polynomial given an array of coefficients.
   *
   * <p>The coefficient array can have trailing zeros, they will not be considered in the
   * representation of the polynomial.
   *
   * @param coefficient the array of coefficients.
   * @return the polynomial represented by the given array of coefficients.
   */
  public static DensePoly ofCoefficients(int[] coefficient) {
    Objects.requireNonNull(coefficient, "The coefficient array must not be null.");
    int degree = coefficient.length - 1;
    while (degree >= 0 && coefficient[degree] == 0) degree--;
    return degree == -1 ? ZERO : new DensePoly(Arrays.copyOf(coefficient, degree + 1));
  }

  /**
   * Initializes this to be the polynomial \(p = cx^n\).
   *
   * @param c the coefficient.
   * @param n the degree.
   * @return the polynomial \( p = cx^n \).
   * @throws IllegalArgumentException if {@code n} &lt; 0.
   */
  public static DensePoly ofCoefficientDegree(int c, int n) throws IllegalArgumentException {
    if (n < 0) throw new IllegalArgumentException("Can't create a monomial with negative exponent");
    if (c == 0) {
      return ZERO;
    } else {
      int[] coefficient = new int[n + 1];
      coefficient[n] = c;
      return new DensePoly(coefficient);
    }
  }

  /**
   * Checks whether this polynomial is the zero polynomial.
   *
   * @return {@code true} if this polynomial is the zero polynomial, {@code false} otherwise.
   */
  public boolean isZero() {
    return this == ZERO;
  }

  // Methods (identical to the previous implementation)

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
    return DensePoly.ofCoefficients(result);
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
    if (isZero() || q.isZero()) return ZERO;
    int[] result = new int[degree() + q.degree() + 1];
    for (int i = 0; i <= degree(); i++)
      for (int j = 0; j <= q.degree(); j++)
        result[i + j] = result[i + j] + coefficient[i] * q.coefficient[j];
    return DensePoly.ofCoefficients(result);
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
    return DensePoly.ofCoefficients(result);
  }

  // Methods inherited from Object

  @Override
  public String toString() {
    if (degree() > 0) {
      StringBuilder sb = new StringBuilder("DensePoly: ");
      int c = coeff(degree());
      if (c < -1) sb.append("-" + (-c));
      else if (c == -1) sb.append("-");
      else if (c > 1) sb.append(c);
      sb.append("x" + (degree() > 1 ? "^" + degree() : ""));
      for (int d = degree() - 1; d > 0; d--) {
        c = coeff(d);
        if (c == 0) continue;
        if (c < -1) sb.append(" - " + (-c));
        else if (c == -1) sb.append(" - ");
        else if (c == 1) sb.append(" + ");
        else sb.append(" + " + c);
        sb.append("x" + (d > 1 ? "^" + d : ""));
      }
      c = coeff(0);
      if (c > 0) sb.append(" + " + c);
      else if (c < 0) sb.append(" - " + (-c));
      return sb.toString();
    } else return "DensePoly: " + coeff(0);
  }

  @Override
  public int hashCode() {
    return Arrays.hashCode(coefficient);
  }

  @Override
  public boolean equals(Object obj) {
    if (this == obj) return true;
    if (!(obj instanceof DensePoly other)) return false;
    return Arrays.equals(coefficient, other.coefficient);
  }
}
