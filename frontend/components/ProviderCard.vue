<template>
  <div class="provider-card">
    <a :href="`/provider/${provider.id}`"
      ><h1>{{ provider.name }}</h1></a
    >

    <div class="services">
      <span
        v-for="service in provider.services"
        :key="`${service}-${provider.id}`"
        >{{ service }}</span
      >
    </div>

    <div class="card_contact">
      <template v-if="provider.online">
        <font-awesome-icon icon="map-pin" /> Online
      </template>
      <template v-else>
        <a target="_blank" :href="googleMapsLocation">
          <font-awesome-icon icon="map-pin" />
          {{ provider.contact.address.city }}
        </a>
      </template>
    </div>
    <div class="card_contact" v-if="provider.contact.number">
      <a :href="contactNumberHref" target="_blank">
        <font-awesome-icon icon="phone" />
        {{ provider.contact.number }}
      </a>
    </div>
    <div class="card_contact" v-if="provider.contact.email">
      <a :href="emailHref" target="_blank">
        <font-awesome-icon icon="envelope" />
        {{ provider.contact.email }}
      </a>
    </div>
    <div class="card_contact" v-if="provider.contact.website">
      <a :href="provider.contact.website" target="_blank">
        <font-awesome-icon icon="globe" />
        {{ provider.contact.website }}
      </a>
    </div>
    <div class="card_contact" v-if="provider.timings">
      <font-awesome-icon icon="clock" />
      {{ provider.timings }}
    </div>
  </div>
</template>

<script>
export default {
  props: {
    provider: {
      type: Object,
      required: true,
    },
  },
  computed: {
    contactNumberHref() {
      return `tel:${this.provider.contact.number}`;
    },
    emailHref() {
      return `mailto:${this.provider.contact.email}`;
    },
    googleMapsLocation() {
      return (
        "https://www.google.com/maps/place/" +
        `${this.provider.contact.address.coordinates[0]}N+${this.provider.contact.address.coordinates[1]}E`
      );
    },
  },
};
</script>

<style lang="scss" scoped>
.provider-card {
  margin: 1rem auto;
  padding: 1rem;
  border-radius: 0.5rem;
  box-shadow: 0px 0px 4px 4px var(--shadow-color);
}

.card_contact {
  margin: 0.5rem;
}
</style>