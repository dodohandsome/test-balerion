// src/api/axiosInstance.ts
import axios from 'axios';

export const api = axios.create({
  baseURL: 'http://localhost:3000', // หรือ URL ของ backend จริงๆ
  // สามารถตั้งค่า header หรือ timeout เพิ่มเติมได้ตามต้องการ
});

api.interceptors.request.use(
    (config) => {
      const token = localStorage.getItem('token');
      if (token) {
        config.headers.Authorization = `Bearer ${token}`;
      }
      return config;
    },
    (error) => Promise.reject(error)
  );
