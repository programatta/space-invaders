# Paso 11: Binario y publicación.
En este paso aprenderás:
* Cómo generar binarios multiplataforma (nativo y WebAssembly) desde un proyecto en Go.
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
* **darwing**: para MacOS e iOS.
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
Al usar **Ebiten** librerías basadas en **OpenGL** como **GLFW**, y tener que ser compiladas nativamente para poder usar los bindings en C, se requiere de un ordenador con MacOS.

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

Se requiere del fichero **wasm_exec.js** que se encuentra en **GOROOT/lib/wasm/** y lo debemos de copiar en **bin/web**.

De esta forma ya podemos generar binarios distribuibles para **linux**, **windows** y **web**. Para **MacOS** deberemos disponer de una máquina **MacOS**. 

Para facilitarnos las cosas, podemos usar un **Makefile**, al igual que en proyectos de **C** y **C++**. El fichero **Makefile** se encuentra en el raiz del proyecto, y una breve descripción en el Apendice II. Compilación y ejecución.

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
