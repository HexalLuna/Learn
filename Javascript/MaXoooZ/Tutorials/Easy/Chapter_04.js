// Chapter 4

// OPERATORS
function fr() {
    /* 
    	Bonjour, bienvenue sur mon avancée sur le javascript (que j'ai déjà appris entièrement).
    	Bonne apprentissage à vous et j'espère vous revoir dev fullstack :').

    	Bon avant de pouvoir executer du code, je vous propose un site pour pouvoir executer vos lignes de code !
    	https://jsbin.com/juwokohedi/edit?js,console
    
    	(Tout est retrouvable dans le README.md)
    */

    /*
    	Chaque variable dans JavaScript est convertie automatiquement de sorte que tout opérateur entre
    	deux variables donnera toujours une sorte de résultat.


    	L'opérateur + (addition) est utilisé à la fois pour l'addition et la concaténation de chaînes.
    	Par exemple, ajouter deux variables est simple:
    */

    var a = 1;
    var b = 2;
    var c = a + b; // C est maintenant égale à 3 !

    // L'opérateur d'addition est utilisé pour concaténer des chaînes en chaînes, des chaînes en nombres et des nombres en chaînes:

    var name = "bébé";
    console.log("Salut " + name + " !"); // Affiche : Salut bébé !
    console.log("Le sens de la vie est " + 42); // Affiche : Le sens de la vie est 42
    console.log(42 + " est le sens de la vie"); // Affiche :  42 est le sens de la vie

    /*
     	JavaScript se comporte différemment lorsque vous essayez de combiner deux opérandes de types différents. 
    	La valeur primitive par défaut est une chaîne, donc lorsque vous essayez d'ajouter un nombre à une chaîne, JavaScript 
    	transformera le nombre en chaîne avant la concaténation.
    */

    console.log(1 + "1"); // Affiche : 11

    // Opérateurs mathématiques : 

    console.log(3 - 5); // Affiche : -2
    console.log(3 * 5); // Affiche : 15
    console.log(3 / 5); // Affiche : 0.6 (Vous pouvez utiliser % aussi à la place de / car javascript le supporte)

    /*
    	JavaScript prend également en charge les opérateurs d'affectation et d'opération combinés. 
    	Donc, au lieu de taper myNumber = myNumber / 2, vous pouvez taper monNumber / = 2.

    	Voici une liste de tous ces opérateurs:
    */

    /*
     	- /=
    	- *=
    	- -=
    	- +=
    	- \%\=
    */
}

function en() {
    /*
    	Hello, welcome to my progress on javascript (which I have already learned completely).
    	Happy learning to you and I hope to see you again dev fullstack: ').

    	Good before being able to execute code, I suggest you a site to be able to execute your lines of code!
    	https://jsbin.com/juwokohedi/edit?js,console

    	(Everything can be found in the README.md)
    */

    /*
    	Every variable in JavaScript is converted automatically so that any operator enters
    	two variables will always give some sort of result.


    	The + (addition) operator is used for both addition and concatenation of strings.
    	For example, adding two variables is simple:
    */

    var a = 1;
    var b = 2;
    var c = a + b; // C is now equal to 3!

    // The addition operator is used to concatenate strings into strings, strings into numbers, and numbers into strings:

    var name = "baby";
    console.log("Hi" + name + "!"); // Poster: Hi baby!
    console.log("The meaning of life is" + 42); // Poster: The meaning of life is 42
    console.log(42 + "is the meaning of life"); // Displays: 42 is the meaning of life

    /*
    	JavaScript behaves differently when you try to combine two operands of different types.
    	The default primitive is a string, so when trying to append a number to a string, JavaScript
    	will transform the number into a string before the concatenation.
    */

    console.log(1 + "1"); // Displays: 11

    // Mathematical operators:

    console.log(3 - 5); // Displays: -2
    console.log(3 * 5); // Displays: 15
    console.log(3 / 5); // Displays: 0.6 (You can also use% instead of / because javascript supports it)

    /*
    	JavaScript also supports combined assignment and operation operators.
    	So instead of typing myNumber = myNumber / 2, you can type myNumber / = 2.

    	Here is a list of all these operators:
    */

    /*
    	- / =
    	- * =
    	- - =
    	- + =
    	- \% \ =
    */
}