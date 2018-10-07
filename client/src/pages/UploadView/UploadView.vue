<template>
  <main class="container mx-auto h-full flex flex-col items-center">
    <div class="sm:w-5/5 md:w-4/5 lg:w-3/5 xl:w-3/5 flex flex-col items-start justify-middle m-4 text-black">
      <h1 class="my-2 font-normal text-4xl">Upload Dive Data</h1>
      <p class="my-2">
        The data you upload will be securely stored in our cloud database.
        As we collect more data, we will begin to release tools for the
        scientific and academic communities to access the database for their research.
      </p>
      <Uploader
        class="uploader-min-height w-full h-full"
        @confirm="onConfirmation"
        :progress="progress" />
      <p class="my-2 text-sm">
        XML files from MacDive and Subcurrent are supported.
      </p>
    </div>
  </main>
</template>
<script>
import UploadIcon from 'vue-feather-icons/icons/UploadIcon';
import CheckIcon from 'vue-feather-icons/icons/CheckIcon';
import head from 'lodash/head';
import {mapGetters, mapActions, mapMutations} from 'vuex';
import {Uploader} from '~/components';

export default {
    components: {
        UploadIcon,
        CheckIcon,
        Uploader
    },
    computed: {
        ...mapGetters({
            progress: 'upload/progress',
            status: 'upload/status'
        })
    },
    mounted () {
        this.resetState();
    },
    methods: {
        ...mapActions({
            uploadFiles: 'upload/uploadFiles'
        }),
        ...mapMutations({
            resetState: 'upload/resetState'
        }),
        onConfirmation(files) {
            this.uploadFiles({
                files
            }).then(() => {
                this;
                console.log(this);
                this.$nextTick(() => {
                    this.$router.push({
                        name: 'thanks'
                    });
                });
                // setTimeout(() => {
                //     this.$router.push({
                //         name: 'thanks'
                //     });
                // }, 1000);
            });
        }
    }
};
</script>
<style>
.uploader-min-height {
    min-height: 12rem;
    height: auto;
}
</style>
