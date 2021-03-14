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
        <p>
          No matching service providers were found near your current location.
          This doesn't look like an issue to us, but the non-availability of
          services around your current location. Please get in touch with the
          team at <a href="mailto:wb.res@outlook.com">wb.res@outlook.com</a> for
          registering new service providers around you or look for an online
          service providers by clicking <a href="/">here</a>.
        </p>
        <br />
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

    const queries = {
      online: "true",
    };
    if (Object.entries(context.route.query).length != 0) {
      let { lat, lon, offline } = context.route.query;
      if (offline) queries.online = "false";
      if (lat && lon) {
        queries.online = "false";
        queries.loc = "true";
        queries.lat = lat;
        queries.lon = lon;
      }
    }

    requestUrl += `?${Object.keys(queries)
      .map(
        (key) =>
          `${encodeURIComponent(key)}=${encodeURIComponent(queries[key])}`
      )
      .join("&")}`;

    console.log(requestUrl);
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
    if (this.$route.query.offline == "true") {
      const existingQuery = this.$route.query;
      if (!(existingQuery.lat && existingQuery.lon)) {
        navigator.geolocation.getCurrentPosition(
          ({ coords }) => {
            existingQuery.lat = coords.latitude;
            existingQuery.lon = coords.longitude;
            window.location =
              this.$route.path +
              `?${Object.keys(existingQuery)
                .map(
                  (key) =>
                    `${encodeURIComponent(key)}=${encodeURIComponent(
                      existingQuery[key]
                    )}`
                )
                .join("&")}`;
          },
          (err) => {
            if (err.code == 1) {
              alert("Error: Access is denied!");
            } else if (err.code == 2) {
              alert("Error: Position is unavailable!");
            } else {
              this.$router.push("/error");
            }
          }
        );
      }
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
