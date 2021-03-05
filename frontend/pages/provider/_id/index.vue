<template>
  <main>
    <site-header />
    <theme-switcher />
    <div class="generic-card">
      <h1>{{ provider.name }}</h1>
      <div class="services">
        <span
          v-for="service in provider.services"
          :key="`${service}-${provider.id}`"
          >{{ service }}</span
        >
      </div>
      <p v-if="provider.description">{{ provider.description }}</p>
      <div class="card-details">
        <div>
          <div v-if="provider.contact.contact_person">
            <font-awesome-icon icon="diagnoses" />
            {{ provider.contact.contact_person }}
          </div>
          <div>
            <font-awesome-icon icon="money-bill" />
            {{ provider.fee_range }} {{ isFeeNegotiable }}
          </div>
          <div v-if="provider.timings">
            <font-awesome-icon icon="clock" />
            {{ provider.timings }}
          </div>
        </div>

        <div>
          <div>
            <template v-if="provider.online">
              <font-awesome-icon icon="map-pin" /> Online
            </template>
            <template v-else>
              <a
                target="_blank"
                :href="googleMapsLocation"
                style="word-wrap: break-word"
              >
                <font-awesome-icon icon="map-pin" />
                {{ provider.contact.address.street_address_1 }},
                {{ provider.contact.address.street_address_2 }},
                {{ provider.contact.address.city }}
              </a>
            </template>
          </div>
          <div v-if="provider.contact.number">
            <a :href="contactNumberHref" target="_blank">
              <font-awesome-icon icon="phone" />
              {{ provider.contact.number }}
            </a>
          </div>
          <div v-if="provider.contact.email">
            <a :href="emailHref" target="_blank">
              <font-awesome-icon icon="envelope" />
              {{ provider.contact.email }}
            </a>
          </div>
          <div v-if="provider.contact.website">
            <a :href="provider.contact.website" target="_blank">
              <font-awesome-icon icon="globe" />
              {{ provider.contact.website }}
            </a>
          </div>
        </div>
      </div>
    </div>
  </main>
</template>

<style lang="scss" scoped>
.generic-card {
  padding: 1rem 2rem;
  h1 {
    margin: 1.5rem auto;
  }

  h2 {
    color: var(--accent-color);
    margin: 1.5rem auto;
  }
  .card-details {
    text-align: left;
  }
  .card-details > div {
    margin: 3rem auto;

    &:last-of-type {
      margin-bottom: 1.5rem;
    }

    div {
      margin: 0.75rem auto;
    }
  }
}
@media screen and (min-width: 550px) {
  .generic-card {
    padding: 1rem 3rem;
    h1 {
      margin: 1.5rem auto;
    }

    h2 {
      color: var(--accent-color);
      margin: 1.5rem auto;
    }
    .card-details > div {
      div {
        text-align: center;
        flex-grow: 1;
        flex-basis: 25%;
      }
      display: flex;
      justify-content: space-between;
      align-items: center;
      margin: 0.5rem auto;

      &:last-of-type {
        margin-bottom: 2rem;
      }

      div {
        margin: 0.5rem auto;
      }
    }
  }
}
</style>

<script>
import axios from "axios";

export default {
  async asyncData(context) {
    const provider = await axios
      .get(
        context.$config.apiUri + `/service_provider/${context.route.params.id}`
      )
      .then((res) => res.data.data)
      .catch(() => context.redirect("/error"));

    return { provider };
  },
  beforeMount() {
    if (!this.$store.getters["theme/customThemeSet"])
      this.$store.dispatch("theme/initTheme");
  },
  computed: {
    isFeeNegotiable() {
      const feeNegotiable = this.provider.fee_negotiable.toLowerCase();
      switch (feeNegotiable) {
        case "yes":
          return "(Negotiable)";
        case "maybe":
          return "(Maybe negotiable)";
        default:
          return "";
      }
    },
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
    }
  },
  methods: {
    getServicesAsTags() {
      return this.provider.services.join(", ")
    }
  },
  head() {
    return {
      bodyAttrs: {
        class: this.$store.getters["theme/theme"],
      },
      title: `${this.provider.name} - Mind Care`,
      meta: [
        { charset: "utf-8" },
        { name: "viewport", content: "width=device-width, initial-scale=1" },
        {
          name: `description`,
          content: this.getServicesAsTags(),
        },
      ]
    };
  },
};
</script>