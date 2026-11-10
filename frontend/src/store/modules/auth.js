import AuthService from "@/services/auth.service";

const state = {
  user: null,
  isAuthenticated: false,
  loading: false
};

const getters = {
  currentUser: state => state.user,
  isAuthenticated: state => state.isAuthenticated,
  isLoading: state => state.loading
};

const mutations = {

  SET_USER(state, user) {
    state.user = user;
    state.isAuthenticated = !!user;
  },

  SET_LOADING(state, value) {
    state.loading = value;
  },

  CLEAR_AUTH(state) {
    state.user = null;
    state.isAuthenticated = false;
  }

};

const actions = {

  async fetchUser({ commit }) {
    commit('SET_LOADING', true);
    try {
      const user = await AuthService.getCurrentUser();

      commit('SET_USER', user);

      return user;
    } catch (error) {
      commit('CLEAR_AUTH');
      throw error;
    } finally {
      commit('SET_LOADING', false);
    }
  },
};

export default {
  namespaced: true,
  state,
  getters,
  mutations,
  actions
};