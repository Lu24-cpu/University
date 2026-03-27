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

import customException.InvalidImportoException;
import customException.InvalidResultException;
import macchinetta.Importo;

public class OperazioniImporti {
    public static void main(String[] args) {
        try(Scanner sc = new Scanner(System.in)) {
            while (sc.hasNext()) {
                String riga = sc.nextLine().replaceAll("\\s+", " ");
                String[] elementi = riga.split(" ");

                try{
                    Importo value1 = Parser.parseImporto(elementi[0]);
                    Importo value2;
                    Importo result = new Importo(0);
                    int value = 0;
                    switch (elementi[1]) {
                        case "+":
                            value2 = Parser.parseImporto(elementi[2]);
                            result = value1.Add(value2);
                            break;
                        case "-":
                            value2 = Parser.parseImporto(elementi[2]);
                            result = value1.Sub(value2);
                            break;
                        case "*":
                            result = value1.Mul(Integer.parseInt(elementi[2]));
                            break;
                        case "/":
                            value2 = Parser.parseImporto(elementi[2]);
                            value = value1.Div(value2);
                            break;
                        default:
                            System.out.println("invalid-operand");
                    }
                    if (value != 0) System.out.println(value);
                    else System.out.println(result.toString());
                } catch (InvalidResultException e) {
                    System.out.println("negative-result");
                } catch (InvalidImportoException e) {
                    System.out.println("invalid-result");
                }
            }
        }
    }
}
