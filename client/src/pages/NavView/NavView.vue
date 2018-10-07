<template>
  <div
    class="relative">
    <NavBar
      v-if="isAuthorized"/>
    <div class="relative">
      <transition name="fade">
        <router-view />
      </transition>
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
<style>
.fade-enter-active, .fade-leave-active {
  transition: opacity 150ms;
}
.fade-enter, .fade-leave-to {
  opacity: 0;
}
.fade-enter-active {
  transition-delay: 150ms;
}
</style>
