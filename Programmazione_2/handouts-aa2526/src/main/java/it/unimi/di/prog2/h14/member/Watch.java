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

package it.unimi.di.prog2.h14.member;

import java.util.Collections;
import java.util.Iterator;
import java.util.List;
import java.util.Objects;
import java.util.StringJoiner;

/** A watch with gears. */
public class Watch implements Iterable<Watch.Gear> {

  /**
   * A gear of a watch.
   *
   * <p>The gear can "age" in the sense that it keeps track of how many rotations it has done.
   */
  public static class Gear {

    /** The diameter of this gear, must be positive. */
    private final double diameter;

    /** The number of rotations this gear has done., must be non-negative. */
    private int rotations;

    /**
     * Creates a gear with the given diameter and number of teeth.
     *
     * @param diameter the diameter of this gear.
     * @throws IllegalArgumentException if diameter is not positive.
     */
    public Gear(double diameter) {
      if (diameter <= 0) throw new IllegalArgumentException();
      this.diameter = diameter;
      rotations = 0;
    }

    /** Simulates a tick of this gear, increasing its number of rotations by one. */
    public void tick() { // this makes the gear mutable
      rotations++;
    }

    /**
     * Returns the number of rotations this gear has done.
     *
     * @return the number of rotations.
     */
    public int rotations() {
      return rotations;
    }

    @Override
    public String toString() {
      return "Gear<diameter=" + diameter + ", rotation(s)=" + rotations + ">";
    }
  }

  /** The model of this watch, must be non-empty (and not null). */
  public final String model;

  /** The gears of this watch, must be non-null and not contain null. */
  public final List<Gear> gears;

  /**
   * Creates a watch with the given model and gears.
   *
   * @param model the model of the watch.
   * @param gears the gears of the watch.
   * @throws IllegalArgumentException if model is empty.
   * @throws NullPointerException if model or gears is null or if gears contains null.
   */
  public Watch(String model, List<Gear> gears) {
    if (Objects.requireNonNull(model).isEmpty()) throw new IllegalArgumentException();
    this.model = model;
    this.gears = List.copyOf(gears);
  }

  /** Simulates a tick of this watch, increasing the number of rotations of all its gears by one. */
  public void tick() {
    for (Gear gear : gears) gear.tick();
  }

  @Override
  public Iterator<Gear> iterator() {
    return Collections.unmodifiableCollection(gears).iterator();
  }

  @Override
  public String toString() {
    StringJoiner sj = new StringJoiner("\n- ", model + "\n- ", "");
    for (Gear gear : gears) {
      sj.add(gear.toString());
    }
    return sj.toString();
  }
}
