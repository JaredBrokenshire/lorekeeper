import store from "@/store";
import api from "@/services/api";

const BASE_URL = "/auth"

const AuthService = {
  // Login
  async login(dto) {
    const response = await api.post(`${BASE_URL}/login`, dto);

    const token = response.data.access_token;

    sessionStorage.setItem('access_token', token);

    return response.data;
  },

  // Register
  async register(dto) {
    const response = await api.post(`${BASE_URL}/register`, dto);

    const token = response.data.access_token;

    sessionStorage.setItem('access_token', token);

    return response.data;
  },

  // Logout
  async logout() {
    const response = await api.post(`${BASE_URL}/logout`);

    sessionStorage.removeItem('access_token');
    store.commit('auth/CLEAR_AUTH');

    return response.data
  },

  // Get authenticated user
  async getCurrentUser() {
    const response = await api.get(`${BASE_URL}/me`);

    sessionStorage.setItem('user', response.data)

    return response.data;
  },

  // Check authentication
  isAuthenticated() {
    return !!sessionStorage.getItem('access_token');
  },

  // Check authentication
  isAdmin() {
    return store.getters["auth/currentUser"].roles.some(role => role.name === "admin");
  },

  // Refresh access token
  async refreshToken() {
    const response = await api.post(`${BASE_URL}/refresh`);

    const token = response.data.access_token;

    if (!token) {
      throw new Error('Refresh response does not contain a token.');
    }

    sessionStorage.setItem('access_token', token);

    return response.data;
  },

  // Get access token
  getToken() {
    return sessionStorage.getItem('access_token');
  }
};

export default AuthService;