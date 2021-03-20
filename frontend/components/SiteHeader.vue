<template>
  <div class="h">
    <NuxtLink to="/"><h1>Mind Care</h1></NuxtLink>
    <NuxtLink to="/register">
      <div
        class="register-btn"
        @mouseover.capture="fullOpen = true"
        @mouseenter.capture="fullOpen = true"
        @mouseleave.capture="fullOpen = false"
        @mouseout.capture="fullOpen = false"
      >
        <span class="register-msg-sm">
          Register a mental health service provider
        </span>
        <span class="register-msg-md">
          <transition name="fade">
            <span v-if="fullOpen">Register!&nbsp;&nbsp;</span>
          </transition>
          <font-awesome-icon icon="arrow-right" />
        </span>
      </div>
    </NuxtLink>
  </div>
</template>

<script>
export default {
  data() {
    return {
      fullOpen: false,
    };
  },
};
</script>

<style lang="scss">
.h {
  position: relative;
  background-image: url("~assets/img/leaf.png");
  background-size: 105vw;
  margin-top: 1.4rem;
  padding: 1rem;
  border-radius: 0 0 0.5rem 0.5rem;
  text-align: center;
  box-shadow: 0px 0px 4px 4px var(--shadow-color);

  h1 {
    cursor: pointer;
  }
}
.register-msg-sm {
  display: block;
}
.register-msg-md {
  display: none;
}
.register-btn {
  font-size: 0.9rem;
  padding: 0.2rem;
  position: absolute;
  top: -1.45rem;
  right: 0;
  width: 100%;
  background: var(--accent-color);
  display: flex;
  align-items: center;
  justify-content: center;
  align-content: center;
  color: #fff;
  cursor: pointer;
  box-shadow: 0px -8px 4px 4px var(--shadow-color);
}

@media screen and (min-width: 550px) {
  .h {
    background-size: 45rem;
    margin-top: 0;
  }
  .register-btn {
    padding: 0;
    top: 0;
    width: 6rem;
    height: 100%;
    border-radius: 0 0 0.5rem 0;
    box-shadow: none;
    transition: all 250ms;
    &:hover {
      width: 10rem;
    }
  }
  .register-msg-sm {
    display: none;
  }
  .register-msg-md {
    display: inline;
  }
}
</style>