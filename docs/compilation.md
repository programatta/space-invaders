# 🛠️ Compilación y ejecución

Este proyecto incluye un **Makefile** con tareas que facilitan la compilación y ejecución en distintos entornos:
* Linux
* Windows
* Web

Se genera un ejecutable autocontenible con todos los recursos embebidos bajo el directorio **bin**.

## 🔨 Compilación.
### 🐧Compilar para tu linux.
Para generar un binario para linux:
~~~bash
make build
~~~
Esto generará el binario en **bin/spaceinvaders**.

### 🪟 Compilar para Windows.
Para generar un binario para MS Windows 64 bits:
~~~bash
make build-win
~~~
Se creará el ejecutable **bin/spaceinvaders.exe**.

### 🌍 Compilar para WebAssembly (WASM).
Para generar un bundle para web:
~~~bash
make build-web
~~~
Esto genera tres ficheros:
* **bin/web/spaceinvaders.wasm**
* **bin/web/wasm_exec.js**
* **bin/web/index.html**

Estos tres ficheros se pueden publicar en **itch.io** si se comprimen juntos en un **zip** y se sube a dicha plataforma.


### 🧼 Limpiar artefactos.
Para limpiar los binarios generados:
~~~bash
make clean
~~~

### 🎯 Generar todos los binarios.
Para generar todos los binarios y web:
~~~bash
make build-all
~~~
La salida como se indica en los anteriores apartados.

## ▶️ Ejecución.
Se puede ejecutar el proyecto sin generar un binario o bundle web:
### 🖥️ Desktop.
~~~bash
make run
~~~

### 🌍 Web.
Levanta un servidor en el puerto 8080 del contendor.
~~~bash
make run-web
~~~
