<template>
  <main>
    <theme-switcher class="light-switcher-up" />
    <div class="login" :class="{ 'reject-animation': isErrorHappened }">
      <div class="banner" v-if="!isSubmitted && isErrorHappened">
        <h3>{{ bannerMessage }}</h3>
      </div>
      <form id="form" :class="{ form: !isSubmitted }">
        <input
          type="text"
          name="username"
          id="username"
          placeholder="Username"
          v-model.trim.lazy="username"
          autocomplete="off"
          @blur="$v.username.$touch()"
          :class="{ 'invalid-field': $v.username.$error }"
        />
        <input
          type="password"
          name="password"
          id="password"
          placeholder="Password"
          v-model.trim.lazy="password"
          @blur="$v.password.$touch()"
          :class="{ 'invalid-field': $v.password.$error }"
        />
        <div id="submit-wrapper" :class="{ 'form-submitted': isSubmitted }">
          <input
            type="submit"
            id="submit"
            :value="submitMessage"
            :disabled="
              $v.username.$invalid || $v.password.$invalid || isSubmitted
            "
            @click.prevent="login()"
          />
        </div>
        <div v-if="showFloatingLogin" class="signup-link">
          <nuxt-link to="/signup">Signup here</nuxt-link>
        </div>
      </form>
    </div>
  </main>
</template>

<style lang="scss" scoped>
main {
  padding-top: 5rem;
  .login {
    background: inherit;
    position: relative;
    width: 95%;
    max-width: 30rem;
    margin: 2.5% auto 2.5% auto;
    box-shadow: 0 0 0.4rem 0.1rem var(--shadow-color);
    border-radius: 0.4rem;
    text-align: center;
    .banner {
      position: absolute;
      z-index: 3;
    }
    form {
      padding: 3rem 2rem 0 2rem;
      display: flex;
      flex-direction: column;
      background: inherit;
      border-radius: 0.2rem;
      * {
        margin: 1rem;
        padding: 0.5rem;
      }
      img {
        margin: 0 auto 1rem auto;
        max-width: 7rem;
        max-height: 7rem;
      }
      #username,
      #password {
        background: inherit;
        color: inherit;
        border: 1px solid var(--shadow-color);
        border-radius: 0.2rem;
      }
      #submit-wrapper {
        position: relative;
        outline: none !important;
        #submit {
          margin: 1rem auto;
          color: var(--text-color);
          border-radius: 1.5rem;
          height: 3rem;
          width: 10rem;
          background: transparent;
          border: 2px solid var(--accent-color);
          outline: none !important;
        }
        &:after {
          content: "";
          position: absolute;
          z-index: 100;
        }
      }
      .signup-link {
        margin: 0;
        padding: 0;
        a {
          padding: 0;
          margin-top: 0;
        }
      }
    }
    > div {
      z-index: -20;
      padding: 0.25rem;
      position: absolute;
      width: 100%;
      top: -3rem;
      border-radius: 0.4rem 0.4rem 0 0;
      background: red;
      color: white;
      transform: translateY(3rem);
    }
    a {
      display: inline-block;
      margin-bottom: 2rem;
      text-decoration: none;
      color: rgb(0, 0, 180);
    }
  }
}
.form-submitted:after {
  outline: none;
  right: 11.2rem !important;
  bottom: 2.6rem !important;
  border: 0.2rem solid grey !important;
  border-top: 0.2rem solid var(--accent-color) !important;
  border-radius: 50%;
  width: 1.4rem !important;
  height: 1.4rem !important;
  animation: spin 0.5s linear infinite;
}
.invalid-field {
  border: 1px solid rgba(255, 0, 0, 0.5) !important;
}
.reject-animation {
  animation: reject 0.5s ease-in-out forwards;
}
@keyframes spin {
  0% {
    transform: rotate(0deg);
  }
  100% {
    transform: rotate(360deg);
  }
}
@keyframes reject {
  0% {
    transform: rotateY(0deg);
  }
  33% {
    transform: rotateY(20deg);
  }
  66% {
    transform: rotateY(-20deg);
  }
}
</style>
<script>
import axios from "axios";
import { required, alphaNum, minLength } from "vuelidate/lib/validators";
import Cookie from "js-cookie";
import moment from "moment";
import { validationMixin } from "vuelidate";

export default {
  data() {
    return {
      currentHour: new Date().getHours(),
      isErrorHappened: false,
      bannerMessage: "",
      username: "",
      password: "",
      isSubmitted: false,
    };
  },
  beforeMount() {
    if (!this.$store.getters.customThemeSet) this.$store.dispatch("initTheme");
  },
  head() {
    return {
      bodyAttrs: {
        class: this.$store.getters.theme,
      },
    };
  },
  methods: {
    login() {
      this.isSubmitted = true;
      this.isErrorHappened = false;
      axios
        .post("admin/login", {
          username: this.username,
          password: this.password,
        })
        .then(
          (response) => {
            const lease_time = response.data.data.lease_time;
            Cookie.set("username", this.username, {
              expires: (1 / 86400) * lease_time,
            });
            Cookie.set("isLoggedIn", "true", {
              expires: (1 / 86400) * lease_time,
            });
            Cookie.set(
              "expires_at",
              moment().add(lease_time, "second").format("x"),
              {
                expires: (1 / 86400) * lease_time,
              }
            );
            Cookie.set("token", response.data.data.token, {
              expires: (1 / 86400) * lease_time,
            });
            this.$store.dispatch("refreshLoginStatus", false);
            this.$store.dispatch("confirmLogin", {
              username: this.username,
              expires_at: moment().add(lease_time, "second").format("x"),
              token: response.data.data.token,
            });
            this.username = "";
            this.password = "";
            this.$v.username.$reset();
            this.$v.password.$reset();
            this.isSubmitted = false;
            this.$router.push({
              name: "dashboard",
            });
          },
          (error) => {
            this.bannerMessage = error.body
              ? error.body.status.message
              : "Something went horribly wrong";
            this.isErrorHappened = true;
            this.isSubmitted = false;
          }
        );
    },
  },
  computed: {
    showFloatingLogin() {
      return true;
    },
    submitMessage() {
      return this.isSubmitted ? "" : "Submit";
    },
  },
  validations: {
    username: {
      required,
      minLen: minLength(5),
      alphaNum,
    },
    password: {
      required,
      minLen: minLength(8),
      isStrong: (value) => {
        if (value === "") {
          return true;
        }
        return new Promise((resolve, reject) => {
          resolve(
            /^(?=.*[A-Z])(?=.*[!@#$&*])(?=.*[0-9])(?=.*[a-z]).*$/.test(value)
          );
        });
      },
    },
  },
};
</script>