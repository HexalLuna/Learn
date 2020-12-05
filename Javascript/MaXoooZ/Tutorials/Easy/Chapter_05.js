// Chapter 4

// CONDITIONS
function fr() {
    /* 
    	Bonjour, bienvenue sur mon avancée sur le javascript (que j'ai déjà appris entièrement).
    	Bonne apprentissage à vous et j'espère vous revoir dev fullstack :').

    	Bon avant de pouvoir executer du code, je vous propose un site pour pouvoir executer vos lignes de code !
    	https://jsbin.com/juwokohedi/edit?js,console
    
    	(Tout est retrouvable dans le README.md)
    */

    /*
    	L'instruction if nous permet de vérifier si une expression est égale à vrai ou faux, et d'exécuter
    	un code différent en fonction du résultat.

    	Par exemple, si nous voulons demander à l'utilisateur si son nom est "Martin", nous pouvons utiliser la fonction de confirmation.
    */

    $name = "Martin";

    if ($name === "Martin" /*La majuscule est importante!*/ ) {
        // Oui il s'appelle bien Martin
    } else {
        // Non il ne s'appelle pas Martin
    }


    // Nous pouvons aussi comparer plusieurs nombres!

    $numberOne = 7;
    $numberTwo = 10;

    if ($numberTwo >= $numberOne) { // ">=" = est ou est plus grand
        // Est plus grand
    }

    // Si vous avez besoin de verifier plusieurs fois quelque chose vous pouvez le faire!

    $name = "Julie";

    switch ($name) {
        case "Julie":
            // Le nom est Julie donc ici ecrivez votre code
            break;

        case "Martin":
            // Le nom n'est pas Julie
            break;

        default:
            // Le nom est ni Julie ni Martin !
            break;
    }

    /*
     	Vous avez plein de chose à apprendre sur les conditions mais pas tout est listable! Donc je vous propose la documentation de firefox.
    	https://developer.mozilla.org/fr/docs/Web/JavaScript/Reference/Instructions/if...else
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
    	The if statement allows us to check if an expression is equal to true or false, and to execute
    	a different code depending on the result.

    	For example, if we want to ask the user if his name is "Martin", we can use the confirmation function.
    */

    $ name = "Martin";

    if ($ name === "Martin" /* Uppercase is important! */ ) {
        // Yes his name is Martin
    } else {
        // No his name is not Martin
    }

    // We can also compare several numbers!

    $ numberOne = 7;
    $ numberTwo = 10;

    if ($ numberTwo > = $ numberOne) { // "> =" = is or is greater
        // Is bigger
    }

    // If you need to check something more than once you can do it!

    $ name = "Julie";

    switch ($ name) {
        box "Julie":
            // The name is Julie so here write your code
            break;

        case "Martin":
            // The name is not Julie
            break;

        default:
            // The name is neither Julie nor Martin!
            break;
    }

    /*
    	You have a lot to learn about the conditions but not everything is listed! So I offer you the documentation of firefox.
    	https://developer.mozilla.org/fr/docs/Web/JavaScript/Reference/Instructions/if...else
    */
}