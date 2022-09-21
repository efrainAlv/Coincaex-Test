# Examen de habilidades de la computación

##### 1. La empresa ADLR ha desarrollado una creama a base de extracto de tomate el cual tiene un efecto despigmentante, esta empresa lo ha contratado para desararrollar un sistema el cual encuentre el punto medio entre la pgmentacion de la piel inicial y pigmentacion de la pierl final, por lo que el gerente general de ADLR le dió estos números `["#7a6736", "#c2982d"]`, en donde la posicion 0 es la inicial y la posicion 1 es la final. Encuenrte el punto medio entre estos dos numeros y escriba el algoritmo utilizado.

#### [Ir al códgio del problema 1](./p1/)

---

##### 2. Realice el algoritmo de multiplicacion de dos matrices, escribe solamente el resultado.

$$ 
\left[ \begin{array}{cc} btc & eth \\ trx & usdt  \end{array} \right]  * 
\left[ \begin{array}{cc} usd & hnl \\ gtq & pen  \end{array} \right] 
$$

#### [Ir al códgio del problema 2](./p2/)

---

##### 5. Escriba las instrucciones JavaScript para poder conectarse a la blockchain web3 (METAMASK), conectese y haga una transaccion de 1 ethereum. (Recuerde el objeto del explorador que permite utilizar metamask).

#### [Ir al códgio del problema 5](./p5/)

---

##### 6. Suponga que un cliente quiere *Q1 000* o *25 000 HNL* de bitcoin al precio *BTC/USD - 1/$20 000* al tipo de cambio de doalar (compra: 7,71 y venta 7,90) 'Cuanto bitcoin recibirá el cliente? Escriba la ecucaión que utilizó y el procedimiento.

$$ 1. transactionAmount = \frac{ \frac{localAmount}{fiatExchangeRate} }{1-userCommission} $$

$$ 2. cryptoAmount = \frac{transactionAmount}{cryptoExchangeRate} $$

$$ transactionAmount = \frac{ \frac{25000}{7.90} }{1-0.05} \longrightarrow transactionAmount = 3326,90$$

$$ cryptoAmount = \frac{3326,90}{20000} \longrightarrow  cryptoAmount = 0.166345 BTC$$

---

##### 7. Realice un diagrama de la infraestructura para poder conectarse a la blockchain (Geth, nodo Bitcoin) y realizar una aplicación WEB el cual pueda crear e importar billeteras de cualquier blockchain. Hint: Intranet


---

##### 8. ¿Qué es la metodología Kanban?

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


##### 9. Escriba todos los comandos necesarios para iniciar un servidor Linux, recuerde que se le ha contratado para levantar un servidor con una aplicación web, el código esta en un repositorio GIT, si usted es QA, agregue su estrategia de testing detallada para la aplicación web (manual).

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


##### 10. En el whitepaper de Satoshi Nakamoto explica acerca de la blockchain, escriba un algoritmo el cual implemente la blockchain, solamente el Árbol de Merkle, en donde el parámetro principal será un arreglo con n transacciones bitcoin.



---

##### 11. Explique la variable “nonce” en un bloque de la blockchain, ¿Para qué sirve?



---

##### 13. Desarrolle 


$$ \overrightarrow{E} = \int d \overrightarrow{E} ; La ley de Gauss $$ 


---


##### 14. Considere el campo vectorial


$$ F = y\hat{i} - x\hat{j} $$

Esboce el campo en los puntos:
$$(± 1,0),(0,±1),(±1,±1)$$
(Puede dibujarlo en papel si gusta).


---

##### 15. Resuelva la siguiente ecuación diferencial utilizando cualquier método (Algebra, Transformada de Fourier o Transformada de Laplace) (Si no sabe resolverlo, inténtelo y deje su procedimiento)

$$
    y^2 = y(a^2 - y^2)
$$

NO PUEDE RESOLVERSE

---

##### 16. Desarrolle el tema Devops.

La mayoría de los problemas que se procura resolver con DevOps parten de que normalmente hay 
diferencias marcadas entre desarrollo y operaciones. Son equipos muy separados, con diferente lenguaje, culturas, habilidades, objetivos... Una de las mayores diferencias es que los desarrolladores (analistas, programadores, testers y responsables de calidad) buscan el cambio continuamente, para corregir errores y añadir funcionalidad. Los administradores (administradores de sistemas, de bases de datos y de redes) buscan estabilidad. Ven cualquier cambio como un riesgo potencial. Nadie les agradecería la nueva funcionalidad. Pero si les culparan de los problemas de explotación provocados por los cambios.

DevOps es el resultado de aplicar los principios más confiables del dominio de fabricación física y liderazgo en el flujo de valor de TI. DevOps confía sobre los cuerpos de conocimiento de Lean, Theory of Constraints, Toyota Sistema de producción, ingeniería de resiliencia, organizaciones de aprendizaje, cultura de seguridad, factores humanos y muchos otros. Otros contextos valiosos que dibuja DevOps incluyen culturas de gestión de alta confianza, liderazgo de servicio y gestión del cambio organizacional. El resultado es calidad, confiabilidad, estabilidad y seguridad a un costo y esfuerzo cada vez más bajos; y flujo acelerado y confiabilidad en todo el flujo de valor de la tecnología, incluido el producto Gestión, Desarrollo, Control de Calidad, Operaciones de TI e Infosec.

---

##### 17. ¿Qué es el liderazgo?



---

##### 18. Se le ha contratado para integrar en un sistema una nueva tecnología llamada SumSub, realice el cronograma de trabajo para este proyecto, puede investigar acerca de la gestión de cronogramas (Si es QA desarrolle la estrategía de testing junto con el cronograma)


---
##### 19. Mencione los 14 dominios de SGSI de ISO/IEC 27001 y desarrolle cada uno.


---
##### 20. ¿Quiénes son los clientes del departamento de IT de coincaex?


---
##### 21. Mencione los departamentos que existen en Coincaex y sus funciones (explique también IT)


---
##### 22. Utilice este espacio para escribir acerca de la empresa, que no le parece y que le parece, que podemos cambiar para que usted se sienta más cómodo y trabaje de la forma más cómoda.



---