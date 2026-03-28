import numpy as np
import matplotlib.pyplot as plt

sprint = [13.24, 13.13]
longrun10 = [57.16, 60.56, 56.31, 59.17, 59.26, 52.46, 54.05, 52.01, 55.42, 60.39]

x = np.arange(len(longrun10)).transpose()
y = longrun10
plt.rc('figure', figsize=(10, 70))
plt.bar(x, y)
plt.show()
