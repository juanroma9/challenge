## Un caso en el que usarías threads para resolver un problema y por qué.?

Los hilos los usuaria para procesos secuenciales en donde se necesite tener el control de las sincronizaciones que se están ejecutando 

## Un caso en el que usarías corrutinas para resolver un problema y por qué.?
Seria para el uso de paralelismo, cuando se evidencia qeu tiene una carga de transaciones y que se puede hacer en paralelo el proceso ahí las utilizaria, de hecho el proyecto los llamados a las apis para obtener los datos complementarios se realiaron en paralelo.

## Análisis de complejidad

de los algorimos presentados considero que el mejor es el logaritmo O(n log n), por que es lineal y no se estaria comiendo los recursos de las maquina o cluster donde se esten ejecutanto si estamos hablando de un gran volumen de datos. 

## Base de datos AlfaDB y BetaDB

La base de datos BetaDB es más estable por que tiene un rendimiento parejo tanto en escritura como lectura al tener un algoritmo de  logaritmo
