<template>
  <main>
    <site-header />
    <theme-switcher />
    <online-offline-toggle />
    <provider-carousel :carouselData="providers" />
  </main>
</template>
<script>
import axios from "axios";

export default {
  transition: "fade",
  name: "index",
  async asyncData(context) {
    let requestUrl = context.$config.apiUri + "/service_providers";
    if (Object.entries(context.route.query).length != 0) {
      let { lat, lon, online } = context.route.query;
      if (online) requestUrl += "?online=true";
      else if (lat && lon) requestUrl += `?loc=true&lat=${lat}&lon=${lon}`;
    }

    let providers = await axios
      .get(requestUrl)
      .then((res) => res.data.data)
      .catch((err) => {
        if (context.isServer) {
          console.error(err);
        }
        context.redirect("/error");
      });

    return { providers };
  },
  mounted() {
    if (Object.entries(this.$route.query).length == 0) {
      navigator.geolocation.getCurrentPosition(
        ({ coords }) => {
          window.location =
            window.location + `?lat=${coords.latitude}&lon=${coords.longitude}`;
        },
        () => {}
      );
    }
  },
  beforeMount() {
    if (!this.$store.getters["theme/customThemeSet"])
      this.$store.dispatch("theme/initTheme");
  },
  head() {
    return {
      bodyAttrs: {
        class: this.$store.getters["theme/theme"],
      },
      title: "Mind Care",
      meta: [
        { charset: "utf-8" },
        { name: "viewport", content: "width=device-width, initial-scale=1" },
        {
          name: `description`,
          content: "Mind Care is a non-profit opensource initiative to helps people catch hold of their mental health.",
        },
      ]
    };
  },
};
</script>
