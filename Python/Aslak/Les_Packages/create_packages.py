# Créer ses propres packages
# Si vous voulez créer vos propres packages, commencez par créer, dans le même dossier que votre programme Python, 
# un répertoire portant le nom du package.

# Dans ce répertoire, vous pouvez soit :

# - mettre vos modules, vos fichiers à l'extension ".py";
# - créer des sous-packages de la même façon, en créant un répertoire dans votre package.


# Le fichier d'initialisation

# En Python, vous trouverez souvent le fichier d'initialisation de package "__init__.py" dans un répertoire destiné à devenir un package.
# Ce fichier est optionnel depuis la version 3.3 de Python. 
# Vous n'êtes pas obligé de le créer mais vous pouvez y mettre du code d'initialisation pour votre package. 
# Je ne vais pas rentrer dans le détail ici (vous avez déjà beaucoup de choses à retenir), 
# mais sachez que ce code d'initialisation est appelé quand vous importez votre package.

# Un dernier exemple
# Voici un dernier exemple, que vous pouvez cette fois faire en même temps que moi pour vous assurer que cela fonctionne.
# Dans votre répertoire de code, là où vous mettez vos exemples Python, créez un fichier ".py" que vous appelerez "test_package.py"
# Créez dans le même répertoire un dossier package . Dedans, créez un fichier "fonctions.py" 
# dans lequel vous recopierez votre fonction "table".

# Dans votre fichier test_package.py, si vous voulez importer votre fonctiontable, vous avez plusieurs solutions :


from package.fonctions import table
table(5) # Appel de la fonction table

# Ou ...
import package.fonctions
package.fonctions.table(5)  # Appel de la fonction table


# Voilà. Il reste bien des choses à dire sur les packages mais je crois que vous avez vu l'essentiel. 
# Cette petite explication révélera son importance quand vous aurez à construire des programmes assez volumineux. 
# Évitez de tout mettre dans un seul module sans chercher à hiérarchiser, profitez de cette possibilité offerte par Python.


# En résumé
# On peut écrire les programmes Python dans des fichiers portant l'extension ".py".

# On peut créer des fichiers contenant des modules pour séparer le code.

# On peut créer des répertoires contenant des packages pour hiérarchiser un programme.