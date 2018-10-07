<template>
  <button
    class="bg-white hover:bg-grey-lightest text-grey-darkest font-semibold px-3 py-2 border border-grey-light shadow">
    <div
        class="flex items-center"
        >
        <img
          :src="googleButton"
          class="h-6 w-6 mr-4"
          alt="Google logo" >
        <span
            class="mx-1 font-normal font-sans leading-loose"
            style="color: #575757"
            >
            Sign in with Google
        </span>
    </div>
  </button>
</template>
<script>
import googleButton from '~/img/Google__G__Logo.svg';
export default {
    data() {
        return {
            googleButton
        };
    }
};
</script>
<style>
</style>
