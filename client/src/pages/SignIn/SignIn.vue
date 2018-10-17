<template>
  <main
    class="sign-in w-screen h-full bg-cover flex flex-col items-center justify-between">
    <div class="flex-1 flex flex-col items-center justify-end">
      <img
        :src="hermesLogo"
        class="pb-10"
        alt="Project Hermes logo">
    </div>
    <div class="flex-2 py-16 px-20 flex flex-col items-center justify-center border-b border-t border-solid border-transparent-white">
      <GoogleButton @click.native="signInWithGoogle" />
    </div>
    <div class="flex-1">
      <div
        v-if="error"
        class="bg-red-lightest border border-red-light text-red-dark px-4 py-3 my-6 rounded relative"
        role="alert">
        <strong class="font-bold">Oops.</strong>
        <span class="block sm:inline">There was an error signing in.</span>
      </div>
    </div>
  </main>
</template>
<script>
import {mapGetters, mapActions} from 'vuex';
import GoogleButton from './GoogleButton.vue';
import hermesLogo from '~/img/logo-full.svg';

export default {
    components: {
        GoogleButton
    },
    data() {
        return {
            error: '',
            hermesLogo
        };
    },
    computed: {
        ...mapGetters({
            isReady: 'auth/isReady'
        })
    },
    methods: {
        ...mapActions({
            authSignInWithGoogle: 'auth/signInWithGoogle'
        }),
        signInWithGoogle() {
            this.authSignInWithGoogle().catch(err => {
                this.error = err.message;
                throw new Error (err.message);
            });
            this.resetError();
        },
        resetError() {
            this.error = '';
        }
    }
};
</script>
<style scoped>
.sign-in {
    background-image: url('../../img/ocean-bg.jpg');
}
</style>
