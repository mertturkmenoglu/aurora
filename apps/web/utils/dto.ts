export type Dto<T> = {
    data: T
}

export type LoginResponseDto = Dto<{
    accessToken: string
    refreshToken: string
}>
