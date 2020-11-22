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

        if(piloto == ""){
            response.send("El dorsal introducido no corresponde con ningún piloto")
        }
    
        response.send(piloto)
        return{
            statusCode: 200
        }
    }

    else if (x[0] != "/"){
        response.send("Petición incorrecta, debe introducir como parámetro GET number=x")
    }

    response.send("Bienvenido, introduzca el dorsal del piloto del cual desea ver la información")
    return{
        statusCode: 200
    }
});