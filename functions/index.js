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
        
        if(!piloto){
            response.status(200).type('html').send(
                `<!DOCTYPE html>
                <html lang="es">
                    <head>
                        <title>Consulta Pilotos</title>
                        <meta charset="utf-8">
                    </head>
                
                    <body>
                        <h1>No existe ningún piloto con este número</h1>
                    </body>
                </html>`
            )
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

    else if (x[0] !== "/"){
        response.status(400).type('html').send(
            `<!DOCTYPE html>
            <html lang="es">
                <head>
                    <title>Consulta Pilotos</title>
                    <meta charset="utf-8">
                </head>
            
                <body>
                    <h1>Error al pasar los parámetros a la función</h1>
                </body>
            </html>`
        )
    }

    response.status(200).type('html').send(
        `<!DOCTYPE html>
        <html lang="es">
            <head>
                <title>Consulta Pilotos</title>
                <meta charset="utf-8">
            </head>
        
            <body>
                <form method="GET">
                    <label>
                        Dorsal del piloto
                        <input type="number" name="number" />
                    </label>
                    <button type="submit">Enviar</button>
                </form>
            </body>
        </html>`
    )
});