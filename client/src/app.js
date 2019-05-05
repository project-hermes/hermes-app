import Vue from 'vue';
import router from '~/router';
import store from '~/store';
import '~/directives';
import App from './App.vue';
import './firebase';
new Vue({
    router,
    store,
    render: createEle => createEle(App)
}).$mount('#app');
