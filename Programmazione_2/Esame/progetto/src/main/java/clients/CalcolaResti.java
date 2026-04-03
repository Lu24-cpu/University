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


public class CalcolaResti {

    public static void main(String[] args) {
        StrategiaResto change;

        try{
            Importo value = Parser.parseImporto(args[1]);

            try(Scanner sc = new Scanner(System.in)) {
                while (sc.hasNext()) {
                    switch (args[0]) {
                        case "H":
                            change = new RestoAlto();
                            break;
                        case "L":
                            change = new RestoBasso();
                            break;
                        case "P":
                            change = new RestoMedio();
                            break;
                        default:
                            throw new AssertionError();
                    }

                    Aggregato cassa = new Aggregato();
                    Aggregato resto;
                    String riga = sc.nextLine().replaceAll("\\, ", ", ");
                    
                    try {
                        riga = riga.replaceAll("\\s+", " ");
                        String[] input = riga.split("; ");
                        input[0] = input[0].replaceAll("\\+ ", "");
                        
                        cassa.Insert(Parser.parseAggregato(new Aggregato(), input));

                        resto = change.Resto(cassa, value);
                        System.out.println(resto);
                    } catch (InsufficentChangeException | InsufficentcoinsException e) {
                        System.out.println("change-not-possible");
                    } catch (InsufficentValueException | TotalvalueException e) {
                        System.out.println("insufficient-value");
                    } catch (InvalidImportoException | InvalidResultException | MonetaException e) {
                        System.out.println(e.getMessage());
                    }
                }
            }
        }catch (InvalidImportoException e) {
            System.out.println("invalid-result");
        }
    }
}
