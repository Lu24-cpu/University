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

import it.unimi.di.prog2.h14.member.Watch.Gear;
import java.util.List;

/** A simple client for the {@link Watch} class. */
public class WatchClient {

  /** . */
  private WatchClient() {}

  /**
   * The main method.
   *
   * @param args the command line arguments (ignored).
   */
  public static void main(String[] args) {
    Watch miniWatch = new Watch("Miniature", List.of(new Gear(2.0), new Gear(3.0)));
    Watch bigWatch = new Watch("Giant", List.of(new Gear(20.0), new Gear(40.0), new Gear(60.0)));
    for (int time = 0; time < 100; time++) {
      miniWatch.tick();
      bigWatch.tick();
    }
    System.out.println(miniWatch);
    System.out.println(bigWatch);
  }
}
