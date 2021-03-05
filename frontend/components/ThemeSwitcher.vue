<template>
  <div class="light-switcher" @click="changeTheme">
    <span><font-awesome-icon :icon="iconToLoad" /></span>
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
  line-height: 2.25rem;
  width: 2.25rem;
  height: 2.25rem;
  position: absolute;
  background: var(--accent-color);
  top: 0.5rem;
  right: 0.5rem;
  text-align: center;
  border-radius: 50%;
  box-shadow: 0 0 0.15rem 0;
  text-decoration: none;
  transition: all 0.2s;
  cursor: pointer;
  font-size: 0.75rem;

  span {
    margin: 0;
  }
  &:hover {
    filter: grayscale(35%);
  }
}

@media screen and (min-width: 550px) {
  div {
  * {
    color: inherit;
  }
  line-height: 1.4rem;
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
}
</style>