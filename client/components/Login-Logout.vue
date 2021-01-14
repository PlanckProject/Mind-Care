<template>
  <nuxt-link to="/login">
    <div>
      <span><font-awesome-icon :icon="['fas', 'sign-in-alt']" /></span>
      <span class="action-text">Login</span>
    </div>
  </nuxt-link>
</template>
<style lang="scss" scoped>
@import "~assets/styles/light.scss";
div {
  * {
    color: inherit;
  }
  position: absolute;
  background: var(--accent-color);
  top: 0.75rem;
  right: 0.75rem;
  padding: 1rem;
  border-radius: 50%;
  box-shadow: 0 0 0.3rem 0;
  text-decoration: none;
  outline: none !important;
  transition: filter 0.2s;
  &:hover {
      filter: grayscale(35%);
  }
}
.action-text {
  display: none;
  transition: all 0.3s linear;
}

@media only screen and (min-width: 40rem) {
  div {
    border-radius: 2rem;
    box-shadow: 0 0 0.4rem 0;
  }
  .action-text {
    display: inline;
  }
}
</style>