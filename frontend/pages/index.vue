<template>
  <main>
    <site-header />
    <theme-switcher />
    <online-offline-toggle />
    <template v-if="providers.length">
      <provider-carousel :carouselData="providers" />
    </template>
    <template v-else>
      <div class="no-data">
        <h1>No service providers :(</h1>
        <p>No matching service providers were found near your current location. This doesn't look like an issue to us, but the non-availability of services around your current location. Please get in touch with the team at <a href="mailto:wb.res@outlook.com">wb.res@outlook.com</a> for registering new service providers around you or look for an online service providers by clicking <a href="/?online=true">here</a>.</p><br>
        <span>We wish you great health!</span>
      </div>
    </template>
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
          hid: "description",
          name: "description",
          content:
            "Mind Care is a non-profit open source initiative to help people catch hold of their mental health.",
        },
        {
          hid: "keywords",
          name: "keywords",
          content:
            "Mind, Care, Mindcare, Mental, Health, Depression, Counselling, Psychiatrist, Mental Health, Talking, Wellbeing",
        },
      ],
    };
  },
};
</script>
