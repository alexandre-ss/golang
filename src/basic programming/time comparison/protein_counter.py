import time
start_time = time.time()
file = open("data.dat").readlines()
file = file[0]

char = [protein for protein in file]

A = C = G = T = 0
for i in range(len(char)):
  if char[i] == "A":
    A = A + 1
  elif char[i] == "C":
    C = C + 1
  elif char[i] == "G":
    G = G + 1
  elif char[i] == "T":
    T = T + 1
print("A = {}\nC = {}\nG = {}\nT = {}".format(A, C, G, T))
print("The program took {}us to execute".format((time.time() - start_time)*(10**6)))

