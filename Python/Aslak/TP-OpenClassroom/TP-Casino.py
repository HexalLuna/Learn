# TP OpenClassroom, un jeu de roulette de casino.
#
# La regle du jeu est :
#	- Le joueur mise sur un numéro compris entre 0 et 49 (50 numéros en tout). En choisissant son numéro, il y dépose la somme qu'il souhaite miser.
#	- La roulette est constituée de 50 cases allant naturellement de 0 à 49. Les numéros pairs sont de couleur noire, les numéros impairs sont de couleur rouge. Le croupier lance la roulette, lâche la bille et quand la roulette s'arrête, relève le numéro de la case dans laquelle la bille s'est arrêtée. C'est le numéro gagnant.
#	- Si le numéro gagnant est celui sur lequel le joueur a misé (probabilité de 1/50, plutôt faible), le croupier lui remet 3 fois la somme misée.
#	- Sinon, le croupier regarde si le numéro misé par le joueur est de la même couleur que le numéro gagnant (s'ils sont tous les deux pairs ou tous les deux impairs). Si c'est le cas, le croupier lui remet 50 % de la somme misée. 
#	- Si ce n'est pas le cas, le joueur perd sa mise.
#
# On notera que les sommes sont arrondites.

# On nous dis qu'il faut importer la fonction randrange du module de random
# On nous dis qu'il faut importer la fonction ceil du module math

from random import randrange
from math import ceil 

# Declaration des variables initiales
argent = 1000	# Somme initiale du joueur
continuerPartie = True	# Booleen de la boucle de jeu

print("""Bienvenu(e) au Casino d'Aslak ! Vous avez""",argent,"""€""")
print("""Vous vous installez à une roulette.\n""")

print("""Souhaitez vous connaître les règles du jeux ? o/n""")	# Test si le joueur veut connaitre les règles
testRegle = input()


if testRegle == "O" or testRegle == "o":	# Boucle de lecture des règles

	print("""Les règles du jeu sont :""")
	print("""- Vous misez sur un numéro compris entre 0 et 49 (50 numéros en tout). En choisissant le numéro, vous déposer la somme que vous souhaitez.""")
	print("""- Le croupier lance la roulette, lâche la bille et quand la roulette s'arrête, relève le numéro de la case dans laquelle la bille s'est arrêtée. C'est le numéro gagnant.""")
	print("""- Les numéros pairs sont de couleur noire, les numéros impairs sont de couleur rouge.""")
	print("""- Si le numéro gagnant est celui que vous aviez choisi, vous gagnez 3 fois la somme misée.""")
	print("""- Si le numéro gagnant est de la même couleur que le numéro que vous aviez choisi, vous gagnez 50% de la somme misé.""")
	print("""- Sinon vous perdre votre mise.""")
	
print("""\n""")	
	
while continuerPartie == True:	
	
	numeroMise = -1
	
	while numeroMise < 0 or numeroMise > 49:
	
		print("""Sur quel numéro souhaitez vous miser ? 0-49""")
		numeroMise = input()
		# print("""\n""")
	
		try:
			numeroMise = int(numeroMise)
		except ValueError:
			print("""Vous n'avez pas saisi de nombre.""")
			numeroMise = -1
			continue
	
		if numeroMise < 0:
		
			print("""Vous avez saisi un nombre inférieur à 0.""")
		
		if numeroMise > 49:
			
			print("""Vous avez saisi un nombre supérieur à 49.""")
		
	print("""Vous choississez le numéro""",numeroMise,""".\n""")
	
	
	argentMise = -1
		
	while argentMise <= 0 or argentMise > argent:
		
		print("""Combien souhaitez vous miser ?""")
		argentMise = input()
		# print("""\n""")
		
		try:
			argentMise = int(argentMise)
		except ValueError:
			print("""Vous n'avez pas saisi de nombre.""")
			argentMise = -1
			continue
	
		if argentMise > argent:
			print("""Vous ne pouvez miser plus que la totalité de votre argent.""")
		
		if argentMise <= 0:
			print("""Vous ne pouvez pas miser une somme nulle ou négative.""")
	
	print("""Vous misez""",argentMise,"""€.\n""")
	
	
	numeroGagnant = randrange(50)
	
	if numeroGagnant%2 == 0:
		print("""Le numéro gagnant est le""",numeroGagnant,"""noir.""")
		
	else:
		print("""Le numéro gagnant est le""",numeroGagnant,"""rouge.""")
	
	if numeroGagnant == numeroMise:
		argent += argentMise*3
		print("""Bien jouer ! Vous avez choisi le bon numéro ! Vous gagnez""",argentMise*3,"""€.\n""")
	elif numeroGagnant%2 == numeroMise%2:
		argent += ceil(argentMise*0.5)
		print("""Bien jouer ! Vous avez choisi la bonne couleur ! Vous gagnez""",ceil(argentMise*0.5),"""€.\n""")
	else:
		argent += -argentMise
		print("""Perdu. Vous n'avez ni le bon numéro ni la bonne couleur.\n""")
		
	if argent <= 0:
		print("""Désolé vous n'avez plus d'argent. Merci d'avoir joué !.""")
		continuerPartie = False
		
	else:
		print("""Vous avez""",argent,"""€""")
		print("""Voulez vous continuez ? o/n""")	# Test si joueur souhaite ou non continuer à jouer
		testContinuerPartie = input()
		print("""\n""")
	
		if testContinuerPartie == """n""" or testContinuerPartie == """N""":
		
			print("""Vous quittez le casino avec""",argent,"""€. Revenez quand vous le voulez !""")
			continuerPartie = False
# Programme Fini.
