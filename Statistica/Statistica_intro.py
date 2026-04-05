#Rifacimento intro a Python

# Dichiarazione di variabili:
x = 3
y = 3.14
z = 'ciao'
w = "come"
k = '''stai?'''

print(type(x), type(y), type(z), type(w), type(k), z[3])

#Dichiarazione liste
iron_man = [
    'Iron Man',
    'Tony Stark',
    'Long Island, New York',
    'Marvel Comics',
    198.51,
    191.85,
    'M',
    1963,
    'Blue',
    'Black',
    85,
    'high'
]

#Dichiarazione tuple
rogue = (
    'Rogue',
    'Anna Marie',
    'Caldecott County, Mississippi',
    'Marvel Comics',
    173.1,
    54.39,
    'F',
    1981,
    'Green',
    'Brown / White',
    10,
    'good'
)

#Dichiarazione dizionari
rogue_one = {
    'name': 'Rogue',
    'identity': 'Anna Marie',
    'birth_place': 'Caldecott County, Mississippi',
    'publisher': 'Marvel Comics',
    'height': 173.1,
    'weight': 54.39,
    'gender': 'F',
    'first_appearance': 1981,
    'eye_color': 'Green',
    'hair_color': 'Brown / White',
    'strength': 10,
    'intelligence': 'good'
}

#Utilizzi delle liste e tuple:
print(iron_man[1])
print(iron_man[-2:])
print(iron_man[4:6])

#liste possono avere delle variazioni nei loro elementi
iron_man[0] = "I'm Iron Man"
print(iron_man[0])

#I dizionari non vengono richiamati tramite l'indice ma tramite la chiave
print(rogue_one['identity'])