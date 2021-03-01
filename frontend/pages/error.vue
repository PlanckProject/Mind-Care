<template>
  <main>
    <site-header />
    <div class="error-d">
        <h3 style="color:red">This looks like an error to us.</h3>
        If this was unexpected, kindly report the error <a href="https://github.com/PlanckProject/Mind-Care" target="_blank">here</a>.
    </div>
  </main>
</template>

<style scoped lang="scss">
.error-d {
    margin-top: 1rem;
    text-align: center;
}
a {
  color: var(--accent-color);
}
</style>

<script>
export default {
  beforeMount() {
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