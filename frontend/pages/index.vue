<template>
  <transition name="fade-in">
    <main>
      <login-logout />
      <theme-switcher />
    </main>
  </transition>
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
    if(!this.$store.getters.customThemeSet) this.$store.dispatch("initTheme");
  },
  head() {
    return {
      bodyAttrs: {
        class: this.$store.getters.theme,
      },
    };
  },  
};
</script>