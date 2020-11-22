# ici nous allons voir la boucle for

# dans la boucle for, nous utilisons la fonction qui se nomme range
# cela fonctionne comme ça : range(initiale, finale, incrémentation ou décrémentation)

for i in range(0,9,1):
    print("i=",i)
    print("--------------------------------------")

# ici, nous obtenons la réponse i = 9 car dans la fonction range, l'indice final n'est pas pris

for i in range(0,9,2):
    print("i=",i)
    print("--------------------------------------")
# Ici, nous obtenons toujours la même réponse

for i in range(9,0,-1):
    print("i=",i)
    print("--------------------------------------")
# ici nous obtenons l'ordre inverse