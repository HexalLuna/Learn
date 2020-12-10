/* Importation et setup de LiquidCrystal */
/* Importation et setup de LiquidCrystal */
#include <LiquidCrystal.h>

LiquidCrystal lcd(8, 9, 4, 5 ,6 ,7);

/* Variables utilisées pour le moteur */
const int MotorForward = 10;
const int MotorReverse = 11;
int Power = 80;

/* Variables utilisées pour le calcul de l'intensité et de la tension du courant */
const int RVar = 0;
int mesureBrute = 0;
float mesure;

/* Variables utilisées pour le bouton et chrono */
int  butonPin = A0;
int boutonState = 0;
bool chronoStatus;
unsigned long millisBoutonON = millis();
unsigned long millisBoutonOff = 0;

/* Fonction d'initialisation */
void setup() {
    Serial.begin(115200);
    motorOn();

    pinMode(butonPin, INPUT_PULLUP);
    chronoStatus = true;

    lcd.begin(20, 4);

    lcd.clear(); //efface LCD + se place en 0, 0
    delay(10); //très très courte pause après clear()

    lcd.print("Bienvenue :) !");
    delay(1000); //pause de une secondes

    lcd.clear(); //efface LCD + se place en 0, 0
    delay(10); //très très courte pause après clear()
}

/* Fonction d'allumage moteur */
void motorOn() {
  pinMode(MotorForward, OUTPUT);
  pinMode(MotorReverse, OUTPUT);
  pinMode(2, OUTPUT);
}

/* Fonction de rotation moteur */
void motorForward(int Power) {
  digitalWrite(2, HIGH);
  analogWrite(MotorForward, Power);
  analogWrite(MotorReverse, 0);
}

/* Fonction d'arrêt moteur */
void motorStop() {
  analogWrite(MotorForward, 0);
  analogWrite(MotorReverse, 0);
  digitalWrite(2, LOW); 
}

/* Fonction boucle */
void loop() {

    boutonState = digitalRead(butonPin);

    if (boutonState == LOW && chronoStatus == false) {
        millisBoutonON = millis();
        Serial.println("Chrono ON");
        chronoStatus = true;
        motorForward(Power);

        mesureBrute = analogRead(RVar);
        mesure = map(mesureBrute, 0, 1023, 0.0, 5000.0);
        float volts = mesure * 1000; //conversion millivolts en volts
        float intensity = volts / 0.3; //calcul de l'intensite par la formule I= U/R

        for(float i=0; i<=100; i= intensity * volts * 0.3) {
          int Rendement = 1.5 / i;

            lcd.setCursor(0, 0);
            lcd.print("I: ");
            lcd.print(intensity, 2);
            lcd.print(" mA "); //display intensity on the screen

            lcd.setCursor(0, 1);
            lcd.print("E: ");
            lcd.print(i);
            lcd.print(" J ");
            
            lcd.setCursor(0, 2);
            lcd.print("U: ");
            lcd.print(volts, 2);
            lcd.print(" V "); //display voltage on the screen
            
            lcd.setCursor(0, 3);
            lcd.print("T: ");
            lcd.print(int(millisBoutonOff));
            lcd.print(" s ");

            lcd.setCursor(0, 4);
            lcd.print("n: ");
            lcd.print(Rendement);
        }
    }
        
    if (boutonState == HIGH && chronoStatus == true) {
        millisBoutonOff = millis() - millisBoutonON;
        Serial.println("Chrono OFF = " + String(millisBoutonOff));
        chronoStatus = false;
        motorStop();
    }
}
