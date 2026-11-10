import Vue from 'vue'
import Toast from 'vue-toastification'
// Import the toast styles!
import "vue-toastification/dist/index.css";

Vue.use(Toast, {
  hideProgressBar: true,
  closeOnClick: false,
  closeButton: false,
  icon: false,
  timeout: 1000,
  transition: 'Vue-Toastification__fade',
})
