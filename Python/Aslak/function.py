# Ici nous apprenons à déclarer une fonction et l'appeler

# Nous utilisons le mot-clé 'def' pour déclarer une fonction

def function():
     print("Tu es dans une fonction")

function()

# Nous pouvons passer 'paremetre' dans une fonction comme ici

def parametre(n,m):
    for _ in range(0,n): # La fonction range () renvoie une séquence de nombres, commençant par 0 par défaut et incrémentée de 1 (par défaut), et s'arrête avant un nombre spécifié.
        for _ in range(0,m):
            print("*"),  # Dans Python 3.x au lieu de "," nous utilisons la fonction end = ""
        print ("")
parametre(2,2)

# Ici nous partons d'une matrice de 2 par 2