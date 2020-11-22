# Nous allons apprendre à compléter notre bloc "try".
# Comme je l'ai indiqué plus haut, la forme minimale est à éviter pour plusieurs raisons.

# D'abord, elle ne différencie pas les exceptions qui pourront être levées dans le bloc "try". 
# Ensuite, Python peut lever des exceptions qui ne signifient pas nécessairement qu'il y a eu une erreur.


# Exécuter le bloc except pour un type d'exception précis

# Dans l'exemple que nous avons vu plus haut, 
# on ne pense qu'à un type d'exceptions susceptible d'être levé : le type "ValueError", 
# qui trahirait une erreur de conversion. Voyons un autre exemple :

try:
    resultat = numerateur / denominateur
except:
    print("Une erreur est survenue... laquelle ?")

# Ici, plusieurs erreurs sont susceptibles d'intervenir, chacune levant une exception différente.

# - NameError: l'une des variables "numerateur" ou "denominateur" n'a pas été définie (elle n'existe pas).
#  Si vous essayez dans l'interpréteur l'instruction "print(numerateur)" alors que vous n'avez pas défini la variable "numerateur", 
# vous aurez la même erreur.

# - TypeError: l'une des variables" numerateur" ou" denominateur" ne peut diviser ou être divisée 
# (les chaînes de caractères ne peuvent être divisées, ni diviser d'autres types, par exemple). 
# Cette exception est levée car vous utilisez l'opérateur de division « / » sur des types qui ne savent pas quoi en faire.

# - ZeroDivisionError: encore elle ! Si "denominateur" vaut 0, cette exception sera levée.

# Tout se joue sur la ligne du "except".
#  Entre ce mot-clé et les deux points, vous pouvez préciser le type de l'exception que vous souhaitez traiter.

try:
    resultat = numerateur / denominateur
except NameError:
    print("La variable numerateur ou denominateur n'a pas été définie.")

# Ce code ne traite que le cas où une exception "NameError" est levée. 
# On peut intercepter les autres types d'exceptions en créant d'autres blocsexceptà la suite :

try:
    resultat = numerateur / denominateur
except NameError:
    print("La variable numerateur ou denominateur n'a pas été définie.")
except TypeError:
    print("La variable numerateur ou denominateur possède un type incompatible avec la division.")
except ZeroDivisionError:
    print("La variable denominateur est égale à 0.")

# Allez un petit dernier !

# On peut capturer l'exception et afficher son message grâce au mot-cléasque vous avez déjà vu dans un autre contexte
# (si si, rappelez-vous de l'importation de modules).

try:
    # Bloc de test
except type_de_l_exception as exception_retournee:
    print("Voici l'erreur :", exception_retournee)

# Dans ce cas, une variable "exception_retournee" est créée par Python si une exception du type précisé est levée dans le bloc "try".