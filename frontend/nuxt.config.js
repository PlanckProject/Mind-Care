
export default {
  /*
  ** Nuxt target
  ** See https://nuxtjs.org/api/configuration-target
  */
  target: 'server',
  /*
  ** Headers of the page
  ** See https://nuxtjs.org/api/configuration-head
  */
  head: {
    title: process.env.npm_package_name || '',
    meta: [
      { charset: 'utf-8' },
      { name: 'viewport', content: 'width=device-width, initial-scale=1' },
      { hid: 'description', name: 'description', content: process.env.npm_package_description || '' }
    ],
    link: [
      { rel: 'icon', type: 'image/x-icon', href: '/leaf.ico' }
    ]
  },
  /*
  ** Global CSS
  */
  css: [
    '~assets/styles/light.scss',
    '~assets/styles/global.scss'
  ],
  /*
  ** Plugins to load before mounting the App
  ** https://nuxtjs.org/guide/plugins
  */
  plugins: [
    '~/plugins/vuelidate.js'
  ],
  /*
  ** Auto import components
  ** See https://nuxtjs.org/api/configuration-components
  */
  components: true,
  /*
  ** Nuxt.js dev-modules
  */
  buildModules: [
    '@nuxtjs/fontawesome',
  ],
  /*
  ** Nuxt.js modules
  */
  modules: [
  ],
  /*
  ** Build configuration
  ** See https://nuxtjs.org/api/configuration-build/
  */
  build: {
  },

  generate: {
    subFolders: false
  },
  // Custom modules
  fontawesome: {
    icons: {
      solid: ['faSignInAlt', 'faSun', 'faMoon', 'faCircle', 'faMapPin', 'faGlobe', 'faPhone', 'faEnvelope', 'faClock', 'faMoneyBill', 'faDiagnoses', 'faLaptopMedical'],
    }
  },

  privateRuntimeConfig: {
    apiUri: process.env.API_SERVER_URI || 'http://localhost:8082'
  },

}
