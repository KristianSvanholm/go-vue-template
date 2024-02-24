import axios from 'axios';

let api;

function startApi(){
    const host = process.env.NODE_ENV === 'development' ? "localhost:8080" : window.location.host;
    api = axios.create({
        withCredentials: true,
        baseURL: `https://${host}/api/v1`
    });

    api.interceptors.response.use((res) => {
        return res.data;
    });
}

startApi();


// Authentication
export const postExample = (payload) => {return api.post('/somepath/', payload)}
export const getExample = (params) => {return api.get('/somepath/', {params})} // Remember the curly braces for URL param-unfolding
export const putExample = (payload)  => {return api.put('/somepath/', payload)}
export const deleteExample = (name,id) => {return api.delete(`/somepath/${name}/${id}`)}

