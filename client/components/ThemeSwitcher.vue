<template>
  <div class="light-switcher" @click="changeTheme">
    <span><font-awesome-icon :icon="['fas', iconToLoad]" /></span>
  </div>
</template>
<script>
export default {
  data() {
    return {
      iconToLoad: "circle"
    }
  },
  beforeMount() {
    document.body.className = this.$store.getters.theme;
    this.iconToLoad = this.$store.getters.theme == "light" ? "moon" : "sun"
  },
  methods: {
    changeTheme() {
      this.$store.dispatch("toggleTheme");
      document.body.className = this.$store.getters.theme;
      this.iconToLoad = this.$store.getters.theme == "light" ? "moon" : "sun";  
    }
  }
};
</script>
<style lang="scss" scoped>
@import "~assets/styles/light.scss";
div {
  * {
    color: inherit;
  }
  position: absolute;
  background: var(--accent-color);
  top: 4.75rem;
  right: 0.75rem;
  padding: 1rem;
  border-radius: 50%;
  box-shadow: 0 0 0.3rem 0;
  text-decoration: none;
  transition: all 0.2s;
  cursor: pointer;

  &:hover {
    filter: grayscale(35%);
  }
}
</style>