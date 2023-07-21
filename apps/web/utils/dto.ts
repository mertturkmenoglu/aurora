export type Dto<T> = {
    data: T
}

export type LoginResponseDto = Dto<{
    accessToken: string
    refreshToken: string
}>

export type UserDto = Dto<User>

export type User = {
    id: string;
    fullName: string;
    email: string;
    phone: string;
    adPreference: AdPreference;
    addresses: Address[];
};

export type AdPreference = {
    id: string;
    email: boolean;
    sms: boolean;
    phone: boolean;
}

export type Address = {
    id: string;
    city: string;
    description: string;
    isDefault: boolean;
    line1: string;
    line2: string;
    name: string;
    phone: string;
    state: string;
    type: string;
    zipCode: string;
}

export type ProductsDto = Dto<Product[]>
export type ProductDto = Dto<Product>

export type Product = {
    id: string;
    createdAt: string;
    updatedAt: string;
    deletedAt: string;
    name: string;
    description: string;
    currentPrice: number;
    oldPrice: number;
    inventory: number;
    images: ProductImage[];
    isFeatured: boolean;
    isNew: boolean;
    isOnSale: boolean;
    isPopular: boolean;
    shippingPrice: number;
    shippingTime: string;
    shippingType: string;
    slug: string;
    brandId: string;
    brand: Brand;
    categoryId: string;
    category: Category;
}

export type ProductImage = {
    id: string;
    productId: string;
    url: string;
}

export type Brand = {
    id: string;
    createdAt: string;
    updatedAt: string;
    deletedAt: string;
    name: string;
    description: string;
}

export type Category = {
    id: string;
    createdAt: string;
    updatedAt: string;
    deletedAt: string;
    name: string;
    parentId: string | null;
    parent: Category | null;
}