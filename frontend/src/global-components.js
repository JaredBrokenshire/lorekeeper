import Vue from 'vue'
import {ValidationObserver, ValidationProvider} from "vee-validate"
import FeatherIcon from '@core/components/feather-icon/FeatherIcon.vue'

// Feather Icons
Vue.component(FeatherIcon.name, FeatherIcon)

// VeeValidate
Vue.component('ValidationObserver', ValidationObserver);
Vue.component('ValidationProvider', ValidationProvider);
