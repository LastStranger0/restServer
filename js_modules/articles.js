import * as serverApi from './db';

async function all(){
    let response = await serverApi.all();
    return parseResponse(response);
}

async function one(id){
    let response = await serverApi.get(id);
    return parseResponse(response);
}

async function remove(id){
    let response = await serverApi.remove(id);
    return parseResponse(response);
}

export {all, one, remove};

function parseResponse( text ){
    try{
        let info = JSON.parse(text);
        if( info.code !== 200 ){
            throw new Error("Code is't 200");
        }
        return info.data;
    }catch{
        throw new Error("Incorrect form server")
    }
    
}