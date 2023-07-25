import {ofetch} from "ofetch";
import {User, UserDto} from "~/utils/dto";

export const useAuth = () => {
    const isAuthenticated = useState(() => false)
    const user = ref<User | null>(null)
    const accessToken = useState(() => localStorage.getItem('accessToken'))
    const refreshToken = useState(() => localStorage.getItem('refreshToken'))

    if (accessToken.value && refreshToken.value) {
        isAuthenticated.value = true

        ofetch<UserDto>('http://localhost:5000/api/v1/users/me', {
            headers: {
                'x-access-token': accessToken.value,
                'x-refresh-token': refreshToken.value,
            },
        }).then((response) => {
            user.value = response.data
        }).catch(() => {
            isAuthenticated.value = false
            user.value = null
        })
    }

    return {
        isAuthenticated,
        user,
    }
}