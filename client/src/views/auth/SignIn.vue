<template>
  <div>
    <input
      v-bind:value="email"
      @input="email = $event.target.value"
      type="text"
      placeholder="Email"
    />
    <br />
    <input
      v-bind:value="password"
      @input="password = $event.target.value"
      type="text"
      placeholder="password"
    />
    <br />
    <button type="text" @click="SignIn">submit</button>
  </div>
</template>

<script>
export default {
  name: "SignIn",
  data() {
    return {
      email: "",
      password: "",
    };
  },
  methods: {
    SignIn() {
      console.log("sign-in");
      fetch("http://127.0.0.1:7000/api/v1/auth/sign-in", {
        method: "POST",
        headers: {
          "Content-Type": "application/json",
        },
        body: JSON.stringify({
          email: this.email,
          password: this.password,
        }),
      }).then((response) => {
        if (response.status != 200) {
          console.log("error");
        } else {
          response.json().then((res) => {
            console.log(res.token);
            this.$store.dispatch("setToken", res.token);
          });
        }
      }).then(setTimeout(() => {this.$store.dispatch("LoadData")}, 600)).then(this.$router.push("/"));
    },
  },
};
</script>
