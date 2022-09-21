# Examen de habilidades de la computación

#### 1. La empresa ADLR ha desarrollado una creama a base de extracto de tomate el cual tiene un efecto despigmentante, esta empresa lo ha contratado para desararrollar un sistema el cual encuentre el punto medio entre la pgmentacion de la piel inicial y pigmentacion de la pierl final, por lo que el gerente general de ADLR le dió estos números `["#7a6736", "#c2982d"]`, en donde la posicion 0 es la inicial y la posicion 1 es la final. Encuenrte el punto medio entre estos dos numeros y escriba el algoritmo utilizado.

#### [Ir al códgio del problema 1](./p1/)

---

#### 2. Realice el algoritmo de multiplicacion de dos matrices, escribe solamente el resultado.

$$ 
\left[ \begin{array}{cc} btc & eth \\ trx & usdt  \end{array} \right]  * 
\left[ \begin{array}{cc} usd & hnl \\ gtq & pen  \end{array} \right] 
$$

#### [Ir al códgio del problema 2](./p2/)

---

#### 5. Escriba las instrucciones JavaScript para poder conectarse a la blockchain web3 (METAMASK), conectese y haga una transaccion de 1 ethereum. (Recuerde el objeto del explorador que permite utilizar metamask).

#### [Ir al códgio del problema 5](./p5/)

---

#### 6. Suponga que un cliente quiere *Q1 000* o *25 000 HNL* de bitcoin al precio *BTC/USD - 1/$20 000* al tipo de cambio de doalar (compra: 7,71 y venta 7,90) 'Cuanto bitcoin recibirá el cliente? Escriba la ecucaión que utilizó y el procedimiento.

$$ 1. transactionAmount = \frac{ \frac{localAmount}{fiatExchangeRate} }{1-userCommission} $$

$$ 2. cryptoAmount = \frac{transactionAmount}{cryptoExchangeRate} $$

$$ transactionAmount = \frac{ \frac{25000}{7.90} }{1-0.05} \longrightarrow transactionAmount = 3326,90$$

$$ cryptoAmount = \frac{3326,90}{20000} \longrightarrow  cryptoAmount = 0.166345 BTC$$

---

#### 7. Realice un diagrama de la infraestructura para poder conectarse a la blockchain (Geth, nodo Bitcoin) y realizar una aplicación WEB el cual pueda crear e importar billeteras de cualquier blockchain. Hint: Intranet


---

#### 8. ¿Qué es la metodología Kanban?

Kanban en una metodología que se apoya de un sistema visual para gestionar una tarea o
trabajo a medida que avanza en un proceso, el objetivo de esta metodología es identificar
posibles “cuellos de botella” y solucionarlos para que el trabajo pueda fluir a una velocidad y rendimiento óptimos. Entonces kanban es una metodología muy eficiente para dar seguimiento a muchos procesos de tecnologías de la información, desarrollo de software y conocimiento en general. A continuacion se describe el flujo por el que se rige Kanban:


1. **Visualizar el flujo de trabajo**: Acá se realiza un tablero simple el cual contiene notas adhesivas o tarjetas, cada tarjeta en el tablero representa una tarea, utilizamos el modelo clásico de tablero kanban donde tenemos 3 columnas:
    - Por hacer: esta columna enumera las tareas que aún no se han iniciado.
    - En proceso: consiste en las tareas que están en progreso (en este punto realizamos las pruebas, nada pasa a la columna de “listo” si no está completamente testeado).
    - Listo: consiste en las tareas que ya se completaron.


2. **Limitar los trabajos en proceso**: Los WIP (los trabajos en proceso) son en pocas palabras todas las tareas en que un equipo está trabajando actualmente lo cual limita la capacidad de flujo de trabajo.

3.  **Administrar el flujo de trabajo**: Administrar .y mejorar el flujo es el quid de un sistema Kanban después de haber implementado los dos primeros pasos, resaltamos la diversas etapas y el estado actual de cada una, aca depende de como establecimos el flujo de trabajo y los límites de los trabajos en proceso se podrá observar un flujo suave o un trabajo acomulado.

4. **Hacer explícitas las reglas del proceso**: Como parte de la visualización de un proceso hay que definir y visualizar explícitamente las reglas sobre cómo se hace un trabajo, es como creamos un acuerdo con los integrantes del grupo para que todos comprendamos cómo describir los trabajos en el sistema.

