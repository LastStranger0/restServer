import axios from "axios";

function makeRequest( route, args = {} )
{
    return fetch(`https://restserver22.herokuapp.com/${route}`, args).then( (response) => {
        
        if( response.status !== 200 && response.status !==201 ){
            throw new Error("Code isn't 200");
        }
        try {
            var result = response.json();
        } catch {
            throw new Error("Json wrong");
        }

        return result;

    } ).then(( data ) => {
        return data;
    })

}

let server = axios.create({
    baseURL: "https://restserver22.herokuapp.com/"
})

server.interceptors.request.use( request => {
    
    if( request.method == 'post' || request.method == 'put'){
        request.headers = {
          "Content-type": "application/json; charset=UTF-8"
        };
    }

    return request;
} )

export {makeRequest, server};