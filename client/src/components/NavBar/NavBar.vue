<template>
  <nav class="h-16 z-10 flex items-center justify-between bg-blue">
    <img
      :src="navLogo"
      class="mx-3 h-8 w-8"
      alt="Project Hermes logo">
  <div class="flex flex-row h-full items-center">
      <NavLink route="/">Dashboard</NavLink>
      <NavLink route="/dives">Dives</NavLink>
      <NavLink route="/sensor">Sensor</NavLink>
      <NavLink route="/upload">Upload</NavLink>
      <UserButton class="mx-1 relative" />
  </div>
</nav>
</template>
<script>
import {mapActions, mapGetters} from 'vuex';
import {UserButton} from '~/components';
import NavLink from './NavLink.vue';
import navLogo from '~/img/logo-h.svg';
export default {
    components: {
        UserButton,
        NavLink
    },
    data() {
        return {
            navLogo
        };
    }
};
</script>
