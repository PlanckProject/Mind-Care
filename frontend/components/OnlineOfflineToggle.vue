<template>
  <div @click="toggleOnlineQuery()">
    <span><font-awesome-icon :style="{'color': isOnline}" icon="laptop-medical" /></span>
  </div>
</template>
<script>
export default {
  computed: {
    isOnline() {
      return (this.$route.query.online == 'true')? 'red': 'var(--text-color)';
    }
  },
  methods: {
    toggleOnlineQuery() {
      if(this.$route.query.online == 'true') {
        window.location = this.$route.path
      } else {
        window.location = this.$route.path + "?online=true"
      }
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
  line-height: 1.5rem;
  width: 1.5rem;
  height: 1.5rem;
  position: absolute;
  background: var(--accent-color);
  top: 2.5rem;
  right: 0.5rem;
  padding: 0rem 0.5rem;
  border-radius: 50%;
  box-shadow: 0 0 0.15rem 0;
  text-decoration: none;
  transition: all 0.2s;
  cursor: pointer;
  font-size: 0.5rem;

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
  top: 5rem;
  right: 1rem;
  padding: 0.75rem 0.875rem;
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