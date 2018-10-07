<template>
  <nav class="pin-t pin-l pin-r z-10 flex items-center justify-between bg-blue p-2">
    <img
      :src="navLogo"
      class="ml-1 mr-6 h-8 w-8 mr-4"
      alt="Project Hermes logo">
    <UserButton class="mr-1 relative" />
</nav>
</template>
<script>
import {mapActions, mapGetters} from 'vuex';
import {UserButton} from '~/components';
import navLogo from '~/img/logo-h.svg';
export default {
    components: {
        UserButton
    },
    data() {
        return {
            navLogo
        };
    }
};
</script>
