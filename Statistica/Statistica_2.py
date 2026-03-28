import numpy as np
import matplotlib.pyplot as plt

names = ['Aquaman', 'Ant-Man', 'Batman', 'Black Widow',
         'Captain America', 'Daredavil', 'Elektra', 'Flash',
         'Green Arrow', 'Human Torch', 'Hancock', 'Iron Man',
         'Mystique', 'Professor X', 'Rogue', 'Superman',
         'Spider-Man', 'Thor', 'Northstar']

years = [1941, 1962, None, None, 1941,
         1964, None, 1940, 1941, 1961,
         None, 1963, None, 1963, 1981,
         None, None, 1962, 1979]

def get_sort_counts(sequence):
    counts = {}

    for x in sequence:
        if x in counts:
            counts[x] += 1
        else:
            counts[x] = 1
    
    pairs = counts.items()
    return sorted(pairs, key=lambda p:p[1], reverse=True)

x, y = np.array(get_sort_counts(years)[1:]).transpose()
plt.rc('figure', figsize=(5.0, 3.0))
plt.bar(x, y)
plt.show()