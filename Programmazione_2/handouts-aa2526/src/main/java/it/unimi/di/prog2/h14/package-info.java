/**
 * Materiale relativo alla lezione 14 per maggiori dettagli si veda il <a
 * href="https://prog2.di.unimi.it/diario">diario del corso</a>.
 *
 * <p>In this package there are several variants of a {@link DecimalDigits}, a class that can be
 * used to compute the digits of the decimal expansion of a given number.
 *
 * <p>All the variants are endowed in addition with an <em>iterator</em> (that is a method that
 * returns a <em>generator</em> in Liskov's parlance) on non zero digits; the variants correspond to
 * various possibilities with regard to the implementation style of the generator:
 *
 * <ul>
 *   <li>{@link DecimalDigitsEG} based on the <strong>external</strong> class {@link
 *       NonZeroDigitsGenerator},
 *   <li>{@link DecimalDigitsNSG} based on a <strong>nested static</strong> class,
 *   <li>{@link DecimalDigitsIG} based on an <strong>inner</strong> class,
 *   <li>{@link DecimalDigitsAG} based on an <strong>anonymous</strong> class.
 * </ul>
 *
 * <p>Such variants are listed in increasing order of preferableness (and implementation
 * complexity).
 */
package it.unimi.di.prog2.h14;
