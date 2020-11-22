# Ici, nous allons former une étoile en utilisant des fonctions et des boucles
# Ici, nous utilisons for pour la boucle et une fonction

def increasing(n):    # Cette fonction est la première ligne pour les étoiles croissantes
    for i in range (0,n):
        for _ in range (0,i+1):
            print("*"),
        print

def decreasing(n):    # Cette fonction est la dernière ligne pour les étoiles décroissantes
    for i in range (1,n):
        for _ in range (n,i,-1):
            print("*"),
        print


n=int(input("Entrez le nombre de ligne que vous voulez, sachant que ça sera multiplié par 2 : "))
increasing(n)
decreasing(n)
