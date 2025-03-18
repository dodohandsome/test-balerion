import { api } from './axiosInstance';

interface LoginResponse {
    token: string;
    profile: { name: string; email: string; role: string; };
}

export async function loginApi(username: string, password: string) {
    const { data } = await api.post<LoginResponse>('/login', {
        username,
        password,
    });
    return data;
}
