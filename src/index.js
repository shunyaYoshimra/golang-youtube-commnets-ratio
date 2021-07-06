import Vue from "vue";
import App from "./App.vue";
import '../node_modules/materialize-css/dist/css/materialize.min.css';
import 'materialize-css/dist/js/materialize.min';

new Vue({
  el: "#app",
  render: h => h(App)
})