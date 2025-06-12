# Paso a paso para la construcción del juego.
Este documento describe el desarrollo completo del clon de **Space Invaders** implementado en Go con **Ebiten**, desde cero hasta su versión final.

Cada paso incluye explicaciones técnicas, ejemplos de código, y una rama específica (`step-*`) en el repositorio para ver el estado del proyecto en ese momento.

|Paso|Descripcion|Rama|Ver|
|----|-----------|----|---|
|Introducción|Creación del módulo de juego, instalación de la librería y primera ventana | step-01-inicial|[▶️](./1_introduccion.md)|
|Cañón y movimiento|Creación de sprite a partir de arrays 2D y movimiento de este | step-02-*|[▶️](./4_canon_y_movimiento.md)|
|Disparos|Lógica de disparos (1 a la vez), se mueven verticalmente| step-03-*|[▶️](./5_disparos.md)|
|Bunkers|Bloques protectores y colisiones| step-04-*|[▶️](./6_bunker.md)|
|Enemigos|Nave, aliens, colisiones y explosiones| step-05-*|[▶️](./7_enemigos.md)|
|Victoria y derrota|Estados de juego| step-06-*|[▶️](./8_victoria_y_derrota.md)|
|Asets empotrados|Fuentes, sonidos y restructuración| step-07-*|[▶️](./9_mejoras_adicionales.md)|
|Mejoras|Mejoras en nave y colosiones y reestructuración| step-08-*|[▶️](./10_mejoras.md)|


> 🔔 **Ramas con asterisco**
>
> Las ramas que llevan un asterisco (*) indican que tienen varios sub pasos.


