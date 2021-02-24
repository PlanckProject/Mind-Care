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
    document.body.className = this.$store.getters["theme/theme"];
    this.iconToLoad = this.$store.getters["theme/theme"] == "light" ? "moon" : "sun"
  },
  methods: {
    changeTheme() {
      this.$store.dispatch("theme/toggleTheme");
      document.body.className = this.$store.getters["theme/theme"];
      this.iconToLoad = this.$store.getters["theme/theme"] == "light" ? "moon" : "sun";  
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
  width: 3rem;
  height: 3rem;
  position: absolute;
  background: var(--accent-color);
  top: 1rem;
  right: 1rem;
  padding: 0.75rem 1rem;
  border-radius: 50%;
  box-shadow: 0 0 0.3rem 0;
  text-decoration: none;
  transition: all 0.2s;
  cursor: pointer;
  font-size: 1rem;

  span {
    margin: 0;
  }
  &:hover {
    filter: grayscale(35%);
  }
}
</style>