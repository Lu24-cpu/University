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

import customException.EmptyRailException;
import customException.InsufficentChangeException;
import customException.InsufficentValueException;
import customException.InsufficentcoinsException;
import customException.InvalidImportoException;
import customException.InvalidResultException;
import customException.MonetaException;
import customException.SlotException;
import customException.TagliaException;
import customException.TotalvalueException;
import macchinetta.Aggregato;
import macchinetta.Binario;
import macchinetta.DistributoreAutomatico;
import macchinetta.StrategiaResto;

public class UsaDistributore {

    public static void main(String[] args) {
        Aggregato cassa = new Aggregato();
        ArrayList<Binario> rails = new ArrayList<>();

        try (Scanner sc = new Scanner(System.in)) {
            try { rails = Parser.parseBinari(sc.nextLine()); }
            catch (TagliaException e) { System.out.println(e.getMessage()); }

            try {
                StringBuilder inputs = new StringBuilder();

                while (sc.hasNext()) {
                    String line = sc.nextLine().trim();

                    if (line.equals(".")) break;
                    line = line.replaceAll("\\s+", " ");
                    String[] input = line.split(" ");

                    line = String.join(" x ", input);
                    inputs.append(line).append(", ");
                }
                if (inputs.length()>2) {
                    inputs.setLength(inputs.length() - 2);
                }

                String[] input = inputs.toString().split(", ");

                cassa.Insert(Parser.parseAggregato(new Aggregato(), input));
            } catch (InvalidImportoException | InsufficentcoinsException | TotalvalueException | InvalidResultException | MonetaException e) {
                System.out.println(e.getMessage());
            }
            
            StrategiaResto strategy = Parser.parseStrategia(sc.nextLine());
            DistributoreAutomatico macchinetta = new DistributoreAutomatico(cassa, strategy, rails);

            while (sc.hasNext()) {
                String line = sc.nextLine().trim();
                try {
                    if (line.equals("?")) System.out.println(macchinetta);
                    else System.out.println(Parser.parseDistributore(macchinetta, line));
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
