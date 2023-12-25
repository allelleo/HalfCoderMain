<template>
    <header class="site-header">
      <div class="site-identity">
        <h1><a href="#">HalfCoder v 0.0.1</a></h1>
      </div>
      <nav class="site-navigation">
        <ul v-if="is_auth" class="nav">
          <li><a href="#" @click="Logout">LogOut</a></li>
          <li><a href="#">Profile</a></li>
          <li><router-link to="/">Home</router-link></li>
        </ul>
        <ul v-else class="nav">
          <li><router-link to="/auth/sign-up">Sign Up</router-link></li>
          <li><a href="/auth/sign-in">Sign In</a></li>
          <li><router-link to="/">Home</router-link></li>
        </ul>
      </nav>
    </header>
     
</template>

<style scoped>
a {
  text-decoration: none;
  color: #000;
}
.site-header { 
  border-bottom: 1px solid #ccc;
  padding: .5em 1em;
  display: flex;
  justify-content: space-between;
}

.site-identity h1 {
  font-size: 1.5em;
  margin: .6em 0;
  display: inline-block;
}


.site-navigation ul, 
.site-navigation li {
  margin: 0; 
  padding: 0;
}

.site-navigation li {
  display: inline-block;
  margin: 1.4em 1em 1em 1em;
}
</style>

<script>
export default {
    
    computed : {
        token() {
            return this.$store.getters.TOKEN;
        },
        user() {
            return this.$store.getters.USER;
        },
        is_auth(){
            return this.$store.getters.IS_AUTH;
        }
    },
    methods:{
        Logout(){
            this.$store.dispatch("Logout");
        }
    }

}
</script>