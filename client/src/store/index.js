import { createStore } from "vuex";
import VueCookies from "vue-cookies";
export default createStore({
  state: {
    token: null,
    is_auth: null,
    user: {
      profile: {
        birthday: null,
        country: null,
        education: null,
        hobby: null,
        id: null,
        phone: null,
        quote: null,
        sex: null,
        website: null,
        work: null,
      },
      ratings: {
        all: {
          id: null,
          score: null,
        },
        month: {
          id: null,
          score: null,
        },
        week: {
          id: null,
          score: null,
        },
      },
      user: {
        email: null,
        first_name: null,
        id: null,
        last_name: null,
        username: null,
      },
    },
  },
  getters: {
    TOKEN: (state) => {
      return state.token;
    },
    IS_AUTH: (state) => {
      return state.is_auth;
    },
    USER: (state) => {
      return state.user;
    },
  },
  mutations: {
    SET_TOKEN: (state, value) => {
      VueCookies.set("token", value);
      state.token = value;
      state.is_auth = true;
    },
    SET_USER: (state, payload) => {
      state.user = payload;
    },
    LOGOUT: (state) => {
      VueCookies.remove("token");
      state.token = null;
      state.is_auth = false;  
    }
  },
  actions: {
    setToken: ({ commit }, value) => {
      commit("SET_TOKEN", value);
    },
    checkToken: ({ commit }) => {
      let token = VueCookies.get("token");
      if (token) {
        commit("SET_TOKEN", token);
      }
    },
    
    Logout: ( {commit} ) => {
      commit("LOGOUT");
    },
    LoadData: ( {commit} ) => {
      fetch("http://127.0.0.1:7000/api/v1/user/me", {
          method: "POST",
          headers: {
            "Content-Type": "application/json",
          },
          body: JSON.stringify({
            token: VueCookies.get("token"),
          }),
        }).then((response) => {
          if (response.status != 200) {
            console.log("error");
          } else {
            response.json().then((res) => {
              console.log(res);
              commit("SET_USER", res);
            });
          }
        });
    }
  },
  modules: {},
});
