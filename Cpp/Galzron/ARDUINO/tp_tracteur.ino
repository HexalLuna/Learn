#include <LiquidCrystal.h>  //Imporation de la librairie "LiquidCrystal" essentiellement programmé en H (fichier header des langages C)

/*
* Une variable dite float est un int décimal, exemple : 1.0 est un float. Un float est donc une valeur numérique décimales (chiffres décimals, nombres décimals)
* Une variable dite int, exemple : 2 est un int. Un int est donc une valeur numérique (chiffres, nombres)
* Une array est un tableau de variables
*/

int Lecture_Tension  = A1; //Déclarations d'une variable de type int
int Lecture_Intensité  = A0; //Déclarations d'une variable de type int
int Te = 30000 ; // Définit la valeur de la période d'échantillonnage Te en ms (valeur à prendre entre 30 000 ms)

int energie[200] = {}; // On définit d'une array appelée "energie" pouvant contenir jusque 201 données et laissé vide pour le moment

const int rs = 2, en = 4, d4 = 9, d5 = 10, d6 = 11, d7 = 12; //Ici nous déclarons des constantes de valeur int
LiquidCrystal lcd(rs, en, d4, d5, d6, d7); //Utilisaion de LiquidCrystal

float Tension = 0.0; //Nous déclarons ici une variable de valeur float
float Intensité = 0.0; //Nous déclarons ici une variable de valeur float

/*
* Initialisation du programme
*/
void setup()  {

    pinMode(13, OUTPUT);
    pinMode(13, OUTPUT);

    lcd.begin(16, 2); //Initialise le lcd 16*2
    Serial.begin(9600);

    lcd.print("---- Groupe 3 ----"); //Message d'intro pour faire jolie
    lcd.setCursor(0, 1);
    lcd.print("---- Arduino ----"); //Message d'intro pour faire jolie

    delay(2000);
    lcd.clear(); //efface les anciennes données de l'écran
}

/*
* Initialisation du programme
*/
void loop() {
    
    digitalWrite(13, HIGH); //demarrage du moteur

    long milisecondes = millis(); //calcul le temps en mililisecondes
    long time = milisecondes / 1000; //convertie les millisecondes en secondes

    while(int i = 0; i < 360000 / Te; i++) { //Boucle d'acquisition de mesures durant 360s
        energie[i] = analogRead(1);
        delay(Te);
    }

    float Tension = analogRead(Lecture_Tension);
    float Intensité = analogRead(Lecture_Intensité);

    Tension = Tension * (5.0 / 1023.0) * 6.46; 
    Intensité = Intensité * (5.0 / 1023.0) * 0.239;
    Energie = energie[i];

    Serial.println(Tension);
    lcd.print(" ");
    Serial.println(Intensité);
    lcd.print(" ");
    Serial.println(Energie);
    
    float Puissance = Tension * Intensité;

    lcd.setCursor(0, 0);

    lcd.print("M="); lcd.print(Rendement); //On affiche le rendement noté M sur le LCD
    lcd.print(" ");
    lcd.print("s="); lcd.print(time); //On affiche le temps du parcours en secondes
    lcd.print(" ");
    lcd.print("Eabs="); lcd.print(Energie); //On affiche l'énergie absorbée noté Eabs sur le LCD

    lcd.setCursor(0, 1);
    delay(1000);
}