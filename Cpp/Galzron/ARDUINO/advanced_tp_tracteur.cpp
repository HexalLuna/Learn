/*
* Made by Visual Studio Code
* User: Daugé Alexis
* Groupe: G3
* Énoncé => Afficher sur un écran LCD le rendement, tension, intensité, durée de parcours et energie absorbée.
*/


//Importation des libraries ARDUINO
#include <LiquidCrystal.h>
#include <LiquidCrystal_I2C.h>
#include <Stepper.h>

//Déclaration des variables
const int enableBridge = 2;
const int MotorForward = 3;
const int MotorReverse = 5;
const int bouton = 1;
const int rs = 2, en = 4, d4 = 9, d5 = 10, d6 = 11, d7 = 12; 

LiquidCrystal lcd(rs, en, d4, d5, d6, d7);
LiquidCrystal_I2C lcd(0x27, 20, 4);

int Lecture_Tension  = A1; 
int Lecture_Intensité  = A0; 
int etatBouton, compteur, loop, loop2, t;

int Power = 80; //Puissance du moteur

float Tension = 0.0; 
float Intensité = 0.0; 


//Initalisation du programme
void setup() {

  lcd.init();
  pinMode(bouton, INPUT);
  etatBouton = 1;

  dcBegin();

  lcd.begin(16, 2);
  Serial.begin(9600);

  lcd.print("---- Groupe 3 ----");
  lcd.setCursor(0, 1);
  lcd.print("---- Arduino ----"); 

  delay(2000);
  lcd.clear();
}


void dcBegin() { // Initialise les pines utilisées pour le moteur
  pinMode(MotorForward, OUTPUT);
  pinMode(MotorReverse, OUTPUT);
  pinMode(enableBridge, OUTPUT);
}

void dcForward(int Power) {
  digitalWrite(enableBridge, HIGH); // Active pont en H
  analogWrite(MotorForward, Power); // Tourne dans le sens Forward à la vitesse spécifiée par Power
  analogWrite(MotorReverse, 0);
}

void dcStop() { //Arrête le moteur
  analogWrite(MotorForward, 0);
  analogWrite(MotorReverse, 0);
  digitalWrite(enableBridge, LOW); 
}

void loop() {

    lcd.backlight();
    etatBouton = digitalRead(bouton);

    if (etatBouton === 0)  { 
        loop = 1;
        loop2 = 0;
        t = 0;
        Eabs = 0;
        compteur = 0; 

        dcBegin(); //Initialisation du moteur
        dcForward(); //Lancement du moteur

        float Tension = analogRead(Lecture_Tension);
        float Intensité = analogRead(Lecture_Intensité);
   
        Serial.println(Tension);
        lcd.print(" ");
        Serial.println(Intensité);

        while (loop2 === 1) {
          t = t + 0.3;

          Tension = Tension * (5.0 / 1023.0) * 6.46; 
          Intensité = Intensité * (5.0 / 1023.0) * 0.239;

          Energie = Tension * Intensité * t
        }

        while (loop == 1) {

          lcd.clear();

          lcd.setCursor(0, 0);
          compteur++;
          lcd.print("s=", compteur);

          delay(1000);

          etatBouton = digitalRead(bouton);

          if (etatBouton == 0) {
              loop = 0;
              etatBouton = digitalRead(bouton);
              dcStop(); //Arrêt du moteur

              delay(1000);
          }
        }
    }
}

