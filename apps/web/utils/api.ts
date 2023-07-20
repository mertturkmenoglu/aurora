import {ofetch} from 'ofetch';

export const api = ofetch.create({
    baseURL: 'http://localhost:5000/api/v1',
});
