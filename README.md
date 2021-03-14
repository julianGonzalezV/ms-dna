# ms-dna
Microservicio para la validacion de DNA mutante, requerido por Magneto para el proceso de seleccion de sus trabajadores y 
combatir a los X-man :) :O 

# Servicios Expuestos:
1) POST isMutant: Recibe un arreglo de string que representan el ADN y devuelve 201 si es un mutante, de lo contrario 403
Ejemplo de consumo:
> URL: https://s2h06thmnd.execute-api.us-east-1.amazonaws.com/test/mutant
> request BODY: {"dna":["ATGCGA","CAGTGC","TTATGT","AGAAGG","CCCCTA","TCACTG"]}
> Response: 200 OK

2) GET Stats: Calcula las estadisticas de los ADNs evaluados, indicando la cantidad de ADNs mutantes, la cantidad de ADNs no mutantes (Humanos) y la correlación entre ambos (Ratio)
Ejemplo de consumo:
> URL https://s2h06thmnd.execute-api.us-east-1.amazonaws.com/test/stats
> Response: 200-OK y {"count_mutant_dna": 2,"count_human_dna": 1,"ratio": 2}


# Detalles de la solucion
> Uso de variables de ambiente para la configuración
> Conexión con base de datos MONGO, proveedor Mongo Atlas que nos ofrece clusterización para mayor disponibilidad
> Se hospeda en AWS Lambda con el fin de soportar Auto-Escalamiento dado el reqierimiento de posibles picos de consumos 
esto nos ayuda a ahorrar costos de mantenimiento de Infraestructura pero a la vez tiene un downside costos cuando la cantidad
de requests empieza a superar el umbral de los 90millones de consumos al mes ..en este caso es mejor explorar opciones de 
CONTENERIZACIÓN y hacer uno de algún servicoi nube de KUBERNETES para la gestión del balanceo, enrrutamiento 
Para el diseño se hace uso de Domian Driven Design 
> Capa de  Aplicacion 
> capa de dominio 
> Capa de Infraestructura
La comunicación entre las capas se hace de manera abstracta, a fin de no depender de la implementación


# Covertura de pruebas 
Se corren sobre el paquete de servicios de dominio (domain/service), se supone que las pruebas deberían concentrarsen allí, 
en los demás paquetes se podrían aplicar otros tipos de  pruebas como por ejemplo las de integración

> $ go test -cover
Respuesta 
> --- PASS: TestSaveDna (0.00s)

PASS
coverage: 83.2% of statements
exit status 1

# Notas a considerar
Si la disponibilidad se ve afaectada dada la concurrencia y tenerlo en lambdas AWS no la soporta entonces pensar en un servicio de
KUBERNETES 


# Instalación Local 
In your local host you have to create the next folders structrure:
XXworkspace
    bin
    src
    pkg

>Then, create the GOPATH environment variable, point to xxworkspace folder--ok

>Additionally, you  have to create  GOPATH/bin into your PATH env variable (This is a pending step, please read https://golang.org/doc/gopath_code.html before )
    

# Commands to execute for local conf
To install gorilla mux
$ go get -u github.com/gorilla/mux

# Install the MongoDB Go Driver
https://blog.friendsofgo.tech/posts/driver-oficial-mongodb-golang/
go get -u go.mongodb.org/mongo-driver

## RUN on local Linux
```bash
$ . start_dev_linux.sh
```

## RUN on local Windows
```bash
$ start.bat
```

# Get zip to AWS lambda: 
Process
- Into your root project folder execute:
> $ go get github.com/aws/aws-lambda-go/lambda   (if required)
> $ GOOS=linux go build main.go --> this command create an executable file called main as the .go name file
> $ zip ms-dna.zip main
Upload zip to S3 via aws cli or manually

