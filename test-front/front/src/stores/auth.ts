// src/stores/auth.ts
import { defineStore } from 'pinia';
import { ref, computed } from 'vue';
import type { User } from '../types';
import { loginApi } from '../api/authApi';
import { useRouter } from 'vue-router';

export const useAuthStore = defineStore('auth', () => {
  const router = useRouter();
  const token = ref<string | null>(localStorage.getItem('token'));
  const storedProfile = localStorage.getItem('profile');
  const user = ref<User | null>(storedProfile ? JSON.parse(storedProfile) : null);

  const isAuthenticated = computed(() => !!token.value);
  const isAdmin = computed(() => user.value?.role === 'ADMIN');

  function setToken(newToken: string, profile: User) {
    token.value = newToken;
    user.value = profile
    localStorage.setItem('token', newToken);
    localStorage.setItem('profile', JSON.stringify(profile));
  }

  function logout() {
    token.value = null;
    user.value = null;
    localStorage.removeItem('token');
    localStorage.removeItem('profile');
    router.push('/auth/login');
  }

  async function login(username: string, password: string) {
    try {
      const { token, profile } = await loginApi(username, password);
      setToken(token, profile);
      return { token, profile };
    } catch (error) {
      throw error;
    }
  }
  return {
    token,
    user,
    isAuthenticated,
    isAdmin,
    setToken,
    logout,
    login,
  };
});
