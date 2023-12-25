<template>
  <div class="main">
    <div class="form" @submit.prevent>
      <input
        v-bind:value="first_name"
        @input="first_name = $event.target.value"
        type="text"
        placeholder="First name"
      />

      <br />
      <input
        v-bind:value="last_name"
        @input="last_name = $event.target.value"
        type="text"
        placeholder="Last name"
      />

      <br />

      <input
        v-bind:value="email"
        @input="email = $event.target.value"
        type="text"
        placeholder="Email"
      />

      <br />

      <input
        v-bind:value="username"
        @input="username = $event.target.value"
        type="text"
        placeholder="Username"
      />

      <br />

      <input
        v-bind:value="password"
        @input="password = $event.target.value"
        type="text"
        placeholder="password"
      />

      <br />

      <button type="text" @click="SignUp">submit</button>
    </div>
  </div>
</template>

<script>
// @ is an alias to /src

export default {
  name: "SignUp",
  data() {
    return {
      first_name: "",
      last_name: "",
      email: "",
      username: "",
      password: "",
    };
  },
  methods: {
    SignUp() {
      fetch("http://127.0.0.1:7000/api/v1/auth/sign-up", {
        method: "POST",
        headers: { "Content-Type": "application/json;charset=utf-8" },
        body: JSON.stringify({
          firstName: this.first_name,
          lastName: this.last_name,
          username: this.username,
          email: this.email,
          password: this.password,
        }),
      }).then((response) => {
        if (response.status != 200) {
          console.log("error");
        } else {
          this.$router.push("/auth/sign-in");
        }
      });
    },
  },
};
</script>
