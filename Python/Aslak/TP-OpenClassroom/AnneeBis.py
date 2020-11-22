import os

annee = input("Saisissez une année : ")

try:
    annee = int(annee)
except:
    print("Erreur lors de la conversion de l'année (un nombre est attendu) !")


if annee % 400 == 0:
    bissextile = True
elif annee % 100 == 0:
    bissextile = False
elif annee % 4 == 0:
    bissextile = True
else:
    bissextile = False

if bissextile:
    print("L'année ", annee, " est bissextile !")
else:
    print("L'année ", annee, " n'est pas bissextile !")

os.system("pause")
