<template>
  <main
    :style="{'opacity': opacity}"
    class="opacity-ease-in container mx-auto my-10 h-full flex flex-col items-center">
    <div class="sm:w-5/5 md:w-3/5 lg:w-2/5 xl:w-2/5 flex flex-col items-center justify-middle m-4 text-black">
      <CheckCircleIcon class="text-green h-24 w-24" />
      <p class="py-3">
        You’re all done, thanks for submitting data!
        <br />
        You can close this window or
        <a
          class="hover:underline text-blue cursor-pointer"
          @click="goToUploadPage">upload more files.</a>
      </p>
    </div>
  </main>
</template>
<script>
import CheckCircleIcon from 'vue-feather-icons/icons/CheckCircleIcon';

export default {
    components: {
        CheckCircleIcon
    },
    data () {
        return {
            opacity: 0
        };
    },
    methods: {
        goToUploadPage() {
            this.$router.push({
                name: 'upload'
            });
        }
    },
    created () {
        setTimeout(() => {
            this.opacity = 1;
        }, 100);
    }
};
</script>
<style scoped>
.opacity-ease-in {
    transition: opacity 0.25s ease-in;
}

</style>
