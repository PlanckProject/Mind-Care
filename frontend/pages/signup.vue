<template>
  <main>
    <theme-switcher class="light-switcher-up" />
    <div class="signup" :class="{ 'reject-animation': isErrorHappened }">
      <div
        class="banner"
        :style="{ background: bannerColor, 'z-index': zValue }"
        v-if="!isSubmitted && (isErrorHappened || isInfo)"
      >
        <h4>{{ bannerMessage }}</h4>
      </div>
      <form id="form" :class="{ form: !isSubmitted }">
        <input
          type="text"
          name="full-name"
          id="full-name"
          placeholder="Full Name"
          v-model.trim.lazy="full_name"
          autocomplete="off"
          @blur="$v.full_name.$touch()"
          :class="{ 'invalid-field': $v.full_name.$error }"
          :disabled="isSubmitted"
        />
        <input
          type="email"
          name="email"
          id="email"
          placeholder="Email"
          v-model.trim.lazy="email"
          autocomplete="off"
          @blur="$v.email.$touch()"
          :class="{ 'invalid-field': $v.email.$error }"
          :disabled="isSubmitted"
        />
        <input
          type="text"
          name="username"
          id="username"
          placeholder="Username"
          v-model.trim.lazy="username"
          autocomplete="off"
          @blur="$v.username.$touch()"
          :class="{ 'invalid-field': $v.username.$error }"
          :disabled="isSubmitted"
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
        <input
          type="password"
          name="confirm_password"
          id="password"
          placeholder="Confirm Password"
          v-model.trim.lazy="confirm_password"
          @blur="$v.confirm_password.$touch()"
          :class="{ 'invalid-field': $v.confirm_password.$error }"
          :disabled="isSubmitted"
        />
        <div id="custom-db__wrapper">
          <input type="checkbox" id="custom-db__checkbox" v-model="use_custom_db" :disabled="isSubmitted"/>
          <label id="custom-db__label" for="custom-db__checkbox" >I want to use my own DB</label>
          <input
            type="text"
            v-if="use_custom_db"
            name="db_connection_string"
            id="db_connection_string"
            placeholder="MongoDB connection string"
            @blur="$v.db_connection_string.$touch()"
            v-model.trim.lazy="db_connection_string"
            :class="{ 'invalid-field': $v.db_connection_string.$error }"
            :disabled="isSubmitted"
          />
        </div>
        <div id="submit-wrapper" :class="{ 'form-submitted': isSubmitted }">
          <input
            type="submit"
            id="submit"
            :value="submitMessage"
            :disabled="
              $v.full_name.$invalid ||
              $v.email.$invalid ||
              $v.username.$invalid ||
              $v.password.$invalid ||
              $v.confirm_password.$invalid ||
              $v.db_connection_string.$invalid ||
              isSubmitted
            "
            @click.prevent="signup()"
          />
        </div>
      </form>
      <nuxt-link to="/login" >Login here</nuxt-link>
    </div>
  </main>
</template>

<script>
import {
  required,
  alphaNum,
  minLength,
  email,
  sameAs,
  helpers,
} from "vuelidate/lib/validators";
import axios from "axios";

const dbConnectionStringValidator = (value, vm) => {
  if (vm.$data.use_custom_db) {
    return vm.$data.db_connection_string.length > 0;
  }
  return true;
};

export default {
  name: "signup",
  data() {
    return {
      currentHour: new Date().getHours(),
      isErrorHappened: false,
      isInfo: false,
      bannerColor: "red",
      bannerMessage: "Failed",
      username: "",
      password: "",
      confirm_password: "",
      full_name: "",
      email: "",
      use_custom_db: false,
      db_connection_string: "",
      isSubmitted: false,
      zValue: -10,
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
  computed: {
    submitMessage() {
      return this.isSubmitted ? "" : "Submit";
    },
  },
  methods: {
    reset() {
      this.isSubmitted = true;
      this.isErrorHappened = false;
      this.isInfo = false;
      this.zValue = -10;
    },
    clearForm() {
      this.username = "";
      this.email = "";
      this.password = "";
      this.confirm_password = "";
      this.full_name = "";
      this.use_custom_db = false;
      this.db_connection_string = "";
      this.$v.full_name.$reset();
      this.$v.email.$reset();
      this.$v.username.$reset();
      this.$v.password.$reset();
      this.$v.confirm_password.$reset();
      this.$v.db_connection_string.$reset();
    },
    signup() {
      this.reset();
      let postData = {
        username: this.username,
        password: this.password,
        email: this.email,
        name: this.full_name,
      };
      if (this.use_custom_db)
        postData["db_connection_string"] = this.db_connection_string;

      axios.post("admin/registration").then(
        (response) => {
          this.bannerMessage = response.ok
            ? "Success. Contact admin for approval"
            : response.data.status.message;
          this.zValue = 10;
          this.bannerColor = "green";
          this.isSubmitted = false;
          this.isInfo = true;
          this.clearForm();
        },
        (error) => {
          setTimeout(() => {
            this.zValue = 10;
          this.bannerMessage = error.body
            ? error.body.status.message
            : "Ugh! that's embarrasing, Sorry!";
          this.isErrorHappened = true;
          this.bannerColor = "red";
          this.isSubmitted = false;
          this.isInfo = false;
          }, 10000)
        }
      );
    },
  },
  watch: {
    use_custom_db(value) {
      if (!value) this.$v.db_connection_string.$reset();
    },
  },
  validations: {
    full_name: {
      required,
    },
    email: {
      required,
      email,
    },
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
    confirm_password: {
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
      sameAs: sameAs("password"),
    },
    db_connection_string: {
      validate: dbConnectionStringValidator,
    },
  },
};
</script>

<style lang="scss" scoped>
.signup {
  background: var(--bg-color);
  position: relative;
  max-width: 30rem;
  margin: 6rem auto 0 auto;
  box-shadow: 0 0 0.4rem 0.1rem var(--shadow-color);
  border-radius: 0.4rem;
  text-align: center;
  form {
    padding: 3rem 2rem 0 2rem;
    display: flex;
    flex-direction: column;
    * {
      margin: 0.5rem;
      padding: 0.5rem;
    }
    img {
      margin: 0 auto 1rem auto;
      max-width: 7rem;
      max-height: 7rem;
    }
    #username,
    #password,
    #full-name,
    #email,
    #db_connection_string {
      border-radius: 0.2rem;
      border: 1px solid var(--shadow-color);
      background: inherit;
      color: inherit;
    }
    #custom-db__wrapper {
      #custom-db__checkbox,
      #custom-db__label{
        cursor: pointer;
      }
      padding: 0;
      #db_connection_string {
        width: 100%;
        margin: 1rem auto;
      }
      label {
        margin: 0;
      }
    }
    #submit-wrapper {
      position: relative;
      margin-top: 0;
      #submit {
        margin: auto;
        margin-top: 0;
        border-radius: 1.5rem;
        color: var(--text-color);
        height: 3rem;
        width: 10rem;
        background: transparent;
        border: 2px solid var(--accent-color);
        outline: none !important;
      }
      &:after {
        content: "";
        z-index: 100;
        position: absolute;
        transition: all 0.3s ease-out;
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
    z-index: 2;
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
.form-submitted:after {
  right: 11.7rem !important;
  bottom: 1.6rem !important;
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
.success {
  background: green !important;
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