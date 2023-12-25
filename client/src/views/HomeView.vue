<template>
  <div>
    <div>Home</div>
    <div v-if="is_auth">Welcome, {{ user.user.username }}</div>
  </div>
</template>

<script>
// @ is an alias to /src

export default {
  name: "HomeView",
  computed: {
    token() {
      return this.$store.getters.TOKEN;
    },
    user() {
      return this.$store.getters.USER;
    },
    is_auth() {
      return this.$store.getters.IS_AUTH;
    },
  },
};
</script>
