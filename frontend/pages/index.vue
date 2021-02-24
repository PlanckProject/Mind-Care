<template>
  <main>
    <site-header />
    <theme-switcher />
    <provider-carousel />
  </main>
</template>
<script>
export default {
  transition: "fade",
  data() {
    return {
      currentHour: new Date().getHours(),
    };
  },
  beforeMount() {
    console.log(this.$store.getters["config/apiServerUri"]);
    if (!this.$store.getters["theme/customThemeSet"])
      this.$store.dispatch("theme/initTheme");
  },
  head() {
    return {
      bodyAttrs: {
        class: this.$store.getters["theme/theme"],
      },
    };
  },
};
</script>
<style lang="scss">
body {
  margin: 0.5rem 0;
}
main {
  margin: auto;
  background: red;
  width: 95vw;
}
</style>