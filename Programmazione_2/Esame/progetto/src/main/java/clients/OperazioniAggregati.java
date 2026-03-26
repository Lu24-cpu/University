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
import macchinetta.*;

public class OperazioniAggregati {

    public static void main(String[] args) {
        Aggregato cassa = new Aggregato();
        
        try (Scanner sc = new Scanner(System.in)) {
            while (sc.hasNext()) {
                String riga = sc.nextLine();

                try {
                    if (riga.contains("+") || !riga.contains("-")) cassa = Parser.parseAggregato(cassa, riga);
                    else {
                        String[] specifiche = riga.split(", ");

                        int i = 0;
                        for (String s : specifiche) {
                            s = s.replaceAll("\\s+", " ");
                            String[] elements = s.split(" ");
                            if (i == 0) cassa.Remove(Moneta.moneta(Parser.parseImporto(elements[elements.length - 1])), Integer.parseInt(elements[1]));
                            else cassa.Remove(Moneta.moneta(Parser.parseImporto(elements[elements.length - 1])), Integer.parseInt(elements[0]));

                            i++;
                        }
                    }
                    System.out.println(cassa.toString());
                }  catch (InsufficentcoinsException e) {
                    System.out.println("missing-coins");
                } catch (TotalvalueException e) {
                    System.out.println("missing-value");
                } catch (InvalidImportoException | InvalidResultException | MonetaException | NumberFormatException e) {
                    System.out.println(e.getMessage());
                }
            }
        }
    }
}
