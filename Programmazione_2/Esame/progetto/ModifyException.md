#Eccezzioni attuali:
1. Capacity
2. EmptyRail
3. InsufficentChange
4. Insufficentcoins
5. Insufficentvalue
6. InvalidImporto
7. InvalidItem
8. InvaldResult
9. Moneta
10. Slot
11. Taglia
12. TotalValue

#Eccezzioni Nuove da implementare:

##QuantityException
Prende il posto di capacity e insufficentcents dato che sono entrambi delle mancanze di quantità

##InvalidResult
Sotto di questa abbiamo anche InsufficenValue, che da togliere, dato che indicano un valore negativo

##Totalvalue
Prende anche InsufficentChange perchè se non c'è abbastanza valore nella cassa e non si può dare il resto

##SlotException
Prende sia slot che InvalidItem perchè è sempre un tentativo di inserire un prodotto troppo grosso o diverso da quello presente
