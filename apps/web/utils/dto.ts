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