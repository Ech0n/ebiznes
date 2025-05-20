import axios from 'axios';

const API_BASE_URL = 'http://localhost:8080/auth';

export const login = async (email: string, password: string) => {
  try {
    const response = await axios.post(`${API_BASE_URL}/login`, { email, password });
    return response.data;
  } catch (error) {
    console.error('Login failed:', error);
    throw error;
  }
};

export const register = async (email: string, password: string) => {
  try {
    const response = await axios.post(`${API_BASE_URL}/register`, { email, password });
    return response.data;
  } catch (error) {
    console.error('Registration failed:', error);
    throw error;
  }
};

export const logout = async () => {
  try {
    const response = await axios.delete(`${API_BASE_URL}/logout`);
    return response.data;
  } catch (error) {
    console.error('Logout failed:', error);
    throw error;
  }
};
