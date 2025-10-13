# Paso 11: Binario y publicación.
En este paso aprenderás:
* Cómo generar binarios multiplataforma (nativo y WebAssembly) desde un proyecto en Go.
* Como aplicar build tags para separar código por plataforma.
* Cómo empaquetar la versión WebAssembly en un .zip adecuado para distribución.
* Cómo subir y publicar tu juego en plataformas como itch.io de forma sencilla.

Ya tenemos el juego como queremos, pero lo estamos ejecutando en un entorno donde tenemos **Golang** instalado.

~~~shell
go run .
~~~

Pero si queremos distribuirlo, tenemos que generar un archivo binario, que como indicamos al inicio de este tutorial, la idea es un binario con los assets autocontenidos sin depender de ningún recurso externo y evitar instalaciones.

## Generación de binarios.
Una de las grandes ventajas que tiene **Golang** es la capacidad multiplataforma, que permite compilar el código para diferentes sistemas operativos y arquitecturas desde un único entorno.

Para llevar a cabo esto, **Go** tiene dos variables de entorno que permiten una compilación cruzada **GOOS** y **GOARCH**.

### La variable GOOS (Go Operating System).
La variable **GOOS** define el sistema operativo sobre el que se genera el binario.

Los valores más usados son:
* **linux**: para la mayoria de los sistemas operativos basados en linux.
* **windows**: para Microsoft Windows.
* **darwin**: para MacOS e iOS.
* **js**: para WebAssembly (este valor va en conjunción con GOARCH=wasm).

### La variable GOARCH (Go Architecture).
La variable **GOARCH** define la arquitectura del procesador en la que va a ser compilado el binario.

Los valores más usados son:
* **amd64**: Para arquitecturas x86-64.
* **arm64**: para arquitecturas ARM de 64 bits.
* **wasm**: para WebAssembly (este valor va en conjunción con GOOS=js).

### Compilación.
El tutorial se ha basado en un entorno linux/debian, por lo que a la hora de crear los binarios es de la siguiente forma, situándonos en el nivel que tenemos el fichero **main.go** hacemos:
* creamos un directorio **bin**:
~~~shell
mkdir bin
~~~

#### Compilar para linux.
Simplemente compilar el proyecto con `build` y lo dejamos en **bin**:
~~~shell
go build -o bin/spaceinvaders main.go
~~~

#### Compilar para MS-Windows.
Compilamos el proyecto con `build` indicando el sistema operativo destino en **GOOS** y la arquitectura **GOARCH** y lo dejamos en **bin**:
~~~shell
GOOS=windows GOARCH=amd64 go build -o bin/spaceinvaders.exe main.go
~~~

#### Compilar para MacOS/iOS.
Al usar **Ebiten**, este usa librerías basadas en **OpenGL** como **GLFW**, y tener que ser compiladas nativamente para poder usar los bindings en C, se requiere de un ordenador con MacOS.

Se puede usar un compilador cruzado (como **osxcross**) y habilitar **Cgo**, pero esto puede ser un proceso complejo ya que se necesita descargar todo el SDK de MacOS y toolchains de Xcode, por lo que lo mas sencillo es usar una máquina MacOS.

#### Compilar para web.
Este proceso requiere de tres ficheros, por lo que creamos un subdirectorio en **bin** llamado **web**:
~~~shell
mkdir -p bin/web
~~~

* Compilar el proyecto para generar el binario **wasm**:
~~~shell
GOOS=js GOARCH=wasm go build -buildvcs=false -o bin/web/spaceinvaders.wasm github.com/programatta/spaceinvaders
~~~

* Creamos un fichero en **bin/web** llamado **index.html** que cargue el fichero **spaceinvaders.wasm** generado:
~~~html
<!DOCTYPE html>
<html>
  <head>
    <meta charset="UTF-8">
    <title>Space Invaders - Ebiten</title>
  </head>
  <body>
    <script src="wasm_exec.js"></script>
    <script>
      const go = new Go();
      WebAssembly.instantiateStreaming(
        fetch("spaceinvaders.wasm"), 
        go.importObject
      ).then(result => { 
        go.run(result.instance);
      });
    </script>
  </body>
</html>
~~~