5. **Implementar bucles de retroalimentación** Nos dimos cuenta que los bucles de retroalimentación son una parte integral de cualquier sistema, el método Kanban impulsa una retroalimentación temprana, especialmente si se cometen errores en un trabajo, para nosotros esto es importante para entregar el trabajo correcto en el menor tiempo posible.

---


#### 9. Escriba todos los comandos necesarios para iniciar un servidor Linux, recuerde que se le ha contratado para levantar un servidor con una aplicación web, el código esta en un repositorio GIT, si usted es QA, agregue su estrategia de testing detallada para la aplicación web (manual).

- sudo apt-get update

- sudo apt-get install git-all

- sudo apt-get install  

- sudo apt-get install apache2

- systemctl status apache2

- systemctl start apache2

- sudo mkdir /var/www/coincaex.com

- sudo chmod -R 755 /var/www/coincaex.com

- cd /var/www/coincaex.com

- git clone `repository`

- systemctl restart apache2

- sudo nano /etc/apache2/sites-available/coincaex.com.conf


Agregar lo siguiente dentro del archivo coincaex.com.conf:


    <VirtualHost *:80>
        ServerAdmin admin@coincaex.com

        ServerName coincaex.com

        ServerAlias www.coincaex.com

        DocumentRoot /var/www/coincaex.com/coincaex/Admin

        ErrorLog ${APACHE_LOG_DIR}/error.log

        CustomLog ${APACHE_LOG_DIR}/access.log combined
    </VirtualHost>

- sudo a2ensite coincaex.com.conf


---


#### 10. En el whitepaper de Satoshi Nakamoto explica acerca de la blockchain, escriba un algoritmo el cual implemente la blockchain, solamente el Árbol de Merkle, en donde el parámetro principal será un arreglo con n transacciones bitcoin.



---

#### 11. Explique la variable “nonce” en un bloque de la blockchain, ¿Para qué sirve?

Es algo parecido a un acertijo, ya que los mineros deben encontrar una cadena (nonce) que al ser combinada con una serie de transacciones sea capaz de generar un hash que empiece con 5 ceros en caso del bitcoin. E primer milnero en descifrar dicho acertijo, será el encargado de agregar el bloque generado dentro del blockchain.

---

#### 13. Desarrolle 


$$ \overrightarrow{E} = \int d \overrightarrow{E} ; La ley de Gauss $$ 

---

#### 14. Considere el campo vectorial


$$ F = y\hat{i} - x\hat{j} $$

Esboce el campo en los puntos:
$$(± 1,0),(0,±1),(±1,±1)$$
(Puede dibujarlo en papel si gusta).


---

#### 15. Resuelva la siguiente ecuación diferencial utilizando cualquier método (Algebra, Transformada de Fourier o Transformada de Laplace) (Si no sabe resolverlo, inténtelo y deje su procedimiento)

$$
    y^2 = y(a^2 - y^2)
$$

NO PUEDE RESOLVERSE

---

#### 16. Desarrolle el tema Devops.

La mayoría de los problemas que se procura resolver con DevOps parten de que normalmente hay 
diferencias marcadas entre desarrollo y operaciones. Son equipos muy separados, con diferente lenguaje, culturas, habilidades, objetivos... Una de las mayores diferencias es que los desarrolladores (analistas, programadores, testers y responsables de calidad) buscan el cambio continuamente, para corregir errores y añadir funcionalidad. Los administradores (administradores de sistemas, de bases de datos y de redes) buscan estabilidad. Ven cualquier cambio como un riesgo potencial. Nadie les agradecería la nueva funcionalidad. Pero si les culparan de los problemas de explotación provocados por los cambios.

DevOps es el resultado de aplicar los principios más confiables del dominio de fabricación física y liderazgo en el flujo de valor de TI. DevOps confía sobre los cuerpos de conocimiento de Lean, Theory of Constraints, Toyota Sistema de producción, ingeniería de resiliencia, organizaciones de aprendizaje, cultura de seguridad, factores humanos y muchos otros. Otros contextos valiosos que dibuja DevOps incluyen culturas de gestión de alta confianza, liderazgo de servicio y gestión del cambio organizacional. El resultado es calidad, confiabilidad, estabilidad y seguridad a un costo y esfuerzo cada vez más bajos; y flujo acelerado y confiabilidad en todo el flujo de valor de la tecnología, incluido el producto Gestión, Desarrollo, Control de Calidad, Operaciones de TI e Infosec.

---

