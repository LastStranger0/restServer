import "@babel/polyfill";
import {server} from "../requests";

function isEmpty(obj) {
    for(var prop in obj) {
        // eslint-disable-next-line no-prototype-builtins
        if(obj.hasOwnProperty(prop)) {
            return false;
        }
    }

    return JSON.stringify(obj) === JSON.stringify({});
}

async function all()
{
    let result = await server.get('products');
    return result.data;
}

async function one(id)
{
    let result = await server.get(`posts/${id}`);
    return result.data;
}

async function remove(id)
{
    let result = await server.post(`deleteProduct/${id}`);
    
    return isEmpty(result.data);
}

async function add( obj )
{
    let result = await server.post('product', obj);
    return obj;
}

async function edit( id, obj )
{
    let result = await server.put(`posts/${id}`, { ...obj, id });
    return result.data;
}

export {all, one, remove, add, edit}