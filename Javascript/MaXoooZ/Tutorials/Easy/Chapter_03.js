// Chapter 3

// ARRAY
function fr() {
    /* 
    	Bonjour, bienvenue sur mon avancée sur le javascript (que j'ai déjà appris entièrement).
    	Bonne apprentissage à vous et j'espère vous revoir dev fullstack :').

    	Bon avant de pouvoir executer du code, je vous propose un site pour pouvoir executer vos lignes de code !
    	https://jsbin.com/juwokohedi/edit?js,console
    
    	(Tout est retrouvable dans le README.md)
    */

    /*
    	L'objet global Array est utilisé pour créer des tableaux.
    	Les tableaux sont des objets de haut-niveau (en termes de complexité homme-machine) semblables à des listes.

    	Créer un tableau:
    */

    var names = ["Galzron", "Aslak", "MaXoooZ"];
    console.log(names.length); // 3 car le tableau contient 3 strings

    // Accéder (via son index) à un élément du tableau

    var first = names[0]; // Le premier élément du tableau
    var last = names[names.length - 1]; // Le dernière élément du tableau

    // Ajouter une chaine de charactère ou un booléen au tableau (à la fin)

    var newLength = names.push("You");
    // ["Galzron", "Aslak", "MaXoooZ", "You"];

    // Puis supprimer le dernier élément du tableau

    var last = names.pop(); // Supprime le dernier élément du tableau
    // ["Galzron", "Aslak", "MaXoooZ"];

    // Retrouve plus d'informations ici : https://developer.mozilla.org/fr/docs/Web/JavaScript/Reference/Objets_globaux/Array x)
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
    	The global Array object is used to create arrays.
    	Tables are high-level objects (in terms of human-machine complexity) similar to lists.

    	Create a table:
    */

    var names = ["Galzron", "Aslak", "MaXoooZ"];
    console.log(names.length); // 3 because the array contains 3 strings

    // Access (via its index) an element of the array

    var first = names[0]; // The first element of the array
    var last = names[names.length - 1]; // The last element of the array

    // Add a character string or a boolean to the array (at the end)

    var newLength = names.push("You");
    // ["Galzron", "Aslak", "MaXoooZ", "You"];

    // Then delete the last element of the array

    var last = names.pop(); // Remove the last element from the array
    // ["Galzron", "Aslak", "MaXoooZ"];

    // Find more information here : https://developer.mozilla.org/fr/docs/Web/JavaScript/Reference/Objets_globaux/Array x)
}