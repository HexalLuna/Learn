/* Importation des libraries */
#include <LiquidCrystal.h>

/* Variables utilisées pour le moteur */
const int enableBridge = 2;
const int MotorForward = 3;
const int MotorReverse = 5;
int Power = 80;

/* Variables utilisées pour le bouton et chrono */
st int  butonPin = 9;
int boutonState = 0;
bool chronoStatus;
unsigned long millisBoutonON = millis();
unsigned long millisBoutonOff = 0;

/* Calcul et autres */
int EnergieABS = float(5 * 40 * (pow(10, -3)); //5 is current voltage, 40 is current intensity, with this calcul EnergieABS = 0.2
int EnergieUtile;                       

/* Fonction d'initialisation */
void setup() {
  lcd.init();
  motorOn();

  lcd.begin(16, 2);
  Serial.begin(9600);

  inMode(butonPin, INPUT_PULLUP);
  chronoStatus = true;

  lcd.setCursor(0, 1);

  delay(2000);
  lcd.clear();
}

/* Fonction d'allumage moteur */
void motorOn() {
  pinMode(MotorForward, OUTPUT);
  pinMode(MotorReverse, OUTPUT);
  pinMode(enableBridge, OUTPUT);
}

/* Fonction de rotation moteur */
void motorForward(int Power) {
  digitalWrite(enableBridge, HIGH);
  analogWrite(MotorForward, Power);
  analogWrite(MotorReverse, 0);
}

/* Fonction d'arrêt moteur */
void motorStop() {
  analogWrite(MotorForward, 0);
  analogWrite(MotorReverse, 0);
  digitalWrite(enableBridge, LOW); 
}

/* Fonction boucle */
void loop() {
    boutonState = digitalRead(butonPin);

    if (boutonState == LOW && chronoStatus == false) {
        motorOn();
        millisBoutonON = millis();
        Serial.println("Chrono ON");
        chronoStatus = true;
        motorForward();
    }
        
    if (boutonState == HIGH && chronoStatus == true) {
        millisBoutonOff = millis() - millisBoutonON;
        Serial.println("Chrono OFF = " + String(millisBoutonOff));
        chronoStatus = false;
        motorStop();
      
        int coeffRendement = x;//j'ai pas le coeff
        float EnergieUtile = coeffRendement * EnergieABS;
        float Rendement = EnergieUtile / EnergieABS;
        
        lcd.setCursor(0, 0);
        lcd.print("U= 5V"); //display voltage on the screen

        lcd.setCursor(0, 1);
        lcd.print("I= 40mA"); //display current intensity on the screen

        lcd.setCursor(1, 0);
        lcd.print(int(millisBoutonOff), "s"); //display the travel time on the screen

        lcd.setCursor(1, 1);
        lcd.print(EnergieABS); //display the absorbed energy on the screen

        lcd.setCursor(2, 0);
        lcd.print(Rendement);
    }
}