Se requiere del fichero **wasm_exec.js**, que se encuentra en **GOROOT/lib/wasm/** y lo debemos de copiar en **bin/web**.

De esta forma ya podemos generar binarios distribuibles para **linux**, **windows** y **web**. Para **MacOS** deberemos disponer de una máquina **MacOS**. 

Para facilitarnos las cosas, podemos usar un **Makefile**, al igual que en proyectos de **C** y **C++**. El fichero **Makefile** se encuentra en el raiz del proyecto, y una breve descripción en [**Compilación y ejecución**](../compilation.md).


### El problema de la multiplataforma.
Con los pasos anteriores, ya hemos podido crear binarios para diferentes plataformas, y vemos nuestro juego funcionando en **Linux**, **Windows** y **Navegadores web**. 

Pero, tenemos un problema, y es que disponemos de salida del juego cuando presionamos la **tecla Escape**, tanto en el estado de **Presentation** como en el estado de **Play**. No presenta ningún problema en desktop, es decir, en **Linux** y **Windows**, pero en los navegadores se queda congelado el juego debido a un error.

Para evitar esto, vamos a hacer uso de las **build tags** que ofrece **Go**.

#### Build tags.
Las **build tags** de **Go** funcionan de forma similar a las **directivas de preprocesador** usadas en **C/C++**, pero de una forma más controlada y estructurada.

Son comentarios especiales que se añaden a comienzo del fichero **.go** para indicarle al compilador cuando incluir ese archivo en la compilación.
Se colocan antes de **package** sin comentarios entre ellas.

>🔔 **Nota.**
>
>Los **build tags** se aplican a nivel de archivo, no dentro del código como se puede hacer con las directivas de preprocesador de C/C++ **`#ifdef`**.

##### Cómo se indican en el código.
* Para versiones **1.16** se especifica como **`// +build linux`** (p.e: compila el fichero para sistemas linux).
* Para versiones **1.17 y superiores** se especifica como **`//go:build linux`**.

>🔔 **Nota.**
>
>Se pueden usar ambas para mantener compatibilidad con  versiones antiguas del compilador.

#### ¿Cómo se usan al compilar?
Si fueran personalizadas, usaríamos el parámetro **-tags=<etiqueta>**, si son estandar no es necesario indicar el parámetro **-tags**.

Para solventar el problema detectado, creamos en **internal** un nuevo paquete llamado **platform** que tendrá funcionalidad por plataforma (desktop y browser). Creamos dos ficheros que tengan la función de salida, uno con implementación y otro sin ella.

##### platform/exit_desktop.go
~~~go
//go:build !js && !wasm

package platform

import "os"

func ExitGame() {
  os.Exit(0)
}
~~~

##### platform/exit_wasm.go
~~~go
//go:build js && wasm

package platform

func ExitGame() {
  //Aquí no hacemos nada, ya que no podemos cerrar la pestaña o ventana
  //del navegador.
}
~~~

##### states/presentation/presentation.go
~~~go
func (ps *PresentationState) ProcessEvents() {
  ...
  if ebiten.IsKeyPressed(ebiten.KeyEscape) {
    platform.ExitGame()
  }
}
~~~

##### states/play/play.go
~~~go
func (ps *PlayState) ProcessEvents() {
  if ebiten.IsKeyPressed(ebiten.KeyEscape) {
    platform.ExitGame()
  }
  ...
}
~~~

Con esta funcionalidad solucionamos el problema detectado en la salida del juego en los navegadores. Al hacer uso de GOOS=js GOARCH=wasm **Go** las va a tratar como **build tags** estándar.

Puede consultar el código de este paso en la rama [step-11-binario_publicacion_1](https://github.com/programatta/space-invaders/tree/step-11-binario_publicacion_1).


## Publicación.
Una vez dispongamos de los binarios, podemos distribuirlos de forma sencilla. En el caso de la versión web (WebAssembly), necesitaremos configurar un servidor web o bien usar una plataforma como **itch.io**, que permite publicar juegos web de forma gratuita.

Para ello, debemos subir los tres ficheros que se encuentran en **bin/web**:

Estos ficheros deben estar en la raíz del .zip, sin carpetas intermedias. Para conseguirlo, ejecutamos:

~~~shell
cd bin/web
zip -r ../spaceinvaders.zip .
~~~

Esto generará un archivo **spaceinvaders.zip** en el directorio **bin**, listo para ser subido a **itch.io** como juego HTML5.

Con esto tendríamos en **bin** los binarios nativos:
* **spaceinvades**
* **spaceinvaders.exe** 

y el de distribución web:
* **spaceinvaders.zip**

Con estos pasos, completamos el ciclo de desarrollo y distribución de nuestro clon de **Space Invaders**, facilitando su publicación tanto en formato nativo como en versión web lista para compartir y jugar desde cualquier navegador.
