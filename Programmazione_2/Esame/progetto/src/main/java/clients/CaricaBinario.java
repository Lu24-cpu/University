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

import java.util.Scanner;

import customException.*;
import macchinetta.Binario;
import macchinetta.Prodotto;

public class CaricaBinario {

    public static void main(String[] args) {
        
        try{
            Binario rail = Parser.parseBinario(args[1], args[0]);
            System.out.println(rail.toString());
            try(Scanner sc = new Scanner(System.in)) {
                while (sc.hasNext())  {
                    String[] specifiche = sc.nextLine().split("; ");
                    int quantity = Integer.parseInt(specifiche[0]);
                    String[] prodotto = specifiche[1].split(" @ ");
                    String pr = String.join("|", prodotto);
                    Prodotto product = Parser.parseProdotto(pr);

                    try {
                        rail.uploadRail(product, quantity);
                        System.out.println(rail.toString());
                    } catch (CapacityException e) {
                        System.out.println("capacity");
                    } catch (TagliaException e) {
                        System.out.println("size");
                    } catch (SlotException e) {
                        System.out.println("item");
                    }
                }
            }
        } catch (InvalidImportoException | TagliaException | NumberFormatException e) {
            System.out.println(e.getMessage());
        }
    }
}
