<template>
  <main class="container mx-auto h-full flex flex-col items-center">
      <div class="sm:w-5/5 md:w-4/5 lg:w-3/5 xl:w-3/5 flex flex-col items-start justify-middle m-4 text-black">
          <h1 class="my-2 font-normal text-4xl">Upload Dive Data</h1>
          <p class="my-2">
              The data you upload will be securely stored in our cloud database.
              As we collect more data, we will begin to release tools for the
              scientific and academic communities to access the database for their research.
          </p>
          <button
            @click="goToUploadPage"
            class="bg-blue hover:bg-blue-dark text-white font-normal my-2 py-2 px-4 rounded">
              Upload File
          </button>
          <p class="my-2 text-sm">
              XML files from MacDive and Subsurface are supported.
          </p>
      </div>
  </main>
</template>
<script>
export default {
    methods: {
        goToUploadPage () {
            this.$router.push({
                name: 'upload'
            });
        }
    }
};
</script>
