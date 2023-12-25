<template>
  <Header />
  <router-view />
</template>

<script>
import Header from "@/components/Header.vue";
export default {
  components: { Header },
  mounted() {
    this.$store.dispatch("checkToken");
    if (this.$store.getters.IS_AUTH) {
      this.$store.dispatch("LoadData")
    }
  },
};
</script>
