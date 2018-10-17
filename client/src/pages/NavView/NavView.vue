<template>
  <div
    class="relative">
    <NavBar
      v-if="isAuthorized"/>
    <div class="relative">
        <router-view />
    </div>
  </div>
</template>

<script>
import {NavBar} from '~/components';
import {mapGetters} from 'vuex';

export default {
    components: {
        NavBar
    },
    computed: {
        ...mapGetters({
            isAuthorized: 'auth/isAuthorized'
        })
    }
};
</script>
