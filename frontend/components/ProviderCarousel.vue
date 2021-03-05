<template>
  <section>
    <template v-for="provider in carouselData">
      <provider-card
        class="generic-card"
        :provider="provider"
        :key="provider.id"
      />
    </template>
  </section>
</template>

<script>
import ProviderCard from "./ProviderCard.vue";
export default {
  components: { ProviderCard },
  props: {
    carouselData: {
      type: Array,
      required: true,
    },
  },
};
</script>