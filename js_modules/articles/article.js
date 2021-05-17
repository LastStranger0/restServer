import {server} from "../requests";
import "@babel/polyfill";

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
    let result = await server.get('posts');
    return result.data;
}

async function one(id)
{
    let result = await server.get(`posts/${id}`);
    return result.data;
}

async function remove(id)
{
    let result = await server.delete(`posts/${id}`);
    return isEmpty(result.data);
}

async function add( obj )
{
    let result = await server.post('posts', obj);
    return result.data;
}

async function edit( id, obj )
{
    let result = await server.put(`posts/${id}`, { ...obj, id });
    return result.data;
}

export {all, one, remove, add, edit};