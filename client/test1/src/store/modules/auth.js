import axios from "axios";

const state = {
  user: null,
  posts: null,
};

const getters = {
  isAuthenticated: (state) => !!state.user,
  StatePosts: (state) => state.posts,
  StateUser: (state) => state.user,
};
const error = (res) => {
  alert("err")
}
const actions = {
  Register({dispatch}, form) {
    let data = {
      "firstName": "John",
      "lastName": "Doe",
      "username": "all1234e1lleo",
      "email": "1134125@example.com",
      "password": "123"
    }
    const url = new URL("http://127.0.0.1:7000/api/v1/auth/sign-up")
   
    fetch(url.toString(), {
      method: "POST",
      headers: {'Content-Type': 'application/json;charset=utf-8'},
      body: JSON.stringify(data)
    }).catch(
      
    ).then((response) => response.json()).then((res) => console.log(res))
  },

  async LogIn({commit}, user) {
    await axios.post("login", user);
    await commit("setUser", user.get("username"));
  },

  async CreatePost({ dispatch }, post) {
    await axios.post("post", post);
    return await dispatch("GetPosts");
  },

  async GetPosts({ commit }) {
    let response = await axios.get("posts");
    commit("setPosts", response.data);
  },

  async LogOut({ commit }) {
    let user = null;
    commit("logout", user);
  },
};

const mutations = {
  setUser(state, username) {
    state.user = username;
  },

  setPosts(state, posts) {
    state.posts = posts;
  },
  logout(state, user) {
    state.user = user;
  },
};

export default {
  state,
  getters,
  actions,
  mutations,
};
