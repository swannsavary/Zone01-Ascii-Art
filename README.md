Ce projet consiste à prendre une chaîne de caractères et à l'afficher en fonction de différents fichiers/polices d'écriture.

Ce projet ne prend en compte que les caractères ASCII de 32 jusqu'à 128. De plus, les « \n » ne sont pas écrits, mais font sauter une ligne.

option:

veiller a  
[outpup][color][word colored][text][police]

output:
renvoie le resultat dans un fichier

--output<nomdufichier.txt>

color:
affiche le resultat en couleur

--color=<couleur>
couleur : [red,green,yellow,blue,magenta,cyan,white,black]

word colored :
apres un --color ecrire le text que vous voulais en couleur specifiquement

--color=<color> word color    text


police:
standard
shadow
thinkertoy
