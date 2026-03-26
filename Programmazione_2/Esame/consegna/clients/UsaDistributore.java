/*

Copyright 2024 Massimo Santini

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

package clients;

import java.util.ArrayList;
import java.util.Scanner;

import customException.*;
import macchinetta.*;

public class UsaDistributore {

    public static void main(String[] args) {
        Aggregato cassa = new Aggregato();
        ArrayList<Binario> rails = new ArrayList<>();

        try (Scanner sc = new Scanner(System.in)) {
            try { rails = Parser.parseBinari(sc.nextLine()); }
            catch (TagliaException e) { System.out.println(e.getMessage()); }

            while (sc.hasNextLine()) {
                String line = sc.nextLine().trim();
                
                if (line.equals(".")) break;
                try { cassa = Parser.parseAggregato(cassa, line); } 
                catch (InvalidImportoException | InvalidResultException | MonetaException e) { System.out.println(e.getMessage()); }
            }

            StrategiaResto strategy = Parser.parseStrategia(sc.nextLine());
            DistributoreAutomatico macchinetta = new DistributoreAutomatico(cassa, strategy, rails);

            while (sc.hasNext()) {
                String line = sc.nextLine().trim();
                try {
                    System.out.println(Parser.parseDistributore(macchinetta, line));
                }  catch (EmptyRailException e) {
                    System.out.println("- empty");
                } catch (InsufficentChangeException e) {
                    System.out.println("- change");
                } catch (InsufficentValueException e) {
                    System.out.println("- value");
                } catch (SlotException e) {
                    System.out.println("- slot");
                }
            }
        }
    }

}
