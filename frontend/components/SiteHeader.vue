<template>
  <div>
    <a href="/"><h1>Mind Care</h1></a>
  </div>
</template>

<style lang="scss" scoped>
div {
  background-image: url('~assets/img/leaf.png');
  background-size: 105vw;
  padding: 1rem;
  border-radius: 0 0 0.5rem 0.5rem;
  text-align: center;
  box-shadow: 0px 0px 4px 4px var(--shadow-color);

  h1 {
    cursor: pointer;
  }
}

@media screen and (min-width: 550px) {
  div {
    background-size: 45rem;
  }
}
</style>