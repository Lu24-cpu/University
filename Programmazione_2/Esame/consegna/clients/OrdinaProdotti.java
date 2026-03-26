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

import customException.InvalidImportoException;
import customException.TagliaException;
import macchinetta.Prodotto;

public class OrdinaProdotti {

    public static void main(String[] args) {
        ArrayList<Prodotto> products = new ArrayList<>();

        try(Scanner sc = new Scanner(System.in)) {
            while (sc.hasNext()) {
                try {
                    products.add(Parser.parseProdotto(sc.nextLine()));
                } catch (InvalidImportoException | TagliaException e) {
                    System.out.println(e.getMessage());
                }
            }
        }

        for (int i = 0; i < products.size(); i++) {
            for (int j = i + 1; j < products.size(); j++) {
                Prodotto prodotto1 = products.get(i), prodotto2 = products.get(j);

                if (prodotto1.size().equals(prodotto2.size()) && prodotto1.value().getTotalCents() > prodotto2.value().getTotalCents()) {
                    products.set(i, prodotto2);
                    products.set(j, prodotto1);
                } else if (prodotto1.size().compareTo(prodotto2.size()) > 0) {
                    products.set(i, prodotto2);
                    products.set(j, prodotto1);
                }
            }
        }

        for (Prodotto product : products) {
            System.out.println(product.toString());
        }
    }
}
