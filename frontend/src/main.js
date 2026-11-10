import Vue from "vue"
import App from "@/App.vue"
import store from "@/store"
import router from "@/router"
import VueRouter from "vue-router"

// 3rd Party Plugins
import '@/libs/toastification'

// Import VeeValidate rules
import {extend, localize} from 'vee-validate';
import * as rules from 'vee-validate/dist/rules';
import en from 'vee-validate/dist/locale/en.json';

// Import Styles
import "@/assets/css/styles.css"

// Import global components
import "@/global-components.js"

// Install VeeValidate rules
Object.keys(rules).forEach(rule => {
  extend(rule, rules[rule]);
});
// Add localisation which provides the messages for default vee-validate.
localize({en});

Vue.config.productionTip = false;

Vue.use(VueRouter)

new Vue({
  router,
  store,
  render: h => h(App),
}).$mount('#app');