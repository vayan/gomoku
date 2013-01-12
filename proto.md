
# IA/CLIENT -> SERVER

* `CONNECT CLIENT` 1er commande envoye par l'IA
* `PLAY X Y` Demande a poser en X Y


# SERVER ->  IA/CLIENT

* `RULES DOUBLE_3 BREAKING_5 TIMEOUT` Confirme la connexion en indiquant les règles optionnelles actives et le temps maximal de réflexion accordé à l'IA par le game server.

* `REM X1 Y1 X2 Y2` Indique une prise, et les coordonnées des 
pierres concernées. La première pierre prise est en (X1 ; Y1), la seconde en (X2 ; Y2).

* `ADD X Y` Indique la pose d'une pierre en (X ; Y).

* `WIN WIN_STATE` Indique la victoire de l'IA et sa raison 

* `LOSE LOSE_STATE` Indique la défaite de l'IA et sa raison 

* `YOURTURN` Indique le début du tour de l'IA, et lui donne la permission de jouer 

# VALEUR POSSIBLE POUR `RULES`


# VALEUR PO