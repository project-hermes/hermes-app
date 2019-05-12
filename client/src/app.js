import Vue from 'vue';
import router from '~/router';
import store from '~/store';
import '~/directives';
import App from './App.vue';
import './firebase';
import {RecycleScroller} from 'vue-virtual-scroller';
import 'vue-virtual-scroller/dist/vue-virtual-scroller.css';

Vue.component('RecycleScroller', RecycleScroller);
import 'leaflet/dist/leaflet.css';
new Vue({
    router,
    store,
    render: createEle => createEle(App)
}).$mount('#app');
