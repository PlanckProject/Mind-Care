export default {
  state() {
    return {
      customTheme: false,
      currentTheme: "light"
    }
  },
  getters: {
    theme: (state) => state.currentTheme,
    customThemeSet: (state) => state.customTheme
  },
  mutations: {
    CHANGE_THEME(state, args) {
      state.currentTheme = args.theme,
      state.customTheme = args.customThemeSet
    }
  },
  actions: {
    toggleTheme(ctx) {
      let theme = {
        theme: ctx.state.currentTheme == "light" ? "dark" : "light",
        customThemeSet: true
      }
      ctx.commit("CHANGE_THEME", theme);
    },
    initTheme(ctx) {
      var hour = new Date().getHours();
      let theme = {
        theme: hour >= 7 && hour < 20 ? "light" : "dark",
        customThemeSet: false
      }
      ctx.commit("CHANGE_THEME", theme);
    }
  }
}