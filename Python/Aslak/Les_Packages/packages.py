# Importer des packages
# Si vous voulez utiliser, dans votre programme, la bibliothèque fictive que nous venons de voir,
# vous avez plusieurs moyens qui tournent tous autour des mots clés from et import :

import nom_bibliotheque

# Cette ligne importe le package contenant la bibliothèque.
# Pour accéder aux sous-packages,
# vous utiliserez un point « . » afin de modéliser le chemin menant au module ou à la fonction que vous voulez utiliser :

nom_bibliotheque.evenements # Pointe vers le sous-package evenements
nom_bibliotheque.evenements.clavier # Pointe vers le module clavier

# Si vous ne voulez importer qu'un seul module (ou qu'une seule fonction) d'un package,
# vous utiliserez une syntaxe similaire, assez intuitive :

from nom_bibliotheque.objets import bouton

# En fonction des besoins, vous pouvez décider d'importer tout un package, un sous-package, un sous-sous-package… 
# ou bien juste un module ou même une seule fonction. Cela dépendra de vos besoins.