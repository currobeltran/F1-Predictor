const functions = require('firebase-functions');

exports.piloto = functions.https.onRequest((request, response) => {
    // Creamos el listado de pilotos
    var c = {
        "44" : {
            "Name": "Lewis Hamilton"
        },
        "77" : {
            "Name": "Valteri Bottas"
        },
        "5" : {
            "Name": "Sebastian Vettel"
        },
        "16" : {
            "Name": "Charles Leclerc"
        },
        "33" : {
            "Name": "Max Verstappen"
        },
        "23" : {
            "Name": "Alexander Albon"
        },
        "55" : {
            "Name": "Carlos Sainz"
        },
        "4" : {
            "Name": "Lando Norris"
        },
        "11" : {
            "Name": "Sergio Perez"
        },
        "18" : {
            "Name": "Lance Stroll"
        },
        "7" : {
            "Name": "Kimi Raikkonen"
        },
        "99" : {
            "Name": "Antonio Giovinazzi"
        },
        "3" : {
            "Name": "Daniel Ricciardo"
        },
        "31" : {
            "Name": "Esteban Ocon"
        },
        "63" : {
            "Name": "George Russell"
        },
        "6" : {
            "Name": "Nicholas Latifi"
        },
        "20" : {
            "Name": "Kevin Magnussen"
        },
        "8" : {
            "Name": "Romain Grosjean"
        },
        "10" : {
            "Name": "Pierre Gasly"
        },
        "26" : {
            "Name": "Daniel Kyvat"
        }
    };

    //El usuario introduce el número del piloto y lo sacará como resultado
    //Si no existe, se le informará de dicho error

    var e = request.url.split("&");
    var x = e[0].split("=")

    if (x[0].includes("number")){
        var piloto = c[x[1]]
        console.log(piloto)
        if(!piloto){
            response.status(200).sendFile('/static/notexist.html', { root : __dirname })
        }
        else{
            response.status(200).type('html').send(
                `<!DOCTYPE html>
                <html lang="es">
                    <head>
                        <title>Consulta Pilotos</title>
                        <meta charset="utf-8">
                    </head>
                
                    <body>
                        <h1>Datos del piloto</h1>
                        <ul>
                            <li>Nombre: ${piloto.Name}</li>
                        </ul>
                    </body>
                </html>`
            )
        }
    }

    else if (x[0] != "/"){
        response.status(400).sendFile('/static/incorrecto.html', { root : __dirname })
    }

    response.status(200).sendFile('/static/index.html', { root : __dirname })
});