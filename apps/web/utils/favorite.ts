import {FetchError} from "ofetch";
import {Dto} from "~/utils/dto";

export type Favorite = {
    id: string;
    productId: string;
}

async function addToFavorite(productId: string) {
    try {
        const res = await api<Dto<Favorite>>('/favorites', {
            method: 'POST',
            headers: {
                'x-access-token': localStorage.getItem('accessToken') || '',
                'x-refresh-token': localStorage.getItem('refreshToken') || '',
            },
            body: {
                productId,
            }
        });

        return res.data
    } catch (e) {
        if (e instanceof FetchError) {
            console.error(e.message);
        }

        return null;
    }
}

async function removeFromFavorite(favoriteId: string) {
    try {
        const res = await api<Dto<Favorite>>(`/favorites/${favoriteId}`, {
            method: 'DELETE',
            headers: {
                'x-access-token': localStorage.getItem('accessToken') || '',
                'x-refresh-token': localStorage.getItem('refreshToken') || '',
            },
        });

        return res.data
    } catch (e) {
        if (e instanceof FetchError) {
            console.error(e.message);
        }

        return null;
    }
}

async function getFavorites() {
    try {
        const res = await api<Dto<Favorite[]>>('/favorites', {
            method: 'GET',
            headers: {
                'x-access-token': localStorage.getItem('accessToken') || '',
                'x-refresh-token': localStorage.getItem('refreshToken') || '',
            },
        });

        return res.data
    } catch (e) {
        if (e instanceof FetchError) {
            console.error(e.message);
        }

        return null;
    }
}

export class FavoriteManager {
    private static instance: FavoriteManager;

    private constructor() {
        this.getFavorites().then((res) => {
            if (res) {
                this.writeFavoritesToSessionStorage(res);
            }
        })
    }

    public static getInstance() {
        if (!FavoriteManager.instance) {
            FavoriteManager.instance = new FavoriteManager();
        }

        return FavoriteManager.instance;
    }

    public static async getAsyncInstance() {
        if (!FavoriteManager.instance) {
            FavoriteManager.instance = new FavoriteManager();
            await FavoriteManager.instance.invalidate();
        }

        return FavoriteManager.instance;
    }

    public async invalidate() {
        const res = await this.getFavorites();

        if (res) {
            this.writeFavoritesToSessionStorage(res);
        }
    }

    public async getFavorites() {
        const res = await getFavorites();

        if (res) {
            this.writeFavoritesToSessionStorage(res);
        }

        return res;
    }

    public async addToFavorite(productId: string) {
        const res = await addToFavorite(productId);

        if (res) {
            this.appendSessionStorage(res);
        }
    }

    public async removeFromFavorite(productId: string) {
        const favs = this.readFavoritesFromSessionStorage();

        const fav = favs.find(f => f.productId === productId);

        if (!fav) {
            return;
        }

        const res = await removeFromFavorite(fav.id);

        if (res) {
            this.removeFromSessionStorage(res);
        }
    }

    public isFavorite(productId: string) {
        const favorites = this.readFavoritesFromSessionStorage();

        return favorites.some(f => f.productId === productId);
    }

    private readFavoritesFromSessionStorage(): Favorite[] {
        const serialized = sessionStorage.getItem('favorites');

        if (!serialized) {
            return [];
        }

        try {
            return JSON.parse(serialized);
        } catch (e) {
            return [];
        }
    }

    private writeFavoritesToSessionStorage(favorites: Favorite[]) {
        sessionStorage.setItem('favorites', JSON.stringify(favorites));
    }

    private appendSessionStorage(data: Favorite) {
        const favorites = this.readFavoritesFromSessionStorage();

        favorites.push(data);

        this.writeFavoritesToSessionStorage(favorites);
    }

    private removeFromSessionStorage(data: Favorite) {
        const favorites = this.readFavoritesFromSessionStorage();

        const index = favorites.findIndex(f => f.id === data.id);

        if (index !== -1) {
            favorites.splice(index, 1);
        }

        this.writeFavoritesToSessionStorage(favorites);
    }
}