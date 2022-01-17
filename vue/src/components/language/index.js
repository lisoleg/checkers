import { createI18n } from "vue-i18n";

const i18n = createI18n({
  locale: localStorage.getItem("lang") ||'ch',
  fallbackLocale: 'ch',
  globalInjection:true,
  legacy: false,
  messages: {
    'ch': require("./ch.js"),
    'en': require("./en.js"),
  },
});

export default i18n;
