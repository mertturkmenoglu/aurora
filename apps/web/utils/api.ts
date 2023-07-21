import {ofetch} from 'ofetch';

export const BASE_URL = 'http://localhost:5000/api/v1';

export const api = ofetch.create({
    baseURL: BASE_URL,
});
