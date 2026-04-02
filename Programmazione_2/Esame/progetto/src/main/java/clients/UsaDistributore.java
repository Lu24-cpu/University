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
import macchinetta.DistributoreAutomatico;

public class UsaDistributore {

    public static void main(String[] args) {
        try (Scanner sc = new Scanner(System.in)) {
            String specs_rails = sc.nextLine();
            DistributoreAutomatico macchinetta;

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
                
                macchinetta = Parser.parseDistributore(specs_rails, inputs, sc.nextLine());
            } catch (InvalidImportoException | InsufficentcoinsException | TotalvalueException | InvalidResultException | MonetaException e) {
                System.out.println(e.getMessage());
            }
            
            while (sc.hasNext()) {
                String[] elements = sc.nextLine().trim().split(" ! ");
                String[] specifiche = elements[1].split(";");
                StringBuilder prodotto = new StringBuilder();
                String[] q = elements[0].split(" ");
                try {

                    if (elements[0].contains("+")) {
                        int quantity = Integer.parseInt(q[1]);

                        prodotto.append(specifiche[0]).append("|").append(specifiche[1]).append("|").append(specifiche[2]);

                        try {
                            macchinetta.caricaBinario(parseProdotto(prodotto.toString()), quantity);
                        } catch (TagliaException | InvalidImportoException e) {
                            System.out.println(e.getMessage());
                        }
                    } else {
                        int rail = Integer.parseInt(q[1]);
                        Aggregato importo = new Aggregato();

                        importo = Parser.parseAggregato(importo, q);

                        return "- " + macchinetta.scaricaBinario(rail, importo);
                    }

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