#### 17. ¿Qué es el liderazgo?

Es la capacidad de guiar a un grupo de personas para alncanzar un objetivo en comun, dando un buen ejemplo que pueda reflejarse en un buen trabajo. 

---

#### 18. Se le ha contratado para integrar en un sistema una nueva tecnología llamada SumSub, realice el cronograma de trabajo para este proyecto, puede investigar acerca de la gestión de cronogramas (Si es QA desarrolle la estrategía de testing junto con el cronograma)


---

#### 19. Mencione los 14 dominios de SGSI de ISO/IEC 27001 y desarrolle cada uno.

1. **Política de la seguridad**
Para operativizar el dominio se desarrollaron Políticas para la seguridad de la informacion y son revisadas periódicamente.


2. **Aspectos organizativos de la seguridad de la información**
Garantizar que la organización confronte las funciones y responsabilidades para
gestionar la seguridad de la información y que defina las responsabilidades de cada empleado o puesto de trabajo en relación a la Seguridad de la información.


3. **Seguridad ligada al recurso humano**
Solicita que se aseguré los procesos de selección, empleo, promociones, desvinculaciones y otros, de empleados y contratistas.


4. **Gestión de activos**
El objetivo de este punto es la  preservación de los activos de  información como soporte del negocio.


5. **Control de accesos**
Orientado a controlar y monitorear los accesos a los medios de información de acuerdo a las políticas definidas por la organización.


6. **Gestión de las Comunicaciones Cifrado**
Las medidas de control para el uso eficaz de la criptografía para proteger la confidencialidad e integridad de la información según el apetito al riesgo plasmado en los procesos que ameritan cifrado, criptografía u ofuscamiento.


7. **Seguridad física y ambiental**
Establecer medidas de control físicas para proteger adecuadamente los activos de información para evitar incidentes que afecten a la integridad física de la información o interferencias no deseada..


8. **Seguridad en la operación**
Una serie de controles que tienen como objetivo garantizar la seguridad en la operación y específicamente en el tema de alcance previamente definido.


9. **Seguridad en las comunicaciones**
Proteger los intercambios de información realizados a través de redes de comunicaciones,
estableciendo controles adecuados para proteger tanto las comunicaciones externas a la organización como las que viajan a través de las redes de la propia.


10. **Adquisición, Desarrollo y mantenimiento de los sistemas**



11. **Relación con los proveedores**
Seguridad de la información en las relaciones con los proveedores, garantizando la protección de los activos de la organización, sin vedar la productividad, a fin que sean accesibles por los proveedores los registros o información necesaria.


12. **Gestión de incidentes de seguridad**
Garantizar un enfoque consistente y eficaz para la gestión de incidentes de seguridad de la información, incluyendo.


13. **Seguridad de la información en la continuidad del negocio**
El objetivo del Plan de recuperación ante desastres es definir en forma precisa cómo  recuperará su infraestructura y servicios dentro de los plazos establecidos en el caso que ocurra un desastre o un incidente disruptivo.


14. **Cumplimiento**
Evitar los incumplimientos de las obligaciones legales, estatutarias, reglamentarias o contractuales relacionadas con la seguridad de la información y con los requisitos de seguridad.


---

#### 20. ¿Quiénes son los clientes del departamento de IT de coincaex?

R. Upcion

---

#### 21. Mencione los departamentos que existen en Coincaex y sus funciones (explique también IT)

**Marketing**: Creacion de contenido digital y promocion de la empresa.
**IT**: Desarrollo de software dedicado al exchange de cripto y operaciones internas.
**Cumplimiento**: Validar la legitimidad de las transacciones.
**Ventas**: Brindar soporte a los clientes.
**Contabilidad**: LLevar las cuentas de la empresa.
**Operaciones**: Asegurar que cada transaccion cumpla con los requerimientos de la empresa.

---

#### 22. Utilice este espacio para escribir acerca de la empresa, que no le parece y que le parece, que podemos cambiar para que usted se sienta más cómodo y trabaje de la forma más cómoda.

Podrían tenerse reuniones periodicas con todos los departamentos para concer la situacion de la empresa y para buscar objetivos en comun para obtener un resultado favorable a nivel general, hay que recordar que el dialogo es una estrategia para desarrollar el cambio de perspectiva.

Tambien tratar escuchar a cada integrante por departemento, siempre hay personas con muy buenas ideas, mas sin embargo, no siempre se cuenta con el espacio o libertad de expresion.

---