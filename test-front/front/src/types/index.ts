export interface User {
  name: string;
  email: string;
  role: string;
}

export interface Card {
  id: string;
  content: string;
  createdAt: string;
  createdBy: string;
  type: 'ADMIN' | 'USER';
}

export interface LoginResponse {
  token: string;
}