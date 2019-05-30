<template>
    <div class="relative h-full flex flex-col">
        <NavBar v-if="isAuthorized" />
        <div class="flex-auto relative z-0">
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
